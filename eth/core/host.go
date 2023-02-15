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

package core

import (
	"context"

	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/params"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `StargazerHostChain` defines the plugins that the chain running Stargazer EVM should implement.
type StargazerHostChain interface {
	// `GetBlockPlugin` returns the `BlockPlugin` of the Stargazer host chain.
	GetBlockPlugin() BlockPlugin
	// `GetConfigurationPlugin` returns the `ConfigurationPlugin` of the Stargazer host chain.
	GetConfigurationPlugin() ConfigurationPlugin
	// `GetGasPlugin` returns the `GasPlugin` of the Stargazer host chain.
	GetGasPlugin() GasPlugin
	// `GetPrecompilePlugin` returns the OPTIONAL `PrecompilePlugin` of the Stargazer host chain.
	GetPrecompilePlugin() PrecompilePlugin
	// `GetStatePlugin` returns the `StatePlugin` of the Stargazer host chain.
	GetStatePlugin() StatePlugin
}

// =============================================================================
// Mandatory Plugins
// =============================================================================

// The following plugins should be implemented by the chain running Stargazer EVM and exposed via
// the `StargazerHostChain` interface. All plugins should be resettable with a given context.
type (
	// `BlockPlugin` defines the methods that the chain running Stargazer EVM should implement to
	// support the `BlockPlugin` interface.
	BlockPlugin interface {
		// `BlockPlugin` implements `libtypes.Resettable`. Calling `Reset` should reset the
		// `BlockPlugin` to a default state.
		libtypes.Resettable
		// `GetStargazerHeaderAtHeight` returns the block header at the given block height.
		GetStargazerHeaderAtHeight(context.Context, uint64) *types.StargazerHeader
		// `BaseFee` returns the base fee of the current block.
		BaseFee() uint64
	}

	// `GasPlugin` is an interface that allows the Stargazer EVM to consume gas on the host chain.
	GasPlugin interface {
		// `GasPlugin` implements `libtypes.Resettable`. Calling `Reset` should reset the
		// `GasPlugin` to a default state.
		libtypes.Resettable
		// `ConsumeGas` consumes the supplied amount of gas. It should not panic due to a
		// `GasOverflow` and should return `core.ErrOutOfGas` if the amount of gas remaining is
		// less than the amount requested.
		ConsumeGas(uint64) error
		// `RefundGas` refunds the supplied amount of gas. It should not panic.
		RefundGas(uint64)
		// `GasRemaining` returns the amount of gas remaining. It should not panic.
		GasRemaining() uint64
		// `GasUsed` returns the amount of gas used during the current transaction. It should not
		// panic.
		GasUsed() uint64
		// `CumulativeGasUsed` returnsthe amount of gas used during the current block. The value
		// returned should include any gas consumed during this transaction. It should not panic.
		CumulativeGasUsed() uint64
		// `MaxFeePerGas` should set the maximum amount of gas that can be consumed by the meter.
		// It should not panic, but instead, return an error, if the new gas limit is less than the
		// currently consumed amount of gas.
		SetGasLimit(uint64) error
	}

	// `StatePlugin` defines the methods that the chain running Stargazer EVM should implement.
	StatePlugin = state.Plugin

	// `ConfigurationPlugin` defines the methods that the chain running Stargazer EVM should
	// implement in order to configuration the parameters of the Stargazer EVM.
	ConfigurationPlugin interface {
		// `ConfigurationPlugin` implements `libtypes.Resettable`. Calling `Reset` should reset the
		// `ConfigurationPlugin` to a default state.
		libtypes.Resettable
		// `ChainConfig` returns the current chain configuration of the Stargazer EVM.
		ChainConfig() *params.ChainConfig
		// `ExtraEips` returns the extra EIPs that the Stargazer EVM supports.
		ExtraEips() []int
	}
)

// =============================================================================
// Optional Plugins
// =============================================================================

// `The following plugins are OPTIONAL to be implemented by the chain running Stargazer EVM.
type (
	// `PrecompilePlugin` defines the methods that the chain running Stargazer EVM should implement
	// in order to support running their own stateful precompiled contracts. Implementing this
	// plugin is optional.
	PrecompilePlugin = precompile.Plugin
)
