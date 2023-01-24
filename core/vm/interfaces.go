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
	"context"
	"math/big"

	"github.com/berachain/stargazer/lib/common"
)

type (
	// `StargazerStateDB` defines an extension to the interface provided by Go-Ethereum to support
	// additional state transition functionalities.
	StargazerStateDB interface {
		GethStateDB
		PrecompileStateDB

		// TransferBalance transfers the balance from one account to another
		TransferBalance(common.Address, common.Address, *big.Int)
	}

	// `PrecompileStateDB` defines the required functions to support execution of stateful
	// precompile contract containers.
	PrecompileStateDB interface {
		// `GetContext` returns the Go context associated to the StateDB.
		GetContext() context.Context
	}

	// `PrecompileRegistry` defines the required functions of a precompile registry.
	PrecompileRegistry interface {
		// `Register` registers a base precompile implementation at its address.
		Register(BasePrecompileImpl) error

		// `Lookup` searches the registry at an address and returns a precompile container if it
		// was found.
		Lookup(common.Address) (PrecompileContainer, bool)
	}

	// `BasePrecompileImpl` is a type for the base precompile implementation, which only needs to
	// provide an Ethereum address of where its contract is found.
	BasePrecompileImpl = ContractRef
)
