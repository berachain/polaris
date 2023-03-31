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
	"errors"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/registry"
	libtypes "pkg.berachain.dev/polaris/lib/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Plugin is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosPolaris
	core.PrecompilePlugin

	KVGasConfig() storetypes.GasConfig
	SetKVGasConfig(storetypes.GasConfig)
	TransientKVGasConfig() storetypes.GasConfig
	SetTransientKVGasConfig(storetypes.GasConfig)
}

// nativeContext is only active during precompile (native host chain) execution.
type nativeContext struct {
	inactive bool // inactive only during EVM reentrancy
	ctx      sdk.Context
	nativeGM storetypes.GasMeter
}

// plugin runs precompile containers in the Cosmos environment with the context gas configs.
type plugin struct {
	libtypes.Registry[common.Address, vm.PrecompileContainer]
	// precompiles is all supported precompile contracts.
	precompiles []precompile.Registrable
	// kvGasConfig is the gas config for the KV store.
	kvGasConfig storetypes.GasConfig
	// transientKVGasConfig is the gas config for the transient KV store.
	transientKVGasConfig storetypes.GasConfig
	// sp allows resetting the context for the reentrancy into the EVM.
	sp StatePlugin
	// ephemeral is the native context active during precompile execution.
	ephemeral *nativeContext
}

// NewPlugin creates and returns a `plugin` with the default kv gas configs.
func NewPlugin(precompiles []precompile.Registrable, sp StatePlugin) Plugin {
	return &plugin{
		Registry:    registry.NewMap[common.Address, vm.PrecompileContainer](),
		precompiles: precompiles,
		// TODO: Re-enable gas config for precompiles.
		// https://github.com/berachain/polaris/issues/393
		kvGasConfig:          storetypes.GasConfig{},
		transientKVGasConfig: storetypes.GasConfig{},
		sp:                   sp,
	}
}

// GetPrecompiles implements `core.PrecompilePlugin`.
func (p *plugin) GetPrecompiles(_ *params.Rules) []precompile.Registrable {
	return p.precompiles
}

func (p *plugin) KVGasConfig() storetypes.GasConfig {
	return p.kvGasConfig
}

func (p *plugin) SetKVGasConfig(kvGasConfig storetypes.GasConfig) {
	p.kvGasConfig = kvGasConfig
}

func (p *plugin) TransientKVGasConfig() storetypes.GasConfig {
	return p.transientKVGasConfig
}

func (p *plugin) SetTransientKVGasConfig(transientKVGasConfig storetypes.GasConfig) {
	p.transientKVGasConfig = transientKVGasConfig
}

// Run runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if the precompile execution returns an
// error or insufficient gas is provided.
//
// Run implements `core.PrecompilePlugin`.
func (p *plugin) Run(
	evm precompile.EVM, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// only run precompile container if ephemeral ctx is nil
	if p.ephemeral != nil {
		return nil, suppliedGas, errors.New("precompile container is already running")
	}
	p.ephemeral = &nativeContext{} // set as active because we are entering native host chain execution

	// use a precompile-specific gas meter for dynamic consumption
	gm := storetypes.NewInfiniteGasMeter()
	// consume static gas from RequiredGas
	gm.ConsumeGas(pc.RequiredGas(input), "RequiredGas")

	// get native Cosmos SDK context from the Polaris StateDB
	sdb := utils.MustGetAs[vm.PolarisStateDB](evm.GetStateDB())
	p.ephemeral.ctx = sdk.UnwrapSDKContext(sdb.GetContext())
	p.ephemeral.nativeGM = p.ephemeral.ctx.GasMeter()

	// begin precompile execution => begin emitting Cosmos event as Eth logs
	cem := utils.MustGetAs[state.ControllableEventManager](p.ephemeral.ctx.EventManager())
	cem.BeginPrecompileExecution(sdb)

	// run precompile container
	ret, err := pc.Run(
		p.ephemeral.ctx.WithGasMeter(gm).
			WithKVGasConfig(p.kvGasConfig).
			WithTransientKVGasConfig(p.transientKVGasConfig),
		evm,
		input,
		caller,
		value,
		readonly,
	)

	// end precompile execution => stop emitting Cosmos event as Eth logs
	cem.EndPrecompileExecution()

	// handle overconsumption of gas
	if gm.GasConsumed() > suppliedGas {
		return nil, 0, vm.ErrOutOfGas
	}

	// clear ephemeral context so that the next precompile container can be run
	p.ephemeral = nil

	// valid precompile gas consumption => return supplied gas
	return ret, suppliedGas - gm.GasConsumed(), err
}

// EnableReentrancy sets the state so that execution can enter the EVM again.
//
// EnableReentrancy implements `core.PrecompilePlugin`.
func (p *plugin) EnableReentrancy() {
	if p.ephemeral == nil || p.ephemeral.inactive {
		// not in precompile execution, cannot (re)enable reentrancy
		return
	}

	// reset the state plugin to the current context so that the next EVM execution can continue
	// normally
	p.sp.Reset(
		sdk.NewContext(
			p.ephemeral.ctx.MultiStore(),
			p.ephemeral.ctx.BlockHeader(),
			p.ephemeral.ctx.IsCheckTx(),
			p.ephemeral.ctx.Logger(),
		).WithGasMeter(p.ephemeral.nativeGM),
	)

	// native code is no longer active, EVM takes over
	p.ephemeral.inactive = true
}

// DisableReentrancy sets the state so that execution cannot enter the EVM again.
//
// DisableReentrancy implements `core.PrecompilePlugin`.
func (p *plugin) DisableReentrancy() {
	if p.ephemeral == nil || !p.ephemeral.inactive {
		// in precompile execution, so cannot (re)disable reentrancy
		return
	}

	// EVM is done, native code is now active again
	p.ephemeral.inactive = false
}
