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

package cachekv_test

import (
	"testing"

	cdb "github.com/cosmos/cosmos-db"

	"cosmossdk.io/store/dbadapter"

	"github.com/berachain/polaris/cosmos/store/cachekv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCacheKV(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/store/cachekv")
}

var _ = Describe("ReadOnly Store", func() {
	var readOnlyStore *cachekv.ReadOnlyStore

	BeforeEach(func() {
		kvStore := dbadapter.Store{DB: cdb.NewMemDB()}
		kvStore.Set([]byte("key"), []byte("value"))
		readOnlyStore = cachekv.NewReadOnlyStoreFor(kvStore)
	})

	It("should panic only on writes", func() {
		Expect(func() {
			Expect(readOnlyStore.Get([]byte("key"))).To(Equal([]byte("value")))
			Expect(readOnlyStore.Has([]byte("KEY"))).To(BeFalse())
		}).NotTo(Panic())

		Expect(func() {
			readOnlyStore.Set([]byte("key"), []byte("new value"))
		}).To(Panic())

		Expect(func() {
			readOnlyStore.Set([]byte("new key"), []byte("value"))
		}).To(Panic())

		Expect(func() {
			readOnlyStore.Delete([]byte("key"))
		}).To(Panic())
	})
})
