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
	"reflect"

	solidity "github.com/berachain/polaris/contracts/bindings/testing"
	pvm "github.com/berachain/polaris/eth/core/vm"
	vmmock "github.com/berachain/polaris/eth/core/vm/mock"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stateful Container", func() {
	var sc vm.PrecompiledContract
	var empty vm.PrecompiledContract
	var blank []byte
	var badInput = []byte{1, 2, 3, 4}
	var err error
	var ctx context.Context

	BeforeEach(func() {
		sc, err = NewStatefulContainer(&mockStateful{&mockBase{}}, mockIdsToMethods)
		Expect(err).ToNot(HaveOccurred())
		empty, err = NewStatefulContainer(nil, nil)
		Expect(empty).To(BeNil())
		Expect(err).To(MatchError("the stateful precompile has no methods to run"))
		ctx = pvm.NewPolarContext(
			context.Background(),
			vmmock.NewEVM(),
			common.Address{},
			big.NewInt(0),
		)
	})

	Describe("Test Required Gas", func() {
		It("should return 0 in all cases", func() {
			// method not found
			Expect(sc.RequiredGas(badInput)).To(Equal(uint64(0)))

			// invalid input
			Expect(sc.RequiredGas(blank)).To(Equal(uint64(0)))
		})

	})

	Describe("Test Run", func() {
		It("should return an error for invalid cases", func() {
			// invalid input
			_, err = sc.Run(
				ctx,
				pvm.UnwrapPolarContext(ctx).Evm(),
				blank,
				pvm.UnwrapPolarContext(ctx).MsgSender(),
				pvm.UnwrapPolarContext(ctx).MsgValue(),
			)
			Expect(err).To(MatchError("input bytes to precompile container are invalid"))

			// method not found
			_, err = sc.Run(
				ctx,
				pvm.UnwrapPolarContext(ctx).Evm(),
				badInput, pvm.UnwrapPolarContext(ctx).MsgSender(),
				pvm.UnwrapPolarContext(ctx).MsgValue(),
			)
			Expect(err).To(MatchError("precompile method not found in contract ABI"))

			// geth unpacking error
			_, err = sc.Run(ctx,
				pvm.UnwrapPolarContext(ctx).Evm(),
				append(getOutputABI.ID, byte(1), byte(2)),
				pvm.UnwrapPolarContext(ctx).MsgSender(),
				pvm.UnwrapPolarContext(ctx).MsgValue(),
			)
			Expect(err).To(HaveOccurred())

			// precompile exec error
			_, err = sc.Run(
				ctx,
				pvm.UnwrapPolarContext(ctx).Evm(),
				getOutputPartialABI.ID,
				pvm.UnwrapPolarContext(ctx).MsgSender(),
				pvm.UnwrapPolarContext(ctx).MsgValue(),
			)
			//nolint:lll // error message.
			Expect(err.Error()).To(Equal(
				"execution reverted: vm error [err during precompile execution] occurred during precompile execution of [getOutputPartial]",
			))
		})

		It("should return properly for valid method calls", func() {
			var inputs []byte
			inputs, err = getOutputABI.Inputs.Pack("string")
			Expect(err).ToNot(HaveOccurred())
			var ret []byte
			ret, err = sc.Run(
				ctx,
				pvm.UnwrapPolarContext(ctx).Evm(),
				append(getOutputABI.ID, inputs...),
				pvm.UnwrapPolarContext(ctx).MsgSender(),
				pvm.UnwrapPolarContext(ctx).MsgValue(),
			)
			Expect(err).ToNot(HaveOccurred())
			var outputs []interface{}
			outputs, err = getOutputABI.Outputs.Unpack(ret)
			Expect(err).ToNot(HaveOccurred())
			Expect(outputs).To(HaveLen(1))
			Expect(
				reflect.ValueOf(outputs[0]).
					Index(0).
					FieldByName("CreationHeight").
					Interface().(*big.Int),
			).To(Equal(big.NewInt(1)))
			Expect(
				reflect.ValueOf(outputs[0]).Index(0).FieldByName("TimeStamp").Interface().(string),
			).To(Equal("string"))
		})
	})
})

// MOCKS BELOW.

var (
	mock, _             = solidity.MockPrecompileMetaData.GetAbi()
	getOutputABI        = mock.Methods["getOutput"]
	getOutputPartialABI = mock.Methods["getOutputPartial"]
	contractFuncAddrABI = mock.Methods["contractFunc"]
	contractFuncStrABI  = mock.Methods["contractFuncStr"]
	overloadedFuncABI   = mock.Methods["overloadedFunc"]
	overloadedFunc0ABI  = mock.Methods["overloadedFunc0"]
	mockStatefulDummy   = &mockStateful{&mockBase{}}
	getOutputFunc, _    = reflect.TypeOf(
		mockStatefulDummy).MethodByName("GetOutput")
	getOutputPartialFunc, _ = reflect.TypeOf(
		mockStatefulDummy).MethodByName("GetOutputPartial")
	contractFuncAddrInputFunc, _ = reflect.TypeOf(
		mockStatefulDummy).MethodByName("ContractFuncAddrInput")
	contractFuncStrInputFunc, _ = reflect.TypeOf(
		mockStatefulDummy).MethodByName("ContractFuncStrInput")
	overloadedFunc, _ = reflect.TypeOf(
		mockStatefulDummy).MethodByName("OverloadedFunc")
	overloadedFunc0, _ = reflect.TypeOf(
		mockStatefulDummy).MethodByName("OverloadedFunc0")
	mockIdsToMethods = map[methodID]*method{
		methodID(getOutputABI.ID): newMethod(
			mockStatefulDummy,
			getOutputABI,
			getOutputFunc,
		),
		methodID(getOutputPartialABI.ID): newMethod(
			mockStatefulDummy,
			getOutputPartialABI,
			getOutputPartialFunc,
		),
		methodID(contractFuncAddrABI.ID): newMethod(
			mockStatefulDummy,
			contractFuncAddrABI,
			contractFuncAddrInputFunc,
		),
		methodID(contractFuncStrABI.ID): newMethod(
			mockStatefulDummy,
			contractFuncStrABI,
			contractFuncStrInputFunc,
		),
		methodID(overloadedFuncABI.ID): newMethod(
			mockStatefulDummy,
			overloadedFuncABI,
			overloadedFunc,
		),
		methodID(contractFuncStrABI.ID): newMethod(
			mockStatefulDummy,
			overloadedFunc0ABI,
			overloadedFunc0,
		),
	}
)

type mockObject struct {
	CreationHeight *big.Int
	TimeStamp      string
}
