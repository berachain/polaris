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

package core

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	lru "github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/consensus"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/params"
)

// By default we are storing up to 1024 items in each cache.
const defaultCacheSize = 1024

// Compile-time check to ensure that `blockchain` implements the `Blockchain` api.
var _ Blockchain = (*blockchain)(nil)

// Blockchain interface defines the methods that a blockchain must have.
type Blockchain interface {
	PreparePlugins(ctx context.Context)
	ChainReader
	ChainWriter
	ChainSubscriber
	ChainResources
	ChainContext
}

// blockchain is the canonical, persistent object that operates the Polaris EVM.
type blockchain struct {
	// the host chain plugins that the Polaris EVM is running on.
	bp BlockPlugin
	hp HistoricalPlugin
	pp PrecompilePlugin
	sp StatePlugin

	engine    consensus.Engine
	processor core.Processor

	// statedb is the state database that is used to mange state during transactions.
	statedb state.StateDB
	// vmConfig is the configuration used to create the EVM.
	vmConfig *vm.Config

	// config represents the chain config.
	config *params.ChainConfig

	// currentBlock is the current/pending block.
	currentBlock atomic.Pointer[types.Block]
	// finalizedBlock is the finalized/latest block.
	finalizedBlock atomic.Pointer[types.Block]
	// currentReceipts is the current/pending receipts.
	currentReceipts atomic.Value
	// currentLogs is the current/pending logs.
	currentLogs atomic.Value

	// receiptsCache is a cache of the receipts for the last `defaultCacheSizeBytes` bytes of
	// blocks. blockHash -> receipts
	receiptsCache *lru.Cache[common.Hash, types.Receipts]
	// blockNumCache is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockNum -> block
	blockNumCache *lru.Cache[uint64, *types.Block]
	// blockHashCache is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockHash -> block
	blockHashCache *lru.Cache[common.Hash, *types.Block]
	// txLookupCache is a cache of the transactions for the last `defaultCacheSizeBytes` bytes of
	// blocks. txHash -> txLookupEntry
	txLookupCache *lru.Cache[common.Hash, *types.TxLookupEntry]

	// subscription event feeds
	scope           event.SubscriptionScope
	chainFeed       event.Feed
	chainHeadFeed   event.Feed
	logsFeed        event.Feed
	pendingLogsFeed event.Feed
	rmLogsFeed      event.Feed // currently never used
	chainSideFeed   event.Feed // currently never used
	logger          log.Logger
}

// =========================================================================
// Constructor
// =========================================================================

// NewChain creates and returns a `api.Chain` with the given EVM chain configuration and host.
func NewChain(
	host PolarisHostChain, config *params.ChainConfig, engine consensus.Engine,
) *blockchain { //nolint:revive // only used as `api.Chain`.
	bc := &blockchain{
		bp:             host.GetBlockPlugin(),
		hp:             host.GetHistoricalPlugin(),
		pp:             host.GetPrecompilePlugin(),
		sp:             host.GetStatePlugin(),
		config:         config,
		vmConfig:       &vm.Config{},
		receiptsCache:  lru.NewCache[common.Hash, types.Receipts](defaultCacheSize),
		blockNumCache:  lru.NewCache[uint64, *types.Block](defaultCacheSize),
		blockHashCache: lru.NewCache[common.Hash, *types.Block](defaultCacheSize),
		txLookupCache:  lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSize),
		chainHeadFeed:  event.Feed{},
		scope:          event.SubscriptionScope{},
		logger:         log.Root(),
		engine:         engine,
	}
	bc.statedb = state.NewStateDB(bc.sp, bc.pp)
	bc.processor = core.NewStateProcessor(bc.config, bc, bc.engine)
	// TODO: hmm...
	bc.currentBlock.Store(
		types.NewBlock(&types.Header{Time: 0, Number: big.NewInt(0),
			BaseFee: big.NewInt(0)}, nil, nil, nil, trie.NewStackTrie(nil)))
	bc.finalizedBlock.Store(nil)
	return bc
}

func (bc *blockchain) LoadLastState(ctx context.Context, number uint64) error {
	// ctx here is the one created from app.CommitMultistore().
	bc.PreparePlugins(ctx)

	return bc.loadLastState(number)
}

func (bc *blockchain) PreparePlugins(ctx context.Context) {
	bc.sp.Reset(ctx)
	bc.bp.Prepare(ctx)
	if bc.hp != nil {
		bc.hp.Prepare(ctx)
	}
}

// ChainConfig returns the Ethereum chain config of the  chain.
func (bc *blockchain) Config() *params.ChainConfig {
	return bc.config
}

// loadLastState loads the last known chain state from the database. This method
// assumes that the chain manager mutex is held.
func (bc *blockchain) loadLastState(number uint64) error {
	bc.logger.Info("loading last state")
	b := bc.GetBlockByNumber(number)
	if number == 0 {
		return nil
	}
	if b == nil {
		return errors.New("block is nil at load last state")
	}
	bc.currentBlock.Store(b)
	return nil
}
