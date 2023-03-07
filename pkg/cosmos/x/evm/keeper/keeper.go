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

package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/eth"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/vm"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	ethrpcconfig "pkg.berachain.dev/polaris/eth/rpc/config"
	"pkg.berachain.dev/polaris/lib/utils"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/configuration"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/gas"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/precompile"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/txpool"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/plugins/txpool/mempool"
	evmrpc "pkg.berachain.dev/polaris/pkg/cosmos/x/evm/rpc"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/store/offchain"
)

// Compile-time interface assertion.
var _ core.PolarisHostChain = (*Keeper)(nil)

type Keeper struct {
	// `provider` is the struct that houses the Polaris EVM.
	polaris *eth.PolarisProvider
	// We store a reference to the `rpcProvider` so that we can register it with
	// the cosmos mux router.
	rpcProvider evmrpc.Provider
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey
	// The offchain KV store.
	offChainKv *offchain.Store
	// `authority` is the bech32 address that is allowed to execute governance proposals.
	authority string

	// The various plugins that are are used to implement `core.PolarisHostChain`.
	bp  block.Plugin
	cp  configuration.Plugin
	gp  gas.Plugin
	pp  precompile.Plugin
	sp  state.Plugin
	txp txpool.Plugin
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	bk state.BankKeeper,
	getPrecompiles func() func() []vm.RegistrablePrecompile,
	authority string,
	appOpts servertypes.AppOptions,
	ethTxMempool sdkmempool.Mempool,
) *Keeper {
	// We setup the keeper with some Cosmos standard sauce.
	k := &Keeper{
		authority: authority,
		storeKey:  storeKey,
	}

	// TODO: parameterize kv store.
	if appOpts != nil {
		k.offChainKv = offchain.NewOffChainKVStore("eth_indexer", appOpts)
	}

	// Setup the RPC Service. // TODO: parameterize config.
	cfg := ethrpcconfig.DefaultServer()
	cfg.BaseRoute = "/eth/rpc"
	k.rpcProvider = evmrpc.NewProvider(cfg)

	// Build the Plugins
	k.bp = block.NewPlugin(k.offChainKv, storeKey)
	k.cp = configuration.NewPlugin(storeKey)
	k.gp = gas.NewPlugin()
	k.txp = txpool.NewPlugin(k.rpcProvider, utils.MustGetAs[*mempool.EthTxPool](ethTxMempool))
	return k
}

// `ConfigureGethLogger` configures the Geth logger to use the Cosmos logger.
func (k *Keeper) ConfigureGethLogger(ctx sdk.Context) {
	ethlog.Root().SetHandler(ethlog.FuncHandler(func(r *ethlog.Record) error {
		logger := ctx.Logger().With("module", "geth")
		switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
		case ethlog.LvlTrace, ethlog.LvlDebug:
			logger.Debug(r.Msg, r.Ctx...)
		case ethlog.LvlInfo, ethlog.LvlWarn:
			logger.Info(r.Msg, r.Ctx...)
		case ethlog.LvlError, ethlog.LvlCrit:
			logger.Error(r.Msg, r.Ctx...)
		}
		return nil
	}))
}

// `Logger` returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}

// `Setup` sets up the precompile and state plugins with the given precompiles and keepers. It also
// sets the query context function for the block and state plugins (to support historical queries).
func (k *Keeper) Setup(
	ak state.AccountKeeper,
	bk state.BankKeeper,
	precompiles []vm.RegistrablePrecompile,
	qc func(height int64, prove bool) (sdk.Context, error),
) {
	// Setup the precompile and state plugins
	k.pp = precompile.NewPlugin(precompiles)
	k.sp = state.NewPlugin(ak, bk, k.storeKey, k.cp, k.pp)

	// Set the query context function for the block and state plugins
	k.sp.SetQueryContextFn(qc)
	k.bp.SetQueryContextFn(qc)

	// Build the Polaris EVM Provider
	k.polaris = eth.NewPolarisProvider(k, k.rpcProvider, nil)
}

// `GetBlockPlugin` returns the block plugin.
func (k *Keeper) GetBlockPlugin() core.BlockPlugin {
	return k.bp
}

// `GetConfigurationPlugin` returns the configuration plugin.
func (k *Keeper) GetConfigurationPlugin() core.ConfigurationPlugin {
	return k.cp
}

// `GetGasPlugin` returns the gas plugin.
func (k *Keeper) GetGasPlugin() core.GasPlugin {
	return k.gp
}

// `GetPrecompilePlugin` returns the precompile plugin.
func (k *Keeper) GetPrecompilePlugin() core.PrecompilePlugin {
	return k.pp
}

// `GetStatePlugin` returns the state plugin.
func (k *Keeper) GetStatePlugin() core.StatePlugin {
	return k.sp
}

// `GetTxPoolPlugin` returns the txpool plugin.
func (k *Keeper) GetTxPoolPlugin() core.TxPoolPlugin {
	return k.txp
}

// `GetAllPlugins` returns all the plugins.
func (k *Keeper) GetAllPlugins() []plugins.BaseCosmosPolaris {
	return []plugins.BaseCosmosPolaris{k.bp, k.cp, k.gp, k.pp, k.sp}
}

// `GetRPCProvider` returns the RPC provider. We use this in `app.go` to register
// the Ethereum JSONRPC server with the application mux server.
func (k *Keeper) GetRPCProvider() evmrpc.Provider {
	return k.rpcProvider
}
