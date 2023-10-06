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
	"fmt"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	ethstate "pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/registry"
	libtypes "pkg.berachain.dev/polaris/lib/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Plugin is the interface that must be implemented by the plugin.
type Plugin interface {
	core.PrecompilePlugin
	RegisterPrecompiles([]ethprecompile.Registrable) error
}

// PolarStateDB is the interface that must be implemented by the state DB.
// The stateDB must allow retrieving the plugin in order to set it's gas config.
type PolarStateDB interface {
	// GetPlugin retrieves the underlying state plugin from the StateDB.
	GetPlugin() ethstate.Plugin
}

// plugin runs precompile containers in the Cosmos environment with the context gas configs.
type plugin struct {
	libtypes.Registry[common.Address, vm.PrecompiledContract]
	// kvGasConfig is the gas config for the KV store.
	kvGasConfig storetypes.GasConfig
	// transientKVGasConfig is the gas config for the transient KV store.
	transientKVGasConfig storetypes.GasConfig
}

// NewPlugin creates and returns a plugin with the default KV store gas configs.
func NewPlugin() Plugin {
	return &plugin{
		Registry: registry.NewMap[common.Address, vm.PrecompiledContract](),
		// NOTE: these are hardcoded as they are also hardcoded in the sdk.
		// This should be updated if it ever changes.
		kvGasConfig:          storetypes.KVGasConfig(),
		transientKVGasConfig: storetypes.TransientGasConfig(),
	}
}

func (p *plugin) Get(addr common.Address, _ *params.Rules) (vm.PrecompiledContract, bool) {
	// TODO: handle rules
	val := p.Registry.Get(addr)
	if val == nil {
		return nil, false
	}
	return val, true
}

func (p *plugin) RegisterPrecompiles(precompiles []ethprecompile.Registrable) error {
	for _, pc := range precompiles {
		// choose the appropriate precompile factory
		var af ethprecompile.AbstractFactory
		switch {
		case utils.Implements[ethprecompile.StatefulImpl](pc):
			af = ethprecompile.NewStatefulFactory()
		case utils.Implements[ethprecompile.StatelessImpl](pc):
			af = ethprecompile.NewStatelessFactory()
		default:
			return fmt.Errorf("unknown precompile type %T", pc)
		}
		// build the precompile container and register with the plugin
		container, err := af.Build(pc, p)
		if err != nil {
			return err
		}

		err = p.Register(container)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetActive implements core.PrecompilePlugin.
func (p *plugin) GetActive(_ params.Rules) []common.Address {
	// TODO: enable hardfork activation and de-activation.
	active := make([]common.Address, 0)
	for k := range p.Registry.Iterate() {
		active = append(active, k)
	}
	return active
}

// Run runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if the precompile execution returns an
// error or insufficient gas is provided.
//
// Run implements core.PrecompilePlugin.
//
//nolint:nonamedreturns // panic recovery.
func (p *plugin) Run(
	evm vm.PrecompileEVM, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readOnly bool,
) (ret []byte, gasRemaining uint64, err error) {
	// get native Cosmos SDK context, MultiStore, and EventManager from the Polaris StateDB
	sdb := utils.MustGetAs[vm.PolarStateDB](evm.GetStateDB())
	ctx := sdk.UnwrapSDKContext(sdb.GetContext())
	ms := utils.MustGetAs[MultiStore](ctx.MultiStore())
	cem := utils.MustGetAs[state.ControllableEventManager](ctx.EventManager())

	requiredGas := pc.RequiredGas(input)
	// handle edge case when not enough gas is provided for even the required gas
	if requiredGas > suppliedGas {
		return nil, 0, vm.ErrOutOfGas
	}

	// make sure the readOnly is only set if we aren't in readOnly yet, which also makes sure that
	// the readOnly flag isn't removed for child calls (taken from geth core/vm/interepreter.go)
	if readOnly && !ms.IsReadOnly() {
		cem.SetReadOnly(true)
		ms.SetReadOnly(true)
		defer func() {
			cem.SetReadOnly(false)
			ms.SetReadOnly(false)
		}()
	}

	// disable reentrancy into the EVM only during precompile execution
	p.disableReentrancy(sdb)
	defer p.enableReentrancy(sdb)

	// recover from any WriteProtection or OutOfGas panic for the EVM to handle as a vm error
	defer RecoveryHandler(&err)

	// use a precompile-specific gas meter for dynamic consumption
	gm := storetypes.NewGasMeter(suppliedGas)
	gm.ConsumeGas(requiredGas, "precompile required gas")

	// run the precompile container
	ret, err = pc.Run(
		ctx.WithGasMeter(gm).
			WithKVGasConfig(p.kvGasConfig).
			WithTransientKVGasConfig(p.transientKVGasConfig),
		evm,
		input,
		caller,
		value,
	)
	gasRemaining = gm.GasRemaining()

	return //nolint:nakedret // named returns.
}

// EnableReentrancy sets the state so that execution can enter the EVM again.
//
// EnableReentrancy implements core.PrecompilePlugin.
func (p *plugin) EnableReentrancy(evm vm.PrecompileEVM) {
	p.enableReentrancy(utils.MustGetAs[vm.PolarStateDB](evm.GetStateDB()))
}

func (p *plugin) enableReentrancy(sdb vm.PolarStateDB) {
	sdkCtx := sdk.UnwrapSDKContext(sdb.GetContext())

	// end precompile execution => stop emitting Cosmos event as Eth logs for now
	cem := utils.MustGetAs[state.ControllableEventManager](sdkCtx.EventManager())
	cem.EndPrecompileExecution()

	// remove Cosmos gas consumption so gas is consumed only per OPCODE
	utils.MustGetAs[state.Plugin](
		utils.MustGetAs[PolarStateDB](sdb).GetPlugin(),
	).SetGasConfig(storetypes.GasConfig{}, storetypes.GasConfig{})
}

// DisableReentrancy sets the state so that execution cannot enter the EVM again.
//
// DisableReentrancy implements core.PrecompilePlugin.
func (p *plugin) DisableReentrancy(evm vm.PrecompileEVM) {
	p.disableReentrancy(utils.MustGetAs[vm.PolarStateDB](evm.GetStateDB()))
}

func (p *plugin) disableReentrancy(sdb vm.PolarStateDB) {
	sdkCtx := sdk.UnwrapSDKContext(sdb.GetContext())

	// resume precompile execution => begin emitting Cosmos event as Eth logs again
	cem := utils.MustGetAs[state.ControllableEventManager](sdkCtx.EventManager())
	cem.BeginPrecompileExecution(sdb)

	// restore ctx gas configs for continuing precompile execution
	utils.MustGetAs[state.Plugin](
		utils.MustGetAs[PolarStateDB](sdb).GetPlugin(),
	).SetGasConfig(p.kvGasConfig, p.transientKVGasConfig)
}
