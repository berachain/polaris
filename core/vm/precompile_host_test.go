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

package vm_test

import (
	"context"
	"errors"
	"math/big"
	"strconv"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	solidity "github.com/berachain/stargazer/testutil/contracts/solidity/generated"
	"github.com/berachain/stargazer/types/abi"

	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Precompile Host", func() {
	var pr *vm.PrecompileRegistry
	var ph *vm.PrecompileHost
	var psdb *mockPSDB

	BeforeEach(func() {
		pr = vm.NewPrecompileRegistry()
		err := pr.Register(&mockStateful{&mockBase{}})
		Expect(err).To(BeNil())
		psdb = &mockPSDB{}
		ph = vm.NewPrecompileHost(pr, psdb)
	})

	Describe("Test Exists", func() {
		It("should not return a container that doesn't exist", func() {
			pc, exists := ph.Exists(common.BytesToAddress([]byte{2}))
			Expect(exists).To(BeFalse())
			Expect(pc).To(BeNil())
		})

		It("should return a container that does exist", func() {
			pc, exists := ph.Exists(addr)
			Expect(exists).To(BeTrue())
			Expect(pc).ToNot(BeNil())
		})
	})

	Describe("Test Run", func() {
		It("should correctly run and build logs", func() {
			abiMethod := solidity.MockPrecompileInterface.ABI.Methods["getOutput"]
			inputs, err := abiMethod.Inputs.Pack("string")
			Expect(err).To(BeNil())
			pc, err := precompile.NewStatefulContainerFactory(precompile.NewLogRegistry()).Build(
				&mockStateful{&mockBase{}},
			)
			Expect(err).To(BeNil())
			_, gas, err := ph.Run(
				pc,
				append(abiMethod.ID, inputs...),
				addr, new(big.Int), 10, false,
			)
			Expect(err).To(BeNil())
			Expect(gas).To(Equal(uint64(9)))
			Expect(len(psdb.logs)).To(Equal(1))
			Expect(psdb.logs[0].Address).To(Equal(addr))
		})
	})

	Describe("Test Stateless Registration", func() {
		It("should properly register via the stateless container factory", func() {
			err := pr.Register(&mockStateless{})
			Expect(err).To(BeNil())
		})
	})
})

// MOCKS BELOW.

type mockPSDB struct {
	logs []*coretypes.Log
}

func (mp *mockPSDB) AddLog(log *coretypes.Log) {
	mp.logs = append(mp.logs, log)
}

func (mp *mockPSDB) GetContext() sdk.Context {
	return testutil.NewContextWithMultistores()
}

type mockBase struct{}

var addr = common.BytesToAddress([]byte{1})

func (mb *mockBase) Address() common.Address {
	return addr
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
		"CancelUnbondingDelegation": mockAbiEvent(),
	}
}

func (ms *mockStateful) CustomValueDecoders() map[precompile.EventType]log.ValueDecoders {
	return nil
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
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		"cancel_unbonding_delegation",
		sdk.NewAttribute("validator", common.EthAddressToValAddress(testutil.Alice).String()),
		sdk.NewAttribute("delegator", common.EthAddressToAccAddress(testutil.Bob).String()),
		sdk.NewAttribute("amount", "10bgt"),
		sdk.NewAttribute("creation_height", strconv.FormatInt(1, 10)),
	))
	return []any{
		[]mockObject{
			{
				CreationHeight: big.NewInt(1),
				TimeStamp:      str,
			},
		},
	}, nil
}

func mockAbiEvent() abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	int64Type, _ := abi.NewType("int64", "int64", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			abi.Argument{
				Name:    "validator",
				Type:    addrType,
				Indexed: true,
			},
			abi.Argument{
				Name:    "delegator",
				Type:    addrType,
				Indexed: true,
			},
			abi.Argument{
				Name:    "amount",
				Type:    uint256Type,
				Indexed: false,
			},
			abi.Argument{
				Name:    "creationHeight",
				Type:    int64Type,
				Indexed: false,
			},
		},
	)
}
