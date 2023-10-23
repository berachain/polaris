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
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/miner"

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

// Miner implements the baseapp.TxSelector interface.
type Miner struct {
	eth.Miner
	serializer     EnvelopeSerializer
	currentPayload *miner.Payload
	payloadTimeout time.Duration
}

// New produces a cosmos miner from a geth miner.
func New(gm eth.Miner, payloadTimeout time.Duration) *Miner {
	return &Miner{
		Miner:          gm,
		payloadTimeout: payloadTimeout,
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
	if payloadEnvelopeBz, err = m.buildBlock(ctx); errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &abci.ResponsePrepareProposal{Txs: [][]byte{payloadEnvelopeBz}}, err
}

// buildBlock builds and submits a payload, it also waits for the txs
// to resolve from the underying worker.
func (m *Miner) buildBlock(ctx context.Context) ([]byte, error) {
	defer m.clearPayload()
	if err := m.submitPayloadForBuilding(ctx); err != nil {
		return nil, err
	}
	return m.resolveEnvelope(ctx, m.payloadTimeout)
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
func (m *Miner) resolveEnvelope(ctx context.Context, timeout time.Duration) ([]byte, error) {
	sCtx := sdk.UnwrapSDKContext(ctx).Logger()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go m.resolvePayload(resultChan, errChan)

	select {
	case <-ctx.Done():
		// If we timed out, return an empty payload.
		// TODO: penalize validators for not being able to deliver the payload?
		sCtx.Error("failed to resolve envelope, proposing empty payload", "err", ctx.Err())
		return m.resolveEmptyPayload()
	case result := <-resultChan:
		sdk.UnwrapSDKContext(ctx).Logger().Info("successfully resolved envelope")
		return result, <-errChan
	}
}

// resolvePayload is a helper function to resolve the payload in a separate goroutine.
func (m *Miner) resolvePayload(resultChan chan []byte, errChan chan error) {
	if m.currentPayload == nil {
		resultChan <- nil
		errChan <- nil
		return
	}
	envelope := m.currentPayload.ResolveFull()
	result, err := m.serializer.ToSdkTxBytes(envelope, envelope.ExecutionPayload.GasLimit)
	resultChan <- result
	errChan <- err
}

// resolveEmptyPayload is a helper function to resolve the empty payload in a separate goroutine.
func (m *Miner) resolveEmptyPayload() ([]byte, error) {
	envelope := m.currentPayload.ResolveEmpty()
	return m.serializer.ToSdkTxBytes(envelope, envelope.ExecutionPayload.GasLimit)
}

// clearPayload clears the payload.
func (m *Miner) clearPayload() {
	m.currentPayload = nil
}
