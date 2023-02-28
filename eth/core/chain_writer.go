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

	"pkg.berachain.dev/stargazer/eth/core/types"
)

// =========================================================================
// Block Processing
// =========================================================================

// `Prepare` prepares the blockchain for processing a new block at the given height.
func (bc *blockchain) Prepare(ctx context.Context, height int64) {
	// If we are processing a new block, then we assume that the previous was finalized.
	// TODO: ensure this is safe. We could build the block in theory by querying the blockplugin
	if bc.processor.block != nil {
		// Cache finalized block.
		bc.finalizedBlock.Store(bc.processor.block)
		bc.blockCache.Add(bc.processor.block.Hash(), bc.processor.block)

		// Cache transaction data
		for _, tx := range bc.processor.block.GetTransactions() {
			bc.txLookupCache.Add(tx.Hash(), tx)
		}
		// Cache receipts.
		bc.receiptsCache.Add(bc.processor.block.Hash(), bc.processor.block.GetReceipts())
	}

	// Prepare the state processor for the next block.
	bc.processor.Prepare(ctx, height)
	bc.chainHeadFeed.Send(ChainHeadEvent{Block: bc.processor.block.EthBlock()})
}

// `ProcessTransaction` processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bc.processor.ProcessTransaction(ctx, tx)
}

// `Finalize` finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) (*types.StargazerBlock, error) {
	return bc.processor.Finalize(ctx)
}
