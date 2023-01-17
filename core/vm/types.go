// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/lib/common"
	gevm "github.com/ethereum/go-ethereum/core/vm"
)

type (
	// `PrecompiledContract` is the basic interface for native Go contracts. The implementation
	// requires a deterministic gas count based on the input size of the `Run` method of the
	// contract.
	PrecompiledContract = gevm.PrecompiledContract

	// `StatefulPrecompiledContract` is the interface for all stateful precompiled contracts, which
	// must expose their functions and gas requirements for stateful execution.
	StatefulPrecompiledContract interface {
		PrecompiledContract

		// `GetFunctionsAndGas` should return all the stateful precompile's functions (and each of
		// their required gas).
		GetFunctionsAndGas() precompile.FnsAndGas
	}

	// `DynamicPrecompiledContract` is the interface for all dynamic stateful precompiled
	// contracts.
	DynamicPrecompiledContract interface {
		StatefulPrecompiledContract

		// `Name` should return a string name of the dynamic contract.
		Name() string
	}

	// `PrecompileGetter` is a type of function used by the EVM to retrieve precompiled contracts
	// during EVM execution.
	PrecompileGetter func(addr common.Address) (PrecompiledContract, bool)
)
