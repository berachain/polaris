package chain

import (
	"context"

	"pkg.berachain.dev/polaris/eth/api"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// `blockProducer` is the block producer.
type blockProducer struct {
	polaris         api.Chain
	currentBlockNum int64
}

// `ProduceBlock` produces a block from the mempool and returns it
func (bp *blockProducer) ProduceBlock() (*types.Block, error) {
	bp.currentBlockNum++

	ctx := context.Background()
	// Prepare Polaris for a new block.
	bp.polaris.Prepare(ctx, bp.currentBlockNum)

	// TODO: get from mempool.
	txs := make(types.Transactions, 0)

	// Iterate through all the transactions in the mempool.
	for _, txn := range txs {
		_, err := bp.polaris.ProcessTransaction(context.Background(), txn)
		if err != nil {
			return nil, err
		}
	}

	// Finalize the block.
	block, _, err := bp.polaris.Finalize(ctx)
	if err != nil {
		return nil, err
	}

	return block, nil
}
