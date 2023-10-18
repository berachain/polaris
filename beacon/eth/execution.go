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
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"

	"pkg.berachain.dev/polaris/beacon/log"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

type (
	// BuilderAPI represents the `Miner` that exists on the backend of the execution layer.
	BuilderAPI interface {
		BuildBlock(context.Context, *miner.BuildPayloadArgs) (*engine.ExecutionPayloadEnvelope, error)
		Etherbase(context.Context) (common.Address, error)
		BlockByNumber(uint64) *coretypes.Block
		CurrentBlock(ctx context.Context) *coretypes.Block
		ConsensusAPI
	}

	// TxPool represents the `TxPool` that exists on the backend of the execution layer.
	TxPoolAPI interface {
		Add([]*coretypes.Transaction, bool, bool) []error
		Stats() (int, int)
		SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
	}

	ConsensusAPI interface {
		NewPayloadV2(ctx context.Context,
			params engine.ExecutableData,
		) (engine.PayloadStatusV1, error)
		ForkchoiceUpdatedV2(ctx context.Context,
			update engine.ForkchoiceStateV1,
			payloadAttributes *engine.PayloadAttributes) (engine.ForkChoiceResponse, error)
	}
)

// ExecutionClient represents the execution layer client.
type ExecutionClient struct {
	BlockBuilder BuilderAPI
	TxPool       TxPoolAPI
	Consensus    ConsensusAPI
}

// NewRemoteExecutionClient creates a new remote execution client.
func NewRemoteExecutionClient(dialURL string, logger log.Logger) (*ExecutionClient, error) {
	var (
		client  *ethclient.Client
		chainID *big.Int
		err     error
	)

	ctx := context.Background()
	for i := 0; i < 100; func() { i++; time.Sleep(time.Second) }() {
		logger.Info("Attempting to connect to execution layer", "attempt", i+1, "dial-url", dialURL)
		client, err = ethclient.DialContext(ctx, dialURL)
		if err != nil {
			continue
		}
		chainID, err = client.ChainID(ctx)
		if err != nil {
			continue
		}
		logger.Info("Successfully connected to execution layer", "ChainID", chainID)
		break
	}
	if client == nil || err != nil {
		return nil, fmt.Errorf("failed to establish connection to execution layer: %w", err)
	}

	return &ExecutionClient{
		BlockBuilder: &builderAPI{Client: client},
		TxPool:       &txPoolAPI{Client: client},
		Consensus:    &builderAPI{Client: client},
	}, nil
}
