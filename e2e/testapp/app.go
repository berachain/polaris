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
	"io"
	"os"
	"path/filepath"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"

	evmv1alpha1 "github.com/berachain/polaris/cosmos/api/polaris/evm/v1alpha1"
	evmconfig "github.com/berachain/polaris/cosmos/config"
	ethcryptocodec "github.com/berachain/polaris/cosmos/crypto/codec"
	signinglib "github.com/berachain/polaris/cosmos/lib/signing"
	polarruntime "github.com/berachain/polaris/cosmos/runtime"
	"github.com/berachain/polaris/cosmos/runtime/ante"
	"github.com/berachain/polaris/cosmos/runtime/miner"
	evmkeeper "github.com/berachain/polaris/cosmos/x/evm/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
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

	polarisConfig := evmconfig.MustReadConfigFromAppOpts(appOpts)
	if polarisConfig.OptimisticExecution {
		baseAppOptions = append(baseAppOptions, baseapp.SetOptimisticExecution())
	}

	// Build the app using the app builder.
	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)
	app.Polaris = polarruntime.New(app,
		polarisConfig, app.Logger(), app.EVMKeeper.Host, nil,
	)

	// Build cosmos ante handler for non-evm transactions.
	cosmHandler, err := authante.NewAnteHandler(
		authante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			FeegrantKeeper:  nil,
			SigGasConsumer:  ante.EthSecp256k1SigVerificationGasConsumer,
			SignModeHandler: app.txConfig.SignModeHandler(),
			TxFeeChecker: func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
				return nil, 0, nil
			},
		},
	)
	if err != nil {
		panic(err)
	}

	// Setup Polaris Runtime.
	if err = app.Polaris.Build(
		app, cosmHandler, app.EVMKeeper, miner.DefaultAllowedMsgs,
	); err != nil {
		panic(err)
	}

	// register streaming services
	if err = app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		panic(err)
	}

	/****  Module Options ****/
	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	app.RegisterUpgradeHandlers()

	// Register eth_secp256k1 keys
	ethcryptocodec.RegisterInterfaces(app.interfaceRegistry)

	// Load the app.
	if err = app.Load(loadLatest); err != nil {
		panic(err)
	}

	// Load the last state of the Polaris EVM.
	// TODO: verify that the version of app CMS is the same as app.LastBlockHeight.
	if err = app.Polaris.LoadLastState(
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

	if err := app.Polaris.SetupServices(apiSvr.ClientCtx); err != nil {
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
