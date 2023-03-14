// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package state

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// Plugin is a plugin which tracks the accounts (balances, nonces, codes, states) in the native
// vm. This also handles removing suicided accounts.
type Plugin interface {
	// Plugin implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// Reset resets the state with the given `context`.
	libtypes.Resettable
	// GetContext returns the current context of the state plugin.
	GetContext() context.Context

	// CreateAccount creates an account with the given `address`.
	CreateAccount(common.Address)
	// Exist reports whether the given account exists in state. Notably this should also return
	// true for suicided accounts.
	Exist(common.Address) bool

	// GetBalance returns the balance of the given account.
	GetBalance(common.Address) *big.Int
	// SetBalance sets the balance of the given account.
	SetBalance(common.Address, *big.Int)
	// AddBalance adds amount to the given account.
	SubBalance(common.Address, *big.Int)
	// SubBalance subtracts amount from the given account.
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
	// GetCodeSize returns the size of the code associated with a given account.
	GetCodeSize(common.Address) int
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
	// DeleteSuicides removes the given accounts from the state.
	DeleteSuicides([]common.Address)
}

type (
	// LogsJournal defines the interface for tracking logs created during a state transition.
	LogsJournal interface {
		// LogsJournal implements `libtypes.Controllable`.
		libtypes.Controllable[string]
		// SetTxContext sets the transaction hash and index for the current transaction.
		SetTxContext(thash common.Hash, ti int)
		// TxIndex returns the current transaction index.
		TxIndex() int
		// AddLog adds a log to the logs journal.
		AddLog(*coretypes.Log)
		// Logs returns the logs of the tx with the exisiting metadata.
		Logs() []*coretypes.Log
		// GetLogs returns the logs of the tx with the given metadata.
		GetLogs(hash common.Hash, blockNumber uint64, blockHash common.Hash) []*coretypes.Log
	}

	// RefundJournal is a `Store` that tracks the refund counter.
	RefundJournal interface {
		// RefundJournal implements `libtypes.Controllable`.
		libtypes.Controllable[string]
		// GetRefund returns the current value of the refund counter.
		GetRefund() uint64
		// AddRefund sets the refund counter to the given `gas`.
		AddRefund(gas uint64)
		// SubRefund subtracts the given `gas` from the refund counter.
		SubRefund(gas uint64)
	}

	AccessListJournal interface {
		// AccessListJournal implements `libtypes.Controllable`.
		libtypes.Controllable[string]
		// AddAddress adds the given `address` to the access list.
		AddAddress(address common.Address) bool
		// AddSlot adds the given `slot` to the access list.
		AddSlot(address common.Address, slot common.Hash) (addrChange bool, slotChange bool)
		// Contains returns whether the given `address` and `slot` are in the access list.
		Contains(address common.Address, slot common.Hash) (addressPresent bool, slotPresent bool)
		// ContainsAddress returns whether the given `address` is in the access list.
		ContainsAddress(address common.Address) bool
	}
)
