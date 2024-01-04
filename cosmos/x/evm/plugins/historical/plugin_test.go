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

package historical

import (
	"math/big"
	"testing"

	"cosmossdk.io/log"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/eth/core/mock"
	"github.com/berachain/polaris/eth/params"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHistoricalPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/historical")
}

var _ = Describe("Historical Data", func() {
	var (
		p   *plugin
		ctx sdk.Context
	)

	BeforeEach(func() {
		ctx = testutil.NewContext(log.NewTestLogger(GinkgoT())).WithBlockHeight(0)
		bp := mock.NewBlockPluginMock()

		genesis := core.DefaultGenesis
		genesis.Config = params.DefaultChainConfig
		p = utils.MustGetAs[*plugin](NewPlugin(params.DefaultChainConfig, bp, nil, testutil.EvmKey))
		Expect(p.InitGenesis(ctx, genesis)).To(Succeed())
	})

	When("Genesis block", func() {
		It("should return the header without error", func() {
			block, err := p.GetBlockByNumber(0)
			Expect(err).ToNot(HaveOccurred())
			header := block.Header()
			Expect(header).ToNot(BeNil())
			blockByHash, err := p.GetBlockByHash(block.Hash())
			Expect(err).ToNot(HaveOccurred())
			Expect(blockByHash).ToNot(BeNil())
			Expect(blockByHash.Hash()).To(Equal(block.Hash()))
		})
	})

	When("Other blocks", func() {
		It("should correctly store and return blocks", func() {
			ctx = ctx.WithBlockHeight(1)
			header := &ethtypes.Header{
				Number:   big.NewInt(1),
				GasLimit: 1000,
			}
			tx := ethtypes.NewTransaction(
				0, common.Address{0x1}, big.NewInt(1), 1000, big.NewInt(1), []byte{0x12},
			)
			txHash := tx.Hash()
			receipts := ethtypes.Receipts{
				{
					Type:              2,
					Status:            1,
					CumulativeGasUsed: 500,
					TxHash:            txHash,
					ContractAddress:   common.Address{0x1},
					GasUsed:           500,
					BlockNumber:       big.NewInt(1),
				},
			}
			txs := ethtypes.Transactions{tx}
			block := ethtypes.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil))
			blockHash := block.Hash()
			receipts[0].BlockHash = blockHash

			Expect(p.StoreBlock(block)).To(Succeed())
			Expect(p.StoreReceipts(blockHash, receipts)).To(Succeed())
			Expect(p.StoreTransactions(1, blockHash, txs)).To(Succeed())

			blockByNum, err := p.GetBlockByNumber(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(blockByNum.Hash()).To(Equal(blockHash))

			blockByHash, err := p.GetBlockByHash(blockHash)
			Expect(err).ToNot(HaveOccurred())
			Expect(blockByHash.Hash()).To(Equal(blockHash))

			receiptsByHash, err := p.GetReceiptsByHash(blockHash)
			Expect(err).ToNot(HaveOccurred())
			Expect(receiptsByHash[0].TxHash).To(Equal(receipts[0].TxHash))
			Expect(receiptsByHash[0].BlockHash).To(Equal(blockHash))

			tleByHash, err := p.GetTransactionByHash(txHash)
			Expect(err).ToNot(HaveOccurred())
			Expect(tleByHash.TxIndex).To(Equal(uint64(0)))
			Expect(tleByHash.BlockHash).To(Equal(blockHash))
			Expect(tleByHash.BlockNum).To(Equal(uint64(1)))
			Expect(tleByHash.Tx.Hash()).To(Equal(txHash))
		})
	})

})
