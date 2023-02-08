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
	"github.com/berachain/stargazer/eth/core"
	"github.com/cosmos/cosmos-sdk/store/types"
)

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
		gasMeter:      gasMeter,
		blockGasMeter: blockGasMeter,
	}
}

// `Setup` implements the core.GasPlugin interface.
func (p *Plugin) Setup() error {
	return nil
}

// `ConsumeGas` implements the core.GasPlugin interface.
func (p *Plugin) ConsumeGas(amount uint64) error {
	p.gasMeter.ConsumeGas(amount, "stargazer-gas-plugin")
	return nil
}

// `RefundGas` implements the core.GasPlugin interface.
func (p *Plugin) RefundGas(amount uint64) {
	p.gasMeter.RefundGas(amount, "stargazer-gas-plugin")
}

// `GasConsumed` implements the core.GasPlugin interface.
func (p *Plugin) GasConsumed() uint64 {
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
