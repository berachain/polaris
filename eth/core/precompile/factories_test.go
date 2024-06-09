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

package precompile

import (
	"context"
	"errors"
	"math/big"

	solidity "github.com/berachain/polaris/contracts/bindings/testing"
	"github.com/berachain/polaris/eth/accounts/abi"
	pvm "github.com/berachain/polaris/eth/core/vm"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

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
			pc, err := scf.Build(&mockStateless{}, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockBase{}, nil)
			Expect(err.Error()).
				To(Equal(
					"wrong container factory for this precompile implementation: StatelessImpl"))
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
			statelessFactory := NewStatelessFactory()
			_, err = statelessFactory.Build(&mockStateless{}, nil)
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
			Expect(err.Error()).
				To(Equal(
					"this ABI method does not have a " +
						"corresponding precompile method: getOutputPartial"))
		})
	})

	Context("Overloaded Stateful Container", func() {
		It("should construct a stateful container with overloaded methods", func() {
			scf := NewStatefulFactory()
			os := &mockStateful{&mockBase{}}
			stateful, err := scf.Build(os, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(stateful).ToNot(BeNil())
		})
	})
})

// MOCKS BELOW.

// ============================================================================

// mockBase is the base contract for STATEFUL impls.
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
type mockStateless struct{}

func (ms *mockStateless) RegistryKey() common.Address {
	return common.Address{}
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

// ============================================================================.
type mockStateful struct {
	*mockBase
}

func (ms *mockStateful) RegistryKey() common.Address {
	return common.HexToAddress("0x696969696969")
}

func (ms *mockStateful) ABIEvents() map[string]abi.Event {
	return mock.Events
}

func (ms *mockStateful) ABIMethods() map[string]abi.Method {
	return mock.Methods
}

func (ms *mockStateful) GetOutput(
	ctx context.Context,
	str string,
) ([]mockObject, error) {
	pvm.UnwrapPolarContext(ctx).Evm().GetStateDB().AddLog(&ethtypes.Log{Address: common.Address{0x1}})
	return []mockObject{
		{
			CreationHeight: big.NewInt(1),
			TimeStamp:      str,
		},
	}, nil
}

func (ms *mockStateful) GetOutputPartial(
	_ context.Context,
) (*mockObject, error) {
	return &mockObject{}, errors.New("err during precompile execution")
}

func (ms *mockStateful) ContractFuncAddrInput(
	_ context.Context,
	_ common.Address,
) (*big.Int, error) {
	return big.NewInt(12112), nil
}

func (ms *mockStateful) ContractFuncStrInput(
	_ context.Context,
	_ string,
) (bool, error) {
	return true, nil
}

func (ms *mockStateful) OverloadedFunc(_ context.Context) (*big.Int, error) {
	return big.NewInt(69), nil
}

func (ms *mockStateful) OverloadedFunc0(_ context.Context, _ *big.Int) (*big.Int, error) {
	return big.NewInt(420), nil
}

// ============================================================================.
type badMockStateful struct {
	*mockBase
}

func (bms *badMockStateful) GetOutput(_ context.Context, _ string) ([]mockObject, error) {
	return nil, nil
}

func (bms *badMockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput":        mock.Methods["getOutput"],
		"getOutputPartial": mock.Methods["getOutputPartial"],
	}
}
