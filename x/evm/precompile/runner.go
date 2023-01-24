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
)

var _ vm.PrecompileRunner = (*CosmosRunner)(nil)

type CosmosRunner struct {
	registry vm.PrecompileRegistry

	psdb vm.PrecompileStateDB
}

// `NewRunner` creates and returns a new `Runner` for the given precompile
// registry `registry` and precompile StateDB `psdb`.
func NewRunner(registry vm.PrecompileRegistry, psdb vm.PrecompileStateDB) *CosmosRunner {
	return &CosmosRunner{
		registry: registry,
		psdb:     psdb,
	}
}

// `Exists` gets a precompile container at the given `addr` from the precompile registry.
//
// `Exists` implements `vm.PrecompileRunner`.
func (cr *CosmosRunner) Exists(addr common.Address) (vm.PrecompileContainer, bool) {
	return cr.registry.Lookup(addr)
}

// `Run` runs the given precompile container and returns the remaining gas after execution. This
// function returns an error if the given statedb is not compatible with precompiles, insufficient
// gas is provided, or the precompile execution returns an error.
//
// `Run` implements `vm.PrecompileRunner`.
func (cr *CosmosRunner) Run(
	pc vm.PrecompileContainer,
	input []byte,
	caller common.Address,
	value *big.Int,
	suppliedGas uint64,
	readonly bool,
) ([]byte, uint64, error) {
	return nil, 0, nil
}
