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

package duocachekv

import (
	"bytes"

	storetypes "cosmossdk.io/store/types"
)

// `store` defines a cachekv store that supports reading from/writing to 2 synchronized cachekv
// stores. The first store is the source-of-truth store which will write to its parent, and the
// second store is the revision store which is used to store the checkpoint.
type store struct {
	// `CacheKVStore` is the source-of-truth store.
	storetypes.CacheKVStore
	// `revision` is the kv store that is stored as a checkpoint.
	revision storetypes.KVStore
}

// `NewStoreFrom` creates and returns a new `store` from a given source-of-truth store `sot` and
// revision store `revision`.
func NewStoreFrom(sot storetypes.CacheKVStore, revision storetypes.KVStore) storetypes.CacheKVStore {
	return &store{
		CacheKVStore: sot,
		revision:     revision,
	}
}

// Get returns nil if key doesn't exist. Panics on nil key.
func (s *store) Get(key []byte) []byte {
	sotVal, revisionVal := s.CacheKVStore.Get(key), s.revision.Get(key)
	if sotVal == nil || revisionVal == nil || !bytes.Equal(sotVal, revisionVal) {
		panic("inconsistent state")
	}
	return sotVal
}

// Has checks if a key exists. Panics on nil key.
func (s *store) Has(key []byte) bool {
	sotHas, revisionHas := s.CacheKVStore.Has(key), s.revision.Has(key)
	if sotHas != revisionHas {
		panic("inconsistent state")
	}
	return sotHas
}

// Set sets the key. Panics on nil key or value.
func (s *store) Set(key, value []byte) {
	s.CacheKVStore.Set(key, value)
	s.revision.Set(key, value)
}

// Delete deletes the key. Panics on nil key.
func (s *store) Delete(key []byte) {
	s.CacheKVStore.Delete(key)
	s.revision.Delete(key)
}
