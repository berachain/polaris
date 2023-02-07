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

// `StargazerHeaderAtHeight` returns the stargazer header at the given height.
func (k *Keeper) StargazerHeaderAtHeight(ctx context.Context, height uint64) *types.StargazerHeader {
	sCtx := sdk.UnwrapSDKContext(ctx)
	if uint64(sCtx.BlockHeight()) == height {
		// If the current block height is the same as the requested height, then we assume that the
		// block has not been written to the store yet. In this case, we build and return a header
		// from the sdk.Context.
		return k.EthHeaderFromSdkContext(sCtx, types.Bloom{}, nil)
	} else if uint64(sCtx.BlockHeight()) < height {
		// If the current block height is less than the requested height, then we assume that the
		// block has been written to the store. In this case, we return the header from the store.
		stargazerBlock, _ := k.GetStargazerBlockAtHeight(sCtx, height)
		return stargazerBlock.StargazerHeader
	}
	// If the current block height is greater than the requested height, then we can't really query can we.
	// In this case, we return an empty header.
	return &types.StargazerHeader{}
}

// `EthHeaderFromSdkContext` builds an ethereum style block header from an `sdk.Context`, `Bloom` and `baseFee`.
func (k *Keeper) EthHeaderFromCosmosContext(ctx sdk.Context, bloom types.Bloom, baseFee *big.Int) *types.StargazerHeader {
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
		k.BlockHashFromCosmosContext(ctx),
	)
}

// `BlockHashFromSdkContext` extracts the block hash from a Cosmos context.
func (k *Keeper) BlockHashFromCosmosContext(ctx sdk.Context) common.Hash {
	headerHash := ctx.HeaderHash()
	if len(headerHash) != 0 {
		return common.BytesToHash(headerHash)
	}

	// only recompute the hash if not set (eg: checkTxState)
	contextBlockHeader := ctx.BlockHeader()
	header, err := tmtypes.HeaderFromProto(&contextBlockHeader)
	if err != nil {
		k.Logger(ctx).Error("failed to cast comet header from proto", "error", err)
		return common.Hash{}
	}

	headerHash = header.Hash()
	return common.BytesToHash(headerHash)
}
