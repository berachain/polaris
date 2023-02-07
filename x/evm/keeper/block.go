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
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/x/evm/storage"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	store.Set(storage.BlockKey(), bz)
	return nil
}

// `GetStargazerBlock` returns the block from the store at the height specified in the context.
func (k *Keeper) GetStargazerBlockAtHeight(
	ctx sdk.Context,
	height uint64,
) (*types.StargazerBlock, error) {
	// Retrieve multi-store at the given height.
	cms, err := ctx.MultiStore().CacheMultiStoreWithVersion(int64(height))
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the store at the given height.
	store := ctx.WithMultiStore(cms).KVStore(k.storeKey)
	bz := store.Get(storage.BlockKey())
	if bz == nil {
		return nil, ErrBlockNotFound
	}

	// Unmarshal the retrieved block.
	block := new(types.StargazerBlock)
	if err = block.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return block, nil
}
