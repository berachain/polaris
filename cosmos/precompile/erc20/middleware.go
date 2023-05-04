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

	pbindings "pkg.berachain.dev/polaris/contracts/bindings/polaris"
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

// ErrTokenDoesNotExist is returned when a token contract does not exist.
var ErrTokenDoesNotExist = errors.New("ERC20 token contract does not exist")

// convertCoinToERC20 converts SDK/Polaris coins to ERC20 tokens for an owner.
func (c *Contract) convertCoinToERC20(
	ctx context.Context,
	evm ethprecompile.EVM,
	value *big.Int,
	denom string,
	owner common.Address,
	recipient common.Address,
	amount *big.Int,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get ERC20 token address pairing with SDK/Polaris coin denomination
	resp, err := c.em.ERC20AddressForCoinDenom(
		ctx, &erc20types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		},
	)
	if err != nil {
		return err
	}

	// burn amount SDK/Polaris coins from owner
	if err = cosmlib.BurnCoinsFromAddress(sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount); err != nil {
		return err
	}

	var token common.Address
	if resp.Token == "" {
		// first occurrence of an IBC originated SDK coin, must be created as a Polaris ERC20 token

		// deploy the new ERC20 token contract (deployer of this contract is the ERC20 module!)
		polarisName := p + denom
		if token, _, err = cosmlib.DeployEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), c.polarisERC20ABI, value,
			pbindings.PolarisERC20MetaData.Bin, polarisName, polarisName,
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

	if erc20types.IsPolarisDenom(denom) {
		// converting Polaris coins to ERC20 originated tokens
		// NOTE: it is guaranteed that the ERC20 tokens were transferred to the ERC20 module
		// precompile contract as escrow before this case is reached.

		// transfer amount ERC20 tokens to the recipient
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, big.NewInt(0),
			transfer, recipient, amount,
		); err != nil {
			return err
		}
	} else {
		// converting IBC-originated SDK coins to Polaris ERC20 tokens

		// mint amount ERC20 tokens to the recipient
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, big.NewInt(0),
			mint, recipient, amount,
		); err != nil {
			return err
		}
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertCoinToERC20,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, denom),
			sdk.NewAttribute(erc20types.AttributeKeyOwner, owner.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyRecipient, recipient.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+denom),
		),
	)
	return nil
}

// convertERC20ToCoin converts ERC20 tokens to SDK/Polaris coins for an owner.
func (c *Contract) convertERC20ToCoin( //nolint:funlen // ok.
	ctx context.Context,
	caller common.Address,
	evm ethprecompile.EVM,
	token common.Address,
	owner common.Address,
	recipient common.Address,
	amount *big.Int,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get SDK/Polaris coin denomination pairing with ERC20 token
	resp, err := c.em.CoinDenomForERC20Address(
		ctx, &erc20types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.AddressToAccAddress(token).String(),
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

	if erc20types.IsPolarisDenom(denom) { //nolint:nestif // necessary cases.
		// converting ERC20 originated tokens to Polaris coins
		// NOTE: owner must approve caller to spend amount ERC20 tokens

		// return an error if the ERC20 token contract does not exist to revert the tx
		if !evm.GetStateDB().Exist(token) {
			return ErrTokenDoesNotExist
		}

		var (
			ret           []any
			balanceBefore *big.Int
			balanceAfter  *big.Int
		)

		// check the ERC20 module's balance of the ERC20-originated token
		ret, err = cosmlib.CallEVMFromPrecompileUnpackArgs(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, big.NewInt(0),
			balanceOf, c.RegistryKey(),
		)
		if err != nil {
			return err
		}
		balanceBefore = utils.MustGetAs[*big.Int](ret[0])

		// caller transfers amount ERC20 tokens from owner to ERC20 module precompile contract in
		// escrow
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			caller, token, c.polarisERC20ABI, big.NewInt(0),
			transferFrom, owner, c.RegistryKey(), amount,
		); err != nil {
			return err
		}

		// check the ERC20 module's balance of the ERC20-originated token after the transfer
		ret, err = cosmlib.CallEVMFromPrecompileUnpackArgs(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, big.NewInt(0),
			balanceOf, c.RegistryKey(),
		)
		if err != nil {
			return err
		}
		balanceAfter = utils.MustGetAs[*big.Int](ret[0])

		// set the amount of Polaris coins to mint as the delta of the ERC20 module's balance
		amount = new(big.Int).Sub(balanceAfter, balanceBefore)
	} else {
		// converting Polaris ERC20 tokens to IBC-originated SDK coins

		// burn amount ERC20 tokens from owner
		if _, err = cosmlib.CallEVMFromPrecompile(
			sdkCtx, c.GetPlugin(), evm,
			c.RegistryKey(), token, c.polarisERC20ABI, big.NewInt(0),
			burn, owner, amount,
		); err != nil {
			return err
		}
	}

	// mint amount SDK/Polaris Coins to recipient
	if err = cosmlib.MintCoinsToAddress(sdkCtx, c.bk, erc20types.ModuleName, recipient, denom, amount); err != nil {
		return err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertERC20ToCoin,
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyOwner, owner.Hex()),
			sdk.NewAttribute(erc20types.AttributeKeyRecipient, recipient.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+denom),
		),
	)
	return nil
}
