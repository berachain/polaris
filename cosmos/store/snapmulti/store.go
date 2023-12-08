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
	"time"

	"cosmossdk.io/store/cachekv"
	storetypes "cosmossdk.io/store/types"

	polariscachekv "github.com/berachain/polaris/cosmos/store/cachekv"
	"github.com/berachain/polaris/lib/ds"
	"github.com/berachain/polaris/lib/ds/stack"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/telemetry"
)

const (
	storeRegistryKey    = `snapmultistore`
	initJournalCapacity = 16
)

// mapMultiStore represents a cached multistore, which is just a map of store keys to its
// corresponding cache kv store currently being used.
type mapMultiStore map[storetypes.StoreKey]storetypes.CacheKVStore

// store is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It journals revisions by cache-wrapping the cachekv stores on a call to `Snapshot`. In this
// store's lifecycle, any operations done before the first call to snapshot will be enforced on the
// root `mapMultiStore`.
type store struct {
	// MultiStore is the underlying multistore
	storetypes.MultiStore
	// root is the mapMultiStore used before the first snapshot is called
	root mapMultiStore
	// journal holds the snapshots of cachemultistores
	journal ds.Stack[mapMultiStore]
	// readOnly is true if the store is in read-only mode
	readOnly bool
}

// NewStoreFrom creates and returns a new `store` from a given Multistore `ms`.
func NewStoreFrom(ms storetypes.MultiStore) *store { //nolint:revive // its okay.
	return &store{
		MultiStore: ms,
		root:       make(mapMultiStore),
		journal:    stack.New[mapMultiStore](initJournalCapacity),
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (s *store) RegistryKey() string {
	return storeRegistryKey
}

// IsReadOnly returns the current read-only mode.
func (s *store) IsReadOnly() bool {
	return s.readOnly
}

// SetReadOnly sets the store to the given read-only mode.
func (s *store) SetReadOnly(readOnly bool) {
	s.readOnly = readOnly
}

// GetCommittedKVStore returns the KV Store from the given Multistore. This function follows
// the Multistore's normal `GetKVStore` code path.
func (s *store) GetCommittedKVStore(key storetypes.StoreKey) storetypes.KVStore {
	return s.MultiStore.GetKVStore(key)
}

// GetKVStore shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	var cms mapMultiStore
	if cms = s.journal.Peek(); cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// if the map multistore does not have the given storekey, get from the underlying multistore
	if cms[key] == nil {
		cms[key] = cachekv.NewStore(s.GetCommittedKVStore(key))
	}

	// if the store is in read-only mode, return a read-only store
	if s.readOnly {
		return polariscachekv.NewReadOnlyStoreFor(cms[key])
	}

	return cms[key]
}

// Snapshot implements `libtypes.Snapshottable`.
func (s *store) Snapshot() int {
	defer telemetry.MeasureSince(time.Now(), MetricKeySnapshot)
	defer telemetry.SetGauge(float32(s.journal.Size()), MetricKeySnapshotSize)

	var cms mapMultiStore
	if cms = s.journal.Peek(); cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// build revision of cms by cachewrapping each cachekv store
	revision := make(mapMultiStore)
	for key, cacheKVStore := range cms {
		revision[key] = utils.MustGetAs[storetypes.CacheKVStore](cacheKVStore.CacheWrap())
	}

	// push the revision to the journal and return the size BEFORE snapshot
	return s.journal.Push(revision) - 1
}

// Revert implements `libtypes.Snapshottable`.
func (s *store) RevertToSnapshot(id int) {
	// id is the new size of the journal we want to maintain.
	defer telemetry.MeasureSince(time.Now(), MetricKeyRevertToSnapshot)
	defer telemetry.SetGauge(float32(s.journal.Size()-id), MetricKeyRevertToSnapshotSize)
	s.journal.PopToSize(id)
}

// Finalize commits each of the individual cachekv stores to its corresponding parent cachekv
// stores in the journal. Finally it commits the root cachekv stores. Skip committing writes
// to the underlying multistore if in read-only mode.
//
// Finalize implements `libtypes.Controllable`.
func (s *store) Finalize() {
	defer telemetry.MeasureSince(time.Now(), MetricKeyFinalize)
	defer telemetry.SetGauge(float32(s.journal.Size()), MetricKeyFinalizeSize)

	// Recursively pop the journal and write each cachekv store to its parent cachekv store.
	for revision := s.journal.Pop(); revision != nil; revision = s.journal.Pop() {
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
