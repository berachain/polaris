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


type transientChange struct {
	account 	*common.Address
	key 		common.Hash
	previous 	common.Hash
}

// `transient` is a `Store` that tracks the transient state.
type transient struct {
	ds.Stack[transientChange]
}

// `NewTransientStorage` returns a new `transient` journal.
func NewTransient() *transient {
	return &transient{
		stack.New[transientChange](initCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (t *transient) RegistryKey() string {
	return transientRegistryKey
}


// `AddTransient` adds a transient change to the `transient` store.
func (t *transient) AddTransient(account common.Address, key common.Hash, previous common.Hash) {
	t.Push(transientChange{
		account: 	&account,
		key: 		key,
		previous: 	previous,
	})
}

// `GetTransient` returns previous transient storage state for a given account + key.
func (t *transient) GetPrevTransient(account common.Address, key common.Hash) *common.Hash {
	size := t.Size()
	for i := size - 1; i >= 0; i-- {
		itemAt := t.PeekAt(i)
		if itemAt.account == &account && itemAt.key == key {
			return &itemAt.previous
		}
	}
	return nil
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
	t.Stack =  stack.New[transientChange](initCapacity)
}