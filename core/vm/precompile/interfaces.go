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
	"context"
	"math/big"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
)

// `Runner` is invoked by the EVM to determine if a particular address is one of a precompiled
// contract and allows the EVM to execute a precompiled contract function.
type RunnerI interface {
	// `Exists` returns if a precompiled contract was found at `addr`.
	Exists(addr common.Address) (types.PrecompileContainer, bool)

	// `Run` runs a precompiled contract container and returns the remaining gas.
	Run(pc types.PrecompileContainer, input []byte, caller common.Address,
		value *big.Int, suppliedGas uint64, readonly bool,
	) (ret []byte, leftOverGas uint64, err error)
}

// `PrecompileStateDB` defines the required functions to support execution of stateful precompiled
// contracts.
type PrecompileStateDB interface { //nolint: revive // todo: come up with better name.
	// `AddLog` adds a log to the StateDB.
	AddLog(*coretypes.Log)

	// `GetContext` returns the Cosmos SDK context with the StateDB Multistore attached.
	GetContext() context.Context
}
