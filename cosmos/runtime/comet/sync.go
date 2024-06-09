// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
