// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package evm_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"cosmossdk.io/log"

	"github.com/berachain/polaris/cosmos/config"
	"github.com/berachain/polaris/cosmos/runtime/chain"
	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/cosmos/x/evm"
	"github.com/berachain/polaris/cosmos/x/evm/keeper"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/eth/core"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/params"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/consensus/beacon"
	ethparams "github.com/ethereum/go-ethereum/params"

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

var _ = Describe("Genesis", func() {
	var (
		ctx sdk.Context
		ak  state.AccountKeeper
		k   *keeper.Keeper
		am  evm.AppModule
		err error
	)

	BeforeEach(func() {
		ctx, ak, _, _ = testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))
		ctx = ctx.WithBlockHeight(0)
		cfg := config.DefaultPolarisConfig()
		ethGen.Config = params.DefaultChainConfig
		cfg.Node.DataDir = GinkgoT().TempDir()
		cfg.Node.KeyStoreDir = GinkgoT().TempDir()
		k = keeper.NewKeeper(
			ak,
			testutil.EvmKey,
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{}...)
			},
			func() func(height int64, prove bool) (sdk.Context, error) {
				return func(height int64, prove bool) (sdk.Context, error) {
					return ctx, nil
				}
			},
			cfg,
		)
		err = k.Setup(
			chain.New(core.NewChain(k.Host, params.DefaultChainConfig, beacon.NewFaker()), nil),
			nil,
		)
		Expect(err).ToNot(HaveOccurred())

		err = k.SetupPrecompiles()
		Expect(err).ToNot(HaveOccurred())
		am = evm.NewAppModule(k, ak)
	})

	Describe("On InitGenesis", func() {
		var bz []byte
		BeforeEach(func() {
			bz, err = json.Marshal(ethGen)
			Expect(err).ToNot(HaveOccurred())
		})
		JustBeforeEach(func() {
			am.InitGenesis(ctx, nil, bz)
		})

		Context("when the genesis is valid", func() {
			It("should succeed without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
			It("should contain the same genesis header values", func() {
				bp := k.Host.GetBlockPlugin()
				expectedHeader := ethGen.ToBlock().Header()
				Expect(bp.GetHeaderByNumber(0)).To(Equal(expectedHeader))
			})
			It("should have the correct balances", func() {
				spf := k.Host.GetStatePluginFactory()
				sp := spf.NewPluginFromContext(ctx)
				for addr, acc := range ethGen.Alloc {
					balance := sp.GetBalance(addr)
					Expect(balance).To(Equal(acc.Balance))
				}
			})
			It("should have the correct code", func() {
				spf := k.Host.GetStatePluginFactory()
				sp := spf.NewPluginFromContext(ctx)
				for addr, acc := range ethGen.Alloc {
					code := sp.GetCode(addr)
					Expect(code).To(Equal(acc.Code))
				}
			})
			It("should have the correct hash", func() {
				spf := k.Host.GetStatePluginFactory()
				sp := spf.NewPluginFromContext(ctx)
				for addr, acc := range ethGen.Alloc {
					for key, expectedHash := range acc.Storage {
						actualHash := sp.GetState(addr, key)
						Expect(actualHash.Bytes()).To(Equal(expectedHash.Bytes()))
					}
				}
			})
		})
	})

	Describe("On ExportGenesis", func() {
		var (
			actualGenesis core.Genesis
			bz            []byte
		)
		BeforeEach(func() {
			bz, err = json.Marshal(ethGen)
			Expect(err).ToNot(HaveOccurred())
			am.InitGenesis(ctx, nil, bz)
		})
		JustBeforeEach(func() {
			data := am.ExportGenesis(ctx, nil)
			Expect(data).ToNot(BeNil())
			err = actualGenesis.UnmarshalJSON(data)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the genesis is valid", func() {
			It("should export without fail", func() {
				ethGen.Config = nil
				ethGen.BaseFee = big.NewInt(int64(ethparams.InitialBaseFee))
				Expect(actualGenesis).To(Equal(*ethGen))
			})
		})
	})
})
