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
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
)

// `registry` stores and provides all stateless and stateful precompile containers. It is a map of
// Ethereum addresses to precompiled contract containers.
type registry map[common.Address]vm.PrecompileContainer

// `NewRegistry` creates and returns a new `vm.PrecompileRegistry`.
func NewRegistry() vm.PrecompileRegistry {
	return make(registry)
}

// `Register` builds a precompile container using a container factory and stores the container
// at the given address. This function returns an error if the given contract is not a properly
// defined precompile or the container factory cannot build the container.
func (r registry) Register(contractImpl vm.BaseContractImpl) error {
	// select the correct container factory based on the contract type.
	var cf AbstractContainerFactory
	//nolint:gocritic // cannot be converted to switch-case.
	if utils.Implements[StatefulContractImpl](contractImpl) {
		cf = nil
	} else if utils.Implements[StatelessContractImpl](contractImpl) {
		cf = nil
	} else {
		return ErrIncorrectPrecompileType
	}

	// build the container and store it at the given address.
	pc, err := cf.Build(contractImpl)
	if err != nil {
		return err
	}
	r[contractImpl.Address()] = pc

	return nil
}

// `Lookup` returns a precompile container at the given address, if it exists.
func (r registry) Lookup(addr common.Address) (vm.PrecompileContainer, bool) {
	pc, found := r[addr]
	return pc, found
}
