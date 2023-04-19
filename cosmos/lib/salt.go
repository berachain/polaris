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

package lib

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"github.com/holiman/uint256"

	storetypes "cosmossdk.io/store/types"
)

const globalNonce = `globalNonce`

var globalNonceKey = []byte(globalNonce)

// UniqueDeterministicSalt returns a unique and deterministic salt for the given input bytes. Uses
// sha256 to hash the input bytes and the global nonce and returns the salt as a *uint256.Int.
func UniqueDeterminsticSalt(nonceStore storetypes.BasicKVStore, input []byte) *uint256.Int {
	// create the sha256 hash and write the input bytes
	h := sha256.New()
	h.Write(input)

	// get the global nonce from the nonce store
	var globalNonce uint64
	if nonceBz := nonceStore.Get(globalNonceKey); nonceBz != nil {
		if err := binary.Read(
			bytes.NewReader(nonceBz), binary.LittleEndian, &globalNonce,
		); err != nil {
			panic(err)
		}
	}

	// write the nonce to the hash
	if err := binary.Write(h, binary.BigEndian, globalNonce); err != nil {
		panic(err)
	}

	// increment the global nonce and update the nonce store
	globalNonce++
	nonceBuf := new(bytes.Buffer)
	if err := binary.Write(nonceBuf, binary.LittleEndian, globalNonce); err != nil {
		panic(err)
	}
	nonceStore.Set(globalNonceKey, nonceBuf.Bytes())

	return new(uint256.Int).SetBytes(h.Sum(nil))
}
