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

package journal

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/utils"
)

var _ = Describe("Logs", func() {
	var l *logs
	var h1 = common.BytesToHash([]byte{1})
	var h2 = common.BytesToHash([]byte{2})
	var a1 = common.BytesToAddress([]byte{3})
	var a2 = common.BytesToAddress([]byte{4})
	var u1 = uint(5)
	var u2 = uint(6)

	BeforeEach(func() {
		l = utils.MustGetAs[*logs](NewLogs())
		Expect(l.Capacity()).To(Equal(32))
	})

	It("should have the correct registry key", func() {
		Expect(l.RegistryKey()).To(Equal("logs"))
	})

	When("adding logs", func() {
		BeforeEach(func() {
			l.AddLog(&coretypes.Log{Address: a1})
			Expect(l.Size()).To(Equal(1))
			Expect(l.PeekAt(0).Address).To(Equal(a1))
		})

		It("should correctly snapshot and revert", func() {
			id := l.Snapshot()

			l.AddLog(&coretypes.Log{Address: a2})
			Expect(l.Size()).To(Equal(2))
			Expect(l.PeekAt(1).Address).To(Equal(a2))

			l.RevertToSnapshot(id)
			Expect(l.Size()).To(Equal(1))
		})

		It("should correctly build logs", func() {
			logs := l.BuildLogsAndClear(h1, h2, u1, u2)
			Expect(len(logs)).To(Equal(1))
			Expect(logs[0].Address).To(Equal(a1))
			Expect(logs[0].TxHash).To(Equal(h1))
			Expect(logs[0].BlockHash).To(Equal(h2))
			Expect(logs[0].TxIndex).To(Equal(u1))
			Expect(logs[0].Index).To(Equal(u2))
		})

		It("should corrctly finalize", func() {
			Expect(func() { l.Finalize() }).ToNot(Panic())
		})
	})
})
