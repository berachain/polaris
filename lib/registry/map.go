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

package registry

import (
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

// `mapRegistry` is a simple implementation of `Registry` that uses a map as the underlying data
// structure.
type mapRegistry[K comparable, T libtypes.Registrable[K]] struct {
	// items is the map of items in the registry.
	items map[K]T
}

// `NewMap` creates and returns a new `mapRegistry`.
//
//nolint:revive // only used as Registry interface.
func NewMap[K comparable, T libtypes.Registrable[K]]() *mapRegistry[K, T] {
	return &mapRegistry[K, T]{
		items: make(map[K]T),
	}
}

// `Get` returns an item using its ID.
func (mr *mapRegistry[K, T]) Get(id K) T {
	return mr.items[id]
}

// `Register` adds an item to the registry.
func (mr *mapRegistry[K, T]) Register(item T) error {
	mr.items[item.RegistryKey()] = item
	return nil
}

// `Remove` removes an item from the registry.
func (mr *mapRegistry[K, T]) Remove(id K) {
	delete(mr.items, id)
}

// `Has` returns true if the item exists in the registry.
func (mr *mapRegistry[K, T]) Has(id K) bool {
	_, ok := mr.items[id]
	return ok
}

// `Iterate` returns the underlying map.
func (mr *mapRegistry[K, T]) Iterate() map[K]T {
	return mr.items
}
