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

package precompile

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("controller", func() {
	var c *manager
	var mr *mockRunner
	var ctx context.Context
	var ms *mockSdb

	BeforeEach(func() {
		mr = &mockRunner{}
		ms = &mockSdb{}
		ctx = context.Background()
		c = utils.MustGetAs[*manager](NewManager(mr, ms))
		err := c.Register(&mockStateless{})
		Expect(err).To(BeNil())
	})

	It("should find and run", func() {
		c.Reset(ctx)

		pc := c.Get(addr)
		Expect(pc).ToNot(BeNil())

		_, _, err := c.Run(pc, []byte{}, addr, new(big.Int), 10, true)
		Expect(err).To(BeNil())
		Expect(mr.called).To(BeTrue())
	})

	It("should not find an unregistered", func() {
		found := c.Has(common.BytesToAddress([]byte{2}))
		Expect(found).To(BeFalse())
	})
})

// MOCKS BELOW.

type mockRunner struct {
	called bool
}

func (mr *mockRunner) Run(
	ctx context.Context, ldb LogsDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	mr.called = true
	return nil, 0, nil
}

type mockSdb struct {
	vm.StargazerStateDB
}

type mockBase struct{}

var addr = common.BytesToAddress([]byte{1})

func (mb *mockBase) RegistryKey() common.Address {
	return addr
}

type mockStateless struct {
	*mockBase
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 0
}

func (ms *mockStateless) Run(
	ctx context.Context, input []byte,
	caller common.Address, value *big.Int, readonly bool,
) ([]byte, error) {
	return nil, nil
}
