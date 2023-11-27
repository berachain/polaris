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

package registry

import (
	libtypes "github.com/berachain/polaris/lib/types"
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
