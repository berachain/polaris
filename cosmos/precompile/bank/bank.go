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
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
		{
			AbiSig:  "getAllBalance(address)",
			Execute: c.GetAllBalance,
		},
		{
			AbiSig:  "getSpendableBalanceByDenom(address,string)",
			Execute: c.GetSpendableBalanceByDenom,
		},
		{
			AbiSig:  "getSpendableBalances(address)",
			Execute: c.GetSpendableBalances,
		},
		{
			AbiSig:  "getSupplyOf(string)",
			Execute: c.GetSupplyOf,
		},
		{
			AbiSig:  "getTotalSupply()",
			Execute: c.GetTotalSupply,
		},
		{
			AbiSig:  "getParams()",
			Execute: c.GetParams,
		},
		{
			AbiSig:  "getDenomMetadata(string)",
			Execute: c.GetDenomMetadata,
		},
		{
			AbiSig:  "getDenomsMetadata()",
			Execute: c.GetDenomsMetadata,
		},
		{
			AbiSig:  "getSendEnabled(string[])",
			Execute: c.GetSendEnabled,
		},
		{
			AbiSig:  "send(address,address,(uint256,string)[])",
			Execute: c.Send,
		},
		{
			AbiSig:  "multiSend((address,(uint256,string)[]),(address,(uint256,string)[])[])",
			Execute: c.MultiSend,
		},
	}
}

// grpc_query functions
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

	fmt.Printf("\nin bank: \nget balance: addr: %v\n", cosmlib.AddressToAccAddress(addr).String())

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
func (c *Contract) GetAllBalance(
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

	// todo: add pagination here
	res, err := c.querier.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
		Address: cosmlib.AddressToAccAddress(addr).String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{sdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSpendableBalanceByDenom implements `getSpendableBalanceByDenom(address,string)` method.
func (c *Contract) GetSpendableBalanceByDenom(
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

	res, err := c.querier.SpendableBalanceByDenom(ctx, &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: cosmlib.AddressToAccAddress(addr).String(),
		Denom:   denom,
	})
	if err != nil {
		return nil, err
	}

	balance := res.GetBalance().Amount
	return []any{balance.BigInt()}, nil
}

// GetSpendableBalances implements `getSpendableBalances(address)` method.
func (c *Contract) GetSpendableBalances(
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

	res, err := c.querier.SpendableBalances(ctx, &banktypes.QuerySpendableBalancesRequest{
		Address: cosmlib.AddressToAccAddress(addr).String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{sdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSupplyOf implements `GetSupplyOf(string)` method.
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

// GetTotalSupply implements `getTotalSupply()` method.
func (c *Contract) GetTotalSupply(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	// todo: add pagination here
	res, err := c.querier.TotalSupply(ctx, &banktypes.QueryTotalSupplyRequest{})
	if err != nil {
		return nil, err
	}

	return []any{sdkCoinsToEvmCoins(res.Supply)}, nil
}

// GetParams implements `getParams()` method.
func (c *Contract) GetParams(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	res, err := c.querier.Params(ctx, &banktypes.QueryParamsRequest{})

	if err != nil {
		return nil, err
	}

	// note: res.Params.SendEnabled is deprecated
	return []any{res.Params}, nil
}

// GetDenomMetadata implements `getDenomMetadata(string)` method.
func (c *Contract) GetDenomMetadata(
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

	res, err := c.querier.DenomMetadata(ctx, &banktypes.QueryDenomMetadataRequest{
		Denom: denom,
	})
	if err != nil {
		return nil, err
	}

	return []any{res.Metadata}, nil
}

// GetDenomsMetadata implements `getDenomsMetadata()` method.
func (c *Contract) GetDenomsMetadata(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	// todo: add pagination here
	res, err := c.querier.DenomsMetadata(ctx, &banktypes.QueryDenomsMetadataRequest{})
	if err != nil {
		return nil, err
	}

	return []any{res.Metadatas}, nil
}

// todo: this function without pagination is a bad idea
// func (c *Contract) GetDenomsOwners(
// 	ctx context.Context,
// 	_ ethprecompile.EVM,
// 	_ common.Address,
// 	_ *big.Int,
// 	readonly bool,
// 	args ...any,
// ) ([]any, error) {
// 	res, err := c.querier.DenomOwners(ctx, &banktypes.QueryDenomOwnersRequest{
// 		Denom: "",
// 		Pagination: nil,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return []any{res.DenomOwners}, nil
// }

// GetSendEnabled implements `getSendEnabled(string[])` method.
func (c *Contract) GetSendEnabled(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	denoms, ok := utils.GetAs[[]string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.SendEnabled(ctx, &banktypes.QuerySendEnabledRequest{
		Denoms: denoms,
	})
	if err != nil {
		return nil, err
	}

	// todo: test if "return []any{res.SendEnabled}, nil" works
	// here we are dereferencing the values for safety
	sendEnableds := make([]banktypes.SendEnabled, len(res.SendEnabled))
	for i, p := range res.SendEnabled {
		sendEnableds[i] = *p
	}

	return []any{sendEnableds}, nil
}

// msg_server functions
// Send implements `send(address,address,(uint256,string))` method.
func (c *Contract) Send(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	fromAddr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	toAddr, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	fmt.Printf("\nin bank: args[2] is: %v\n", args[2])

	coins, err := extractCoinsFromInput(args[2])
	fmt.Printf("\nin bank: amount: %v\n", coins)

	fmt.Printf("\nin bank: fromAddr: %v\n", cosmlib.AddressToAccAddress(fromAddr).String())
	fmt.Printf("\nin bank: toAddr: %v\n", cosmlib.AddressToAccAddress(toAddr).String())

	if err != nil {
		return nil, err
	}

	_, err = c.msgServer.Send(ctx, &banktypes.MsgSend{
		FromAddress: cosmlib.AddressToAccAddress(fromAddr).String(),
		ToAddress:   cosmlib.AddressToAccAddress(toAddr).String(),
		Amount:      coins,
	})
	return []any{err == nil}, err
}

// MultiSend implements `multiSend((address,(uint256,string)[]),(address,(uint256,string)[])[])` method.
func (c *Contract) MultiSend(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	evmInput, ok := utils.GetAs[generated.IBankModuleBalance](args[0])
	if !ok {
		return nil, precompile.ErrInvalidAny
	}
	evmOutputs, ok := utils.GetAs[[]generated.IBankModuleBalance](args[1])
	if !ok {
		return nil, precompile.ErrInvalidAny
	}

	totalOutputCoins := sdk.NewCoins()

	// input params for c.msgServer.MultiSend
	sdkInputs := make([]banktypes.Input, 1)
	sdkOutputs := make([]banktypes.Output, len(evmOutputs))

	inputSdkCoins := sdk.NewCoins()
	for _, coin := range evmInput.Coins {
		inputSdkCoins = append(inputSdkCoins, sdk.NewCoin(coin.Denom, sdk.NewIntFromBigInt(coin.Amount)))
	}

	sdkInputs[0] = banktypes.NewInput(
		cosmlib.AddressToAccAddress(evmInput.Addr),
		inputSdkCoins,
	)

	for i, evmOutput := range evmOutputs {
		sdkCoins := sdk.NewCoins()
		for _, coin := range evmOutput.Coins {
			sdkCoins = append(sdkCoins, sdk.NewCoin(coin.Denom, sdk.NewIntFromBigInt(coin.Amount)))
		}

		totalOutputCoins = totalOutputCoins.Add(sdkCoins...)

		sdkOutputs[i] = banktypes.NewOutput(
			cosmlib.AddressToAccAddress(evmOutput.Addr),
			sdkCoins,
		)
	}

	// Check input amount and total amounts for outputs are equal
	if !inputSdkCoins.Equal(totalOutputCoins) {
		return nil, precompile.ErrInvalidAny
	}

	_, err := c.msgServer.MultiSend(ctx, &banktypes.MsgMultiSend{
		Inputs:  sdkInputs,
		Outputs: sdkOutputs,
	})
	return []any{err == nil}, err
}
