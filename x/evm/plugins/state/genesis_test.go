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

// import (
// 	. "github.com/onsi/ginkgo/v2"

// 	. "github.com/onsi/gomega"
// )

// var (
// 	alice = testutil.Alice
// )

// var _ = Describe("Genesis", func() {
// 	var (
// 		ctx sdk.Context
// 		sp  Plugin
// 	)

// 	BeforeEach(func() {
// 		var ak AccountKeeper
// 		var bk BankKeeper
// 		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
// 		sp = NewPlugin(ak, bk, testutil.EvmKey, "abera", nil)
// 		sp.InitGenesis(ctx, &types.GenesisState{
// 			CodeRecords: []types.CodeRecord{
// 				{
// 					Address: alice.Hex(),
// 					Code:    []byte("code"),
// 				},
// 			},
// 			StateRecords: []types.StateRecord{
// 				{
// 					Address: alice.Hex(),
// 					Slot:    []byte("slot"),
// 					Value:   []byte("value"),
// 				},
// 			},
// 		})
// 	})

// 	It("should export current state", func() {
// 		var gs types.GenesisState
// 		sp.ExportGenesis(ctx, &gs)

// 		Expect(gs.CodeRecords).To(HaveLen(1))
// 		Expect(gs.CodeRecords[0].Address).To(Equal(alice.Hex()))
// 	})
// })
