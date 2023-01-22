// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package container_test

import (
	"context"
	"errors"
	"math/big"
	"reflect"

	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	mock                = solidity.MockPrecompileInterface
	getOutputABI        = mock.ABI.Methods["getOutput"]
	getOutputPartialABI = mock.ABI.Methods["getOutputPartial"]
	contractFuncAddrABI = mock.ABI.Methods["contractFunc"]
	contractFuncStrABI  = mock.ABI.Methods["contractFuncStr"]
)

var _ = Describe("Stateful Container", func() {
	var sc *container.StatefulContainer
	var empty *container.StatefulContainer
	var ctx context.Context
	var addr common.Address
	var readonly bool
	var value *big.Int
	var blank []byte
	var badInput = []byte{1, 2, 3, 4}

	BeforeEach(func() {
		ctx = testutil.NewContextWithMultistores()
		sc = container.NewStatefulContainer(mockIdsToMethods())
		empty = container.NewStatefulContainer(nil)
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

func mockIdsToMethods() map[string]*types.Method {
	idsToMethods := make(map[string]*types.Method)
	idsToMethods[utils.UnsafeBytesToStr(getOutputABI.ID)] = &types.Method{
		AbiSig:      getOutputABI.Sig,
		AbiMethod:   &getOutputABI,
		Execute:     getOutput,
		RequiredGas: 1,
	}
	idsToMethods[utils.UnsafeBytesToStr(getOutputPartialABI.ID)] = &types.Method{
		AbiSig:      getOutputPartialABI.Sig,
		AbiMethod:   &getOutputPartialABI,
		Execute:     getOutputPartial,
		RequiredGas: 10,
	}
	idsToMethods[utils.UnsafeBytesToStr(contractFuncAddrABI.ID)] = &types.Method{
		AbiSig:      contractFuncAddrABI.Sig,
		AbiMethod:   &contractFuncAddrABI,
		Execute:     contractFuncAddrInput,
		RequiredGas: 100,
	}
	idsToMethods[utils.UnsafeBytesToStr(contractFuncStrABI.ID)] = &types.Method{
		AbiSig:      contractFuncStrABI.Sig,
		AbiMethod:   &contractFuncStrABI,
		Execute:     contractFuncStrInput,
		RequiredGas: 1000,
	}

	return idsToMethods
}
