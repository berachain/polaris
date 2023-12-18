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
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/berachain/polaris/eth/consensus"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/eth/core/state"
	polarapi "github.com/berachain/polaris/eth/polar/api"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	gethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
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
	allowUnprotectedTxs bool,
	logHandler log.Handler,
) *Polaris {
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
		blockchain: core.NewChain(host, &config.Chain, engine),
	}

	// Build the backend api object.
	pl.apiBackend = NewAPIBackend(pl, stack.ExtRPCEnabled(), allowUnprotectedTxs, pl.config)

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
		&pl.config.Chain, stack.EventMux(), pl.engine,
		func(header *ethtypes.Header) bool { return true },
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
	apis := ethapi.GetAPIs(pl.apiBackend, pl.blockchain)

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
		{
			Namespace: "debug",
			Service:   tracers.NewAPI(pl.apiBackend),
		},
	}...)
}

// RegisterSyncStatusProvider registers a sync status provider.
func (pl *Polaris) RegisterSyncStatusProvider(
	syncStatus SyncStatusProvider,
) {
	pl.syncStatus = syncStatus
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

// stateAtBlock retrieves the state database associated with a certain block.
// If no state is locally available for the given block, a number of blocks
// are attempted to be reexecuted to generate the desired state. The optional
// base layer statedb can be provided which is regarded as the statedb of the
// parent block.
//
// An additional release function will be returned if the requested state is
// available. Release is expected to be invoked when the returned state is no
// longer needed. Its purpose is to prevent resource leaking. Though it can be
// noop in some cases.
//
// Parameters:
//   - block:      The block for which we want the state(state = block.Root)
//   - reexec:     The maximum number of blocks to reprocess trying to obtain the desired state
//   - base:       If the caller is tracing multiple blocks, the caller can provide the parent
//     state continuously from the callsite.
//   - readOnly:   If true, then the live 'blockchain' state database is used. No mutation should
//     be made from caller, e.g. perform Commit or other 'save-to-disk' changes.
//     Otherwise, the trash generated by caller may be persisted permanently.
//   - preferDisk: This arg can be used by the caller to signal that even though the 'base' is
//     provided, it would be preferable to start from a fresh state, if we have it
//     on disk.
func (pl *Polaris) stateAtBlock(
	_ context.Context, block *ethtypes.Block, _ uint64,
	_ state.StateDB, _ bool, _ bool,
) (state.StateDB, tracers.StateReleaseFunc, error) {
	return pl.pathState(block)
}

// stateAtTransaction returns the execution environment of a certain transaction.
func (pl *Polaris) stateAtTransaction(
	ctx context.Context, block *ethtypes.Block,
	txIndex int, reexec uint64,
) (*gethcore.Message, vm.BlockContext, state.StateDB, tracers.StateReleaseFunc, error) {
	// Short circuit if it's genesis block.
	if block.NumberU64() == 0 {
		return nil, vm.BlockContext{}, nil, nil, errors.New("no transaction in genesis")
	}
	// Create the parent state database
	parent := pl.blockchain.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return nil, vm.BlockContext{}, nil, nil, fmt.Errorf("parent %#x not found", block.ParentHash())
	}
	// Lookup the statedb of parent block from the live database,
	// otherwise regenerate it on the flight.
	statedb, release, err := pl.stateAtBlock(ctx, parent, reexec, nil, true, false)
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, err
	}
	if txIndex == 0 && len(block.Transactions()) == 0 {
		return nil, vm.BlockContext{}, statedb, release, nil
	}
	// Recompute transactions up to the target index.
	signer := ethtypes.MakeSigner(pl.blockchain.Config(), block.Number(), block.Time())
	for idx, tx := range block.Transactions() {
		// Assemble the transaction call message and return if the requested offset
		msg, _ := gethcore.TransactionToMessage(tx, signer, block.BaseFee())
		txContext := gethcore.NewEVMTxContext(msg)
		context := gethcore.NewEVMBlockContext(block.Header(), pl.blockchain, nil)
		if idx == txIndex {
			return msg, context, statedb, release, nil
		}
		// Not yet the searched for transaction, execute on top of the current state
		vmenv := vm.NewEVM(context, txContext, statedb, pl.blockchain.Config(), vm.Config{})
		statedb.SetTxContext(tx.Hash(), idx)
		if _, err = gethcore.ApplyMessage(vmenv,
			msg, new(gethcore.GasPool).AddGas(tx.Gas())); err != nil {
			return nil, vm.BlockContext{}, nil, nil,
				fmt.Errorf("transaction %s failed: %w", tx.Hash().Hex(), err)
		}
		// Ensure any modifications are committed to the state
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		statedb.Finalise(vmenv.ChainConfig().IsEIP158(block.Number()))
	}
	return nil, vm.BlockContext{}, nil, nil,
		fmt.Errorf("transaction index %d out of range for block %#x", txIndex, block.Hash())
}

// pathState function returns the state at a specific block.
func (pl *Polaris) pathState(block *ethtypes.Block) (state.StateDB, func(), error) {
	// Check if the requested state is available in the live chain.
	// StateAt returns the state by root hash.
	statedb, err := pl.blockchain.StateAtBlockNumber(block.Number().Uint64() + 1)
	if err == nil {
		// If there is no error, return the state, a no-op function, and no error.
		return statedb, func() {}, nil
	}
	// If there is an error, it means the state is not available.
	// TODO: Historic state is not supported in path-based scheme.
	// Fully archive node in pbss will be implemented by relying
	// on state history, but needs more work on top.
	return nil, nil, errors.New("historical state not available in path scheme yet")
}
