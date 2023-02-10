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

// `StargazerHostChain` defines the plugins that the chain running Stargazer EVM must implement.
type StargazerHostChain interface {
	GetBlockPlugin() BlockPlugin
	GetGasPlugin() GasPlugin
	GetStatePlugin() StatePlugin
	GetPrecompilePlugin() PrecompilePlugin
	GetConfigurationPlugin() ConfigurationPlugin

	// Temporary
	GetStargazerHeaderAtHeight(context.Context, uint64) *types.StargazerHeader
}

// =============================================================================
// Mandatory Plugins
// =============================================================================

// The following plugins MUST be implemented by the chain running Stargazer EVM and exposed via the
// `StargazerHostChain` interface. All plugins should be resettable with a given context.
type (
	// `BlockPlugin` defines the methods that the chain running Stargazer EVM must implement to\
	// support the `BlockPlugin` interface.
	BlockPlugin interface {
		libtypes.Resettable
		BaseFee() uint64
	}

	GasPlugin interface {
		// `GasPlugin` implements `libtypes.Resettable`. Calling Reset() MUST reset the GasPlugin to a
		// default state.
		libtypes.Resettable

		// `ConsumeGas` MUST consume the supplied amount of gas. It MUST not panic due to a GasOverflow
		// and must return core.ErrOutOfGas if the amount of gas remaining is less than the amount
		// requested.
		ConsumeGas(uint64) error

		// `RefundGas` MUST refund the supplied amount of gas. It MUST not panic.
		RefundGas(uint64)

		// `GasRemaining` MUST return the amount of gas remaining. It MUST not panic.
		GasRemaining() uint64

		// `GasUsed` MUST return the amount of gas used during the current transaction. It MUST not panic.
		GasUsed() uint64

		// `CumulativeGasUsed` MUST return the amount of gas used during the current block. The value returned
		// MUST include any gas consumed during this transaction. It MUST not panic.
		CumulativeGasUsed() uint64

		// `MaxFeePerGas` MUST set the maximum amount of gas that can be consumed by the meter. It MUST not panic, but
		// instead, return an error, if the new gas limit is less than the currently consumed amount of gas.
		SetGasLimit(uint64) error
	}

	// `StatePlugin` defines the methods that the chain running Stargazer EVM must implement.
	StatePlugin = state.Plugin

	// `ConfigurationPlugin` defines the methods that the chain running Stargazer EVM must implement
	// in order to configuration the parameters of the Stargazer EVM.
	ConfigurationPlugin interface {
		libtypes.Resettable
		ChainConfig() *params.ChainConfig
		UpdateChainConfig(*params.ChainConfig)
	}
)

// =============================================================================
// Optional Plugins
// =============================================================================

// `The following plugins are OPTIONAL to be implemented by the chain running Stargazer EVM.
type (
	// `PrecompilePlugin` defines the methods that the chain running Stargazer EVM must implement
	// in order to support running their own stateful precompiled contracts.
	//
	// Implementing this plugin is optional.
	PrecompilePlugin = precompile.Plugin
)
