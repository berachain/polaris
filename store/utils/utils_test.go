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

	storeutils "github.com/berachain/stargazer/store/utils"
	"github.com/berachain/stargazer/testutil"
	store "github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tm-db"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "utils")
}

var testKey = storetypes.NewKVStoreKey("test")

var _ = Describe("TestKVStoreReaderAtHeight", func() {
	var (
		ctx sdk.Context
		ms  = store.NewCommitMultiStore(dbm.NewMemDB())
	)

	BeforeEach(func() {
		ms.MountStoreWithDB(testKey, storetypes.StoreTypeIAVL, dbm.NewMemDB())
		err := ms.LoadLatestVersion()
		Expect(err).ToNot(HaveOccurred())
		ctx = testutil.NewContextWithMultiStore(ms)
	})

	It("should work as intended", func() {
		store := ctx.KVStore(testKey)
		store.Set([]byte("foo"), []byte("bar"))
		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()).
			Get([]byte("foo"))).To(Equal([]byte("bar")))
		Expect(ms.LatestVersion()).To(Equal(ctx.BlockHeight()))

		ms.Commit()
		Expect(ms.LatestVersion()).To(Equal(ctx.BlockHeight() + 1))

		// Move forward one block.
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
		// Update the key.
		store = ctx.KVStore(testKey)
		store.Set([]byte("foo"), []byte("notbar"))
		// Old version should still be the same
		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()-1).
			Get([]byte("foo"))).To(Equal([]byte("bar")))
		// New version should be updated
		Expect(storeutils.KVStoreReaderAtBlockHeight(ctx, testKey, ctx.BlockHeight()).
			Get([]byte("foo"))).To(Equal([]byte("notbar")))
	})
})
