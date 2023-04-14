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

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// convertCoinToERC20 converts SDK/Polaris coins to ERC20 tokens for an owner.
func (c *Contract) convertCoinToERC20(
	ctx context.Context,
	caller common.Address,
	evm ethprecompile.EVM,
	value *big.Int,
	denom string,
	owner common.Address,
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

	var token common.Address
	if resp.Token == "" {
		// first occurrence of an IBC originated SDK coin, must be created as a Polaris ERC20 token

		// deploy the new ERC20 token contract (deployer of this contract is the ERC20 module!)
		if token, err = c.deployPolarisERC20Contract(sdkCtx, evm, c.RegistryKey(), denom, value); err != nil {
			return err
		}

		// create the new ERC20 token contract pairing with SDK coin denomination
		c.em.RegisterCoinERC20Pair(sdkCtx, denom, token)
	} else {
		// subsequent occurence of an IBC-originated SDK coin OR an ERC20 originated token's
		// Polaris coin counterpart

		// convert ERC20 token bech32 address to common.Address
		var tokenAcc sdk.AccAddress
		if tokenAcc, err = sdk.AccAddressFromBech32(resp.Token); err != nil {
			return err
		}
		token = cosmlib.AccAddressToEthAddress(tokenAcc)
	}

	if erc20types.IsPolarisDenom(denom) {
		// converting Polaris coins to ERC20 originated tokens
		// NOTE: it is guaranteed that the ERC20 tokens were transferred to the ERC20 module
		// precompile contract as escrow before this case is reached.

		// transfer amount ERC20 tokens to the owner
		if err = c.callERC20Transfer(sdkCtx, evm, c.RegistryKey(), token, owner, amount); err != nil {
			return err
		}
	} else {
		// converting IBC-originated SDK coins to Polaris ERC20 tokens

		// mint amount ERC20 tokens to the owner
		if err = c.callPolarisERC20Mint(sdkCtx, evm, c.RegistryKey(), token, owner, amount); err != nil {
			return err
		}
	}

	// burn amount SDK/Polaris coins from owner
	if err := cosmlib.BurnCoinsFromAddress(
		sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount,
	); err != nil {
		return err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertCoinToERC20,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, denom),
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+denom),
		),
	)
	return nil
}

// convertERC20ToCoin converts ERC20 tokens to SDK/Polaris coins for an owner.
func (c *Contract) convertERC20ToCoin(
	ctx context.Context,
	caller common.Address,
	evm ethprecompile.EVM,
	token common.Address,
	owner common.Address,
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
	if resp.Denom == "" {
		// if denomination not found, create new pair with ERC20 token <> Polaris coin denomination
		resp.Denom = c.em.RegisterERC20CoinPair(sdkCtx, token)
	}

	if erc20types.IsPolarisDenom(resp.Denom) {
		// converting ERC20 originated tokens to Polaris coins
		// NOTE: owner must approve caller to spend amount ERC20 tokens

		// return an error if the ERC20 token contract does not exist to revert the tx
		if !evm.GetStateDB().Exist(token) {
			return errors.New("ERC20 token contract does not exist")
		}

		// caller transfers amount ERC20 tokens from owner to ERC20 module precompile contract in
		// escrow
		if err := c.callERC20TransferFrom(sdkCtx, evm, caller, token, owner, c.RegistryKey(), amount); err != nil {
			return err
		}
	} else {
		// converting Polaris ERC20 tokens to IBC-originated SDK coins

		// burn amount ERC20 tokens from owner
		if err := c.callPolarisERC20Burn(sdkCtx, evm, c.RegistryKey(), token, owner, amount); err != nil {
			return err
		}
	}

	// mint amount SDK/Polaris Coins to owner
	if err = cosmlib.MintCoinsToAddress(sdkCtx, c.bk, erc20types.ModuleName, owner, resp.Denom, amount); err != nil {
		return err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertERC20ToCoin,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, resp.Denom),
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()+resp.Denom),
		),
	)
	return nil
}
