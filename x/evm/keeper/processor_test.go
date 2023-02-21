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

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/common"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/crypto"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/keeper"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	"github.com/berachain/stargazer/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	dummyContract = common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
	key, _        = crypto.GenerateEthKey()
	signer        = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)

	legacyTxData = &coretypes.LegacyTx{
		Nonce:    0,
		To:       &dummyContract,
		Gas:      100000,
		GasPrice: big.NewInt(2),
		Data:     []byte("abcdef"),
	}
)

var _ = Describe("Processor", func() {
	var (
		k             *keeper.Keeper
		ak            state.AccountKeeper
		bk            state.BankKeeper
		ctx           sdk.Context
		blockGasLimit = uint64(1000000)
		tx            *coretypes.Transaction
	)

	BeforeEach(func() {
		// before chain, init genesis state
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(ak, bk, "authority")
		for _, plugin := range k.GetAllPlugins() {
			plugin.InitGenesis(ctx, types.DefaultGenesis())
		}
		tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
		addr, err := signer.Sender(tx)
		Expect(err).To(BeNil())
		k.GetStatePlugin().CreateAccount(addr)
		k.GetStatePlugin().AddBalance(addr, big.NewInt(10000000))
		k.GetStatePlugin().Finalize()

		// before every block
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(blockGasLimit))
		k.BeginBlocker(ctx)
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		It("should panic on nil, empty transaction", func() {
			Expect(func() { k.ProcessTransaction(ctx, nil) }).To(Panic())
			Expect(func() { k.ProcessTransaction(ctx, &coretypes.Transaction{}) }).To(Panic())
		})

		It("should handle legacy tx", func() {
			// process tx
			receipt, err := k.ProcessTransaction(ctx, tx)
			Expect(err).To(BeNil())
			Expect(receipt.BlockNumber.Int64()).To(Equal(ctx.BlockHeight()))
		})
	})
})
