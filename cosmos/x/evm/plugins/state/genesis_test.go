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
