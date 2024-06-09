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

package block

import (
	"context"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins"
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Plugin interface {
	plugins.HasGenesis
	core.BlockPlugin
}

type plugin struct {
	// ctx is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// storekey is the store key for the header store.
	storekey storetypes.StoreKey
	// getQueryContext allows for querying block headers.
	getQueryContext func() func(height int64, prove bool) (sdk.Context, error)
}

func NewPlugin(
	storekey storetypes.StoreKey,
	qfn func() func(height int64, prove bool) (sdk.Context, error),
) Plugin {
	return &plugin{
		storekey:        storekey,
		getQueryContext: qfn,
	}
}

// Prepare implements core.BlockPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}
