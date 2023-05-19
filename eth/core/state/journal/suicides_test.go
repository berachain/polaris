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
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/state/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Suicides", func() {
	var s *suicides

	BeforeEach(func() {
		sp := mock.NewEmptyStatePlugin()
		sp.CreateAccount(common.HexToAddress("0x1"))
		sp.CreateAccount(common.HexToAddress("0x3"))
		sp.GetCodeHashFunc = func(address common.Address) common.Hash {
			if address == common.HexToAddress("0x1") || address == common.HexToAddress("0x3") {
				return common.Hash{0x1}
			}
			return common.Hash{}
		}
		s = NewSuicides(sp)
	})

	It("should have the correct registry key", func() {
		Expect(s.RegistryKey()).To(Equal(suicidesRegistryKey))
	})

	It("should work correctly in the scope of a tx", func() {
		Expect(s.GetSuicides()).To(BeEmpty())

		s.Snapshot()
		Expect(s.Suicide(common.HexToAddress("0x2"))).To(BeFalse())
		Expect(s.Suicide(common.HexToAddress("0x1"))).To(BeTrue())
		Expect(s.HasSuicided(common.HexToAddress("0x2"))).To(BeFalse())
		Expect(s.HasSuicided(common.HexToAddress("0x1"))).To(BeTrue())

		snap2 := s.Snapshot()
		Expect(s.Suicide(common.HexToAddress("0x3"))).To(BeTrue())
		Expect(s.HasSuicided(common.HexToAddress("0x3"))).To(BeTrue())
		Expect(s.HasSuicided(common.HexToAddress("0x1"))).To(BeTrue())
		Expect(s.GetSuicides()).To(HaveLen(2))

		s.RevertToSnapshot(snap2)
		Expect(s.HasSuicided(common.HexToAddress("0x1"))).To(BeTrue())
		Expect(s.HasSuicided(common.HexToAddress("0x3"))).To(BeFalse())
		Expect(s.GetSuicides()).To(HaveLen(1))

		s.Finalize()
		Expect(s.lastSnapshot).To(Equal(-1))
		Expect(s.journal.Size()).To(Equal(0))
	})

	It("should not suicide when snapshot is not called", func() {
		Expect(s.Suicide(common.HexToAddress("0x1"))).To(BeFalse())
		Expect(s.HasSuicided(common.HexToAddress("0x1"))).To(BeFalse())
	})
})
