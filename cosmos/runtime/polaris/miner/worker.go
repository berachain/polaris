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
	"context"
	"errors"

	"cosmossdk.io/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/types/mempool"
)

const maxTxBytes = 1000000 // 1mb

type worker struct {
	logger     log.Logger
	mempool    mempool.Mempool
	txVerifier baseapp.ProposalTxVerifier

	prepCh     chan *abci.RequestPrepareProposal
	procCh     chan *abci.RequestProcessProposal
	prepRespCh chan *ProposedBlock
	procRespCh chan *ProcessedBlock
	stopCh     chan struct{}

	// snapshotMu       sync.RWMutex
	// snapshotBlock *types.Block
	// snapshotReceipts types.Receipts
}

// newWorker returns a new worker.
func newWorker(
	logger log.Logger, mp mempool.Mempool,
	txVerifier baseapp.ProposalTxVerifier,
	proposalHandler PolarisProposalHandler,
) *worker {
	return &worker{
		logger:     logger.With("module", "polaris-miner"),
		mempool:    mp,
		txVerifier: txVerifier,
		prepCh:     proposalHandler.prepCh,
		procCh:     proposalHandler.procCh,
		prepRespCh: proposalHandler.prepRespCh,
		procRespCh: proposalHandler.procRespCh,
		stopCh:     make(chan struct{}),
	}
}

// Start starts the miner.
func (m *worker) start() {
	go m.loop()
}

// Stop stops the miner.
func (m *worker) stop() {
	close(m.prepCh)
	close(m.procCh)
	close(m.prepRespCh)
	close(m.procRespCh)

	m.stopCh <- struct{}{}
}

// loop is the main loop of the miner.
func (m *worker) loop() {
	for {
		select {
		case req := <-m.prepCh:
			// TODO should we forward the context from the baseapp?
			m.prepRespCh <- m.buildBlock(context.Background(), req.Txs)
		case req := <-m.procCh:
			m.procRespCh <- m.processBlock(context.Background(), req.Txs)
		case <-m.stopCh:
			return
		}
	}
}

// BuildBlock builds a block using the provided mempool and txs.s.
func (m *worker) buildBlock(ctx context.Context, txs [][]byte) *ProposedBlock {
	_, isNoOp := m.mempool.(mempool.NoOpMempool)
	if m.mempool == nil || isNoOp {
		panic("mempool must be set")
	}
	var (
		selectedTxs  [][]byte
		totalTxBytes int64
	)

	iterator := m.mempool.Select(ctx, txs)

	for iterator != nil {
		memTx := iterator.Tx()

		// NOTE: Since transaction verification was already executed in CheckTx,
		// which calls mempool.Insert, in theory everything in the pool should be
		// valid. But some mempool implementations may insert invalid txs, so we
		// check again.
		bz, err := m.txVerifier.PrepareProposalVerifyTx(memTx)
		if err != nil {
			err = m.mempool.Remove(memTx)
			if err != nil && !errors.Is(err, mempool.ErrTxNotFound) {
				panic(err)
			}
		} else {
			// TODO track gas consumption fuck byte size.
			txSize := int64(len(bz))
			if totalTxBytes += txSize; totalTxBytes <= maxTxBytes {
				selectedTxs = append(selectedTxs, bz)
			} else {
				// more transactions.
				break
			}
		}

		iterator = iterator.Next()
	}

	m.logger.Info("⛏️ mining block", "num_txs", len(selectedTxs), "total_tx_bytes", totalTxBytes)
	return &ProposedBlock{&abci.ResponsePrepareProposal{Txs: selectedTxs}, nil}
}

// ProcessBlock processes a block using the provided mempool and txs.
func (m *worker) processBlock(_ context.Context, txs [][]byte) *ProcessedBlock {
	m.logger.Info("🤨 processing block", "num_txs", len(txs))
	for _, txBytes := range txs {
		_, err := m.txVerifier.ProcessProposalVerifyTx(txBytes)
		if err != nil {
			return &ProcessedBlock{
				&abci.ResponseProcessProposal{
					Status: abci.ResponseProcessProposal_REJECT,
				}, nil}
		}
	}
	return &ProcessedBlock{
		&abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT},
		nil,
	}
}
