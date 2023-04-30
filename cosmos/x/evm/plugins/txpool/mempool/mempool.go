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

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// EthTxPool is a mempool for Ethereum transactions. It wraps a
// PriorityNonceMempool and caches transactions that are added to the mempool by
// ethereum transaction hash.
type EthTxPool struct {
	sdkmempool.Mempool

	// ethTxCache caches transactions that are added to the mempool
	// so that they can be retrieved later
	ethTxCache map[common.Hash]*coretypes.Transaction

	// // nonceCache caches the pending nonce by txhash
	// nonceCache map[common.Address]*coretypes.Transaction

	// // minedBlockCache caches the mined transaction by block hash
	// minedBlockCache map[common.Hash][]*coretypes.Transaction

	// blockNumberCache

}

// New is called when the mempool is created.
func NewEthTxPoolFrom(m sdkmempool.Mempool) *EthTxPool {
	return &EthTxPool{
		Mempool:    m,
		ethTxCache: make(map[common.Hash]*coretypes.Transaction),
	}
}

// Insert is called when a transaction is added to the mempool.
func (etp *EthTxPool) Insert(ctx context.Context, tx sdk.Tx) error {
	// Call the base mempool's Insert method
	if err := etp.Mempool.Insert(ctx, tx); err != nil {
		return err
	}

	// We want to cache
	etr, ok := utils.GetAs[*types.EthTransactionRequest](tx.GetMsgs()[0])
	if !ok {
		return nil
	}

	t := etr.AsTransaction()
	etp.ethTxCache[t.Hash()] = t
	return nil
}

// GetTx is called when a transaction is retrieved from the mempool.
func (etp *EthTxPool) GetTransaction(hash common.Hash) *coretypes.Transaction {
	return etp.ethTxCache[hash]
}

// GetPoolTransactions is called when the mempool is retrieved.
func (etp *EthTxPool) GetPoolTransactions() coretypes.Transactions {
	txs := make(coretypes.Transactions, 0, len(etp.ethTxCache))
	for _, tx := range etp.ethTxCache {
		txs = append(txs, tx)
	}
	return txs
}

// Remove is called when a transaction is removed from the mempool.
func (etp *EthTxPool) Remove(tx sdk.Tx) error {
	// Call the base mempool's Remove method
	if err := etp.Mempool.Remove(tx); err != nil {
		return err
	}

	// We want to cache this tx.
	etr, ok := utils.GetAs[*types.EthTransactionRequest](tx)
	if !ok {
		return nil
	}

	delete(etp.ethTxCache, etr.AsTransaction().Hash())
	return nil
}
