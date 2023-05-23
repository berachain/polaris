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

package mint

import (
	"context"
	"math/big"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/mint"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// Contract is the precompile contract for the mint module.
type Contract struct {
	ethprecompile.BaseContract
	querier minttypes.QueryServer
}

// NewPrecompileContract returns a new instance of the mint module precompile contract.
func NewPrecompileContract(
	q minttypes.QueryServer,
) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.MintModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(minttypes.ModuleName)),
		),
		querier: q,
	}
}

// CustomValueDecoders overrides the `coreprecompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{}
}

// PrecompileMethods implements the `coreprecompile.StatefulImpl` interface.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "annualProvisions()",
			Execute: c.AnnualProvisions,
		},
		{
			AbiSig:  "inflation()",
			Execute: c.Inflation,
		},
	}
}

// SetWithdrawAddress is the precompile contract method for the `setWithdrawAddress(address)` method.
func (c *Contract) AnnualProvisions(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	res, err := c.querier.AnnualProvisions(ctx, &minttypes.QueryAnnualProvisionsRequest{})
	if err != nil {
		return nil, err
	}
	return []any{res.AnnualProvisions.BigInt()}, err
}

// SetWithdrawAddressBech32 is the precompile contract method for the `setWithdrawAddress(string)` method.
func (c *Contract) Inflation(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	res, err := c.querier.Inflation(ctx, &minttypes.QueryInflationRequest{})
	if err != nil {
		return nil, err
	}
	return []any{res.Inflation.BigInt()}, err
}
