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

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/lib/utils"
)

// `CurrentBlock` returns the current block of the blockchain.
func (bc *blockchain) CurrentBlock() (*types.StargazerBlock, error) {
	if bc.processor.block == nil {
		return nil, errors.New("BING BONG ewrror")
	}
	bc.blockCache.Add(bc.processor.block.Hash(), bc.processor.block)
	return bc.processor.block, nil
}

// `FinalizedBlock` returns the last finalized block of the blockchain.
func (bc *blockchain) FinalizedBlock() (*types.StargazerBlock, error) {
	fb, ok := utils.GetAs[*types.StargazerBlock](bc.finalizedBlock.Load())
	if fb == nil || !ok {
		return nil, errors.New("BING BONG ewrror")
	}
	bc.blockCache.Add(fb.Hash(), fb)
	return fb, nil
}

// GetBlock retrieves a block from the database by hash and number,
// caching it if found.
func (bc *blockchain) GetStargazerBlockByNumber(number int64) (*types.StargazerBlock, error) {
	// Short circuit if the block's already in the cache, retrieve otherwise
	if bc.processor.block != nil && bc.processor.block.Number.Int64() == number {
		return bc.processor.block, nil
	}

	fp := bc.finalizedBlock.Load()
	if fp != nil {
		block, ok := fp.(*types.StargazerBlock)
		if !ok {
			return nil, errors.New("BING BONG ewrror")
		}
		if block.Number.Int64() == number {
			return block, nil
		}
	}

	block, err := bc.Host().GetBlockPlugin().GetStargazerBlockByNumber(number)
	if err != nil {
		return nil, err
	}

	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block, nil
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *blockchain) GetStargazerBlockByHash(hash common.Hash) (*types.StargazerBlock, error) {
	// Short circuit if the block's already in the cache, retrieve otherwise
	if bc.processor.block != nil && bc.processor.block.Hash() == hash {
		return bc.processor.block, nil
	}

	fp := bc.finalizedBlock.Load()
	if fp != nil {
		block, ok := fp.(*types.StargazerBlock)
		if !ok {
			return nil, errors.New("BING BONG ewrror")
		}
		if block.Hash() == hash {
			return block, nil
		}
	}

	if block, ok := bc.blockCache.Get(hash); ok {
		return block, nil
	}

	block, err := bc.Host().GetBlockPlugin().GetStargazerBlockByHash(hash)
	if block == nil {
		return nil, err
	}

	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block, nil
}

// // SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
// func (bc *blockchain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
// 	return bc.scope.Track(bc.rmLogsFeed.Subscribe(ch))
// }

// // SubscribeChainEvent registers a subscription of ChainEvent.
// func (bc *blockchain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
// 	return bc.scope.Track(bc.chainFeed.Subscribe(ch))
// }

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *blockchain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	return bc.scope.Track(bc.chainHeadFeed.Subscribe(ch))
}

// // SubscribeChainSideEvent registers a subscription of ChainSideEvent.
// func (bc *blockchain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
// 	return bc.scope.Track(bc.chainSideFeed.Subscribe(ch))
// }

// // SubscribeLogsEvent registers a subscription of []*types.Log.
// func (bc *blockchain) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
// 	return bc.scope.Track(bc.logsFeed.Subscribe(ch))
// }

// // SubscribeBlockProcessingEvent registers a subscription of bool where true means
// // block processing has started while false means it has stopped.
// func (bc *blockchain) SubscribeBlockProcessingEvent(ch chan<- bool) event.Subscription {
// 	return bc.scope.Track(bc.blockProcFeed.Subscribe(ch))
// }

func (bc *blockchain) GetStateByNumber(number int64) (vm.GethStateDB, error) {
	return bc.host.GetStatePlugin().GetStateByNumber(number)
}

func (bc *blockchain) GetEVM(ctx context.Context, txContext vm.TxContext, state vm.GethStateDB,
	header *types.Header, vmConfig *vm.Config) *vm.GethEVM {
	blockContext := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(header, &chainContext{bc.processor}),
		Coinbase:    header.Coinbase, // todo: check for fee collector
		GasLimit:    header.GasLimit,
		BlockNumber: header.Number,
		Time:        header.Time,
		Difficulty:  header.Difficulty,
		BaseFee:     header.BaseFee,
		// Random:      header.Ra,
	}

	chainCfg := bc.processor.cp.ChainConfig() // todo: get chain config at height.
	return vm.NewGethEVMWithPrecompiles(
		// todo: get precompile controller
		blockContext, txContext, state, chainCfg, *vmConfig, nil,
	)
}
