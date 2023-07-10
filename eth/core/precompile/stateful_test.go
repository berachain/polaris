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
	"errors"
	"math/big"
	"reflect"

	solidity "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stateful Container", func() {
	var sc vm.PrecompileContainer
	var empty vm.PrecompileContainer
	var ctx context.Context
	var addr common.Address
	var readonly bool
	var value *big.Int
	var blank []byte
	var badInput = []byte{1, 2, 3, 4}
	var err error
	BeforeEach(func() {
		ctx = context.Background()
		sc, err = precompile.NewStateful(&mockStateful{&mockBase{}}, mockIdsToMethods)
		Expect(err).ToNot(HaveOccurred())
		empty, err = precompile.NewStateful(nil, nil)
		Expect(empty).To(BeNil())
		Expect(err).To(MatchError("the stateful precompile has no methods to run"))
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
			_, err = sc.Run(ctx, &mockEVM{}, blank, addr, value, readonly)
			Expect(err).To(MatchError("input bytes to precompile container are invalid"))

			// method not found
			_, err = sc.Run(ctx, &mockEVM{}, badInput, addr, value, readonly)
			Expect(err).To(MatchError("precompile method not found in contract ABI"))

			// geth unpacking error
			_, err = sc.Run(ctx, &mockEVM{}, append(getOutputABI.ID, byte(1), byte(2)), addr, value, readonly)
			Expect(err).To(HaveOccurred())

			// precompile exec error
			_, err = sc.Run(ctx, &mockEVM{}, getOutputPartialABI.ID, addr, value, readonly)
			Expect(err.Error()).To(Equal(
				"execution reverted: vm error [err during precompile execution] occurred during precompile execution of [getOutputPartial]", //nolint:lll // test.
			))

			// precompile returns vals when none expected
			var inputs []byte
			inputs, err = contractFuncStrABI.Inputs.Pack("string")
			Expect(err).ToNot(HaveOccurred())
			_, err = sc.Run(ctx, &mockEVM{}, append(contractFuncStrABI.ID, inputs...), addr, value, readonly)
			Expect(err).To(HaveOccurred())

			// geth output packing error
			inputs, err = contractFuncAddrABI.Inputs.Pack(addr)
			Expect(err).ToNot(HaveOccurred())
			_, err = sc.Run(ctx, &mockEVM{}, append(contractFuncAddrABI.ID, inputs...), addr, value, readonly)
			Expect(err).To(HaveOccurred())
		})

		It("should return properly for valid method calls", func() {
			// sc.WithStateDB(sdb)
			var inputs []byte
			inputs, err = getOutputABI.Inputs.Pack("string")
			Expect(err).ToNot(HaveOccurred())
			var ret []byte
			ret, err = sc.Run(ctx, &mockEVM{}, append(getOutputABI.ID, inputs...), addr, value, readonly)
			Expect(err).ToNot(HaveOccurred())
			var outputs []interface{}
			outputs, err = getOutputABI.Outputs.Unpack(ret)
			Expect(err).ToNot(HaveOccurred())
			Expect(outputs).To(HaveLen(1))
			Expect(
				reflect.ValueOf(outputs[0]).Index(0).FieldByName("CreationHeight").
					Interface().(*big.Int)).To(Equal(big.NewInt(1)))
			Expect(reflect.ValueOf(outputs[0]).Index(0).FieldByName("TimeStamp").
				Interface().(string)).To(Equal("string"))
		})
	})
})

// MOCKS BELOW.

type mockEVM struct {
	precompile.EVM
}

var (
	mock, _             = solidity.MockPrecompileMetaData.GetAbi()
	getOutputABI        = mock.Methods["getOutput"]
	getOutputPartialABI = mock.Methods["getOutputPartial"]
	contractFuncAddrABI = mock.Methods["contractFunc"]
	contractFuncStrABI  = mock.Methods["contractFuncStr"]
	mockIdsToMethods    = map[string]*precompile.Method{
		utils.UnsafeBytesToStr(getOutputABI.ID): precompile.NewMethod(
			&getOutputABI,
			getOutputABI.Sig,
			reflect.ValueOf(getOutput),
		),
		utils.UnsafeBytesToStr(getOutputPartialABI.ID): precompile.NewMethod(
			&getOutputPartialABI,
			getOutputPartialABI.Sig,
			reflect.ValueOf(getOutputPartial),
		),
		utils.UnsafeBytesToStr(contractFuncAddrABI.ID): precompile.NewMethod(
			&contractFuncAddrABI,
			contractFuncAddrABI.Sig,
			reflect.ValueOf(contractFuncAddrInput),
		),
		utils.UnsafeBytesToStr(contractFuncStrABI.ID): precompile.NewMethod(
			&contractFuncStrABI,
			contractFuncStrABI.Sig,
			reflect.ValueOf(contractFuncStrInput),
		),
	}
)

type mockObject struct {
	CreationHeight *big.Int
	TimeStamp      string
}

//revive:disable
func getOutput(
	_ *mockStateful,
	ctx context.Context,
	_ precompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	_ = ctx
	str, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, errors.New("cast error")
	}
	return []any{
		[]mockObject{
			{
				CreationHeight: big.NewInt(1),
				TimeStamp:      str,
			},
		},
	}, nil
}

func getOutputPartial(
	_ *mockStateful,
	ctx context.Context,
	_ precompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	_ = ctx
	return nil, errors.New("err during precompile execution")
}

func contractFuncAddrInput(
	_ *mockStateful,
	ctx context.Context,
	_ precompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	_ = ctx
	_, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, errors.New("cast error")
	}
	return []any{"invalid - should be *big.Int here"}, nil
}

func contractFuncStrInput(
	_ *mockStateful,
	ctx context.Context,
	_ precompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	_ = ctx
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, errors.New("cast error")
	}
	ans := big.NewInt(int64(len(addr)))
	return []any{ans}, nil
}
