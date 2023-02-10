// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/x/evm/key"
)

// ===========================================================================
// Stargazer Block Tracking
// ===========================================================================.
const entryNum = 256

// TrackHistoricalStargazerBlocks saves the latest historical-info and deletes the oldest
// heights that are below pruning height.
func (k Keeper) TrackHistoricalStargazerBlocks(ctx sdk.Context, block *types.StargazerBlock) {
	// Prune store to ensure we only have parameter-defined historical entries.
	// In most cases, this will involve removing a single historical entry.
	// In the rare scenario when the historical entries gets reduced to a lower value k'
	// from the original value k. k - k' entries must be deleted from the store.
	// Since the entries to be deleted are always in a continuous range, we can iterate
	// over the historical entries starting from the most recent version to be pruned
	// and then return at the first empty entry.
	for i := ctx.BlockHeight() - int64(entryNum); i >= 0; i-- {
		toPrune, found := k.GetStargazerBlockAtHeight(ctx, uint64(i))
		if found {
			if err := k.PruneStargazerBlock(ctx, toPrune); err != nil {
				panic(err)
			}
		} else {
			break
		}
	}
	if err := k.SetLatestStargazerBlock(ctx, block); err != nil {
		panic(err)
	}
}

// ===========================================================================
// Stargazer Block By Height
// ===========================================================================

// `GetStargazerBlock` returns the block from the store at the height specified in the context.
func (k *Keeper) GetStargazerBlockAtHeight(
	ctx sdk.Context,
	height uint64,
) (*types.StargazerBlock, bool) {
	bz := ctx.KVStore(k.storeKey).Get(key.BlockAtHeight(height))
	if bz == nil {
		return nil, false
	}

	// Unmarshal the retrieved block.
	block := new(types.StargazerBlock)
	if err := block.UnmarshalBinary(bz); err != nil {
		return nil, false
	}
	return block, true
}

// `MustStoreStargazerBlock` saves a block to the store.
func (k *Keeper) SetLatestStargazerBlock(
	ctx sdk.Context,
	block *types.StargazerBlock,
) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := block.MarshalBinary()
	if err != nil {
		return err
	}
	// Store the full block at the block key. (Overrides the old spot on the tree.)
	store.Set(key.BlockAtHeight(block.Number.Uint64()), bz)

	// // Store a mapping of block hashes to block heights. (Grows over time)
	// store.Set(key.BlockHashToHeight(block.Hash()), sdk.Uint64ToBigEndian(block.Number.Uint64()))
	return nil
}

// `PruneStargazerBlock` prunes a stargazer block from the store.
func (k *Keeper) PruneStargazerBlock(
	ctx sdk.Context,
	block *types.StargazerBlock,
) error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(key.BlockAtHeight(block.Number.Uint64()))
	// Notably, we don't delete the store key mapping hash to height as we want this
	// to persist at the application layer in order to query by hash. (TODO? Tendermint?)
	return nil
}

// ===========================================================================
// Stargazer Block By Hash
// ===========================================================================

// // `GetStargazerBlockByHash` returns the block from the store with a given hash.
// func (k *Keeper) GetStargazerBlockByHash(
// 	ctx sdk.Context,
// 	hash common.Hash,
// ) (*types.StargazerBlock, bool) {
// 	bz := ctx.KVStore(k.storeKey).Get(key.BlockHashToHeight(hash))
// 	if bz == nil {
// 		return nil, false
// 	}
// 	return k.GetStargazerBlockAtHeight(ctx, sdk.BigEndianToUint64(bz))
// }

// ===========================================================================
// Transactions
// ===========================================================================

// `GetStargazerBlockTransactionCountByNumber` returns the number of transactions in a block from a block
// matching the given block number.
func (k *Keeper) GetStargazerBlockTransactionCountByNumber(ctx sdk.Context, number uint64) uint64 {
	// store := storeutils.KVStoreReaderAtBlockHeight(ctx, k.storeKey, int64(number))
	block, found := k.GetStargazerBlockAtHeight(ctx, number)
	if !found {
		return 0
	}

	return uint64(len(block.Transactions))
}

// // `GetBlockTransactionCountByHash` returns the number of transactions in a block from a block
// // matching the given block hash.
// func (k *Keeper) GetStargazerBlockTransactionCountByHash(ctx sdk.Context, hash common.Hash) uint64 {
// 	bz := ctx.KVStore(k.storeKey).Get(key.BlockHashToHeight(hash))
// 	if bz == nil {
// 		return 0
// 	}
// 	// Now that we have recovered the height from the block hash, we can go and query using it.
// 	return k.GetStargazerBlockTransactionCountByNumber(ctx, sdk.BigEndianToUint64(bz))
// }
