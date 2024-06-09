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
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logs", func() {
	var l *logs
	var thash = common.BytesToHash([]byte{1})
	var ti = uint(1)
	var bnum = uint64(2)
	var bhash = common.BytesToHash([]byte{2})
	var a1 = common.BytesToAddress([]byte{3})
	var a2 = common.BytesToAddress([]byte{4})
	var a3 = common.BytesToAddress([]byte{5})

	BeforeEach(func() {
		l = utils.MustGetAs[*logs](NewLogs())
		l.SetTxContext(thash, int(ti))
		Expect(l.Capacity()).To(Equal(32))
	})

	It("should have the correct registry key", func() {
		Expect(l.RegistryKey()).To(Equal("logs"))
	})

	When("adding logs", func() {
		BeforeEach(func() {
			l.AddLog(&ethtypes.Log{Address: a1})
			Expect(l.Size()).To(Equal(1))
			Expect(l.PeekAt(0).Address).To(Equal(a1))
			Expect(l.PeekAt(0).TxHash).To(Equal(thash))
			Expect(l.PeekAt(0).TxIndex).To(Equal(ti))
		})

		It("should correctly snapshot and revert", func() {
			id := l.Snapshot()

			l.AddLog(&ethtypes.Log{Address: a2})
			Expect(l.Size()).To(Equal(2))
			Expect(l.PeekAt(1).Address).To(Equal(a2))

			l.RevertToSnapshot(id)
			Expect(l.Size()).To(Equal(1))
		})

		It("should correctly get logs", func() {
			logs := l.Logs()
			Expect(logs).To(HaveLen(1))
			Expect(logs[0].TxHash).To(Equal(thash))
			Expect(logs[0].BlockHash).To(Equal(common.Hash{}))
			Expect(logs[0].BlockNumber).To(Equal(uint64(0)))

			logs = l.GetLogs(thash, bnum, bhash)
			Expect(logs).To(HaveLen(1))
			Expect(logs[0].BlockHash).To(Equal(bhash))
			Expect(logs[0].BlockNumber).To(Equal(bnum))
		})

		It("should corrctly finalize", func() {
			Expect(func() { l.Finalize() }).ToNot(Panic())
		})

		It("should correctly clone", func() {
			l.AddLog(&ethtypes.Log{Address: a2})
			Expect(l.Size()).To(Equal(2))
			Expect(l.PeekAt(1).Address).To(Equal(a2))

			l2 := utils.MustGetAs[*logs](l.Clone())
			Expect(l2.Size()).To(Equal(2))
			Expect(l2.PeekAt(0).Address).To(Equal(a1))
			Expect(l2.PeekAt(1).Address).To(Equal(a2))

			l2.AddLog(&ethtypes.Log{Address: a3})
			Expect(l2.Size()).To(Equal(3))
			Expect(l2.PeekAt(2).Address).To(Equal(a3))
			Expect(l.Size()).To(Equal(2))
		})
	})
})
