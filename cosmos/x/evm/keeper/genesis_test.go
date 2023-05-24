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
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	enclib "pkg.berachain.dev/polaris/lib/encoding"

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
		ethGenesis   *core.Genesis
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

		lib.MintCoinsToAddress(ctx, bk, types.ModuleName, testutil.Alice, "abera", big.NewInt(69000)) //nolint:errcheck,lll // test mint must succeed

		genesisState = types.DefaultGenesis()
		ethGenesis = enclib.MustUnmarshalJSON[core.Genesis]([]byte(genesisState.Params.EthGenesis))
	})

	Context("InitGenesis is called", func() {
		JustBeforeEach(func() {
			err = k.InitGenesis(ctx, *genesisState)
		})

		When("the genesis is valid", func() {
			It("should execute without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		When("the coinbase is invalid", func() {
			BeforeEach(func() {
				ethGenesis.Coinbase = testutil.Bob
				genesisState.Params.EthGenesis = string(enclib.MustMarshalJSON(*ethGenesis))
			})
			It("should report a coinbase mismatch error", func() {
				Expect(err).To(Equal(fmt.Errorf("coinbase of the genesis block must be the null address, not: %s", testutil.Bob)))
			})
		})
		When("the balance is invalid", func() {
			BeforeEach(func() {
				ethGenesis.Alloc[testutil.Bob] = core.GenesisAccount{
					Balance: big.NewInt(100),
				}
				genesisState.Params.EthGenesis = string(enclib.MustMarshalJSON(*ethGenesis))
			})
			Context("the account does not exist", func() {
				It("should report a balance mismatch error", func() {
					Expect(err).To(Equal(fmt.Errorf("account %s balance mismatch: expected 0, got %v", testutil.Bob, 100)))
				})
			})
			Context("the account exists but the balance is mismatched", func() {
				BeforeEach(func() {
					lib.MintCoinsToAddress(ctx, bk, types.ModuleName, testutil.Bob, "abera", big.NewInt(50)) //nolint:errcheck,lll // test mint must succeed
				})
				It("should report a balance mismatch error", func() {
					Expect(err).To(Equal(fmt.Errorf("account %s balance mismatch: expected %v, got %v", testutil.Bob, 50, 100)))
				})
			})
		})
	})
})
