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

	"github.com/berachain/stargazer/eth/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// `BeginBlocker` is called during the BeginBlock processing of the ABCI lifecycle.
func (k *Keeper) BeginBlocker(ctx context.Context, req *abci.RequestBeginBlock) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("BeginBlocker")
	k.stateProcessor.Prepare(ctx, k.GetStargazerHeaderAtHeight(ctx, uint64(sCtx.BlockHeight())))
}

// `ProcessTransaction` is called during the DeliverTx processing of the ABCI lifecycle.
func (k *Keeper) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("Begin ProcessTransaction()")

	// Process the transaction and return the receipt.
	receipt, err := k.stateProcessor.ProcessTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	// TODO: note if we get a Block Error out of gas here, we need the transaction to be included
	// in the block. This is because the transaction was included in the block, but something
	// happened to put it into a situation where it really should have, this will traditionally
	// cause the cosmos transaction to fail, which is correct, but not what we want here. What
	// we need to do, is edit the gas consumption to consume the remaining gas in the block,
	//  modifying the receipt, and return a failed EVM tx, but a successful cosmos tx.

	// TODO: Need to emit event to create a map of TendermintHash EthereumTxHash mapping
	// TODO: BUT should we just yeet receipts into tendermint? (TMHash -> Receipt)
	// This would give us Tendermint Hash -> Receipt mapping.
	// https://github.com/evmos/ethermint/issues/1075
	// https://github.com/crypto-org-chain/cronos/issues/455
	// TODO: figure out how the tendermint indexer works.
	// 	Indexer DB: Key: ethereum_tx.ethereumTxHash/{ETH_HASH}/{res.Height}/{res.Index}, Value: tm hash.
	// Indexer DB: Key: tm hash, Value: abci.TxResult.
	// State DB: Key: abciResponsesKey:{height}, Value: tmstate.ABCIResponses.
	// TODO: We don't have access to the TM TxHash in the state machine?
	k.Logger(sCtx).Info("End ProcessTransaction()")
	return receipt, err
}

// `EndBlocker` is called during the EndBlock processing of the ABCI lifecycle.
func (k *Keeper) EndBlocker(ctx context.Context, req *abci.RequestEndBlock) []abci.ValidatorUpdate {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("EndBlocker")

	// Finalize the stargazer block and retrieve it from the processor.
	stargazerBlock, err := k.stateProcessor.Finalize(ctx)
	if err != nil {
		panic(err)
	}

	// Save the historical stargazer block.
	k.TrackHistoricalStargazerBlocks(sCtx, stargazerBlock)
	// TODO: should we just yeet stargazer blocks into tendermint event filter? (sans receipts)?

	return []abci.ValidatorUpdate{}
}
