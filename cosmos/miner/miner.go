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
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/miner"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// emptyHash is a common.Hash initialized to all zeros.
var emptyHash = common.Hash{}

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	*miner.Miner
	serializer     evmtypes.TxSerializer
	currentPayload *miner.Payload
}

// New produces a cosmos miner from a geth miner.
func New(gm *miner.Miner) *Miner {
	return &Miner{
		Miner: gm,
	}
}

// Init sets the transaction serializer.
func (m *Miner) Init(serializer evmtypes.TxSerializer) {
	m.serializer = serializer
}

// PrepareProposal implements baseapp.PrepareProposal.
func (m *Miner) PrepareProposal(
	ctx sdk.Context, _ *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	var txs [][]byte
	var err error
	if txs, err = m.buildBlock(ctx); err != nil {
		return nil, err
	}
	return &abci.ResponsePrepareProposal{Txs: txs}, err
}

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underying worker.
func (m *Miner) buildBlock(ctx sdk.Context) ([][]byte, error) {
	defer m.clearPayload()
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		return nil, err
	}
	return m.resolveTxs(), nil
}

// submitPayloadForBuilding submits a payload for building.
func (m *Miner) submitPayloadForBuilding(ctx context.Context) error {
	var (
		err     error
		payload *miner.Payload
		sCtx    = sdk.UnwrapSDKContext(ctx)
	)

	// Build Payload
	if payload, err = m.BuildPayload(m.constructPayloadArgs(sCtx)); err != nil {
		sCtx.Logger().Error("failed to build payload", "err", err)
		return err
	}
	m.currentPayload = payload
	sCtx.Logger().Info("submitted payload for building")
	return nil
}

// constructPayloadArgs builds a payload to submit to the miner.
func (m *Miner) constructPayloadArgs(ctx sdk.Context) *miner.BuildPayloadArgs {
	return &miner.BuildPayloadArgs{
		Timestamp:    uint64(ctx.BlockTime().Unix()),
		FeeRecipient: common.Address{}, /* todo: set etherbase */
		Random:       common.Hash{},    /* todo: generated random */
		Withdrawals:  make(types.Withdrawals, 0),
		BeaconRoot:   &emptyHash,
	}
}

// resolveTxs resolves the transactions from the payload.
func (m *Miner) resolveTxs() [][]byte {
	if m.currentPayload == nil {
		return nil
	}
	envelope := m.currentPayload.ResolveFull()
	ethTxBzs := envelope.ExecutionPayload.Transactions
	txs := make([][]byte, len(envelope.ExecutionPayload.Transactions))

	// encode to sdk.txs and then
	for i, ethTxBz := range ethTxBzs {
		var tx types.Transaction
		if err := tx.UnmarshalBinary(ethTxBz); err != nil {
			return nil
		}
		bz, err := m.serializer.SerializeToBytes(&tx)
		if err != nil {
			panic(err)
		}
		txs[i] = bz
	}
	return txs
}

// clearPayload clears the payload.
func (m *Miner) clearPayload() {
	m.currentPayload = nil
}
