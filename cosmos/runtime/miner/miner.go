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
	"errors"
	"sync"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/berachain/polaris/eth"
	"github.com/berachain/polaris/eth/core"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/miner"
)

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	miner eth.Miner
	app   TxDecoder
	bc    core.Blockchain

	valTxSelector  baseapp.TxSelector
	serializer     EnvelopeSerializer
	allowedValMsgs map[string]sdk.Msg
	currentPayload *miner.Payload

	blockBuilderMu *sync.RWMutex
}

// New produces a cosmos miner from a geth miner.
func New(
	miner eth.Miner, app TxDecoder, allowedValMsgs map[string]sdk.Msg,
	bc core.Blockchain, blockBuilderMu *sync.RWMutex,
) *Miner {
	return &Miner{
		miner:          miner,
		app:            app,
		bc:             bc,
		allowedValMsgs: allowedValMsgs,
		valTxSelector:  baseapp.NewDefaultTxSelector(),
		blockBuilderMu: blockBuilderMu,
	}
}

// Init sets the transaction serializer.
func (m *Miner) Init(serializer EnvelopeSerializer) {
	m.serializer = serializer
}

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underlying worker.
func (m *Miner) buildBlock(ctx sdk.Context) ([]byte, uint64, error) {
	defer m.clearPayload()

	// Record the time it takes to build a payload.
	defer telemetry.MeasureSince(time.Now(), MetricKeyBuildBlock)

	// Miner locks for block building to occupy the txpool.
	m.blockBuilderMu.Lock()
	defer m.blockBuilderMu.Unlock()

	// Submit payload for building with the given context.
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		return nil, 0, err
	}
	env, gasUsed := m.resolveEnvelope()

	return env, gasUsed, nil
}

// submitPayloadForBuilding submits a payload for building.
func (m *Miner) submitPayloadForBuilding(ctx context.Context) error {
	var (
		err         error
		payload     *miner.Payload
		sCtx        = sdk.UnwrapSDKContext(ctx)
		payloadArgs = m.constructPayloadArgs(uint64(sCtx.BlockTime().Unix()))
	)

	// Set the mining context for geth to build the payload with.
	m.bc.StatePluginFactory().SetLatestMiningContext(ctx)
	m.bc.PrimePlugins(ctx)

	// Build Payload.
	if payload, err = m.miner.BuildPayload(payloadArgs); err != nil {
		sCtx.Logger().Error("failed to build payload", "err", err)
		return err
	}
	m.currentPayload = payload
	sCtx.Logger().Info("submitted payload for building")
	return nil
}

// constructPayloadArgs builds a payload to submit to the miner.
func (m *Miner) constructPayloadArgs(blockTime uint64) *miner.BuildPayloadArgs {
	return &miner.BuildPayloadArgs{
		Timestamp:    blockTime,
		FeeRecipient: m.miner.Etherbase(),
		Random:       common.Hash{}, /* todo: generated random */
		Withdrawals:  make(ethtypes.Withdrawals, 0),
		BeaconRoot:   nil, // Add this when implementing Cancun.
	}
}

// resolveEnvelope resolves the payload.
func (m *Miner) resolveEnvelope() ([]byte, uint64) {
	if m.currentPayload == nil {
		return nil, 0
	}
	envelope := m.currentPayload.ResolveFull()
	payload := envelope.ExecutionPayload

	// Record metadata about the payload
	defer telemetry.SetGauge(float32(payload.GasUsed), MetricKeyBlockGasUsed)
	defer telemetry.SetGauge(float32(len(payload.Transactions)), MetricKeyTransactions)

	bz, err := m.serializer.ToSdkTxBytes(envelope, payload.GasLimit)
	if err != nil {
		panic(err)
	}

	return bz, payload.GasUsed
}

// clearPayload clears the payload.
func (m *Miner) clearPayload() {
	m.currentPayload = nil
}

// processValidatorMsgs processes the validator messages.
func (m *Miner) processValidatorMsgs(
	ctx sdk.Context, maxTxBytes int64, ethGasUsed uint64, txs [][]byte,
) ([][]byte, error) {
	b := ctx.ConsensusParams().Block
	if b == nil {
		return nil, errors.New("consensus params block is nil")
	}
	if uint64(b.MaxGas) < ethGasUsed {
		return nil, errors.New("eth gas used exceeds comet block max gas")
	}
	blockGasRemaining := uint64(b.MaxGas) - ethGasUsed

	for _, txBz := range txs {
		tx, err := m.app.TxDecode(txBz)
		if err != nil {
			continue
		}

		includeTx := true
		for _, msg := range tx.GetMsgs() {
			if _, ok := m.allowedValMsgs[proto.MessageName(msg)]; !ok {
				includeTx = false
				break
			}
		}

		if includeTx {
			stop := m.valTxSelector.SelectTxForProposal(
				ctx, uint64(maxTxBytes), blockGasRemaining, tx, txBz,
			)
			if stop {
				break
			}
		}
	}
	return m.valTxSelector.SelectedTxs(ctx), nil
}
