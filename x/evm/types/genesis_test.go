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
		state := NewGenesisState(*params, nil)
		err := ValidateGenesis(*state)
		Expect(err).To(HaveOccurred())
	})

	It("should write address to code hash", func() {
		addressToCodeHash := make(map[string]string)
		addr := common.HexToAddress("0x1")
		ch := common.HexToHash("0x2")
		WriteAddressToCodeHash(addr, ch, &addressToCodeHash)

		Expect(addressToCodeHash[addr.Hex()]).To(Equal(ch.Hex()))
	})

	It("should write codeHash to code", func() {
		codeHashToCode := make(map[string]string)
		ch := common.HexToHash("0x1")
		code := []byte("0x2")
		WriteCodeHashToCode(ch, code, &codeHashToCode)

		Expect(codeHashToCode[ch.Hex()]).To(Equal(string(code)))
	})

	It("should write slot to value", func() {
		var state = StateRecord{
			State: make(map[string]string),
		}
		slot := common.HexToHash("0x1")
		value := common.HexToHash("0x2")
		WriteSlotToValue(slot, value, &state)

		Expect(state.State[slot.Hex()]).To(Equal(value.Hex()))
	})

	It("should write address to state data", func() {
		addressToStateData := make(map[string]*StateRecord)
		addr := common.HexToAddress("0x1")
		state := StateRecord{
			State: make(map[string]string),
		}

		WriteSlotToValue(common.HexToHash("0x1"), common.HexToHash("0x2"), &state)
		WriteAddressToStateData(addr, &state, &addressToStateData)

		Expect(addressToStateData[addr.Hex()].State).To(Equal(state.State))
	})

	It("should create new contract state", func() {
		addressToCodeHash := make(map[string]string)
		codeHashToCode := make(map[string]string)
		addressToStateData := make(map[string]*StateRecord)
		cs := NewContractState(addressToCodeHash, codeHashToCode, addressToStateData)
		Expect(cs.AddressToCodeHash).To(Equal(addressToCodeHash))
		Expect(cs.CodeHashToCode).To(Equal(codeHashToCode))
		Expect(cs.AddressToStateData).To(Equal(addressToStateData))
	})

	It("should create default genesis", func() {
		genesis := DefaultGenesis()
		Expect(genesis.Params).To(Equal(*DefaultParams()))
		Expect(genesis.AddressToContractState).To(Equal(make(map[string]*ContractState)))
	})
})
