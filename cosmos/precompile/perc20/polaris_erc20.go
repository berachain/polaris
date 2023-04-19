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

	pbindings "pkg.berachain.dev/polaris/contracts/bindings/polaris/precompile"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// Contract is the PolarisERC20 precompiled contract implementation. Adheres to both the ERC-20 and
// ERC-2612 standards.
type Contract struct {
	precompile.BaseContract

	name   string
	symbol string
}

// NewPrecompileContract returns a new instance of the PolarisERC20 precompiled contract with the
// given name and symbol.
func NewPrecompileContract(name, symbol string) ethprecompile.DynamicImpl {
	address := common.HexToAddress(name) // TODO: use hash with nonce.
	return &Contract{
		BaseContract: precompile.NewBaseContract(pbindings.PolarisERC20MetaData.ABI, address),
		name:         name,
		symbol:       symbol,
	}
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
