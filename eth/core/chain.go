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
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/params"
)

// By default we are storing up to 64mb of historical data for each cache.
const defaultCacheSizeBytes = 1024 * 1024 * 64

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
	gp GasPlugin
	sp StatePlugin
	tp TxPoolPlugin

	// StateProcessor is the canonical, persistent state processor that runs the EVM.
	processor *StateProcessor
	// statedb is the state database that is used to mange state during transactions.
	statedb vm.PolarisStateDB
	// vmConfig is the configuration used to create the EVM.
	vmConfig *vm.Config

	// genesisBlock is the genesis block of the blockchain.
	genesisBlock *types.Block

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
		bp:       host.GetBlockPlugin(),
		cp:       host.GetConfigurationPlugin(),
		hp:       host.GetHistoricalPlugin(),
		gp:       host.GetGasPlugin(),
		sp:       host.GetStatePlugin(),
		tp:       host.GetTxPoolPlugin(),
		vmConfig: &vm.Config{},
		genesisBlock: types.NewBlock(
			&types.Header{
				BaseFee: big.NewInt(int64(params.InitialBaseFee)), Number: big.NewInt(0),
			},
			nil, nil, nil, trie.NewStackTrie(nil),
		),
		receiptsCache:  lru.NewCache[common.Hash, types.Receipts](defaultCacheSizeBytes),
		blockNumCache:  lru.NewCache[uint64, *types.Block](defaultCacheSizeBytes),
		blockHashCache: lru.NewCache[common.Hash, *types.Block](defaultCacheSizeBytes),
		txLookupCache:  lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSizeBytes),
		chainHeadFeed:  event.Feed{},
		scope:          event.SubscriptionScope{},
		logger:         log.Root(),
	}
	bc.statedb = state.NewStateDB(bc.sp)
	bc.processor = NewStateProcessor(
		bc.cp, bc.gp, host.GetPrecompilePlugin(), bc.statedb, bc.vmConfig,
	)
	bc.currentBlock.Store(bc.genesisBlock)
	bc.finalizedBlock.Store(bc.genesisBlock)

	return bc
}

// ChainConfig returns the Ethereum chain config of the  chain.
func (bc *blockchain) Config() *params.ChainConfig {
	return bc.cp.ChainConfig()
}
