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

package snapmulti

import (
	"cosmossdk.io/store/cachekv"
	storetypes "cosmossdk.io/store/types"

	polarisstore "pkg.berachain.dev/polaris/cosmos/store"
	"pkg.berachain.dev/polaris/cosmos/store/duocachekv"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
	"pkg.berachain.dev/polaris/lib/utils"
)

const (
	storeRegistryKey    = `snapmultistore`
	initJournalCapacity = 16
)

// `mapMultiStore` represents a cached multistore, which is just a map of store keys to its
// corresponding cache kv store currently being used.
type mapMultiStore map[storetypes.StoreKey]storetypes.CacheKVStore

// `store` is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It journals revisions by cache-wrapping the cachekv stores on a call to `Snapshot`. In this
// store's lifecycle, any operations done before the first call to snapshot will be enforced on the
// root `mapMultiStore`.
type store struct {
	// `MultiStore` is the underlying multistore
	storetypes.MultiStore
	// `root` is the mapMultiStore used before the first snapshot is called
	root mapMultiStore
	// `sot` is the source of truth mapMultiStore
	sot mapMultiStore
	// `journal` holds the snapshots of cachemultistores
	journal ds.Stack[mapMultiStore]
}

// `NewStoreFrom` creates and returns a new `store` from a given Multistore `ms`.
func NewStoreFrom(ms storetypes.MultiStore) polarisstore.ControllableMulti {
	return &store{
		MultiStore: ms,
		root:       make(mapMultiStore),
		sot:        make(mapMultiStore),
		journal:    stack.New[mapMultiStore](initJournalCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (s *store) RegistryKey() string {
	return storeRegistryKey
}

// `GetCommittedKVStore` returns the KV Store from the given Multistore. This function follows
// the Multistore's normal `GetKVStore` code path.
func (s *store) GetCommittedKVStore(key storetypes.StoreKey) storetypes.KVStore {
	return s.MultiStore.GetKVStore(key)
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	cms := s.journal.Peek()
	if cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// check if cache kv store already used
	if cacheKVStore, found := cms[key]; found {
		return cacheKVStore
	}

	// get kvstore from mapMultiStore and set duocachekv to memory
	sotKV := cachekv.NewStore(s.GetCommittedKVStore(key))
	cms[key] = duocachekv.NewStoreFrom(sotKV, sotKV)
	return cms[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *store) Snapshot() int {
	cms := s.journal.Peek()
	if cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// build revision of cms by cachewrapping each cachekv store
	revision := make(mapMultiStore)
	for key, cacheKVStore := range cms {
		revisionKV := utils.MustGetAs[storetypes.CacheKVStore](cacheKVStore.CacheWrap())
		revision[key] = duocachekv.NewStoreFrom(s.sot[key], revisionKV)
	}

	// push the revision to the journal and return the size BEFORE snapshot
	return s.journal.Push(revision) - 1
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *store) RevertToSnapshot(id int) {
	// `id` is the new size of the journal we want to maintain.
	s.journal.PopToSize(id)
}

// `Finalize` commits each of the individual cachekv stores to its corresponding parent cachekv stores
// in the journal. Finally it commits the root cachekv stores.
//
// `Finalize` implements `libtypes.Controllable`.
func (s *store) Finalize() {
	// Recursively pop the journal and write each cachekv store to its parent cachekv store.
	for revision := s.journal.Peek(); s.journal.Size() > 0; revision = s.journal.Pop() {
		for key, cacheKVStore := range revision {
			cacheKVStore.Write()
			delete(revision, key)
		}
	}

	// We must handle the root separately.
	for key, cacheKVStore := range s.root {
		cacheKVStore.Write()
		delete(s.root, key)
	}
}
