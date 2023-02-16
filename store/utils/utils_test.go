// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	_ = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "store/utils")
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
