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

package key

var (
	// `SGHeaderPrefix` is the prefix for storing headers.
	SGHeaderPrefix = []byte("block")

	// receiptKey = []byte("receipt")
	// hashKey    = []byte("hash").
)

// func BlockAtHeight(height uint64) []byte {
// 	return append(blockKey, sdk.Uint64ToBigEndian(height)...)
// }

// `HashToTxIndex` returns the key for a receipt lookup.
// func HashToTxIndex(h []byte) []byte {
// 	return append(hashKey, h...)
// }

// // `TxIndexToReciept` returns the key for the receipt lookup for a given block.
// func TxIndexToReciept(txIndex uint64) []byte {
// 	return append(receiptKey, sdk.Uint64ToBigEndian(txIndex)...)
// }
