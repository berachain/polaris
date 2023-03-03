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

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/state"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/lib/utils"
)

// `ChainReader` defines methods that are used to read the state and blocks of the chain.
type ChainReader interface {
	ChainBlockReader
	ChainTxPoolReader
	ChainSubscriber
	GetStateByNumber(int64) (vm.GethStateDB, error)
	GetEVM(context.Context, vm.TxContext, vm.StargazerStateDB, *types.Header, *vm.Config) *vm.GethEVM
	ChainConfig() *params.ChainConfig
}

type ChainBlockReader interface {
	CurrentBlock() (*types.Block, error)
	CurrentBlockAndReceipts() (*types.Block, types.Receipts, error)
	FinalizedBlock() (*types.Block, error)
	GetReceipts(common.Hash) (types.Receipts, error)
	GetStargazerBlockByHash(common.Hash) (*types.Block, error)
	GetStargazerBlockByNumber(int64) (*types.Block, error)
	GetTransaction(common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
}

type ChainTxPoolReader interface {
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(common.Hash) *types.Transaction
	GetPoolNonce(common.Address) (uint64, error)
}

// =========================================================================
// BlockReader
// =========================================================================

// `CurrentBlock` returns the current block of the blockchain.
func (bc *blockchain) CurrentBlock() (*types.Block, error) {
	cb, ok := utils.GetAs[*types.Block](bc.currentBlock.Load())
	if cb == nil || !ok {
		return nil, errors.New("current block cannot be loaded from cache")
	}
	bc.blockNumCache.Add(cb.Number().Int64(), cb)
	bc.blockHashCache.Add(cb.Hash(), cb)
	return cb, nil
}

// `CurrentReceipts` returns the current receipts of the blockchain.
func (bc *blockchain) CurrentBlockAndReceipts() (*types.Block, types.Receipts, error) {
	cb, err := bc.CurrentBlock()
	if err != nil {
		return nil, nil, err
	}
	cr, ok := utils.GetAs[types.Receipts](bc.currentReceipts.Load())
	if cb == nil || !ok {
		return nil, nil, errors.New("current receipts cannot be loaded from cache")
	}
	bc.receiptsCache.Add(cb.Hash(), cr)
	return cb, cr, nil
}

// `FinalizedBlock` returns the last finalized block of the blockchain.
func (bc *blockchain) FinalizedBlock() (*types.Block, error) {
	fb, ok := utils.GetAs[*types.Block](bc.finalizedBlock.Load())
	if fb == nil || !ok {
		return nil, errors.New("finalized block cannot be loaded from cache")
	}
	bc.blockNumCache.Add(fb.Number().Int64(), fb)
	bc.blockHashCache.Add(fb.Hash(), fb)
	return fb, nil
}

func (bc *blockchain) GetReceipts(blockHash common.Hash) (types.Receipts, error) {
	// check the cache
	if receipts, ok := bc.receiptsCache.Get(blockHash); ok {
		return receipts, nil
	}

	// check the block plugin
	receipts, err := bc.host.GetBlockPlugin().GetReceiptsByHash(blockHash)
	if err != nil {
		return nil, err
	}

	// cache the found receipts for next time and return
	bc.receiptsCache.Add(blockHash, receipts)
	return nil, nil
}

func (bc *blockchain) GetTransaction(
	txHash common.Hash,
) (*types.Transaction, common.Hash, uint64, uint64, error) {
	// check the cache
	if txLookupEntry, ok := bc.txLookupCache.Get(txHash); ok {
		return txLookupEntry.Tx, txLookupEntry.BlockHash,
			txLookupEntry.BlockNum, txLookupEntry.TxIndex, nil
	}

	// check the block plugin
	txLookupEntry, err := bc.host.GetBlockPlugin().GetTransactionByHash(txHash)
	if err != nil {
		return nil, common.Hash{}, 0, 0, err
	}

	// cache the found transaction for next time and return
	bc.txLookupCache.Add(txHash, txLookupEntry)
	return txLookupEntry.Tx, txLookupEntry.BlockHash,
		txLookupEntry.BlockNum, txLookupEntry.TxIndex, nil
}

// GetBlock retrieves a block from the database by hash and number, caching it if found.
func (bc *blockchain) GetStargazerBlockByNumber(number int64) (*types.Block, error) {
	// check the block number cache
	if block, ok := bc.blockNumCache.Get(number); ok {
		bc.blockHashCache.Add(block.Hash(), block)
		return block, nil
	}

	// check the block plugin
	block, err := bc.host.GetBlockPlugin().GetBlockByNumber(number)
	if err != nil {
		return nil, err
	}

	// Cache the found block for next time and return
	bc.blockNumCache.Add(block.Number().Int64(), block)
	bc.blockHashCache.Add(block.Hash(), block)
	return block, nil
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *blockchain) GetStargazerBlockByHash(hash common.Hash) (*types.Block, error) {
	// check the block hash cache
	if block, ok := bc.blockHashCache.Get(hash); ok {
		bc.blockNumCache.Add(block.Number().Int64(), block)
		return block, nil
	}

	// check the block plugin
	block, err := bc.host.GetBlockPlugin().GetBlockByHash(hash)
	if err != nil {
		return nil, err
	}

	// Cache the found block for next time and return
	bc.blockNumCache.Add(block.Number().Int64(), block)
	bc.blockHashCache.Add(block.Hash(), block)
	return block, nil
}

func (bc *blockchain) GetStateByNumber(number int64) (vm.GethStateDB, error) {
	sp, err := bc.host.GetStatePlugin().GetStateByNumber(number)
	if err != nil {
		return nil, err
	}
	return state.NewStateDB(sp), nil
}

func (bc *blockchain) GetEVM(
	_ context.Context, txContext vm.TxContext, state vm.StargazerStateDB,
	header *types.Header, vmConfig *vm.Config,
) *vm.GethEVM {
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
	}

	chainCfg := bc.processor.cp.ChainConfig() // todo: get chain config at height.
	return vm.NewGethEVMWithPrecompiles(
		blockContext, txContext, state, chainCfg, *vmConfig, bc.processor.pp,
	)
}

// =========================================================================
// TransactionPoolReader
// =========================================================================

func (bc *blockchain) GetPoolTransactions() (types.Transactions, error) {
	return bc.host.GetTxPoolPlugin().GetAllTransactions()
}

func (bc *blockchain) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return bc.host.GetTxPoolPlugin().GetTransaction(hash)
}

func (bc *blockchain) GetPoolNonce(addr common.Address) (uint64, error) {
	return bc.host.GetTxPoolPlugin().GetNonce(addr)
}

func (bc *blockchain) ChainConfig() *params.ChainConfig {
	return bc.host.GetConfigurationPlugin().ChainConfig()
}
