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

package container

import (
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
)

type (
	// `AbstractFactory` is an interface that all precompile container factories must adhere to.
	AbstractFactory interface {
		// `Build` builds and returns the precompile container for the type of container/factory.
		Build(vm.RegistrablePrecompile) (vm.PrecompileContainer, error)
	}
)

type (
	// `StatelessPrecompileImpl` is the interface for all stateless precompiled contract
	// implementations. A stateless contract must provide its own precompile container, as it is
	// stateless in nature. This requires a deterministic gas count, `RequiredGas`, and an
	// executable function `Run`.
	StatelessPrecompileImpl interface {
		vm.RegistrablePrecompile

		vm.PrecompileContainer
	}

	// `StatefulPrecompileImpl` is the interface for all stateful precompiled contracts, which must
	// expose their ABI methods and precompile methods for stateful execution.
	StatefulPrecompileImpl interface {
		vm.RegistrablePrecompile

		// `ABIMethods` should return a map of Ethereum method names to Go-Ethereum abi `Method`
		// structs. NOTE: this can be directly loaded from the `Methods` field of a Go-Ethereum ABI
		// struct, which can be built for a solidity interface or contract.
		ABIMethods() map[string]abi.Method

		// `PrecompileMethods` should return all the stateful precompile's functions (and each of
		// their required gas).
		PrecompileMethods() Methods
	}

	// `DynamicPrecompileImpl` is the interface for all dynamic stateful precompiled contracts.
	DynamicPrecompileImpl interface {
		StatefulPrecompileImpl

		// `Name` should return a string name of the dynamic contract.
		Name() string
	}
)
