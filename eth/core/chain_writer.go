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
	"fmt"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/trie"
	"pkg.berachain.dev/stargazer/eth/core/types"
)

// =========================================================================
// Block Processing
// =========================================================================

// `Prepare` prepares the blockchain for processing a new block.
func (bc *blockchain) Prepare(ctx context.Context, number int64) {
	// Finalize the previous block and load it into the cache.
	block, ok := bc.currentBlock.Load().(*types.Block)
	fmt.Println("PREPARE TOP:")
	if ok {
		bc.finalizedBlock.Store(block)
		// Cache transaction data
		fmt.Println("PREPARE TX:", block.Transactions())
		for i, tx := range block.Transactions() {
			bc.txLookupCache.Add(tx.Hash(), &rawdb.LegacyTxLookupEntry{
				BlockHash:  block.Hash(),
				BlockIndex: block.Number().Uint64(),
				Index:      uint64(i),
			})
		}
		// We also add to the block cache.
		bc.blockCache.Add(block.Hash(), block)
		fmt.Println("PREPARE RECE:", bc.processor.receipts)
		bc.receiptsCache.Add(block.Hash(), bc.processor.receipts)
	}

	// Prepare the plugins for the new block.
	bp := bc.host.GetBlockPlugin()
	bp.Prepare(ctx)

	header, err := bc.HeaderByNumber(number)
	if err != nil {
		fmt.Println("Error getting header by number", err)
		panic("ERR")
	}
	bc.currentHeader.Store(header)

	// We wipe the last current block to ensure no state is carried over from the
	// previous cycle.
	block = types.NewBlock(header, nil, nil, nil, trie.NewStackTrie(nil))
	bc.currentBlock.Store(block)

	bc.processor.Prepare(ctx, bc.cc, header)
}

// `ProcessTransaction` processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bc.processor.ProcessTransaction(ctx, tx)
}

// `Finalize` finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) (*types.Header, types.Transactions, types.Receipts) {
	// We create a new block from the output of the state processor.
	header, txs, receipts := bc.processor.Finalize(ctx)
	block := types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil))
	fmt.Println("FINAL BLOCK", block)
	// Now that we've executed the block, we then then say its been added to the top of the chain.
	bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})
	bc.currentBlock.Store(block)
	bc.currentHeader.Store(header)
	return header, txs, receipts
}
