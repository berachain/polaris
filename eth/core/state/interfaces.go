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

	"github.com/berachain/stargazer/lib/common"
)

type Store interface {
	// `GetName` returns the name of the store
	Name() string

	// `RevertToSnapshot` reverts the state to a previous version
	RevertToSnapshot(int)

	// `Snapshot` returns an identifier for the current revision of the state.
	Snapshot() int
}

type AccountPlugin interface {
	// `CreateAccount` creates a new account with the given address
	CreateAccount(context.Context, common.Address)

	// `GetNonce` returns the nonce of the account associated with the given address
	GetNonce(context.Context, common.Address) uint64

	// `SetNonce` sets the nonce of the account associated with the given address
	SetNonce(context.Context, common.Address, uint64)
}

type BalancePlugin interface {
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

type StoragePlugin interface {
	// `GetState` returns the value of key in the storage of the account associated with the given address
	GetState(context.Context, common.Address, common.Hash) common.Hash

	// `SetState` sets the value of key in the storage of the account associated with the given address
	SetState(context.Context, common.Address, common.Hash, common.Hash)

	// `DeleteState` deletes the value of key in the storage of the account associated with the given address
	DeleteState(context.Context, common.Address, common.Hash)
}
