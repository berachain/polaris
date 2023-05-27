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
	"pkg.berachain.dev/polaris/eth/common"
)

// DefaultGenesis is the default genesis state.
func DefaultGenesis() *GenesisState {
	atc := make(map[string]*Contract)
	htc := make(map[string]string)

	return &GenesisState{
		Params:            *DefaultParams(),
		AddressToContract: atc,
		HashToCode:        htc,
	}
}

// ValidateGenesis is used to validate the genesis state.
func ValidateGenesis(data GenesisState) error {
	return data.Params.ValidateBasic()
}

// NewGenesisState creates a new `GenesisState` object.
func NewGenesisState(params Params, atc map[string]*Contract, htc map[string]string) *GenesisState {
	return &GenesisState{
		Params:            params,
		AddressToContract: atc,
		HashToCode:        htc,
	}
}

// NewContract creates a new `Contract` object.
func NewContract(codeHash common.Hash, slotToValue map[string]string) *Contract {
	return &Contract{
		CodeHash:    codeHash.Hex(),
		SlotToValue: slotToValue,
	}
}

// WriteToSlot takes in a slot, value and pointer to a contract and writes the value to the slot.
func WriteToSlot(slot common.Hash, value common.Hash, contract *Contract) {
	contract.SlotToValue[slot.Hex()] = value.Hex()
}

// WriteCodeToHash takes in a code hash, code and map of code hashes to code and writes the code to the code hash.
func WriteCodeToHash(codeHash common.Hash, code []byte, htc map[string]string) {
	htc[codeHash.Hex()] = string(code)
}
