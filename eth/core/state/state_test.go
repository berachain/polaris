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

package state_test

import (
	"math/big"
	"testing"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/state"
	vmmock "github.com/berachain/stargazer/eth/core/vm/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/core/state")
}

var _ = Describe("EVM Test Suite", func() {
	var sdb *vmmock.StargazerStateDBMock
	var addr common.Address

	BeforeEach(func() {
		sdb = vmmock.NewEmptyStateDB()
	})

	Context("Test CanTransfer", func() {
		It("should return true if the account has enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			ok := state.CanTransfer(sdb, addr, big.NewInt(100))
			Expect(ok).To(BeTrue())
		})

		It("should return false if the account does not have enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			ok := state.CanTransfer(sdb, addr, big.NewInt(101))
			Expect(ok).To(BeFalse())
		})
	})

	Context("Test Transfer", func() {
		It("should state.Transfer the amount if the account has enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			sdb.SubBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(0)
				}
			}
			sdb.AddBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(100)
				}
			}
			state.Transfer(sdb, addr, addr, big.NewInt(100))
			Expect(sdb.GetBalanceFunc(addr).Cmp(big.NewInt(100))).To(Equal(0))
		})

		It("should not state.Transfer the amount if the account does not have enough balance", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(100)
			}
			sdb.SubBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(0)
				}
			}
			sdb.AddBalanceFunc = func(addr common.Address, amount *big.Int) {
				sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
					return big.NewInt(100)
				}
			}
			state.Transfer(sdb, addr, addr, big.NewInt(101))
			Expect(sdb.GetBalanceFunc(addr).Cmp(big.NewInt(100))).To(Equal(0))
		})
	})
})
