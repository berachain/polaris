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
	"io"
	"os"
	"path/filepath"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	evmv1alpha1 "pkg.berachain.dev/polaris/cosmos/api/polaris/evm/v1alpha1"
	evmconfig "pkg.berachain.dev/polaris/cosmos/config"
	signinglib "pkg.berachain.dev/polaris/cosmos/lib/signing"
	polarruntime "pkg.berachain.dev/polaris/cosmos/runtime"
	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
)

//nolint:gochecknoinits // from sdk.
func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".polard")
}

// DefaultNodeHome default home directories for the application daemon.
var DefaultNodeHome string

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*runtime.App
	*polarruntime.Polaris
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             *govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	// polaris required keeper
	EVMKeeper *evmkeeper.Keeper
}

// NewPolarisApp returns a reference to an initialized SimApp.
//
//nolint:funlen // from sdk.
func NewPolarisApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	bech32Prefix string,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	var (
		app        = &SimApp{}
		appBuilder *runtime.AppBuilder
		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			MakeAppConfig(bech32Prefix),
			depinject.Provide(
				signinglib.ProvideNoopGetSigners[*evmv1alpha1.WrappedEthereumTransaction],
				signinglib.ProvideNoopGetSigners[*evmv1alpha1.WrappedPayloadEnvelope],
			),
			depinject.Supply(
				// supply the application options
				appOpts,
				// supply the logger
				logger,
				// ADVANCED CONFIGURATION\
				PolarisConfigFn(evmconfig.MustReadConfigFromAppOpts(appOpts)),
				PrecompilesToInject(app),
				QueryContextFn(app),
				//
				// AUTH
				//
				// For providing a custom function required in auth to generate custom account types
				// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
				//
				// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),

				// For providing a custom a base account type add it below.
				// By default the auth module uses authtypes.ProtoBaseAccount().
				//
				// func() sdk.AccountI { return authtypes.ProtoBaseAccount() },

				//
				// MINT
				//

				// For providing a custom inflation function for x/mint add here your
				// custom function that implements the minttypes.InflationCalculationFn
				// interface.
			),
		)
	)

	if err := depinject.Inject(appConfig,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.DistrKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.EvidenceKeeper,
		&app.ConsensusParamsKeeper,
		&app.EVMKeeper,
	); err != nil {
		panic(err)
	}

	// Build the app using the app builder.
	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)
	app.Polaris = polarruntime.New(
		evmconfig.MustReadConfigFromAppOpts(appOpts), app.Logger(), app.EVMKeeper.Host,
	)

	// Setup Polaris Runtime.
	if err := app.Polaris.Setup(app.BaseApp, app.EVMKeeper); err != nil {
		panic(err)
	}

	// register streaming services
	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		panic(err)
	}

	/****  Module Options ****/
	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	app.RegisterUpgradeHandlers()

	// Load the app.
	if err := app.Load(loadLatest); err != nil {
		panic(err)
	}

	// Load the last state of the polaris evm.
	if err := app.Polaris.LoadLastState(
		app.CommitMultiStore(), uint64(app.LastBlockHeight()),
	); err != nil {
		panic(err)
	}

	return app
}

// Name returns the name of the App.
func (app *SimApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

func (app *SimApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// SimulationManager implements the SimulationApp interface.
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return nil
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(
		apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger,
	); err != nil {
		panic(err)
	}

	if err := app.Polaris.Init(apiSvr.ClientCtx, app.Logger(), app.EVMKeeper); err != nil {
		panic(err)
	}
}

// Close shuts down the application.
func (app *SimApp) Close() error {
	if pl := app.Polaris; pl != nil {
		return pl.Close()
	}
	return app.BaseApp.Close()
}
