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
	"math/big"
	"os"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/accounts/abi"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/eth/testutil/contracts/solidity/generated"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/keeper"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

var _ = Describe("Processor", func() {
	var (
		k            *keeper.Keeper
		ak           state.AccountKeeper
		bk           state.BankKeeper
		ctx          sdk.Context
		key, _       = crypto.GenerateEthKey()
		signer       = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
		legacyTxData *coretypes.LegacyTx
	)

	BeforeEach(func() {
		err := os.RemoveAll("tmp/berachain")
		Expect(err).To(BeNil())

		legacyTxData = &coretypes.LegacyTx{
			Nonce:    0,
			Gas:      10000000,
			Data:     []byte("abcdef"),
			GasPrice: big.NewInt(1),
		}

		// before chain, init genesis state
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("evm"),
			ak, bk,
			func() []vm.RegistrablePrecompile { return nil },
			"authority",
			sims.NewAppOptionsWithFlagHome("tmp/berachain"),
			nil,
		)
		for _, plugin := range k.GetAllPlugins() {
			plugin.InitGenesis(ctx, types.DefaultGenesis())
		}

		// before every block
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(100000000)).
			WithKVGasConfig(storetypes.GasConfig{}).
			WithBlockHeight(1)
		k.BeginBlocker(ctx)
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		AfterEach(func() {
			k.EndBlocker(ctx)
			err := os.RemoveAll("tmp/berachain")
			Expect(err).To(BeNil())
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

		It("should successfully deploy a valid contract and call it", func() {
			legacyTxData.Data = common.FromHex(generated.SolmateERC20Bin)
			tx := coretypes.MustSignNewTx(key, signer, legacyTxData)
			addr, err := signer.Sender(tx)
			Expect(err).To(BeNil())
			k.GetStatePlugin().CreateAccount(addr)
			k.GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
			k.GetStatePlugin().Finalize()

			// create the contract
			receipt, err := k.ProcessTransaction(ctx, tx)
			Expect(err).To(BeNil())
			Expect(receipt.BlockNumber.Int64()).To(Equal(ctx.BlockHeight()))
			Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
			Expect(len(k.GetStatePlugin().GetCode(receipt.ContractAddress))).NotTo(Equal(0))

			// call the contract non-view function
			legacyTxData.To = &receipt.ContractAddress
			var solmateABI abi.ABI
			err = solmateABI.UnmarshalJSON([]byte(generated.SolmateERC20ABI))
			Expect(err).To(BeNil())
			input, err := solmateABI.Pack("mint", common.BytesToAddress([]byte{0x88}), big.NewInt(8888888))
			Expect(err).To(BeNil())
			legacyTxData.Data = input
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			receipt, err = k.ProcessTransaction(ctx, tx)
			Expect(err).To(BeNil())
			Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
			Expect(len(receipt.Logs)).To(Equal(1))

			// call the contract view function
			legacyTxData.Data = crypto.Keccak256Hash([]byte("totalSupply()")).Bytes()[:4]
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			receipt, err = k.ProcessTransaction(ctx, tx)
			Expect(err).To(BeNil())
			Expect(receipt.Status).To(Equal(coretypes.ReceiptStatusSuccessful))
		})
	})
})
