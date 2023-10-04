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
)

func (k *Keeper) ProcessPayloadEnvelope(
	ctx context.Context, msg *evmtypes.WrappedPayloadEnvelope,
) (*evmtypes.WrappedPayloadEnvelopeResponse, error) {
	envelope := engine.ExecutionPayloadEnvelope{}
	err := envelope.UnmarshalJSON(msg.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload envelope: %w", err)
	}

	sCtx := sdk.UnwrapSDKContext(ctx)
	gasMeter := sCtx.GasMeter()

	// x := new(common.Hash)
	fmt.Println("PROCESSED", envelope.ExecutionPayload)
	fmt.Println("PROCESSED", envelope.ExecutionPayload.BlockHash)
	block, err := engine.ExecutableDataToBlock(*envelope.ExecutionPayload, nil, nil)
	if err != nil {
		k.Logger(sCtx).Error("failed to build evm block", "err", err)
		return nil, err
	}

	fmt.Println(block.Withdrawals(), block.Header().WithdrawalsHash, "BIyNG")

	// bz, _ := block.Header().MarshalJSON()

	if err = k.polaris.Blockchain().InsertBlockWithoutSetHead(block); err != nil {
		return nil, err
	}

	// Consume the gas used by the execution of the ethereum block.
	gasMeter.ConsumeGas(block.GasUsed(), "block gas used")

	return &evmtypes.WrappedPayloadEnvelopeResponse{}, nil
}
