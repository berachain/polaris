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
	"github.com/berachain/stargazer/eth/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// `StargazerHeader` represents a wrapped Ethereum header that allows for specifying a custom
// blockhash to make it compatible with a non-ethereum chain.
//
//go:generate rlpgen -type StargazerHeader -out header.rlpgen.go -decoder
type StargazerHeader struct {
	// `Header` is an embedded ethereum header.
	*Header
	// `hostHash` is the block hash on the host chain.
	hostHash common.Hash
}

// `NewEmptyStargazerHeader` returns an empty `StargazerHeader`.
func NewEmptyStargazerHeader() *StargazerHeader {
	return &StargazerHeader{Header: &Header{}}
}

// `NewStargazerHeader` returns a `StargazerHeader` with the given `header` and `hash`.
func NewStargazerHeader(header *Header, hash common.Hash) *StargazerHeader {
	return &StargazerHeader{Header: header, hostHash: hash}
}

// `UnmarshalBinary` decodes a block from the Ethereum RLP format.
func (h *StargazerHeader) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, h)
}

// `MarshalBinary` encodes the block into the Ethereum RLP format.
func (h *StargazerHeader) MarshalBinary() ([]byte, error) {
	bz, err := rlp.EncodeToBytes(h)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// `Author` returns the address of the original block producer.
func (h *StargazerHeader) Author() common.Address {
	return h.Coinbase
}

// `Hash` returns the block hash of the header, we override the geth implementation
// to use the hash of the host chain, as the implementing chain might want to use it's
// real block hash opposed to hashing the "fake" header.
func (h *StargazerHeader) Hash() common.Hash {
	// if h.hostHash == (common.Hash{}) {
	// 	h.hostHash = h.Header.Hash()
	// }
	// return h.hostHash
	return h.Header.Hash()
}

// `SetHash` sets the hash of the header.
func (h *StargazerHeader) SetHash(hash common.Hash) {
	h.hostHash = hash
}
