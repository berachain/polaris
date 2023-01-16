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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/types/abi"
)

var _ = Describe("Precompile Registry", func() {
	var pr *vm.PrecompileRegistry
	var mc *mockStatefulContract
	var ms *mockStateless
	var mf *mockFactory
	var ctx sdk.Context
	var addr common.Address
	var moduleAddr common.Address

	BeforeEach(func() {
		pr = vm.NewPrecompileRegistry(testutil.EvmKey)
		ms = &mockStateless{}
		mc = &mockStatefulContract{mockStateless: ms}
		mf = &mockFactory{mockStatefulContract: mc}
		ctx = testutil.NewContextWithMultistores()
		addr = common.BytesToAddress([]byte{1})
		moduleAddr = common.BytesToAddress(authtypes.NewModuleAddress("test").Bytes())
	})

	Describe("Test Event Factory", func() {
		It("should load a non-nil event factory", func() {
			Expect(pr.GetEventFactory()).ToNot(BeNil())
		})

		It("should correctly register events for registered modules", func() {
			err := pr.RegisterModule("test", mc)
			Expect(err).To(BeNil())
			cosmosEvent := sdk.NewEvent("cosmos_event_type")
			log, err := pr.GetEventFactory().BuildLog(&cosmosEvent)
			Expect(err).To(BeNil())
			Expect(log.Address).To(Equal(moduleAddr))
		})
	})

	Describe("Test Stateless Precompiles", func() {
		It("should inject and get a stateless precompile", func() {
			pr.InjectStatelessContract(addr, ms)
			sc, found := pr.GetPrecompileFn(ctx)(addr)
			Expect(found).To(BeTrue())
			Expect(sc).ToNot(BeNil())
		})
	})

	Describe("Test Stateful Precompile", func() {
		It("should register and get a stateful precompile", func() {
			err := pr.RegisterModule("test", mc)
			Expect(err).To(BeNil())
			spc, found := pr.GetPrecompileFn(ctx)(moduleAddr)
			Expect(found).To(BeTrue())
			Expect(spc).ToNot(BeNil())
		})
	})

	Describe("Test Factory Precompile", func() {
		It("should inject and get a factory Precompile", func() {
			err := pr.InjectFactoryContract(ctx, addr, mf)
			Expect(err).To(BeNil())
			fc, found := pr.GetPrecompileFn(ctx)(addr)
			Expect(found).To(BeTrue())
			Expect(fc).ToNot(BeNil())
		})
	})
})

// Mocks below.

type mockStateless struct{}

func (*mockStateless) RequiredGas(input []byte) uint64 {
	return 0
}

func (*mockStateless) Run(input []byte) ([]byte, error) {
	return nil, nil
}

type mockStatefulContract struct {
	*mockStateless
}

func (mc *mockStatefulContract) ABIEvents() map[string]abi.Event {
	return map[string]abi.Event{
		"CosmosEventType": abi.NewEvent(
			"CosmosEventType",
			"CosmosEventType",
			false,
			nil,
		),
	}
}

func (mc *mockStatefulContract) GetFunctionsAndGas() precompile.FnsAndGas {
	return nil
}

type mockFactory struct {
	*mockStatefulContract
}

func (mf *mockFactory) Name() string {
	return "name"
}
