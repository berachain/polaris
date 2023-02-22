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

package keeper_test

import (
	"fmt"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethapi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/eth/testutil/contracts/solidity"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/keeper"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

var _ = Describe("Processor", func() {
	var (
		k      *keeper.Keeper
		ak     state.AccountKeeper
		bk     state.BankKeeper
		ctx    sdk.Context
		key, _ = crypto.GenerateEthKey()
		signer = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
		// legacyTxData *coretypes.LegacyTx
		gas    = hexutil.Uint64(10000000)
		txArgs *ethapi.TransactionArgs
	)

	BeforeEach(func() {
		// legacyTxData = &coretypes.LegacyTx{
		// 	Nonce:    0,
		// 	Gas:      10000000,
		// 	Data:     []byte("abcdef"),
		// 	GasPrice: big.NewInt(1),
		// }

		// before chain, init genesis state
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(ak, bk, "authority")
		for _, plugin := range k.GetAllPlugins() {
			plugin.InitGenesis(ctx, types.DefaultGenesis())
		}

		// before every block
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(100000000)).
			WithKVGasConfig(storetypes.GasConfig{})
		k.BeginBlocker(ctx)
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// setup
			txArgs = &ethapi.TransactionArgs{
				Gas:      &gas,
				GasPrice: (*hexutil.Big)(big.NewInt(1)),
				Value:    (*hexutil.Big)(big.NewInt(1)),
				Nonce:    (*hexutil.Uint64)(new(uint64)),
			}

			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		It("should panic on nil, empty transaction", func() {
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, nil)
				Expect(err).ToNot(BeNil())
			}).To(Panic())
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, &coretypes.Transaction{})
				Expect(err).ToNot(BeNil())
			}).To(Panic())
		})

		// It("should call a dummy contract", func() {
		// 	// setup state for legacy tx
		// 	dummyContract := common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
		// 	txArgs.To = &dummyContract
		// 	signedTx, err := coretypes.SignTx(txArgs.ToTransaction(), signer, key)
		// 	Expect(err).To(BeNil())
		// 	addr, err := signer.Sender(signedTx)
		// 	Expect(err).To(BeNil())
		// 	k.GetStatePlugin().CreateAccount(addr)
		// 	k.GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
		// 	k.GetStatePlugin().Finalize()

		// 	// process tx
		// 	receipt, err := k.ProcessTransaction(ctx, signedTx)
		// 	Expect(err).To(BeNil())
		// 	Expect(receipt.BlockNumber.Int64()).To(Equal(ctx.BlockHeight()))
		// 	Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
		// })

		// It("should create a bad contract and call it", func() {
		// 	// setup state for contract creation
		// 	legacyTxData.Value = big.NewInt(10)
		// 	legacyTxData.To = nil
		// 	legacyTxData.Data = common.FromHex(generated.RevertableTxMetaData.Bin)
		// 	tx := coretypes.MustSignNewTx(key, signer, legacyTxData)

		// 	addr, err := signer.Sender(tx)
		// 	Expect(err).To(BeNil())
		// 	k.GetStatePlugin().CreateAccount(addr)
		// 	k.GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
		// 	k.GetStatePlugin().Finalize()

		// 	// process tx
		// 	receipt, err := k.ProcessTransaction(ctx, tx)
		// 	Expect(err).To(BeNil())
		// 	Expect(receipt.BlockNumber.Int64()).To(Equal(ctx.BlockHeight()))
		// 	fmt.Println("receipt", receipt)
		// 	Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
		// 	Expect(receipt.ContractAddress).ToNot(BeNil())
		// 	Expect(k.GetStatePlugin().GetCode(receipt.ContractAddress)).To(Equal(legacyTxData.Data))
		// 	Expect(k.GetStatePlugin().Exist(receipt.ContractAddress)).To(BeTrue())
		// 	Expect(k.GetStatePlugin().GetCodeHash(receipt.ContractAddress)).To(Equal(crypto.Keccak256Hash(legacyTxData.Data)))

		// 	contractAddr := common.BytesToAddress(receipt.ContractAddress.Bytes())
		// 	legacyTxData := &coretypes.DynamicFeeTx{
		// 		Nonce:     1,
		// 		Gas:       10000000,
		// 		To:        &contractAddr,
		// 		GasTipCap: big.NewInt(1),
		// 		GasFeeCap: big.NewInt(1),
		// 		Data:      common.Hex2Bytes("0x34234"),
		// 	}
		// 	tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
		// 	receipt, err = k.ProcessTransaction(ctx, tx)
		// 	Expect(err).To(BeNil())
		// 	fmt.Println(*receipt)
		// 	Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusFailed))
		// })

		It("should successfully deploy a valid contract and call it", func() {
			// setup state for contract creation
			data := new(hexutil.Bytes)
			bz, err := solidity.ERC20Contract.Bin.MarshalJSON()
			Expect(err).To(BeNil())
			data.UnmarshalJSON(bz)
			txArgs.To = nil
			txArgs.Data = data
			tx, err := coretypes.SignTx(txArgs.ToTransaction(), signer, key)
			Expect(err).To(BeNil())
			addr, err := signer.Sender(tx)
			Expect(err).To(BeNil())
			k.GetStatePlugin().CreateAccount(addr)
			k.GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
			k.GetStatePlugin().Finalize()

			// process tx
			receipt, err := k.ProcessTransaction(ctx, tx)
			Expect(err).To(BeNil())
			Expect(receipt.BlockNumber.Int64()).To(Equal(ctx.BlockHeight()))
			Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
			fmt.Println("receipt", receipt)
			Expect(k.GetStatePlugin().GetCode(receipt.ContractAddress)).To(Equal(*txArgs.Data))

			// setup state for contract call

		})
	})
})
