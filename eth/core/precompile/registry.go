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
	"github.com/berachain/stargazer/eth/core/precompile/container"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
)

// CURRENTLY UNUSED -- MOVE THE BUILD CONTAINER LOGIC ELSEWHERE.

// `Registry` stores and provides all stateless, stateful, and dynamic precompile containers. It is
// a map of Ethereum addresses to precompiled contract containers.
type Registry map[common.Address]vm.PrecompileContainer

// `NewRegistry` creates and returns a new `Registry`.
func NewRegistry() Registry {
	return make(Registry)
}

// `Register` builds a precompile container using a container factory and stores the container
// at the given address. This function returns an error if the given contract is not a properly
// defined precompile or the container factory cannot build the container.
func (r Registry) Register(contractImpl vm.BasePrecompileImpl) error {
	// select the correct container factory based on the contract type.
	var cf container.AbstractFactory
	//nolint:gocritic // cannot be converted to switch-case.
	if utils.Implements[container.DynamicPrecompileImpl](contractImpl) {
		cf = container.NewDynamicFactory()
	} else if utils.Implements[container.StatefulPrecompileImpl](contractImpl) {
		cf = container.NewStatefulFactory()
	} else if utils.Implements[container.StatelessPrecompileImpl](contractImpl) {
		cf = container.NewStatelessFactory()
	} else {
		return ErrIncorrectPrecompileType
	}

	// build the container and store at its address.
	pc, err := cf.Build(contractImpl)
	if err != nil {
		return err
	}
	r[contractImpl.RegistryKey()] = pc

	return nil
}

// `lookup` returns a precompile container at the given address, if it exists.
func (r Registry) lookup(addr common.Address) (vm.PrecompileContainer, bool) {
	pc, found := r[addr]
	return pc, found
}
