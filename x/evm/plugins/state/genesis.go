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

	for _, cr := range data.CodeRecords {
		addr := common.HexToAddress(cr.Address)
		p.SetCode(addr, cr.Code)
	}

	for _, sr := range data.StateRecords {
		addr := common.HexToAddress(sr.Address)
		slot := common.BytesToHash(sr.Slot)
		value := common.BytesToHash(sr.Value)
		p.SetState(addr, slot, value)
	}

	p.Finalize()
}

// `Export` genesis modifies a pointer to a genesis state object and populates it.
func (p *plugin) ExportGenesis(ctx sdk.Context, data *types.GenesisState) {
	p.Reset(ctx)

	// Iterate over all the code records and add them to the genesis state.
	p.IterateCode(func(addr common.Address, code []byte) bool {
		data.CodeRecords = append(data.CodeRecords, types.CodeRecord{
			Address: addr.Hex(),
			Code:    code,
		})
		return false
	})

	// Iterate over all the state records and add them to the genesis state.
	p.IterateState(func(addr common.Address, key, value common.Hash) bool {
		data.StateRecords = append(data.StateRecords, types.StateRecord{
			Address: addr.Hex(),
			Slot:    key.Bytes(),
			Value:   value.Bytes(),
		})
		return false
	})

	p.Finalize()
}
