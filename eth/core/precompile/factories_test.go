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

package precompile_test

import (
	"context"
	"math/big"

	solidity "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	mockPrecompile, _ = solidity.MockPrecompileMetaData.GetAbi()
)

var _ = Describe("Container Factories", func() {

	Context("Stateless Container Factory", func() {
		var scf *precompile.StatelessFactory

		BeforeEach(func() {
			scf = precompile.NewStatelessFactory()
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
		var scf *precompile.StatefulFactory

		BeforeEach(func() {
			scf = precompile.NewStatefulFactory()
		})

		It("should correctly build stateful containers and log events", func() {
			pc, err := scf.Build(&mockStateful{&mockBase{}}, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockStateless{&mockBase{}}, nil)
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: StatefulContainerImpl"))
		})
	})

	Context("Bad Stateful Container", func() {
		var scf *precompile.StatefulFactory

		BeforeEach(func() {
			scf = precompile.NewStatefulFactory()
		})

		It("should error on missing precompile method for ABI method", func() {
			_, err := scf.Build(&badMockStateful{&mockStateful{&mockBase{}}}, nil)
			Expect(err.Error()).To(Equal("this ABI method does not have a corresponding precompile method: getOutputPartial()"))
		})

		It("should error on invalid precompile methods", func() {
			_, err := scf.Build(&invalidMockStateful{&mockStateful{&mockBase{}}}, nil)
			Expect(err.Error()).To(Equal("incomplete precompile Method"))
		})
	})

	Context("Dynamic Container Factory", func() {
		var dcf *precompile.DynamicFactory

		BeforeEach(func() {
			dcf = precompile.NewDynamicFactory()
		})

		It("should properly build dynamic container", func() {
			pc, err := dcf.Build(&mockDynamic{&mockStateful{&mockBase{}}}, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(pc).ToNot(BeNil())

			_, err = dcf.Build(&mockStateful{&mockBase{}}, nil)
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: DynamicContainerImpl"))
		})
	})
})

// MOCKS BELOW.

type mockBase struct{}

func (mb *mockBase) RegistryKey() common.Address {
	return common.Address{}
}

type mockStateless struct {
	*mockBase
}

func (ms *mockStateless) RequiredGas(_ []byte) uint64 {
	return 10
}

func (ms *mockStateless) Run(
	_ context.Context, _ precompile.EVM, _ []byte,
	_ common.Address, _ *big.Int, _ bool,
) ([]byte, error) {
	return nil, nil
}

func (ms *mockStateless) WithStateDB(vm.GethStateDB) vm.PrecompileContainer {
	return ms
}

type mockStateful struct {
	*mockBase
}

func (ms *mockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput": mockPrecompile.Methods["getOutput"],
	}
}

func (ms *mockStateful) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		{
			AbiSig:      "getOutput(string)",
			Execute:     getOutput,
			RequiredGas: 1,
		},
	}
}

func (ms *mockStateful) ABIEvents() map[string]abi.Event {
	return nil
}

func (ms *mockStateful) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}

func (ms *mockStateful) SetPlugin(precompile.Plugin) {}

type badMockStateful struct {
	*mockStateful
}

func (bms *badMockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput":        mock.Methods["getOutput"],
		"getOutputPartial": mock.Methods["getOutputPartial"],
	}
}

type invalidMockStateful struct {
	*mockStateful
}

func (ims *invalidMockStateful) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		{
			AbiSig:      "getOutput(string)",
			RequiredGas: 1,
		},
	}
}

type mockDynamic struct {
	*mockStateful
}

func (md *mockDynamic) Name() string {
	return ""
}
