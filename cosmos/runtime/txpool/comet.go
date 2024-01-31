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

package txpool

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
)

const (
	defaultCacheSize = 100000
)

// CometRemoteCache is used to mark which txs are added to our Polaris node remotely from
// Comet CheckTX and when.
type CometRemoteCache interface {
	IsRemoteTx(txHash common.Hash) bool
	MarkRemoteSeen(txHash common.Hash) bool
	TimeFirstSeen(txHash common.Hash) int64 // Unix timestamp
}

// Thread-safe implementation of CometRemoteCache.
type cometRemoteCache struct {
	timeInserted *lru.Cache[common.Hash, int64]
}

func newCometRemoteCache() *cometRemoteCache {
	return &cometRemoteCache{

		timeInserted: lru.NewCache[common.Hash, int64](defaultCacheSize),
	}
}

func (crc *cometRemoteCache) IsRemoteTx(txHash common.Hash) bool {
	return crc.timeInserted.Contains(txHash)
}

// Record the time the tx was inserted from Comet successfully.
func (crc *cometRemoteCache) MarkRemoteSeen(txHash common.Hash) bool {
	if !crc.timeInserted.Contains(txHash) {
		crc.timeInserted.Add(txHash, time.Now().Unix())
		return true
	}
	return false
}

func (crc *cometRemoteCache) TimeFirstSeen(txHash common.Hash) int64 {
	i, _ := crc.timeInserted.Get(txHash)
	return i
}
