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
)

type AccountPlugin interface {
	plugin.Base
	// `CreateAccount` creates a new account with the given address
	CreateAccount(context.Context, common.Address)

	// `HasAccount` returns true if the account associated with the given address exists
	HasAccount(context.Context, common.Address) bool

	// `DeleteAccount` deletes the account associated with the given address
	DeleteAccount(context.Context, common.Address)

	// `GetNonce` returns the nonce of the account associated with the given address
	GetNonce(context.Context, common.Address) uint64

	// `SetNonce` sets the nonce of the account associated with the given address
	SetNonce(context.Context, common.Address, uint64)
}

type BalancePlugin interface {
	plugin.Base
	// `GetBalance` returns the balance of the account associated with the given address
	GetBalance(context.Context, common.Address) *big.Int

	// `AddBalance` adds amount to the balance of the account associated with the given address
	AddBalance(context.Context, common.Address, *big.Int)

	// `SubBalance` subtracts amount from the balance of the account associated with the given address
	SubBalance(context.Context, common.Address, *big.Int)

	// `TransferBalance` transfers amount from the balance of the account associated with the
	// given from address to the balance of the account associated with the given to address
	TransferBalance(context.Context, common.Address, common.Address, *big.Int)
}

type CodePlugin interface {
	plugin.Base
	// `GetCodeHash` returns the code hash of the account associated with the given address
	GetCodeHash(context.Context, common.Address) common.Hash

	// `SetCodeHash` sets the code hash of the account associated with the given address
	SetCodeHash(context.Context, common.Address, common.Hash)

	// `GetCodeFromHash` returns the code associated with the given hash
	GetCodeFromHash(context.Context, common.Hash) []byte

	// `SetCode` sets the code of the account associated with the given hash
	SetCode(context.Context, common.Hash, []byte)

	// `DeleteCode` deletes the code of the account associated with the given address
	DeleteCode(context.Context, common.Address)
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

type StoragePlugin interface {
	plugin.Base
	// `GetState` returns the value of key in the storage of the account associated with the given address
	GetState(context.Context, common.Address, common.Hash) common.Hash

	// `GetCommittedState` returns the value of key in the storage of the account associated with the given address
	GetCommittedState(context.Context, common.Address, common.Hash) common.Hash

	// `ForEachStorage` iterates over the storage of the account associated with the given address
	ForEachStorage(context.Context, common.Address, func(common.Hash, common.Hash) bool) error

	// `SetState` sets the value of key in the storage of the account associated with the given address
	SetState(context.Context, common.Address, common.Hash, common.Hash)

	// `DeleteState` deletes the value of key in the storage of the account associated with the given address
	DeleteState(context.Context, common.Address, common.Hash)
}
