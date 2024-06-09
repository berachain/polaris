// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package state

import (
	"context"
	"math/big"

	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

// Plugin is a plugin which tracks the accounts (balances, nonces, codes, states) in the native
// vm. This also handles removing suicided accounts.
type Plugin interface {
	// Plugin implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// Reset resets the state with the given `context`.
	libtypes.Resettable
	// Plugin implements `libtypes.Cloneable`.
	libtypes.Cloneable[Plugin]
	// GetContext returns the current context of the state plugin.
	GetContext() context.Context
	// Error returns the current saved error of the state plugin.
	Error() error

	// CreateAccount creates an account with the given `address`.
	CreateAccount(common.Address)
	// Exist reports whether the given account exists in state. Notably this should also return
	// true for suicided accounts.
	Exist(common.Address) bool
	// Empty returns whether the given account is considered empty. Empty is defined according to
	// EIP161 (balance = nonce = code = 0).
	Empty(common.Address) bool
	// `DeleteAccounts` removes the given accounts from the state.
	DeleteAccounts([]common.Address)

	// GetBalance returns the balance of the given account.
	GetBalance(common.Address) *big.Int
	// SetBalance sets the balance of the given account.
	SetBalance(common.Address, *big.Int)
	// SubBalance subtracts amount from the given account.
	SubBalance(common.Address, *big.Int)
	// AddBalance adds amount to the given account.
	AddBalance(common.Address, *big.Int)

	// GetNonce returns the nonce of the given account.
	GetNonce(common.Address) uint64
	// SetNonce sets the nonce of the given account.
	SetNonce(common.Address, uint64)

	// GetCodeHash returns the code hash of the given account.
	GetCodeHash(common.Address) common.Hash
	// GetCode returns the code associated with a given account.
	GetCode(common.Address) []byte
	// SetCode sets the code associated with a given account.
	SetCode(common.Address, []byte)

	// GetCommittedState returns the committed value from account storage.
	GetCommittedState(common.Address, common.Hash) common.Hash
	// GetState returns the value from account storage.
	GetState(common.Address, common.Hash) common.Hash
	// SetState sets the value for a given key in account storage.
	SetState(common.Address, common.Hash, common.Hash)
	// SetStorage sets the storage of the given account.
	SetStorage(addr common.Address, storage map[common.Hash]common.Hash)
	// ForEachStorage iterates over the storage of an account and calls the given callback
	// function.
	ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error
}
