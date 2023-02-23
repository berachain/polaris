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

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

// `InitGenesis` takes in a pointer to a genesis state object and populates the KV store.
func (p *plugin) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	p.Reset(ctx)

	for contract, contractState := range data.AddressToContract {
		// Get the code from the code hash to code map.
		code := common.Hex2Bytes(data.HashToCode[contractState.CodeHash])
		addr := common.HexToAddress(contract)
		p.SetCode(addr, code)

		// Loop through the contract slot to value map and set the state.
		for k, v := range contractState.SlotToValue {
			slot := common.HexToHash(k)
			value := common.HexToHash(v)
			p.SetState(addr, slot, value)
		}
	}

	p.Finalize()
}

// `Export` genesis modifies a pointer to a genesis state object and populates it.
func (p *plugin) ExportGenesis(ctx sdk.Context, data *types.GenesisState) {
	p.Reset(ctx)

	// Address to contract state map.
	atc := make(map[string]*types.Contract)
	// Code hash to code map.
	htc := make(map[string]string)

	// Iterate over all the code records and add them to the genesis state.
	p.IterateCode(func(address common.Address, code []byte) bool {
		// Get the contract code hash.
		codeHash := common.BytesToHash(code).Hex()

		// Add the code hash and code to the code hash to code map.
		htc[codeHash] = common.Bytes2Hex(code)

		return false // keep iterating
	})

	// Iterate over all the contract state records and add them to the map.
	p.IterateState(func(address common.Address, key, value common.Hash) bool {
		// If the slot to value map is nil, allocate memory for it.
		if atc[address.Hex()] == nil {
			atc[address.Hex()] = &types.Contract{
				SlotToValue: make(map[string]string),
			}
		}

		// Set the slots to values map.
		atc[address.Hex()].SlotToValue[key.Hex()] = value.Hex()
		return false // keep iterating
	})

	// Add to the genesis state.
	data.AddressToContract = atc
	data.HashToCode = htc

	p.Finalize()
}
