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
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/ethereum/go-ethereum/trie"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block", func() {
	var r types.Receipts
	var txs types.Transactions
	var sh *types.StargazerHeader
	var sb *types.StargazerBlock

	BeforeEach(func() {
		txs = types.Transactions{
			types.NewTx(&types.DynamicFeeTx{}),
			types.NewTx(&types.LegacyTx{}),
		}
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
		sh = types.NewEmptyStargazerHeader()
		sb = types.NewStargazerBlock(sh)
	})

	It("should be marshallable", func() {
		sb.Bloom = types.CreateBloom(r)
		data, err := sb.MarshalBinary()
		Expect(err).To(BeNil())
		sb2 := &types.StargazerBlock{}
		err = sb2.UnmarshalBinary(data)
		Expect(err).To(BeNil())
		Expect(sb2.Bloom).To(Equal(sb.Bloom))
		Expect(sb2.Bloom).To(Equal(types.CreateBloom(r)))
	})

	When("building a block", func() {
		BeforeEach(func() {
			Expect(sb.TxIndex()).To(Equal(uint(0)))
			Expect(sb.LogIndex()).To(Equal(uint(0)))

			sb.AppendTx(txs[0], r[0])
			Expect(sb.TxIndex()).To(Equal(uint(1)))
			Expect(sb.LogIndex()).To(Equal(uint(2)))

			sb.AppendTx(txs[1], r[1])
			Expect(sb.TxIndex()).To(Equal(uint(2)))
			Expect(sb.LogIndex()).To(Equal(uint(4)))
		})

		It("should convert receipts to storage receipts", func() {
			sr := sb.GetReceiptsForStorage()
			Expect(len(sr)).To(Equal(len(r)))
			Expect(sr[0].Logs).To(Equal(r[0].Logs))
			Expect(sr[1].Logs).To(Equal(r[1].Logs))
		})

		It("should finalize", func() {
			sb.Finalize(uint64(100))
			Expect(sb.GasUsed).To(Equal(uint64(100)))
			Expect(sb.TxHash).To(Equal(types.DeriveSha(txs, trie.NewStackTrie(nil))))
			Expect(sb.ReceiptHash).To(Equal(types.DeriveSha(r, trie.NewStackTrie(nil))))
			Expect(sb.Bloom).To(Equal(types.CreateBloom(r)))
		})

		It("should finalize empty txs", func() {
			sb2 := types.NewStargazerBlock(sh)
			sb2.Finalize(uint64(0))
			Expect(sb2.GasUsed).To(Equal(uint64(0)))
			Expect(sb2.TxHash).To(Equal(types.EmptyRootHash))
			Expect(sb2.ReceiptHash).To(Equal(types.EmptyRootHash))
		})
	})
})
