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
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// convertCoinToERC20 converts SDK coins to ERC20 tokens for an owner.
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

	// burn amount SDK coins from owner
	err := cosmlib.BurnCoinsFromAddress(
		sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount,
	)
	if err != nil {
		return err
	}

	// get ERC20 token address pairing with SDK coin denomination
	resp, err := c.em.ERC20AddressForCoinDenom(
		ctx, &erc20types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		},
	)
	if err != nil {
		return err
	}

	var token common.Address
	if resp.Token == "" { //nolint:nestif // okay here.
		// deploy the new ERC20 token contract (deployer of this contract is the ERC20 module!)
		token, err = c.deployPolarisERC20Contract(sdkCtx, evm, c.RegistryKey(), denom, value)
		if err != nil {
			return err
		}

		// create the new ERC20 token contract pairing with SDK coin denomination
		c.em.RegisterCoinERC20Pair(sdkCtx, denom, token)

		// mint amount ERC20 tokens to the owner
		if err = c.callPolarisERC20Mint(sdkCtx, evm, c.RegistryKey(), token, owner, amount); err != nil {
			return err
		}
	} else {
		// convert ERC20 token bech32 address to common.Address
		var tokenAcc sdk.AccAddress
		tokenAcc, err = sdk.AccAddressFromBech32(resp.Token)
		if err != nil {
			return err
		}
		token = cosmlib.AccAddressToEthAddress(tokenAcc)

		// approve the caller to transfer amountERC20 tokens from ERC20 module precompile contract
		if err = c.callERC20Approve(sdkCtx, evm, c.RegistryKey(), token, caller, amount); err != nil {
			return err
		}

		// transfer amount ERC20 tokens from ERC20 module precompile contract to owner
		if err = c.callERC20TransferFrom(sdkCtx, evm, caller, token, c.RegistryKey(), owner, amount); err != nil {
			return err
		}
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertCoinToERC20,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, denom),
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
		),
	)
	return nil
}

// convertERC20ToCoin converts ERC20 tokens to SDK coins for an owner.
func (c *Contract) convertERC20ToCoin(
	ctx context.Context,
	caller common.Address,
	evm ethprecompile.EVM,
	token common.Address,
	owner common.Address,
	amount *big.Int,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// approve the caller to transfer amount ERC20 tokens from owner
	if err := c.callERC20Approve(sdkCtx, evm, owner, token, caller, amount); err != nil {
		return err
	}

	// transfer amount ERC20 tokens from owner to ERC20 module precompile contract
	if err := c.callERC20TransferFrom(sdkCtx, evm, caller, token, owner, c.RegistryKey(), amount); err != nil {
		return err
	}

	// get SDK coin denomination pairing with ERC20 token
	tokenString := cosmlib.AddressToAccAddress(token).String()
	resp, err := c.em.CoinDenomForERC20Address(
		ctx, &erc20types.CoinDenomForERC20AddressRequest{Token: tokenString},
	)
	if err != nil {
		return err
	}

	// denomination not found, create new pair
	denom := resp.Denom
	if denom == "" {
		c.em.RegisterERC20CoinPair(sdkCtx, token)
		// get the newly registered Polaris coin denomination
		resp, err = c.em.CoinDenomForERC20Address(
			ctx, &erc20types.CoinDenomForERC20AddressRequest{Token: tokenString},
		)
		if err != nil {
			return err
		}
		denom = resp.Denom
	}

	// mint amount SDK Coins and send to owner
	if err = cosmlib.MintCoinsToAddress(sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount); err != nil {
		return err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			erc20types.EventTypeConvertERC20ToCoin,
			sdk.NewAttribute(erc20types.AttributeKeyDenom, denom),
			sdk.NewAttribute(erc20types.AttributeKeyToken, token.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
		),
	)
	return nil
}
