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
	"sync"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// txChanSize is the size of channel listening to NewTxsEvent. The number is referenced from the
// size of tx pool.
const txChanSize = 4096

// Handler listens for new insertions into the geth txpool and broadcasts them to the CometBFT
// layer for p2p and ABCI.
type Handler struct {
	// Cosmos
	logger     log.Logger
	clientCtx  client.Context
	serializer TxSerializer

	// Ethereum
	txPool *txpool.TxPool
	txsCh  chan core.NewTxsEvent
	txsSub event.Subscription
	wg     sync.WaitGroup
}

func NewHandler(logger log.Logger, clientCtx client.Context, txPool *txpool.TxPool) *Handler {
	return &Handler{
		logger:     logger.With("module", "miner-handler"),
		clientCtx:  clientCtx.WithBroadcastMode(flags.BroadcastSync),
		serializer: NewTxSerializer(clientCtx),
		txPool:     txPool,
	}
}

// Start starts the Handler.
// TODO: when is this called?
func (h *Handler) Start() {
	h.wg.Add(1)
	h.txsCh = make(chan core.NewTxsEvent, txChanSize)
	h.txsSub = h.txPool.SubscribeNewTxsEvent(h.txsCh)
	h.logger.Info("handler started")
	go h.txBroadcastLoop() // start broadcast handlers
}

// Stop stops the Handler.
// TODO: when is this called?
func (h *Handler) Stop() {
	// Triggers txBroadcastLoop to quit.
	h.txsSub.Unsubscribe()

	// Leave the channels.
	close(h.txsCh)
	h.wg.Wait()

	h.logger.Info("handler stopped")
}

// txBroadcastLoop announces new transactions to connected peers.
func (h *Handler) txBroadcastLoop() {
	defer h.wg.Done()
	for {
		select {
		case event := <-h.txsCh:
			h.logger.Debug("received new transactions", "numTxs", len(event.Txs))
			h.BroadcastTransactions(event.Txs)
		case <-h.txsSub.Err():
			h.logger.Error("tx subscription error", "err", h.txsSub.Err())
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
			h.logger.Error("failed to serialize transaction", "err", err)
			continue
		}

		// Send the transaction to the CometBFT mempool, which will gossip it to peers via
		// CometBFT's p2p layer.
		rsp, err := h.clientCtx.BroadcastTx(txBytes)

		// If we see an ABCI response error.
		if rsp != nil && rsp.Code != 0 {
			h.logger.Error("failed to broadcast transaction", "rsp", rsp, "err", err)
			continue
		}

		// If we see any other type of error.
		if err != nil {
			h.logger.Error("error on transactions broadcast", "err", err)
			continue
		}
	}
}
