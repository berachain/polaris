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

	"github.com/ethereum/go-ethereum/consensus/misc"

	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	// Prepare prepares the chain for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, int64)
	// ProcessTransaction processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*ExecutionResult, error)
	// Finalize finalizes the block and returns the block. This method is called after the last
	// tx in the block.
	Finalize(context.Context) error
	// Commit commits the current block to the blockchain and emits chain events.
	Commit(context.Context)
	// SendTx sends the given transaction to the tx pool.
	SendTx(ctx context.Context, signedTx *types.Transaction) error
}

// =========================================================================
// Block Processing
// =========================================================================

// Prepare prepares the blockchain for processing a new block at the given height.
func (bc *blockchain) Prepare(ctx context.Context, height int64) {
	bc.logger.Info("Preparing block", "height", height)

	// Prepare the Block, Configuration, Gas, and Historical plugins for the block.
	bc.bp.Prepare(ctx)
	bc.cp.Prepare(ctx)
	bc.gp.Prepare(ctx)
	if bc.hp != nil {
		bc.hp.Prepare(ctx)
	}

	// If we are processing a new block, then we assume that the previous was finalized.
	header := bc.bp.NewHeaderWithBlockNumber(height)

	// We can use the finalized block's header here because we assume that the previous block was
	// finalized.
	if parent := bc.finalizedBlock.Load(); parent != nil {
		header.BaseFee = misc.CalcBaseFee(bc.ChainConfig(), parent.Header())
	}

	bc.processor.Prepare(
		ctx,
		bc.GetEVM(ctx, vm.TxContext{}, bc.statedb, header, bc.vmConfig),
		header,
	)
}

// ProcessTransaction processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*ExecutionResult, error) {
	bc.logger.Info("Processing transaction", "tx hash", tx.Hash().Hex())

	// Reset the Gas and State plugins for the tx.
	bc.gp.Reset(ctx)
	bc.sp.Reset(ctx)

	return bc.processor.ProcessTransaction(ctx, tx)
}

// Finalize finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) error {
	block, receipts, logs, err := bc.processor.Finalize(ctx)
	if err != nil {
		return err
	}

	// store the pending logs
	bc.pendingLogsFeed.Send(logs)
	bc.currentLogs.Store(logs)

	blockHash, blockNum := block.Hash(), block.Number().Int64()
	bc.logger.Info("Finalizing block", "block", blockHash.Hex(), "num txs", len(receipts))

	// mark the current block and receipts
	if block != nil {
		bc.currentBlock.Store(block)
	}
	if receipts != nil {
		bc.currentReceipts.Store(receipts)
	}

	// store the block header on the host chain
	err = bc.bp.SetHeaderByNumber(blockNum, block.Header())
	if err != nil {
		return err
	}

	// store the block, receipts, and txs on the host chain if historical plugin is supported
	if bc.hp != nil {
		err = bc.hp.StoreBlock(block)
		if err != nil {
			return err
		}
		err = bc.hp.StoreReceipts(blockHash, receipts)
		if err != nil {
			return err
		}
		err = bc.hp.StoreTransactions(blockNum, blockHash, block.Transactions())
		return err
	}

	return nil
}

// Commit commits the current block to the blockchain and emits chain events.
func (bc *blockchain) Commit(ctx context.Context) {
	if block := bc.currentBlock.Load(); block != nil {
		// Cache finalized block.
		blockHash, blockNum := block.Hash(), block.NumberU64()
		bc.finalizedBlock.Store(block)
		bc.blockNumCache.Add(int64(blockNum), block)
		bc.blockHashCache.Add(blockHash, block)

		// Cache transaction data.
		for txIndex, tx := range block.Transactions() {
			bc.txLookupCache.Add(
				tx.Hash(),
				&types.TxLookupEntry{
					Tx:        tx,
					TxIndex:   uint64(txIndex),
					BlockNum:  blockNum,
					BlockHash: blockHash,
				},
			)
		}

		// Cache receipts.
		if receipts, ok := utils.GetAs[types.Receipts](bc.currentReceipts.Load()); ok {
			bc.receiptsCache.Add(blockHash, receipts)
		}

		// Send logs and chain events.
		if logs, ok := utils.GetAs[[]*types.Log](bc.currentLogs.Load()); ok {
			if len(logs) > 0 {
				bc.logsFeed.Send(logs)
			}
			bc.chainFeed.Send(ChainEvent{Block: block, Hash: blockHash, Logs: logs})
		}

		// Send chain head event.
		bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})
	}
}

func (bc *blockchain) SendTx(_ context.Context, signedTx *types.Transaction) error {
	return bc.tp.SendTx(signedTx)
}
