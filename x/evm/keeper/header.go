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
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"

	tmtypes "github.com/tendermint/tendermint/types"
)

// `StargazerHeaderAtHeight`.
func (k *Keeper) StargazerHeaderAtHeight(ctx context.Context, height uint64) *types.StargazerHeader {
	sCtx := sdk.UnwrapSDKContext(ctx)
	if sCtx.BlockHeight() == int64(height) {
		return k.EthHeaderFromSdkContext(sCtx, types.Bloom{}, nil)
	}
	stargazerBlock, _ := k.GetStargazerBlockAtHeight(sCtx, height)
	return stargazerBlock.StargazerHeader
}

// `EthHeaderFromSdkContext` builds an ethereum style block header from an `sdk.Context`, `Bloom` and `baseFee`.
func (k *Keeper) EthHeaderFromSdkContext(ctx sdk.Context, bloom types.Bloom, baseFee *big.Int) *types.StargazerHeader {
	cometHeader := ctx.BlockHeader()
	txHash := types.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	return types.NewStargazerHeader(
		&types.Header{
			ParentHash:  common.BytesToHash(cometHeader.LastBlockId.Hash),
			UncleHash:   types.EmptyUncleHash,
			Coinbase:    common.BytesToAddress(cometHeader.ProposerAddress),
			Root:        common.BytesToHash(cometHeader.AppHash),
			TxHash:      txHash,
			ReceiptHash: types.EmptyRootHash,
			Bloom:       bloom,
			Difficulty:  big.NewInt(0),
			Number:      big.NewInt(cometHeader.Height),
			GasLimit:    ctx.BlockGasMeter().Limit(),
			GasUsed:     ctx.BlockGasMeter().GasConsumed(),
			Time:        uint64(cometHeader.Time.UTC().Unix()),
			Extra:       []byte{},
			MixDigest:   common.Hash{},
			Nonce:       types.BlockNonce{},
			BaseFee:     baseFee,
		},
		k.BlockHashFromSdkContext(ctx),
	)
}

// // `StargazerHeaderAtHeight` returns the header at the given height.
// func (k Keeper) StargazerBlockHashAtHeight(gctx context.Context, height uint64) common.Hash {
// 	ctx := sdk.UnwrapSDKContext(gctx)
// 	if height > math.MaxInt64 {
// 		panic("height is greater than max int64")
// 	}
// 	h := int64(height)
// 	ctxHeight := ctx.BlockHeight()
// 	switch {
// 	case ctxHeight == h:
// 		return k.BlockHashFromSdkContext(ctx)
// 	case ctxHeight > h:
// 		return k.BlockHashFromHistoricalInfo(ctx, h)
// 	default:
// 		return common.Hash{}
// 	}
// }

// `HashFromSdkContext` extracts the block hash from the Cosmos context.
func (k *Keeper) BlockHashFromSdkContext(ctx sdk.Context) common.Hash {
	headerHash := ctx.HeaderHash()
	if len(headerHash) != 0 {
		return common.BytesToHash(headerHash)
	}

	// only recompute the hash if not set (eg: checkTxState)
	contextBlockHeader := ctx.BlockHeader()
	header, err := tmtypes.HeaderFromProto(&contextBlockHeader)
	if err != nil {
		k.Logger(ctx).Error("failed to cast tendermint header from proto", "error", err)
		return common.Hash{}
	}

	headerHash = header.Hash()
	return common.BytesToHash(headerHash)
}

// // `HashFromSdkContext` extracts the block has using historical information saved
// // in the staking keeper.
// func (k *Keeper) BlockHashFromHistoricalInfo(ctx sdk.Context, height int64) common.Hash {
// 	histInfo, found := k.stakingKeeper.GetHistoricalInfo(ctx, height)
// 	if !found {
// 		k.Logger(ctx).Debug("historical info not found", "height", height)
// 		return common.Hash{}
// 	}

// 	header, err := tmtypes.HeaderFromProto(&histInfo.Header)
// 	if err != nil {
// 		k.Logger(ctx).Error("failed to cast tendermint header from proto", "error", err)
// 		return common.Hash{}
// 	}

// 	return common.BytesToHash(header.Hash())
// }
