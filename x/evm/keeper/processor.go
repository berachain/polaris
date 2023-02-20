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

	coretypes "github.com/berachain/stargazer/eth/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `BeginBlocker` is called during the BeginBlock processing of the ABCI lifecycle.
func (k *Keeper) BeginBlocker(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("BeginBlocker")
	k.stargazer.Prepare(ctx, sCtx.BlockHeight())
}

// `ProcessTransaction` is called during the DeliverTx processing of the ABCI lifecycle.
func (k *Keeper) ProcessTransaction(ctx context.Context, tx *coretypes.Transaction) (*coretypes.Receipt, error) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("Begin ProcessTransaction()")

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
	// But we do have access to the ethereum tx hash.
	// We could expose a get txn by hash in our app side mempool that allows use to query the txn by hash.
	// Basically just have a cache of eth hashes in the mempool.
	// App-side mempool good project.

	// TODO: In theory, the TendermintTxHash is the Sha256 hash of the fully populated
	// EthereumMsgTx (after from and hash and stuff are filled in).
	// This should be doable at the application layer, and means that given an EthereumHash
	// we can calculate a TendermintHash. But not vice versa.
	k.Logger(sCtx).Info("End ProcessTransaction()")
	return receipt, err
}

// `EndBlocker` is called during the EndBlock processing of the ABCI lifecycle.
func (k *Keeper) EndBlocker(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sCtx).Info("EndBlocker")

	// Finalize the stargazer block and retrieve it from the processor.
	stargazerBlock, err := k.stargazer.Finalize(ctx)
	if err != nil {
		panic(err)
	}

	// Save the historical stargazer header.
	k.TrackHistoricalStargazerHeader(sCtx, stargazerBlock.StargazerHeader)

	// TODO: this is sketchy and needs to be refactored later.
	go k.UpdateOffChainStorage(sCtx, stargazerBlock)

	// do all the off chain storage
	// probably in a go routine
}
