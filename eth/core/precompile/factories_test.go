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

	solidity "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	_, _ = solidity.MockPrecompileMetaData.GetAbi()
)

var _ = Describe("Container Factories", func() {

	Context("Stateless Container Factory", func() {
		var scf *StatelessFactory

		BeforeEach(func() {
			scf = NewStatelessFactory()
		})

		It("should build stateless precompile containers", func() {
			pc, err := scf.Build(&mockStateless{&mockBase{}}, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockBase{}, nil)
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: StatelessContainerImpl"))
		})
	})

	Context("Stateful Container Factory", func() {
		var scf *StatefulFactory

		BeforeEach(func() {
			scf = NewStatefulFactory()
		})

		It("should correctly build stateful containers and log events", func() {
			pc, err := scf.Build(&mockStateful{&mockBase{}}, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockStateless{&mockBase{}}, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("Bad Stateful Container", func() {
		var scf *StatefulFactory

		BeforeEach(func() {
			scf = NewStatefulFactory()
		})

		It("should error on missing precompile method for ABI method", func() {
			_, err := scf.Build(&badMockStateful{&mockBase{}}, nil)
			var _ StatefulImpl = (*mockBase)(nil)
			Expect(err.Error()).To(Equal("this ABI method does not have a corresponding precompile method: getOutputPartial"))
		})
	})

	Context("Overloaded Stateful Container", func() {

		It("should construct a stateful container with overloaded methods", func() {
			scf := NewStatefulFactory()
			os := &overloadedStateful{&mockBase{}}
			stateful, err := scf.Build(os, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(stateful).ToNot(BeNil())
		})

	})
})

// MOCKS BELOW.

// ============================================================================

type mockBase struct{}

func (mb *mockBase) RegistryKey() common.Address {
	return common.Address{}
}

func (mb *mockBase) ABIMethods() map[string]abi.Method { return nil }

func (mb *mockBase) ABIEvents() map[string]abi.Event { return nil }

// CustomValueDecoders should return a map of event attribute keys to value decoder
// functions. This is used to decode event attribute values that require custom decoding
// logic.
func (mb *mockBase) CustomValueDecoders() ValueDecoders { return nil }

func (mb *mockBase) SetPlugin(_ Plugin) {}

// ============================================================================.
type mockStateless struct {
	*mockBase
}

func (ms *mockStateless) RequiredGas(_ []byte) uint64 {
	return 10
}

func (ms *mockStateless) Run(
	_ context.Context, _ vm.PrecompileEVM, _ []byte,
	_ common.Address, _ *big.Int,
) ([]byte, error) {
	return nil, nil
}

func (ms *mockStateless) WithStateDB(vm.GethStateDB) vm.PrecompileContainer {
	return ms
}

// ============================================================================.
type mockStateful struct {
	*mockBase
}

// ============================================================================.
type badMockStateful struct {
	*mockBase
}

func (bms *badMockStateful) GetOutput(_ context.Context, _ string) ([]byte, error) {
	return nil, nil
}

func (bms *badMockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput":        mock.Methods["getOutput"],
		"getOutputPartial": mock.Methods["getOutputPartial"],
	}
}

// ============================================================================

type overloadedStateful struct {
	*mockBase
}

func (os *overloadedStateful) OverloadedFunc(_ context.Context) (*big.Int, error) {
	return big.NewInt(69), nil
}

func (os *overloadedStateful) OverloadedFunc0(_ context.Context, _ *big.Int) (*big.Int, error) {
	return big.NewInt(420), nil
}

func (os *overloadedStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"overloadedFunc":  mock.Methods["overloadedFunc"],
		"overloadedFunc0": mock.Methods["overloadedFunc0"],
	}
}
