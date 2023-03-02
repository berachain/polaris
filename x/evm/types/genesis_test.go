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

package types

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
)

var _ = Describe("Genesis", func() {
	It("fail if genesis is invalid", func() {
		params := DefaultParams()
		params.EvmDenom = ""
		state := NewGenesisState(*params, nil, nil)
		err := ValidateGenesis(*state)
		Expect(err).To(HaveOccurred())
	})

	It("should return default genesis", func() {
		state := DefaultGenesis()
		Expect(state).ToNot(BeNil())
	})

	It("should return new genesis state", func() {
		atc := make(map[string]*Contract)
		htc := make(map[string]string)
		params := DefaultParams()
		state := NewGenesisState(*params, atc, htc)
		Expect(state).ToNot(BeNil())
		Expect(state.Params).To(Equal(*params))
		Expect(state.AddressToContract).To(Equal(atc))
		Expect(state.HashToCode).To(Equal(htc))
	})

	It("should return a new contract code", func() {
		codeHash := common.HexToHash("0x123")
		code := []byte("0x123")
		slotToValue := make(map[string]string)
		contract := NewContract(codeHash, code, slotToValue)
		Expect(contract).ToNot(BeNil())
		Expect(contract.CodeHash).To(Equal(codeHash.Hex()))
	})

	It("should write to slot", func() {
		slot := common.HexToHash("0x123")
		value := common.HexToHash("0x123")
		contract := NewContract(common.HexToHash("0x123"), []byte("0x123"), make(map[string]string))
		WriteToSlot(slot, value, contract)
		Expect(contract).ToNot(BeNil())
		Expect(contract.SlotToValue[slot.Hex()]).To(Equal(value.Hex()))
	})

	It("should return a new contract code", func() {
		state := DefaultGenesis()
		codeHash := common.HexToHash("0x123")
		code := []byte("0x123")
		WriteCodeToHash(codeHash, code, state.HashToCode)
		code2 := state.HashToCode[codeHash.Hex()]
		Expect(code2).To(Equal(string(code)))

	})
})
