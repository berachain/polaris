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

package historical

import (
	"context"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins"
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/params"
)

// Plugin is the interface that must be implemented by the plugin.
type Plugin interface {
	core.HistoricalPlugin
	plugins.HasGenesis
}

// plugin keeps track of polaris blocks via headers.
type plugin struct {
	// ctx is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// chainConfig stores the chain configuration for the evm chain.make
	chainConfig *params.ChainConfig
	// bp represents the block plugin, used for accessing historical block headers.
	bp core.BlockPlugin
	// storekey is the store key for the header store.
	storeKey storetypes.StoreKey
}

// NewPlugin creates a new instance of the block plugin from the given context.
func NewPlugin(
	chainConfig *params.ChainConfig, bp core.BlockPlugin,
	_ storetypes.StoreKey, storekey storetypes.StoreKey,
) Plugin {
	return &plugin{
		chainConfig: chainConfig,
		bp:          bp,
		storeKey:    storekey,
	}
}

// Prepare implements core.HistoricalPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}
