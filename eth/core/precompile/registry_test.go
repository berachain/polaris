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

package precompile

import (
	"context"
	"errors"
	"math/big"

	"github.com/berachain/stargazer/eth/core/precompile/container"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("registry", func() {
	pr := NewRegistry()

	It("should error on incorrect precompile types", func() {
		err := pr.Register(&mockBase{})
		Expect(err.Error()).To(Equal("this contract does not implement the required precompile contract interface"))
	})

	It("should create a stateless container", func() {
		err := pr.Register(&mockStateless{&mockBase{}})
		Expect(err).To(BeNil())
	})

	It("should create a stateful container", func() {
		err := pr.Register(&mockStateful{&mockBase{}})
		Expect(err).To(BeNil())
	})

	It("should create a dynamic container", func() {
		err := pr.Register(&mockDynamic{&mockStateful{&mockBase{}}})
		Expect(err).To(BeNil())
	})

	It("should error on building the container", func() {
		err := pr.Register(&mockStatefulBad{&mockStateful{&mockBase{}}})
		Expect(err.Error()).To(Equal("this ABI method does not have a corresponding precompile method: getOutputPartial()"))
	})
})

// MOCKS BELOW.

type mockBase struct{}

var addr = common.BytesToAddress([]byte{1})

func (mb *mockBase) RegistryKey() common.Address {
	return addr
}

type mockStateless struct {
	*mockBase
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 0
}

func (ms *mockStateless) Run(
	ctx context.Context, input []byte,
	caller common.Address, value *big.Int, readonly bool,
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
		"getOutput": solidity.MockPrecompileInterface.ABI.Methods["getOutput"],
	}
}

func (ms *mockStateful) PrecompileMethods() container.Methods {
	return container.Methods{
		{
			AbiSig:      "getOutput(string)",
			Execute:     getOutput,
			RequiredGas: 1,
		},
	}
}

type mockStatefulBad struct {
	*mockStateful
}

func (msb *mockStatefulBad) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput":        solidity.MockPrecompileInterface.ABI.Methods["getOutput"],
		"getOutputPartial": solidity.MockPrecompileInterface.ABI.Methods["getOutputPartial"],
	}
}

type mockDynamic struct {
	*mockStateful
}

func (md *mockDynamic) Name() string {
	return "name"
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
) ([]any, []*types.Log, error) {
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
	}, nil, nil
}
