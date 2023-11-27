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
