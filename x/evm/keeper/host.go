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
	"math"

	errorsmod "cosmossdk.io/errors"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Compile-time assertion to ensure `Host` adheres to `core.Host`.
var _ core.Host = (*Host)(nil)

// `Host` is the Cosmos host for Ethereum.
type Host struct {
	// `sk` is used to access the staking keeper.
	sk StakingKeeper
}

func NewHost(sk StakingKeeper) *Host {
	return &Host{
		sk: sk,
	}
}

// `Logger` returns a module-specific logger.
func (h *Host) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "evm")
}

// `GetCoinbase` implements `core.Host`.
func (h *Host) GetCoinbase(gctx context.Context) (common.Address, error) {
	ctx := sdk.UnwrapSDKContext(gctx)
	// TODO: add redundancy here, incrase BlockHeader().ProposerAddress is not found, we want
	// to make sure that life is gucci as sometimes it doesn't matter.
	validator, found := h.sk.GetValidatorByConsAddr(ctx, ctx.BlockHeader().ProposerAddress)
	if !found {
		return common.Address{}, errorsmod.Wrapf(
			stakingtypes.ErrNoValidatorFound,
			"failed to retrieve validator operator from block proposer address %s",
			ctx.BlockHeader().ProposerAddress,
		)
	}

	return common.BytesToAddress(validator.GetOperator()), nil
}

// `GasMeter` implements `core.Host`.
func (h *Host) GasMeter(ctx context.Context) core.StargazerGasMeter {
	return sdk.UnwrapSDKContext(ctx).GasMeter()
}

// `GetBlockHashFunc` implements `core.Host`.
func (h *Host) GetBlockHashFunc(gctx context.Context) vm.GetHashFunc {
	ctx := sdk.UnwrapSDKContext(gctx)
	return func(height uint64) common.Hash {
		blockHeight := uint64(ctx.BlockHeight())
		switch {
		case blockHeight == height:
			// Case 1: The requested height matches the one from the context so we can
			// retrieve the header hash directly from the context.
			// Note: The headerHash is only set at begin block, it will be nil in case of a
			// query context
			headerHash := ctx.HeaderHash()
			if len(headerHash) != 0 {
				return common.BytesToHash(headerHash)
			}

			// only recompute the hash if not set (eg: checkTxState)
			contextBlockHeader := ctx.BlockHeader()
			header, err := tmtypes.HeaderFromProto(&contextBlockHeader)
			if err != nil {
				h.Logger(ctx).Error("failed to cast tendermint header from proto", "error", err)
				return common.Hash{}
			}

			return common.BytesToHash(header.Hash())

		case blockHeight > height:
			// Case 2: if the chain is not the current height we need to retrieve the hash from
			// the store for the current chain epoch. This only applies if the current height is
			// greater than the requested height.

			// If the requested height is greater than the max uint64 value, we return an empty
			// hash.
			if height > uint64(math.MaxInt64) {
				return common.Hash{}
			}

			histInfo, found := h.sk.GetHistoricalInfo(ctx, int64(height))
			if !found {
				h.Logger(ctx).Debug("historical info not found", "height", height)
				return common.Hash{}
			}

			header, err := tmtypes.HeaderFromProto(&histInfo.Header)
			if err != nil {
				h.Logger(ctx).Error("failed to cast tendermint header from proto", "error", err)
				return common.Hash{}
			}

			return common.BytesToHash(header.Hash())
		default:
			// Case 3: heights greater than the current one returns an empty hash.
			return common.Hash{}
		}
	}
}
