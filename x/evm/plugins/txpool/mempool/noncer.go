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

package mempool

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// `NonceRetriever` is an interface that allows for the
// TxPool plugin to retrieve the nonce of an account.
type NonceRetriever interface {
	GetNonce(addr common.Address) uint64
}

// `noncer` is a struct that implements the NonceRetriever interface
// and caches the nonce of an account.
type noncer struct {
	fallback NonceRetriever
	nonces   map[common.Address]uint64
	lock     sync.Mutex
}

// `newNoncer` returns a new noncer.
func newNoncer(nr NonceRetriever) *noncer {
	return &noncer{
		fallback: nr,
		nonces:   make(map[common.Address]uint64),
	}
}

// `GetNonce` returns the nonce of an account.
func (txn *noncer) get(addr common.Address) uint64 {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	if _, ok := txn.nonces[addr]; !ok {
		if nonce := txn.fallback.GetNonce(addr); nonce != 0 {
			txn.nonces[addr] = nonce
		}
	}
	return txn.nonces[addr]
}

// `SetNonce` sets the nonce of an account.
func (txn *noncer) set(addr common.Address, nonce uint64) {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	txn.nonces[addr] = nonce
}

// `SetIfLower` sets the nonce of an account if the nonce is lower than the
// current nonce.
func (txn *noncer) setIfLower(addr common.Address, txNonce uint64) {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	if _, ok := txn.nonces[addr]; !ok {
		if sdbNonce := txn.fallback.GetNonce(addr); sdbNonce != 0 {
			txn.nonces[addr] = sdbNonce
		}
	}
	if txn.nonces[addr] <= txNonce {
		return
	}
	txn.nonces[addr] = txNonce
}
