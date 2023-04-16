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
	"cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
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
	testdata_pulsar "github.com/cosmos/cosmos-sdk/testutil/testdata/testpb"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	authprecompile "pkg.berachain.dev/polaris/cosmos/precompile/auth"
	bankprecompile "pkg.berachain.dev/polaris/cosmos/precompile/bank"
	distrprecompile "pkg.berachain.dev/polaris/cosmos/precompile/distribution"
	govprecompile "pkg.berachain.dev/polaris/cosmos/precompile/governance"
	stakingprecompile "pkg.berachain.dev/polaris/cosmos/precompile/staking"
	simappconfig "pkg.berachain.dev/polaris/cosmos/runtime/config"
	"pkg.berachain.dev/polaris/cosmos/x/evm"
	evmante "pkg.berachain.dev/polaris/cosmos/x/evm/ante"
	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/provider"
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
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		evm.AppModuleBasic{},
	)
	// application configuration (used by depinject).
	AppConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: simappconfig.DefaultModule,
	})
)

var (
	_ runtime.AppI            = (*PolarisApp)(nil)
	_ servertypes.Application = (*PolarisApp)(nil)
)

// PolarisApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PolarisApp struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry
	autoCliOpts       autocli.AppOptions

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
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	// polaris keepers
	EVMKeeper *evmkeeper.Keeper

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
func NewPolarisApp( //nolint: funlen // from sdk.
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

	// Polaris Configuration
	cfg := provider.DefaultConfig()
	cfg.NodeConfig.DataDir = DefaultNodeHome + "/data"

	// setup evm keeper and all of its plugins.
	app.EVMKeeper.Setup(
		app.AccountKeeper,
		app.BankKeeper,
		[]vm.RegistrablePrecompile{
			// TODO: register more precompiles here.
			stakingprecompile.NewPrecompileContract(app.StakingKeeper),
			bankprecompile.NewPrecompileContract(app.BankKeeper),
			authprecompile.NewPrecompileContract(),
			distrprecompile.NewPrecompileContract(),
			govprecompile.NewPrecompileContract(app.GovKeeper),
		},
		app.CreateQueryContext,
		cfg,
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

	// Set the query context function for the evm module.

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	app.RegisterUpgradeHandlers()

	// add test gRPC service for testing gRPC queries in isolation
	testdata_pulsar.RegisterQueryServer(app.GRPCQueryRouter(), testdata_pulsar.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper,
			authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)

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
func (app *PolarisApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns PolarisApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns PolarisApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns PolarisApp's InterfaceRegistry.
func (app *PolarisApp) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns PolarisApp's TxConfig.
func (app *PolarisApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// AutoCliOpts returns the autocli options for the app.
func (app *PolarisApp) AutoCliOpts() autocli.AppOptions {
	return app.autoCliOpts
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *PolarisApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	kvStoreKey, ok := utils.GetAs[*storetypes.KVStoreKey](app.UnsafeFindStoreKey(storeKey))
	if !ok {
		return nil
	}
	return kvStoreKey
}

func (app *PolarisApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
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
func (app *PolarisApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface.
func (app *PolarisApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *PolarisApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
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
