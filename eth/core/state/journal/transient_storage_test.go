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
