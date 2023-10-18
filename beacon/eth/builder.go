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

package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/rpc"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

type builderAPI struct {
	*ethclient.Client
}

func NewBuilderAPI(c *ethclient.Client) BuilderAPI {
	return &builderAPI{c}
}

func (api *builderAPI) CurrentBlock(ctx context.Context) *coretypes.Block {
	b, err := api.Client.BlockByNumber(ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		b, err = api.Client.BlockByNumber(ctx, nil)
		if err != nil {
			return nil
		}
	}
	return b
}

func (api *builderAPI) BlockByNumber(num uint64) *coretypes.Block {
	b, _ := api.Client.BlockByNumber(context.Background(), new(big.Int).SetUint64(num))
	return b
}

func (api *builderAPI) Etherbase(ctx context.Context) (common.Address, error) {
	var miner common.Address
	if err := api.Client.Client().CallContext(ctx, &miner, "miner_etherbase"); err != nil {
		return common.Address{}, err
	}
	return miner, nil
}

func (api *builderAPI) BuildBlock(
	ctx context.Context, attrs *miner.BuildPayloadArgs,
) (*engine.ExecutionPayloadEnvelope, error) {
	var payload engine.ExecutionPayloadEnvelope
	if err := api.Client.Client().CallContext(ctx, &payload, "miner_buildBlock", attrs); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (api *builderAPI) NewPayloadV2(
	ctx context.Context, params engine.ExecutableData,
) (engine.PayloadStatusV1, error) {
	var payloadStatus engine.PayloadStatusV1
	err := api.Client.Client().CallContext(ctx, &payloadStatus, "engine_newPayloadV2", params)
	return payloadStatus, err
}

func (api *builderAPI) ForkchoiceUpdatedV2(
	ctx context.Context, update engine.ForkchoiceStateV1, payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	var response engine.ForkChoiceResponse
	err := api.Client.Client().CallContext(
		ctx, &response, "engine_forkchoiceUpdatedV2", update, payloadAttributes)
	return response, err
}
