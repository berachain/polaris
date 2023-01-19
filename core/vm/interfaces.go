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
	"math/big"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/lib/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `StargazerStateDB` defines an extension to the interface provided by go-ethereum to
// support additional state transition functionalities that are useful in a Cosmos SDK context.
type StargazerStateDB interface {
	GethStateDB

	// TransferBalance transfers the balance from one account to another
	TransferBalance(common.Address, common.Address, *big.Int)

	// `GetContext` returns the cosmos sdk context with the statedb multistore attached.
	GetContext() sdk.Context
}

// `PrecompileStateDB` defines an extension to the interface provided by go-ethereum to
// support additional state transition functionalities that are useful in a Cosmos SDK context.
type PrecompileStateDB interface {
	// `AddLog` adds a log to the statedb.
	AddLog(log *coretypes.Log)

	// `GetContext` returns the cosmos sdk context with the statedb multistore attached.
	GetContext() sdk.Context
}
