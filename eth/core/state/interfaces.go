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
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core/state/plugin"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	libtypes "github.com/berachain/stargazer/lib/types"
)

type StateDBBackend interface { //nolint:revive // vibes.
	libtypes.Snapshottable

	GetContext() context.Context
	CreateAccount(common.Address)

	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	GetBalance(common.Address) *big.Int
	TransferBalance(common.Address, common.Address, *big.Int)

	GetNonce(common.Address) uint64
	SetNonce(common.Address, uint64)

	GetCodeHash(common.Address) common.Hash
	GetCode(common.Address) []byte
	SetCode(common.Address, []byte)
	GetCodeSize(common.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	GetCommittedState(common.Address, common.Hash) common.Hash
	GetState(common.Address, common.Hash) common.Hash
	SetState(common.Address, common.Hash, common.Hash)

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(common.Address) bool
	ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error

	Finalize() error
}

type LogsPlugin interface {
	plugin.Base
	// `Prepare` prepares the state for a txHash
	Prepare(common.Hash, uint)

	// `AddLog` adds a log to the state
	AddLog(*coretypes.Log)

	// `GetLogs` returns the logs of the state
	GetLogs(common.Hash, common.Hash) []*coretypes.Log
}

type RefundPlugin interface {
	plugin.Base
	// `AddRefund` adds amount to the refund counter
	Add(uint64)

	// `SubRefund` subtracts amount from the refund counter
	Sub(uint64)

	// `GetRefund` returns the refund counter
	Get() uint64
}
