// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
