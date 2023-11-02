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

	"github.com/cosmos/gogoproto/proto"

	storetypes "cosmossdk.io/store/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/miner"

	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/eth"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// emptyHash is a common.Hash initialized to all zeros.
var emptyHash = common.Hash{}

// EnvelopeSerializer is used to convert an envelope into a byte slice that represents
// a cosmos sdk.Tx.
type EnvelopeSerializer interface {
	ToSdkTxBytes(*engine.ExecutionPayloadEnvelope, uint64) ([]byte, error)
}

type App interface {
	BeginBlocker(sdk.Context) (sdk.BeginBlock, error)
	PreBlocker(sdk.Context, *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error)
	baseapp.ProposalTxVerifier
}

// EVMKeeper is an interface that defines the methods needed for the EVM setup.
type EVMKeeper interface {
	// Setup initializes the EVM keeper.
	Setup(evmkeeper.Blockchain) error
	SetLatestQueryContext(context.Context) error
}

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	eth.Miner
	app            App
	keeper         EVMKeeper
	valTxSelector  baseapp.TxSelector
	serializer     EnvelopeSerializer
	allowedValMsgs map[string]sdk.Msg
	currentPayload *miner.Payload
}

// New produces a cosmos miner from a geth miner.
func New(gm eth.Miner, app App, keeper EVMKeeper, allowedValMsgs map[string]sdk.Msg) *Miner {
	return &Miner{
		Miner:          gm,
		keeper:         keeper,
		app:            app,
		allowedValMsgs: allowedValMsgs,
		valTxSelector:  baseapp.NewDefaultTxSelector(),
	}
}

// Init sets the transaction serializer.
func (m *Miner) Init(serializer EnvelopeSerializer) {
	m.serializer = serializer
}

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

	// We have to run the PreBlocker && BeginBlocker to get the chain into the state
	// it'll be in when the EVM transaction actually runs.
	if _, err = m.app.PreBlocker(ctx, nil); err != nil {
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

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underying worker.
func (m *Miner) buildBlock(ctx sdk.Context) ([]byte, uint64, error) {
	defer m.clearPayload()
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		return nil, 0, err
	}
	env, gasUsed := m.resolveEnvelope()
	return env, gasUsed, nil
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
		FeeRecipient: m.Etherbase(),
		Random:       common.Hash{}, /* todo: generated random */
		Withdrawals:  make(types.Withdrawals, 0),
		BeaconRoot:   &emptyHash,
	}
}

// resolveEnvelope resolves the payload.
func (m *Miner) resolveEnvelope() ([]byte, uint64) {
	if m.currentPayload == nil {
		return nil, 0
	}
	envelope := m.currentPayload.ResolveFull()
	bz, err := m.serializer.ToSdkTxBytes(envelope, envelope.ExecutionPayload.GasLimit)
	if err != nil {
		panic(err)
	}
	return bz, envelope.ExecutionPayload.GasUsed
}

// clearPayload clears the payload.
func (m *Miner) clearPayload() {
	m.currentPayload = nil
}

// processValidatorMsgs processes the validator messages.
func (m *Miner) processValidatorMsgs(
	ctx sdk.Context, maxTxBytes int64, ethGasUsed uint64, txs [][]byte,
) ([][]byte, error) {
	var maxBlockGas uint64
	if b := ctx.ConsensusParams().Block; b != nil {
		maxBlockGas = uint64(b.MaxGas)
	}

	blockGasRemaining := maxBlockGas - ethGasUsed

	for _, txBz := range txs {
		tx, err := m.app.TxDecode(txBz)
		if err != nil {
			return nil, err
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
