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
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/mock"
	"github.com/berachain/stargazer/lib/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ethsecp256k1 "github.com/berachain/stargazer/lib/crypto"
)

var (
	william = common.HexToAddress("0x123")
	key, _  = ethsecp256k1.GenerateEthKey()
	signer  = types.LatestSignerForChainID(params.MainnetChainConfig.ChainID)
)

var _ = Describe("StateProcessor", func() {
	var (
		// evm *vmmock.StargazerEVMMock
		// sdb  *vmmock.StargazerStateDBMock
		// msg  = new(mock.MessageMock)
		host = mock.NewMockHost()
		sp   = core.NewStateProcessor(params.MainnetChainConfig, host)
	)

	Context("Empty block", func() {
		BeforeEach(func() {
			sp.Prepare(context.Background(), 0)
			It("should return an error if the state is missing", func() {
				Expect(len(host.StargazerHeaderAtHeightCalls())).To(Equal(1))
			})

			It("should build a an empty block", func() {
				block, err := sp.Finalize(context.Background(), 0)
				Expect(err).To(BeNil())
				Expect(block).ToNot(BeNil())
				Expect(len(block.Transactions)).To(Equal(0))
			})
		})
	})

	Context("Block with transactions", func() {
		BeforeEach(func() {
			sp.Prepare(context.Background(), params.MainnetChainConfig.LondonBlock.Uint64()+1)
		})

		It("should error on an unsigned transaction", func() {
			tx := types.NewTx(&types.LegacyTx{
				Nonce:    0,
				GasPrice: big.NewInt(0),
				Gas:      100000,
				To:       &william,
				Value:    big.NewInt(0),
				Data:     []byte{},
				V:        big.NewInt(0),
				R:        big.NewInt(0),
				S:        big.NewInt(0),
			})
			receipt, err := sp.ProcessTransaction(context.Background(), tx)
			Expect(err).ToNot(BeNil())
			Expect(receipt).To(BeNil())
		})

		It("should not error on a signed transaction", func() {
			txData := &types.LegacyTx{
				Nonce:    0,
				To:       &william,
				Gas:      100000,
				GasPrice: big.NewInt(2),
				Data:     []byte("abcdef"),
			}
			signedTx := types.MustSignNewTx(key, signer, txData)
			_, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).To(BeNil())
			// Expect(receipt).To(BeNil())
		})
	})
})
