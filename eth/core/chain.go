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

package core

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/berachain/polaris/eth/consensus"
	"github.com/berachain/polaris/eth/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

// By default we are storing up to 512 items in each cache.
const defaultCacheSize = 512

// Compile-time check to ensure that `blockchain` implements the `Blockchain` api.
var _ Blockchain = (*blockchain)(nil)

// Blockchain interface defines the methods that a blockchain must have.
type Blockchain interface {
	ChainReader
	ChainWriter
	ChainSubscriber
	ChainResources
	core.ChainContext

	PrimePlugins(ctx context.Context)
	StatePluginFactory() StatePluginFactory
}

// blockchain is the canonical, persistent object that operates the Polaris EVM.
type blockchain struct {
	// the host chain plugins that the Polaris EVM is running on.
	bp  BlockPlugin
	hp  HistoricalPlugin
	pp  PrecompilePlugin
	spf StatePluginFactory

	engine    consensus.Engine
	processor core.Processor
	validator core.Validator

	// vmConfig is the configuration used to create the EVM.
	vmConfig *vm.Config

	// config represents the chain config.
	config *params.ChainConfig

	// currentBlock is the current/pending block.
	currentBlock atomic.Pointer[ethtypes.Block]
	// finalizedBlock is the finalized/latest block.
	finalizedBlock atomic.Pointer[ethtypes.Block]

	// receiptsCache is a cache of the receipts for the last `defaultCacheSizeBytes` bytes of
	// blocks. blockHash -> receipts
	receiptsCache *lru.Cache[common.Hash, ethtypes.Receipts]
	// blockNumCache is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockNum -> block
	blockNumCache *lru.Cache[uint64, *ethtypes.Block]
	// blockHashCache is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockHash -> block
	blockHashCache *lru.Cache[common.Hash, *ethtypes.Block]
	// txLookupCache is a cache of the transactions for the last `defaultCacheSizeBytes` bytes of
	// blocks. txHash -> txLookupEntry
	txLookupCache *lru.Cache[common.Hash, *types.TxLookupEntry]

	// subscription event feeds
	scope         event.SubscriptionScope
	chainFeed     event.Feed
	chainHeadFeed event.Feed
	logsFeed      event.Feed
	rmLogsFeed    event.Feed // currently never used
	chainSideFeed event.Feed // currently never used
	logger        log.Logger
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
		spf:            host.GetStatePluginFactory(),
		config:         config,
		vmConfig:       &vm.Config{},
		receiptsCache:  lru.NewCache[common.Hash, ethtypes.Receipts](defaultCacheSize),
		blockNumCache:  lru.NewCache[uint64, *ethtypes.Block](defaultCacheSize),
		blockHashCache: lru.NewCache[common.Hash, *ethtypes.Block](defaultCacheSize),
		txLookupCache:  lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSize),
		chainHeadFeed:  event.Feed{},
		scope:          event.SubscriptionScope{},
		logger:         log.Root(),
		engine:         engine,
	}
	bc.processor = core.NewStateProcessor(bc.config, bc, bc.engine)
	bc.validator = core.NewBlockValidator(bc.config, bc, bc.engine)
	// TODO: bug fix required.
	bc.currentBlock.Store(
		ethtypes.NewBlock(&ethtypes.Header{Time: 0, Number: big.NewInt(0),
			BaseFee: big.NewInt(0)}, nil, nil, nil, trie.NewStackTrie(nil)))
	bc.finalizedBlock.Store(nil)
	return bc
}

func (bc *blockchain) PrimePlugins(ctx context.Context) {
	if bc.bp != nil {
		bc.bp.Prepare(ctx)
	}
	if bc.hp != nil {
		bc.hp.Prepare(ctx)
	}
}

// LoadLastState loads the last known chain state from the database. This method
// assumes that the chain manager mutex is held.
func (bc *blockchain) LoadLastState(number uint64) error {
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

func (bc *blockchain) StatePluginFactory() StatePluginFactory {
	return bc.spf
}
