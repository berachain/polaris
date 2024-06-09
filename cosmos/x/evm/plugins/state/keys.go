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
