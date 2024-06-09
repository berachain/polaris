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
