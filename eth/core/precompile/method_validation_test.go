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
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Method", func() {
	var precompileABI map[string]abi.Method
	var m *mockImpl
	BeforeEach(func() {
		precompileABI = map[string]abi.Method{
			"exampleFunc": {
				Name:    "exampleFunc",
				RawName: "exampleFunc",
				Inputs: []abi.Argument{
					{
						Name: "a",
						Type: abi.Type{T: abi.IntTy},
					},
					{
						Name: "b",
						Type: abi.Type{T: abi.AddressTy},
					},
					{
						Name: "c",
						Type: abi.Type{
							T:         abi.SliceTy,
							Elem:      &abi.Type{T: abi.TupleTy},
							TupleType: reflect.TypeOf(mockStruct{}),
						},
					},
				},
				Outputs: []abi.Argument{
					{
						Name: "d",
						Type: abi.Type{T: abi.BoolTy},
					},
				},
			},
			"zeroReturn": {
				Name:    "zeroReturn",
				RawName: "zeroReturn",
				Outputs: []abi.Argument{},
			},
		}
		m = &mockImpl{}
	})
	Context("findMatchingAbiMethod", func() {
		It("should validate args successfully", func() {
			exampleFuncValue, found := reflect.TypeOf(m).MethodByName("ExampleFunc")
			Expect(found).To(BeTrue())

			methodName, err := findMatchingABIMethod(exampleFuncValue, precompileABI)
			Expect(err).ToNot(HaveOccurred())
			Expect(methodName).To(Equal("exampleFunc"))

			sliceA := []uint64{}
			sliceB := []*big.Int{}

			Expect(validateArg(
				reflect.ValueOf(sliceA),
				reflect.ValueOf(sliceB)).Error()).To(Equal(
				"type mismatch: []uint64 != []*big.Int",
			))
		})
	})

	Context("validateArg", func() {
		It("should error when struct fields aren't the same", func() {
			sliceA := []mockStruct{}
			sliceB := []mockStructBad{}
			Expect(validateArg(
				reflect.ValueOf(sliceA),
				reflect.ValueOf(sliceB)).Error()).To(Equal(
				"type mismatch: *big.Int != uint64",
			))
		})
		It("should error when we point to a non-struct", func() {
			randomPointer := 69
			implMethodVarType := &randomPointer
			abiMethodVarType := &mockStruct{}
			Expect(validateArg(
				reflect.ValueOf(implMethodVarType).Elem(),
				reflect.ValueOf(abiMethodVarType)).Error()).To(Equal(
				"type mismatch: int != *precompile.mockStruct",
			))
		})
	})

	It("should panic when our ABI method does not return anything", func() {
		zeroReturn := precompileABI["zeroReturn"]
		mockMethod, _ := reflect.TypeOf(m).MethodByName("MockMethod")

		//nolint:errcheck // it's going to panic
		Expect(func() { validateOutputs(mockMethod, &zeroReturn) }).To(Panic())
	})
	It("should error when we have different structs as params", func() {
		m := mockStruct{}
		mb := mockStructBad{}

		Expect(validateArg(reflect.New(reflect.TypeOf(m)).Elem(), reflect.New(reflect.TypeOf(mb)).Elem())).To(HaveOccurred())

		Expect(validateStruct(reflect.TypeOf(m), reflect.TypeOf(mb))).To(HaveOccurred())
		mbn := mockStructBadNumFields{}

		Expect(validateStruct(reflect.TypeOf(m), reflect.TypeOf(mbn))).To(HaveOccurred())

		notAStruct := 69

		Expect(validateStruct(reflect.TypeOf(m), reflect.TypeOf(notAStruct)).Error()).To(Equal(
			"validateStruct: not a struct"))
	})

	Context("validateOutputs", func() {
		It("should error when our impl and abi outputs aren't correct", func() {
			exampleFunc := precompileABI["exampleFunc"]

			noErrorReturn, found := reflect.TypeOf(m).MethodByName("NoErrorReturn")
			Expect(found).To(BeTrue())
			Expect(validateOutputs(noErrorReturn, &exampleFunc).Error()).To(Equal(
				"last return type must be error, got bool"))

			numReturnMismatch, found := reflect.TypeOf(m).MethodByName("NumReturnMismatch")
			Expect(found).To(BeTrue())
			Expect(validateOutputs(numReturnMismatch, &exampleFunc).Error()).To(Equal(
				"number of return types mismatch"))

			returnTypeMismatch, found := reflect.TypeOf(m).MethodByName("ReturnTypeMismatch")
			Expect(found).To(BeTrue())
			Expect(validateOutputs(returnTypeMismatch, &exampleFunc).Error()).To(Equal(
				"return type mismatch: string != bool"))
		})
	})

	Context("findMatchingABIMethod", func() {

		It("should return ErrNoImplMethodSubstringMatchesABIMethods", func() {
			mockMethod, found := reflect.TypeOf(m).MethodByName("MockMethod")
			Expect(found).To(BeTrue())
			methodName, err := findMatchingABIMethod(mockMethod, precompileABI)
			Expect(methodName).To(Equal(""))
			Expect(err.Error()).To(Equal(ErrNoImplMethodSubstringMatchesABIMethods.Error() + ": MockMethod"))
		})
	})
})

type mockImpl struct{}

type mockStruct struct {
	_ *big.Int
}

type mockStructBad struct {
	_ uint64
}

type mockStructBadNumFields struct {
	_ *big.Int
	_ *big.Int
}

func (m *mockImpl) MockMethod() error { return nil }

func (m *mockImpl) ExampleFunc(_ *big.Int, _ common.Address, _ []mockStruct) (bool, error) {
	return true, nil
}

func (m *mockImpl) ExampleFuncBad(_ *big.Int, _ common.Address, _ []mockStructBad) (bool, error) {
	return true, nil
}

func (m *mockImpl) NoErrorReturn(_ *big.Int) (bool, bool) {
	return true, true
}

func (m *mockImpl) NumReturnMismatch(_ *big.Int) error {
	return nil
}

func (m *mockImpl) ReturnTypeMismatch(_ *big.Int) (string, error) {
	return "bera", nil
}
