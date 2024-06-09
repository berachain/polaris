// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
