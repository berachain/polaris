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
	defer func() { m.currentPayload = nil }()
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		return nil, err
	}
	return [][]byte{m.resolveEnvelope()}, nil
}

// submitPayloadForBuilding submits a payload for building.
func (m *Miner) submitPayloadForBuilding(ctx context.Context) error {
	sCtx := sdk.UnwrapSDKContext(ctx)
	sCtx.Logger().Info("Submitting payload for building")
	payload, err := m.BuildPayload(&miner.BuildPayloadArgs{
		Parent:    common.Hash{}, // Empty parent is fine, geth miner will handle.
		Timestamp: uint64(sCtx.BlockTime().Unix()),
		// TODO: properly fill in the rest of the payload.
		FeeRecipient: common.Address{},
		Random:       common.Hash{},
		Withdrawals:  make(types.Withdrawals, 0),
		BeaconRoot:   &emptyHash,
	})
	if err != nil {
		sCtx.Logger().Error("Failed to build payload", "err", err)
		return err
	}
	m.currentPayload = payload
	return nil
}

// resolveEnvelope resolves the payload.
func (m *Miner) resolveEnvelope() []byte {
	if m.currentPayload == nil {
		return nil
	}

	bz, err := m.serializer.PayloadToBytes(m.currentPayload.ResolveFull())
	if err != nil {
		panic(err)
	}
	return bz
}
