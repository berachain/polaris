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

package evm

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	store "cosmossdk.io/store/types"

	modulev1alpha1 "github.com/berachain/polaris/cosmos/api/polaris/evm/module/v1alpha1"
	"github.com/berachain/polaris/cosmos/config"
	"github.com/berachain/polaris/cosmos/x/evm/keeper"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:gochecknoinits // GRRRR fix later.
func init() {
	appmodule.Register(&modulev1alpha1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	ModuleKey depinject.OwnModuleKey
	Config    *modulev1alpha1.Module
	Key       *store.KVStoreKey

	AppOpts           servertypes.AppOptions
	PolarisCfg        func() *config.Config
	CustomPrecompiles func() *ethprecompile.Injector `optional:"true"`
	QueryContextFn    func() func(height int64, prove bool) (sdk.Context, error)

	AccountKeeper AccountKeeper
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

// ProvideModule is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) DepInjectOutput {
	// Default to empty precompile injector if not provided.
	if in.CustomPrecompiles == nil {
		in.CustomPrecompiles = func() *ethprecompile.Injector { return &ethprecompile.Injector{} }
	}

	k := keeper.NewKeeper(
		in.AccountKeeper,
		in.Key,
		in.CustomPrecompiles,
		in.QueryContextFn,
		in.PolarisCfg(),
	)
	m := NewAppModule(k, in.AccountKeeper)

	return DepInjectOutput{
		Keeper: k,
		Module: m,
	}
}
