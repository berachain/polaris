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

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
)

// `CurrentHeader` retrieves the current head header of the canonical chain. The
// header is retrieved from the HeaderChain's internal cache.
func (bc *blockchain) CurrentHeader() *types.Header {
	return bc.currentHeader.Load().(*types.Header)
}

// `CurrentBlock` returns the current block of the blockchain.
func (bc *blockchain) CurrentBlock() *types.Block {
	return bc.currentBlock.Load().(*types.Block)
}

// `FinalizedBlock` returns the last finalized block of the blockchain.
func (bc *blockchain) FinalizedBlock() *types.Block {
	return bc.finalizedBlock.Load().(*types.Block)
}

func (bc *blockchain) HeaderByHash(bhash common.Hash) (*types.Header, error) {
	block, err := bc.BlockByHash(bhash)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

func (bc *blockchain) HeaderByNumber(number int64) (*types.Header, error) {
	block, err := bc.BlockByNumber(number)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

// GetBlock retrieves a block from the database by hash and number,
// caching it if found.
func (bc *blockchain) BlockByNumber(number int64) (*types.Block, error) {
	block := new(types.Block)
	if hash, ok := bc.blockNumCache.Get(number); ok {
		if block, ok = bc.blockCache.Get(hash); ok {
			if block != nil {
				return block, nil
			}
		}
	}

	// var ok bool
	// if cached, ok := bc.blockCache.Get(blockHash); ok {
	// 	return cached, nil
	// }

	// sgBlock := bc.Host().GetBlockPlugin().GetStargazerBlockByNumber(number)
	// fmt.Println("number", number)
	// fmt.Println("SGBLOCK", sgBlock)
	// if sgBlock == nil {
	// 	return nil, errors.New("BING BONG ewrror")
	// }

	header := bc.Host().GetBlockPlugin().GetStargazerHeaderByNumber(number)

	// block = sgBlock.EthBlock()

	block = types.NewBlockWithHeader(header.Header)

	// TODO GET FROM OFFCHAIN and cache
	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block, nil
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *blockchain) BlockByHash(hash common.Hash) (*types.Block, error) {
	var ok bool
	block := new(types.Block)
	if block, ok = bc.blockCache.Get(hash); ok {
		if block != nil {
			return block, nil
		}
	}

	// block := bc.Host().GetBlockPlugin().GetStargazerBlockByHash(hash)
	// if block == nil {
	// 	return nil, errors.New("BING BONG")
	// }

	// TODO GET FROM OFFCHAIN and cache
	// // Cache the found block for next time and return
	// bc.blockCache.Add(block.Hash(), block)
	return block, nil
}

func (bc *blockchain) GetReceipts(bhash common.Hash) (types.Receipts, error) {
	if cached, ok := bc.receiptsCache.Get(bhash); ok {
		return cached, nil
	}

	// TODO GET FROM OFFCHAIN and cache
	return nil, nil
}

func (bc *blockchain) GetLogs(ctx context.Context, bhash common.Hash, number uint64) ([][]*types.Log, error) {
	receipts, err := bc.GetReceipts(bhash)
	if err != nil {
		return nil, err
	}

	logs := make([][]*types.Log, len(receipts))
	for _, receipt := range receipts {
		logs = append(logs, receipt.Logs)
	}

	return logs, nil
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
		GetHash:     GetHashFn(header, bc.cc),
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
