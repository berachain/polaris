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

package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	_ = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "host/cosmos/store/utils")
}

// var testKey = storetypes.NewKVStoreKey("test")

var _ = Describe("TestKVStoreReaderAtHeight", func() {
	// var (
	// 	ctx sdk.Context
	// 	ms  = store.NewCommitMultiStore(dbm.NewMemDB())
	// )
})

// 	BeforeEach(func() {
// 		ms.MountStoreWithDB(testKey, storetypes.StoreTypeIAVL, dbm.NewMemDB())
// 		err := ms.LoadLatestVersion()
// 		Expect(err).ToNot(HaveOccurred())
// 		ctx = testutil.NewContextWithMultiStore(ms).WithBlockHeight(1)
// 		ms.Commit()
// 		// version == 1
// 	})

// 	It("should work as intended", func() {
// 		Expect(ctx.BlockHeight()).To(Equal(int64(1)))
// 		Expect(ctx.MultiStore().LatestVersion()).To(Equal(ctx.BlockHeight()))

// 		store := ctx.MultiStore().GetKVStore(testKey)
// 		store.Set([]byte("foo"), []byte("bar"))
// 		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()).
// 			Get([]byte("foo"))).To(Equal([]byte("bar")))

// 		// Move forward one block.
// 		ms.Commit()
// 		Expect(ms.LatestVersion()).To(Equal(int64(2)))
// 		// version == 2
// 		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
// 		// height == 2
// 		Expect(ctx.BlockHeight()).To(Equal(int64(2)))
// 		Expect(ctx.MultiStore().LatestVersion()).To(Equal(ctx.BlockHeight()))

// 		// Update the key.
// 		ctx = ctx.WithMultiStore(ms)
// 		store = ctx.MultiStore().GetKVStore(testKey)
// 		store.Set([]byte("foo"), []byte("notbar"))

// 		ms.Commit()
// 		// version == 3
// 		Expect(ms.LatestVersion()).To(Equal(int64(3)))
// 		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
// 		// height == 3

// 		// New version should be updated.
// 		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()).
// 			Get([]byte("foo"))).To(Equal([]byte("notbar")))

// 		// Old version should still be the same.
// 		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()-1).
// 			Get([]byte("foo"))).To(Equal([]byte("bar")))
// 		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()-2).
// 			Get([]byte("foo"))).To(Equal([]byte(nil)))

// 	})
// })
