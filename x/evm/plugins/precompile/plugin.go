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
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/lib/registry"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/x/evm/plugins"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state"
)

// `Plugin` is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosStargazer
	core.PrecompilePlugin
}

// `plugin` runs precompile containers in the Cosmos environment with the context gas configs.
type plugin struct {
	sdk.Context
	libtypes.Registry[common.Address, vm.PrecompileContainer]

	// `getPrecompiles` returns all supported precompile contracts.
	getPrecompiles func() []vm.RegistrablePrecompile
}

// `NewPlugin` creates and returns a `plugin` with the given precompile getter function.
func NewPlugin(getPrecompiles func() []vm.RegistrablePrecompile) Plugin {
	return &plugin{
		Registry:       registry.NewMap[common.Address, vm.PrecompileContainer](),
		getPrecompiles: getPrecompiles,
	}
}

// `Reset` implements `core.PrecompilePlugin`.
func (p *plugin) Reset(ctx context.Context) {
	p.Context = sdk.UnwrapSDKContext(ctx)
}

// `GetPrecompiles` implements `core.PrecompilePlugin`.
func (p *plugin) GetPrecompiles(_ *params.Rules) []vm.RegistrablePrecompile {
	return p.getPrecompiles()
}

// `Run` runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if the precompile execution returns an
// error or insufficient gas is provided.
//
// `Run` implements `core.PrecompilePlugin`.
func (p *plugin) Run(
	sdb vm.GethStateDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// use a precompile-specific gas meter for dynamic consumption
	gm := storetypes.NewInfiniteGasMeter()
	// consume static gas from RequiredGas
	gm.ConsumeGas(pc.RequiredGas(input), "RequiredGas")

	// begin precompile execution => begin emitting Cosmos event as Eth logs
	cem := utils.MustGetAs[state.ControllableEventManager](p.Context.EventManager())
	cem.BeginPrecompileExecution(sdb)

	// run precompile container
	ret, err := pc.Run(
		p.Context.WithGasMeter(gm),
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

	// valid precompile gas consumption => return supplied gas
	return ret, suppliedGas - gm.GasConsumed(), err
}
