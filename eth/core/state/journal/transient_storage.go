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

package journal

import (
	"github.com/ethereum/go-ethereum/common"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
)

// transientStorage is a representation of EIP-1153 "Transient Storage".
type transientStorage map[common.Address]map[common.Hash]common.Hash

// Set sets the transient-storage `value` for `key` at the given `addr`.
func (t transientStorage) Set(addr common.Address, key, value common.Hash) {
	if _, ok := t[addr]; !ok {
		t[addr] = make(map[common.Hash]common.Hash)
	}
	t[addr][key] = value
}

// Get gets the transient storage for `key` at the given `addr`.
func (t transientStorage) Get(addr common.Address, key common.Hash) common.Hash {
	val, ok := t[addr]
	if !ok {
		return common.Hash{}
	}
	return val[key]
}

// Copy does a deep copy of the transientStorage
func (t transientStorage) Copy() transientStorage {
	storage := make(transientStorage)
	for storKey, storVal := range t {
		valDeepCopy := make(map[common.Hash]common.Hash, len(storVal))
		for key, val := range storVal {
			valDeepCopy[key] = val
		}
		storage[storKey] = valDeepCopy
	}
	return storage
}

// `transient` is a `Store` that tracks the transient state.
type transient struct {
	ds.Stack[transientStorage]
}

// `NewTransientStorage` returns a new `transient` journal.
func NewTransient() *transient {
	return &transient{
		stack.New[transientStorage](initCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (t *transient) RegistryKey() string {
	return transientRegistryKey
}


// `AddTransient` adds a transient change to the `transient` store.
func (t *transient) AddTransient(account common.Address, key common.Hash, val common.Hash) {
	currentState := t.Peek().Copy()
	currentState.Set(account, key, val)
	t.Push(currentState)
}

// `GetTransient` returns previous transient storage state for a given account + key.
func (t *transient) GetTransient(account common.Address, key common.Hash) common.Hash {
	return t.Peek().Get(account, key)
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (t *transient) Snapshot() int {
	return t.Size()
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (t *transient) RevertToSnapshot(id int) {
	t.PopToSize(id)
}

// `Finalize` implements `libtypes.Controllable`.
func (t *transient) Finalize() {
	t.Stack =  stack.New[transientStorage](initCapacity)
}