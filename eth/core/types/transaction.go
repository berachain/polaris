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
