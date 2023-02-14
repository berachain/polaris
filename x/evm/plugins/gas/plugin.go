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
