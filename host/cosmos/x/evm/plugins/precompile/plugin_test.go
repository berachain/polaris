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

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state/events/mock"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx sdk.Context

var _ = Describe("plugin", func() {
	var p *plugin
	var sdb *mockSDB

	BeforeEach(func() {
		ctx = testutil.NewContext()
		ctx = ctx.WithEventManager(
			events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory()),
		)
		p = utils.MustGetAs[*plugin](NewPlugin(nil))
		sdb = &mockSDB{}
	})

	It("should use correctly consume gas", func() {
		_, remainingGas, err := p.Run(sdb, &mockStateless{}, []byte{}, addr, new(big.Int), 30, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(remainingGas).To(Equal(uint64(10)))
	})

	It("should error on insufficient gas", func() {
		_, _, err := p.Run(sdb, &mockStateless{}, []byte{}, addr, new(big.Int), 5, true)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should plug in custom gas configs", func() {
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(0)))
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(0)))

		p.SetKVGasConfig(storetypes.KVGasConfig())
		Expect(p.KVGasConfig().DeleteCost).To(Equal(uint64(1000)))
		p.SetTransientKVGasConfig(storetypes.TransientGasConfig())
		Expect(p.TransientKVGasConfig().DeleteCost).To(Equal(uint64(100)))
	})
})

// MOCKS BELOW.

type mockSDB struct {
	vm.PolarisStateDB
}

func (ms *mockSDB) GetContext() context.Context {
	return ctx
}

type mockStateless struct{}

var addr = common.BytesToAddress([]byte{1})

func (ms *mockStateless) RegistryKey() common.Address {
	return addr
}

func (ms *mockStateless) Run(
	ctx context.Context, input []byte,
	caller common.Address, value *big.Int, readonly bool,
) ([]byte, error) {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(10, "")
	return nil, nil
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 10
}

func (ms *mockStateless) WithStateDB(vm.GethStateDB) vm.PrecompileContainer {
	return ms
}
