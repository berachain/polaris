// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package miner implements the Ethereum miner.
package miner

import (
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
