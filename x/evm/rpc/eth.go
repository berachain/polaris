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

package rpc

// `EthReaderBackend` is the backend for the `eth` namespace of the JSON-RPC API.
// It is only able to retrieve information about the current state of the chain by
// number. For querying data by hash, one must determine the block number that the
// data is stored at and then query by number.
// type EthReaderBackend struct {
// 	k keeper.Keeper
// }

// ==============================================================================
// EthReaderBackend
// ==============================================================================

// // `BlockNumber` implements the `eth_blockNumber` JSON-RPC method.
// func (eb *EthReaderBackend) BlockNumber(ctx sdk.Context) uint64 {
// 	return uint64(ctx.BlockHeight())
// }

// // `GetBlockByNumber` is used to implement the `eth_getBlockByNumber` JSON-RPCÍ.
// func (eb *EthReaderBackend) GetBlockByNumber(
// 	ctx sdk.Context, number uint64, fullTx bool,
// ) (*types.StargazerBlock, error) {
// 	block, found := eb.k.GetStargazerBlockAtHeight(ctx, number)
// 	if !found {
// 		return nil, errors.New("no block found")
// 	}
// 	return block, nil
// }

// // `GetStargazerBlockTransactionCountByNumber` returns the number of transactions in a block from a block
// // matching the given block number.
// func (eb *EthReaderBackend) BlockTransactionCountByNumber(ctx sdk.Context, number uint64) uint64 {
// 	// store := storeutils.KVStoreReaderAtBlockHeight(ctx, k.storeKey, int64(number))
// 	block, found := eb.k.GetStargazerBlockAtHeight(ctx, number)
// 	if !found {
// 		return 0
// 	}

// 	return uint64(block.TxIndex())
// }
