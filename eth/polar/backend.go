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

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/consensus"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/log"
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

// executionLayerNode defines methods that allow a Polaris chain to build and expose JSON-RPC API.
type executionLayerNode interface {
	ExtRPCEnabled() bool
	RegisterAPIs([]rpc.API)
	RegisterLifecycle(node.Lifecycle)
	EventMux() *event.TypeMux //nolint:staticcheck // deprecated but still in geth.
}

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	config *Config
	// core pieces of the polaris stack
	host       core.PolarisHostChain
	blockchain core.Blockchain
	txPool     *txpool.TxPool
	miner      *miner.Miner

	// apiBackend is utilize by the api handlers as a middleware between the
	// JSON-RPC APIs and the core pieces.
	apiBackend APIBackend
	syncStatus SyncStatusProvider

	// engine represents the consensus engine for the backend.
	engine consensus.Engine

	// filterSystem is the filter system that is used by the filter API.
	// TODO: relocate
	filterSystem *filters.FilterSystem
}

// New creates a new backend for the Polaris EVM.
func New(
	config *Config,
	host core.PolarisHostChain,
	engine consensus.Engine,
	stack executionLayerNode,
	logHandler log.Handler,
) *Polaris {
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

	if config.Miner.GasPrice == nil || config.Miner.GasPrice.Cmp(common.Big0) <= 0 {
		log.Warn("Sanitizing invalid miner gas price",
			"provided", config.Miner.GasPrice, "updated", ethconfig.Defaults.Miner.GasPrice)
		config.Miner.GasPrice = new(big.Int).Set(ethconfig.Defaults.Miner.GasPrice)
	}

	if engine == nil {
		engine = beacon.New(&consensus.DummyEthOne{})
	}

	pl := &Polaris{
		config:     config,
		host:       host,
		engine:     engine,
		blockchain: core.NewChain(host, config.Genesis.Config, engine),
	}

	// Build the backend api object.
	pl.apiBackend = NewAPIBackend(pl, stack.ExtRPCEnabled(), pl.config)

	// Run safety message for feedback to the user if they are running
	// with development configs.
	pl.config.SafetyMessage()

	// For now, we only have a legacy pool, we will implement blob pool later.
	legacyPool := legacypool.New(
		pl.config.LegacyTxPool, pl.Blockchain(),
	)

	// Setup the transaction pool and attach the legacyPool.
	var err error
	if pl.txPool, err = txpool.New(
		new(big.Int).SetUint64(pl.config.LegacyTxPool.PriceLimit),
		pl.blockchain,
		[]txpool.SubPool{legacyPool},
	); err != nil {
		panic(err)
	}

	// Setup the miner, we use a dummy isLocal function, since it is not used.
	pl.miner = miner.New(pl, &pl.config.Miner,
		pl.config.Genesis.Config, stack.EventMux(), pl.engine,
		func(header *types.Header) bool { return true },
	)

	// Register the backend on the node
	stack.RegisterAPIs(pl.APIs())
	stack.RegisterLifecycle(pl)

	// Register the filter API separately in order to get access to the filterSystem
	pl.filterSystem = utils.RegisterFilterAPI(stack, pl.apiBackend, &defaultEthConfig)
	return pl
}

// Start implements node.Lifecycle, starting all internal goroutines needed by the
// Polaris protocol implementation.
func (pl *Polaris) Start() error {
	return nil
}

// Stop implements node.Lifecycle, terminating all internal goroutines used by the
// Polaris protocol.
func (pl *Polaris) Stop() error {
	return nil
}

// APIs return the collection of RPC services the polar package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (pl *Polaris) APIs() []rpc.API {
	// Grab a bunch of the apis from go-Polaris (thx bae)
	apis := polarapi.GethAPIs(pl.apiBackend, pl.blockchain)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "net",
			Service:   polarapi.NewNetAPI(pl.apiBackend),
		},
		{
			Namespace: "web3",
			Service:   polarapi.NewWeb3API(pl.apiBackend),
		},
	}...)
}

// Host returns the Polaris host chain.
func (pl *Polaris) Host() core.PolarisHostChain {
	return pl.host
}

// Engine returns the consensus engine.
func (pl *Polaris) Engine() consensus.Engine { return pl.engine }

// Miner returns the miner.
func (pl *Polaris) Miner() *miner.Miner {
	return pl.miner
}

// TxPool returns the transaction pool.
func (pl *Polaris) TxPool() *txpool.TxPool {
	return pl.txPool
}

// MinerChain returns the blockchain.
func (pl *Polaris) MinerChain() miner.BlockChain {
	return pl.Blockchain()
}

// Blockchain returns the blockchain.
func (pl *Polaris) Blockchain() core.Blockchain {
	return pl.blockchain
}
