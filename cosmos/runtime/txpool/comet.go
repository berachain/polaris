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
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	defaultCacheSize = 4096
)

// CometRemoteCache is used to mark which txs are added to our Polaris node remotely from
// Comet CheckTX and when.
type CometRemoteCache interface {
	IsRemoteTx(txHash common.Hash) bool
	MarkRemoteSeen(txHash common.Hash)
	TimeFirstSeen(txHash common.Hash) int64 // UNIX timestamp
	DropRemoteTx(txHash common.Hash)
}

// Thread-safe implementation of CometRemoteCache
type cometRemoteCache struct {
	timeInserted   map[common.Hash]int64
	timeInsertedMu sync.RWMutex
}

func newCometRemoteCache() *cometRemoteCache {
	return &cometRemoteCache{
		timeInserted: make(map[common.Hash]int64, defaultCacheSize),
	}
}

func (crc *cometRemoteCache) IsRemoteTx(txHash common.Hash) bool {
	crc.timeInsertedMu.RLock()
	defer crc.timeInsertedMu.RUnlock()
	_, ok := crc.timeInserted[txHash]
	return ok
}

func (crc *cometRemoteCache) MarkRemoteSeen(txHash common.Hash) {
	// Record the time the tx was inserted from Comet successfully.
	crc.timeInsertedMu.Lock()
	crc.timeInserted[txHash] = time.Now().Unix()
	crc.timeInsertedMu.Unlock()
}

func (crc *cometRemoteCache) TimeFirstSeen(txHash common.Hash) int64 {
	crc.timeInsertedMu.RLock()
	defer crc.timeInsertedMu.RUnlock()
	return crc.timeInserted[txHash]
}

func (crc *cometRemoteCache) DropRemoteTx(txHash common.Hash) {
	crc.timeInsertedMu.Lock()
	delete(crc.timeInserted, txHash)
	crc.timeInsertedMu.Unlock()
}
