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
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/state"
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
}

// The following plugins must be implemented by the chain running Stargazer EVM and exposed via the
// `StargazerHostChain` interface. All plugins should be resettable with a given context.
type (
	BlockPlugin interface {
		libtypes.Resettable
	}

	GasPlugin interface {
		libtypes.Resettable
		ConsumeGas(amount uint64) error
		RefundGas(amount uint64)
		GasRemaining() uint64
		GasUsed() uint64
		CumulativeGasUsed() uint64
		SetGasLimit(limit uint64) error
	}

	StatePlugin = state.Plugin

	PrecompilePlugin = precompile.Plugin

	ConfigurationPlugin interface {
		libtypes.Resettable
		BaseFee() uint64
		ChainConfig() *params.ChainConfig
		UpdateChainConfig(*params.ChainConfig)
	}
)
