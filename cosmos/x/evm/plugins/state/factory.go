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
	"context"
	"errors"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events"
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// var _ core.State = (*SPFactory)(nil)

type SPFactory struct {
	// keepers used for balance and account information.
	ak       AccountKeeper
	storeKey storetypes.StoreKey
	plf      events.PrecompileLogFactory

	// Contexts for queries
	latestQueryContext sdk.Context // "latest"
	minerBuildContext  sdk.Context // "miner"
	insertChainContext sdk.Context
	qfn                func() func(height int64, prove bool) (sdk.Context, error) // "historical"
}

// NewSPFactory creates a new SPFactory instance with the provided AccountKeeper,
// store key, query function, and PrecompileLogFactory.
func NewSPFactory(
	ak AccountKeeper,
	storeKey storetypes.StoreKey,
	qfn func() func(height int64, prove bool) (sdk.Context, error),
) *SPFactory {
	return &SPFactory{
		ak:       ak,
		storeKey: storeKey,
		qfn:      qfn,
	}
}

// NewPluginFromContext creates a new Plugin instance using the current SPFactory's
// configuration and the provided context.
func (spf *SPFactory) NewPluginWithMode(mode string) core.StatePlugin {
	p := NewPlugin(spf.ak, spf.storeKey, spf.qfn, spf.plf)
	switch mode {
	case "miner":
		p.Reset(spf.minerBuildContext)
	case "chain":
		p.Reset(spf.insertChainContext)
	default:
		p.Reset(spf.latestQueryContext)
	}
	return p
}

// NewPluginWithContext creates a new StatePlugin instance using the provided context.
// It initializes the plugin with the current SPFactory's account keeper, store key,
// query function, and precompile log factory, then resets the plugin's context to the
// one provided.
func (spf *SPFactory) NewPluginFromContext(ctx context.Context) core.StatePlugin {
	p := NewPlugin(spf.ak, spf.storeKey, spf.qfn, spf.plf)
	p.Reset(ctx)
	return p
}

// NewPluginAtBlockNumber creates a new StatePlugin instance using the provided block.
func (spf *SPFactory) NewPluginAtBlockNumber(blockNumber uint64) (core.StatePlugin, error) {
	var ctx sdk.Context
	// Ensure the query context function is set.
	if spf.qfn == nil {
		return nil, errors.New("no query context function set in host chain")
	}

	// // NOTE: the PreBlock and BeginBlock state changes will not have been applied to the state
	// // at this point.
	// // This is kind of bad since queries from JSON-RPC (i.e eth_call estimateGas etc.)
	// // won't be able to do this
	// // ontop of a state that has these updates for the block.
	// // TODO: Fix this.
	int64Number := int64(blockNumber)
	// TODO: the GTE may be hiding a larger issue with the timing of the NewHead channel stuff.
	// Investigate and hopefully remove this GTE.
	if int64Number >= spf.latestQueryContext.BlockHeight() {
		// TODO: Manager properly
		if spf.latestQueryContext.MultiStore() == nil {
			ctx = spf.latestQueryContext.WithEventManager(sdk.NewEventManager())
		} else {
			ctx, _ = spf.latestQueryContext.CacheContext()
		}
	} else {
		// Get the query context at the given height.
		var err error
		ctx, err = spf.qfn()(int64Number, false)
		if err != nil {
			return nil, err
		}
	}

	return spf.NewPluginFromContext(ctx), nil
	// // Create a State Plugin with the requested chain height.
	// sp := NewPlugin(p.ak, p.storeKey, p.qfn, p.plf)

	// // TODO: Manager properly
	// if p.latestQueryContext.MultiStore() != nil {
	// 	sp.Reset(ctx)
	// }
	// // return sp, nil
	// ctx, err := spf.qfn()(int64(blockNumber), false)
	// if err != nil {
	// 	return nil, err
	// }
}

// SetMiningContext updates the SPFactory's minerBuildContext to the provided context.
func (spf *SPFactory) SetLatestMiningContext(ctx context.Context) {
	spf.minerBuildContext = sdk.UnwrapSDKContext(ctx)
}

// SetInsertChainContext updates the SPFactory's minerBuildContext to the provided context.
func (spf *SPFactory) SetInsertChainContext(ctx context.Context) {
	spf.minerBuildContext = sdk.UnwrapSDKContext(ctx)
}

// SetLatestQueryContext updates the SPFactory's latestQueryContext to the provided context.
// This context will be used for subsequent state queries.
func (spf *SPFactory) SetLatestQueryContext(ctx context.Context) {
	spf.latestQueryContext = sdk.UnwrapSDKContext(ctx)
}

// SetPrecompileLogFactory sets the PrecompileLogFactory in the SPFactory.
func (spf *SPFactory) SetPrecompileLogFactory(plf events.PrecompileLogFactory) {
	spf.plf = plf
}
