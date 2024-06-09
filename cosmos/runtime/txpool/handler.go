// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package txpool

import (
	"errors"
	"sync/atomic"
	"time"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// txChanSize is the size of channel listening to NewTxsEvent. The number is referenced from the
// size of tx pool.
const (
	txChanSize = 4096
	maxRetries = 5
	retryDelay = 50 * time.Millisecond
	statPeriod = 60 * time.Second
)

// SdkTx is used to generate mocks.
type SdkTx interface {
	sdk.Tx
}

// TxSubProvider.
type TxSubProvider interface {
	SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription
	Stats() (int, int)
}

// TxSerializer provides an interface to Serialize Geth Transactions to Bytes (via sdk.Tx).
type TxSerializer interface {
	ToSdkTx(signedTx *ethtypes.Transaction, gasLimit uint64) (sdk.Tx, error)
	ToSdkTxBytes(signedTx *ethtypes.Transaction, gasLimit uint64) ([]byte, error)
}

// TxBroadcaster provides an interface to broadcast TxBytes to the comet p2p layer.
type TxBroadcaster interface {
	BroadcastTxSync(txBytes []byte) (res *sdk.TxResponse, err error)
}

// Subscription represents a subscription to the txpool.
type Subscription interface {
	event.Subscription
}

// failedTx represents a transaction that failed to broadcast.
type failedTx struct {
	tx      *ethtypes.Transaction
	retries int
}

// handler listens for new insertions into the geth txpool and broadcasts them to the CometBFT
// layer for p2p and ABCI.
type handler struct {
	// Cosmos
	logger     log.Logger
	clientCtx  TxBroadcaster
	serializer TxSerializer
	crc        CometRemoteCache

	// Ethereum
	txPool  TxSubProvider
	txsCh   chan core.NewTxsEvent
	stopCh  chan struct{}
	txsSub  Subscription
	running atomic.Bool

	// Queue for failed transactions
	failedTxs chan *failedTx
}

// newHandler creates a new handler.
func newHandler(
	clientCtx TxBroadcaster, txPool TxSubProvider, serializer TxSerializer,
	crc CometRemoteCache, logger log.Logger,
) *handler {
	h := &handler{
		logger:     logger,
		clientCtx:  clientCtx,
		serializer: serializer,
		crc:        crc,
		txPool:     txPool,
		txsCh:      make(chan core.NewTxsEvent, txChanSize),
		stopCh:     make(chan struct{}),
		failedTxs:  make(chan *failedTx, txChanSize),
	}
	return h
}

// Start starts the handler.
func (h *handler) Start() error {
	if h.running.Load() {
		return errors.New("handler already started")
	}
	go h.mainLoop()
	go h.failedLoop() // Start the retry policy
	go h.statLoop()
	return nil
}

// Stop stops the handler.
func (h *handler) Stop() error {
	if !h.Running() {
		return errors.New("handler already stopped")
	}

	// Push two stop signals to the stop channel to ensure that both loops stop.
	h.stopCh <- struct{}{}
	h.stopCh <- struct{}{}
	return nil
}

// mainLoop start handles the subscription to the txpool and broadcasts transactions.
func (h *handler) mainLoop() {
	// Connect to the subscription.
	h.txsSub = h.txPool.SubscribeTransactions(h.txsCh, true)
	h.logger.With("module", "txpool-handler").Info("starting txpool handler")
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
			telemetry.IncrCounter(float32(len(event.Txs)), MetricKeyCometLocalTxs)
			h.broadcastTransactions(event.Txs)
		}
	}
}

// failedLoop will periodically attempt to re-broadcast failed transactions.
func (h *handler) failedLoop() {
	for {
		select {
		case <-h.stopCh:
			return
		case failed := <-h.failedTxs:
			if failed.retries == 0 {
				h.logger.Error("failed to broadcast transaction after max retries", "tx", maxRetries)
				continue
			}
			telemetry.IncrCounter(float32(1), MetricKeyBroadcastRetry)
			h.broadcastTransaction(failed.tx, failed.retries-1)
		}

		// We slightly space out the retries in order to prioritize new transactions.
		time.Sleep(retryDelay)
	}
}

func (h *handler) statLoop() {
	ticker := time.NewTicker(statPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-h.stopCh:
			return
		case <-ticker.C:
			pending, queue := h.txPool.Stats()
			telemetry.SetGauge(float32(pending), MetricKeyTxPoolPending)
			telemetry.SetGauge(float32(queue), MetricKeyTxPoolQueue)
		}
	}
}

// Running returns true if the handler is running.
func (h *handler) Running() bool {
	return h.running.Load()
}

// stop stops the handler.
func (h *handler) stop(err error) {
	// Mark as not running to prevent further events.
	h.running.Store(false)

	// If we are stopping because of an error, log it.
	if err != nil {
		h.logger.Error("txpool handler", "error", err)
	}

	// Triggers txBroadcastLoop to quit.
	h.txsSub.Unsubscribe()

	// Close channels.
	close(h.txsCh)
	close(h.stopCh)
	close(h.failedTxs)
}

// broadcastTransactions will propagate a batch of transactions to the CometBFT mempool.
func (h *handler) broadcastTransactions(txs ethtypes.Transactions) {
	numBroadcasted := 0
	for _, signedEthTx := range txs {
		if !h.crc.IsRemoteTx(signedEthTx.Hash()) {
			h.broadcastTransaction(signedEthTx, maxRetries)
			numBroadcasted++
		}
	}
	h.logger.Debug(
		"broadcasting transactions", "num_received", len(txs), "num_broadcasted", numBroadcasted,
	)
}

// broadcastTransaction will propagate a transaction to the CometBFT mempool.
func (h *handler) broadcastTransaction(tx *ethtypes.Transaction, retries int) {
	txBytes, err := h.serializer.ToSdkTxBytes(tx, tx.Gas())
	if err != nil {
		h.logger.Error("failed to serialize transaction", "err", err)
		return
	}

	// Send the transaction to the CometBFT mempool, which will gossip it to peers via
	// CometBFT's p2p layer.
	rsp, err := h.clientCtx.BroadcastTxSync(txBytes)
	if err != nil {
		h.logger.Error("error on transactions broadcast", "err", err)
		h.failedTxs <- &failedTx{tx: tx, retries: retries}
		return
	}

	// If rsp == 1, likely the txn is already in a block, and thus the broadcast failing is actually
	// the desired behaviour.
	if rsp == nil || rsp.Code == 0 || rsp.Code == 1 {
		return
	}

	switch rsp.Code {
	case sdkerrors.ErrMempoolIsFull.ABCICode():
		h.logger.Error("failed to broadcast: comet-bft mempool is full", "tx_hash", tx.Hash())
		telemetry.IncrCounter(float32(1), MetricKeyMempoolFull)
	case
		sdkerrors.ErrTxInMempoolCache.ABCICode():
		return
	default:
		h.logger.Error("failed to broadcast transaction",
			"codespace", rsp.Codespace, "code", rsp.Code, "info", rsp.Info, "tx_hash", tx.Hash())
		telemetry.IncrCounter(float32(1), MetricKeyBroadcastFailure)
	}

	h.failedTxs <- &failedTx{tx: tx, retries: retries}
}
