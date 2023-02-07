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

	"cosmossdk.io/api/tendermint/abci"
	"github.com/berachain/stargazer/eth/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `BeginBlocker` is called during the BeginBlock processing of the ABCI lifecycle.
func (k *Keeper) BeginBlocker(ctx context.Context, req *abci.RequestBeginBlock) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("BeginBlocker")
	k.stateProcessor.Prepare(ctx, uint64(sCtx.BlockHeight()))
}

// `ProcessTransaction` is called during the DeliverTx processing of the ABCI lifecycle.
func (k *Keeper) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("ProcessTransaction")
	receipt, err := k.stateProcessor.ProcessTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	k.Logger(sCtx).Info("ProcessTransaction done")

	// Store the receipt in the state
	k.SetReceipt(sCtx, receipt)
	return receipt, err
}

// `EndBlocker` is called during the EndBlock processing of the ABCI lifecycle.
func (k *Keeper) EndBlocker(ctx context.Context, req *abci.RequestEndBlock) []abci.ValidatorUpdate {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("EndBlocker")
	reciepts, bloom, err := k.stateProcessor.Finalize(ctx, uint64(sCtx.BlockHeight()))
	if err != nil {
		panic(err)
	}
	// TODO: Store receipts and/or logs and/or blocks.
	_ = reciepts
	k.SetBlockBloom(sCtx, bloom)
	return []abci.ValidatorUpdate{}
}
