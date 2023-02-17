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

package core_test

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/mock"
	"github.com/berachain/stargazer/eth/core/vm"
	vmmock "github.com/berachain/stargazer/eth/core/vm/mock"
	"github.com/berachain/stargazer/eth/params"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateTransition", func() {
	var (
		evm *vmmock.StargazerEVMMock
		sdb *vmmock.StargazerStateDBMock
		msg mock.MessageMock
		gp  *mock.GasPluginMock
	)

	BeforeEach(func() {
		msg = *mock.NewEmptyMessage()
		evm = vmmock.NewStargazerEVM()
		sdb, _ = evm.StateDB().(*vmmock.StargazerStateDBMock)
		msg.FromFunc = func() common.Address {
			return common.Address{1}
		}

		msg.GasPriceFunc = func() *big.Int {
			return big.NewInt(123456789)
		}

		msg.ToFunc = func() *common.Address {
			return &common.Address{1}
		}

		gp = mock.NewGasPluginMock()
	})

	When("Contract Creation", func() {
		BeforeEach(func() {
			msg.ToFunc = func() *common.Address {
				return nil
			}
			gp.SetBlockGasLimit(1000000)
			gp.Prepare(context.Background())
		})
		It("should call create", func() {
			gp.Reset(context.Background())

			msg.GasFunc = func() uint64 {
				return 53000 // exact intrinsic gas for create after homestead
			}
			res, err := core.ApplyMessage(evm, gp, &msg, true)
			Expect(res.UsedGas).To(Equal(uint64(53000)))
			Expect(len(evm.CreateCalls())).To(Equal(1))
			Expect(err).To(BeNil())
		})
		When("we have less than the intrinsic gas", func() {
			msg.GasFunc = func() uint64 {
				return 53000 - 1
			}
			It("should return error", func() {
				gp.Reset(context.Background())
				_, err := core.ApplyMessage(evm, gp, &msg, true)
				Expect(err).To(MatchError(core.ErrIntrinsicGas))
			})
		})

		It("should call create with commit", func() {
			gp.Reset(context.Background())
			msg.GasFunc = func() uint64 {
				return 53000
			}
			_, err := core.ApplyMessage(evm, gp, &msg, true)
			Expect(err).To(BeNil())
		})

		It("should handle transition error", func() {
			gp.Reset(context.Background())
			msg.GasFunc = func() uint64 {
				return 0
			}
			_, err := core.ApplyMessage(evm, gp, &msg, true)
			Expect(err).To(Not(BeNil()))
		})

		When("We call with a tracer", func() {
			var tracer *vmmock.EVMLoggerMock
			BeforeEach(func() {
				tracer = vmmock.NewEVMLoggerMock()
				evm.ConfigFunc = func() vm.Config {
					return vm.Config{
						Debug:  true,
						Tracer: tracer,
					}
				}
				gp.Reset(context.Background())
			})

			It("should call create with tracer", func() {
				msg.GasFunc = func() uint64 {
					return 53000 // exact intrinsic gas for create after homestead
				}
				_, err := core.ApplyMessage(evm, gp, &msg, false)
				Expect(len(tracer.CaptureTxStartCalls())).To(Equal(1))
				Expect(len(tracer.CaptureTxEndCalls())).To(Equal(1))
				Expect(err).To(BeNil())
			})
			It("should call create with tracer and commit", func() {
				msg.GasFunc = func() uint64 {
					return 53000 // exact intrinsic gas for create after homestead
				}
				sdb = vmmock.NewEmptyStateDB()
				evm.StateDBFunc = func() vm.StargazerStateDB {
					return sdb
				}
				_, err := core.ApplyMessage(evm, gp, &msg, true)
				Expect(err).To(BeNil())
				Expect(len(tracer.CaptureTxStartCalls())).To(Equal(1))
				Expect(len(tracer.CaptureTxEndCalls())).To(Equal(1))
				Expect(len(sdb.FinalizeCalls())).To(Equal(1))
			})
			It("should handle abort error", func() {
				msg.GasFunc = func() uint64 {
					return 0
				}
				_, err := core.ApplyMessage(evm, gp, &msg, false)
				Expect(len(tracer.CaptureTxStartCalls())).To(Equal(1))
				Expect(len(tracer.CaptureTxEndCalls())).To(Equal(1))
				Expect(err).To(Not(BeNil()))
			})
			It("should handle abort error with commit", func() {
				msg.GasFunc = func() uint64 {
					return 0
				}
				_, err := core.ApplyMessage(evm, gp, &msg, true)
				Expect(len(tracer.CaptureTxStartCalls())).To(Equal(1))
				Expect(len(tracer.CaptureTxEndCalls())).To(Equal(1))
				Expect(err).To(Not(BeNil()))
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
			gp.SetBlockGasLimit(1000000)
			gp.Prepare(context.Background())
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
				gp.Reset(context.Background())
			})

			When("we are in london", func() {
				It("should call call", func() {
					res, err := core.ApplyMessage(evm, gp, &msg, true)
					Expect(len(evm.CallCalls())).To(Equal(1))
					Expect(res.UsedGas).To(Equal(uint64(16000))) // refund is capped to 1/5th
					Expect(err).To(BeNil())
				})
			})

			When("we are not in london", func() {
				It("should call and cap refund properly", func() {
					evm.ChainConfigFunc = func() *params.ChainConfig {
						return &params.ChainConfig{
							LondonBlock:    big.NewInt(1000000000),
							HomesteadBlock: big.NewInt(0),
						}
					}
					res, err := core.ApplyMessage(evm, gp, &msg, true)
					Expect(len(evm.CallCalls())).To(Equal(1))
					Expect(res.UsedGas).To(Equal(uint64(10000))) // refund is capped to 1/2
					Expect(err).To(BeNil())
				})
			})
		})

		It("should check to ensure required funds are available", func() {
			gp.Reset(context.Background())
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
			_, err := core.ApplyMessage(evm, gp, &msg, true)
			Expect(err).To(MatchError(core.ErrInsufficientFundsForTransfer))
		})

		It("should error on block gas limit", func() {
			gp.Reset(context.Background())

			msg.GasFunc = func() uint64 {
				return 1000000 / 2
			}
			evm.CallFunc = func(
				caller vm.ContractRef,
				addr common.Address, input []byte,
				gas uint64,
				value *big.Int,
			) ([]byte, uint64, error) {
				return nil, 0, nil
			}
			res, err := core.ApplyMessage(evm, gp, &msg, false)
			Expect(err).To(BeNil())
			Expect(res.Err).To(BeNil())

			gp.Reset(context.Background())
			msg.GasFunc = func() uint64 {
				return 1000000/2 + 1
			}
			evm.CallFunc = func(
				caller vm.ContractRef,
				addr common.Address, input []byte,
				gas uint64,
				value *big.Int,
			) ([]byte, uint64, error) {
				return nil, 0, nil
			}
			res, err = core.ApplyMessage(evm, gp, &msg, false)
			Expect(err).To(BeNil())
			Expect(res.Err.Error()).To(Equal("out of gas"))
			Expect(gp.CumulativeGasUsed()).To(Equal(uint64(1000000)))
		})

		When("the message has data", func() {
			It("should cost more gas", func() {
				gp.Reset(context.Background())

				msg.GasFunc = func() uint64 {
					return 6969699669
				}

				msg.DataFunc = func() []byte {
					return []byte{1, 2, 3}
				}

				// Call the intrinsic gas function with data
				st := core.NewStateTransition(evm, gp, &msg)
				Expect(gp.SetTxGasLimit(10000000)).To(BeNil())
				Expect(st.ConsumeEthIntrinsicGas(true, true, true, false)).To(BeNil())
				consumedWithData := gp.CumulativeGasUsed()

				// Reset the gas meter.
				gp.Prepare(context.Background())

				// Call the intrinsic gas function with no data
				msg.DataFunc = func() []byte {
					return []byte{}
				}
				Expect(st.ConsumeEthIntrinsicGas(true, true, true, false)).To(BeNil())

				// We expect that the call with Data will consume more gas.
				Expect(consumedWithData).To(BeNumerically(">", gp.CumulativeGasUsed()))
			})

			It("should cost more gas, shanghai fork", func() {
				gp.Reset(context.Background())

				msg.GasFunc = func() uint64 {
					return 6969699669
				}

				msg.DataFunc = func() []byte {
					return []byte{1, 2, 3}
				}

				// Call the intrinsic gas function with data
				st := core.NewStateTransition(evm, gp, &msg)
				Expect(gp.SetTxGasLimit(10000000)).To(BeNil())
				Expect(st.ConsumeEthIntrinsicGas(true, true, true, true)).To(BeNil())
				consumedWithData := gp.CumulativeGasUsed()

				// Reset the gas meter.
				gp.Prepare(context.Background())

				// Call the intrinsic gas function with no data
				msg.DataFunc = func() []byte {
					return []byte{}
				}
				Expect(st.ConsumeEthIntrinsicGas(true, true, true, true)).To(BeNil())

				// We expect that the call with Data will consume more gas.
				Expect(consumedWithData).To(BeNumerically(">", gp.CumulativeGasUsed()))
			})
		})
	})
})
