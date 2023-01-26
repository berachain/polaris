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

	"github.com/berachain/stargazer/core/precompile/container"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	"github.com/berachain/stargazer/x/evm/plugins/state"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stateful Container", func() {
	var sc *container.Stateful
	var empty *container.Stateful
	var ctx context.Context
	var addr common.Address
	var readonly bool
	var value *big.Int
	var blank []byte
	var badInput = []byte{1, 2, 3, 4}
	var sdb *mockSdb

	BeforeEach(func() {
		ctx = testutil.NewContextWithMultistores()
		sc = container.NewStateful(mockIdsToMethods)
		empty = container.NewStateful(nil)
		sdb = &mockSdb{&state.StateDB{}, 0}
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
			_, err := empty.Run(ctx, sdb, blank, addr, value, readonly)
			Expect(err).To(MatchError("the stateful precompile has no methods to run"))

			// invalid input
			_, err = sc.Run(ctx, sdb, blank, addr, value, readonly)
			Expect(err).To(MatchError("input bytes to precompile container are invalid"))

			// missing statedb
			_, err = sc.Run(ctx, nil, blank, addr, value, readonly)
			Expect(err).To(MatchError("statedb is not compatible with Stargazer"))

			// method not found
			_, err = sc.Run(ctx, sdb, badInput, addr, value, readonly)
			Expect(err).To(MatchError("precompile method not found in contract ABI"))

			// geth unpacking error
			_, err = sc.Run(ctx, sdb, append(getOutputABI.ID, byte(1), byte(2)), addr, value, readonly)
			Expect(err).ToNot(BeNil())

			// precompile exec error
			_, err = sc.Run(ctx, sdb, getOutputPartialABI.ID, addr, value, readonly)
			Expect(err.Error()).To(Equal("err during precompile execution: getOutputPartial"))

			// precompile returns vals when none expected
			inputs, err := contractFuncStrABI.Inputs.Pack("string")
			Expect(err).To(BeNil())
			_, err = sc.Run(ctx, sdb, append(contractFuncStrABI.ID, inputs...), addr, value, readonly)
			Expect(err).ToNot(BeNil())

			// geth output packing error
			inputs, err = contractFuncAddrABI.Inputs.Pack(addr)
			Expect(err).To(BeNil())
			_, err = sc.Run(ctx, sdb, append(contractFuncAddrABI.ID, inputs...), addr, value, readonly)
			Expect(err).ToNot(BeNil())
		})

		It("should return properly for valid method calls", func() {
			inputs, err := getOutputABI.Inputs.Pack("string")
			Expect(err).To(BeNil())
			ret, err := sc.Run(ctx, sdb, append(getOutputABI.ID, inputs...), addr, value, readonly)
			Expect(err).To(BeNil())
			Expect(sdb.count).To(Equal(2))
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
	mock                = solidity.MockPrecompileInterface
	getOutputABI        = mock.ABI.Methods["getOutput"]
	getOutputPartialABI = mock.ABI.Methods["getOutputPartial"]
	contractFuncAddrABI = mock.ABI.Methods["contractFunc"]
	contractFuncStrABI  = mock.ABI.Methods["contractFuncStr"]
	mockIdsToMethods    = map[string]*container.Method{
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

type mockSdb struct {
	vm.StargazerStateDB
	count int
}

func (ms *mockSdb) AddLog(log *coretypes.Log) {
	ms.count++
}

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
) ([]any, []*coretypes.Log, error) {
	str, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, nil, errors.New("cast error")
	}
	return []any{
		[]mockObject{
			{
				CreationHeight: big.NewInt(1),
				TimeStamp:      str,
			},
		},
	}, []*coretypes.Log{{}, {}}, nil
}

func getOutputPartial(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, []*coretypes.Log, error) {
	return nil, nil, errors.New("err during precompile execution")
}

func contractFuncAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, []*coretypes.Log, error) {
	_, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, nil, errors.New("cast error")
	}
	return []any{"invalid - should be *big.Int here"}, nil, nil
}

func contractFuncStrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, []*coretypes.Log, error) {
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, nil, errors.New("cast error")
	}
	ans := big.NewInt(int64(len(addr)))
	return []any{ans}, nil, nil
}
