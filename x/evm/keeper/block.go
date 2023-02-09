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
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/x/evm/key"

	storeutils "github.com/berachain/stargazer/store/utils"
)

// ===========================================================================
// Stargazer Block
// ===========================================================================

// `MustStoreStargazerBlock` saves a block to the store.
func (k *Keeper) SetStargazerBlockForCurrentHeight(
	ctx sdk.Context,
	block *types.StargazerBlock,
) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := block.MarshalBinary()
	if err != nil {
		return err
	}
	// Store the full block at the block key. (Overrides the old spot on the tree.)
	store.Set(key.Block, bz)
	// Store the number of transactions in the block. (Overrides the old spot on the tree.)
	store.Set(key.BlockNumTx, sdk.Uint64ToBigEndian(uint64(len(block.Transactions))))

	// Store a mapping of block hashes to block heights. (Grows over time)
	store.Set(key.BlockHashToHeight(block.Hash()), sdk.Uint64ToBigEndian(block.Number.Uint64()))
	return nil
}

// `GetStargazerBlock` returns the block from the store at the height specified in the context.
func (k *Keeper) GetStargazerBlockAtHeight(
	ctx sdk.Context,
	height uint64,
) (*types.StargazerBlock, error) {
	bz := storeutils.KVStoreReaderAtBlockHeight(ctx, k.storeKey, int64(height)).Get(key.Block)
	if bz == nil {
		return nil, ErrBlockNotFound
	}

	// Unmarshal the retrieved block.
	block := new(types.StargazerBlock)
	if err := block.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return block, nil
}

// `GetStargazerBlockByHash` returns the block from the store with a given hash.
func (k *Keeper) GetStargazerBlockByHash(
	ctx sdk.Context,
	hash common.Hash,
) (*types.StargazerBlock, error) {
	bz := ctx.KVStore(k.storeKey).Get(key.BlockHashToHeight(hash))
	if bz == nil {
		return nil, ErrBlockNotFound
	}
	return k.GetStargazerBlockAtHeight(ctx, sdk.BigEndianToUint64(bz))
}

// ===========================================================================
// Transactions
// ===========================================================================

// `GetStargazerBlockTransactionCountByNumber` returns the number of transactions in a block from a block
// matching the given block number.
func (k *Keeper) GetStargazerBlockTransactionCountByNumber(ctx sdk.Context, number uint64) uint64 {
	store := storeutils.KVStoreReaderAtBlockHeight(ctx, k.storeKey, int64(number))
	return sdk.BigEndianToUint64(store.Get(key.BlockNumTx))
}

// `GetBlockTransactionCountByHash` returns the number of transactions in a block from a block
// matching the given block hash.
func (k *Keeper) GetStargazerBlockTransactionCountByHash(ctx sdk.Context, hash common.Hash) uint64 {
	bz := ctx.KVStore(k.storeKey).Get(key.BlockHashToHeight(hash))
	if bz == nil {
		return 0
	}
	// Now that we have recovered the height from the block hash, we can go and query using it.
	return k.GetStargazerBlockTransactionCountByNumber(ctx, sdk.BigEndianToUint64(bz))
}
