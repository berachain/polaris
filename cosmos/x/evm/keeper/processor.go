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

package keeper

import (
	"context"
	"fmt"
	"time"

	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// ProcessPayloadEnvelope uses Geth's beacon engine API to build a block from a execution payload
// request. It is called by Cosmos-SDK during ABCI DeliverTx phase (1 cosmos tx to build the entire
// eth block).
func (k *Keeper) ProcessPayloadEnvelope(
	ctx context.Context, msg *evmtypes.WrappedPayloadEnvelope,
) (*evmtypes.WrappedPayloadEnvelopeResponse, error) {
	var (
		err      error
		block    *ethtypes.Block
		envelope engine.ExecutionPayloadEnvelope
	)
	// TODO: maybe we just consume the block gas limit and call it a day?
	sCtx := sdk.UnwrapSDKContext(ctx)
	gasMeter := sCtx.GasMeter()
	blockGasMeter := sCtx.BlockGasMeter()

	// Reset GasMeter to 0.
	//
	// TODO: we need to remove this next re-genesis.
	defer gasMeter.RefundGas(gasMeter.GasConsumed(), "reset after evm")
	defer blockGasMeter.RefundGas(blockGasMeter.GasConsumed(), "reset after evm")

	if err = envelope.UnmarshalJSON(msg.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload envelope: %w", err)
	}

	if block, err = engine.ExecutableDataToBlock(*envelope.ExecutionPayload, nil, nil); err != nil {
		k.Logger(sCtx).Error("failed to build evm block", "err", err)
		return nil, err
	}

	// Record how long it takes to insert the new block into the chain.
	defer telemetry.ModuleMeasureSince(evmtypes.ModuleName,
		time.Now(), evmtypes.MetricKeyInsertBlock)

	// Set the finalize block context on the state plugin factory. Set the finalize block context,
	// which will be written to by InsertBlock. This is a runMsgs cache context, which is
	// only written once ProcessPayloadEnvelope executes without error.
	k.spf.SetFinalizeBlockContext(ctx)
	defer k.spf.SetLatestQueryContext(ctx)
	k.chain.PrimePlugins(ctx)

	// Insert the finalized block and set the chain head.
	if err = k.chain.InsertBlockAndSetHead(block); err != nil {
		return nil, err
	}

	return &evmtypes.WrappedPayloadEnvelopeResponse{}, nil
}

// EthTransaction implements the MsgServer interface. It is intentionally a no-op, but is required
// for the cosmos-sdk to not freak out.
func (k *Keeper) EthTransaction(
	context.Context, *evmtypes.WrappedEthereumTransaction,
) (*evmtypes.WrappedEthereumTransactionResult, error) {
	panic("intentionally not implemented")
}
