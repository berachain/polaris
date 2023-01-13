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

package cachekv_test

import (
	"testing"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/journal"
	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	dbm "github.com/tendermint/tm-db"
)

func TestJournalManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "core/state/store/cachemulti")
}

var _ = Describe("CacheMulti", func() {
	var (
		byte0           = []byte{0}
		byte1           = []byte{1}
		nonZeroCodeHash = common.BytesToHash([]byte{0x05})
		zeroCodeHash    = common.Hash{}
		parent          storetypes.KVStore
		evmStore        *cachekv.EvmStore
	)

	BeforeEach(func() {
		parent = sdkcachekv.NewStore(dbadapter.Store{DB: dbm.NewMemDB()})
		evmStore = cachekv.NewEvmStore(parent, journal.NewManager())
	})

	It("TestWarmSlotVia0", func() {
		evmStore.Set(byte1, zeroCodeHash.Bytes())
		evmStore.Write()
		Expect(parent.Get(byte1)).To(Equal(zeroCodeHash.Bytes()))
	})
	It("TestWriteZeroValParentNotNil", func() {
		evmStore.Set(byte0, zeroCodeHash.Bytes())
		evmStore.Write()
		Expect(parent.Get(byte0)).To(Equal(zeroCodeHash.Bytes()))
	})
	It("TestWriteNonZeroValParentNil", func() {
		evmStore.Set(byte0, nonZeroCodeHash.Bytes())
		Expect(parent.Get(byte0)).To(BeNil())
		evmStore.Write()
		Expect(parent.Get(byte0)).To(Equal(nonZeroCodeHash.Bytes()))
	})
	It("TestWriteNonZeroValParentNotNil", func() {
		evmStore.Set(byte0, nonZeroCodeHash.Bytes())
		Expect(parent.Get(byte0)).To(BeNil())
		evmStore.Write()
		Expect(parent.Get(byte0)).To(Equal(nonZeroCodeHash.Bytes()))
	})
	It("TestWriteAfterDelete", func() {
		evmStore.Set(byte1, zeroCodeHash.Bytes())
		Expect(evmStore.Get(byte1)).To(Equal(zeroCodeHash.Bytes()))
		evmStore.Delete(byte1)
		Expect(evmStore.Get(byte1)).To(BeNil())
		evmStore.Write()
		Expect(parent.Get(byte1)).To(BeNil())
	})
})
