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
// `StargazerHostChain` interface.
type (
	BlockPlugin interface {
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
