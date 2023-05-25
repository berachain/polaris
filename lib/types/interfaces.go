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

package types

import "context"

// Cloneable is an interface that defines a `Clone` method.
type Cloneable[T any] interface {
	Clone() T
}

// Snapshottable is an interface that defines methods for snapshotting and reverting
// a logical unit of data.
type Snapshottable interface {
	// Snapshot returns an identifier for the current revision of the data.
	Snapshot() int

	// RevertToSnapshot reverts the data to a previous version
	RevertToSnapshot(int)
}

// Registrable is an interface that all objects that can be registered in a
// Registry must implement.
type Registrable[K comparable] interface {
	// RegistryKey returns the key that will be used to register the object.
	RegistryKey() K
}

// Registry is an interface that all objects that can be used as a registry
// must implement.
type Registry[K comparable, T Registrable[K]] interface {
	// Get return an item using its ID. It returns nil if the ID does not exist.
	Get(K) T

	// Register adds an item to the registry, indexed on the item's `RegistryKey`.
	Register(T) error

	// Remove removes an item from the registry.
	Remove(K)

	// Has returns true if the item exists in the registry.
	Has(K) bool

	// Iterate returns an iterable map of the registry.
	Iterate() map[K]T
}

// Controllable defines a type which can be controlled.
type Controllable[K comparable] interface {
	Snapshottable
	Registrable[K]
	Finalizeable
}

// Controller is an interface for controller types.
type Controller[K comparable, T Controllable[K]] interface {
	Snapshottable
	Registry[K, T]
	Finalizeable
}

// Finalizeable is an interface that defines a `Finalize` method.
type Finalizeable interface {
	// Finalize finalizes the state of the object.
	Finalize()
}

// Resettable is an interface that defines a `Reset` method. The `Reset` method is usually used
// to reset the state per transaction.
type Resettable interface {
	// Reset resets the state of the object with the new given context.
	Reset(context.Context)
}

// Preparable is an interface that defines a `Prepare` method. The `Prepare` method is usually
// used to prepare the state per block.
type Preparable interface {
	// Prepare prepares the state of the object with the new given context.
	Prepare(context.Context)
}
