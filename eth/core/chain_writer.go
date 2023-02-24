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
