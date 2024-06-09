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

package journal

import (
	"github.com/berachain/polaris/lib/ds/stack"
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

// transientState is a representation of EIP-1153 "Transient Storage".
type transientState map[common.Address]map[common.Hash]common.Hash

// Set sets the transient storage state `value` for `key` at the given `addr`.
func (t transientState) Set(addr common.Address, key, value common.Hash) {
	if _, ok := t[addr]; !ok {
		t[addr] = make(map[common.Hash]common.Hash)
	}
	t[addr][key] = value
}

// Get gets the transient storage state for `key` at the given `addr`.
func (t transientState) Get(addr common.Address, key common.Hash) common.Hash {
	val, ok := t[addr]
	if !ok {
		return common.Hash{}
	}
	return val[key]
}

// Copy does a deep copy of the transientState.
func (t transientState) Copy() transientState {
	storage := make(transientState)
	for storKey, storVal := range t {
		valDeepCopy := make(map[common.Hash]common.Hash, len(storVal))
		for key, val := range storVal {
			valDeepCopy[key] = val
		}
		storage[storKey] = valDeepCopy
	}
	return storage
}

type TransientStorage interface {
	// TransientStorage implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// TransientStorage implements `libtypes.Cloneable`.
	libtypes.Cloneable[TransientStorage]
	// GetTransientState returns a transient storage for a given account.
	GetTransientState(addr common.Address, key common.Hash) common.Hash
	// SetTransientState sets a given transient storage change to the transient journal.
	SetTransientState(addr common.Address, key, value common.Hash)
}

// `transientStorage` is a journal that tracks the transient state.
type transientStorage struct {
	baseJournal[transientState]
}

// `NewTransientStorage` returns a new `transient` journal.
func NewTransientStorage() TransientStorage {
	return &transientStorage{
		newBaseJournal[transientState](initCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (t *transientStorage) RegistryKey() string {
	return transientRegistryKey
}

// `AddTransient` adds a transient change to the `transient` store.
func (t *transientStorage) SetTransientState(addr common.Address, key, value common.Hash) {
	currentState := t.Peek()
	if currentState.Get(addr, key) == value {
		return
	}
	currentState = currentState.Copy()
	currentState.Set(addr, key, value)
	t.Push(currentState)
}

// `GetTransient` returns previous transient storage state for a given account + key.
func (t *transientStorage) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	return t.Peek().Get(addr, key)
}

// `Finalize` implements `libtypes.Controllable`.
func (t *transientStorage) Finalize() {
	t.Stack = stack.New[transientState](initCapacity)
}

// Clone implements `libtypes.Cloneable`.
func (t *transientStorage) Clone() TransientStorage {
	clone := &transientStorage{
		newBaseJournal[transientState](t.Capacity()),
	}

	// copy every individual transient state
	for i := 0; i < t.Size(); i++ {
		clone.Push(t.PeekAt(i).Copy())
	}

	return clone
}
