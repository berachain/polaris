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
	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Genesis", func() {
	var (
		ctx sdk.Context
		sp  state.Plugin
		// code []byte
	)

	BeforeEach(func() {
		var ak state.AccountKeeper
		var bk state.BankKeeper
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		sp = state.NewPlugin(ak, bk, testutil.EvmKey, &mockConfigurationPlugin{}, nil)

		// Create account for alice.
		sp.Reset(ctx)
		sp.CreateAccount(alice)
		sp.Finalize()

		// code = []byte("code")
	})

	It("should init and export genesis", func() {
		// genesis := types.DefaultGenesis()

		// // New Contract.
		// contract := types.Contract{
		// 	CodeHash: codeHash.Hex(),
		// 	SlotToValue: map[string]string{
		// 		slot.Hex(): value.Hex(),
		// 	},
		// }

		// // Set the address to contract.
		// genesis.AddressToContract[alice.Hex()] = &contract

		// // Set the code hash to code.
		// genesis.HashToCode[codeHash.Hex()] = string(code)

		// // Init Genesis.
		// sp.InitGenesis(ctx, genesis)

		// // Check that the code is set.
		// sp.Reset(ctx)
		// Expect(sp.GetCode(alice)).To(Equal(code))
		// sp.Finalize()

		// // Check that the code hash is set.
		// sp.Reset(ctx)
		// Expect(sp.GetCodeHash(alice)).To(Equal(codeHash))
		// sp.Finalize()

		// // Check that the storage is set.
		// sp.Reset(ctx)
		// Expect(sp.GetState(alice, slot)).To(Equal(value))
		// sp.Finalize()

		// // Export Genesis.
		// exportedGenesis := types.GenesisState{}
		// sp.ExportGenesis(ctx, &exportedGenesis)

	})
})
