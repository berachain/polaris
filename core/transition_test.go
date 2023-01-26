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
	"math/big"

	"github.com/berachain/stargazer/core"
	"github.com/berachain/stargazer/core/mock"
	vmmock "github.com/berachain/stargazer/core/vm/mock"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/params"
	"github.com/berachain/stargazer/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateTransition", func() {
	var (
		st  *core.StateTransition
		evm *vmmock.StargazerEVMMock
		sdb *vmmock.StargazerStateDBMock
		msg = new(mock.MessageMock)
	)

	BeforeEach(func() {
		msg = mock.NewEmptyMessage()
		evm = vmmock.NewStargazerEVM()
		sdb, _ = evm.StateDB().(*vmmock.StargazerStateDBMock)
		_ = sdb
		msg.FromFunc = func() common.Address {
			return testutil.Alice
		}

		msg.GasPriceFunc = func() *big.Int {
			return big.NewInt(123456789)
		}

		msg.GasFunc = func() uint64 {
			return 100000
		}

		msg.ToFunc = func() *common.Address {
			return &common.Address{1}
		}

		evm.ChainConfigFunc = func() *params.EthChainConfig {
			return &params.EthChainConfig{}
		}

		st = core.NewStateTransition(evm, msg)

	})
	When("Contract Creation", func() {
		BeforeEach(func() {
			msg.ToFunc = func() *common.Address {
				return nil
			}
		})
		It("should call create", func() {
			res, err := st.TransitionDB()
			Expect(len(evm.CreateCalls())).To(Equal(1))
			Expect(res.UsedGas).To(Equal(uint64(100000)))
			Expect(err).To(BeNil())
		})
	})

	When("Contract Call", func() {
		BeforeEach(func() {
			msg.ToFunc = func() *common.Address {
				return &common.Address{1}
			}

			sdb.GetCodeHashFunc = func(addr common.Address) common.Hash {
				return common.Hash{1}
			}
		})
		It("should call call", func() {
			res, err := st.TransitionDB()
			Expect(len(evm.CallCalls())).To(Equal(1))
			Expect(res.UsedGas).To(Equal(uint64(100000)))
			Expect(err).To(BeNil())
		})
	})
})
