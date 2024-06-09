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

var (
	alice  = common.Address{1}
	bob    = common.Address{2}
	key    = common.Hash{0x01}
	value  = common.Hash{0x02}
	value2 = common.Hash{0x03}
)

var _ = Describe("TransientStorage", func() {
	var ts *transientStorage

	BeforeEach(func() {
		ts = utils.MustGetAs[*transientStorage](NewTransientStorage())
	})

	It("should add without impacting previous state", func() {
		ts.SetTransientState(alice, key, value)
		ts.SetTransientState(bob, key, value)

		// manually ensure the first transient state is not overwritten
		Expect(ts.PeekAt(0).Get(alice, key)).To(Equal(value))
		Expect(ts.PeekAt(0).Get(bob, key)).To(Equal(common.Hash{}))

		// the current transient state should have all state changes
		Expect(ts.GetTransientState(alice, key)).To(Equal(value))
		Expect(ts.GetTransientState(bob, key)).To(Equal(value))
	})

	It("should have consistent gets and sets", func() {
		ts.SetTransientState(alice, key, value) // {alice:value}
		Expect(ts.GetTransientState(alice, key)).To(Equal(value))

		before := ts.Snapshot()
		ts.SetTransientState(alice, key, value2) // {alice:value2}
		Expect(ts.GetTransientState(alice, key)).To(Equal(value2))

		ts.SetTransientState(bob, key, value) // {alice:value2, bob: value}
		ts.RevertToSnapshot(before)           // {alice:value}
		Expect(ts.GetTransientState(alice, key)).To(Equal(value))
		Expect(ts.GetTransientState(bob, key)).To(Equal(common.Hash{}))
	})

	It("should correctly finalize", func() {
		ts.SetTransientState(alice, key, value)
		ts.Finalize()
		Expect(ts.Size()).To(Equal(0))
		Expect(func() { ts.Finalize() }).ToNot(Panic())
	})

	It("should correctly clone", func() {
		ts.SetTransientState(bob, key, value)
		Expect(ts.GetTransientState(alice, key)).To(Equal(common.Hash{}))
		Expect(ts.GetTransientState(bob, key)).To(Equal(value))

		ts2 := utils.MustGetAs[*transientStorage](ts.Clone())
		Expect(ts2.GetTransientState(alice, key)).To(Equal(common.Hash{}))
		Expect(ts2.GetTransientState(bob, key)).To(Equal(value))

		ts2.SetTransientState(alice, key, value2)
		Expect(ts2.GetTransientState(alice, key)).To(Equal(value2))
		Expect(ts.GetTransientState(alice, key)).To(Equal(common.Hash{}))
	})
})
