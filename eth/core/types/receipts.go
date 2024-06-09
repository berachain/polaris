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
