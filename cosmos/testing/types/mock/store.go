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

package mock

import (
	cdb "github.com/cosmos/cosmos-db"

	"cosmossdk.io/store/dbadapter"
	"cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/cosmos/testing/types/mock/interfaces/mock"
	"pkg.berachain.dev/polaris/lib/utils"
)

// MultiStore is a simple multistore used for testing.
type MultiStore struct {
	kvstore map[string]types.KVStore
	*mock.MultiStoreMock
}

// MultiStore implements precompile.MultiStore.
func (m MultiStore) SetReadOnly(bool) {}

// MultiStore implements precompile.MultiStore.
func (m MultiStore) IsReadOnly() bool { return false }

// CachedMultiStore is a simple chached multistore for testing.
type CachedMultiStore struct {
	kvstore map[string]types.KVStore
	*mock.CacheMultiStoreMock
}

// NewMultiStore returns a new Multistore instance used for testing.
func NewMultiStore() types.MultiStore {
	ms := MultiStore{
		kvstore:        map[string]types.KVStore{},
		MultiStoreMock: &mock.MultiStoreMock{},
	}
	ms.GetKVStoreFunc = func(storeKey types.StoreKey) types.KVStore {
		if store, ok := ms.kvstore[storeKey.String()]; ok {
			return store
		}
		store := newTestKVStore(cdb.NewMemDB())
		ms.kvstore[storeKey.String()] = store
		return store
	}

	ms.CacheMultiStoreFunc = func() types.CacheMultiStore {
		return NewCachedMultiStore(ms)
	}

	return ms
}

// NewCachedMultiStore returns a new CacheMultiStore instance for testing.
func NewCachedMultiStore(ms MultiStore) types.CacheMultiStore {
	kvstore := map[string]types.KVStore{}

	for key, store := range ms.kvstore {
		kvstore[key] = utils.MustGetAs[types.KVStore](store.CacheWrap())
	}

	cached := CachedMultiStore{
		kvstore:             kvstore,
		CacheMultiStoreMock: &mock.CacheMultiStoreMock{},
	}
	cached.GetKVStoreFunc = func(storeKey types.StoreKey) types.KVStore {
		if store, ok := cached.kvstore[storeKey.String()]; ok {
			return store
		}
		store := newTestKVStore(cdb.NewMemDB())
		cached.kvstore[storeKey.String()] = store
		return store
	}

	cached.WriteFunc = func() {
		for _, store := range cached.kvstore {
			utils.MustGetAs[types.CacheKVStore](store).Write()
		}
	}
	return cached
}

// TestKVStore is a kv store for testing.
type TestKVStore struct {
	dbadapter.Store
}

// newTestKVStore returns a new kv store instance for testing.
func newTestKVStore(db cdb.DB) dbadapter.Store {
	return dbadapter.Store{DB: db}
}
