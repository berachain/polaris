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
	"github.com/berachain/stargazer/x/evm/plugins/state"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
)

// Compile-time assertion to ensure `CosmosRunner` adheres to `vm.PrecompileRunner`.
var _ precompile.Runner = (*CosmosRunner)(nil)

// `CosmosRunner` runs precompile containers in a Cosmos environment with the given gas configs.
type CosmosRunner struct {
	// `kvGasConfig` is the gas config for execution of kv store operations in native precompiles.
	kvGasConfig *storetypes.GasConfig
	// `transientKVGasConfig` is the gas config for execution transient kv store operations in
	// native precompiles.
	transientKVGasConfig *storetypes.GasConfig
}

// `NewCosmosRunner` creates and returns a `CosmosRunner` with the SDK default gas configs.
func NewCosmosRunner() *CosmosRunner {
	defaultKVGasConfig := storetypes.KVGasConfig()
	defaultTransientKVGasConfig := storetypes.TransientGasConfig()

	return &CosmosRunner{
		kvGasConfig:          &defaultKVGasConfig,
		transientKVGasConfig: &defaultTransientKVGasConfig,
	}
}

// `KVGasConfig` returns the `CosmosRunner`'s `kvGasConfig`.
func (cr *CosmosRunner) KVGasConfig() *storetypes.GasConfig {
	return cr.kvGasConfig
}

// `TransientKVGasConfig` returns the `CosmosRunner`'s `transientKVGasConfig`.
func (cr *CosmosRunner) TransientKVGasConfig() *storetypes.GasConfig {
	return cr.transientKVGasConfig
}

// `SetKVGasConfig` sets the `CosmosRunner` to have `kvGasConfig`.
func (cr *CosmosRunner) SetKVGasConfig(kvGasConfig *storetypes.GasConfig) {
	cr.kvGasConfig = kvGasConfig
}

// `SetTransientKVGasConfig` sets the `CosmosRunner` to have `transientKVGasConfig`.
func (cr *CosmosRunner) SetTransientKVGasConfig(transientKVGasConfig *storetypes.GasConfig) {
	cr.transientKVGasConfig = transientKVGasConfig
}

// `Run` runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if insufficient gas is provided or the
// precompile execution returns an error.
//
// `Run` implements `vm.PrecompileRunner`.
func (cr *CosmosRunner) Run(
	ctx context.Context, ldb precompile.LogsDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// use a precompile-specific gas meter for dynamic consumption
	gm := storetypes.NewInfiniteGasMeter()
	// consume static gas from RequiredGas
	gm.ConsumeGas(pc.RequiredGas(input), "RequiredGas")

	// begin precompile execution => begin emitting Cosmos event as Eth logs
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	cem := utils.MustGetAs[state.ControllableEventManager](sdkCtx.EventManager()) // TODO: okay to panic here?
	cem.BeginPrecompileExecution(ldb)

	// run precompile container
	ret, err := pc.Run(
		sdkCtx.
			WithGasMeter(gm).
			WithKVGasConfig(*cr.kvGasConfig).
			WithTransientKVGasConfig(*cr.transientKVGasConfig),
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

	return ret, suppliedGas - gm.GasConsumed(), err
}
