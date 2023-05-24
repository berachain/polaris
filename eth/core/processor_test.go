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

	bindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/mock"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	vmmock "pkg.berachain.dev/polaris/eth/core/vm/mock"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	dummyContract = common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
	key, _        = crypto.GenerateEthKey()
	signer        = types.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
	_             = key
	_             = signer
	blockGasLimit = 10000000
	dummyHeader   = &types.Header{
		Number:   big.NewInt(1),
		BaseFee:  big.NewInt(1),
		GasLimit: uint64(blockGasLimit),
	}
	legacyTxData = &types.LegacyTx{
		Nonce:    0,
		To:       &dummyContract,
		Gas:      1000000,
		GasPrice: big.NewInt(1),
		Data:     []byte("abcdef"),
	}
)

var _ = Describe("StateProcessor", func() {
	var (
		sdb *vmmock.PolarisStateDBMock
		bp  *mock.BlockPluginMock
		gp  *mock.GasPluginMock
		cp  *mock.ConfigurationPluginMock
		pp  *mock.PrecompilePluginMock
		sp  *core.StateProcessor
		evm *vm.GethEVM
	)

	BeforeEach(func() {
		sdb = vmmock.NewEmptyStateDB()
		_, bp, cp, gp, _, pp, _, _ = mock.NewMockHostAndPlugins()
		bp.GetNewBlockMetadataFunc = func(n uint64) (common.Address, uint64) {
			return common.BytesToAddress([]byte{2}), uint64(3)
		}
		pp.HasFunc = func(addr common.Address) bool {
			return false
		}
		gp.SetBlockGasLimit(uint64(blockGasLimit))
		sdb.SetTxContextFunc = func(thash common.Hash, ti int) {}
		sdb.TxIndexFunc = func() int { return 0 }
		sp = core.NewStateProcessor(cp, gp, pp, sdb, &vm.Config{})
		Expect(sp).ToNot(BeNil())
		evm = vm.NewGethEVMWithPrecompiles(
			vm.BlockContext{
				Transfer:    core.Transfer,
				CanTransfer: core.CanTransfer,
			}, vm.TxContext{}, sdb, cp.ChainConfig(), vm.Config{}, pp,
		)
		sp.Prepare(evm, dummyHeader)
	})

	Context("Empty block", func() {
		It("should build a an empty block", func() {
			block, receipts, logs, err := sp.Finalize(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(block).ToNot(BeNil())
			Expect(receipts).To(BeEmpty())
			Expect(logs).To(BeEmpty())
		})
	})

	Context("Block with transactions", func() {
		BeforeEach(func() {
			_, _, _, err := sp.Finalize(context.Background())
			Expect(err).ToNot(HaveOccurred())
			sp.Prepare(evm, dummyHeader)
		})

		It("should error on an unsigned transaction", func() {
			Expect(gp.SetTxGasLimit(1000002)).ToNot(HaveOccurred())
			receipt, err := sp.ProcessTransaction(context.Background(), types.NewTx(legacyTxData))
			Expect(err).To(HaveOccurred())
			Expect(receipt).To(BeNil())
			block, receipts, logs, err := sp.Finalize(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(block).ToNot(BeNil())
			Expect(receipts).To(BeEmpty())
			Expect(logs).To(BeEmpty())
		})

		It("should not error on a signed transaction", func() {
			signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(1000001)
			}
			sdb.FinaliseFunc = func(bool) {}
			Expect(gp.SetTxGasLimit(1000002)).ToNot(HaveOccurred())
			result, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).ToNot(BeNil())
			Expect(result.Err).ToNot(HaveOccurred())
			Expect(result.UsedGas).ToNot(BeZero())
			block, receipts, logs, err := sp.Finalize(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(block).ToNot(BeNil())
			Expect(receipts).To(HaveLen(1))
			Expect(logs).To(BeEmpty())
		})

		It("should handle", func() {
			sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
				return big.NewInt(1000001)
			}
			sdb.GetCodeFunc = func(addr common.Address) []byte {
				if addr != dummyContract {
					return nil
				}
				return common.Hex2Bytes(bindings.PrecompileConstructorMetaData.Bin)
			}
			sdb.GetCodeHashFunc = func(addr common.Address) common.Hash {
				if addr != dummyContract {
					return common.Hash{}
				}
				return crypto.Keccak256Hash(common.Hex2Bytes(bindings.PrecompileConstructorMetaData.Bin))
			}
			sdb.ExistFunc = func(addr common.Address) bool {
				return addr == dummyContract
			}
			sdb.FinaliseFunc = func(bool) {}
			legacyTxData.To = nil
			legacyTxData.Value = big.NewInt(0)
			signedTx := types.MustSignNewTx(key, signer, legacyTxData)
			Expect(gp.SetTxGasLimit(1000002)).ToNot(HaveOccurred())
			result, err := sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).ToNot(BeNil())
			Expect(result.Err).ToNot(HaveOccurred())

			// Now try calling the contract
			legacyTxData.To = &dummyContract
			signedTx = types.MustSignNewTx(key, signer, legacyTxData)
			Expect(gp.SetTxGasLimit(1000002)).ToNot(HaveOccurred())
			result, err = sp.ProcessTransaction(context.Background(), signedTx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).ToNot(BeNil())
			Expect(result.Err).ToNot(HaveOccurred())
			block, receipts, logs, err := sp.Finalize(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(block).ToNot(BeNil())
			Expect(receipts).To(HaveLen(2))
			Expect(logs).To(BeEmpty())
		})
	})
})

var _ = Describe("No precompile plugin provided", func() {
	It("should use the default plugin if none is provided", func() {
		_, bp, cp, gp, _, _, _, _ := mock.NewMockHostAndPlugins()
		gp.SetBlockGasLimit(uint64(blockGasLimit))
		bp.GetNewBlockMetadataFunc = func(n uint64) (common.Address, uint64) {
			return common.BytesToAddress([]byte{2}), uint64(3)
		}
		sp := core.NewStateProcessor(cp, gp, nil, vmmock.NewEmptyStateDB(), &vm.Config{})
		Expect(func() {
			sp.Prepare(nil, &types.Header{
				GasLimit: uint64(blockGasLimit),
			})
		}).ToNot(Panic())
	})
})
