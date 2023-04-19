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

package perc20

import (
	"context"
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	pbindings "pkg.berachain.dev/polaris/contracts/bindings/polaris/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

// bankDenomPrefix is the prefix for the bank denom of the Polaris ERC20 tokens.
const bankDenomPrefix = `perc20/`

// Contract is the PolarisERC20 precompiled contract implementation. Adheres to both the ERC-20 and
// ERC-2612 standards.
type Contract struct {
	precompile.BaseContract

	bk        bankkeeper.Keeper
	bankDenom string

	name   string
	symbol string
}

// NewPrecompileContract returns a new instance of the PolarisERC20 precompiled contract with the
// given name, symbol, and endowment.
func NewPrecompileContract(
	ctx sdk.Context,
	bk bankkeeper.Keeper,
	name, symbol string,
	endowment *big.Int,
) (ethprecompile.DynamicImpl, common.Address, error) {
	address := common.HexToAddress(name) // TODO: use hash with nonce.
	pc := &Contract{
		BaseContract: precompile.NewBaseContract(pbindings.PolarisERC20MetaData.ABI, address),
		bk:           bk,
		bankDenom:    bankDenomPrefix + address.Hex(),
		name:         name,
		symbol:       symbol,
	}

	// mint endowment coins to the PolarisERC20 contract address.
	if endowment.Cmp(big.NewInt(0)) > 0 {
		if err := cosmlib.MintCoinsToAddress(ctx, bk, erc20types.ModuleName, address, pc.bankDenom, endowment); err != nil {
			return nil, common.Address{}, errors.Wrap(err, "failed to deploy PolarisERC20")
		}
	}

	return pc, address, nil
}

func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "name()",
			Execute: c.ERC20Name,
		},
		{
			AbiSig:  "symbol()",
			Execute: c.Symbol,
		},
		{
			AbiSig:  "decimals()",
			Execute: c.Decimals,
		},
		{
			AbiSig:  "totalSupply()",
			Execute: c.TotalSupply,
		},
		{
			AbiSig:  "balanceOf(address)",
			Execute: c.BalanceOf,
		},
	}
}

// Name implements DynamicImpl.
func (c *Contract) Name() string {
	return c.name
}

// ERC20Name returns the name of the PolarisERC20 token.
func (c *Contract) ERC20Name(
	_ context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return []any{c.name}, nil
}

// Symbol returns the symbol of the PolarisERC20 token.
func (c *Contract) Symbol(
	_ context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return []any{c.symbol}, nil
}

// The Cosmos SDK implementation of the PolarisERC20 always uses 18 decimals.
func (c *Contract) Decimals(
	_ context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return []any{uint8(sdkmath.LegacyPrecision)}, nil
}

// TotalSupply returns the total supply of the PolarisERC20 token from the bank module.
func (c *Contract) TotalSupply(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return []any{c.bk.GetSupply(sdk.UnwrapSDKContext(ctx), c.bankDenom).Amount.BigInt()}, nil
}

// BalanceOf returns the balance of the PolarisERC20 token from the bank module.
func (c *Contract) BalanceOf(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return []any{
		c.bk.GetBalance(
			sdk.UnwrapSDKContext(ctx),
			cosmlib.AddressToAccAddress(addr),
			c.bankDenom,
		).Amount.BigInt(),
	}, nil
}
