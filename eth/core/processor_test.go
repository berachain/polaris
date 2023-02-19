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
	"time"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/mock"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	vmmock "github.com/berachain/stargazer/eth/core/vm/mock"
	"github.com/berachain/stargazer/eth/crypto"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/eth/testutil/contracts/solidity/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	dummyContract = common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
	key, _        = crypto.GenerateEthKey()
	signer        = types.LatestSignerForChainID(params.DefaultChainConfig.ChainID)

	legacyTxData = &types.LegacyTx{
		Nonce:    0,
		To:       &dummyContract,
		Gas:      100000,
		GasPrice: big.NewInt(2),
		Data:     []byte("abcdef"),
	}
)

var _ = Describe("StateProcessor", func() {
	var (
		// evm         *vmmock.StargazerEVMMock
		sdb *vmmock.StargazerStateDBMock
		// msg         *mock.MessageMock
		host          *mock.StargazerHostChainMock
		bp            *mock.BlockPluginMock
		gp            *mock.GasPluginMock
		cp            *mock.ConfigurationPluginMock
		pp            *mock.PrecompilePluginMock
		sp            *core.StateProcessor
		blockNumber   uint64
		blockGasLimit uint64
	)

	BeforeEach(func() {
		// evm = vmmock.NewStargazerEVM()
		sdb = vmmock.NewEmptyStateDB()
		// msg = mock.NewEmptyMessage()
		host = mock.NewMockHost()
		bp = mock.NewBlockPluginMock()
		gp = mock.NewGasPluginMock()
		cp = mock.NewConfigurationPluginMock()
		pp = &mock.PrecompilePluginMock{}
		host.GetBlockPluginFunc = func() core.BlockPlugin {
			return bp
		}
		host.GetGasPluginFunc = func() core.GasPlugin {
			return gp
		}
		host.GetConfigurationPluginFunc = func() core.ConfigurationPlugin {
			return cp
		}
		host.GetPrecompilePluginFunc = func() core.PrecompilePlugin {
			return pp
		}
		sp = core.NewStateProcessor(host, sdb, vm.Config{}, true)
		Expect(sp).ToNot(BeNil())
		blockNumber = params.DefaultChainConfig.LondonBlock.Uint64() + 1
		blockGasLimit = 1000000

		bp.GetStargazerHeaderAtHeightFunc = func(height int64) *types.StargazerHeader {
			header := types.NewEmptyStargazerHeader()
			header.GasLimit = blockGasLimit
			header.BaseFee = big.NewInt(1)
			header.Coinbase = common.BytesToAddress([]byte{2})
			header.Number = big.NewInt(int64(blockNumber))
			header.Time = uint64(3)
			header.Difficulty = new(big.Int)
			header.MixDigest = common.BytesToHash([]byte{})
			return header
		}
		pp.HasFunc = func(addr common.Address) bool {
			return false
		}

		sdb.PrepareFunc = func(rules params.Rules, sender common.Address,
			coinbase common.Address, dest *common.Address,
			precompiles []common.Address, txAccesses types.AccessList,
		) {
			// no-op
		}

		gp.SetBlockGasLimit(blockGasLimit)
		sp.Prepare(context.Background(), 0)
	})

	Context("Empty block", func() {
		It("should build a an empty block", func() {
			block, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())
			Expect(block).ToNot(BeNil())
			Expect(block.TxIndex()).To(Equal(uint(0)))
		})
	})

	Context("Block with transactions", func() {
		BeforeEach(func() {
			_, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())

			pp.ResetFunc = func(ctx context.Context) {
				// no-op
			}

			sp.Prepare(context.Background(), int64(blockNumber))
		})

		It("should error on an unsigned transaction", func() {
			receipt, err := sp.ProcessTransaction(context.Background(), types.NewTx(legacyTxData))
			Expect(err).ToNot(BeNil())
			Expect(receipt).To(BeNil())
			block, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())
			Expect(block).ToNot(BeNil())
			Expect(block.TxIndex()).To(Equal(uint(0)))
		})

		It("should not error on a signed transaction", func() {
			signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(200000)
			}
			result, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Status).To(Equal(types.ReceiptStatusSuccessful))
			Expect(result.BlockNumber).To(Equal(big.NewInt(int64(blockNumber))))
			Expect(result.TransactionIndex).To(Equal(uint(0)))
			Expect(result.TxHash.Hex()).To(Equal(signedTx.Hash().Hex()))
			Expect(result.GasUsed).ToNot(BeZero())
			block, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())
			Expect(block).ToNot(BeNil())
			Expect(block.TxIndex()).To(Equal(uint(1)))
		})

		It("should add a contract address to the receipt", func() {
			legacyTxDataCopy := *legacyTxData
			legacyTxDataCopy.To = nil
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(200000)
			}
			signedTx := types.MustSignNewTx(key, signer, &legacyTxDataCopy)
			result, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.ContractAddress).ToNot(BeNil())
			block, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())
			Expect(block).ToNot(BeNil())
			Expect(block.TxIndex()).To(Equal(uint(1)))
		})

		It("should mark a receipt with a virtual machine error as failed", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(200000)
			}
			sdb.GetCodeFunc = func(addr common.Address) []byte {
				if addr != dummyContract {
					return nil
				}
				return []byte(generated.RevertableTxMetaData.Bin)
			}
			sdb.GetCodeHashFunc = func(addr common.Address) common.Hash {
				if addr != dummyContract {
					return common.Hash{}
				}
				return crypto.Keccak256Hash([]byte(generated.RevertableTxMetaData.Bin))
			}
			sdb.ExistFunc = func(addr common.Address) bool {
				return addr == dummyContract
			}
			legacyTxData.Value = big.NewInt(0)
			signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			result, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Status).To(Equal(types.ReceiptStatusFailed))
			block, err := sp.Finalize(context.Background())
			Expect(err).To(BeNil())
			Expect(block).ToNot(BeNil())
			Expect(block.TxIndex()).To(Equal(uint(1)))
		})

		// It("should panic when block gas limit is exceeded", func() {
		// 	blockGasLimit = uint64(20000)
		// 	bp.GetStargazerHeaderAtHeightFunc = func(height int64) *types.StargazerHeader {
		// 		header := types.NewEmptyStargazerHeader()
		// 		header.GasLimit = blockGasLimit
		// 		header.BaseFee = big.NewInt(1)
		// 		header.Coinbase = common.BytesToAddress([]byte{2})
		// 		header.Number = big.NewInt(int64(blockNumber))
		// 		header.Time = uint64(3)
		// 		return header
		// 	}
		// 	gp.SetBlockGasLimit(blockGasLimit)

		// 	signedTx := types.MustSignNewTx(key, signer, legacyTxData)
		// 	Expect(func() {
		// 		sp.ProcessTransaction(context.Background(), signedTx)
		// 	}).To(Panic())
		// })
	})
})

var _ = Describe("Stargazer", func() {
	var (
		sdb           *vmmock.StargazerStateDBMock
		host          *mock.StargazerHostChainMock
		bp            *mock.BlockPluginMock
		gp            *mock.GasPluginMock
		cp            *mock.ConfigurationPluginMock
		pp            *mock.PrecompilePluginMock
		sp            *core.StateProcessor
		blockGasLimit uint64
	)

	BeforeEach(func() {
		sdb = vmmock.NewEmptyStateDB()
		host = mock.NewMockHost()
		bp = mock.NewBlockPluginMock()
		gp = mock.NewGasPluginMock()
		cp = mock.NewConfigurationPluginMock()
		pp = &mock.PrecompilePluginMock{}
		host.GetBlockPluginFunc = func() core.BlockPlugin {
			return bp
		}
		host.GetGasPluginFunc = func() core.GasPlugin {
			return gp
		}
		host.GetConfigurationPluginFunc = func() core.ConfigurationPlugin {
			return cp
		}
		host.GetPrecompilePluginFunc = func() core.PrecompilePlugin {
			return pp
		}
		sp = core.NewStateProcessor(host, sdb, vm.Config{}, true)
		Expect(sp).ToNot(BeNil())
		blockGasLimit = 1000000

		bp.GetStargazerHeaderAtHeightFunc = func(height int64) *types.StargazerHeader {
			return types.NewStargazerHeader(
				&types.Header{
					Number:     big.NewInt(height + 1),
					BaseFee:    big.NewInt(69),
					GasLimit:   blockGasLimit,
					ParentHash: crypto.Keccak256Hash([]byte{byte(height)}),
					Time:       uint64(time.Now().Unix()),
					Difficulty: big.NewInt(0),
					MixDigest:  common.Hash{},
				},
				crypto.Keccak256Hash([]byte{byte(height)}),
			)
		}
		pp.HasFunc = func(addr common.Address) bool {
			return false
		}

		gp.SetBlockGasLimit(blockGasLimit)
		sp.Prepare(context.Background(), 0)
	})

	It("should return the correct hash", func() {
		// hashFn := sp.GetHashFn()
		// Expect(hashFn(112)).To(Equal(crypto.Keccak256Hash([]byte{byte(112)})))
	})
})
