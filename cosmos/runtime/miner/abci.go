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

// Package miner implements the Ethereum miner.
package miner

import (
	storetypes "cosmossdk.io/store/types"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
)

// emptyHash is a common.Hash initialized to all zeros.
var emptyHash = common.Hash{}

// PrepareProposal implements baseapp.PrepareProposal.
func (m *Miner) PrepareProposal(
	ctx sdk.Context, req *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	var (
		payloadEnvelopeBz []byte
		err               error
		valTxs            [][]byte
		ethGasUsed        uint64
	)

	// m.keeper.GetStatePlugin().SetStateOverride(ctx)
	// defer m.keeper.GetStatePlugin().ClearStateOverride()

	// We have to run the PreBlocker && BeginBlocker to get the chain into the state
	// it'll be in when the EVM transaction actually runs.
	if _, err = m.app.PreBlocker(ctx, &abci.RequestFinalizeBlock{
		Txs:                req.Txs,
		Time:               req.Time,
		Misbehavior:        req.Misbehavior,
		Height:             req.Height,
		NextValidatorsHash: req.NextValidatorsHash,
		ProposerAddress:    req.ProposerAddress,
	}); err != nil {
		return nil, err
	} else if _, err = m.app.BeginBlocker(ctx); err != nil {
		return nil, err
	}

	ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "prepare proposal")
	ctx.BlockGasMeter().RefundGas(ctx.BlockGasMeter().GasConsumed(), "prepare proposal")
	ctx = ctx.WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{}).
		WithGasMeter(storetypes.NewInfiniteGasMeter())

	// We have to prime the state plugin.
	if err = m.keeper.SetLatestQueryContext(ctx); err != nil {
		return nil, err
	}

	// Trigger the geth miner to build a block.
	if payloadEnvelopeBz, ethGasUsed, err = m.buildBlock(ctx); err != nil {
		return nil, err
	}

	// Process the validator messages.
	if valTxs, err = m.processValidatorMsgs(ctx, req.MaxTxBytes, ethGasUsed, req.Txs); err != nil {
		return nil, err
	}

	// Combine the payload envelope with the validator transactions.
	allTxs := [][]byte{payloadEnvelopeBz}

	// If there are validator transactions, append them to the allTxs slice.
	if len(valTxs) > 0 {
		allTxs = append(allTxs, valTxs...)
	}

	// Return the payload and validator transactions as a transaction in the proposal.
	return &abci.ResponsePrepareProposal{Txs: allTxs}, err
}
