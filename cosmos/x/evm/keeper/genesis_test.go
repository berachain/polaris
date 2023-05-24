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
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"

	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keeper", func() {
	var (
		k            *keeper.Keeper
		ak           state.AccountKeeper
		bk           state.BankKeeper
		sk           stakingkeeper.Keeper
		sc           ethprecompile.StatefulImpl
		ctx          sdk.Context
		genesisState *types.GenesisState
		err          error
	)

	BeforeEach(func() {
		// setup keepers for genesis
		ctx, ak, bk, sk = testutil.SetupMinimalKeepers()
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(30000000))

		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("evm"),
			ak, bk, sk,
			"authority",
			simtestutil.NewAppOptionsWithFlagHome("tmp/berachain"),
			evmmempool.NewEthTxPoolFrom(evmmempool.DefaultPriorityMempool()),
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
		)

		lib.MintCoinsToAddress(ctx, bk, types.ModuleName, testutil.Alice, "abera", big.NewInt(69000))

		genesisState = types.DefaultGenesis()
		err = k.InitGenesis(ctx, *genesisState)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("InitGenesis is called", func() {
		ReportAfterEach(func(report SpecReport) {
			coins := []sdk.Coin{bk.GetBalance(ctx, lib.AddressToAccAddress(testutil.Alice), "abera"),
				bk.GetBalance(ctx, lib.AddressToAccAddress(testutil.Bob), "abera")}
			fmt.Println(coins)
			fmt.Printf("chainID: %s\n", ctx.ChainID())
			fmt.Printf("GasLimit: %d\n", ctx.GasMeter().Limit())
		})

		When("the genesis is valid", func() {
			// BeforeEach(func() {

			// })

			It("should execute without error", func() {
				print(k)
				Expect(err).To(BeNil())
			})
		})
		// When("the GasLimit is invalid", func() {
		// 	// BeforeEach(func() {
		// 	// 	genesisState = *types.DefaultGenesis()
		// 	// })

		// 	It("should report a GasLimit mismatch error", func() {
		// 	})
		// })
		// When("the ChainID is invalid", func() {
		// 	It("should report a ChainID mismatch error", func() {
		// 	})
		// })
		// When("the coinbase is invalid", func() {
		// 	It("should report a coinbase mismatch error", func() {
		// 	})
		// })
		// When("the balance is invalid", func() {
		// 	It("should report a balance mismatch error", func() {
		// lib.MintCoinsToAddress(ctx, bk, types.ModuleName, testutil.Bob, "abera", big.NewInt(69000))
		// 	})
		// })
	})

	// DescribeTable("InitGenesis",
	// 	func(ctx sdk.Context, genesisState types.GenesisState, expectedErr error) {
	// 		err := k.InitGenesis(ctx, genesisState)
	// 		Expect(err).To(Equal(expectedErr))
	// 	},

	// 	Entry("the genesis is valid", ctx, *types.DefaultGenesis(), nil),
	// 	Entry("the GasLimit is invalid", ctx, genesisState, fmt.Errorf("gas limit mismatch: expected %d, got %d", ethGenesis.GasLimit, ctx.GasMeter().Limit())),
	// 	Entry("the ChainID is invalid", ctx, genesisState, fmt.Errorf("invalid ChainID: 0")),
	// 	Entry("the coinbase is invalid", ctx, genesisState, fmt.Errorf("invalid coinbase: ")),
	// 	Entry("the timestamp is invalid", ctx, genesisState, fmt.Errorf("invalid timestamp: 0")),
	// 	Entry("the balance is invalid", ctx, genesisState, fmt.Errorf("invalid balance: []")),
	// )
})
