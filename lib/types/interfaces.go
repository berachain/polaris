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

package types

import "context"

// `Cloneable` is an interface that defines a `Clone` method.
type Cloneable[T any] interface {
	Clone() T
}

// `Snapshottable` is an interface that defines methods for snapshotting and reverting
// a logical unit of data.
type Snapshottable interface {
	// `Snapshot` returns an identifier for the current revision of the data.
	Snapshot() int

	// `RevertToSnapshot` reverts the data to a previous version
	RevertToSnapshot(int)
}

// `Registrable` is an interface that all objects that can be registered in a
// `Registry` must implement.
type Registrable[K comparable] interface {
	// `RegistryKey` returns the key that will be used to register the object.
	RegistryKey() K
}

// `Registry` is an interface that all objects that can be used as a registry
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

// `Controllable` defines a type which can be controlled.
type Controllable[K comparable] interface {
	Snapshottable
	Registrable[K]
	Finalizeable
}

// `Controller` is an interface for controller types.
type Controller[K comparable, T Controllable[K]] interface {
	Snapshottable
	Registry[K, T]
	Finalizeable
}

// `Finalizeable` is an interface that defines a `Finalize` method.
type Finalizeable interface {
	// `Finalize` finalizes the state of the object.
	Finalize()
}

// `Resettable` is an interface that defines a `Reset` method.
type Resettable interface {
	// `Reset` resets the state of the object with the new given context.
	Reset(context.Context)
}
