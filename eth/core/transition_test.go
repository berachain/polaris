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

	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/mock"
	"github.com/berachain/stargazer/eth/core/vm"
	vmmock "github.com/berachain/stargazer/eth/core/vm/mock"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateTransition", func() {
	var (
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

		msg.ToFunc = func() *common.Address {
			return &common.Address{1}
		}

	})

	When("Contract Creation", func() {
		BeforeEach(func() {
			msg.ToFunc = func() *common.Address {
				return nil
			}
		})
		It("should call create", func() {
			msg.GasFunc = func() uint64 {
				return 53000 // exact intrinsic gas for create after homestead
			}
			res, err := core.ApplyMessage(evm, msg)
			Expect(len(evm.CreateCalls())).To(Equal(1))
			Expect(res.UsedGas).To(Equal(uint64(53000)))
			Expect(err).To(BeNil())
		})
		When("we have less than the intrinsic gas", func() {
			msg.GasFunc = func() uint64 {
				return 53000 - 1
			}
			It("should return error", func() {
				_, err := core.ApplyMessage(evm, msg)
				Expect(err).To(MatchError(core.ErrIntrinsicGas))
			})
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
			msg.GasFunc = func() uint64 {
				return 100000
			}
		})
		Context("Gas Refund", func() {
			BeforeEach(func() {
				sdb.GetRefundFunc = func() uint64 {
					return 20000
				}
				evm.StateDBFunc = func() vm.StargazerStateDB {
					return sdb
				}
				evm.CallFunc = func(caller vm.ContractRef, addr common.Address,
					input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
					return []byte{}, 80000, nil
				}
			})
			When("we are in london", func() {
				It("should call call", func() {
					res, err := core.ApplyMessage(evm, msg)
					Expect(len(evm.CallCalls())).To(Equal(1))
					Expect(res.UsedGas).To(Equal(uint64(16000))) // refund is capped to 1/5th
					Expect(err).To(BeNil())
				})
			})

			When("we are not in london", func() {
				It("should call and cap refund properly", func() {
					evm.ChainConfigFunc = func() *params.EthChainConfig {
						return &params.EthChainConfig{
							LondonBlock:    big.NewInt(1000000000),
							HomesteadBlock: big.NewInt(0),
						}
					}
					res, err := core.ApplyMessage(evm, msg)
					Expect(len(evm.CallCalls())).To(Equal(1))
					Expect(res.UsedGas).To(Equal(uint64(10000))) // refund is capped to 1/2
					Expect(err).To(BeNil())
				})
			})
		})
		It("should check to ensure required funds are available", func() {
			msg.ValueFunc = func() *big.Int {
				return big.NewInt(1)
			}
			evm.ContextFunc = func() vm.BlockContext {
				return vm.BlockContext{
					CanTransfer: func(db vm.GethStateDB, addr common.Address, amount *big.Int) bool {
						return false
					},
				}
			}
			_, err := core.ApplyMessage(evm, msg)
			Expect(err).To(MatchError(core.ErrInsufficientFundsForTransfer))
		})
	})
})
