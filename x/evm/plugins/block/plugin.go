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
	cbft "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
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
	// `SetQueryContextFn` sets the query context func for the plugin.
	SetQueryContextFn(fn func(height int64, prove bool) (sdk.Context, error))

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
	// getQueryContext allows for querying block headers.
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

// blockHashFromCosmosContext returns the block hash from the provided Cosmos SDK context.
// If the context contains a valid header hash, it is converted to a common.Hash and returned.
// Otherwise, if the header hash is not set (e.g., for checkTxState), the hash is computed
// from the context's block header and returned as a common.Hash. If the block header is invalid,
// the function returns an empty common.Hash and logs an error.
func blockHashFromCosmosContext(ctx sdk.Context) common.Hash {
	// Check if the context contains a header hash
	headerHash := ctx.HeaderHash()
	if len(headerHash) != 0 {
		return common.BytesToHash(headerHash)
	}

	// If the header hash is not set, compute the hash from the context's block header
	contextBlockHeader := ctx.BlockHeader()
	header, err := cbft.HeaderFromProto(&contextBlockHeader)
	if err != nil {
		// If the block header is invalid, return an empty hash
		return common.Hash{}
	}

	// Convert the computed hash to a common.Hash and return it
	return common.BytesToHash(header.Hash())
}

// `blockGasLimitFromCosmosContext` returns the maximum gas limit for the current block, as defined
// by either the block gas meter or the consensus parameters if the gas meter is not set or is an
// InfiniteGasMeter. If neither the gas meter nor the consensus parameters are available, it
// returns 0. This shouldn't be an issue in practice but we include this function for completeness
// defensive programming purposes.
func blockGasLimitFromCosmosContext(ctx sdk.Context) uint64 {
	blockGasMeter := ctx.BlockGasMeter()
	if blockGasMeter == nil || blockGasMeter.Limit() == 0 {
		cp := ctx.ConsensusParams()
		if cp == nil || cp.Block == nil {
			return 0
		}
		return uint64(cp.Block.MaxGas)
	}

	return blockGasMeter.Limit()
}
