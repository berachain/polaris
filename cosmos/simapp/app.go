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
package simapp

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
	"cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"

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
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/consensus"
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

	ethcryptocodec "pkg.berachain.dev/polaris/cosmos/crypto/codec"
	polarisruntime "pkg.berachain.dev/polaris/cosmos/runtime"
	simappconfig "pkg.berachain.dev/polaris/cosmos/simapp/config"
	"pkg.berachain.dev/polaris/cosmos/x/erc20"
	erc20keeper "pkg.berachain.dev/polaris/cosmos/x/erc20/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm"
	evmante "pkg.berachain.dev/polaris/cosmos/x/evm/ante"
	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
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
	ModuleBasics = module.NewBasicManager([]module.AppModuleBasic{
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
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		consensus.AppModuleBasic{},
		evm.AppModuleBasic{},
		erc20.AppModuleBasic{},
	}...,
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

// polarisruntime extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PolarisApp struct {
	*polarisruntime.PolarisApp

	LegacyAminoCodec       *codec.LegacyAmino
	ApplicationCodec       codec.Codec
	TxnConfig              client.TxConfig
	CodecInterfaceRegistry codectypes.InterfaceRegistry
	AutoCliOptions         autocli.AppOptions

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
	ConsensusParamsKeeper consensuskeeper.Keeper

	// polaris keepers
	ERC20Keeper *erc20keeper.Keeper
	EVMKeeper   *evmkeeper.Keeper

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
		app          = &PolarisApp{}
		appBuilder   = &polarisruntime.AppBuilder{}
		ethTxMempool = evmmempool.NewPolarisEthereumTxPool()
		appConfig    = depinject.Configs(
			AppConfig,
			depinject.Supply(
				appOpts,
				logger,
				ethTxMempool,
				PrecompilesToInject(app),
			),
		)
	)

	if err := depinject.Inject(appConfig,
		&appBuilder.AppBuilder,
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
		&app.ConsensusParamsKeeper,
		&app.EVMKeeper,
		&app.ERC20Keeper,
	); err != nil {
		panic(err)
	}

	// Build app with the provided options.
	app.PolarisApp = appBuilder.Build(db, traceStore, append(baseAppOptions, baseapp.SetMempool(ethTxMempool))...)

	// ===============================================================
	// THE "DEPINJECT IS CAUSING PROBLEMS" SECTION
	// ===============================================================

	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		homePath = DefaultNodeHome
	}

	// setup evm keeper and all of its plugins.
	app.EVMKeeper.Setup(
		app.GetKey("offchain-evm"),
		app.CreateQueryContext,
		// TODO: clean this up.
		homePath+"/config/polaris.toml",
		homePath+"/data/polaris",
		logger,
	)

	opt := ante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		SignModeHandler: app.TxConfig().SignModeHandler(),
		FeegrantKeeper:  nil,
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
	// testdata_pulsar.RegisterQueryServer(app.GRPCQueryRouter(), testdata_pulsar.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required for apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.ApplicationCodec, app.AccountKeeper,
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

// SimulationManager implements the SimulationApp interface.
func (app *PolarisApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// LegacyAmino returns polarisruntime's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisApp) LegacyAmino() *codec.LegacyAmino {
	return app.LegacyAminoCodec
}

// AppCodec returns polarisruntime's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PolarisApp) AppCodec() codec.Codec {
	return app.ApplicationCodec
}

// InterfaceRegistry returns polarisruntime's InterfaceRegistry.
func (app *PolarisApp) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.CodecInterfaceRegistry
}

// TxConfig returns polarisruntime's TxConfig.
func (app *PolarisApp) TxConfig() client.TxConfig {
	return app.TxnConfig
}

// AutoCliOpts returns the autocli options for the app.
func (app *PolarisApp) AutoCliOpts() autocli.AppOptions {
	return app.AutoCliOptions
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *PolarisApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterEthSecp256k1SignatureType registers the eth_secp256k1 signature type.
func (app *PolarisApp) RegisterEthSecp256k1SignatureType() {
	ethcryptocodec.RegisterInterfaces(app.CodecInterfaceRegistry)
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
