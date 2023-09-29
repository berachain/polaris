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

package abci

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	prepare "pkg.berachain.dev/polaris/cosmos/abci/prepare"
	process "pkg.berachain.dev/polaris/cosmos/abci/process"
	"pkg.berachain.dev/polaris/eth/polar"
)

type (
	// ProposalTxVerifier defines the interface that is implemented by BaseApp,
	// that any custom ABCI PrepareProposal and ProcessProposal handler can use
	// to verify a transaction.
	ProposalTxVerifier interface {
		PrepareProposalVerifyTx(tx sdk.Tx) ([]byte, error)
		ProcessProposalVerifyTx(txBz []byte) (sdk.Tx, error)
	}

	// DefaultProposalHandler defines the default ABCI PrepareProposal and
	// ProcessProposal handlers.
	DefaultProposalHandler struct {
		proposer  prepare.Handler
		processor process.Handler
	}
)

func NewDefaultProposalHandler(polaris *polar.Polaris, txVerifier ProposalTxVerifier) *DefaultProposalHandler {
	return &DefaultProposalHandler{
		proposer:  prepare.NewHandler(polaris, txVerifier),
		processor: process.NewHandler(txVerifier),
	}
}

func (h *DefaultProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return h.proposer.PrepareProposal
}

func (h *DefaultProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return h.processor.ProcessProposal
}
