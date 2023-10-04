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
	"math/big"
	"sync/atomic"

	lru "github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/eth/common"
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

type Blockchain interface {
	Config() *params.ChainConfig
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
	cp ConfigurationPlugin
	hp HistoricalPlugin
	pp PrecompilePlugin
	sp StatePlugin

	processor core.Processor

	// statedb is the state database that is used to mange state during transactions.
	statedb vm.PolarisStateDB
	// vmConfig is the configuration used to create the EVM.
	vmConfig *vm.Config

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
func NewChain(host PolarisHostChain) *blockchain { //nolint:revive // only used as `api.Chain`.
	bc := &blockchain{
		bp:             host.GetBlockPlugin(),
		cp:             host.GetConfigurationPlugin(),
		hp:             host.GetHistoricalPlugin(),
		pp:             host.GetPrecompilePlugin(),
		sp:             host.GetStatePlugin(),
		vmConfig:       &vm.Config{},
		receiptsCache:  lru.NewCache[common.Hash, types.Receipts](defaultCacheSize),
		blockNumCache:  lru.NewCache[uint64, *types.Block](defaultCacheSize),
		blockHashCache: lru.NewCache[common.Hash, *types.Block](defaultCacheSize),
		txLookupCache:  lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSize),
		chainHeadFeed:  event.Feed{},
		scope:          event.SubscriptionScope{},
		logger:         log.Root(),
	}
	bc.statedb = state.NewStateDB(bc.sp, bc.pp)
	bc.processor = core.NewStateProcessor(bc.cp.ChainConfig(), bc, beacon.New(nil))
	// TODO: hmm...
	bc.currentBlock.Store(
		types.NewBlock(&types.Header{Number: big.NewInt(0),
			BaseFee: big.NewInt(0)}, nil, nil, nil, trie.NewStackTrie(nil)))
	// bc.currentBlock.Store(nil)
	bc.finalizedBlock.Store(nil)

	if err := bc.loadLastState(); err != nil {
		panic(err)
	}
	return bc
}

// ChainConfig returns the Ethereum chain config of the  chain.
func (bc *blockchain) Config() *params.ChainConfig {
	return bc.cp.ChainConfig()
}

// loadLastState loads the last known chain state from the database. This method
// assumes that the chain manager mutex is held.
func (bc *blockchain) loadLastState() error {
	return nil
	// // Restore the last known head block
	// head := rawdb.ReadHeadBlockHash(bc.db)
	// if head == (common.Hash{}) {
	// 	// Corrupt or empty database, init from scratch
	// 	log.Warn("Empty database, resetting chain")
	// 	return bc.Reset()
	// }
	// // Make sure the entire head block is available
	// headBlock := bc.GetBlockByHash(head)
	// if headBlock == nil {
	// 	// Corrupt or empty database, init from scratch
	// 	log.Warn("Head block missing, resetting chain", "hash", head)
	// 	return bc.Reset()
	// }
	// // Everything seems to be fine, set as the head block
	// bc.currentBlock.Store(headBlock.Header())
	// headBlockGauge.Update(int64(headBlock.NumberU64()))

	// // Restore the last known head header
	// headHeader := headBlock.Header()
	// if head := rawdb.ReadHeadHeaderHash(bc.db); head != (common.Hash{}) {
	// 	if header := bc.GetHeaderByHash(head); header != nil {
	// 		headHeader = header
	// 	}
	// }
	// bc.hc.SetCurrentHeader(headHeader)

	// // Restore the last known head snap block
	// bc.currentSnapBlock.Store(headBlock.Header())
	// headFastBlockGauge.Update(int64(headBlock.NumberU64()))

	// if head := rawdb.ReadHeadFastBlockHash(bc.db); head != (common.Hash{}) {
	// 	if block := bc.GetBlockByHash(head); block != nil {
	// 		bc.currentSnapBlock.Store(block.Header())
	// 		headFastBlockGauge.Update(int64(block.NumberU64()))
	// 	}
	// }

	// // Restore the last known finalized block and safe block
	// // Note: the safe block is not stored on disk and it is set to the last
	// // known finalized block on startup
	// if head := rawdb.ReadFinalizedBlockHash(bc.db); head != (common.Hash{}) {
	// 	if block := bc.GetBlockByHash(head); block != nil {
	// 		bc.currentFinalBlock.Store(block.Header())
	// 		headFinalizedBlockGauge.Update(int64(block.NumberU64()))
	// 		bc.currentSafeBlock.Store(block.Header())
	// 		headSafeBlockGauge.Update(int64(block.NumberU64()))
	// 	}
	// }
	// // Issue a status log for the user
	// var (
	// 	currentSnapBlock  = bc.CurrentSnapBlock()
	// 	currentFinalBlock = bc.CurrentFinalBlock()

	// 	headerTd = bc.GetTd(headHeader.Hash(), headHeader.Number.Uint64())
	// 	blockTd  = bc.GetTd(headBlock.Hash(), headBlock.NumberU64())
	// )
	// if headHeader.Hash() != headBlock.Hash() {
	// 	log.Info("Loaded most recent local header", "number", headHeader.Number, "hash", headHeader.Hash(), "td", headerTd, "age", common.PrettyAge(time.Unix(int64(headHeader.Time), 0)))
	// }
	// log.Info("Loaded most recent local block", "number", headBlock.Number(), "hash", headBlock.Hash(), "td", blockTd, "age", common.PrettyAge(time.Unix(int64(headBlock.Time()), 0)))
	// if headBlock.Hash() != currentSnapBlock.Hash() {
	// 	snapTd := bc.GetTd(currentSnapBlock.Hash(), currentSnapBlock.Number.Uint64())
	// 	log.Info("Loaded most recent local snap block", "number", currentSnapBlock.Number, "hash", currentSnapBlock.Hash(), "td", snapTd, "age", common.PrettyAge(time.Unix(int64(currentSnapBlock.Time), 0)))
	// }
	// if currentFinalBlock != nil {
	// 	finalTd := bc.GetTd(currentFinalBlock.Hash(), currentFinalBlock.Number.Uint64())
	// 	log.Info("Loaded most recent local finalized block", "number", currentFinalBlock.Number, "hash", currentFinalBlock.Hash(), "td", finalTd, "age", common.PrettyAge(time.Unix(int64(currentFinalBlock.Time), 0)))
	// }
	// if pivot := rawdb.ReadLastPivotNumber(bc.db); pivot != nil {
	// 	log.Info("Loaded last snap-sync pivot marker", "number", *pivot)
	// }
	// return nil
}
