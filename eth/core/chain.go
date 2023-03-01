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
	ProcessTransaction(context.Context, *types.Transaction) (*types.Receipt, error)
	// `Finalize` finalizes the block and returns the block. This method is called after the last
	// tx in the block.
	Finalize(context.Context) (*types.StargazerBlock, error)
}

// `ChainReader` defines methods that are used to read the state and blocks of the chain.
type ChainReader interface {
	CurrentHeader() *types.StargazerHeader
	CurrentBlock() *types.StargazerBlock
	FinalizedBlock() *types.StargazerBlock
	GetStargazerBlockByHash(common.Hash) *types.StargazerBlock
	GetStargazerBlockByNumber(int64) *types.StargazerBlock
	GetStateByNumber(int64) (vm.GethStateDB, error)
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	GetEVM(context.Context, vm.TxContext, vm.GethStateDB, *types.Header, *vm.Config) *vm.GethEVM
	GetTransaction(txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
}

// Compile-time check to ensure that `blockchain` implements the `ChainReaderWriter` interface.
var _ ChainReaderWriter = (*blockchain)(nil)

// `blockchain` is the canonical, persistent object that operates the Stargazer EVM.
type blockchain struct {
	// `StateProcessor` is the canonical, persistent state processor that runs the EVM.
	processor *StateProcessor
	// `host` is the host chain that the Stargazer EVM is running on.
	host StargazerHostChain

	// `finalizedBlock` is the last finalized block.
	finalizedBlock atomic.Value

	// `receiptsCache` is a cache of the receipts for the last `defaultCacheSizeBytes` bytes of blocks.
	receiptsCache *lru.Cache[common.Hash, types.Receipts]
	// `blockCache` is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	blockCache *lru.Cache[common.Hash, *types.StargazerBlock]
	// `txLookupCache` is a cache of the transactions for the last `defaultCacheSizeBytes` bytes of blocks.
	txLookupCache *lru.Cache[common.Hash, *types.TxLookupEntry]

	chainHeadFeed event.Feed
	scope         event.SubscriptionScope
}

// =========================================================================
// Constructor
// =========================================================================

// `NewChain` creates and returns a `api.Chain` with the given EVM chain configuration and host.
func NewChain(host StargazerHostChain) *blockchain { //nolint:revive // temp.
	bc := &blockchain{
		host:          host,
		receiptsCache: lru.NewCache[common.Hash, types.Receipts](defaultCacheSizeBytes),
		blockCache:    lru.NewCache[common.Hash, *types.StargazerBlock](defaultCacheSizeBytes),
		txLookupCache: lru.NewCache[common.Hash, *types.TxLookupEntry](defaultCacheSizeBytes),
		chainHeadFeed: event.Feed{},
		scope:         event.SubscriptionScope{},
	}
	bc.processor = bc.buildStateProcessor(vm.Config{}, true)
	return bc
}

// `Host` returns the host chain that the Stargazer EVM is running on.
func (bc *blockchain) Host() StargazerHostChain {
	return bc.host
}

// `buildStateProcessor` builds and returns a `StateProcessor` with the given EVM configuration and
// commit flag.
func (bc *blockchain) buildStateProcessor(vmConfig vm.Config, commit bool) *StateProcessor {
	return NewStateProcessor(bc.host, state.NewStateDB(bc.host.GetStatePlugin()), vmConfig, commit)
}
