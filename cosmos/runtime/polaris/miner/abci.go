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

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Payload[T any] struct {
	ctx     context.Context
	content *T
	err     error
}

type (
	// PrepareProposal ABCI.
	ProposalRequest = Payload[abci.RequestPrepareProposal]
	ProposedBlock   = Payload[abci.ResponsePrepareProposal]

	// ProcessorProposal ABCI.
	ProcessRequest = Payload[abci.RequestProcessProposal]
	ProcessedBlock = Payload[abci.ResponseProcessProposal]
)

// PolarisProposalHandler defines the default ABCI PrepareProposal and
// ProcessProposal handlers.
type PolarisProposalHandler struct {
	prepCh     chan *ProposalRequest
	procCh     chan *ProcessRequest
	prepRespCh chan *ProposedBlock
	procRespCh chan *ProcessedBlock
}

// NewPolarisProposalHandler returns a new default.
func NewPolarisProposalHandler() PolarisProposalHandler {
	return PolarisProposalHandler{
		prepCh:     make(chan *ProposalRequest),
		procCh:     make(chan *ProcessRequest),
		prepRespCh: make(chan *ProposedBlock),
		procRespCh: make(chan *ProcessedBlock),
	}
}

// PrepareProposalHandler returns the default implementation for processing an
// ABCI proposal. The application's mempool is enumerated and all valid
// transactions are added to the proposal. Transactions are valid if they:.
func (h PolarisProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		// Fire off a request to build a block to propose.
		h.prepCh <- &ProposalRequest{ctx: ctx, content: req}

		// Wait for the response.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case block := <-h.prepRespCh:
			return block.content, block.err
		}
	}
}

// PrepareProposalHandler returns the default implementation for processing an
// ABCI proposal. The application's mempool is enumerated and all valid
// transactions are added to the proposal. Transactions are valid if they:.
func (h PolarisProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		// Fire off the request to process the proposal.
		h.procCh <- &ProcessRequest{ctx: ctx, content: req}

		// Wait for the response.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case block := <-h.procRespCh:
			return block.content, block.err
		}
	}
}
