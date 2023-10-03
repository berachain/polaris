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
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/params"
)

// emptyHash is a common.Hash initialized to all zeros.
var emptyHash = common.Hash{}

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	*miner.Miner
	serializer     evmtypes.TxSerializer
	currentPayload *miner.Payload
}

// NewMiner returns a new instance of the Miner.
func NewMiner(
	eth miner.Backend, config *miner.Config, chainConfig *params.ChainConfig,
	mux *event.TypeMux, engine consensus.Engine, //nolint:staticcheck // its okay for now.
	isLocalBlock func(header *types.Header) bool,
) *Miner {
	return &Miner{
		Miner: miner.New(eth, config, chainConfig, mux, engine, isLocalBlock),
	}
}

// PrepareProposal implements baseapp.PrepareProposal.
func (m *Miner) PrepareProposal(
	ctx sdk.Context, _ *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	// TODO: maybe add some safety checks against `req` args.
	return &abci.ResponsePrepareProposal{Txs: m.buildBlock(ctx)}, nil
}

// SetSerializer sets the transaction serializer.
func (m *Miner) SetSerializer(serializer evmtypes.TxSerializer) {
	m.serializer = serializer
}

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underying worker.
func (m *Miner) buildBlock(ctx sdk.Context) [][]byte {
	defer func() { m.currentPayload = nil }()
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		panic(err)
	}
	return [][]byte{m.resolveEnvelope()}
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

func (m *Miner) resolveEnvelope() []byte {
	if m.currentPayload == nil {
		return nil
	}
	envelope := m.currentPayload.ResolveFull()
	bz, err := m.serializer.PayloadToBytes(envelope)
	if err != nil {
		panic(err)
	}
	return bz
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
