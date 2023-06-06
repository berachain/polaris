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

package evm_test

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/cosmos/x/evm"
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	var (
		cdc    codec.JSONCodec
		ctx    sdk.Context
		sc     ethprecompile.StatefulImpl
		ak     state.AccountKeeper
		sk     stakingkeeper.Keeper
		k      *keeper.Keeper
		ethGen *types.GenesisState
		am     evm.AppModule
		err    error
	)

	BeforeEach(func() {
		ctx, ak, _, sk = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			ak, sk,
			storetypes.NewKVStoreKey("evm"),
			"authority",
			evmmempool.NewPolarisEthereumTxPool(),
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
		)
	})

	Context("On ValidateGenesis", func() {
		BeforeEach(func() {
		})

		When("", func() {
			It("", func() {

			})
		})
	})

	Context("On InitGenesis", func() {
		BeforeEach(func() {
			ethGen = core.DefaultGenesis()
			am = evm.NewAppModule(k, ak)
		})

		JustBeforeEach(func() {
			var bz []byte
			bz, err = ethGen.Marshal()
			if err != nil {
				panic(err)
			}
			am.InitGenesis(ctx, cdc, bz)
		})

		When("the genesis is valid", func() {
			BeforeEach(func() {

			})
			It("should succeed without error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(k.GetHost().GetBlockPlugin().GetHeaderByNumber(0)).To(Equal(ethGen.ToBlock().Header()))
			})
		})
	})

	Context("On ExportGenesis", func() {
		BeforeEach(func() {

		})

		When("", func() {
			It("", func() {

			})
		})
	})
})
