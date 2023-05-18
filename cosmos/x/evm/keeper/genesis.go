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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// InitGenesis is called during the InitGenesis.
func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) error {
	// We configure the logger here because we want to get the logger off the context opposed to allocating a new one.
	k.ConfigureGethLogger(ctx)

	// Initialize all the plugins.
	for _, plugin := range k.host.GetAllPlugins() {
		// checks whether plugin implements methods of HasGenesis and executes them if it does
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			plugin.InitGenesis(ctx, &genState)
		}
	}

	go func() {
		time.Sleep(2 * time.Second) //nolint:gomnd // we will fix this eventually.
		// Start the polaris "Node" in order to spin up things like the JSON-RPC server.
		if err := k.polaris.StartServices(); err != nil {
			return
		}
	}()

	return nil
}

// ExportGenesis returns the exported genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesisState := new(types.GenesisState)
	for _, plugin := range k.host.GetAllPlugins() {
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			plugin.ExportGenesis(ctx, genesisState)
		}
	}
	return genesisState
}
