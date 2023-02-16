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

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	cbft "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: change this.
const bf = int64(1e9)

type StargazerHeaderGetter interface {
	GetStargazerHeader(ctx sdk.Context, height int64) (*coretypes.StargazerHeader, bool)
}

type plugin struct {
	ctx sdk.Context
	shg StargazerHeaderGetter
}

func NewPluginFrom(ctx sdk.Context, shg StargazerHeaderGetter) core.BlockPlugin {
	return &plugin{
		ctx: ctx,
		shg: shg,
	}
}

func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}

// `BaseFee` returns the base fee for the current block.
// TODO: implement properly with DynamicFee Module of some kind.
func (p *plugin) BaseFee() uint64 {
	return uint64(bf)
}

func (p *plugin) GetStargazerHeaderAtHeight(height int64) *coretypes.StargazerHeader {
	// If the current block height is the same as the requested height, then we assume that the
	// block has not been written to the store yet. In this case, we build and return a header
	// from the sdk.Context.
	if p.ctx.BlockHeight() == height {
		return p.getStargazerHeaderFromCosmosContext()
	}

	// If the current block height is less than (or technically also greater than) the requested
	// height, then we assume that the block has been written to the store. In this case, we
	// return the header from the store.
	if header, found := p.shg.GetStargazerHeader(p.ctx, height); found {
		return header
	}

	return &coretypes.StargazerHeader{}
}

// `getStargazerHeaderFromCosmosContext` builds an ethereum style block header from an
// `sdk.Context`, `Bloom` and `baseFee`.
func (p *plugin) getStargazerHeaderFromCosmosContext() *coretypes.StargazerHeader {
	cometHeader := p.ctx.BlockHeader()

	// We retrieve the `TxHash` from the `DataHash` field of the `sdk.Context` opposed to deriving it
	// from solely the ethereum transaction information.
	txHash := coretypes.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	return coretypes.NewStargazerHeader(
		&coretypes.Header{
			// `ParentHash` is set to the hash of the previous block.
			ParentHash: common.BytesToHash(cometHeader.LastBlockId.Hash),
			// `UncleHash` is set empty as CometBFT does not have uncles.
			UncleHash: coretypes.EmptyUncleHash,
			// TODO: Use staking keeper to get the operator address.
			Coinbase: common.BytesToAddress(cometHeader.ProposerAddress),
			// `Root` is set to the hash of the state after the transactions are applied.
			Root: common.BytesToHash(cometHeader.AppHash),
			// `TxHash` is set to the hash of the transactions in the block. We take the
			// `DataHash` from the `sdk.Context` opposed to using DeriveSha on the StargazerBlock,
			// in order to include non-evm transactions block hash.
			TxHash: txHash,
			// We simply map the cosmos "BlockHeight" to the ethereum "BlockNumber".
			Number: big.NewInt(cometHeader.Height),
			// `GasLimit` is set to the block gas limit.
			GasLimit: p.ctx.BlockGasMeter().Limit(),
			// `Time` is set to the block timestamp.
			Time: uint64(cometHeader.Time.UTC().Unix()),
			// `BaseFee` is set to the block base fee.
			BaseFee: big.NewInt(int64(p.BaseFee())),
			// `ReceiptHash` set to empty. It is filled during `Finalize` in the StateProcessor.
			ReceiptHash: common.Hash{},
			// `Bloom` is set to empty. It is filled during `Finalize` in the StateProcessor.
			Bloom: coretypes.Bloom{},
			// `GasUsed` is set to 0. It is filled during `Finalize` in the StateProcessor.
			GasUsed: 0,
			// `Difficulty` is set to 0 as it is only used in PoW consensus.
			Difficulty: big.NewInt(0),
			// `MixDigest` is set empty as it is only used in PoW consensus.
			MixDigest: common.Hash{},
			// `Nonce` is set empty as it is only used in PoW consensus.
			Nonce: coretypes.BlockNonce{},
			// `Extra` is unused in Stargazer.
			Extra: []byte(nil),
		},
		blockHashFromCosmosContext(p.ctx),
	)
}

// `blockHashFromSdkContext` extracts the block hash from a Cosmos context.
func blockHashFromCosmosContext(ctx sdk.Context) common.Hash {
	headerHash := ctx.HeaderHash()
	if len(headerHash) != 0 {
		return common.BytesToHash(headerHash)
	}

	// only recompute the hash if not set (eg: checkTxState)
	contextBlockHeader := ctx.BlockHeader()
	header, err := cbft.HeaderFromProto(&contextBlockHeader)
	if err != nil {
		return common.Hash{}
	}

	return common.BytesToHash(header.Hash())
}
