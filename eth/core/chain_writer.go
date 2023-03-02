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

	"github.com/ethereum/go-ethereum/core/vm"

	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/utils"
)

// =========================================================================
// Block Processing
// =========================================================================

// `Prepare` prepares the blockchain for processing a new block at the given height.
func (bc *blockchain) Prepare(ctx context.Context, height int64) {
	bc.host.GetBlockPlugin().Prepare(ctx)
	bc.host.GetGasPlugin().Prepare(ctx)
	bc.host.GetConfigurationPlugin().Prepare(ctx)

	// If we are processing a new block, then we assume that the previous was finalized.
	if block, ok := utils.GetAs[*types.Block](bc.currentBlock.Load()); ok {
		// Cache finalized block.
		blockHash, blockNum := block.Hash(), block.NumberU64()
		bc.finalizedBlock.Store(block)
		bc.blockNumCache.Add(int64(blockNum), block)
		bc.blockHashCache.Add(blockHash, block)

		// Cache transaction data
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
		var receipts types.Receipts
		if receipts, ok = utils.GetAs[types.Receipts](bc.currentReceipts.Load()); ok {
			bc.receiptsCache.Add(blockHash, receipts)
		}

		// TODO: synchronize chain head feed.
		bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})
	}

	header := bc.host.GetBlockPlugin().NewHeaderWithBlockNumber(height)
	bc.processor.Prepare(
		ctx,
		bc.GetEVM(ctx, vm.TxContext{}, bc.statedb, header, bc.vmConfig),
		header,
	)
}

// `ProcessTransaction` processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*ExecutionResult, error) {
	return bc.processor.ProcessTransaction(ctx, tx)
}

// `Finalize` finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) (*types.Block, types.Receipts, error) {
	block, receipts, err := bc.processor.Finalize(ctx)
	if block != nil {
		bc.currentBlock.Store(block)
	}
	if receipts != nil {
		bc.currentReceipts.Store(receipts)
	}
	return block, receipts, err
}

func (bc *blockchain) SendTx(_ context.Context, signedTx *types.Transaction) error {
	return bc.host.GetTxPoolPlugin().SendTx(signedTx)
}
