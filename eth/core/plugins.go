package core

import (
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/state"
)

type (
	ChainPlugin interface {
		BasePlugin
	}

	GasPlugin interface {
		BasePlugin
		ConsumeGas(amount uint64) error
		RefundGas(amount uint64)
		GasConsumed() uint64
		CumulativeGasUsed() uint64
	}

	StatePlugin = state.StatePlugin

	PrecompilePlugin interface {
		BasePlugin
		precompile.Runner
	}

	ConfigurationPlugin interface {
		BasePlugin
	}

	BasePlugin interface {
		Setup() error
	}
)
