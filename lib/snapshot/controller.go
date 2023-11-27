// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package snapshot

import (
	"github.com/berachain/polaris/lib/ds"
	"github.com/berachain/polaris/lib/ds/stack"
	"github.com/berachain/polaris/lib/registry"
	libtypes "github.com/berachain/polaris/lib/types"
)

// initJournalCapacity is the initial capacity of the `journal` stack.
// TODO: determine better initial capacity.
const initJournalCapacity = 32

// revision is a snapshot revision, which holds all `Controllable`s' snapshot ids.
// Specifically, it is a map of a `Controllable`'s `RegistryKey` to its corresponding current
// snapshot revision id.
type revision[K comparable] map[K]int

// controller conforms to the `libtypes.Controller` interface and is used to register and sync
// snapshotting across multiple `libtypes.Controllable` objects.
type controller[K comparable, T libtypes.Controllable[K]] struct {
	// Registry stores the `Controllable` objects.
	libtypes.Registry[K, T]

	// journal is a stack of `revision`s. All `Controllable` objects are currently on the
	// snapshot revision id at the top (`Peek()`) of the journal stack. If the stack is empty, all
	// Controllable objects have no snapshot.
	journal ds.Stack[revision[K]]
}

// NewController returns a new `Controller` object.
func NewController[K comparable, T libtypes.Controllable[K]]() libtypes.Controller[K, T] {
	return &controller[K, T]{
		Registry: registry.NewMap[K, T](),
		journal:  stack.New[revision[K]](initJournalCapacity),
	}
}

// Snapshot takes a snapshot for all controllable objects and returns the controller's snapshot
// id.
//
// Snapshot implements `libtypes.Snapshottable`.
func (c *controller[K, T]) Snapshot() int {
	newRevision := make(revision[K])
	for key, controllable := range c.Iterate() {
		newRevision[key] = controllable.Snapshot()
	}

	// push the new revision and return the size BEFORE snapshot
	return c.journal.Push(newRevision) - 1
}

// RevertToSnapshot reverts all controllable objects to their own snapshot id corresponding to
// id.
//
// RevertToSnapshot implements `libtypes.Snapshottable`.
func (c *controller[K, T]) RevertToSnapshot(id int) {
	// id is the new size of the journal we want to maintain.
	for key, revertedSnapshot := range c.journal.PopToSize(id) {
		// revert all `Controllable` objects to their corresponding revision
		c.Get(key).RevertToSnapshot(revertedSnapshot)
	}
}

// Finalize writes all the controllables controlled by this controller.
//
// Finalize implements `libtypes.Controller`.
func (c *controller[K, T]) Finalize() {
	for _, controllable := range c.Iterate() {
		controllable.Finalize()
	}
}
