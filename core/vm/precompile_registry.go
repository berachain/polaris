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

package vm

import (
	"reflect"

	"cosmossdk.io/errors"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
)

var statefulContractType = reflect.TypeOf(precompile.StatefulContractImpl(nil))
var statelessContractType = reflect.TypeOf(precompile.StatelessContractImpl(nil))

// `PrecompileRegistry` stores and provides stateless and stateful precompile containers to a
// precompile host.
type PrecompileRegistry struct {
	// `precompiles` is a map of Ethereum addresses to precompiled contract containers. Only
	// supporting stateless and stateful precompiles for now.
	precompiles map[common.Address]types.PrecompileContainer

	// `logFactory` is the Ethereum log builder for all Cosmos events emitted during precompile
	// execution.
	logFactory *container.LogFactory
}

// `NewPrecompileRegistry` creates and returns a new `PrecompileRegistry`.
func NewPrecompileRegistry() *PrecompileRegistry {
	return &PrecompileRegistry{
		precompiles: make(map[common.Address]types.PrecompileContainer),
		logFactory:  container.NewLogFactory(),
	}
}

// `Register` builds a precompile container using a container factory and stores the container
// at the given address. This function returns an error if the given contract is not a properly
// defined precompile or the container factory cannot build the container.
func (pr *PrecompileRegistry) Register(
	addr common.Address,
	contract precompile.BaseContractImpl,
) error {
	var cf precompile.AbstractContainerFactory
	contractType := reflect.ValueOf(contract).Type()
	//nolint:gocritic // cannot be converted to switch-case.
	if contractType.Implements(statefulContractType) {
		cf = precompile.NewStatefulContainerFactory()
	} else if contractType.Implements(statelessContractType) {
		cf = precompile.NewStatelessContainerFactory()
	} else {
		return errors.Wrap(ErrIncorrectPrecompileType, contractType.Name())
	}

	pc, err := cf.Build(contract)
	if err != nil {
		return err
	}
	pr.precompiles[addr] = pc

	return nil
}

// `Get` returns a precompile container at the given address, if it exists.
func (pr *PrecompileRegistry) Get(addr common.Address) (types.PrecompileContainer, bool) {
	pc, found := pr.precompiles[addr]
	return pc, found
}
