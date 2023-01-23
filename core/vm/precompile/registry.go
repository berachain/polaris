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
	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
)

// `Registry` stores and provides all stateless and stateful precompile containers to a
// precompile host.
type Registry struct {
	// `precompiles` is a map of Ethereum addresses to precompiled contract containers. Only
	// supporting stateless and stateful precompiles for now.
	precompiles map[common.Address]types.PrecompileContainer

	// `Registry` is the registry for all Cosmos events emitted during precompile
	// execution.
	Registry *log.Registry
}

// `NewRegistry` creates and returns a new `Registry`.
func NewRegistry(logTranslator log.Translator) *Registry {
	return &Registry{
		precompiles: make(map[common.Address]types.PrecompileContainer),
		Registry:    log.NewRegistry(logTranslator),
	}
}

// `Register` builds a precompile container using a container factory and stores the container
// at the given address. This function returns an error if the given contract is not a properly
// defined precompile or the container factory cannot build the container.
func (pr *Registry) Register(contractImpl container.BaseContractImpl) error {
	var cf container.AbstractContainerFactory
	//nolint:gocritic // cannot be converted to switch-case.
	if utils.Implements[container.StatefulContractImpl](contractImpl) {
		cf = container.NewStatefulContainerFactory(pr.Registry)
	} else if utils.Implements[container.StatelessContractImpl](contractImpl) {
		cf = container.NewStatelessContainerFactory()
	} else {
		return ErrIncorrectPrecompileType
	}

	pc, err := cf.Build(contractImpl)
	if err != nil {
		return err
	}
	pr.precompiles[contractImpl.Address()] = pc

	return nil
}

// `Get` returns a precompile container at the given address, if it exists.
func (pr *Registry) Get(addr common.Address) (types.PrecompileContainer, bool) {
	pc, found := pr.precompiles[addr]
	return pc, found
}
