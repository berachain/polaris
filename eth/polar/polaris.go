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

package polar

import (
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/consensus"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/log"
	oldminer "pkg.berachain.dev/polaris/eth/miner"
	polarapi "pkg.berachain.dev/polaris/eth/polar/api"
	"pkg.berachain.dev/polaris/eth/rpc"
)

// TODO: break out the node into a separate package and then fully use the
// abstracted away networking stack, by extension we will need to improve the registration
// architecture.

var defaultEthConfig = ethconfig.Config{
	SyncMode:           0,
	FilterLogCacheSize: 0,
}

// NetworkingStack defines methods that allow a Polaris chain to build and expose JSON-RPC API.
type NetworkingStack interface {
	// IsExtRPCEnabled returns true if the networking stack is configured to expose JSON-RPC API.
	ExtRPCEnabled() bool

	// RegisterHandler manually registers a new handler into the networking stack.
	RegisterHandler(string, string, http.Handler)

	// RegisterAPIs registers JSON-RPC handlers for the networking stack.
	RegisterAPIs([]rpc.API)

	// RegisterLifecycles registers objects to have their lifecycle manged by the stack.
	RegisterLifecycle(node.Lifecycle)

	// Start starts the networking stack.
	Start() error

	// Close stops the networking stack
	Close() error
}

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	cfg *Config
	// NetworkingStack represents the networking stack responsible for exposes the JSON-RPC
	// APIs. Although possible, it does not handle p2p networking like its sibling in geth
	// would.
	stack NetworkingStack

	// core pieces of the polaris stack
	host       core.PolarisHostChain
	blockchain core.Blockchain
	txPool     *txpool.TxPool
	spminer    oldminer.Miner
	miner      *miner.Miner

	// backend is utilize by the api handlers as a middleware between the JSON-RPC APIs
	// and the core pieces.
	backend Backend

	// engine represents the consensus engine for the backend.
	engine core.EnginePlugin
	beacon consensus.Engine

	// filterSystem is the filter system that is used by the filter API.
	// TODO: relocate
	filterSystem *filters.FilterSystem
}

func NewWithNetworkingStack(
	cfg *Config,
	host core.PolarisHostChain,
	stack NetworkingStack,
	logHandler log.Handler,
) *Polaris {
	pl := &Polaris{
		cfg:        cfg,
		blockchain: core.NewChain(host),
		stack:      stack,
		host:       host,
		engine:     host.GetEnginePlugin(),
		beacon:     &consensus.DummyEthOne{},
	}
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		log.Root().SetHandler(logHandler)
	}

	pl.backend = NewBackend(pl, pl.stack.ExtRPCEnabled(), pl.cfg)

	return pl
}

// Init initializes the Polaris struct.
func (pl *Polaris) Init() error {
	var err error
	pl.spminer = oldminer.New(pl)

	// For now, we only have a legacy pool, we will implement blob pool later.
	legacyPool := legacypool.New(
		pl.cfg.LegacyTxPool, pl.Blockchain(),
	)

	// Setup the transaction pool and attach the legacyPool.
	if pl.txPool, err = txpool.New(
		new(big.Int).SetUint64(pl.cfg.LegacyTxPool.PriceLimit),
		pl.blockchain,
		[]txpool.SubPool{legacyPool},
	); err != nil {
		return err
	}

	mux := new(event.TypeMux) //nolint:staticcheck // deprecated but still in geth.
	// TODO: miner config to app.toml
	pl.miner = miner.New(pl, &pl.cfg.Miner,
		pl.host.GetConfigurationPlugin().ChainConfig(), mux, pl.beacon, pl.IsLocalBlock)
	// extra data must be nil until 1 block 1 transaction.
	// eth.miner.SetExtra(makeExtraData(config.Miner.ExtraData))
	// Build and set the RPC Backend and other services.

	// if eth.APIBackend.allowUnprotectedTxs {
	// 	log.Info("Unprotected transactions allowed")
	// }

	return nil
}

// APIs return the collection of RPC services the polar package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (pl *Polaris) APIs() []rpc.API {
	// Grab a bunch of the apis from go-ethereum (thx bae)
	apis := polarapi.GethAPIs(pl.backend, pl.blockchain)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "net",
			Service:   polarapi.NewNetAPI(pl.backend),
		},
		{
			Namespace: "web3",
			Service:   polarapi.NewWeb3API(pl.backend),
		},
	}...)
}

// StartServices notifies the NetworkStack to spin up (i.e json-rpc).
func (pl *Polaris) StartServices() error {
	// Register the JSON-RPCs with the networking stack.
	pl.stack.RegisterAPIs(pl.APIs())

	// Register the filter API separately in order to get access to the filterSystem
	pl.filterSystem = utils.RegisterFilterAPI(pl.stack, pl.backend, &defaultEthConfig)

	go func() {
		time.Sleep(3 * time.Second) //nolint:gomnd // TODO:fix, this is required for hive...
		if err := pl.stack.Start(); err != nil {
			panic(err)
		}
	}()
	return nil
}

// RegisterService adds a service to the networking stack.
func (pl *Polaris) RegisterService(lc node.Lifecycle) {
	pl.stack.RegisterLifecycle(lc)
}

func (pl *Polaris) Close() error {
	return pl.stack.Close()
}

func (pl *Polaris) Host() core.PolarisHostChain {
	return pl.host
}

func (pl *Polaris) Miner() *miner.Miner {
	return pl.miner
}

func (pl *Polaris) SPMiner() oldminer.Miner {
	return pl.spminer
}

func (pl *Polaris) TxPool() *txpool.TxPool {
	return pl.txPool
}

func (pl *Polaris) MinerChain() miner.BlockChain {
	return pl.blockchain
}

func (pl *Polaris) Blockchain() core.Blockchain {
	return pl.blockchain
}

func (pl *Polaris) IsLocalBlock(_ *types.Header) bool {
	// Unused.
	return true
}
