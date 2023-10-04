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

package core

import (
	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

type ChainSubscriber interface {
	SubscribeRemovedLogsEvent(chan<- RemovedLogsEvent) event.Subscription // currently not used
	SubscribeChainEvent(chan<- ChainEvent) event.Subscription
	SubscribeChainHeadEvent(chan<- ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription // currently not used
	SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription
	SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription
	EmitCurrentBlockEvents()
}

// SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
func (bc *blockchain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
	return bc.scope.Track(bc.rmLogsFeed.Subscribe(ch))
}

// SubscribeChainEvent registers a subscription of ChainEvent.
func (bc *blockchain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
	return bc.scope.Track(bc.chainFeed.Subscribe(ch))
}

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *blockchain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	return bc.scope.Track(bc.chainHeadFeed.Subscribe(ch))
}

// SubscribeChainSideEvent registers a subscription of ChainSideEvent.
func (bc *blockchain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
	return bc.scope.Track(bc.chainSideFeed.Subscribe(ch))
}

// SubscribeLogsEvent registers a subscription of []*types.Log.
func (bc *blockchain) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return bc.scope.Track(bc.logsFeed.Subscribe(ch))
}

// SubscribePendingLogsEvent registers a subscription of []*types.Log.
func (bc *blockchain) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return bc.scope.Track(bc.pendingLogsFeed.Subscribe(ch))
}

// EmitCurrentBlockEvents emits chain events for the current block.
func (bc *blockchain) EmitCurrentBlockEvents() {
	// Send the pending/current logs on the logs feeds.
	logs, ok := utils.GetAs[[]*types.Log](bc.currentLogs.Load())
	if ok {
		bc.pendingLogsFeed.Send(logs)
		if len(logs) > 0 {
			bc.logsFeed.Send(logs)
		}
	}

	// Send the current block on the chain(head) feeds.
	if block, ok := utils.GetAs[*types.Block](bc.currentBlock.Load()); ok {
		bc.chainFeed.Send(ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
		bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})
	}
}
