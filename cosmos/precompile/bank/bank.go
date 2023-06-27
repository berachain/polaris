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
	"reflect"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the bank module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer banktypes.MsgServer
	querier   banktypes.QueryServer
}

// NewPrecompileContract returns a new instance of the bank precompile contract.
func NewPrecompileContract(ms banktypes.MsgServer, qs banktypes.QueryServer) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.BankModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
		),
		msgServer: ms,
		querier:   qs,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	contractVal := reflect.ValueOf(c)
	return ethprecompile.GeneratePrecompileMethod(c.ABIMethods(), contractVal)
}

// GetBalance implements `getBalance(address,string)` method.
func (c *Contract) GetBalance(
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
	denom, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
		Denom:   denom,
	})
	if err != nil {
		return nil, err
	}

	balance := res.GetBalance().Amount
	return []any{balance.BigInt()}, nil
}

// // GetAllBalances implements `getAllBalances(address)` method.
func (c *Contract) GetAllBalances(
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

	// todo: add pagination here
	res, err := c.querier.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
	})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSpendableBalanceByDenom implements `getSpendableBalanceByDenom(address,string)` method.
func (c *Contract) GetSpendableBalanceByDenom(
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
	denom, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.SpendableBalanceByDenom(ctx, &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
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
	_ bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	res, err := c.querier.SpendableBalances(ctx, &banktypes.QuerySpendableBalancesRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
	})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSupplyOf implements `GetSupplyOf(string)` method.
func (c *Contract) GetSupplyOf(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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
	_ bool,
	_ ...any,
) ([]any, error) {
	// todo: add pagination here
	res, err := c.querier.TotalSupply(ctx, &banktypes.QueryTotalSupplyRequest{})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Supply)}, nil
}

// GetDenomMetadata implements `getDenomMetadata(string)` method.
func (c *Contract) GetDenomMetadata(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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

	denomUnits := make([]generated.IBankModuleDenomUnit, len(res.Metadata.DenomUnits))
	for i, d := range res.Metadata.DenomUnits {
		denomUnits[i] = generated.IBankModuleDenomUnit{
			Denom:    d.Denom,
			Aliases:  d.Aliases,
			Exponent: d.Exponent,
		}
	}

	result := generated.IBankModuleDenomMetadata{
		Description: res.Metadata.Description,
		DenomUnits:  denomUnits,
		Base:        res.Metadata.Base,
		Display:     res.Metadata.Display,
		Name:        res.Metadata.Name,
		Symbol:      res.Metadata.Symbol,
	}
	return []any{result}, nil
}

// GetSendEnabled implements `getSendEnabled(string[])` method.
func (c *Contract) GetSendEnabled(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	res, err := c.querier.SendEnabled(ctx, &banktypes.QuerySendEnabledRequest{
		Denoms: []string{denom},
	})
	if err != nil {
		return nil, err
	}
	if len(res.SendEnabled) == 0 {
		return nil, precompile.ErrInvalidString
	}

	return []any{res.SendEnabled[0].Enabled}, nil
}

// Send implements `send(address,address,(uint256,string))` method.
func (c *Contract) Send(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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
	coins, err := cosmlib.ExtractCoinsFromInput(args[2])

	if err != nil {
		return nil, err
	}

	_, err = c.msgServer.Send(ctx, &banktypes.MsgSend{
		FromAddress: cosmlib.Bech32FromEthAddress(fromAddr),
		ToAddress:   cosmlib.Bech32FromEthAddress(toAddr),
		Amount:      coins,
	})
	return []any{err == nil}, err
}
