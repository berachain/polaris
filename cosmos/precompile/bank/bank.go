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
	"reflect"

	"github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
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
	pc ethprecompile.PolarContext,
	addr common.Address,
	denom string,
) ([]any, error) {
	res, err := c.querier.Balance(pc.Ctx, &banktypes.QueryBalanceRequest{
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
	pc ethprecompile.PolarContext,
	addr common.Address,
) ([]any, error) {
	// todo: add pagination here
	res, err := c.querier.AllBalances(pc.Ctx, &banktypes.QueryAllBalancesRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
	})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSpendableBalanceByDenom implements `getSpendableBalanceByDenom(address,string)` method.
func (c *Contract) GetSpendableBalanceByDenom(
	pc ethprecompile.PolarContext,
	addr common.Address,
	denom string,
) ([]any, error) {

	res, err := c.querier.SpendableBalanceByDenom(pc.Ctx, &banktypes.QuerySpendableBalanceByDenomRequest{
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
	pc ethprecompile.PolarContext,
	addr common.Address,
) ([]any, error) {

	res, err := c.querier.SpendableBalances(pc.Ctx, &banktypes.QuerySpendableBalancesRequest{
		Address: cosmlib.Bech32FromEthAddress(addr),
	})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Balances)}, nil
}

// GetSupplyOf implements `GetSupplyOf(string)` method.
func (c *Contract) GetSupplyOf(
	pc ethprecompile.PolarContext,
	denom string,
) ([]any, error) {

	res, err := c.querier.SupplyOf(pc.Ctx, &banktypes.QuerySupplyOfRequest{
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
	pc ethprecompile.PolarContext,
) ([]any, error) {
	// todo: add pagination here
	res, err := c.querier.TotalSupply(pc.Ctx, &banktypes.QueryTotalSupplyRequest{})
	if err != nil {
		return nil, err
	}

	return []any{cosmlib.SdkCoinsToEvmCoins(res.Supply)}, nil
}

// GetDenomMetadata implements `getDenomMetadata(string)` method.
func (c *Contract) GetDenomMetadata(
	pc ethprecompile.PolarContext,
	denom string,
) ([]any, error) {

	res, err := c.querier.DenomMetadata(pc.Ctx, &banktypes.QueryDenomMetadataRequest{
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
	pc ethprecompile.PolarContext,
	denom string,
) ([]any, error) {
	res, err := c.querier.SendEnabled(pc.Ctx, &banktypes.QuerySendEnabledRequest{
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
	pc ethprecompile.PolarContext,
	fromAddr common.Address,
	toAddr common.Address,
	coins types.Coins,
) ([]any, error) {
	_, err := c.msgServer.Send(pc.Ctx, &banktypes.MsgSend{
		FromAddress: cosmlib.Bech32FromEthAddress(fromAddr),
		ToAddress:   cosmlib.Bech32FromEthAddress(toAddr),
		Amount:      coins,
	})
	return []any{err == nil}, err
}
