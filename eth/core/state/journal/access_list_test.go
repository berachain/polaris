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

package journal

import (
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AccessList", func() {
	var (
		al *accessList
		a1 = common.BytesToAddress([]byte{1})
		a2 = common.BytesToAddress([]byte{2})
		s1 = common.BytesToHash([]byte{1})
		s2 = common.BytesToHash([]byte{2})
	)

	BeforeEach(func() {
		al = utils.MustGetAs[*accessList](NewAccesslist())
	})

	It("should have the correct registry key", func() {
		Expect(al.RegistryKey()).To(Equal("accessList"))
	})

	It("should support controllable access list operations", func() {
		al.AddAddressToAccessList(a1)
		Expect(al.AddressInAccessList(a1)).To(BeTrue())
		Expect(al.AddressInAccessList(a2)).To(BeFalse())
		al.Peek().DeleteAddress(a1)
		Expect(al.AddressInAccessList(a1)).To(BeFalse())

		al.AddSlotToAccessList(a1, s1)
		al.AddSlotToAccessList(a1, s2)

		id := al.Snapshot()
		al.AddSlotToAccessList(a2, s1)

		Expect(al.AddressInAccessList(a2)).To(BeTrue())

		al.RevertToSnapshot(id)
		Expect(al.AddressInAccessList(a2)).To(BeFalse())

		Expect(func() { al.Finalize() }).ToNot(Panic())
		Expect(al.Size()).To(Equal(1))
	})

	It("should clone correctly", func() {
		al.AddSlotToAccessList(a1, s1)
		al.AddSlotToAccessList(a1, s2)

		al2 := utils.MustGetAs[*accessList](al.Clone())
		Expect(al2.AddressInAccessList(a1)).To(BeTrue())
		Expect(al2.AddressInAccessList(a2)).To(BeFalse())

		al2.AddSlotToAccessList(a2, s1)
		Expect(al2.AddressInAccessList(a2)).To(BeTrue())
		Expect(al.AddressInAccessList(a2)).To(BeFalse())
	})
})
