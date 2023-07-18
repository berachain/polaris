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
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/core/txpool"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Compile-time interface assertion.
var _ sdkmempool.Mempool = (*WrappedGethTxPool)(nil)

// WrappedGethTxPool is a mempool for Ethereum transactions. It wraps a Geth TxPool.
// NOTE: currently does not support adding `sdk.Tx`s that do NOT have a `WrappedEthereumTransaction`
// as the tx Msg.
type WrappedGethTxPool struct {
	// The underlying Geth mempool implementation.
	*txpool.TxPool

	// serializer converts eth txs to sdk txs when being iterated over.
	serializer SdkTxSerializer

	// pendingBaseFee is set by the miner, as is the signer.
	pendingBaseFee *big.Int
	signer         coretypes.Signer

	// iterator is used to iterate over the txpool.
	iterator *iterator
}

// NewWrappedGethTxPool creates a new Ethereum transaction pool.
func NewWrappedGethTxPool() *WrappedGethTxPool {
	return &WrappedGethTxPool{}
}

// Setup sets the chain config and sdk tx serializer on the wrapped Geth TxPool.
func (gtp *WrappedGethTxPool) Setup(txPool *txpool.TxPool, serializer SdkTxSerializer) {
	gtp.TxPool = txPool
	gtp.serializer = serializer
}

// Prepare prepares the txpool for the next pending block.
func (gtp *WrappedGethTxPool) Prepare(pendingBaseFee *big.Int, signer coretypes.Signer) {
	gtp.pendingBaseFee = pendingBaseFee
	gtp.signer = signer
}

// Insert is called when a transaction is added to the mempool.
func (gtp *WrappedGethTxPool) Insert(_ context.Context, tx sdk.Tx) error {
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		err := gtp.AddRemotes(coretypes.Transactions{ethTx})[0]
		// If we see ErrAlreadyKnown, we can ignore it, since this is likely from the ABCI broadcast.
		// TODO: we should do a check here to make sure that the ErrAlreadyKnown is happening because of
		// the fact that InsertLocal was called. If this is a genuine p2p broadcast of a tx, we may want to
		// actually handle the error if already known, in the case where two indepdent peers are sending us the
		// same transaction. TODO verify this.
		if errors.Is(err, txpool.ErrAlreadyKnown) {
			return nil
		}
		return err
	}
	return nil
}

// InsertSync is called when a transaction is added to the mempool (for testing purposes).
func (gtp *WrappedGethTxPool) InsertSync(_ context.Context, tx sdk.Tx) error {
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		return gtp.AddRemotesSync(coretypes.Transactions{ethTx})[0]
	}
	return nil
}

// Remove is called when a transaction is removed from the mempool.
func (gtp *WrappedGethTxPool) Remove(tx sdk.Tx) error {
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		if gtp.iterator != nil {
			gtp.iterator.txs.Pop()
		} else if gtp.RemoveTx(ethTx.Hash(), true) < 1 {
			// remove from the pending queue of txs in the geth mempool.
			// Note: RemoveTx will return 0 if the tx was removed from future queue. Generally, any
			// tx in the future queue will not be removed because only the pending txs get
			// selected by prepare proposal.
			return sdkmempool.ErrTxNotFound
		}
	}
	return nil
}

// Select returns an Iterator over the app-side mempool. If txs are specified, then they shall be
// incorporated into the Iterator. The Iterator must closed by the caller.
func (gtp *WrappedGethTxPool) Select(context.Context, [][]byte) sdkmempool.Iterator {
	// return nil if there are no pending txs
	pendingTxs := gtp.Pending(true)
	if len(pendingTxs) == 0 {
		return nil
	}

	// return an iterator over the pending txs, sorted by price and nonce
	gtp.iterator = &iterator{
		txs: coretypes.NewTransactionsByPriceAndNonce(
			gtp.signer,
			pendingTxs,
			gtp.pendingBaseFee,
		),
		serializer: gtp.serializer,
	}
	return gtp.iterator
}

// CountTx returns the number of transactions currently in the mempool.
func (gtp *WrappedGethTxPool) CountTx() int {
	pending, queued := gtp.Stats()
	return pending + queued
}
