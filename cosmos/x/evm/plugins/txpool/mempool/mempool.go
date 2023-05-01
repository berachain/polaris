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

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
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

	// nonces is a cache of the pending nonces by sender address
	nonces map[common.Address]uint64

	// NonceRetriever is used to retrieve the nonce for a given address.
	// (this is typically a reference to the StateDB)
	nr NonceRetriever

	// We have a mutex to protect the ethTxCache and nonces maps since they are accessed
	// concurrently by multiple goroutines.
	mu sync.RWMutex

	// txFeed and scope is used to send new batch transactions to new txs subscribers when the
	// batch is added to the mempool.
	txFeed event.Feed
	scope  event.SubscriptionScope
}

// New is called when the mempool is created.
func NewEthTxPoolFrom(mp *PriorityNonceMempool[int64]) *EthTxPool {
	return &EthTxPool{
		PriorityNonceMempool: mp,
		ethTxCache:           make(map[common.Hash]*coretypes.Transaction),
		nonces:               make(map[common.Address]uint64),
	}
}

// SetNonceRetriever sets the nonce retriever db for the mempool.
func (etp *EthTxPool) SetNonceRetriever(nr NonceRetriever) {
	etp.nr = nr
}

// GetNewTxsEventSubscription returns a new event subscription for the new txs feed.
func (etp *EthTxPool) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	// Currently sending an individual new txs event for every new tx added to the mempool.
	// TODO: support sending batch new txs events when adding queued txs to the pending txs.
	return etp.scope.Track(etp.txFeed.Subscribe(ch))
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
	etr, ok := utils.GetAs[*types.EthTransactionRequest](tx.GetMsgs()[0])
	if !ok {
		return nil
	}

	ethTx := etr.AsTransaction()
	etp.ethTxCache[ethTx.Hash()] = ethTx
	sender, _ := coretypes.Sender(coretypes.LatestSignerForChainID(ethTx.ChainId()), ethTx)
	etp.nonces[sender] = ethTx.Nonce() + 1

	etp.txFeed.Send(core.NewTxsEvent{Txs: coretypes.Transactions{ethTx}})

	return nil
}

// GetTx is called when a transaction is retrieved from the mempool.
func (etp *EthTxPool) Get(hash common.Hash) *coretypes.Transaction {
	return etp.ethTxCache[hash]
}

// GetTransactions is called when the mempool is retrieved.
func (etp *EthTxPool) Pending(bool) map[common.Address]coretypes.Transactions {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	// TODO: implement me.
	return nil
}

// GetNonce returns the nonce for the given address from the mempool if the address has sent a tx
// in the mempool.
func (etp *EthTxPool) Nonce(addr common.Address) uint64 {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	nonce, found := etp.nonces[addr]
	if !found {
		// fallback to nonce retrieval from db
		if nonce = etp.nr.GetNonce(addr); nonce > 0 {
			etp.nonces[addr] = nonce
		}
	}
	return nonce
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
	etr, ok := utils.GetAs[*types.EthTransactionRequest](tx)
	if !ok {
		return nil
	}

	ethTx := etr.AsTransaction()
	sender, _ := coretypes.Sender(coretypes.LatestSignerForChainID(ethTx.ChainId()), ethTx)
	delete(etp.ethTxCache, ethTx.Hash())
	delete(etp.nonces, sender)
	return nil
}

// Stats returns the number of currently pending (locally created) transactions.
func (etp *EthTxPool) Stats() (int, int) {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	pending, queued := etp.Content()
	return len(pending), len(queued)
}

// ContentFrom retrieves the data content of the transaction pool, returning the
// pending as well as queued transactions of this address, grouped by nonce.
func (etp *EthTxPool) ContentFrom(addr common.Address) (coretypes.Transactions, coretypes.Transactions) {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	pending, queued := etp.Content()
	return pending[addr], queued[addr]
}

// Content retrieves the data content of the transaction pool, returning all the
// pending as well as queued transactions, grouped by account and nonce.
func (etp *EthTxPool) Content() (map[common.Address]coretypes.Transactions, map[common.Address]coretypes.Transactions) {
	etp.mu.RLock()
	defer etp.mu.RUnlock()

	// TODO: implement me.
	return nil, nil
}
