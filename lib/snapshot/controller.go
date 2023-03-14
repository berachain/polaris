// SPDX-License-Identifier: Apache-2.0
//

package snapshot

import (
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
	"pkg.berachain.dev/polaris/lib/registry"
	libtypes "pkg.berachain.dev/polaris/lib/types"
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
