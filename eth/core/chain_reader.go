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
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/utils"
)

// ChainReader defines methods that are used to read the state and blocks of the chain.
type ChainReader interface {
	Config() *params.ChainConfig
	ChainBlockReader
	ChainTxPoolReader
	ChainSubscriber
}

// ChainBlockReader defines methods that are used to read information about blocks in the chain.
type ChainBlockReader interface {
	CurrentBlock() *types.Block
	CurrentBlockAndReceipts() (*types.Block, types.Receipts)
	FinalizedBlock() *types.Block
	GetReceiptsByHash(common.Hash) types.Receipts
	GetBlockByHash(common.Hash) *types.Block
	GetBlockByNumber(int64) *types.Block
	GetTransaction(common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
}

// ChainTxPoolReader defines methods that are used to read information about the state
// of the mempool.
type ChainTxPoolReader interface {
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(common.Hash) *types.Transaction
	GetPoolNonce(common.Address) (uint64, error)
	GetPoolStats() (int, int)
	GetPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	GetPoolContentFrom(addr common.Address) (types.Transactions, types.Transactions)
}

// =========================================================================
// Configuration
// =========================================================================

// ChainConfig returns the Ethereum chain config of the  chain.
func (bc *blockchain) Config() *params.ChainConfig {
	return bc.cp.ChainConfig()
}

// =========================================================================
// BlockReader
// =========================================================================

// CurrentHeader returns the current header of the blockchain.
func (bc *blockchain) CurrentBlock() *types.Block {
	block, ok := utils.GetAs[*types.Block](bc.currentBlock.Load())
	if block == nil || !ok {
		return nil
	}
	bc.blockNumCache.Add(block.Number().Int64(), block)
	bc.blockHashCache.Add(block.Hash(), block)
	return block
}

// CurrentReceipts returns the current receipts of the blockchain.
func (bc *blockchain) CurrentBlockAndReceipts() (*types.Block, types.Receipts) {
	var err error

	// Get current block.
	block := bc.CurrentBlock()
	if block == nil {
		bc.logger.Error("current block is nil")
		return nil, nil
	}

	// Get receipts from cache.
	receipts, ok := utils.GetAs[types.Receipts](bc.currentReceipts.Load())
	if receipts == nil || !ok {
		bc.logger.Error("current receipts are nil")
		return nil, nil
	}

	// Derive receipts from block.
	receipts, err = bc.deriveReceipts(receipts, block.Hash())
	if err != nil {
		bc.logger.Error("failed to derive receipts", "err", err)
		return nil, nil
	}

	// Add to cache.
	bc.receiptsCache.Add(block.Hash(), receipts)
	return block, receipts
}

// FinalizedBlock returns the last finalized block of the blockchain.
func (bc *blockchain) FinalizedBlock() *types.Block {
	fb, ok := utils.GetAs[*types.Block](bc.finalizedBlock.Load())
	if fb == nil || !ok {
		return nil
	}
	bc.blockNumCache.Add(fb.Number().Int64(), fb)
	bc.blockHashCache.Add(fb.Hash(), fb)
	return fb
}

// GetReceipts gathers the receipts that were created in the block defined by
// the given hash.
func (bc *blockchain) GetReceiptsByHash(blockHash common.Hash) types.Receipts {
	// check the cache
	if receipts, ok := bc.receiptsCache.Get(blockHash); ok {
		derived, err := bc.deriveReceipts(receipts, blockHash)
		if err != nil {
			bc.logger.Error("failed to derive receipts", "err", err)
			return nil
		}
		return derived
	}

	// check if historical plugin is supported by host chain
	if bc.hp == nil {
		bc.logger.Debug("historical plugin not supported by host chain")
		return nil
	}

	// check the historical plugin
	receipts, err := bc.hp.GetReceiptsByHash(blockHash)
	if err != nil {
		bc.logger.Error("failed to get receipts from historical plugin", "err", err)
		return nil
	}

	// cache the found receipts for next time and return
	bc.receiptsCache.Add(blockHash, receipts)
	derived, err := bc.deriveReceipts(receipts, blockHash)
	if err != nil {
		bc.logger.Error("failed to derive receipts", "err", err)
		return nil
	}
	return derived
}

// GetTransaction gets a transaction by hash. It also returns the block hash of the
// block that the transaction was included in, the block number, and the index of the
// transaction in the block. It only retrieves transactions that are included in the chain
// and does not acquire transactions that are in the mempool.
func (bc *blockchain) GetTransaction(
	txHash common.Hash,
) (*types.Transaction, common.Hash, uint64, uint64, error) {
	// check the cache
	if txLookupEntry, ok := bc.txLookupCache.Get(txHash); ok {
		return txLookupEntry.Tx, txLookupEntry.BlockHash,
			txLookupEntry.BlockNum, txLookupEntry.TxIndex, nil
	}

	// check if historical plugin is supported by host chain
	if bc.hp == nil {
		bc.logger.Debug("historical plugin not supported by host chain")
		return nil, common.Hash{}, 0, 0, ErrTxNotFound
	}

	// check the historical plugin
	txLookupEntry, err := bc.hp.GetTransactionByHash(txHash)
	if err != nil {
		return nil, common.Hash{}, 0, 0, err
	}

	// cache the found transaction for next time and return
	bc.txLookupCache.Add(txHash, txLookupEntry)
	return txLookupEntry.Tx, txLookupEntry.BlockHash,
		txLookupEntry.BlockNum, txLookupEntry.TxIndex, nil
}

// GetBlock retrieves a block from the database by hash and number, caching it if found.
func (bc *blockchain) GetBlockByNumber(number int64) *types.Block {
	// check the block number cache
	if block, ok := bc.blockNumCache.Get(number); ok {
		bc.blockHashCache.Add(block.Hash(), block)
		return block
	}

	// check if historical plugin is supported by host chain
	if bc.hp == nil {
		bc.logger.Debug("historical plugin not supported by host chain")
		return nil
	}

	// check the historical plugin
	block, err := bc.hp.GetBlockByNumber(number)
	if err != nil {
		return nil
	}

	// Cache the found block for next time and return
	bc.blockNumCache.Add(number, block)
	bc.blockHashCache.Add(block.Hash(), block)
	return block
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *blockchain) GetBlockByHash(hash common.Hash) *types.Block {
	// check the block hash cache
	if block, ok := bc.blockHashCache.Get(hash); ok {
		bc.blockNumCache.Add(block.Number().Int64(), block)
		return block
	}

	// check if historical plugin is supported by host chain
	if bc.hp == nil {
		bc.logger.Debug("historical plugin not supported by host chain")
		return nil
	}

	// check the historical plugin
	block, err := bc.hp.GetBlockByHash(hash)
	if err != nil {
		bc.logger.Error("failed to get block by hash", "err", err)
		return nil
	}

	// Cache the found block for next time and return
	bc.blockNumCache.Add(block.Number().Int64(), block)
	bc.blockHashCache.Add(hash, block)
	return block
}

// =========================================================================
// TransactionPoolReader
// =========================================================================

// GetPoolTransactions returns all of the transactions that are currently in
// the mempool.
func (bc *blockchain) GetPoolTransactions() (types.Transactions, error) {
	pending := bc.tp.Pending(false)
	txs := make(types.Transactions, len(pending))
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

// GetPoolTransaction returns a transaction from the mempool by hash.
func (bc *blockchain) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return bc.tp.Get(hash)
}

// GetPoolNonce returns the pending nonce of addr from the mempool.
func (bc *blockchain) GetPoolNonce(addr common.Address) (uint64, error) {
	return bc.tp.Nonce(addr), bc.statedb.Error()
}

// GetPoolStats returns the number of pending and queued txs in the mempool.
func (bc *blockchain) GetPoolStats() (int, int) {
	return bc.tp.Stats()
}

// GetPoolContent returns the pending and queued txs in the mempool.
func (bc *blockchain) GetPoolContent() (
	map[common.Address]types.Transactions, map[common.Address]types.Transactions,
) {
	return bc.tp.Content()
}

// GetPoolContentFrom returns the pending and queued txs in the mempool for the given address.
func (bc *blockchain) GetPoolContentFrom(addr common.Address) (
	types.Transactions, types.Transactions,
) {
	return bc.tp.ContentFrom(addr)
}
