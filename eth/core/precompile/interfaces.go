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
	"math/big"

	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	libtypes "github.com/berachain/stargazer/lib/types"
)

type (
	// `LogsDB` defines the required function to add a log to the StateDB. This ensures a precompile
	// runner can only add logs to the StateDB and not modify any other state on the StateDB.
	LogsDB interface {
		// `AddLog` adds a log to the StateDB.
		AddLog(*coretypes.Log)
	}

	// `Plugin` defines the required functions that a precompile plugin must implement.
	Plugin interface {
		libtypes.Resettable
		// `Run` runs a precompile container with its natives context and returns the remaining gas.
		Run(ldb LogsDB, pc vm.PrecompileContainer, input []byte,
			caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
		) (ret []byte, remainingGas uint64, err error)
	}
)
