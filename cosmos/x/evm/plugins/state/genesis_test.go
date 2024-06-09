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

package state_test

import (
	"math/big"

	"cosmossdk.io/log"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Genesis", func() {
	var (
		ctx  sdk.Context
		sp   state.Plugin
		code = []byte("code")
	)

	BeforeEach(func() {
		var ak state.AccountKeeper
		ctx, ak, _, _ = testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))
		sp = state.NewPlugin(ak, testutil.EvmKey, nil, &mockPLF{})

		// Create account for alice, bob
		acc := ak.NewAccountWithAddress(ctx, bob[:])
		Expect(acc.SetSequence(2)).To(Succeed())
		ak.SetAccount(ctx, acc)
		sp.Reset(ctx)
	})

	It("should fail init genesis on bad data", func() {
		genesis := new(core.Genesis)
		genesis.Alloc = make(core.GenesisAlloc)
		genesis.Alloc[bob] = core.GenesisAccount{Nonce: 1}
		// Call Init Genesis and expect bob's case to error because of nonce mismatch.
		Expect(sp.InitGenesis(ctx, genesis)).To(
			MatchError("account nonce mismatch for (0x0000000000000000000000000000000000626f62) between auth (2) and evm (1) genesis state"), //nolint:lll // test.
		)
	})

	It("should init and export genesis", func() {
		genesis := new(core.Genesis)
		genesis.Alloc = make(core.GenesisAlloc)
		genesis.Alloc[alice] = core.GenesisAccount{
			Balance: big.NewInt(5e18),
			Storage: map[common.Hash]common.Hash{
				common.BytesToHash([]byte("key")): common.BytesToHash([]byte("value")),
			},
			Code:  code,
			Nonce: 1,
		}
		genesis.Alloc[bob] = core.GenesisAccount{
			Balance: big.NewInt(2e18),
			Nonce:   2,
		}

		// Call Init Genesis
		Expect(sp.InitGenesis(ctx, genesis)).To(Succeed())

		// Check the code, hash, balance.
		Expect(sp.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash(code)))
		Expect(sp.GetCodeHash(bob)).To(Equal(crypto.Keccak256Hash(nil)))
		Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(5e18)))
		Expect(sp.GetBalance(bob)).To(Equal(big.NewInt(2e18)))
		Expect(sp.GetCode(alice)).To(Equal(code))

		// Very exported genesis is equal.
		var exportedGenesis core.Genesis
		sp.ExportGenesis(ctx, &exportedGenesis)
		Expect(exportedGenesis.Alloc).To(Equal(genesis.Alloc))
	})
})
