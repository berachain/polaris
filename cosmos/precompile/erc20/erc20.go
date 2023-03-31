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

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the auth module.
type Contract struct {
	precompile.BaseContract

	querier erc20types.QueryServiceServer
}

// NewPrecompileContract returns a new instance of the auth module precompile contract.
func NewPrecompileContract(querier erc20types.QueryServiceServer) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			generated.ERC20ModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(
				authtypes.NewModuleAddress(erc20types.ModuleName),
			),
		),
		querier: querier,
	}
}

// CustomValueDecoders implements StatefulImpl.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		"token": log.ConvertAccAddressFromBech32,
		"denom": log.ConvertString,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "coinDenomForERC20Address(address)",
			Execute: c.CoinDenomForERC20AddressAddrInput,
		},
		{
			AbiSig:  "coinDenomForERC20Address(string)",
			Execute: c.CoinDenomForERC20AddressStringInput,
		},
		{
			AbiSig:  "erc20AddressForCoinDenom(string)",
			Execute: c.ERC20AddressForCoinDenom,
		},
		{
			AbiSig:  "convertCoinToERC20(IERC20,address,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertCoinToERC20(IERC20,string,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertCoinToERC20(string,address,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertCoinToERC20string,string,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertERC20ToCoin(IERC20,address,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertERC20ToCoin(IERC20,string,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertERC20ToCoin(string,address,uint256)",
			Execute: nil,
		},
		{
			AbiSig:  "convertERC20ToCoin(string,string,uint256)",
			Execute: nil,
		},
	}
}

// CoinDenomForERC20AddressAddrInput returns the SDK coin denomination for the given ERC20 address.
func (c *Contract) CoinDenomForERC20AddressAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	resp, err := c.querier.CoinDenomForERC20Address(
		ctx,
		&erc20types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.AddressToAccAddress(addr).String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{resp.Denom}, nil
}

// CoinDenomForERC20AddressStringInput returns the SDK coin denomination for the given ERC20 address.
func (c *Contract) CoinDenomForERC20AddressStringInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	resp, err := c.querier.CoinDenomForERC20Address(
		ctx,
		&erc20types.CoinDenomForERC20AddressRequest{
			Token: addr,
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{resp.Denom}, nil
}

// ERC20AddressForCoinDenom returns the ERC20 address for the given SDK coin denomination.
func (c *Contract) ERC20AddressForCoinDenom(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	resp, err := c.querier.ERC20AddressForCoinDenom(
		ctx,
		&erc20types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{resp.Token}, nil
}
