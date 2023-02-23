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

import "pkg.berachain.dev/stargazer/eth/common"

// `DefaultGenesis` is the default genesis state.
func DefaultGenesis() *GenesisState {
	acs := make(map[string]*ContractState)
	return &GenesisState{
		Params:                 *DefaultParams(),
		AddressToContractState: acs,
	}
}

// `ValidateGenesis` is used to validate the genesis state.
func ValidateGenesis(data GenesisState) error {
	return data.Params.ValidateBasic()
}

// `NewGenesisState` creates a new `GenesisState` object.
func NewGenesisState(params Params, acs map[string]*ContractState) *GenesisState {
	return &GenesisState{
		Params:                 params,
		AddressToContractState: acs,
	}
}

// `NewContractState` creates a new `ContractState` object.
func NewContractState(
	addressToCodeHash map[string]string,
	codeHashToCode map[string]string,
	addressToStateData map[string]*StateRecord,
) *ContractState {
	return &ContractState{
		AddressToCodeHash:  addressToCodeHash,
		CodeHashToCode:     codeHashToCode,
		AddressToStateData: addressToStateData,
	}
}

// `WriteAddressToCodeHash` writes an address to a code hash in the contract state.
func WriteAddressToCodeHash(addr common.Address, codeHash common.Hash, addressToCodeHashMap *map[string]string) {
	(*addressToCodeHashMap)[addr.Hex()] = codeHash.Hex()
}

// `WriteCodeHashToCode` writes a code hash to code in the contract state.
func WriteCodeHashToCode(codeHash common.Hash, code []byte, codeHashToCodeMap *map[string]string) {
	(*codeHashToCodeMap)[codeHash.Hex()] = string(code)
}

// `WriteSlotToValue` writes a slot to a value in the state record.
func WriteSlotToValue(slot, value common.Hash, state *StateRecord) {
	state.State[slot.Hex()] = value.Hex()
}

// `WriteAddressToStateData` writes an address to state record in the contract state.
func WriteAddressToStateData(addr common.Address, state *StateRecord, addressToStateDataMap *map[string]*StateRecord) {
	(*addressToStateDataMap)[addr.Hex()] = state
}
