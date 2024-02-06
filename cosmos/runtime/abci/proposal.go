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
	"time"

	"cosmossdk.io/log"

	"github.com/berachain/polaris/cosmos/runtime/chain"
	"github.com/berachain/polaris/cosmos/runtime/miner"

	cometabci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProposalProvider is a struct that provides the abci functions required
// for validators to propose blocks and validators/full nodes to process
// said proposals.
type ProposalProvider struct {
	logger            log.Logger
	preBlocker        sdk.PreBlocker
	beginBlocker      sdk.BeginBlocker
	wrappedMiner      *miner.Miner
	wrappedBlockchain *chain.WrappedBlockchain

	// TODO: refactor validator commands out of the wbc and miner.
	// valCmdProcessor   *ValidatorCommands
	// *eth.ExecutionLayer
}

// NewProposalProvider creates a new ProposalProvider instance.
// It takes a miner.Miner and a chain.WrappedBlockchain as
// arguments and returns a pointer to the initialized ProposalProvider.
func NewProposalProvider(
	preBlocker sdk.PreBlocker, beginBlocker sdk.BeginBlocker,
	wrappedMiner *miner.Miner, wrappedBlockchain *chain.WrappedBlockchain,
	logger log.Logger,
) *ProposalProvider {
	return &ProposalProvider{
		preBlocker:        preBlocker,
		beginBlocker:      beginBlocker,
		wrappedMiner:      wrappedMiner,
		wrappedBlockchain: wrappedBlockchain,
		logger:            logger,
	}
}

// PrepareProposal is responsible for preparing a proposal for the next block.
// It takes a context and a RequestPrepareProposal, simulates finalizing the block,
// and if successful, delegates the proposal preparation to the wrapped miner.
// It returns a ResponsePrepareProposal and an error if any occurs during the process.
func (pp *ProposalProvider) PrepareProposal(
	ctx sdk.Context, req *cometabci.RequestPrepareProposal,
) (*cometabci.ResponsePrepareProposal, error) {
	var (
		start  = time.Now()
		height = ctx.BlockHeight()
	)

	pp.logger.Info(
		"entering prepare proposal",
		"timestamp", start, "height", height)
	defer func() {
		pp.logger.Info(
			"exiting prepare proposal",
			"timestamp", time.Now(),
			"duration", time.Since(start),
			"height", height)
	}()

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
	var (
		start  = time.Now()
		height = ctx.BlockHeight()
	)

	pp.logger.Info(
		"entering process proposal",
		"timestamp", start, "height", height)
	defer func() {
		pp.logger.Info(
			"exiting process proposal",
			"timestamp", time.Now(),
			"duration", time.Since(start),
			"height", height)
	}()

	if err := pp.simulateFinalizeBlock(ctx, req); err != nil {
		return nil, err
	}

	// We set this preblocked, beginblocked, and processed context to the state plugin factory for
	// queries on the node.
	spf := pp.wrappedBlockchain.StatePluginFactory()

	// Technically a race condition here, between here and emitting the chain head
	// event but it is so small and the network latency will most definitely overshadow.
	defer spf.SetLatestQueryContext(ctx)

	// Set the insert chain context for processing the block. NOTE: We insert to the chain but do
	// NOT set the chain head using this context.
	spf.SetInsertChainContext(ctx)
	pp.wrappedBlockchain.PrimePlugins(ctx)

	return pp.wrappedBlockchain.ProcessProposal(ctx, req)
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

	// First check for an abort signal after beginBlock, as it's the first place
	// we spend any significant amount of time.
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// continue
	}

	return nil
}
