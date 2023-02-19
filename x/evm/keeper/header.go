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
	"math/big"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/x/evm/key"
)

// ===========================================================================
// Stargazer Block Tracking
// ===========================================================================.

// `numHistoricalBlocks` is the number of historical blocks to keep in the store. This is set to
// 256, as this is the furthest back the BLOCKHASH opcode is allowed to look back.
const numHistoricalBlocks int64 = 256

// `TrackHistoricalStargazerHeader` saves the latest historical-info and deletes the oldest
// heights that are below pruning height.
func (k Keeper) TrackHistoricalStargazerHeader(ctx sdk.Context, header *types.StargazerHeader) {
	// Prune the store to ensure we only maintain the last numHistoricalBlocks.
	// In most cases, this will involve removing a single block from the store.
	// In the rare scenario when the historical blocks gets reduced to a lower value k'
	// from the original value k. k - k' blocks must be deleted from the store.
	// Since the entries to be deleted are always in a continuous range, we can iterate
	// over the historical entries starting from the most recent version to be pruned
	// and then return at the first empty entry.
	for i := ctx.BlockHeight() - numHistoricalBlocks; i >= 0; i-- {
		toPrune, found := k.GetStargazerHeader(ctx, i)
		if found {
			if err := k.PruneStargazerHeader(ctx, toPrune); err != nil {
				panic(err)
			}
		} else {
			break
		}
	}
	if err := k.SetStargazerHeader(ctx, header); err != nil {
		panic(err)
	}
}

// `GetStargazerBlock` returns the block from the store at the height specified in the context.
func (k *Keeper) GetStargazerHeader(ctx sdk.Context, height int64) (*types.StargazerHeader, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), key.SGHeaderPrefix)
	bz := store.Get(big.NewInt(height).Bytes())
	if bz == nil {
		return nil, false
	}

	// Unmarshal the retrieved header.
	header := new(types.StargazerHeader)
	if err := header.UnmarshalBinary(bz); err != nil {
		return nil, false
	}
	return header, true
}

// `SetStargazerHeader` saves a block to the store.
func (k *Keeper) SetStargazerHeader(ctx sdk.Context, header *types.StargazerHeader) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), key.SGHeaderPrefix)
	bz, err := header.MarshalBinary()
	if err != nil {
		return err
	}
	// Store the full block at the block key. (Overrides the old spot on the tree.)
	store.Set(header.Number.Bytes(), bz)

	// // Store a mapping of block hashes to block heights. (Grows over time)
	// store.Set(key.BlockHashToHeight(block.Hash()), sdk.Uint64ToBigEndian(block.Number.Uint64()))
	return nil
}

// `PruneStargazerHeader` prunes a stargazer block from the store.
func (k *Keeper) PruneStargazerHeader(ctx sdk.Context, header *types.StargazerHeader) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), key.SGHeaderPrefix)
	store.Delete(header.Number.Bytes())
	// Notably, we don't delete the store key mapping hash to height as we want this
	// to persist at the application layer in order to query by hash. (TODO? Tendermint?)
	return nil
}
