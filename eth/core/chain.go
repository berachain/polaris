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
