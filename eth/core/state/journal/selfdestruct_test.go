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
	"github.com/berachain/polaris/eth/core/state/journal/mock"
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SelfDestructs", func() {
	var s *selfDestructs
	var a1 = common.HexToAddress("0x1")
	var a2 = common.HexToAddress("0x2")
	var a3 = common.HexToAddress("0x3")
	var a4 = common.HexToAddress("0x4")

	BeforeEach(func() {
		s = utils.MustGetAs[*selfDestructs](NewSelfDestructs(mock.NewSelfDestructsStatePluginMock()))
	})

	It("should have the correct registry key", func() {
		Expect(s.RegistryKey()).To(Equal(suicidesRegistryKey))
	})

	It("should work correctly in the scope of a tx", func() {
		Expect(s.GetSelfDestructs()).To(BeEmpty())

		s.Snapshot()
		s.SelfDestruct(a2)
		s.SelfDestruct(a1)
		Expect(s.HasSelfDestructed(a2)).To(BeFalse())
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())

		snap2 := s.Snapshot()
		s.SelfDestruct(a3)
		Expect(s.HasSelfDestructed(a3)).To(BeTrue())
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())
		Expect(s.GetSelfDestructs()).To(HaveLen(2))

		s.RevertToSnapshot(snap2)
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())
		Expect(s.HasSelfDestructed(a3)).To(BeFalse())
		Expect(s.GetSelfDestructs()).To(HaveLen(1))

		s.Finalize()
		Expect(s.lastSnapshot).To(Equal(-1))
		Expect(s.Size()).To(Equal(0))
	})

	It("should not suicide when snapshot is not called", func() {
		s.SelfDestruct(a1)
		Expect(s.HasSelfDestructed(a1)).To(BeFalse())
	})

	It("should clone correctly", func() {
		s.Snapshot()
		s.SelfDestruct(a1)
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())

		s.Snapshot()
		s.SelfDestruct(a3)
		Expect(s.HasSelfDestructed(a3)).To(BeTrue())
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())
		Expect(s.GetSelfDestructs()).To(HaveLen(2))

		s2 := utils.MustGetAs[*selfDestructs](s.Clone())
		Expect(s.HasSelfDestructed(a3)).To(BeTrue())
		Expect(s.HasSelfDestructed(a1)).To(BeTrue())
		Expect(s2.GetSelfDestructs()).To(HaveLen(2))

		s.Snapshot()
		s2.Snapshot()

		s2.SelfDestruct(a4)
		Expect(s.HasSelfDestructed(a4)).To(BeFalse())
	})
})
