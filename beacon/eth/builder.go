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

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

type builderAPI struct {
	*ethclient.Client
}

func (api *builderAPI) BuildBlock(
	ctx context.Context,
	attrs *miner.BuildPayloadArgs,
) (*engine.ExecutionPayloadEnvelope, error) {
	return api.Client.BuildBlock(ctx, attrs)
}

func (api *builderAPI) NewPayloadV3(ctx context.Context, params engine.ExecutableData,
	versionedHashes []common.Hash, beaconRoot *common.Hash) (engine.PayloadStatusV1, error) {
	return api.Client.NewPayloadV3(ctx, params, versionedHashes, beaconRoot)
}

func (api *builderAPI) Etherbase() common.Address {
	etherbase, _ := api.Client.Etherbase(context.Background())
	return etherbase
}

func (api *builderAPI) BlockByNumber(num uint64) *coretypes.Block {
	b, _ := api.Client.BlockByNumber(context.Background(), new(big.Int).SetUint64(num))
	return b
}
