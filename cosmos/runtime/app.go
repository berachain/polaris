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
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	testdata_pulsar "github.com/cosmos/cosmos-sdk/testutil/testdata/testpb"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	polarisbaseapp "pkg.berachain.dev/polaris/cosmos/runtime/baseapp"
	simappconfig "pkg.berachain.dev/polaris/cosmos/runtime/config"
	evmante "pkg.berachain.dev/polaris/cosmos/x/evm/ante"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"

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
	_ runtime.AppI            = (*PolarisApp)(nil)
	_ servertypes.Application = (*PolarisApp)(nil)
)

// PolarisBaseApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PolarisApp struct {
	polarisbaseapp.PolarisBaseApp

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

// NewPolarisApp returns a reference to an initialized PolarisApp.
func NewPolarisApp( //nolint:funlen // as defined by the sdk.
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *PolarisApp {
	var (
		app        = &PolarisApp{}
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
		&app.ApplicationCodec,
		&app.LegacyAminoCodec,
		&app.TxnConfig,
		&app.CodecInterfaceRegistry,
		&app.AutoCliOptions,
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

	// Build app with the provided options.
	app.App = appBuilder.Build(logger, db, traceStore, append(baseAppOptions, mempoolOpt)...)
	// TODO: move this somewhere better, introduce non IAVL enforced module keys as a PR to the SDK
	// we ask @tac0turtle how 2 fix
	offchainKey := storetypes.NewKVStoreKey("offchain-evm")
	app.PolarisBaseApp.MountCustomStores(offchainKey)

	// ===============================================================
	// THE "DEPINJECT IS CAUSING PROBLEMS" SECTION
	// ===============================================================

	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		homePath = DefaultNodeHome
	}

	// setup evm keeper and all of its plugins.
	app.EVMKeeper.Setup(
		offchainKey,
		app.CreateQueryContext,
		// TODO: clean this up.
		homePath+"/config/polaris.toml",
		homePath+"/data/polaris",
	)

	opt := ante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		SignModeHandler: app.TxConfig().SignModeHandler(),
		FeegrantKeeper:  app.FeeGrantKeeper,
		SigGasConsumer:  evmante.SigVerificationGasConsumer,
	}
	ch, _ := evmante.NewAnteHandler(
		opt,
	)
	app.SetAnteHandler(
		ch,
	)

	// We must register the EthSecp256k1 signature type because it is not registered by default.
	// TODO: remove once upstreamed to the SDK.
	app.RegisterEthSecp256k1SignatureType()

	if err := app.RegisterStreamingServices(appOpts, app.KVStoreKeys()); err != nil {
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

// SimulationManager implements the SimulationApp interface.
func (app *PolarisApp) SimulationManager() *module.SimulationManager {
	return app.sm
}
