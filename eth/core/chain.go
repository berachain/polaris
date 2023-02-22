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

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/state"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
)

// By default we are storing up to 64mb of historical data for each cache.
const defaultCacheSizeBytes = 1024 * 1024 * 64

// `blockchain` is the canonical, persistent object that operates the Stargazer EVM.
type blockchain struct {
	// `StateProcessor` is the canonical, persistent state processor that runs the EVM.
	processor *StateProcessor
	// `host` is the host chain that the Stargazer EVM is running on.
	host StargazerHostChain

	finalizedBlock atomic.Value

	// `receiptsCache` is a cache of the receipts for the last `defaultCacheSizeBytes` bytes of blocks.
	receiptsCache *lru.Cache[common.Hash, types.Receipts]
	// `blockCache` is a cache of the blocks for the last `defaultCacheSizeBytes` bytes of blocks.
	blockCache *lru.Cache[common.Hash, *types.StargazerBlock]
	// `txLookupCache` is a cache of the transactions for the last `defaultCacheSizeBytes` bytes of blocks.
	txLookupCache *lru.Cache[common.Hash, *types.Transaction]
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
		txLookupCache: lru.NewCache[common.Hash, *types.Transaction](defaultCacheSizeBytes),
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

// =========================================================================
// Block Processing
// =========================================================================

// `Prepare` prepares the blockchain for processing a new block at the given height.
func (bc *blockchain) Prepare(ctx context.Context, height int64) {
	// If we are processing a new block, then we assume that the previous was finalized.
	// TODO: ensure this is safe. We could build the block in theory by querying the blockplugin
	if bc.processor.block != nil {
		bc.finalizedBlock.Store(bc.processor.block)
		bc.blockCache.Add(bc.processor.block.Hash(), bc.processor.block)
	}
	bc.processor.Prepare(ctx, height)
}

// `ProcessTransaction` processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bc.processor.ProcessTransaction(ctx, tx)
}

// `Finalize` finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) (*types.StargazerBlock, error) {
	return bc.processor.Finalize(ctx)
}
