// SPDX-License-Identifier: Apache-2.0
//

package registry

import (
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// mapRegistry is a simple implementation of `Registry` that uses a map as the underlying data
// structure.
type mapRegistry[K comparable, T libtypes.Registrable[K]] struct {
	// items is the map of items in the registry.
	items map[K]T
}

// NewMap creates and returns a new `mapRegistry`.
//
//nolint:revive // only used as Registry interface.
func NewMap[K comparable, T libtypes.Registrable[K]]() *mapRegistry[K, T] {
	return &mapRegistry[K, T]{
		items: make(map[K]T),
	}
}

// Get returns an item using its ID.
func (mr *mapRegistry[K, T]) Get(id K) T {
	return mr.items[id]
}

// Register adds an item to the registry.
func (mr *mapRegistry[K, T]) Register(item T) error {
	mr.items[item.RegistryKey()] = item
	return nil
}

// Remove removes an item from the registry.
func (mr *mapRegistry[K, T]) Remove(id K) {
	delete(mr.items, id)
}

// Has returns true if the item exists in the registry.
func (mr *mapRegistry[K, T]) Has(id K) bool {
	_, ok := mr.items[id]
	return ok
}

// Iterate returns the underlying map.
func (mr *mapRegistry[K, T]) Iterate() map[K]T {
	return mr.items
}
