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

package precompile_test

import (
	"context"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	"github.com/berachain/stargazer/types/abi"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Factories", func() {
	var lr *precompile.LogRegistry

	Context("Stateless Container Factory", func() {
		var scf *precompile.StatelessContainerFactory

		BeforeEach(func() {
			scf = precompile.NewStatelessContainerFactory()
		})

		It("should build stateless precompile containers", func() {
			pc, err := scf.Build(&mockStateless{&mockBase{}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockBase{})
			Expect(err.Error()).To(Equal("StatelessContainerImpl: this precompile contract implementation is not implemented"))
		})
	})

	Context("Stateful Container Factory", func() {
		var scf *precompile.StatefulContainerFactory

		BeforeEach(func() {
			lr = precompile.NewLogRegistry()
			scf = precompile.NewStatefulContainerFactory(lr)
		})

		It("should correctly build stateful containers and log events", func() {
			pc, err := scf.Build(&mockStateful{&mockBase{}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = scf.Build(&mockStateless{&mockBase{}})
			Expect(err.Error()).To(Equal("StatefulContainerImpl: this precompile contract implementation is not implemented"))
		})
	})

	Context("Dynamic Container Factory", func() {
		var dcf *precompile.DynamicContainerFactory

		BeforeEach(func() {
			lr = precompile.NewLogRegistry()
			dcf = precompile.NewDynamicContainerFactory(lr)
		})

		It("should properly build dynamic container", func() {
			pc, err := dcf.Build(&mockDynamic{&mockStateful{&mockBase{}}})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())

			_, err = dcf.Build(&mockStateful{&mockBase{}})
			Expect(err.Error()).To(Equal("DynamicContainerImpl: this precompile contract implementation is not implemented"))
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

func (ms *mockStateful) CustomValueDecoders() map[precompile.EventType]log.ValueDecoders {
	return map[precompile.EventType]log.ValueDecoders{
		precompile.EventType("Event"): make(log.ValueDecoders),
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
	ctx sdk.Context,
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
