package miner

// func (w *miner) fillTransactions(interrupt *atomic.Int32, env *environment) error {
// 	pending := w.txPool.Pending(true)

// 	// Split the pending transactions into locals and remotes.
// 	localTxs, remoteTxs := make(map[common.Address][]*txpool.LazyTransaction), pending
// 	for _, account := range w.txPool.Locals() {
// 		if txs := remoteTxs[account]; len(txs) > 0 {
// 			delete(remoteTxs, account)
// 			localTxs[account] = txs
// 		}
// 	}

// 	// Fill the block with all available pending transactions.
// 	if len(localTxs) > 0 {
// 		txs := miner.NewTransactionsByPriceAndNonce(env.signer, localTxs, env.header.BaseFee)
// 		if err := w.commitTransactions(env, txs, interrupt); err != nil {
// 			return err
// 		}
// 	}
// 	if len(remoteTxs) > 0 {
// 		txs := miner.NewTransactionsByPriceAndNonce(env.signer, remoteTxs, env.header.BaseFee)
// 		if err := w.commitTransactions(env, txs, interrupt); err != nil {
// 			return err
// 		}
// 	}
// }
