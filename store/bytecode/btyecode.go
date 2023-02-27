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

package bytecode

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/crypto"
)

var (
	byteCodePrefix = []byte{0x1}
	versionPrefix  = []byte{0x2}
)

// `StoreByteCode` stores the byte code of the given address.
func (s Store) StoreByteCode(addr common.Address, code []byte) {
	prefix.NewStore(s.Store, byteCodePrefix).Set(addr.Bytes(), code)
}

// `GetByteCode` returns the byte code of the given address, compares it with the given
// code hash, and returns the byte code if the code hash matches.
func (s Store) GetByteCode(addr common.Address, codeHash common.Hash) ([]byte, error) {
	code := prefix.NewStore(s.Store, byteCodePrefix).Get(addr.Bytes())
	if codeHash != crypto.Keccak256Hash(code) {
		return nil, ErrByteCodeDoesNotMatch
	}

	return code, nil
}

// `IterateByteCode` iterates over the byte code and calls the given callback function. Break the
// iteration if the callback function returns true.
func (s Store) IterateByteCode(start, end []byte, cb func(addr common.Address, code []byte) bool) {
	iter := prefix.NewStore(s.Store, byteCodePrefix).Iterator(start, end)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		if cb(common.BytesToAddress(iter.Key()), iter.Value()) {
			break
		}
	}
}

// `SetVersion` sets the version of the byte code store. The version is used for the store snapshots.
func (s Store) SetVersion(version int64) {
	prefix.NewStore(s.Store, versionPrefix).Set(versionPrefix, sdk.Uint64ToBigEndian(uint64(version)))
}

// `GetVersion` returns the version of the byte code store.
func (s Store) GetVersion() int64 {
	return int64(sdk.BigEndianToUint64(prefix.NewStore(s.Store, versionPrefix).Get(versionPrefix)))
}
