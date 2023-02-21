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

package types_test

import (
	"github.com/ethereum/go-ethereum/trie"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
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
