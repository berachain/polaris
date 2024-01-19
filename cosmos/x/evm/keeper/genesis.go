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

package keeper

import (
	"github.com/berachain/polaris/cosmos/x/evm/plugins"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis is called during the InitGenesis.
func (k *Keeper) InitGenesis(ctx sdk.Context, genState *core.Genesis) error {
	// TODO: Feels jank as fuck lol, but it works.
	genState.Config = k.chain.Config()

	// Initialize all the plugins.
	for _, plugin := range k.Host.GetAllPlugins() {
		// checks whether plugin implements methods of HasGenesis and executes them if it does
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			if err := plugin.InitGenesis(ctx, genState); err != nil {
				return err
			}
		}
	}

	// Insert to chain with the genesis context. The plugins are already prepared with their
	// InitGenesis.
	k.spf.SetGenesisContext(ctx)
	return k.chain.WriteGenesisBlock(genState.ToBlock())
}

// ExportGenesis returns the exported genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *core.Genesis {
	genesisState := new(core.Genesis)
	for _, plugin := range k.Host.GetAllPlugins() {
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			plugin.ExportGenesis(ctx, genesisState)
		}
	}
	return genesisState
}
