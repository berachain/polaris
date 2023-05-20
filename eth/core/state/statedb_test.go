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

package state_test

import (
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/state/mock"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	alice = common.Address{1}
	bob   = common.Address{2}
	slot  = common.Hash{1}
)

var _ = Describe("StateDB", func() {
	var sdb vm.PolarisStateDB

	BeforeEach(func() {
		sdb = state.NewStateDB(mock.NewEmptyStatePlugin())
	})

	It("Should suicide correctly", func() {
		sdb.Snapshot()

		sdb.CreateAccount(alice)
		Expect(sdb.Suicide(alice)).To(BeFalse())
		Expect(sdb.HasSuicided(alice)).To(BeFalse())

		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		Expect(sdb.Suicide(bob)).To(BeTrue())
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSuicided(bob)).To(BeTrue())
	})

	It("should snapshot/revert", func() {
		Expect(func() {
			id := sdb.Snapshot()
			sdb.RevertToSnapshot(id)
		}).ToNot(Panic())
	})

	It("should handle access lists", func() {
		sdb.Prepare(
			params.Rules{IsBerlin: true, IsShanghai: true},
			alice, bob, &common.Address{3},
			[]common.Address{{4}},
			coretypes.AccessList{
				coretypes.AccessTuple{
					Address:     common.Address{5},
					StorageKeys: []common.Hash{{2}, {3}},
				},
			},
		)
		Expect(sdb.AddressInAccessList(alice)).To(BeTrue())
		Expect(sdb.AddressInAccessList(common.Address{3})).To(BeTrue())
		ap, sp := sdb.SlotInAccessList(common.Address{5}, common.Hash{2})
		Expect(ap).To(BeTrue())
		Expect(sp).To(BeTrue())
		Expect(sdb.AddressInAccessList(common.Address{3})).To(BeTrue())

		sdb.AddAddressToAccessList(alice)
		Expect(sdb.AddressInAccessList(alice)).To(BeTrue())
		ap, sp = sdb.SlotInAccessList(alice, slot)
		Expect(ap).To(BeTrue())
		Expect(sp).To(BeFalse())
		sdb.AddSlotToAccessList(alice, slot)
		ap, sp = sdb.SlotInAccessList(alice, slot)
		Expect(ap).To(BeTrue())
		Expect(sp).To(BeTrue())
	})

	It("should delete suicides on finalize", func() {
		sdb.Snapshot()
		sdb.SetTxContext(common.Hash{}, 0)

		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		Expect(sdb.Suicide(bob)).To(BeTrue())
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSuicided(bob)).To(BeTrue())

		sdb.Finalise(true)
		Expect(sdb.HasSuicided(bob)).To(BeFalse())
	})
})
