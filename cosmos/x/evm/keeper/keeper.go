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
	"math/big"
	"time"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/engine"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/polar"
	"pkg.berachain.dev/polaris/lib/utils"
)

type Keeper struct {
	// ak is the reference to the AccountKeeper.
	ak state.AccountKeeper
	// provider is the struct that houses the Polaris EVM.
	polaris *polar.Polaris
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey
	// The host contains various plugins that are are used to implement `core.PolarisHostChain`.
	host Host

	// temp syncing
	lock bool
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	ak state.AccountKeeper,
	sk block.StakingKeeper,
	storeKey storetypes.StoreKey,
	ethTxMempool sdkmempool.Mempool,
	pcs func() *ethprecompile.Injector,
) *Keeper {
	// We setup the keeper with some Cosmos standard sauce.
	k := &Keeper{
		ak:       ak,
		storeKey: storeKey,
		lock:     true,
	}

	k.host = NewHost(
		storeKey,
		sk,
		ethTxMempool,
		pcs,
	)
	return k
}

// Setup sets up the plugins in the Host. It also build the Polaris EVM Provider.
func (k *Keeper) Setup(
	_ *storetypes.KVStoreKey,
	qc func(height int64, prove bool) (sdk.Context, error),
	polarisConfigPath string,
	polarisDataDir string,
	logger log.Logger,
) {
	// Setup plugins in the Host
	k.host.Setup(k.storeKey, nil, k.ak, qc)

	// Build the Polaris EVM Provider
	cfg, err := polar.LoadConfigFromFilePath(polarisConfigPath)
	// TODO: fix properly.
	if err != nil || cfg.GPO == nil {
		// TODO: log warning for this case.
		logger.Error("failed to load polaris config", "falling back to defaults")
		cfg = polar.DefaultConfig()
	}

	// TODO: PARSE POLARIS.TOML CORRECT AGAIN
	nodeCfg := polar.DefaultGethNodeConfig()
	nodeCfg.DataDir = polarisDataDir
	node, err := polar.NewGethNetworkingStack(nodeCfg)
	if err != nil {
		panic(err)
	}

	k.polaris = polar.NewWithNetworkingStack(cfg, k.host, node, ethlog.FuncHandler(
		func(r *ethlog.Record) error {
			polarisGethLogger := logger.With("module", "polaris-geth")
			switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
			case ethlog.LvlTrace, ethlog.LvlDebug:
				polarisGethLogger.Debug(r.Msg, r.Ctx...)
			case ethlog.LvlInfo, ethlog.LvlWarn:
				polarisGethLogger.Info(r.Msg, r.Ctx...)
			case ethlog.LvlError, ethlog.LvlCrit:
				polarisGethLogger.Error(r.Msg, r.Ctx...)
			}
			return nil
		}),
	)
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}

// GetHost returns the Host that contains all plugins.
func (k *Keeper) GetHost() Host {
	return k.host
}

// StartServices waits until the first block is being processed for the lock to unlock before
// starting the networking stack and txpool service.
func (k *Keeper) StartServices(clientContext client.Context) {
	// Set the Polaris blockchain.
	k.SetupBlockchain()

	// spin lock until the first block is being processed
	for ; k.lock; time.Sleep(1000 * time.Millisecond) { //nolint:gomnd // 100ms is fine.
		continue
	}

	// start the txpool service
	txpool := txpool.NewTxPool(
		txpool.DefaultConfig, k.host.GetConfigurationPlugin().ChainConfig(), k.polaris.Blockchain(),
	)
	k.polaris.SetTxPool(txpool)
	k.polaris.SetEngine(k.host.GetEnginePlugin())
	k.host.GetTxPoolPlugin().Start(txpool, clientContext)
	utils.MustGetAs[engine.Plugin](k.host.GetEnginePlugin()).Start(clientContext)

	// start the networking stack (json-rpc, graphql, etc.)
	if err := k.polaris.StartServices(); err != nil {
		panic(err)
	}
}

func (k *Keeper) SetupBlockchain() {
	// Set the Polaris blockchain.
	k.polaris.SetBlockchain(core.NewChain(k.host))
}

// GetPolaris returns the Polaris instance.
func (k *Keeper) GetPolaris() *polar.Polaris {
	return k.polaris
}

// TODO: Remove these, because they're hacky af.
// Required temporarily for BGT plugin.
func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress) *big.Int {
	ethAddr := common.BytesToAddress(addr)
	return new(big.Int).SetBytes(ctx.KVStore(k.storeKey).Get(state.BalanceKeyFor(ethAddr)))
}

func (k *Keeper) SetBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(state.BalanceKeyFor(ethAddr), amount.Bytes())
}

func (k *Keeper) AddBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(state.BalanceKeyFor(ethAddr), new(big.Int).Add(k.GetBalance(ctx, addr), amount).Bytes())
}

func (k *Keeper) SubBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(state.BalanceKeyFor(ethAddr), new(big.Int).Sub(k.GetBalance(ctx, addr), amount).Bytes())
}
