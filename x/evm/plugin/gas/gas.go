package gas

import (
	"github.com/berachain/stargazer/eth/core"
	"github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time interface assertions.
var _ core.GasPlugin = (*Plugin)(nil)

// `Plugin implements the core.GasPlugin interface.`
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
