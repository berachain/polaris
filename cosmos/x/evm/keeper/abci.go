// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlock runs on the Cosmos-SDK lifecycle EndBlock() during ABCI Finalize.
func (k *Keeper) EndBlock(ctx context.Context) error {
	// Verify that the EVM block was written.
	blockNum := uint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	newHead := k.chain.GetBlockByNumber(blockNum)
	if newHead == nil {
		return fmt.Errorf(
			"evm block %d failed to process", blockNum,
		)
	} else if newHead.NumberU64() != blockNum {
		return fmt.Errorf(
			"evm block [%d] does not match comet block [%d]", newHead.NumberU64(), blockNum,
		)
	}

	// Set the finalized eth block once we know it has been finalized successfully by Cosmos.
	return k.chain.SetFinalizedBlock()
}

// PrepareCheckState runs on the Cosmos-SDK lifecycle PrepareCheckState() during ABCI Commit.
func (k *Keeper) PrepareCheckState(ctx context.Context) error {
	k.spf.SetLatestQueryContext(ctx)
	k.chain.PrimePlugins(ctx)
	return nil
}
