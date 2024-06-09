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

package chain

import (
	"fmt"

	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (wbc *WrappedBlockchain) ProcessProposal(
	ctx sdk.Context, req *abci.RequestProcessProposal,
) (*abci.ResponseProcessProposal, error) {
	var (
		err error
	)

	// Pull an execution payload out of the proposal.
	var envelope *engine.ExecutionPayloadEnvelope
	for _, tx := range req.Txs {
		var sdkTx sdk.Tx
		sdkTx, err = wbc.app.TxDecode(tx)
		if err != nil {
			// should have been verified in prepare proposal, we
			// ignore it for now (i.e VE extensions will fail decode).
			continue
		}

		if len(sdkTx.GetMsgs()) == 1 {
			protoEnvelope := sdkTx.GetMsgs()[0]
			if env, ok := protoEnvelope.(*evmtypes.WrappedPayloadEnvelope); ok {
				envelope = env.UnwrapPayload()
				break
			}
		}
	}

	// If the proposal doesn't contain an ethereum envelope, we should reject it.
	if envelope == nil {
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, fmt.Errorf("failed to find envelope in proposal")
	}

	// Convert it to a block.
	var block *ethtypes.Block
	if block, err = engine.ExecutableDataToBlock(*envelope.ExecutionPayload, nil, nil); err != nil {
		ctx.Logger().Error("failed to build evm block", "err", err)
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	}

	// Insert the block into the chain.
	if err = wbc.InsertBlock(block); err != nil {
		ctx.Logger().Error("failed to insert block", "err", err)
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	}

	return &abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_ACCEPT,
	}, nil
}
