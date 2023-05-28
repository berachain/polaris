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

package mempool

import (
	"context"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Get is called when a transaction is retrieved from the mempool.
func (etp *EthTxPool) Get(hash common.Hash) *coretypes.Transaction {
	return etp.ethTxCache[hash]
}

// Pending is called when txs in the mempool are retrieved.
func (etp *EthTxPool) Pending(bool) map[common.Address]coretypes.Transactions {
	pendingNonces := make(map[common.Address]uint64)
	pending := make(map[common.Address]coretypes.Transactions)

	etp.mu.Lock()
	defer etp.mu.Unlock()

	for iter := etp.PriorityNonceMempool.Select(context.Background(), nil); iter != nil; iter = iter.Next() {
		tx := iter.Tx()
		if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
			addr := coretypes.GetSender(ethTx)
			pendingNonce := pendingNonces[addr]
			switch {
			case pendingNonce == 0:
				// If its the first tx, set the pending nonce to the nonce of the tx.
				txNonce := ethTx.Nonce()
				// If on the first lookup the nonce delta is more than 0, then there is a gap
				// and thus no pending transactions, but there are queued transactions. We
				// continue.
				if sdbNonce := etp.nr.GetNonce(addr); txNonce-sdbNonce >= 1 {
					continue
				}
				// this is a pending tx, add it to the pending map.
				pendingNonces[addr] = txNonce
				pending[addr] = append(pending[addr], ethTx)
			case ethTx.Nonce() == pendingNonce+1:
				// If its not the first tx, but the nonce is the same as the pending nonce, add
				// it to the list.
				pending[addr] = append(pending[addr], ethTx)
				pendingNonces[addr] = pendingNonce + 1
			default:
				// If we see an out of order nonce, we break since the rest should be "queued".
				break
			}
		}
	}

	return pending
}

// queued retrieves the content of the mempool.
//

func (etp *EthTxPool) queued() map[common.Address]coretypes.Transactions {
	pendingNonces := make(map[common.Address]uint64)
	queued := make(map[common.Address]coretypes.Transactions)

	etp.mu.Lock()
	defer etp.mu.Unlock()

	// After the lock is released we can iterate over the mempool.
	for iter := etp.PriorityNonceMempool.Select(context.Background(), nil); iter != nil; iter = iter.Next() {
		if ethTx := evmtypes.GetAsEthTx(iter.Tx()); ethTx != nil {
			addr := coretypes.GetSender(ethTx)
			pendingNonce, seenTransaction := pendingNonces[addr]
			switch {
			case !seenTransaction:
				// When we see a transaction, mark it as the pending nonce.
				pendingNonce = ethTx.Nonce()
				// If on the first lookup the nonce delta is more than 0, then there is a gap
				// and thus no pending transactions, but there are queued transactions.
				if pendingNonce-etp.nr.GetNonce(addr) >= 1 {
					queued[addr] = append(queued[addr], ethTx)
				} else {
					// this is a pending tx, add it to the pending map.
					pendingNonces[addr] = pendingNonce
				}
			case ethTx.Nonce() == pendingNonces[addr]+1:
				// If we are still contiguous and the nonce is the same as the pending nonce,
				// increment the pending nonce.
				pendingNonce++
				pendingNonces[addr] = pendingNonce
			default:
				// If we are still contiguous and the nonce is greater than the pending nonce, we are no longer contiguous.
				// Add to the queued list.
				queued[addr] = append(queued[addr], ethTx)
				// All other transactions in the skip list should be queued.
			}
		}
	}

	return queued
}

// Nonce returns the nonce for the given address from the mempool if the address has sent a tx
// in the mempool.
func (etp *EthTxPool) Nonce(addr common.Address) uint64 {
	pendingNonces := make(map[common.Address]uint64)
	etp.mu.Lock()

	// search for the first pending ethTx
	for iter := etp.PriorityNonceMempool.Select(context.Background(), nil); iter != nil; iter = iter.Next() {
		if ethTx := evmtypes.GetAsEthTx(iter.Tx()); ethTx != nil {
			txAddr := coretypes.GetSender(ethTx)
			if addr != txAddr {
				continue
			}
			_, ok := pendingNonces[addr]
			txNonce := ethTx.Nonce()
			switch {
			case !ok:
				// If on the first lookup the nonce delta is more than 0, then there is a gap
				// and thus no pending transactions, but there are queued transactions.
				if sdbNonce := etp.nr.GetNonce(addr); txNonce-sdbNonce >= 1 {
					return sdbNonce
				}
				// this is a pending tx, add it to the pending map.
				pendingNonces[addr] = txNonce
			case txNonce == pendingNonces[addr]+1:
				// If we are still contiguous and the nonce is the same as the pending nonce,
				// increment the pending nonce.
				pendingNonces[addr]++
			case txNonce > pendingNonces[addr]+1:
				// As soon as we see a non contiguous nonce we break.
				break
			}
		}
	}

	// We move this here instead of defer as a slight optimization.
	etp.mu.Unlock()

	// if the addr has no eth txs, fallback to the nonce retriever db
	if _, ok := pendingNonces[addr]; !ok {
		return etp.nr.GetNonce(addr)
	}

	// pending nonce is 1 more than the current nonce
	return pendingNonces[addr] + 1
}

// Stats returns the number of currently pending and queued (locally created) transactions.
func (etp *EthTxPool) Stats() (int, int) {
	var pendingTxsLen, queuedTxsLen int
	pending, queued := etp.Content()

	etp.mu.RLock()
	defer etp.mu.RUnlock()

	for _, txs := range pending {
		pendingTxsLen += len(txs)
	}
	for _, txs := range queued {
		queuedTxsLen += len(txs)
	}
	return pendingTxsLen, queuedTxsLen
}

// ContentFrom retrieves the data content of the transaction pool, returning the pending as well as
// queued transactions of this address, grouped by nonce.
func (etp *EthTxPool) ContentFrom(addr common.Address) (coretypes.Transactions, coretypes.Transactions) {
	pending, queued := etp.Content()
	return pending[addr], queued[addr]
}

// Content retrieves the data content of the transaction pool, returning all the pending as well as
// queued transactions, grouped by account and nonce.
func (etp *EthTxPool) Content() (
	map[common.Address]coretypes.Transactions, map[common.Address]coretypes.Transactions,
) {
	return etp.Pending(false), etp.queued()
}
