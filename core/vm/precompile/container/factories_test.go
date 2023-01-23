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

	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/core/vm/precompile/log/cosmos"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	"github.com/berachain/stargazer/types/abi"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Factories", func() {
	var lr *log.Registry

	Context("Stateless Container Factory", func() {
		var scf *container.StatelessContainerFactory

		BeforeEach(func() {
			scf = container.NewStatelessContainerFactory()
		})

		It("should build stateless precompile containers", func() {
			pc, err := scf.Build(&mockStateless{&mockBase{}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockBase{})
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: StatelessContainerImpl"))
		})
	})

	Context("Stateful Container Factory", func() {
		var scf *container.StatefulContainerFactory

		BeforeEach(func() {
			lr = log.NewRegistry(cosmos.NewTranslator(nil))
			scf = container.NewStatefulContainerFactory(lr)
		})

		It("should correctly build stateful containers and log events", func() {
			pc, err := scf.Build(&mockStateful{&mockBase{}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockStateless{&mockBase{}})
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: StatefulContainerImpl"))
		})
	})

	Context("Bad Stateful Container", func() {
		var scf *container.StatefulContainerFactory

		BeforeEach(func() {
			lr = log.NewRegistry(cosmos.NewTranslator(nil))
			scf = container.NewStatefulContainerFactory(lr)
		})

		It("should error on missing precompile method for ABI method", func() {
			_, err := scf.Build(&badMockStateful{&mockStateful{&mockBase{}}})
			Expect(err.Error()).To(Equal("this ABI method does not have a corresponding precompile method: getOutputPartial()"))
		})
	})

	Context("Dynamic Container Factory", func() {
		var dcf *container.DynamicContainerFactory

		BeforeEach(func() {
			lr = log.NewRegistry(cosmos.NewTranslator(nil))
			dcf = container.NewDynamicContainerFactory(lr)
		})

		It("should properly build dynamic container", func() {
			pc, err := dcf.Build(&mockDynamic{&mockStateful{&mockBase{}}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = dcf.Build(&mockStateful{&mockBase{}})
			Expect(err.Error()).To(Equal("this precompile contract implementation is not implemented: DynamicContainerImpl"))
		})
	})
})

// MOCKS BELOW.

type mockBase struct{}

func (mb *mockBase) Address() common.Address {
	return common.Address{}
}

type mockStateless struct {
	*mockBase
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 0
}

func (ms *mockStateless) Run(
	ctx context.Context, input []byte, caller common.Address,
	value *big.Int, readonly bool,
) ([]byte, error) {
	return nil, nil
}

type mockStateful struct {
	*mockBase
}

func (ms *mockStateful) ABIEvents() map[string]abi.Event {
	return map[string]abi.Event{
		"Event": {Name: "Event"},
	}
}

func (ms *mockStateful) CustomValueDecoders() map[string]cosmos.ValueDecoders {
	return map[string]cosmos.ValueDecoders{
		string("Event"): make(cosmos.ValueDecoders),
	}
}

func (ms *mockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput": solidity.MockPrecompileInterface.ABI.Methods["getOutput"],
	}
}

func (ms *mockStateful) PrecompileMethods() types.Methods {
	return types.Methods{
		{
			AbiSig:      "getOutput(string)",
			Execute:     getOutput,
			RequiredGas: 1,
		},
	}
}

type badMockStateful struct {
	*mockStateful
}

func (bms *badMockStateful) ABIMethods() map[string]abi.Method {
	return map[string]abi.Method{
		"getOutput":        solidity.MockPrecompileInterface.ABI.Methods["getOutput"],
		"getOutputPartial": solidity.MockPrecompileInterface.ABI.Methods["getOutputPartial"],
	}
}

type mockDynamic struct {
	*mockStateful
}

func (md *mockDynamic) Name() string {
	return ""
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
