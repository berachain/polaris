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

//nolint:revive // embed.
package runtime

import (
	"io"
	"os"
	"path/filepath"

	dbm "github.com/cosmos/cosmos-db"

	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	testdata_pulsar "github.com/cosmos/cosmos-sdk/testutil/testdata/testpb"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	polarisbaseapp "pkg.berachain.dev/polaris/cosmos/runtime/baseapp"
	simappconfig "pkg.berachain.dev/polaris/cosmos/runtime/config"
	evmante "pkg.berachain.dev/polaris/cosmos/x/evm/ante"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/lib/utils"

	_ "embed"

	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
)

var (
	// DefaultNodeHome default home directories for the application daemon.
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(polarisbaseapp.ModuleBasics...)

	// application configuration (used by depinject).
	AppConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: simappconfig.DefaultModule,
	})
)

var (
	_ runtime.AppI            = (*PolarisBaseApp)(nil)
	_ servertypes.Application = (*PolarisBaseApp)(nil)
)

// PolarisBaseApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PolarisBaseApp struct {
	polarisbaseapp.PolarisBaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry
	autoCliOpts       autocli.AppOptions

	// simulation manager
	sm *module.SimulationManager
}

//nolint:gochecknoinits // its okay.
func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".polard")

	simappconfig.SetupCosmosConfig()
}

// NewPolarisBaseApp returns a reference to an initialized PolarisBaseApp.
func NewPolarisBaseApp( //nolint: funlen // from sdk.
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *PolarisBaseApp {
	var (
		app        = &PolarisBaseApp{}
		appBuilder *runtime.AppBuilder
		// Below we could construct and set an application specific mempool and ABCI 1.0 Prepare and Process Proposal
		// handlers. These defaults are already set in the SDK's BaseApp, this shows an example of how to override
		// them.
		//
		// nonceMempool = mempool.NewSenderNonceMempool()
		// ethTxMempool = mempool.NewEthTxPool()
		ethTxMempool mempool.Mempool = evmmempool.NewEthTxPoolFrom(
			mempool.NewPriorityMempool(mempool.DefaultPriorityNonceMempoolConfig()),
		)
		mempoolOpt = baseapp.SetMempool(
			ethTxMempool,
		)

		// prepareOpt   = func(app *baseapp.BaseApp) {
		// 	app.SetPrepareProposal(app.DefaultPrepareProposal())
		// }
		// processOpt = func(app *baseapp.BaseApp) {
		// 	app.SetProcessProposal(app.DefaultProcessProposal())
		// }
		//

		// Further down we'd set the options in the AppBuilder like below.

		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			AppConfig,
			depinject.Supply(
				app.App,
				// supply the application options
				appOpts,
				// ADVANCED CONFIGURATION
				//
				// ETH TX MEMPOOL
				ethTxMempool,
				// evmtx.CustomSignModeHandlers,
				//
				//
				func() []signing.SignModeHandler {
					return []signing.SignModeHandler{evmante.SignModeEthTxHandler{}}
				},
				polarisbaseapp.PrecompilesToInject(&app.PolarisBaseApp),
				// AUTH
				//
				// For providing a custom function required in auth to generate custom account types
				// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
				//
				// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),
				//
				// For providing a custom a base account type add it below.
				// By default the auth module uses authtypes.ProtoBaseAccount().
				//
				// func() sdk.AccountI { return authtypes.ProtoBaseAccount() },
				//
				// MINT
				//
				// For providing a custom inflation function for x/evm add here your
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
		&app.autoCliOpts,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.DistrKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.GroupKeeper,
		&app.ConsensusParamsKeeper,
		&app.EVMKeeper,
		&app.ERC20Keeper,
	); err != nil {
		panic(err)
	}

	// Build app
	app.App = appBuilder.Build(logger, db, traceStore, PolarisAppOptions(
		app.interfaceRegistry, append(baseAppOptions, mempoolOpt)...)...,
	)

	// ===============================================================
	// THE "DEPINJECT IS CAUSING PROBLEMS" SECTION
	// ===============================================================

	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		homePath = DefaultNodeHome
	}

	// setup evm keeper and all of its plugins.
	app.EVMKeeper.Setup(
		app.CreateQueryContext,
		// TODO: clean this up.
		homePath+"/config/polaris.toml",
		homePath+"/data/polaris",
	)

	opt := ante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		SignModeHandler: app.txConfig.SignModeHandler(),
		FeegrantKeeper:  app.FeeGrantKeeper,
		SigGasConsumer:  evmante.SigVerificationGasConsumer,
	}
	ch, _ := evmante.NewAnteHandler(
		opt,
	)
	app.SetAnteHandler(
		ch,
	)

	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	app.RegisterUpgradeHandlers()

	// add test gRPC service for testing gRPC queries in isolation
	testdata_pulsar.RegisterQueryServer(app.GRPCQueryRouter(), testdata_pulsar.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required for apps that don't use the simulator for fuzz testing
	// transactions
	// overrideModules := map[string]module.AppModuleSimulation{
	// 	authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper,
	// 		authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	// }
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, nil)

	app.sm.RegisterStoreDecoders()

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init
	// genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring),
	// the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	//
	// app.SetInitChainer(func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// 	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	// 	return app.App.InitChainer(ctx, req)
	// })

	if err := app.Load(loadLatest); err != nil {
		panic(err)
	}

	return app
}

// Name returns the name of the App.
func (app *PolarisBaseApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns PolarisBaseApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisBaseApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns PolarisBaseApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisBaseApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns PolarisBaseApp's InterfaceRegistry.
func (app *PolarisBaseApp) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns PolarisBaseApp's TxConfig.
func (app *PolarisBaseApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// AutoCliOpts returns the autocli options for the app.
func (app *PolarisBaseApp) AutoCliOpts() autocli.AppOptions {
	return app.autoCliOpts
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *PolarisBaseApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	kvStoreKey, ok := utils.GetAs[*storetypes.KVStoreKey](app.UnsafeFindStoreKey(storeKey))
	if !ok {
		return nil
	}
	return kvStoreKey
}

func (app *PolarisBaseApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := utils.GetAs[*storetypes.KVStoreKey](k); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *PolarisBaseApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface.
func (app *PolarisBaseApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *PolarisBaseApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
	app.EVMKeeper.SetClientCtx(apiSvr.ClientCtx)
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dup := make(map[string][]string)
	for _, perms := range simappconfig.ModuleAccPerms {
		dup[perms.Account] = perms.Permissions
	}

	return dup
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)

	if len(simappconfig.BlockAccAddrs) > 0 {
		for _, addr := range simappconfig.BlockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}

	return result
}

func GetHomePath(appOpts servertypes.AppOptions) string {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return DefaultNodeHome
	}
	return homePath
}
