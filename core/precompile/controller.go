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

	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
)

// Compile-time assertion to ensure `Controller` adheres to `PrecompileController`.
var _ vm.PrecompileController = (*Controller)(nil)

// `Controller` retrieves and runs precompile containers with an ephemeral context.
type Controller struct {
	// `ephemeralSDB` is the StargazerStateDB for a current state transition.
	ephemeralSDB vm.StargazerStateDB

	// `runner` will run the precompile in a custom precompile environment for a given context.
	runner vm.PrecompileRunner

	// `registry` allows the `Controller` to search for a precompile container at an address.
	registry registry
}

// `NewPrecompileController` creates and returns a `Controller` with the given precompile
// registry and precompile runner.
func NewPrecompileController(registry registry, runner vm.PrecompileRunner) *Controller {
	return &Controller{
		runner:   runner,
		registry: registry,
	}
}

// `PrepareForStateTransition` sets the precompile's native environment statedb.
//
// `PrepareForStateTransition` implements `vm.PrecompileController`.
func (c *Controller) PrepareForStateTransition(sdb vm.GethStateDB) error {
	ssdb, ok := utils.GetAs[vm.StargazerStateDB](sdb)
	if !ok {
		return ErrIncompatibleStateDB
	}

	c.ephemeralSDB = ssdb
	return nil
}

// `Exists` searches the registry at the given `addr` for a precompile container.
//
// `Exists` implements `vm.PrecompileController`.
func (c *Controller) Exists(addr common.Address) (vm.PrecompileContainer, bool) {
	return c.registry.lookup(addr)
}

// `Run` runs the precompile container using its runner and its ephemeral context.
//
// `Run` implements `vm.PrecompileController`.
func (c *Controller) Run(
	pc vm.PrecompileContainer, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	return c.runner.Run(pc, c.ephemeralSDB, input, caller, value, suppliedGas, readonly)
}
