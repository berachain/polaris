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
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/core/types"
)

// Mempool defines a mempool interface that can be used to query transactions.
type Miner struct {
	baseTxVerifier baseapp.ProposalTxVerifier
	mempool        Mempool
	logger         log.Logger
	pendingBlock   *types.Block
}

// NewMiner returns a new instance of a miner.
func NewMiner(
	mempool Mempool,
	txVerifier baseapp.ProposalTxVerifier,
	logger log.Logger,
) *Miner {
	return &Miner{
		baseTxVerifier: txVerifier,
		mempool:        mempool,
		logger:         logger,
	}
}

func (miner *Miner) Logger() log.Logger {
	return miner.logger.With("module", "miner")
}

// PrepareProposalVerifyTx performs transaction verification when a proposer is
// creating a block proposal during PrepareProposal. Any state committed to the
// PrepareProposal state internally will be discarded. <nil, err> will be
// returned if the transaction cannot be encoded. <bz, nil> will be returned if
// the transaction is valid, otherwise <bz, err> will be returned.
func (miner *Miner) PrepareProposalVerifyTx(tx sdk.Tx) ([]byte, error) {
	return miner.baseTxVerifier.PrepareProposalVerifyTx(tx)
}

// ProcessProposalVerifyTx performs transaction verification when receiving a
// block proposal during ProcessProposal. Any state committed to the
// ProcessProposal state internally will be discarded. <nil, err> will be
// returned if the transaction cannot be decoded. <Tx, nil> will be returned if
// the transaction is valid, otherwise <Tx, err> will be returned.
func (miner *Miner) ProcessProposalVerifyTx(txBz []byte) (sdk.Tx, error) {
	// TODO: figure out if this is even possible to update.
	return miner.baseTxVerifier.ProcessProposalVerifyTx(txBz)
}
