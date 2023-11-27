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
