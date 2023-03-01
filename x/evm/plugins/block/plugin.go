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

package block

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/core"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/x/evm/plugins"
)

// TODO: change this.
const bf = uint64(1)

// `Plugin` is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosStargazer
	UpdateOffChainStorage(sdk.Context, *coretypes.StargazerBlock)
	// `TrackHistoricalStargazerHeader` saves the latest historical-info and deletes the oldest
	// heights that are below pruning height.
	TrackHistoricalStargazerHeader(ctx sdk.Context, header *coretypes.StargazerHeader)
	// `GetStargazerBlock` returns the block from the store at the height specified in the context.
	GetStargazerHeader(ctx sdk.Context, height int64) (*coretypes.StargazerHeader, bool)
	// `SetStargazerHeader` saves a block to the store.
	SetStargazerHeader(ctx sdk.Context, header *coretypes.StargazerHeader) error
	// `PruneStargazerHeader` prunes a stargazer block from the store.
	PruneStargazerHeader(ctx sdk.Context, header *coretypes.StargazerHeader) error
	core.BlockPlugin
}

// `plugin` keeps track of stargazer blocks via headers.
type plugin struct {
	// `ctx` is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// `storekey` is the store key for the header store.
	storekey storetypes.StoreKey
	//  `offchainStore` is the offchain store, used for accessing offchain data.
	offchainStore storetypes.CacheKVStore
	// `getQueryContext` allows for querying block headers.
	getQueryContext func(height int64, prove bool) (sdk.Context, error)
}

// `NewPlugin` creates a new instance of the block plugin from the given context.
func NewPlugin(offchainStore storetypes.CacheKVStore, storekey storetypes.StoreKey) Plugin {
	return &plugin{
		offchainStore: offchainStore,
		storekey:      storekey,
	}
}

// `Prepare` implements core.BlockPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}

// `BaseFee` returns the base fee for the current block.
// TODO: implement properly with DynamicFee Module of some kind.
//
// `BaseFee` implements core.BlockPlugin.
func (p *plugin) BaseFee() uint64 {
	return bf
}
