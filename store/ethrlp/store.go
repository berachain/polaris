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

package ethrlp

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/common"
)

// `EthereumRlpEncoded` is an interface that should be used to work with all ethereum rlp encoded data.
type EthereumRlpEncoded interface {
	Hash() common.Hash
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

// `EthereumRlpStore` is a wrapper around the underlying store that allows to store and retrieve, implement `EthereumRlpEncoded` interface.
type EthereumRlpStore[T EthereumRlpEncoded] struct {
	underlying storetypes.KVStore
}

// `NewRlpEncodedStore` creates a new instance of `EthereumRlpStore` from provided underlying store and key prefix.
func NewRlpEncodedStore[T EthereumRlpEncoded](underlying storetypes.KVStore, keyPrefix []byte) *EthereumRlpStore[T] {
	return &EthereumRlpStore[T]{underlying: prefix.NewStore(underlying, keyPrefix)}
}

// `Set` stores the provided data in the underlying store.
func (rlps *EthereumRlpStore[T]) Set(data T) {
	bz, err := data.MarshalBinary()
	if err != nil {
		// TODO: operate in mode without offchain storage if this fails
		// ctx.Logger().Error("MarshalBinary for block. Failed to update offchain storagae", "err", err)
		return
	}
	rlps.underlying.Set(data.Hash().Bytes(), bz)
}

// `Get` retrieves the unmarshalled data from the underlying store.
func (rlps *EthereumRlpStore[T]) Get(key []byte) (*T, bool) {
	bz := rlps.underlying.Get(key)
	if bz == nil {
		return nil, false
	}
	var data T
	err := data.UnmarshalBinary(bz)
	if err != nil {
		return nil, false
	}
	return &data, true
}
