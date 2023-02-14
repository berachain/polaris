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

package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega".
)

var (
// william = common.HexToAddress("0x123")
// key, _  = ethsecp256k1.GenerateEthKey()
// signer  = types.LatestSignerForChainID(params.MainnetChainConfig.ChainID)

//	legacyTxData = &types.LegacyTx{
//		Nonce:    0,
//		To:       &william,
//		Gas:      100000,
//		GasPrice: big.NewInt(2),
//		Data:     []byte("abcdef"),
//	}
)

var _ = Describe("StateProcessor", func() {
	var (
	// evm *vmmock.StargazerEVMMock
	// sdb  *vmmock.StargazerStateDBMock
	// msg  = new(mock.MessageMock)
	// host = mock.NewMockHost()
	// // sp          = core.NewStateProcessor(params.MainnetChainConfig, host)
	// blockNumber = params.MainnetChainConfig.LondonBlock.Uint64() + 1
	)

	Context("Empty block", func() {
		BeforeEach(func() {
			// sp.Prepare(context.Background(), 0)
			// It("should return an error if the state is missing", func() {
			// 	Expect(len(host.StargazerHeaderAtHeightCalls())).To(Equal(1))
			// })

			// It("should build a an empty block", func() {
			// 	block, err := sp.Finalize(context.Background(), 0)
			// 	Expect(err).To(BeNil())
			// 	Expect(block).ToNot(BeNil())
			// 	Expect(len(block.Transactions)).To(Equal(0))
			// })
		})
	})

	Context("Block with transactions", func() {
		BeforeEach(func() {
			// sp.Prepare(context.Background(), blockNumber)
		})

		It("should error on an unsigned transaction", func() {
			// receipt, err := sp.ProcessTransaction(context.Background(), types.NewTx(legacyTxData))
			// Expect(err).ToNot(BeNil())
			// Expect(receipt).To(BeNil())
			// block, err := sp.Finalize(context.Background(), blockNumber)
			// Expect(err).To(BeNil())
			// Expect(block).ToNot(BeNil())
			// Expect(len(block.Transactions)).To(Equal(0))
		})

		It("should not error on a signed transaction", func() {
			// signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			// result, err := sp.ProcessTransaction(context.Background(), signedTx)
			// Expect(err).To(BeNil())
			// Expect(result).ToNot(BeNil())
			// Expect(result.Status).To(Equal(types.ReceiptStatusSuccessful))
			// Expect(result.BlockNumber).To(Equal(big.NewInt(int64(blockNumber))))
			// Expect(result.TransactionIndex).To(Equal(uint(0)))
			// Expect(result.TxHash.Hex()).To(Equal(signedTx.Hash().Hex()))
			// Expect(result.GasUsed).ToNot(BeZero())
			// block, err := sp.Finalize(context.Background(), blockNumber)
			// Expect(err).To(BeNil())
			// Expect(block).ToNot(BeNil())
			// Expect(len(block.Transactions)).To(Equal(1))
		})

		It("should add a contract address to the receipt", func() {
			// legacyTxDataCopy := *legacyTxData
			// legacyTxDataCopy.To = nil
			// signedTx := types.MustSignNewTx(key, signer, &legacyTxDataCopy)
			// result, err := sp.ProcessTransaction(context.Background(), signedTx)
			// Expect(err).To(BeNil())
			// Expect(result).ToNot(BeNil())
			// Expect(result.ContractAddress).ToNot(BeNil())
			// block, err := sp.Finalize(context.Background(), blockNumber)
			// Expect(err).To(BeNil())
			// Expect(block).ToNot(BeNil())
			// Expect(len(block.Transactions)).To(Equal(1))
		})

		It("should mark a receipt with a virtual machine error as failed", func() {
			// signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			// result, err := sp.ProcessTransaction(context.Background(), signedTx)
			// Expect(err).To(BeNil())
			// Expect(result).ToNot(BeNil())
			// Expect(result.Status).To(Equal(types.ReceiptStatusFailed))
			// block, err := sp.Finalize(context.Background(), blockNumber)
			// Expect(err).To(BeNil())
			// Expect(block).ToNot(BeNil())
			// Expect(len(block.Transactions)).To(Equal(1))
		})

		It("should not include consensus breaking transactions", func() {
			// signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			// result, err := sp.ProcessTransaction(context.Background(), signedTx)
			// Expect(err).To(BeNil())
			// Expect(result).ToNot(BeNil())
			// Expect(result.Status).To(Equal(types.ReceiptStatusFailed))
			// block, err := sp.Finalize(context.Background(), blockNumber)
			// Expect(err).To(BeNil())
			// Expect(block).ToNot(BeNil())
			// Expect(len(block.Transactions)).To(Equal(1))
		})
	})
})

var _ = Describe("EVM Test Suite", func() {
	// var host *mock.StargazerHostChainMock

	// hash1 := common.Hash{1}
	// hash2 := common.Hash{2}
	// hash3 := common.Hash{3}
	// hash4 := common.Hash{4}

	// currentHeader := &types.StargazerHeader{
	// 	Header: &types.Header{
	// 		Number:     big.NewInt(int64(123)),
	// 		BaseFee:    big.NewInt(69),
	// 		ParentHash: common.Hash{111},
	// 	},
	// 	CachedHash: common.Hash{1},
	// }

	Context("TestGetHashFunc", func() {
		BeforeEach(func() {
			// host = mock.NewMockHost()
		})
		// It("should return the correct hash", func() {
		// 	host.StargazerHeaderAtHeightFunc = func(ctx context.Context, height uint64) *types.StargazerHeader {
		// 		return &types.StargazerHeader{
		// 			Header: &types.Header{
		// 				Number:     big.NewInt(int64(height)),
		// 				BaseFee:    big.NewInt(69),
		// 				ParentHash: common.Hash{123},
		// 			},
		// 			CachedHash: crypto.Keccak256Hash([]byte{byte(height)}),
		// 		}
		// 	}
		// 	hash := core.GetHashFn(context.Background(), currentHeader, host)
		// 	Expect(hash(112)).To(Equal(crypto.Keccak256Hash([]byte{byte(112)})))
		// })
	})
})
