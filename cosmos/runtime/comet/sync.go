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

package comet

import (
	"context"

	"github.com/berachain/polaris/eth/polar"

	cmtclient "github.com/cometbft/cometbft/rpc/client"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/ethereum/go-ethereum"
)

var _ polar.SyncStatusProvider = (*cometSyncStatus)(nil)

// cometSyncStatus is a view of the sync status from CometBFT.
type cometSyncStatus struct {
	cometClient client.CometRPC
}

// NewSyncProvider returns a new cometSyncStatus.
func NewSyncProvider(clientCtx client.Context) polar.SyncStatusProvider {
	return &cometSyncStatus{
		cometClient: clientCtx.Client,
	}
}

// PeerCount returns the number of peers currently connected to the node.
func (sv *cometSyncStatus) PeerCount(ctx context.Context) (uint64, error) {
	result, err := sv.cometClient.(cmtclient.NetworkClient).NetInfo(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(result.NPeers), nil
}

// Listening returns whether the node is currently listening for incoming
// connections from other nodes.
func (sv *cometSyncStatus) Listening(ctx context.Context) (bool, error) {
	result, err := sv.cometClient.(cmtclient.NetworkClient).NetInfo(ctx)
	if err != nil {
		return false, err
	}

	return result.Listening, nil
}

// SyncProgress returns the current sync progress.
func (sv *cometSyncStatus) SyncProgress(ctx context.Context) (ethereum.SyncProgress, error) {
	resultStatus, err := sv.cometClient.(cmtclient.StatusClient).Status(ctx)
	if err != nil {
		return ethereum.SyncProgress{}, err
	}

	resultInfo, err := sv.cometClient.(cmtclient.HistoryClient).
		BlockchainInfo(
			ctx,
			resultStatus.SyncInfo.EarliestBlockHeight,
			resultStatus.SyncInfo.LatestBlockHeight,
		)
	if err != nil {
		return ethereum.SyncProgress{}, err
	}

	return ethereum.SyncProgress{
		StartingBlock: uint64(resultStatus.SyncInfo.EarliestBlockHeight),
		CurrentBlock:  uint64(resultInfo.LastHeight),
		HighestBlock:  uint64(resultStatus.SyncInfo.LatestBlockHeight),
	}, nil
}
