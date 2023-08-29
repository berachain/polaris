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

package precompile

import (
	"context"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmock "pkg.berachain.dev/polaris/cosmos/testing/types/mock"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events/mock"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin", func() {
	var p *plugin
	var e vm.PrecompileEVM
	var ctx sdk.Context

	BeforeEach(func() {
		ctx = testutil.NewContext()
		ctx = ctx.WithEventManager(
			events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory()),
		)
		p = utils.MustGetAs[*plugin](NewPlugin(nil, &mockSP{ctx}))
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

	It("should plug in custom gas configs", func() {
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(1000)))
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(100)))

		p.SetKVGasConfig(storetypes.GasConfig{
			DeleteCost: 2,
		})
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(2)))
		p.SetTransientKVGasConfig(storetypes.GasConfig{
			DeleteCost: 3,
		})
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(3)))
	})

	It("should handle read-only static calls", func() {
		ms := utils.MustGetAs[tmock.MultiStore](ctx.MultiStore())
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
})

var (
	addr  = common.BytesToAddress([]byte{1})
	addr2 = common.BytesToAddress([]byte{2})
)

// MOCKS BELOW.

type mockSP struct {
	ctx sdk.Context
}

func (msp *mockSP) SetGasConfig(kvg storetypes.GasConfig, tkvg storetypes.GasConfig) {
	msp.ctx = msp.ctx.WithKVGasConfig(kvg).WithTransientKVGasConfig(tkvg)
}

type mockEVM struct {
	vm.PrecompileEVM
	ctx sdk.Context
	ms  *mockSDB
}

func (me *mockEVM) GetStateDB() vm.GethStateDB {
	return me.ms
}

type mockSDB struct {
	vm.PolarisStateDB
	ctx  sdk.Context
	logs int
}

func (ms *mockSDB) GetContext() context.Context {
	return ms.ctx
}

func (ms *mockSDB) AddLog(*coretypes.Log) {
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
