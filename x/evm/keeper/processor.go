// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
)

// `BeginBlocker` is called during the BeginBlock processing of the ABCI lifecycle.
func (k *Keeper) BeginBlocker(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("keeper.BeginBlocker")
	k.stargazer.Prepare(ctx, sCtx.BlockHeight())
}

// `ProcessTransaction` is called during the DeliverTx processing of the ABCI lifecycle.
func (k *Keeper) ProcessTransaction(ctx context.Context, tx *coretypes.Transaction) (*coretypes.Receipt, error) {
	// Process the transaction and return the receipt.
	receipt, err := k.stargazer.ProcessTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	// TODO: note if we get a Block Error out of gas here, we need the transaction to be included
	// in the block. This is because the transaction was included in the block, but something
	// happened to put it into a situation where it really should have, this will traditionally
	// cause the cosmos transaction to fail, which is correct, but not what we want here. What
	// we need to do, is edit the gas consumption to consume the remaining gas in the block,
	//  modifying the receipt, and return a failed EVM tx, but a successful cosmos tx.
	// TODO: determine if the above is actually correct.

	k.Logger(sdk.UnwrapSDKContext(ctx)).Info("End ProcessTransaction()")
	return receipt, err
}

// `EndBlocker` is called during the EndBlock processing of the ABCI lifecycle.
func (k *Keeper) EndBlocker(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)

	// Finalize the stargazer block and retrieve it from the processor.
	stargazerBlock, err := k.stargazer.Finalize(ctx)
	if err != nil {
		panic(err)
	}

	k.Logger(sCtx).Info("keeper.EndBlocker", "block header:", stargazerBlock.Header)

	// Save the historical stargazer header in the IAVL Tree.
	k.bp.ProcessHeader(sCtx, stargazerBlock.StargazerHeader)

	// TODO: this is sketchy and needs to be refactored later.
	// Save the block data to the off-chain storage.
	if k.offChainKv != nil {
		k.bp.UpdateOffChainStorage(sCtx, stargazerBlock)
	}
}
