package miner

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

import (
	"errors"
	"math/big"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/eth/core/types"
)

type Mempool = sdkmempool.Mempool

func (miner *Miner) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		var (
			selectedTxs  [][]byte
			totalTxBytes int64
		)

		// TODO: this is just an example.
		miner.pendingBlock = types.NewBlock(
			&types.Header{
				Number:     big.NewInt(0),
				Difficulty: big.NewInt(0),
				GasLimit:   0,
				GasUsed:    0,
				Time:       0,
				Extra:      []byte{},
			},
			nil,
			nil,
			nil,
			trie.NewStackTrie(nil),
		)

		iterator := miner.mempool.Select(ctx, req.Txs)

		for iterator != nil {
			memTx := iterator.Tx()

			// NOTE: Since transaction verification was already executed in CheckTx,
			// which calls mempool.Insert, in theory everything in the pool should be
			// valid. But some mempool implementations may insert invalid txs, so we
			// check again.
			bz, err := miner.baseTxVerifier.PrepareProposalVerifyTx(memTx)
			if err != nil {
				err = miner.mempool.Remove(memTx)
				if err != nil && !errors.Is(err, sdkmempool.ErrTxNotFound) {
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

		return &abci.ResponsePrepareProposal{Txs: selectedTxs}, nil
	}
}

func (miner *Miner) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		for _, txBytes := range req.Txs {
			_, err := miner.baseTxVerifier.ProcessProposalVerifyTx(txBytes)
			if err != nil {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, err
			}
		}

		miner.Logger().Info("⛏️ block was mined", "number", ctx.BlockHeight())
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}
