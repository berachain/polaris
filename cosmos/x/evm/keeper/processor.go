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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core/types"
)

func (k *Keeper) ProcessPayloadEnvelope(
	ctx context.Context, msg *evmtypes.WrappedPayloadEnvelope,
) (*evmtypes.WrappedPayloadEnvelopeResponse, error) {
	var (
		err      error
		block    *types.Block
		envelope engine.ExecutionPayloadEnvelope
	)
	// TODO: maybe we just consume the block gas limit and call it a day?
	sCtx := sdk.UnwrapSDKContext(ctx)
	gasMeter := sCtx.GasMeter()
	blockGasMeter := sCtx.BlockGasMeter()

	// Reset GasMeter to 0.
	gasMeter.RefundGas(gasMeter.GasConsumed(), "reset before evm block")
	blockGasMeter.RefundGas(blockGasMeter.GasConsumed(), "reset before evm block")
	defer gasMeter.ConsumeGas(gasMeter.GasConsumed(), "reset after evm")

	if err = envelope.UnmarshalJSON(msg.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload envelope: %w", err)
	}

	if block, err = engine.ExecutableDataToBlock(*envelope.ExecutionPayload, nil, nil); err != nil {
		k.Logger(sCtx).Error("failed to build evm block", "err", err)
		return nil, err
	}

	// Prepare should be moved to the blockchain? THIS IS VERY HOOD YES NEEDS TO BE MOVED.
	k.chain.PreparePlugins(ctx)
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
