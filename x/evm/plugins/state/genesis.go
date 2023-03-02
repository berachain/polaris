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

	for addr, contract := range data.AddressToContract {
		// Set the contract code.
		address := common.HexToAddress(addr)
		code := []byte(data.HashToCode[contract.CodeHash])
		p.SetCode(address, code)

		// Set the contract state.
		for k, v := range contract.SlotToValue {
			slot := common.HexToHash(k)
			value := common.HexToHash(v)
			p.SetState(address, slot, value)
		}
	}

	p.Finalize()
}

// `Export` genesis modifies a pointer to a genesis state object and populates it.
func (p *plugin) ExportGenesis(ctx sdk.Context, data *types.GenesisState) {
	p.Reset(ctx)

	// Allocate memory for the address to contract map if it is nil.
	if data.AddressToContract == nil {
		data.AddressToContract = make(map[string]*types.Contract)
	}
	// Allocate memory for the hash to code map if it is nil.
	if data.HashToCode == nil {
		data.HashToCode = make(map[string]string)
	}

	p.IterateCode(func(address common.Address, code []byte) bool {
		// Get the contract code hash.
		codeHash := p.GetCodeHash(address)
		// If the contract is nil, allocate memory for it.
		if data.AddressToContract[address.Hex()] == nil {
			data.AddressToContract[address.Hex()] = &types.Contract{}
		}
		data.AddressToContract[address.Hex()].CodeHash = codeHash.Hex()
		// Add the code hash and code to the code hash to code map.
		data.HashToCode[codeHash.Hex()] = string(code)
		return false // keep iterating
	})

	p.IterateState(func(addr common.Address, key, value common.Hash) bool {
		// if the slot to value map is nil on the contract, allocate memory for it.
		if data.AddressToContract[addr.Hex()].SlotToValue == nil {
			data.AddressToContract[addr.Hex()].SlotToValue = make(map[string]string)
		}

		// Set the slots to value map.
		data.AddressToContract[addr.Hex()].SlotToValue[key.Hex()] = value.Hex()

		return false // keep iterating
	})

	p.Finalize()
}
