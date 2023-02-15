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
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/params"
)

// `blockchain` is the canonical, persistent object that operates the Stargazer EVM.
type blockchain struct {
	// `host` is the host chain that the Stargazer EVM is running on.
	host StargazerHostChain
	// `csp` is the canonical, persistent state processor that runs the EVM.
	csp *StateProcessor
	// config is the state factory that builds state processors and statedbs.
	config *params.ChainConfig
}

// `NewChain` creates and returns a `blockchain` with the given EVM chain configuration and
// host chain. TODO: return public, exported `api.Chain` interface instead of `*blockchain`.
//
//nolint:revive // will be changed.
func NewChain(config *params.ChainConfig, host StargazerHostChain) *blockchain {
	bc := &blockchain{
		host:   host,
		config: config,
	}

	bc.csp = bc.BuildStateProcessor(vm.Config{}, true)
	return bc
}

// `BuildStateProcessor` builds and returns a `StateProcessor` with the given EVM configuration and
// commit flag.
func (bc *blockchain) BuildStateProcessor(vmConfig vm.Config, commit bool) *StateProcessor {
	return NewStateProcessor(bc.host, state.NewStateDB(bc.host.GetStatePlugin()), vmConfig, commit)
}
