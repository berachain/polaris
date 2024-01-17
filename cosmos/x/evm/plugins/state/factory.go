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

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/eth/core/state"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ core.StatePluginFactory = (*SPFactory)(nil)

type SPFactory struct {
	// keepers used for balance and account information.
	ak       AccountKeeper
	storeKey storetypes.StoreKey
	plf      events.PrecompileLogFactory

	// Contexts for state plugins
	genesisContext       sdk.Context // "genesis" ---> set in InitGenesis
	minerBuildContext    sdk.Context // "miner" -----> set in PrepareProposal
	insertChainContext   sdk.Context // "insert" ----> set in ProcessProposal
	finalizeBlockContext sdk.Context // "finalize" --> set in Finalize
	latestQueryContext   sdk.Context // "latest" ----> set in PrepareCheckState

	// Query function for getting the context at a given height.
	qfn func() func(height int64, prove bool) (sdk.Context, error) // "historical"
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
func (spf *SPFactory) NewPluginWithMode(mode state.Mode) core.StatePlugin {
	p := NewPlugin(spf.ak, spf.storeKey, spf.qfn, spf.plf)
	switch mode {
	case state.Genesis:
		p.Reset(spf.genesisContext)
	case state.Miner:
		p.Reset(spf.minerBuildContext)
	case state.Insert:
		p.Reset(spf.insertChainContext)
	case state.Finalize:
		p.Reset(spf.finalizeBlockContext)
	case state.Latest:
		fallthrough
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
func (spf *SPFactory) NewPluginAtBlockNumber(blockNumber int64) (core.StatePlugin, error) {
	var (
		ctx sdk.Context
		err error
	)

	if blockNumber >= spf.latestQueryContext.BlockHeight() {
		ctx, _ = spf.latestQueryContext.CacheContext()
	} else {
		// Get the query context at the given height.
		ctx, err = spf.qfn()(blockNumber, false)
		if err != nil {
			return nil, err
		}
	}

	return spf.NewPluginFromContext(ctx), nil
}

// SetGenesisContext updates the SPFactory's genesis context to the provided context.
func (spf *SPFactory) SetGenesisContext(ctx context.Context) {
	spf.genesisContext = sdk.UnwrapSDKContext(ctx)
}

// SetMiningContext updates the SPFactory's minerBuildContext to the provided context.
func (spf *SPFactory) SetLatestMiningContext(ctx context.Context) {
	spf.minerBuildContext = sdk.UnwrapSDKContext(ctx)
}

// SetInsertChainContext updates the SPFactory's insertChainContext to the provided context.
func (spf *SPFactory) SetInsertChainContext(ctx context.Context) {
	spf.insertChainContext = sdk.UnwrapSDKContext(ctx)
}

// SetFinalizeBlockContext updates the SPFactory's finalizeBlockContext to the provided context.
func (spf *SPFactory) SetFinalizeBlockContext(ctx context.Context) {
	spf.finalizeBlockContext = sdk.UnwrapSDKContext(ctx)
}

// SetLatestQueryContext updates the SPFactory's latestQueryContext to the provided context.
// This context will be used for subsequent state queries.
//
// NOTE: From ABCI, this may be the UNSAFE PrepareCheckState context, which should NOT be written
// to.
func (spf *SPFactory) SetLatestQueryContext(ctx context.Context) {
	spf.latestQueryContext = sdk.UnwrapSDKContext(ctx)
}

// SetPrecompileLogFactory sets the PrecompileLogFactory in the SPFactory.
func (spf *SPFactory) SetPrecompileLogFactory(plf events.PrecompileLogFactory) {
	spf.plf = plf
}
