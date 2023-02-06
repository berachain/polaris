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

package precompile

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/registry"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `manager` retrieves and runs precompile containers with an ephemeral context.
type manager struct {
	// `Registry` allows the `Controller` to search for a precompile container at an address.
	libtypes.Registry[common.Address, vm.PrecompileContainer]
	// `ctx` is the ephemeral native context, updated on every state transition.
	ctx context.Context
	// `runner` will run the precompile in a custom precompile environment for a given context.
	runner vm.PrecompileRunner
	// `ssdb` is a reference to the StateDB used to add logs from the precompile's execution.
	ssdb vm.StargazerStateDB
}

// `NewManager` creates and returns a `Controller` with a new precompile registry and precompile
// runner.
func NewManager(runner vm.PrecompileRunner, ssdb vm.StargazerStateDB) vm.PrecompileManager {
	return &manager{
		Registry: registry.NewMap[common.Address, vm.PrecompileContainer](),
		runner:   runner,
		ssdb:     ssdb,
	}
}

// `Reset` sets the precompile's native environment context.
//
// `Reset` implements `vm.PrecompileController`.
func (m *manager) Reset(ctx context.Context) {
	m.ctx = ctx
}

// `Run` runs the precompile container using its runner and its ephemeral context.
//
// `Run` implements `vm.PrecompileController`.
func (m *manager) Run(
	pc vm.PrecompileContainer, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	return m.runner.Run(m.ctx, m.ssdb, pc, input, caller, value, suppliedGas, readonly)
}
