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

package state

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

var (
	alice = testutil.Alice
)

var _ = Describe("Genesis", func() {
	var (
		ctx      sdk.Context
		sp       Plugin
		codeHash common.Hash
		code     []byte
		slot     common.Hash
		value    common.Hash
	)

	BeforeEach(func() {
		var ak AccountKeeper
		var bk BankKeeper
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		sp = NewPlugin(ak, bk, testutil.EvmKey, "abera", nil)

		// Create account for alice.
		sp.Reset(ctx)
		sp.CreateAccount(alice)
		sp.Finalize()

		code = []byte("code")
		codeHash = crypto.Keccak256Hash(code)
		slot = common.HexToHash("0x456")
		value = common.HexToHash("0x789")
	})

	It("should init and export genesis", func() {
		genesis := types.DefaultGenesis()

		// New Contract.
		contract := types.Contract{
			CodeHash: codeHash.Hex(),
			SlotToValue: map[string]string{
				slot.Hex(): value.Hex(),
			},
		}

		// Set the address to contract.
		genesis.AddressToContract[alice.Hex()] = &contract

		// Set the code hash to code.
		genesis.HashToCode[codeHash.Hex()] = string(code)

		// Init Genesis.
		sp.InitGenesis(ctx, genesis)

		// Check that the code is set.
		sp.Reset(ctx)
		Expect(sp.GetCode(alice)).To(Equal(code))
		sp.Finalize()

		// Check that the code hash is set.
		sp.Reset(ctx)
		Expect(sp.GetCodeHash(alice)).To(Equal(codeHash))
		sp.Finalize()

		// Check that the storage is set.
		sp.Reset(ctx)
		Expect(sp.GetState(alice, slot)).To(Equal(value))
		sp.Finalize()

		// Export Genesis.
		exportedGenesis := types.GenesisState{}
		sp.ExportGenesis(ctx, &exportedGenesis)

		// Check that the code is exported.
		Expect(exportedGenesis.AddressToContract).To(Equal(genesis.AddressToContract))
		// Check that the hash to code is exported.
		Expect(exportedGenesis.HashToCode).To(Equal(genesis.HashToCode))
		// Check that the storage is exported.
		Expect(
			exportedGenesis.AddressToContract[alice.Hex()].SlotToValue).
			To(Equal(genesis.AddressToContract[alice.Hex()].SlotToValue))
	})
})
