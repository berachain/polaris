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
	"github.com/berachain/polaris/cosmos/x/evm/types"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StorageKeyFor", func() {
	It("returns a prefix to iterate over a given account storage", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		prefix := StorageKeyFor(address)
		Expect(prefix).To(HaveLen(1 + common.AddressLength))
		Expect(prefix[0]).To(Equal(types.StorageKeyPrefix))
		Expect(prefix[1:]).To(Equal(address.Bytes()))
	})
})

var _ = Describe("AddressFromStorageKey", func() {
	It("should return the address from a storage key", func() {
		addr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		prefix := StorageKeyFor(addr)

		addr2 := AddressFromStorageKey(prefix)
		Expect(addr2).To(Equal(addr))
	})
})

var _ = Describe("SlotKeyFor", func() {
	It("returns a storage key for a given account and storage slot", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		slot := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
		key := SlotKeyFor(address, slot)
		Expect(key).To(HaveLen(1 + common.AddressLength + common.HashLength))
		Expect(key[0]).To(Equal(types.StorageKeyPrefix))
		Expect(key[1 : 1+common.AddressLength]).To(Equal(address.Bytes()))
		Expect(key[1+common.AddressLength:]).To(Equal(slot.Bytes()))
	})
})

var _ = Describe("SlotFromSlotKey", func() {
	It("should return the slot from the key", func() {
		addr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		slot := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
		key := SlotKeyFor(addr, slot)

		addr2 := AddressFromSlotKey(key)
		slot2 := SlotFromSlotKey(key)
		Expect(addr2).To(Equal(addr))
		Expect(slot2).To(Equal(slot))
	})
})

var _ = Describe("CodeHashKeyFor or a given account", func() {
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	key := CodeHashKeyFor(address)
	Expect(key).To(HaveLen(1 + common.AddressLength))
	Expect(key[0]).To(Equal(types.CodeHashKeyPrefix))
	Expect(key[1:]).To(Equal(address.Bytes()))
})

var _ = Describe("AddressFromCodeHashKey", func() {
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	key := CodeHashKeyFor(address)

	address2 := AddressFromCodeHashKey(key)
	Expect(address2).To(Equal(address))
})

var _ = Describe("CodeKeyFor", func() {
	It("returns a code key for a given account", func() {
		address := common.HexToHash("0x1234567890abcdef1234567890abcdef12345678")
		key := CodeKeyFor(address)
		Expect(key).To(HaveLen(1 + common.HashLength))
		Expect(key[0]).To(Equal(types.CodeKeyPrefix))
		Expect(key[1:]).To(Equal(address.Bytes()))
	})
})
