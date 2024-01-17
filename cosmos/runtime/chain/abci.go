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
