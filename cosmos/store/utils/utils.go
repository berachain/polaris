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

package utils

// // KVStoreReader is a subset of the `KVStore` interface that only exposes read
// // methods.
// type KVStoreReader interface {
// 	// Get returns nil if key doesn't exist. Panics on nil key.
// 	Get(key []byte) []byte

// 	// Has checks if a key exists. Panics on nil key.
// 	Has(key []byte) bool
// }

// // KVStoreReaderAtBlockHeight returns a KVStoreReader at a given height. If the height is greater
// // than or equal to the current height, the reader will be at the latest height. We return the store
// // with the modified height as a `KVStoreReader` since it does not make any sense to return a `KVStore`
// // since we cannot update historical versions of the tree.
// func KVStoreReaderAtBlockHeight(ctx sdk.Context, storeKey storetypes.StoreKey, height int64) KVStoreReader {
// 	if height >= ctx.BlockHeight() {
// 		return ctx.KVStore(storeKey)
// 	}
// 	fmt.Println("KVStoreReaderAtBlockHeight", height, ctx.BlockHeight())
// 	cms, ok := ctx.MultiStore().(storetypes.CommitMultiStore)
// 	if !ok {
// 		panic("REE")
// 	}
// 	cms.LoadVersion(height)
// 	return ctx.WithMultiStore(cms).KVStore(storeKey)
// }
