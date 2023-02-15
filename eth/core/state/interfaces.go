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

	"github.com/berachain/stargazer/eth/common"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `Plugin` is a plugin which tracks the accounts (balances, nonces, codes, states) in the native
// vm. This also handles removing suicided accounts.
type Plugin interface {
	// `Plugin` implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// `Reset` resets the state with the given `context`.
	libtypes.Resettable

	// `CreateAccount` creates an account with the given `address`.
	CreateAccount(common.Address)
	// `Exist` reports whether the given account exists in state. Notably this should also return
	// true for suicided accounts.
	Exist(common.Address) bool

	// `GetBalance` returns the balance of the given account.
	GetBalance(common.Address) *big.Int
	// `AddBalance` adds amount to the given account.
	SubBalance(common.Address, *big.Int)
	// `SubBalance` subtracts amount from the given account.
	AddBalance(common.Address, *big.Int)
	// `TransferBalance` transfers amount from one account to the other.
	TransferBalance(common.Address, common.Address, *big.Int)

	// `GetNonce` returns the nonce of the given account.
	GetNonce(common.Address) uint64
	// `SetNonce` sets the nonce of the given account.
	SetNonce(common.Address, uint64)

	// `GetCodeHash` returns the code hash of the given account.
	GetCodeHash(common.Address) common.Hash
	// `GetCode` returns the code associated with a given account.
	GetCode(common.Address) []byte
	// `SetCode` sets the code associated with a given account.
	SetCode(common.Address, []byte)
	// `GetCodeSize` returns the size of the code associated with a given account.
	GetCodeSize(common.Address) int

	// `GetCommittedState` returns the committed value from account storage.
	GetCommittedState(common.Address, common.Hash) common.Hash
	// `GetState` returns the value from account storage.
	GetState(common.Address, common.Hash) common.Hash
	// `SetState` sets the value for a given key in account storage.
	SetState(common.Address, common.Hash, common.Hash)
	// `ForEachStorage` iterates over the storage of an account and calls the given callback
	// function.
	ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error

	// `DeleteSuicides` removes the given accounts from the state.
	DeleteSuicides([]common.Address)
}

// `LogsJournal` defines the interface for tracking logs created during a state transition.
type LogsJournal interface {
	// `LogsJournal` implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// `AddLog` adds a log to the state
	AddLog(*coretypes.Log)
	// `BuildLogsAndClear` returns the logs of the tx with the given metadata
	BuildLogsAndClear(common.Hash, common.Hash, uint, uint) []*coretypes.Log
}

// `RefundJournal` is a `Store` that tracks the refund counter.
type RefundJournal interface {
	// `RefundJournal` implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// `GetRefund` returns the current value of the refund counter.
	GetRefund() uint64
	// `AddRefund` sets the refund counter to the given `gas`.
	AddRefund(gas uint64)
	// `SubRefund` subtracts the given `gas` from the refund counter.
	SubRefund(gas uint64)
}
