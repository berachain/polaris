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

package gas

import (
	"context"
	"math"

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/vm"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `gasMeterDescriptor` is the descriptor for the gas meter used in the plugin.
const gasMeterDescriptor = `stargazer-gas-plugin`

// `plugin` wraps a Cosmos context and utilize's the underlying `GasMeter` and `BlockGasMeter`
// to implement the core.GasPlugin interface.
type plugin struct {
	sdk.Context
}

// `NewPluginFrom` creates a new instance of the gas plugin from a given context.
func NewPluginFrom(ctx sdk.Context) core.GasPlugin {
	return &plugin{
		Context: ctx,
	}
}

// `Reset` implements the core.GasPlugin interface.
func (p *plugin) Reset(ctx context.Context) {
	p.Context = sdk.UnwrapSDKContext(ctx)
}

// `SetGasLimit` resets the gas limit of the underlying GasMeter.
func (p *plugin) SetGasLimit(limit uint64) error {
	consumed := p.GasMeter().GasConsumed()
	// The gas meter is reset to the new limit.
	p.Context = p.WithGasMeter(storetypes.NewGasMeter(limit))
	// Re-consume the gas that was already consumed.
	return p.ConsumeGas(consumed)
}

// `ConsumeGas` implements the core.GasPlugin interface.
func (p *plugin) ConsumeGas(amount uint64) error {
	// We don't want to panic if we overflow so we do some safety checks.
	if newConsumed, overflow := addUint64Overflow(p.GasMeter().GasConsumed(), amount); overflow {
		return core.ErrGasUintOverflow
	} else if newConsumed > p.GasMeter().Limit() {
		return vm.ErrOutOfGas
	}
	p.GasMeter().ConsumeGas(amount, gasMeterDescriptor)
	return nil
}

// `RefundGas` implements the core.GasPlugin interface.
func (p *plugin) RefundGas(amount uint64) {
	p.GasMeter().RefundGas(amount, gasMeterDescriptor)
}

// `GasRemaining` implements the core.GasPlugin interface.
func (p *plugin) GasRemaining() uint64 {
	return p.GasMeter().GasRemaining()
}

// `GasUsed` implements the core.GasPlugin interface.
func (p *plugin) GasUsed() uint64 {
	return p.GasMeter().GasConsumed()
}

// `CumulativeGasUsed` returns the cumulative gas used during the current block. If the cumulative
// gas used is greater than the block gas limit, it returns the block gas limit, but the tx will
// still fail in `runTx`.
//
// `CumulativeGasUsed` implements the core.GasPlugin interface.
func (p *plugin) CumulativeGasUsed() uint64 {
	used := p.GasUsed() + p.BlockGasMeter().GasConsumed()
	if limit := p.BlockGasMeter().Limit(); used > limit {
		used = limit
	}
	return used
}

// `addUint64Overflow` performs the addition operation on two uint64 integers and returns a boolean
// on whether or not the result overflows.
func addUint64Overflow(a, b uint64) (uint64, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}
