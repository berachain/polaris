// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package erc20

import (
	"context"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abi "github.com/ethereum/go-ethereum/accounts/abi"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

const (
	p            = `p`
	balanceOf    = `balanceOf`
	transfer     = `transfer`
	transferFrom = `transferFrom`
	mint         = `mint`
	burn         = `burn`
)

var (
	// ErrTokenDoesNotExist is returned when a token contract does not exist.
	ErrTokenDoesNotExist = errors.New("ERC20 token contract does not exist")
	// ErrInvalidAmount is returned when an amount is invalid.
	ErrInvalidAmount = errors.New("amount is negative or 0")
)

// transferCoinToERC20 transfers SDK/Polaris coins to ERC20 tokens for an owner.
//

func (c *Contract) transferCoinToERC20(
	ctx context.Context,
	evm ethprecompile.EVM,
	value *big.Int,
	denom string,
	owner common.Address,
	recipient common.Address,
	amount *big.Int,
) error {
	if amount.Cmp(common.Big0) <= 0 {
		return ErrInvalidAmount
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	resp, err := c.em.ERC20AddressForCoinDenom(
		sdkCtx, &erc20types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		},
	)
	if err != nil {
		return err
	}

	// burn the incoming SDK/Polaris coins from owner
	if err = cosmlib.BurnCoinsFromAddress(sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount); err != nil {
		return err
	}

	// determine the outgoing (Polaris)ERC20 token
	var token common.Address
	if resp.Token == "" {
		// first occurrence of an IBC originated SDK coin, must be created as a Polaris ERC20 token

		// deploy the new ERC20 token contract (deployer of this contract is the ERC20 module!)
		if token, _, err = cosmlib.DeployOnEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), c.polarisERC20ABI, value,
			c.polarisERC20Bin, p+denom,
		); err != nil {
			return err
		}

		// create the new ERC20 token contract pairing with SDK coin denomination
		c.em.RegisterCoinERC20Pair(sdkCtx, denom, token)
	} else {
		// subsequent occurrence of an IBC-originated SDK coin OR an ERC20 originated token's
		// Polaris coin counterpart

		// convert ERC20 token bech32 address to common.Address
		var tokenAcc sdk.AccAddress
		if tokenAcc, err = sdk.AccAddressFromBech32(resp.Token); err != nil {
			return err
		}
		token = cosmlib.AccAddressToEthAddress(tokenAcc)

		// return an error if the ERC20 token contract does not exist to revert the tx
		if !evm.GetStateDB().Exist(token) {
			return ErrTokenDoesNotExist
		}
	}

	// handle the outgoing (Polaris)ERC20 tokens
	if erc20types.IsPolarisDenom(denom) {
		// converting Polaris coins to ERC20 originated tokens
		// NOTE: it is guaranteed that the ERC20 tokens were transferred to the ERC20 module
		// precompile contract as escrow before this case is reached.

		// transfer amount ERC20 tokens to the recipient
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, common.Big0,
			transfer, recipient, amount,
		); err != nil {
			return err
		}
	} else {
		// converting IBC-originated SDK coins to Polaris ERC20 tokens

		// mint amount PolarisERC20 tokens to the recipient
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, common.Big0,
			mint, recipient, amount,
		); err != nil {
			return err
		}
	}

	// emit an event at the end of this successful transfer
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeTransferCoinToERC20,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, denom),
			sdk.NewAttribute(erc20types.AttributeKeyOwner, owner.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyRecipient, recipient.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+denom),
		),
	)
	return nil
}

// transferERC20ToCoin transfers ERC20 tokens to SDK/Polaris coins for an owner.
func (c *Contract) transferERC20ToCoin(
	ctx context.Context,
	_ common.Address,
	evm ethprecompile.EVM,
	token common.Address,
	owner common.Address,
	recipient common.Address,
	amount *big.Int,
) error {
	if amount.Cmp(common.Big0) <= 0 {
		return ErrInvalidAmount
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get SDK/Polaris coin denomination pairing with ERC20 token
	resp, err := c.em.CoinDenomForERC20Address(
		ctx, &erc20types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.Bech32FromEthAddress(token),
		},
	)
	if err != nil {
		return err
	}

	denom := resp.Denom
	if denom == "" {
		// if denomination not found, create new pair with ERC20 token <> Polaris coin denomination
		denom = c.em.RegisterERC20CoinPair(sdkCtx, token)
	}

	if erc20types.IsPolarisDenom(denom) { //nolint:nestif // readability.
		// transferring ERC20 originated tokens to Polaris coins

		// return an error if the ERC20 token contract does not exist to revert the tx
		if !evm.GetStateDB().Exist(token) {
			return ErrTokenDoesNotExist
		}

		var (
			balanceBefore *big.Int
			balanceAfter  *big.Int
			plugin        = c.GetPlugin()
			erc20Module   = c.RegistryKey()
		)

		// check the ERC20 module's balance of the ERC20-originated token
		if balanceBefore, err = getBalanceOf(
			sdkCtx, plugin, evm, erc20Module, token, c.polarisERC20ABI, erc20Module,
		); err != nil {
			return err
		}

		// caller transfers amount ERC20 tokens from owner to ERC20 module in escrow
		// NOTE: owner must have previously approved the ERC20 Module to spend amount ERC20 tokens
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, plugin, evm,
			erc20Module, token, c.polarisERC20ABI, common.Big0,
			transferFrom, owner, erc20Module, amount,
		); err != nil {
			return err
		}

		// check the ERC20 module's balance of the ERC20-originated token
		if balanceAfter, err = getBalanceOf(
			sdkCtx, plugin, evm, erc20Module, token, c.polarisERC20ABI, erc20Module,
		); err != nil {
			return err
		}

		amount = new(big.Int).Sub(balanceAfter, balanceBefore)
	} else {
		// transferring Polaris ERC20 tokens to IBC-originated SDK coins

		// burn amount ERC20 tokens from owner
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, common.Big0,
			burn, owner, amount,
		); err != nil {
			return err
		}
	}

	// mint amount SDK/Polaris Coins to recipient
	if err = cosmlib.MintCoinsToAddress(sdkCtx, c.bk, erc20types.ModuleName, recipient, denom, amount); err != nil {
		return err
	}

	// emit an event at the end of this successful transfer
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeTransferERC20ToCoin,
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyOwner, owner.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyRecipient, recipient.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+denom),
		),
	)
	return nil
}

// getBalanceOf returns the balanceOf `address` for a ERC20 token at `contractAddr`.
func getBalanceOf(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm ethprecompile.EVM,
	caller common.Address,
	contractAddr common.Address,
	contract abi.ABI,
	address common.Address,
) (*big.Int, error) {
	ret, err := cosmlib.StaticCallEVMFromPrecompileUnpackArgs(
		ctx, plugin, evm,
		caller, contractAddr, contract,
		balanceOf, address,
	)
	if err != nil {
		return nil, err
	}
	return utils.MustGetAs[*big.Int](ret[0]), nil
}
