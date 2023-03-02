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
	"sync/atomic"

	lru "github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/state"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/params"
)

// By default we are storing up to 64mb of historical data for each cache.
const defaultCacheSizeBytes = 1024 * 1024 * 64

// `ChainReaderWriter` is the interface that wraps the basic methods of the EVM chain.
type ChainReaderWriter interface {
	ChainWriter
	ChainReader
}

// `ChainWriter` defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	// `Prepare` prepares the chain for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, int64)
	// `ProcessTransaction` processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*ExecutionResult, error)
	// `Finalize` finalizes the block and returns the block. This method is called after the last
	// tx in the block.
	Finalize(context.Context) (*types.Block, types.Receipts, error)

	// `SendTx` sends the given transaction to the tx pool.
	SendTx(ctx context.Context, signedTx *types.Transaction) error
}

// `ChainReader` defines methods that are used to read the state and blocks of the chain.
type ChainReader interface {
	CurrentBlock() (*types.Block, error)
	CurrentBlockAndReceipts() (*types.Block, types.Receipts, error)
	FinalizedBlock() (*types.Block, error)
	GetReceipts(common.Hash) (types.Receipts, error)
	GetTransaction(common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
	GetStargazerBlockByHash(common.Hash) (*types.Block, error)
	GetStargazerBlockByNumber(int64) (*types.Block, error)
	GetStateByNumber(int64) (vm.GethStateDB, error)
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	GetEVM(context.Context, vm.TxContext, vm.StargazerStateDB, *types.Header, *vm.Config) *vm.GethEVM
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(common.Hash) *types.Transaction
	GetPoolNonce(common.Address) (uint64, error)
	ChainConfig() *params.ChainConfig
}

// Compile-time check to ensure that `blockchain` implements the `ChainReaderWriter` interface.
var _ ChainReaderWriter = (*blockchain)(nil)

// `blockchain` is the canonical, persistent object that operates the Stargazer EVM.
type blockchain struct {
	// `host` is the host chain that the Stargazer EVM is running on.
	host StargazerHostChain
	// `StateProcessor` is the canonical, persistent state processor that runs the EVM.
	processor *StateProcessor
	// `statedb` is the state database that is used to mange state during transactions.
	statedb vm.StargazerStateDB
	// vmConfig is the configuration used to create the EVM.
	vmConfig *vm.Config

	// `currentBlock` is the current/pending block.
	currentBlock atomic.Value
	// `finalizedBlock` is the finalized/latest block.
	finalizedBlock atomic.Value
	// `currentReceipts` is the current/pending receipts.
	currentReceipts atomic.Value

	// `receiptsCache` is a cache of the receipts for the last `defaultCacheSizeBytes` bytes of
	// blocks. blockHash -> receipts
	receiptsCache *lru.Cache[common.Hash, types.Receipts]
	// `blockNumCache` is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockNum -> block
	blockNumCache *lru.Cache[int64, *types.Block]
	// `blockHashCache` is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	// blockHash -> block
	blockHashCache *lru.Cache[common.Hash, *types.Block]
	// `txLookupCache` is a cache of the transactions for the last `defaultCacheSizeBytes` bytes of
	// blocks. txHash -> txLookupEntry
	txLookupCache *lru.Cache[common.Hash, *types.TxLookupEntry]

	cc            ChainContext
	chainHeadFeed event.Feed
	scope         event.SubscriptionScope
}

// =========================================================================
// Constructor
// =========================================================================

// `NewChain` creates and returns a `api.Chain` with the given EVM chain configuration and host.
func NewChain(host StargazerHostChain) *blockchain { //nolint:revive // temp.
	bc := &blockchain{
		host:           host,
		statedb:        state.NewStateDB(host.GetStatePlugin()),
		vmConfig:       &vm.Config{},
		receiptsCache:  lru.NewCache[common.Hash, types.Receipts](defaultCacheSizeBytes),
		blockNumCache:  lru.NewCache[int64, *types.Block](defaultCacheSizeBytes),
		blockHashCache: lru.NewCache[common.Hash, *types.Block](defaultCacheSizeBytes),
		txLookupCache:  lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSizeBytes),
		chainHeadFeed:  event.Feed{},
		scope:          event.SubscriptionScope{},
	}
	bc.cc = &chainContext{bc}
	bc.processor = NewStateProcessor(bc.host, bc.statedb, bc.vmConfig)
	return bc
}
