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

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// convertCoinToERC20 converts SDK coins to ERC20 tokens for an owner.
func (c *Contract) convertCoinToERC20() {}

// convertERC20ToCoin converts ERC20 tokens to SDK coins for an owner.
func (c *Contract) convertERC20ToCoin(
	ctx context.Context,
	caller common.Address,
	evm ethprecompile.EVM,
	value *big.Int,
	token common.Address,
	owner common.Address,
	amount *big.Int,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// transfer amount ERC20 tokens from owner to ERC20 module precompile contract
	err := c.transferERC20ToModule(evm, sdkCtx.GasMeter(), caller, owner, value, token, amount)
	if err != nil {
		return err
	}

	// get SDK coin denomination pairing with ERC20 token
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
		// denomination not found, create new pair
		c.em.RegisterDenomTokenPair(sdkCtx, token)
	}

	// mint amount SDK Coins and send to owner
	cosmlib.MintCoinsToAddress(sdkCtx, c.bk, erc20types.ModuleName, owner, denom, amount)

	return nil
}

// transferERC20ToModule transfers ERC20 tokens from an owner to ERC20 module precompile contract
// by calling back into the EVM.
func (c *Contract) transferERC20ToModule(
	evm ethprecompile.EVM,
	gasMeter storetypes.GasMeter,
	caller common.Address,
	owner common.Address,
	value *big.Int,
	token common.Address,
	amount *big.Int,
) error {
	suppliedGas := gasMeter.GasRemaining()

	c.GetPlugin().EnableReentrancy()

	// call ERC20 contract to transferFrom owner to ERC20 module precompile contract
	input, err := c.erc20ABI.Pack("transferFrom", owner, c.RegistryKey(), amount)
	if err != nil {
		return err
	}
	_, gasRemaining, err := evm.Call(
		vm.AccountRef(caller), token, input, suppliedGas, value,
	)

	c.GetPlugin().DisableReentrancy()

	// consume gas used by EVM during ERC20 transfer
	defer gasMeter.ConsumeGas(suppliedGas-gasRemaining, "ERC20 transfer")

	return err
}
