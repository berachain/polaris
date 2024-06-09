// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
