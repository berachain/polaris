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
	gasMeter      storetypes.GasMeter
	blockGasMeter storetypes.GasMeter
}

// `NewPluginFrom` creates a new instance of the gas plugin from a given context.
func NewPluginFrom(ctx sdk.Context) core.GasPlugin {
	return &plugin{
		ctx.GasMeter(), ctx.BlockGasMeter(),
	}
}

// `Prepare` implements the core.GasPlugin interface.
func (p *plugin) Prepare(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	p.gasMeter = sCtx.GasMeter()
	p.blockGasMeter = sCtx.BlockGasMeter()
}

// `Reset` implements the core.GasPlugin interface.
func (p *plugin) Reset(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	p.gasMeter = sCtx.GasMeter()
	p.blockGasMeter = sCtx.BlockGasMeter()
}

// `SetGasLimit` resets the gas limit of the underlying GasMeter.
func (p *plugin) SetTxGasLimit(limit uint64) error {
	consumed := p.gasMeter.GasConsumed()
	// The gas meter is reset to the new limit.
	p.gasMeter = storetypes.NewGasMeter(limit)
	// Re-consume the gas that was already consumed.
	return p.TxConsumeGas(consumed)
}

// `BlockGasLimit` implements the core.GasPlugin interface.
func (p *plugin) BlockGasLimit() uint64 {
	return p.blockGasMeter.Limit()
}

// `TxConsumeGas` implements the core.GasPlugin interface.
func (p *plugin) TxConsumeGas(amount uint64) error {
	// We don't want to panic if we overflow so we do some safety checks.
	//nolint:gocritic // can't convert cleanly.
	if newConsumed, overflow := addUint64Overflow(p.TxGasUsed(), amount); overflow {
		return core.ErrGasUintOverflow
	} else if newConsumed > p.gasMeter.Limit() {
		return vm.ErrOutOfGas
	} else if newConsumed > p.blockGasMeter.Limit()-p.blockGasMeter.GasConsumed() {
		return core.ErrBlockOutOfGas
	}
	p.gasMeter.ConsumeGas(amount, gasMeterDescriptor)
	return nil
}

// `TxRefundGas` implements the core.GasPlugin interface.
func (p *plugin) TxRefundGas(amount uint64) {
	p.gasMeter.RefundGas(amount, gasMeterDescriptor)
}

// `TxGasRemaining` implements the core.GasPlugin interface.
func (p *plugin) TxGasRemaining() uint64 {
	return p.gasMeter.GasRemaining()
}

// `TxGasUsed` implements the core.GasPlugin interface.
func (p *plugin) TxGasUsed() uint64 {
	return p.gasMeter.GasConsumed()
}

// `CumulativeGasUsed` returns the cumulative gas used during the current block. If the cumulative
// gas used is greater than the block gas limit, we expect for Stargazer to handle it.
//
// `CumulativeGasUsed` implements the core.GasPlugin interface.
func (p *plugin) CumulativeGasUsed() uint64 {
	return p.TxGasUsed() + p.blockGasMeter.GasConsumed()
}

// `addUint64Overflow` performs the addition operation on two uint64 integers and returns a boolean
// on whether or not the result overflows.
func addUint64Overflow(a, b uint64) (uint64, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}
