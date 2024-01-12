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

package state

import (
	types "github.com/berachain/polaris/cosmos/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
)

// NOTE: we use copy to build keys for max performance: https://github.com/golang/go/issues/55905

func BalanceKeyFor(address common.Address) []byte {
	bz := make([]byte, 1+common.AddressLength)
	copy(bz, []byte{types.BalanceKeyPrefix})
	copy(bz[1:], address[:])
	return bz
}

// StorageKeyFor returns a prefix to iterate over a given account storage (multiple slots).
func StorageKeyFor(address common.Address) []byte {
	bz := make([]byte, 1+common.AddressLength)
	copy(bz, []byte{types.StorageKeyPrefix})
	copy(bz[1:], address[:])
	return bz
}

// AddressFromStorageKey returns the address from a storage key.
func AddressFromStorageKey(key []byte) common.Address {
	return common.BytesToAddress(key[1:])
}

// SlotKeyFor defines the full key under which an account storage slot is stored.
func SlotKeyFor(address common.Address, slot common.Hash) []byte {
	bz := make([]byte, 1+common.AddressLength+common.HashLength)
	copy(bz, []byte{types.StorageKeyPrefix})
	copy(bz[1:], address[:])
	copy(bz[1+common.AddressLength:], slot[:])
	return bz
}

// SlotFromSlotKeyFor returns the slot from a slot key.
func SlotFromSlotKey(key []byte) common.Hash {
	return common.BytesToHash(key[1+common.AddressLength:])
}

// AddressFromSlotKey returns the address from a slot key.
func AddressFromSlotKey(key []byte) common.Address {
	return common.BytesToAddress(key[1 : 1+common.AddressLength])
}

// CodeHashKeyFor defines the full key under which an addresses codehash is stored.
func CodeHashKeyFor(address common.Address) []byte {
	bz := make([]byte, 1+common.AddressLength)
	copy(bz, []byte{types.CodeHashKeyPrefix})
	copy(bz[1:], address[:])
	return bz
}

// CodeKeyFor defines the full key under which an addresses code is stored.
func CodeKeyFor(codeHash common.Hash) []byte {
	bz := make([]byte, 1+common.HashLength)
	copy(bz, []byte{types.CodeKeyPrefix})
	copy(bz[1:], codeHash[:])
	return bz
}

// AddressFromCodeHashKey returns the address from a code hash key.
func AddressFromCodeHashKey(key []byte) common.Address {
	return common.BytesToAddress(key[1:])
}

// AddressFromBalanceKey returns the address from a balance key.
func AddressFromBalanceKey(key []byte) common.Address {
	return common.BytesToAddress(key[1:])
}
