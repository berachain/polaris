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

package types_test

import (
	"unsafe"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block", func() {
	var r types.Receipts
	var sr *types.StargazerReceipts
	var sh *types.StargazerHeader
	var sb *types.StargazerBlock

	BeforeEach(func() {
		r = types.Receipts{
			{
				Type: 1,
				Logs: []*types.Log{
					{Address: common.BytesToAddress([]byte{1})},
					{Address: common.BytesToAddress([]byte{2})},
				},
			},
			{
				Type: 2,
				Logs: []*types.Log{
					{Address: common.BytesToAddress([]byte{3})},
					{Address: common.BytesToAddress([]byte{4})},
				},
			},
		}
		sr = types.StargazerReceiptsFromReceipts(r)
		sh = types.NewEmptyStargazerHeader()
		sb = types.NewStargazerBlock(
			sh,
			types.Transactions{
				types.NewTx(&types.DynamicFeeTx{}),
				types.NewTx(&types.LegacyTx{}),
			},
			*sr,
		)
	})

	It("should conform to the standard create bloom", func() {
		sb.CreateBloom()
		Expect(sb.StargazerHeader.Bloom).To(Equal(types.CreateBloom(r)))
	})

	It("should work", func() {
		sb.SetGasUsed(uint64(100))
		Expect(sb.GasUsed).To(Equal(uint64(100)))

		sb.SetReceiptHash()
		Expect(sb.ReceiptHash).To(Equal(types.DeriveSha(
			//#nosec:G103
			*(*(types.Receipts))((unsafe.Pointer(&sr.Receipts))), trie.NewStackTrie(nil),
		)))

		sb.CreateBloom()
		Expect(sb.StargazerHeader.Bloom).To(Equal(
			types.CreateBloom(*(*(types.Receipts))((unsafe.Pointer(&sr.Receipts))))),
		)

		data, err := sb.MarshalBinary()
		Expect(err).To(BeNil())
		sb2 := &types.StargazerBlock{}
		err = sb2.UnmarshalBinary(data)
		Expect(err).To(BeNil())
		Expect(sb2.Bloom).To(Equal(sb.Bloom))
		Expect(sb2.Bloom).To(Equal(types.CreateBloom(r))) // conforms to standard create bloom
	})

	It("should set to empty root hash on no receipts", func() {
		b := types.NewStargazerBlock(sh, types.Transactions{}, *types.NewStargazerReceipts())
		b.SetReceiptHash()
		Expect(b.ReceiptHash).To(Equal(types.EmptyRootHash))
	})
})
