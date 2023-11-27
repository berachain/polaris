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

package types

import (
	"math/big"
	"unsafe"

	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

// DeriveReceiptsFromBlock is a helper function for deriving receipts from a block.
func DeriveReceiptsFromBlock(
	chainConfig *params.ChainConfig, receipts ethtypes.Receipts, block *ethtypes.Block,
) (ethtypes.Receipts, error) {
	// calculate the blobGasPrice according to the excess blob gas.
	var blobGasPrice = new(big.Int)
	if chainConfig.IsCancun(block.Number(), block.Time()) {
		blobGasPrice = eip4844.CalcBlobFee(*block.ExcessBlobGas())
	}

	// Derive receipts from block.
	if err := receipts.DeriveFields(
		chainConfig, block.Hash(), block.Number().Uint64(), block.Time(),
		block.BaseFee(), blobGasPrice, block.Transactions(),
	); err != nil {
		return nil, err
	}
	return receipts, nil
}

// MarshalReceipts marshals `Receipts`, as type `[]*ReceiptForStorage`, to bytes using rlp
// encoding.
func MarshalReceipts(receipts ethtypes.Receipts) ([]byte, error) {
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	receiptsForStorage := *(*[]*ethtypes.ReceiptForStorage)(unsafe.Pointer(&receipts))

	bz, err := rlp.EncodeToBytes(receiptsForStorage)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// UnmarshalReceipts unmarshals receipts from bytes to `[]*ReceiptForStorage` to `Receipts` using
// rlp decoding.
func UnmarshalReceipts(bz []byte) (ethtypes.Receipts, error) {
	var receiptsForStorage []*ethtypes.ReceiptForStorage
	if err := rlp.DecodeBytes(bz, &receiptsForStorage); err != nil {
		return nil, err
	}
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	return *(*ethtypes.Receipts)(unsafe.Pointer(&receiptsForStorage)), nil
}
