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
	"errors"
	"sync/atomic"
	"time"

	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	queuelib "pkg.berachain.dev/polaris/lib/queue"
)

// txChanSize is the size of channel listening to NewTxsEvent. The number is referenced from the
// size of tx pool.
const (
	txChanSize        = 4096
	retryDelay        = 50 * time.Millisecond
	emptyQueueBackoff = 250 * time.Millisecond
)

// SdkTx is used to generate mocks.
type SdkTx interface {
	sdk.Tx
}

// TxSubProvider.
type TxSubProvider interface {
	SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription
}

// TxSerializer provides an interface to Serialize Geth Transactions to Bytes (via sdk.Tx).
type TxSerializer interface {
	ToSdkTx(signedTx *coretypes.Transaction, gasLimit uint64) (sdk.Tx, error)
	ToSdkTxBytes(signedTx *coretypes.Transaction, gasLimit uint64) ([]byte, error)
}

// TxBroadcaster provides an interface to broadcast TxBytes to the comet p2p layer.
type TxBroadcaster interface {
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
	clientCtx  TxBroadcaster
	serializer TxSerializer

	// Ethereum
	txPool  TxSubProvider
	txsCh   chan core.NewTxsEvent
	stopCh  chan struct{}
	txsSub  Subscription
	running atomic.Bool // Running method returns true if the handler is running.

	txQueue *queuelib.LockFreeQueue[[]byte]
}

// newHandler function creates a new handler.
func newHandler(
	clientCtx TxBroadcaster, txPool TxSubProvider, serializer TxSerializer, logger log.Logger,
) *handler {
	h := &handler{
		logger:     logger,
		clientCtx:  clientCtx,
		serializer: serializer,
		txPool:     txPool,
		txsCh:      make(chan core.NewTxsEvent, txChanSize),
		stopCh:     make(chan struct{}),
		txQueue:    queuelib.NewLockFreeQueue[[]byte](),
	}
	return h
}

// Running method returns true if the handler is running.
func (h *handler) Running() bool {
	return h.running.Load()
}

// Start method starts the handler.
func (h *handler) Start() error {
	if h.running.Load() {
		return errors.New("handler has already been started")
	}
	go h.queueLoop()
	go h.broadcastLoop() // This starts the retry policy
	return nil
}

// Stop method stops the handler.
func (h *handler) Stop() error {
	if !h.Running() {
		return errors.New("handler has already been stopped")
	}

	// Push two stop signals to the stop channel to ensure that both loops stop.
	h.stopCh <- struct{}{}
	h.stopCh <- struct{}{}
	return nil
}

// queueLoop method handles the subscription to the txpool and broadcasts transactions.
func (h *handler) queueLoop() {
	// Connect to the subscription.
	h.txsSub = h.txPool.SubscribeNewTxsEvent(h.txsCh)
	h.logger.With("module", "txpool-handler").Info("txpool handler is starting")
	h.running.Store(true)

	// Handle events.
	var err error
	for {
		select {
		case <-h.stopCh:
			h.stop(err)
			return
		case err = <-h.txsSub.Err():
			h.stopCh <- struct{}{}
		case event := <-h.txsCh:
			h.queueTransactions(event.Txs)
		}
	}
}

// broadcastLoop function is responsible for continuously broadcasting transactions from the
// queue. It runs in a loop and performs the following steps:
// broadcastLoop method broadcasts transactions from the queue.
func (h *handler) broadcastLoop() {
	for {
		select {
		case <-h.stopCh:
			return
		case <-time.After(emptyQueueBackoff):
			if tx := h.dequeueTxBytes(); tx != nil {
				h.broadcastTx(tx)
			}
		}
	}
}

// broadcastTx method dequeues the next transaction off the queue and attempts to broadcast it.
func (h *handler) broadcastTx(txBytes []byte) {
	if err := h.broadcastTransaction(txBytes); errors.Is(err, sdkerrors.ErrMempoolIsFull) {
		// If the mempool is full, we need to re-enqueue this transaction to be broadcast.
		h.enqueueTxBytes(txBytes)
	} else if errors.Is(err, sdkerrors.ErrTxInMempoolCache) {
		// Do nothing, since this transaction already exists at the CometBFT layer.
		return
	} else if err != nil {
		h.logger.Error("An error occurred during the transaction broadcast", "err", err)
	}
}

// enqueueTxBytes method adds a transaction to the queue and signals that a
// new transaction is ready. It takes a byte slice representing the transaction as an argument.
func (h *handler) enqueueTxBytes(txBytes []byte) {
	h.txQueue.Enqueue(txBytes)
}

// dequeueTxBytes method removes and returns the next transaction from the queue.
// It waits for a signal that a transaction is ready before attempting to dequeue.
func (h *handler) dequeueTxBytes() []byte {
	return h.txQueue.Dequeue()
}

// stop method stops the handler.
func (h *handler) stop(err error) {
	// Mark as not running to prevent further events.
	h.running.Store(false)

	// If we are stopping because of an error, log it.
	if err != nil {
		h.logger.Error("An error occurred in the txpool handler", "error", err)
	}

	// This triggers the txBroadcastLoop to quit.
	h.txsSub.Unsubscribe()

	// Close channels.
	close(h.txsCh)
	close(h.stopCh)
	close(h.stopCh)
}

// queueTransactions method will propagate a batch of transactions to the CometBFT mempool.
func (h *handler) queueTransactions(txs coretypes.Transactions) {
	h.logger.Debug("The transactions are being broadcasted", "num_txs", len(txs))
	for _, signedEthTx := range txs {
		txBytes, err := h.serializer.ToSdkTxBytes(signedEthTx, signedEthTx.Gas())
		if err != nil {
			h.logger.Error("Failed to serialize the transaction", "err", err)
			return
		}
		h.enqueueTxBytes(txBytes)
	}
}

// broadcastTransaction method will propagate a transaction to the CometBFT mempool.
func (h *handler) broadcastTransaction(txBytes []byte) error {
	// Send the transaction to the CometBFT mempool, which will gossip it to peers via
	// CometBFT's p2p layer.
	rsp, err := h.clientCtx.BroadcastTxSync(txBytes)

	if rsp == nil || rsp.Code == 0 {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
