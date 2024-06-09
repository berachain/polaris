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

package testapp

import (
	evmconfig "github.com/berachain/polaris/cosmos/config"
	bankprecompile "github.com/berachain/polaris/cosmos/precompile/bank"
	distrprecompile "github.com/berachain/polaris/cosmos/precompile/distribution"
	govprecompile "github.com/berachain/polaris/cosmos/precompile/governance"
	stakingprecompile "github.com/berachain/polaris/cosmos/precompile/staking"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

// PrecompilesToInject returns a function that provides the initialization of the standard
// set of precompiles.
func PrecompilesToInject(
	app *SimApp,
	customPcs ...ethprecompile.Registrable,
) func() *ethprecompile.Injector {
	return func() *ethprecompile.Injector {
		// Create the precompile injector with the standard precompiles.
		pcs := ethprecompile.NewPrecompiles([]ethprecompile.Registrable{
			bankprecompile.NewPrecompileContract(
				app.AccountKeeper,
				bankkeeper.NewMsgServerImpl(app.BankKeeper),
				app.BankKeeper,
			),
			distrprecompile.NewPrecompileContract(
				app.AccountKeeper,
				app.StakingKeeper,
				distrkeeper.NewMsgServerImpl(app.DistrKeeper),
				distrkeeper.NewQuerier(app.DistrKeeper),
			),
			govprecompile.NewPrecompileContract(
				app.AccountKeeper,
				govkeeper.NewMsgServerImpl(app.GovKeeper),
				govkeeper.NewQueryServer(app.GovKeeper),
				app.interfaceRegistry,
			),
			stakingprecompile.NewPrecompileContract(app.AccountKeeper, app.StakingKeeper),
		}...)

		// Add the custom precompiles to the injector.
		for _, pc := range customPcs {
			pcs.AddPrecompile(pc)
		}
		return pcs
	}
}

// PrecompilesToInject returns a function that provides the initialization of the standard
// set of precompiles.
func QueryContextFn(app *SimApp) func() func(height int64, prove bool) (sdk.Context, error) {
	return func() func(height int64, prove bool) (sdk.Context, error) {
		return app.BaseApp.CreateQueryContext
	}
}

// PolarisConfigFn returns a function that provides the initialization of the standard
// set of precompiles.
func PolarisConfigFn(cfg *evmconfig.Config) func() *evmconfig.Config {
	return func() *evmconfig.Config {
		return cfg
	}
}
