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

package block

// import (
// 	"context"
// 	"math/big"

// 	"github.com/berachain/stargazer/eth/core/types"
// 	storetypes "github.com/cosmos/cosmos-sdk/store/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// type Plugin struct {
// 	storeKey storetypes.StoreKey
// }

// // `BaseFee()` returns the base fee of the current block.
// func (p *Plugin) BaseFee(ctx context.Context) *big.Int {
// 	// sCtx := sdk.UnwrapSDKContext(ctx)
// 	return new(big.Int).SetInt64(10000)
// }

// // TrackHistoricalStargazerHeaders saves the latest historical-info and deletes the oldest
// // heights that are below pruning height
// func (p Plugin) TrackHistoricalStargazerHeaders(ctx sdk.Context, block *types.StargazerBlock) {
// 	entryNum := 256

// 	// Prune store to ensure we only have parameter-defined historical entries.
// 	// In most cases, this will involve removing a single historical entry.
// 	// In the rare scenario when the historical entries gets reduced to a lower value k'
// 	// from the original value k. k - k' entries must be deleted from the store.
// 	// Since the entries to be deleted are always in a continuous range, we can iterate
// 	// over the historical entries starting from the most recent version to be pruned
// 	// and then return at the first empty entry.
// 	for i := ctx.BlockHeight() - int64(entryNum); i >= 0; i-- {
// 		block, found := p.GetStargazerHeaderAtHeight(ctx, uint64(i))
// 		if found {
// 			p.PruneStargazerBlock(ctx, block)
// 		} else {
// 			break
// 		}
// 	}
// 	p.SetLatestStargazerBlock(ctx, block)
// }

// // `StargazerHeaderAtHeight` returns the stargazer header at the given height.
// func (p *Plugin) GetStargazerHeaderAtHeight(ctx context.Context, height uint64) (*types.StargazerHeader, bool) {
// 	sCtx := sdk.UnwrapSDKContext(ctx)
// 	if uint64(sCtx.BlockHeight()) == height {
// 		// If the current block height is the same as the requested height, then we assume that the
// 		// block has not been written to the store yet. In this case, we build and return a header
// 		// from the sdk.Context.
// 		return p.StargazerHeaderFromCosmosContext(sCtx, types.Bloom{}, k.BaseFee(ctx)), true
// 	} else if uint64(sCtx.BlockHeight()) < height {
// 		// If the current block height is less than the requested height, then we assume that the
// 		// block has been written to the store. In this case, we return the header from the store.
// 		bz := sdk.UnwrapSDKContext(ctx).KVStore(k.storeKey).Get(storage.BlockKey())
// 		if bz == nil {
// 			return nil, false
// 		}

// 		// Unmarshal the retrieved block
// 		header := new(types.StargazerHeader)
// 		if err := header.UnmarshalBinary(bz); err != nil {
// 			return nil, false
// 		}
// 		return header, true
// 	}
// 	// If the current block height is greater than the requested height, then we can't really query can we.
// 	// In this case, we return an empty header.
// 	return &types.StargazerHeader{}, false
// }
