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
