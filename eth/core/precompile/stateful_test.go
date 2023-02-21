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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	solidity "pkg.berachain.dev/stargazer/eth/testutil/contracts/solidity/generated"
	"pkg.berachain.dev/stargazer/lib/utils"
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

	BeforeEach(func() {
		ctx = context.Background()
		sc = precompile.NewStateful(&mockStateful{&mockBase{}}, mockIdsToMethods)
		empty = precompile.NewStateful(nil, nil)
	})

	Describe("Test Required Gas", func() {
		It("should return 0 for invalid cases", func() {
			// empty input
			Expect(empty.RequiredGas(blank)).To(Equal(uint64(0)))

			// method not found
			Expect(sc.RequiredGas(badInput)).To(Equal(uint64(0)))

			// invalid input
			Expect(sc.RequiredGas(blank)).To(Equal(uint64(0)))
		})

		It("should properly return the required gas for valid methods", func() {
			Expect(sc.RequiredGas(getOutputABI.ID)).To(Equal(uint64(1)))
			Expect(sc.RequiredGas(getOutputPartialABI.ID)).To(Equal(uint64(10)))
			Expect(sc.RequiredGas(contractFuncAddrABI.ID)).To(Equal(uint64(100)))
			Expect(sc.RequiredGas(contractFuncStrABI.ID)).To(Equal(uint64(1000)))
		})
	})

	Describe("Test Run", func() {
		It("should return an error for invalid cases", func() {
			// empty input
			_, err := empty.Run(ctx, blank, addr, value, readonly)
			Expect(err).To(MatchError("the stateful precompile has no methods to run"))

			// invalid input
			_, err = sc.Run(ctx, blank, addr, value, readonly)
			Expect(err).To(MatchError("input bytes to precompile container are invalid"))

			// method not found
			_, err = sc.Run(ctx, badInput, addr, value, readonly)
			Expect(err).To(MatchError("precompile method not found in contract ABI"))

			// geth unpacking error
			_, err = sc.Run(ctx, append(getOutputABI.ID, byte(1), byte(2)), addr, value, readonly)
			Expect(err).ToNot(BeNil())

			// precompile exec error
			_, err = sc.Run(ctx, getOutputPartialABI.ID, addr, value, readonly)
			Expect(err.Error()).To(Equal("getOutputPartial: err during precompile execution"))

			// precompile returns vals when none expected
			inputs, err := contractFuncStrABI.Inputs.Pack("string")
			Expect(err).To(BeNil())
			_, err = sc.Run(ctx, append(contractFuncStrABI.ID, inputs...), addr, value, readonly)
			Expect(err).ToNot(BeNil())

			// geth output packing error
			inputs, err = contractFuncAddrABI.Inputs.Pack(addr)
			Expect(err).To(BeNil())
			_, err = sc.Run(ctx, append(contractFuncAddrABI.ID, inputs...), addr, value, readonly)
			Expect(err).ToNot(BeNil())
		})

		It("should return properly for valid method calls", func() {
			// sc.WithStateDB(sdb)
			inputs, err := getOutputABI.Inputs.Pack("string")
			Expect(err).To(BeNil())
			ret, err := sc.Run(ctx, append(getOutputABI.ID, inputs...), addr, value, readonly)
			Expect(err).To(BeNil())
			outputs, err := getOutputABI.Outputs.Unpack(ret)
			Expect(err).To(BeNil())
			Expect(len(outputs)).To(Equal(1))
			Expect(
				reflect.ValueOf(outputs[0]).Index(0).FieldByName("CreationHeight").
					Interface().(*big.Int)).To(Equal(big.NewInt(1)))
			Expect(reflect.ValueOf(outputs[0]).Index(0).FieldByName("TimeStamp").
				Interface().(string)).To(Equal("string"))
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
	mockIdsToMethods    = map[string]*precompile.Method{
		utils.UnsafeBytesToStr(getOutputABI.ID): {
			AbiSig:      getOutputABI.Sig,
			AbiMethod:   &getOutputABI,
			Execute:     getOutput,
			RequiredGas: 1,
		},
		utils.UnsafeBytesToStr(getOutputPartialABI.ID): {
			AbiSig:      getOutputPartialABI.Sig,
			AbiMethod:   &getOutputPartialABI,
			Execute:     getOutputPartial,
			RequiredGas: 10,
		},
		utils.UnsafeBytesToStr(contractFuncAddrABI.ID): {
			AbiSig:      contractFuncAddrABI.Sig,
			AbiMethod:   &contractFuncAddrABI,
			Execute:     contractFuncAddrInput,
			RequiredGas: 100,
		},
		utils.UnsafeBytesToStr(contractFuncStrABI.ID): {
			AbiSig:      contractFuncStrABI.Sig,
			AbiMethod:   &contractFuncStrABI,
			Execute:     contractFuncStrInput,
			RequiredGas: 1000,
		},
	}
)

type mockObject struct {
	CreationHeight *big.Int
	TimeStamp      string
}

func getOutput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
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
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	return nil, errors.New("err during precompile execution")
}

func contractFuncAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	_, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, errors.New("cast error")
	}
	return []any{"invalid - should be *big.Int here"}, nil
}

func contractFuncStrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, errors.New("cast error")
	}
	ans := big.NewInt(int64(len(addr)))
	return []any{ans}, nil
}
