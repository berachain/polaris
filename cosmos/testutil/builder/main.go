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

package main

import (
	"context"
	"math/big"
	"os"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/miner"
)

// Block Builder Test.
func main() {
	// Configure logger, client, etherbase.
	logger := log.NewLogger(os.Stdout).With("module", "main")
	client, _ := ethclient.Dial("http://localhost:8545")
	etherbase, _ := client.Etherbase(context.Background())

	// Get Parent Header
	latestBlockNumber, _ := client.BlockNumber(context.Background())
	parentHeader, _ := client.HeaderByNumber(context.Background(),
		big.NewInt(int64(latestBlockNumber)))
	logger.Info("parent located", "parent-header", parentHeader.Hash(),
		"parent-header-time", parentHeader.Time, "parent-header-number", parentHeader.Number)

	// Build block using the miner on the execution client.
	builderResponse, err := client.BuildBlock(context.Background(), &miner.BuildPayloadArgs{
		Timestamp:    parentHeader.Time + 5, //nolint:gomnd // testing script.
		FeeRecipient: etherbase,
		Random:       common.Hash{},
		Withdrawals:  nil,
		Parent:       parentHeader.Hash(),
		BeaconRoot:   nil,
	})
	logger.Info("block built", "builder-response", builderResponse, "err", err)

	// SubmitNewPayload
	payloadResponse, err := client.NewPayloadV3(context.Background(),
		*builderResponse.ExecutionPayload, nil, nil)
	logger.Info("block submitted to chain", "payload-response", payloadResponse, "err", err)
}
