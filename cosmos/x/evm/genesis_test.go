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
	"bytes"
	"encoding/json"
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"pkg.berachain.dev/polaris/cosmos/precompile/staking"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm"
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"

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
		ethGen *core.Genesis
		am     evm.AppModule
		err    error
	)

	BeforeEach(func() {
		ethGen = core.DefaultGenesis
		ctx, ak, _, sk = testutil.SetupMinimalKeepers()
		sc = staking.NewPrecompileContract(&sk)
		k = keeper.NewKeeper(
			ak, sk,
			storetypes.NewKVStoreKey("evm"),
			"authority",
			evmmempool.NewWrappedGethTxPool(),
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
		)
		k.Setup(storetypes.NewKVStoreKey("offchain-evm"), nil, "", GinkgoT().TempDir(), log.NewNopLogger())

		am = evm.NewAppModule(k, ak)
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
		})
		JustBeforeEach(func() {
			var bz []byte
			bz, err = json.Marshal(ethGen)
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
			})
			It("should contain the same genesis header values", func() {
				bp := k.GetHost().GetBlockPlugin()
				expectedHeader := ethGen.ToBlock().Header()
				Expect(bp.GetHeaderByNumber(0)).To(Equal(expectedHeader))
			})
			It("should contain the correct chain config", func() {
				actualConfig := k.GetHost().GetConfigurationPlugin().ChainConfig()
				expectedConfig := ethGen.Config
				Expect(actualConfig).To(Equal(expectedConfig))
			})
			It("should have the correct balances", func() {
				sp := k.GetHost().GetStatePlugin()
				for addr, acc := range ethGen.Alloc {
					balance := sp.GetBalance(addr)
					cmp := balance.Cmp(acc.Balance)
					Expect(cmp).To(BeZero())
				}
			})
			It("should have the correct code", func() {
				sp := k.GetHost().GetStatePlugin()
				for addr, acc := range ethGen.Alloc {
					code := sp.GetCode(addr)
					cmp := bytes.Compare(code, acc.Code)
					Expect(cmp).To(BeZero())
				}
			})
			It("should have the correct hash", func() {
				sp := k.GetHost().GetStatePlugin()
				for addr, acc := range ethGen.Alloc {
					for key, expectedHash := range acc.Storage {
						actualHash := sp.GetState(addr, key)
						cmp := bytes.Compare(actualHash.Bytes(), expectedHash.Bytes())
						Expect(cmp).To(BeZero())
					}
				}
			})
		})
	})

	Context("On ExportGenesis", func() {
		var (
			actualGenesis core.Genesis
		)
		JustBeforeEach(func() {
			var bz []byte
			bz, err = json.Marshal(ethGen)
			if err != nil {
				panic(err)
			}
			am.InitGenesis(ctx, cdc, bz)

			data := am.ExportGenesis(ctx, cdc)
			if data == nil {
				panic(fmt.Errorf("data is nil"))
			}
			if err = actualGenesis.UnmarshalJSON(data); err != nil {
				panic(err)
			}
		})

		When("the genesis is valid", func() {
			It("should export without fail", func() {
				Expect(actualGenesis).To(Equal(*ethGen))
			})
		})
	})
})
