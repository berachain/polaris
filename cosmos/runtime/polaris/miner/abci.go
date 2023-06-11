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
	"errors"

	"cosmossdk.io/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
)

// PolarisProposalHandler defines the default ABCI PrepareProposal and
// ProcessProposal handlers.
type PolarisProposalHandler struct {
	logger     log.Logger
	mempool    mempool.Mempool
	txVerifier baseapp.ProposalTxVerifier
}

// NewPolarisProposalHandler returns a new default.
func NewPolarisProposalHandler(
	logger log.Logger, mp mempool.Mempool, txVerifier baseapp.ProposalTxVerifier,
) PolarisProposalHandler {
	return PolarisProposalHandler{
		logger:     logger.With("module", "polaris-miner"),
		mempool:    mp,
		txVerifier: txVerifier,
	}
}

// PrepareProposalHandler returns the default implementation for processing an
// ABCI proposal. The application's mempool is enumerated and all valid
// transactions are added to the proposal. Transactions are valid if they:.
func (h PolarisProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		// If the mempool is nil or NoOp we simply return the transactions
		// requested from CometBFT, which, by default, should be in FIFO order.
		_, isNoOp := h.mempool.(mempool.NoOpMempool)
		if h.mempool == nil || isNoOp {
			panic("mempool must be set")
		}
		var (
			selectedTxs  [][]byte
			totalTxBytes int64
		)

		iterator := h.mempool.Select(ctx, req.Txs)

		for iterator != nil {
			memTx := iterator.Tx()

			// NOTE: Since transaction verification was already executed in CheckTx,
			// which calls mempool.Insert, in theory everything in the pool should be
			// valid. But some mempool implementations may insert invalid txs, so we
			// check again.
			bz, err := h.txVerifier.PrepareProposalVerifyTx(memTx)
			if err != nil {
				err = h.mempool.Remove(memTx)
				if err != nil && !errors.Is(err, mempool.ErrTxNotFound) {
					panic(err)
				}
			} else {
				txSize := int64(len(bz))
				if totalTxBytes += txSize; totalTxBytes <= req.MaxTxBytes {
					selectedTxs = append(selectedTxs, bz)
				} else {
					// We've reached capacity per req.MaxTxBytes so we cannot select any
					// more transactions.
					break
				}
			}

			iterator = iterator.Next()
		}

		h.logger.Info("🤡 proposing block", "num_txs", len(selectedTxs), "total_tx_bytes", totalTxBytes)
		return &abci.ResponsePrepareProposal{Txs: selectedTxs}, nil
	}
}

// ProcessProposalHandler returns the default implementation for processing an
// ABCI proposal. Every transaction in the proposal must pass 2 conditions:
//
// 1. The transaction bytes must decode to a valid transaction.
// 2. The transaction must be valid (i.e. pass runTx, AnteHandler only)
//
// If any transaction fails to pass either condition, the proposal is rejected.
// Note that step (2) is identical to the validation step performed in
// DefaultPrepareProposal. It is very important that the same validation logic
// is used in both steps, and applications must ensure that this is the case in
// non-default handlers.
func (h PolarisProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	// If the mempool is nil or NoOp we simply return ACCEPT,
	// because PrepareProposal may have included txs that could fail verification.
	_, isNoOp := h.mempool.(mempool.NoOpMempool)
	if h.mempool == nil || isNoOp {
		panic("mempool must be set")
	}

	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		h.logger.Info("🤨 processing block", "num_txs", len(req.Txs))
		for _, txBytes := range req.Txs {
			_, err := h.txVerifier.ProcessProposalVerifyTx(txBytes)
			if err != nil {
				return &abci.ResponseProcessProposal{ //nolint:nilerr // must return nil.
					Status: abci.ResponseProcessProposal_REJECT,
				}, nil
			}
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}
