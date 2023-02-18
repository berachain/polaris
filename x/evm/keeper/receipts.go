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

// // SetReceipt stores the receipt indexed by the tx index.
// func (k *Keeper) SetReceipts(ctx sdk.Context, receipt *types.Receipt) {
// 	rs := (coretypes.ReceiptForStorage)(*receipt)
// 	bz, err := rs.MarshalBinary()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// save a mapping: txHash to BlockNumber

// 	// save a mapping txHash

// 	// We need to store the receipt for the block numer + tx index for efficient iteration., but
// 	// we also need to allow for a way to lookup a receipt by hash.
// 	receiptKey := key.TxIndexToReciept(
// 		uint64(receipt.TransactionIndex),
// 	)

// 	// Store the receiptKey in the store with a key of the tx hash.
// 	ctx.KVStore(k.storeKey).Set(key.HashToTxIndex(receipt.TxHash.Bytes()), receiptKey)

// 	// Store the receipt indexed by tx index.
// 	ctx.KVStore(k.storeKey).Set(receiptKey, bz)
// }

// // `GetReceipt` gets the receipt indexed by the receipt hash.
// func (k *Keeper) GetReceipt(ctx sdk.Context, txIndex uint64) *types.Receipt {
// 	receiptKey := key.TxIndexToReciept(txIndex)
// 	bz := ctx.KVStore(k.storeKey).Get(receiptKey)
// 	if bz == nil {
// 		return nil
// 	}
// 	receipt := new(types.Receipt)
// 	if err := receipt.UnmarshalBinary(bz); err != nil {
// 		panic(err)
// 	}
// 	return receipt
// }

// // `GetReceiptByTxHash` gets the receipt indexed by the transaction hash.
// func (k *Keeper) GetReceiptByTxHash(ctx sdk.Context, txHash common.Hash) *types.Receipt {
// 	receiptKey := ctx.KVStore(k.storeKey).Get(key.HashToTxIndex(txHash.Bytes()))
// 	if receiptKey == nil {
// 		return nil
// 	}
// 	bz := ctx.KVStore(k.storeKey).Get(receiptKey)
// 	if bz == nil {
// 		return nil
// 	}
// 	receipt := new(types.Receipt)
// 	if err := receipt.UnmarshalBinary(bz); err != nil {
// 		panic(err)
// 	}
// 	return receipt
// }
