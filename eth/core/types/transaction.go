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
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// TxLookupEntry is a positional metadata to help looking up a transaction by hash.
//
//go:generate rlpgen -type TxLookupEntry -out transaction.rlpgen.go -decoder
type TxLookupEntry struct {
	Tx        *ethtypes.Transaction
	TxIndex   uint64
	BlockNum  uint64
	BlockHash common.Hash
}

// UnmarshalBinary decodes a tx lookup entry from the Ethereum RLP format.
func (tle *TxLookupEntry) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, tle)
}

// MarshalBinary encodes the tx lookup enßtry into the Ethereum RLP format.
func (tle *TxLookupEntry) MarshalBinary() ([]byte, error) {
	bz, err := rlp.EncodeToBytes(tle)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
