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
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// EthTxPool is a mempool for Ethereum transactions. It wraps a PriorityNonceMempool and caches
// transactions that are added to the mempool by ethereum transaction hash.
type EthTxPool struct {
	// The underlying mempool implementation.
	*PriorityNonceMempool[int64]

	// ethTxCache caches transactions that are added to the mempool so that they can be retrieved
	// later
	ethTxCache map[common.Hash]*coretypes.Transaction

	// NonceRetriever is used to retrieve the nonce for a given address (this is typically a
	// reference to the StateDB).
	nr NonceRetriever

	// We have a mutex to protect the ethTxCache and nonces maps since they are accessed
	// concurrently by multiple goroutines.
	mu sync.RWMutex
}

// New is called when the mempool is created.
func NewEthTxPoolFrom(mp *PriorityNonceMempool[int64]) *EthTxPool {
	return &EthTxPool{
		PriorityNonceMempool: mp,
		ethTxCache:           make(map[common.Hash]*coretypes.Transaction),
	}
}

// SetNonceRetriever sets the nonce retriever db for the mempool.
func (etp *EthTxPool) SetNonceRetriever(nr NonceRetriever) {
	etp.nr = nr
}

// Insert is called when a transaction is added to the mempool.
func (etp *EthTxPool) Insert(ctx context.Context, tx sdk.Tx) error {
	etp.mu.Lock()
	defer etp.mu.Unlock()

	// Call the base mempool's Insert method
	if err := etp.PriorityNonceMempool.Insert(ctx, tx); err != nil {
		return err
	}

	// We want to cache
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		etp.ethTxCache[ethTx.Hash()] = ethTx
	}

	return nil
}

// Remove is called when a transaction is removed from the mempool.
func (etp *EthTxPool) Remove(tx sdk.Tx) error {
	etp.mu.Lock()
	defer etp.mu.Unlock()

	// Call the base mempool's Remove method
	if err := etp.PriorityNonceMempool.Remove(tx); err != nil {
		return err
	}

	// We want to remove the caches of this tx.
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		delete(etp.ethTxCache, ethTx.Hash())
	}

	return nil
}

// Get is called when a transaction is retrieved from the mempool.
func (etp *EthTxPool) Get(hash common.Hash) *coretypes.Transaction {
	return etp.ethTxCache[hash]
}

// Pending is called when txs in the mempool are retrieved.
func (etp *EthTxPool) Pending(bool) map[common.Address]coretypes.Transactions {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	allNonces := etp.senderIndices
	pending := make(map[common.Address]coretypes.Transactions)
	for sender, list := range allNonces {
		// get Eth Address of sender
		addrBech32, _ := sdk.AccAddressFromBech32(sender)
		addr := cosmlib.AccAddressToEthAddress(addrBech32)

		// add the first eth tx in the list, if it exists
		var ethTx *coretypes.Transaction
		for elem := list.Front(); elem != nil; elem = elem.Next() {
			if ethTx = evmtypes.GetAsEthTx(utils.MustGetAs[sdk.Tx](elem.Value)); ethTx != nil {
				pending[addr] = coretypes.Transactions{ethTx}
				break
			}
		}
	}

	return pending
}

// queued is called when content of the mempool is retrieved.
func (etp *EthTxPool) queued() map[common.Address]coretypes.Transactions {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	allNonces := etp.senderIndices
	queued := make(map[common.Address]coretypes.Transactions)
	for sender, list := range allNonces {
		// get Eth Address of sender
		addrBech32, _ := sdk.AccAddressFromBech32(sender)
		addr := cosmlib.AccAddressToEthAddress(addrBech32)

		// skip the first ethTx seen, add the rest to the queued list
		var (
			ethTxs  coretypes.Transactions
			seenOne bool
		)
		for elem := list.Front(); elem != nil; elem = elem.Next() {
			if ethTx := evmtypes.GetAsEthTx(utils.MustGetAs[sdk.Tx](elem.Value)); ethTx != nil {
				if seenOne {
					ethTxs = append(ethTxs, ethTx)
				}
				seenOne = true
			}
		}
		queued[addr] = ethTxs
	}

	return queued
}

// GetNonce returns the nonce for the given address from the mempool if the address has sent a tx
// in the mempool.
func (etp *EthTxPool) Nonce(addr common.Address) uint64 {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	// search the addr's txs for the first eth tx nonce (first pending nonce)
	if txs := etp.senderIndices[cosmlib.AddressToAccAddress(addr).String()]; txs != nil {
		for elem := txs.Front(); elem != nil; elem = elem.Next() {
			if ethTx := evmtypes.GetAsEthTx(utils.MustGetAs[sdk.Tx](elem.Value)); ethTx != nil {
				// pending nonce is the account's incremented nonce
				return ethTx.Nonce() + 1
			}
		}
	}

	// if the addr has no eth txs, fallback to the nonce retriever db
	return etp.nr.GetNonce(addr)
}

// Stats returns the number of currently pending (locally created) transactions.
func (etp *EthTxPool) Stats() (int, int) {
	pending, queued := etp.Content()
	return len(pending), len(queued)
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
