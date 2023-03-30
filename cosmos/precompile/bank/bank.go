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

package bank

import (
	"context"
	"math/big"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the bank module.
type Contract struct {
	precompile.BaseContract

	msgServer banktypes.MsgServer
	querier   banktypes.QueryServer
}

// NewPrecompileContract returns a new instance of the bank precompile contract.
func NewPrecompileContract(bk bankkeeper.Keeper) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			generated.BankModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
		),
		msgServer: bankkeeper.NewMsgServerImpl(bk),
		querier:   bk,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "getBalance(address,string)",
			Execute: c.GetBalance,
		},
		// {
		// 	AbiSig:  "getAllBalance(address)",
		// 	Execute: c.GetAllBalance,
		// },
		{
			AbiSig:  "getSupplyOf(string)",
			Execute: c.GetSupplyOf,
		},
	}
}

// GetBalance implements `getBalance(address,string)` method.
func (c *Contract) GetBalance(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	denom, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: cosmlib.AddressToAccAddress(addr).String(),
		Denom:   denom,
	})
	if err != nil {
		return nil, err
	}

	balance := res.GetBalance().Amount
	return []any{balance.BigInt()}, nil
}

// // GetAllBalance implements `getAllBalance(address)` method.
// func (c *Contract) GetAllBalance(
// 	ctx context.Context,
// 	_ ethprecompile.EVM,
// 	_ common.Address,
// 	_ *big.Int,
// 	readonly bool,
// 	args ...any,
// ) ([]any, error) {
// 	addr, ok := utils.GetAs[common.Address](args[0])
// 	if !ok {
// 		return nil, precompile.ErrInvalidHexAddress
// 	}

// 	// AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)

// 	// type QueryAllBalancesRequest struct {
// 	// 	// address is the address to query balances for.
// 	// 	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
// 	// 	// pagination defines an optional pagination for the request.
// 	// 	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
// 	// 	// resolve_denom is the flag to resolve the denom into a human-readable form from the metadata.
// 	// 	//
// 	// 	// Since: cosmos-sdk 0.48
// 	// 	ResolveDenom bool `protobuf:"varint,3,opt,name=resolve_denom,json=resolveDenom,proto3" json:"resolve_denom,omitempty"`
// 	// }

// 	res, err := c.querier.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
// 		Address: cosmlib.AddressToAccAddress(addr).String(),
// 	})
// 	if status.Code(err) == codes.NotFound {
// 		// handle the case where the delegation does not exist
// 		return []any{big.NewInt(0)}, nil
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	// type QueryAllBalancesResponse struct {
// 	// 	// balances is the balances of all the coins.
// 	// 	Balances github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances"`
// 	// 	// pagination defines the pagination in the response.
// 	// 	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
// 	// }

// 	// ask: ?????????????????????????????????
// 	// res.Balances
// 	// res.Pagination

//		return []any{[]Coin}, nil
//	}
func (c *Contract) GetSupplyOf(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.SupplyOf(ctx, &banktypes.QuerySupplyOfRequest{
		Denom: denom,
	})
	if err != nil {
		return nil, err
	}

	supply := res.GetAmount().Amount
	return []any{supply.BigInt()}, nil
}
