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

import (
	"sync"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	"pkg.berachain.dev/polaris/eth/core/types"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// txChanSize is the size of channel listening to NewTxsEvent. The number is referenced from the
// size of tx pool.
const txChanSize = 4096

// Handler listens for new insertions into the geth txpool and broadcasts them to the CometBFT
// layer for p2p and ABCI.
type Handler struct {
	// Cosmos
	serializer TxSerializer
	clientCtx  client.Context

	// Ethereum
	txPool *txpool.TxPool
	txsCh  chan core.NewTxsEvent
	txsSub event.Subscription
	wg     sync.WaitGroup
}

func NewHandler(txPool *txpool.TxPool, s TxSerializer) *Handler {
	return &Handler{
		txPool:     txPool,
		serializer: s,
	}
}

// Start starts the Handler.
// TODO: when is this called?
func (h *Handler) Start() {
	h.wg.Add(1)
	h.txsCh = make(chan core.NewTxsEvent, txChanSize)
	h.txsSub = h.txPool.SubscribeNewTxsEvent(h.txsCh)
	go h.txBroadcastLoop() // start broadcast handlers
}

// Stop stops the Handler.
// TODO: when is this called?
func (h *Handler) Stop() {
	h.txsSub.Unsubscribe() // quits txBroadcastLoop

	// Quit new txs channel
	close(h.txsCh)
	h.wg.Wait()

	// log.Info("Ethereum protocol stopped")
}

// txBroadcastLoop announces new transactions to connected peers.
func (h *Handler) txBroadcastLoop() {
	defer h.wg.Done()
	for {
		select {
		case event := <-h.txsCh:
			h.BroadcastTransactions(event.Txs)
		case <-h.txsSub.Err():
			return
		}
	}
}

// BroadcastTransactions will propagate a batch of transactions to the CometBFT mempool.
func (h *Handler) BroadcastTransactions(txs types.Transactions) {
	for _, signedEthTx := range txs {
		// Serialize the transaction to Bytes
		txBytes, err := h.serializer.SerializeToBytes(signedEthTx)
		if err != nil {
			// TODO: log error?
			panic(errorslib.Wrap(err, "failed to serialize transaction"))
		}

		// Send the transaction to the CometBFT mempool, which will gossip it to peers via
		// CometBFT's p2p layer.
		syncCtx := h.clientCtx.WithBroadcastMode(flags.BroadcastAsync)
		rsp, err := syncCtx.BroadcastTx(txBytes)
		if rsp != nil && rsp.Code != 0 {
			// TODO: log error?
			panic(errorsmod.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog))
		}
		if err != nil {
			// TODO: log error?
			panic(err)
		}
	}
}
