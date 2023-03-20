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

package auth

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	coreprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the auth module.
type Contract struct {
	precompile.BaseContract
}

// NewPrecompileContract returns a new instance of the auth module precompile contract.
func NewPrecompileContract() coreprecompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.AuthModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			contractAbi, cosmlib.AccAddressToEthAddress(
				authtypes.NewModuleAddress(authtypes.ModuleName),
			),
		),
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() coreprecompile.Methods {
	return coreprecompile.Methods{
		{
			AbiSig:      "convertHexToBech32(address)",
			Execute:     c.ConvertHexToBech32,
			RequiredGas: params.IdentityBaseGas,
		},
		{
			AbiSig:      "convertBech32ToHexAddress(string)",
			Execute:     c.ConvertBech32ToHexAddress,
			RequiredGas: params.IdentityBaseGas,
		},
	}
}

// ConvertHexToBech32 converts a common.Address to a bech32 string.
func (c *Contract) ConvertHexToBech32(
	ctx context.Context,
	_ coreprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	hexAddr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	// try val address first
	valAddr, err := sdk.ValAddressFromHex(hexAddr.String())
	if err == nil {
		return []any{valAddr.String()}, nil
	}

	// try account address
	accAddr, err := sdk.AccAddressFromHexUnsafe(hexAddr.String())
	if err == nil {
		return []any{accAddr.String()}, nil
	}

	return nil, precompile.ErrInvalidHexAddress
}

// ConvertBech32ToHexAddress converts a bech32 string to a common.Address.
func (c *Contract) ConvertBech32ToHexAddress(
	ctx context.Context,
	_ coreprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	// try account address first
	accAddr, err := sdk.AccAddressFromBech32(bech32Addr)
	if err == nil {
		return []any{cosmlib.AccAddressToEthAddress(accAddr)}, nil
	}

	// try validator address
	valAddr, err := sdk.ValAddressFromBech32(bech32Addr)
	if err == nil {
		return []any{cosmlib.ValAddressToEthAddress(valAddr)}, nil
	}

	return nil, precompile.ErrInvalidBech32Address
}
