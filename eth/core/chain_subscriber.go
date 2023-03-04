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

	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
)

type ChainSubscriber interface {
	SubscribeRemovedLogsEvent(chan<- RemovedLogsEvent) event.Subscription
	SubscribeChainEvent(chan<- ChainEvent) event.Subscription
	SubscribeChainHeadEvent(chan<- ChainHeadEvent) event.Subscription
}

// SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
func (bc *blockchain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
	// return bc.scope.Track(bc.rmLogsFeed.Subscribe(ch))
	return nil
}

// SubscribeChainEvent registers a subscription of ChainEvent.
func (bc *blockchain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
	// return bc.scope.Track(bc.chainFeed.Subscribe(ch))
	return nil
}

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *blockchain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	// TODO: synchronize chain head feed.
	return bc.scope.Track(bc.chainHeadFeed.Subscribe(ch))
}

// SubscribeChainSideEvent registers a subscription of ChainSideEvent.
func (bc *blockchain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
	// return bc.scope.Track(bc.chainSideFeed.Subscribe(ch))
	return nil
}

// SubscribeLogsEvent registers a subscription of []*types.Log.
func (bc *blockchain) SubscribeLogsEvent(ch chan<- []*coretypes.Log) event.Subscription {
	// return bc.scope.Track(bc.logsFeed.Subscribe(ch))
	return nil
}

// SubscribeBlockProcessingEvent registers a subscription of bool where true means
// block processing has started while false means it has stopped.
func (bc *blockchain) SubscribeBlockProcessingEvent(ch chan<- bool) event.Subscription {
	// return bc.scope.Track(bc.blockProcFeed.Subscribe(ch))
	return nil
}
