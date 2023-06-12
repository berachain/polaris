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

const maxTxBytes = 1000000000

type Worker struct {
	logger     log.Logger
	mempool    mempool.Mempool
	txVerifier baseapp.ProposalTxVerifier
	prepChan   chan *abci.RequestPrepareProposal
	procChan   chan *abci.RequestProcessProposal
	prepResp   chan *ProposedBlock
	procResp   chan *ProcessedBlock
	stop       chan struct{}
}

// NewWorker returns a new miner.
func NewWorker(
	logger log.Logger, mp mempool.Mempool, txVerifier baseapp.ProposalTxVerifier, proposalHandler PolarisProposalHandler,
) *Worker {
	return &Worker{
		logger:     logger.With("module", "polaris-miner"),
		mempool:    mp,
		txVerifier: txVerifier,
		prepChan:   proposalHandler.prepChan,
		procChan:   proposalHandler.procChan,
		prepResp:   proposalHandler.prepResp,
		procResp:   proposalHandler.procResp,
		stop:       make(chan struct{}),
	}
}

// Start starts the miner.
func (m *Worker) Start() {
	go m.loop()
}

// Stop stops the miner.
func (m *Worker) Stop() {
	close(m.prepChan)
	close(m.procChan)
	close(m.prepResp)
	close(m.procResp)

	m.stop <- struct{}{}
}

// loop is the main loop of the miner.
func (m *Worker) loop() {
	for {
		select {
		case req := <-m.prepChan:
			m.prepResp <- m.BuildBlock(context.Background(), req.Txs)
		case req := <-m.procChan:
			m.procResp <- m.ProcessBlock(context.Background(), req.Txs)
		case <-m.stop:
			return
		}
	}
}

// BuildBlock builds a block using the provided mempool and txs.s.
func (m *Worker) BuildBlock(ctx context.Context, txs [][]byte) *ProposedBlock {
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
				// We've reached capacity per req.MaxTxBytes so we cannot select any
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
func (m *Worker) ProcessBlock(_ context.Context, txs [][]byte) *ProcessedBlock {
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
