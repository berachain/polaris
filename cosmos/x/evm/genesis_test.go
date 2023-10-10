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
	"math/big"
	"testing"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"pkg.berachain.dev/polaris/cosmos/config"
	"pkg.berachain.dev/polaris/cosmos/precompile/staking"
	testutil "pkg.berachain.dev/polaris/cosmos/testutil"
	"pkg.berachain.dev/polaris/cosmos/x/evm"
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ethGen = core.DefaultGenesis
)

func TestGenesis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm")
}

var _ = Describe("", func() {
	var (
		cdc codec.JSONCodec
		ctx sdk.Context
		sc  ethprecompile.StatefulImpl
		ak  state.AccountKeeper
		sk  stakingkeeper.Keeper
		k   *keeper.Keeper
		am  evm.AppModule
		err error
	)

	BeforeEach(func() {
		ctx, ak, _, sk = testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))
		ctx = ctx.WithBlockHeight(0)
		sc = staking.NewPrecompileContract(ak, &sk)
		cfg := config.DefaultConfig()
		ethGen.Config = params.DefaultChainConfig
		cfg.Node.DataDir = GinkgoT().TempDir()
		cfg.Node.KeyStoreDir = GinkgoT().TempDir()
		k = keeper.NewKeeper(
			ak, sk,
			testutil.EvmKey,
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
			func() func(height int64, prove bool) (sdk.Context, error) {
				return func(height int64, prove bool) (sdk.Context, error) {
					return ctx, nil
				}
			},
			log.NewTestLogger(GinkgoT()),
			cfg,
		)
		err = k.SetupPrecompiles()
		Expect(err).ToNot(HaveOccurred())
		am = evm.NewAppModule(k, ak)
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
			It("should succeed without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
			It("should contain the same genesis header values", func() {
				bp := k.Host.GetBlockPlugin()
				expectedHeader := ethGen.ToBlock().Header()
				Expect(bp.GetHeaderByNumber(0)).To(Equal(expectedHeader))
			})
			It("should have the correct balances", func() {
				sp := k.Host.GetStatePlugin()
				for addr, acc := range ethGen.Alloc {
					balance := sp.GetBalance(addr)
					cmp := balance.Cmp(acc.Balance)
					Expect(cmp).To(BeZero())
				}
			})
			It("should have the correct code", func() {
				sp := k.Host.GetStatePlugin()
				for addr, acc := range ethGen.Alloc {
					code := sp.GetCode(addr)
					cmp := bytes.Compare(code, acc.Code)
					Expect(cmp).To(BeZero())
				}
			})
			It("should have the correct hash", func() {
				sp := k.Host.GetStatePlugin()
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
				ethGen.Config = nil
				ethGen.BaseFee = big.NewInt(int64(params.InitialBaseFee))
				Expect(actualGenesis).To(Equal(*ethGen))
			})
		})
	})
})
