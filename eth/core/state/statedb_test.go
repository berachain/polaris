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
