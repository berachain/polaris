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
	"errors"
	"math"

	"github.com/berachain/stargazer/eth/core"
	"github.com/cosmos/cosmos-sdk/store/types"
)

// `gasMeterDescriptor` is the descriptor for the gas meter used in the plugin.
const gasMeterDescriptor = `stargazer-gas-plugin`

// Compile-time interface assertions.
var _ core.GasPlugin = (*Plugin)(nil)

// `Plugin implements the core.GasPlugin interface.`.
type Plugin struct {
	gasMeter      types.GasMeter
	blockGasMeter types.GasMeter
}

// `NewPlugin` creates a new instance of the gas plugin.
func NewPlugin(gasMeter, blockGasMeter types.GasMeter) *Plugin {
	return &Plugin{
		gasMeter:      types.NewInfiniteGasMeter(),
		blockGasMeter: blockGasMeter,
	}
}

// `Setup` implements the core.GasPlugin interface.
func (p *Plugin) Setup() error {
	return nil
}

// `SetGasLimit` resets the gas limit of the underlying GasMeter.
func (p *Plugin) SetGasLimit(limit uint64) error {
	consumed := p.gasMeter.GasConsumed()
	// The gas meter is reset to the new limit.
	p.gasMeter = types.NewGasMeter(limit)
	// Re-consume the gas that was already consumed.
	return p.ConsumeGas(consumed)
}

// `ConsumeGas` implements the core.GasPlugin interface.
func (p *Plugin) ConsumeGas(amount uint64) error {
	// We don't want to panic if we overflow so we do some safety checks.
	if newConsumed, overflow := addUint64Overflow(p.gasMeter.GasConsumed(), amount); overflow {
		return core.ErrGasUintOverflow
	} else if newConsumed > p.gasMeter.Limit() {
		return errors.New("out of gas")
	}
	p.gasMeter.ConsumeGas(amount, gasMeterDescriptor)
	return nil
}

// `RefundGas` implements the core.GasPlugin interface.
func (p *Plugin) RefundGas(amount uint64) {
	p.gasMeter.RefundGas(amount, gasMeterDescriptor)
}

// `GasRemaining` implements the core.GasPlugin interface.
func (p *Plugin) GasRemaining() uint64 {
	return p.gasMeter.GasRemaining()
}

// `GasUsed` implements the core.GasPlugin interface.
func (p *Plugin) GasUsed() uint64 {
	return p.gasMeter.GasConsumed()
}

// `CumulativeGasUsed` implements the core.GasPlugin interface.
func (p *Plugin) CumulativeGasUsed() uint64 {
	used := p.gasMeter.GasConsumed()
	limit := p.blockGasMeter.Limit()
	used += p.blockGasMeter.GasConsumed()
	if used > limit {
		used = limit
	}

	return used
}

// addUint64Overflow performs the addition operation on two uint64 integers and
// returns a boolean on whether or not the result overflows.
func addUint64Overflow(a, b uint64) (uint64, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}
