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
	"errors"
	"math/big"
	"testing"

	tmock "github.com/stretchr/testify/mock"

	"github.com/berachain/polaris/eth/core/state"
	"github.com/berachain/polaris/eth/core/state/mock"
	"github.com/berachain/polaris/eth/core/state/mocks"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	alice = common.Address{1}
	bob   = common.Address{2}
	slot  = common.Hash{1}
)

func TestState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/core/state")
}

var _ = Describe("StateDB", func() {
	var sdb state.StateDB
	var sp *mock.PluginMock
	var pp *mocks.PrecompilePlugin

	BeforeEach(func() {
		sp = mock.NewEmptyStatePlugin()
		pp = mocks.NewPrecompilePlugin(GinkgoT())
		sdb = state.NewStateDB(sp, pp)
	})

	It("Should SelfDestruct correctly", func() {
		sdb.Snapshot()

		sdb.CreateAccount(alice)
		sdb.SelfDestruct(alice)
		Expect(sdb.HasSelfDestructed(alice)).To(BeFalse())

		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		sdb.SelfDestruct(bob)
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSelfDestructed(bob)).To(BeTrue())
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
			ethtypes.AccessList{
				ethtypes.AccessTuple{
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

	It("should delete SelfDestructs on finalize", func() {
		sdb.Snapshot()
		sdb.SetTxContext(common.Hash{}, 0)

		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		sdb.SelfDestruct(bob)
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSelfDestructed(bob)).To(BeTrue())

		sdb.Finalise(true)
		Expect(sdb.HasSelfDestructed(bob)).To(BeFalse())
	})

	It("should handle saved errors", func() {
		sp.ErrorFunc = func() error {
			return errors.New("mocked saved error")
		}

		// check that sdb correctly uses the saved error from state plugin
		err := sdb.Error()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("mocked saved error"))
	})

	It("should return code for precompiles", func() {
		pp.On("Get", common.Address{0x7}, tmock.Anything).Return(nil, true).Once()
		Expect(sdb.GetCode(common.Address{0x7})).To(Equal([]byte{0x1}))
		pp.On("Get", common.Address{0x7}, tmock.Anything).Return(nil, false).Once()
		Expect(sdb.GetCode(common.Address{0x7})).To(Equal([]byte{}))
	})
})
