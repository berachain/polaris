// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package precompile

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/store/snapmulti"
	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state/events/mock"
	ethstate "github.com/berachain/polaris/eth/core/state"
	pvm "github.com/berachain/polaris/eth/core/vm"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/precompile")
}

var _ = Describe("plugin", func() {
	var p *plugin
	var e vm.PrecompileEVM
	var ctx sdk.Context

	BeforeEach(func() {
		ctx = testutil.NewContext(log.NewTestLogger(GinkgoT()))
		ctx = ctx.WithEventManager(
			events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory()),
		)
		ctx = ctx.WithMultiStore(snapmulti.NewStoreFrom(ctx.MultiStore()))
		p = utils.MustGetAs[*plugin](NewPlugin())
		e = &mockEVM{nil, ctx, &mockSDB{nil, ctx, 0}}
	})

	It("should use correctly consume gas", func() {
		_, remainingGas, err := p.Run(e, &mockStateless{}, []byte{}, addr, new(big.Int), 30, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(remainingGas).To(Equal(uint64(10)))
	})

	It("should error on insufficient gas", func() {
		_, _, err := p.Run(e, &mockStateless{}, []byte{}, addr, new(big.Int), 5, false)
		Expect(err).To(MatchError("out of gas"))
	})

	It("should handle read-only static calls", func() {
		ms := utils.MustGetAs[MultiStore](ctx.MultiStore())
		cem := utils.MustGetAs[state.ControllableEventManager](ctx.EventManager())
		// verify its not read-only right now
		Expect(ms.IsReadOnly()).To(BeFalse())
		Expect(cem.IsReadOnly()).To(BeFalse())

		// run read only precompile
		_, _, err := p.Run(e, &mockStateful{}, []byte{2}, addr2, new(big.Int), 5, true)
		Expect(err.Error()).To(ContainSubstring(vm.ErrWriteProtection.Error()))
		_, _, err = p.Run(e, &mockStateful{}, []byte{3}, addr2, new(big.Int), 5, true)
		Expect(err.Error()).To(ContainSubstring(vm.ErrWriteProtection.Error()))

		// check that the multistore and event manager is set back to read-only false
		Expect(ms.IsReadOnly()).To(BeFalse())
		Expect(cem.IsReadOnly()).To(BeFalse())
	})

	It("should catch panics and return a geth error type", func() {
		_, remainingGas, err := p.Run(e,
			&mockPanicking{err: storetypes.ErrorNegativeGasConsumed{Descriptor: "henlo"}},
			[]byte{}, addr, new(big.Int), 30, false)
		Expect(err).To(MatchError(vm.ErrOutOfGas.Error()))
		Expect(remainingGas).To(Equal(uint64(0)))
	})

	It("should catch panics and propagate", func() {
		_, _, vmErr := p.Run(e, &mockPanicking{
			err: errors.New("error"),
		}, []byte{}, addr, new(big.Int), 30, false)
		Expect(errors.Is(vmErr, vm.ErrExecutionReverted)).To(BeTrue())
	})
})

var (
	addr  = common.BytesToAddress([]byte{1})
	addr2 = common.BytesToAddress([]byte{2})
)

type mockEVM struct {
	vm.PrecompileEVM
	ctx sdk.Context
	ms  *mockSDB
}

func (me *mockEVM) GetStateDB() vm.StateDB {
	return me.ms
}

type mockSDB struct {
	pvm.PolarStateDB
	ctx  sdk.Context
	logs int
}

func (ms *mockSDB) GetPlugin() ethstate.Plugin {
	return state.NewPlugin(
		nil, nil, nil, nil,
	)
}

func (ms *mockSDB) GetContext() context.Context {
	return ms.ctx
}

func (ms *mockSDB) AddLog(*ethtypes.Log) {
	ms.logs++
}

type mockStateless struct{} // at addr 1

func (ms *mockStateless) RegistryKey() common.Address {
	return addr
}

func (ms *mockStateless) Run(
	ctx context.Context, _ vm.PrecompileEVM, _ []byte,
	_ common.Address, _ *big.Int,
) ([]byte, error) {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(10, "")
	return nil, nil
}

func (ms *mockStateless) RequiredGas(_ []byte) uint64 {
	return 10
}

type mockStateful struct{} // at addr 2

func (msf *mockStateful) RegistryKey() common.Address {
	return addr
}

// panics if modifying state on read-only.
func (msf *mockStateful) Run(
	ctx context.Context, _ vm.PrecompileEVM, input []byte,
	_ common.Address, _ *big.Int,
) ([]byte, error) {
	if input[0] == byte(2) {
		panic(vm.ErrWriteProtection)
	} else if input[0] == byte(3) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		sdkCtx.EventManager().EmitEvent(sdk.NewEvent("test"))
	}
	return nil, nil
}

func (msf *mockStateful) RequiredGas(_ []byte) uint64 {
	return 1
}

type mockPanicking struct {
	err any
} // at addr 1

func (mp *mockPanicking) RegistryKey() common.Address {
	return addr
}

func (mp *mockPanicking) Run(
	_ context.Context, _ vm.PrecompileEVM, _ []byte,
	_ common.Address, _ *big.Int,
) ([]byte, error) {
	panic(mp.err)
}

func (*mockPanicking) RequiredGas(_ []byte) uint64 {
	return 1
}
