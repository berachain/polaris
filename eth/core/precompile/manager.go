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
	"math/big"

	"github.com/berachain/stargazer/eth/core/precompile/container"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/registry"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/lib/utils"
)

// `manager` retrieves and runs precompile containers with an ephemeral context.
type manager struct {
	// `Registry` allows the `Controller` to search for a precompile container at an address.
	libtypes.Registry[common.Address, vm.PrecompileContainer]

	// `ephemeralSDB` is the StargazerStateDB for a current state transition.
	ephemeralSDB vm.StargazerStateDB

	// `runner` will run the precompile in a custom precompile environment for a given context.
	runner vm.PrecompileRunner
}

// `NewManager` creates and returns a `Controller` with a new precompile registry and precompile
// runner.
func NewManager(runner vm.PrecompileRunner) vm.PrecompileManager {
	return &manager{
		Registry: registry.NewMap[common.Address, vm.PrecompileContainer](),
		runner:   runner,
	}
}

// `PrepareForStateTransition` sets the precompile's native environment statedb.
//
// `PrepareForStateTransition` implements `vm.PrecompileController`.
func (c *manager) PrepareForStateTransition(sdb vm.GethStateDB) error {
	ssdb, ok := utils.GetAs[vm.StargazerStateDB](sdb)
	if !ok {
		return container.ErrIncompatibleStateDB
	}

	c.ephemeralSDB = ssdb
	return nil
}

// `Run` runs the precompile container using its runner and its ephemeral context.
//
// `Run` implements `vm.PrecompileController`.
func (c *manager) Run(
	pc vm.PrecompileContainer, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	return c.runner.Run(pc, c.ephemeralSDB, input, caller, value, suppliedGas, readonly)
}
