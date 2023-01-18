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

package state

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/berachain/stargazer/lib/common"
)

type GethStateDB = vm.StateDB

// `StargazerStateDB` defines an extension to the interface provided by go-ethereum to
// support additional state transition functionalities that are useful in a Cosmos SDK context.
type StargazerStateDB interface {
	GethStateDB

	// TransferBalance transfers the balance from one account to another
	TransferBalance(common.Address, common.Address, *big.Int)
}

// `PrecompileStateDB` defines an extension to the interface provided by the Go-Ethereum codebase
// to support additional state transition functionalities. In particular it supports getting the
// cosmos sdk context for natively running stateful precompiled contracts.
type PrecompileStateDB interface {
	GethStateDB

	// `GetContext` returns the cosmos sdk context with the statedb multistore attached.
	GetContext() sdk.Context

	// `EnableEventLogging` enables Cosmos events to be added to Ethereum logs.
	EnableEventLogging()

	// `DisableEventLogging` disables Cosmos events to be added to Ethereum logs.
	DisableEventLogging()
}
