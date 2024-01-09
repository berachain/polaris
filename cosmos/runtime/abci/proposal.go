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
	"github.com/berachain/polaris/cosmos/runtime/chain"
	"github.com/berachain/polaris/cosmos/runtime/miner"
	"github.com/berachain/polaris/eth"

	cometabci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProposalProvider is a struct that provides the abci functions required
// for validators to propose blocks and validators/full nodes to process
// said proposals.
type ProposalProvider struct {
	*eth.ExecutionLayer
	preBlocker   sdk.PreBlocker
	beginBlocker sdk.BeginBlocker
	// valCmdProcessor   *ValidatorCommands
	wrappedMiner      *miner.Miner
	wrappedBlockchain *chain.WrappedBlockchain
}

// NewProposalProvider creates a new ProposalProvider instance.
// It takes a miner.Miner and a chain.WrappedBlockchain as
// arguments and returns a pointer to the initialized ProposalProvider.
func NewProposalProvider(
	wrappedMiner *miner.Miner, wrappedBlockchain *chain.WrappedBlockchain,
) *ProposalProvider {
	return &ProposalProvider{
		wrappedMiner:      wrappedMiner,
		wrappedBlockchain: wrappedBlockchain,
	}
}

// PrepareProposal is responsible for preparing a proposal for the next block.
// It takes a context and a RequestPrepareProposal, simulates finalizing the block,
// and if successful, delegates the proposal preparation to the wrapped miner.
// It returns a ResponsePrepareProposal and an error if any occurs during the process.
func (pp *ProposalProvider) PrepareProposal(
	ctx sdk.Context, req *cometabci.RequestPrepareProposal,
) (*cometabci.ResponsePrepareProposal, error) {
	if err := pp.simulateFinalizeBlock(ctx, req); err != nil {
		return nil, err
	}

	return pp.wrappedMiner.PrepareProposal(ctx, req)
}

// ProcessProposal processes the incoming proposal.
// It takes a context and a RequestProcessProposal, simulates finalizing the block,
// and if successful, delegates the proposal processing to the wrapped blockchain.
// It returns a ResponseProcessProposal and an error if any occurs during the process.
func (pp *ProposalProvider) ProcessProposal(
	ctx sdk.Context, req *cometabci.RequestProcessProposal,
) (*cometabci.ResponseProcessProposal, error) {
	if err := pp.simulateFinalizeBlock(ctx, req); err != nil {
		return nil, err
	}
	resp, err := pp.wrappedBlockchain.ProcessProposal(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// simulateFinalizeBlock simulates the execution of a block.
// We have to run the PreBlocker && BeginBlocker to get the chain into the state
// it'll be in when the EVM transaction actually runs.
func (pp *ProposalProvider) simulateFinalizeBlock(ctx sdk.Context, req abciRequest) error {
	if _, err := pp.preBlocker(ctx, &cometabci.RequestFinalizeBlock{
		Txs:                req.GetTxs(),
		Time:               req.GetTime(),
		Misbehavior:        req.GetMisbehavior(),
		Height:             req.GetHeight(),
		NextValidatorsHash: req.GetNextValidatorsHash(),
		ProposerAddress:    req.GetProposerAddress(),
	}); err != nil {
		return err
	}

	if _, err := pp.beginBlocker(ctx); err != nil {
		return err
	}

	return nil
}
