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

package txpool

import (
	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// txChanSize is the size of channel listening to NewTxsEvent. The number is referenced from the
// size of tx pool.
const txChanSize = 4096

// TxSubProvider.
type TxSubProvider interface {
	SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription
}

// TxSerializer provides an interface to Serialize Geth Transactions to Bytes (via sdk.Tx).
type TxSerializer interface {
	SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error)
}

// Broadcaster provides an interface to broadcast TxBytes to the comet p2p layer.
type Broadcaster interface {
	BroadcastTxSync(txBytes []byte) (res *sdk.TxResponse, err error)
}

// Subscription represents a subscription to the txpool.
type Subscription interface {
	event.Subscription
}

// handler listens for new insertions into the geth txpool and broadcasts them to the CometBFT
// layer for p2p and ABCI.
type handler struct {
	// Cosmos
	logger     log.Logger
	clientCtx  Broadcaster
	serializer TxSerializer

	// Ethereum
	txPool  TxSubProvider
	txsCh   chan core.NewTxsEvent
	stopCh  chan struct{}
	txsSub  Subscription
	running bool
}

// newHandler creates a new handler and starts the broadcast loop.
func newHandler(
	clientCtx Broadcaster, txPool TxSubProvider, serializer TxSerializer, logger log.Logger,
) *handler {
	txsCh := make(chan core.NewTxsEvent, txChanSize)
	h := &handler{
		logger:     logger,
		clientCtx:  clientCtx,
		serializer: serializer,
		txPool:     txPool,
		txsCh:      txsCh,
		stopCh:     make(chan struct{}),
	}
	return h
}

// Start starts the handler.
func (h *handler) Start() {
	go h.start()
}

// start handles the subscription to the txpool and broadcasts transactions.
func (h *handler) start() {
	// Connect to the subscription.
	h.txsSub = h.txPool.SubscribeNewTxsEvent(h.txsCh)
	h.running = true

	// Handle events.
	var err error
	for {
		select {
		case event := <-h.txsCh:
			h.broadcastTransactions(event.Txs)
		case err = <-h.txsSub.Err():
			h.stopCh <- struct{}{}
		case <-h.stopCh:
			h.stop(err)
		}
	}
}

// Running returns true if the handler is running.
func (h *handler) Running() bool {
	return h.running
}

// Stop stops the handler.
func (h *handler) Stop() {
	h.stopCh <- struct{}{}
}

// stop stops the handler.
func (h *handler) stop(err error) {
	if err != nil {
		h.logger.Error("tx subscription error", "err", err)
	}

	// Triggers txBroadcastLoop to quit.
	h.txsSub.Unsubscribe()
	h.running = false

	// Leave the channels.
	close(h.txsCh)
}

// broadcastTransactions will propagate a batch of transactions to the CometBFT mempool.
func (h *handler) broadcastTransactions(txs coretypes.Transactions) {
	h.logger.Debug("broadcasting transactions", "num_txs", len(txs))
	for _, signedEthTx := range txs {
		// Serialize the transaction to Bytes
		txBytes, err := h.serializer.SerializeToBytes(signedEthTx)
		if err != nil {
			h.logger.Error("failed to serialize transaction", "err", err)
			continue
		}

		// Send the transaction to the CometBFT mempool, which will gossip it to peers via
		// CometBFT's p2p layer.
		rsp, err := h.clientCtx.BroadcastTxSync(txBytes)

		// If we see an ABCI response error.
		if rsp != nil && rsp.Code != 0 {
			h.logger.Error("failed to broadcast transaction", "rsp", rsp, "err", err)
		} else if err != nil {
			h.logger.Error("error on transactions broadcast", "err", err)
		}
	}
}
