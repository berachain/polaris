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
