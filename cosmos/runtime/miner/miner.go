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
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/miner"

	"pkg.berachain.dev/polaris/beacon/eth"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// emptyHash is a common.Hash initialized to all zeros.
var emptyHash = common.Hash{}

// EnvelopeSerializer is used to convert an envelope into a byte slice that represents
// a cosmos sdk.Tx.
type EnvelopeSerializer interface {
	ToSdkTxBytes(*engine.ExecutionPayloadEnvelope, uint64) ([]byte, error)
}

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	eth.BuilderAPI
	serializer EnvelopeSerializer
}

// New produces a cosmos miner from a geth miner.
func New(gm eth.BuilderAPI) *Miner {
	return &Miner{
		BuilderAPI: gm,
	}
}

// Init sets the transaction serializer.
func (m *Miner) Init(serializer EnvelopeSerializer) {
	m.serializer = serializer
}

// PrepareProposal implements baseapp.PrepareProposal.
func (m *Miner) PrepareProposal(
	ctx sdk.Context, _ *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	var payloadEnvelopeBz []byte
	var err error
	if payloadEnvelopeBz, err = m.buildBlock(ctx); err != nil {
		return nil, err
	}
	return &abci.ResponsePrepareProposal{Txs: [][]byte{payloadEnvelopeBz}}, err
}

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underying worker.
func (m *Miner) buildBlock(ctx sdk.Context) ([]byte, error) {
	var (
		err      error
		envelope *engine.ExecutionPayloadEnvelope
		sCtx     = sdk.UnwrapSDKContext(ctx)
	)
	// Build Payload
	parent := m.CurrentBlock(ctx)
	if envelope, err = m.BuildBlock(ctx, m.constructPayloadArgs(sCtx, parent)); err != nil {
		sCtx.Logger().Error("failed to build payload", "err", err)
		return nil, err
	}

	bz, err := m.serializer.ToSdkTxBytes(envelope, envelope.ExecutionPayload.GasLimit)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// constructPayloadArgs builds a payload to submit to the miner.
func (m *Miner) constructPayloadArgs(
	ctx sdk.Context, parent *types.Block) *miner.BuildPayloadArgs {
	etherbase, err := m.Etherbase(ctx)
	if err != nil {
		ctx.Logger().Error("failed to get etherbase", "err", err)
		return nil
	}

	return &miner.BuildPayloadArgs{
		Timestamp:    parent.Header().Time + 2, //nolint:gomnd // todo fix this arbitrary number.
		FeeRecipient: etherbase,
		Random:       common.Hash{}, /* todo: generated random */
		Withdrawals:  make(types.Withdrawals, 0),
		BeaconRoot:   &emptyHash,
		Parent:       parent.Hash(),
	}
}
