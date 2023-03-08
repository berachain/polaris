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
	"reflect"
	"testing"

	dbm "github.com/cosmos/cosmos-db"

	sdkcachekv "cosmossdk.io/store/cachekv"
	sdkcachemulti "cosmossdk.io/store/cachemulti"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnapMulti(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "host/cosmos/store/snapmulti")
}

var _ = Describe("Snapmulti Store", func() {
	var (
		byte1          = []byte{1}
		cms            *store
		ms             storetypes.MultiStore
		accStoreParent storetypes.KVStore
		accStoreCache  storetypes.KVStore
		accStoreKey    = storetypes.NewKVStoreKey("acc")
		evmStoreParent storetypes.KVStore
		evmStoreCache  storetypes.KVStore
		evmStoreKey    = storetypes.NewKVStoreKey("evm")
	)

	BeforeEach(func() {
		stores := map[storetypes.StoreKey]storetypes.CacheWrapper{
			evmStoreKey: dbadapter.Store{DB: dbm.NewMemDB()},
			accStoreKey: dbadapter.Store{DB: dbm.NewMemDB()},
		}
		ms = sdkcachemulti.NewStore(
			dbm.NewMemDB(),
			stores, map[string]storetypes.StoreKey{},
			nil,
			nil,
		)
		accStoreParent = ms.GetKVStore(accStoreKey)
		evmStoreParent = ms.GetKVStore(evmStoreKey)
		cms = NewStoreFrom(ms)
		accStoreCache = cms.GetKVStore(accStoreKey)
		evmStoreCache = cms.GetKVStore(evmStoreKey)
	})

	It("CorrectStoreType", func() {
		// Test that the correct store type is returned
		Expect(reflect.TypeOf(cms.GetKVStore(evmStoreKey))).To(Equal(reflect.TypeOf(&sdkcachekv.Store{})))
		Expect(reflect.TypeOf(cms.GetKVStore(accStoreKey))).To(Equal(reflect.TypeOf(&sdkcachekv.Store{})))
	})

	It("TestWrite", func() {
		// Test that the cache multi store writes to the underlying stores
		evmStoreCache.Set(byte1, byte1)
		accStoreCache.Set(byte1, byte1)
		Expect(evmStoreParent.Get(byte1)).To(BeNil())
		Expect(accStoreParent.Get(byte1)).To(BeNil())
		Expect(evmStoreCache.Get(byte1)).To(Equal(byte1))
		Expect(accStoreCache.Get(byte1)).To(Equal(byte1))

		cms.Finalize()

		Expect(evmStoreParent.Get(byte1)).To(Equal(byte1))
		Expect(evmStoreParent.Get(byte1)).To(Equal(byte1))
		Expect(evmStoreCache.Get(byte1)).To(Equal(byte1))
		Expect(accStoreCache.Get(byte1)).To(Equal(byte1))
	})

	It("TestWriteCacheMultiStore", func() {
		// check that accStoreCache is not equal to accStoreParent
		accStoreCache.Set(byte1, byte1)
		Expect(accStoreCache.Has(byte1)).To(BeTrue())
		Expect(accStoreParent.Has(byte1)).To(BeFalse())

		// check that getting accStore from cms is not the same as parent
		accStoreCache2 := cms.GetKVStore(accStoreKey)
		Expect(accStoreCache2.Has(byte1)).To(BeTrue())
	})

	It("should have the correct registry key", func() {
		Expect(cms.RegistryKey()).To(Equal("snapmultistore"))
	})

	When("snapshots and reverts", func() {
		var snapshot1 int
		BeforeEach(func() {
			cms.GetKVStore(accStoreKey).Set(byte1, byte1)
			snapshot1 = cms.Snapshot()
		})

		It("should correctly revert", func() {
			cms.GetKVStore(accStoreKey).Set(byte1, []byte{2})
			Expect(cms.GetKVStore(accStoreKey).Get(byte1)).To(Equal([]byte{2}))

			cms.RevertToSnapshot(snapshot1)
			Expect(cms.GetKVStore(accStoreKey).Get(byte1)).To(Equal(byte1))
		})

		It("should handle nested snapshots", func() {
			cms.Snapshot()
			cms.GetKVStore(accStoreKey).Set(byte1, []byte{3})
			Expect(cms.GetKVStore(accStoreKey).Get(byte1)).To(Equal([]byte{3}))

			cms.RevertToSnapshot(snapshot1)
			Expect(cms.GetKVStore(accStoreKey).Get(byte1)).To(Equal(byte1))
		})

		It("should finalize properly", func() {
			cms.GetKVStore(accStoreKey).Set(byte1, []byte{2})
			Expect(cms.GetKVStore(accStoreKey).Get(byte1)).To(Equal([]byte{2}))

			cms.Finalize()
			Expect(accStoreParent.Get(byte1)).To(Equal([]byte{2}))
		})
	})
})
