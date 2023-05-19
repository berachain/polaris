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

package journal

import (
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
	libtypes "pkg.berachain.dev/polaris/lib/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// emptyCodeHash is the Keccak256 Hash of empty code
// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
var emptyCodeHash = crypto.Keccak256Hash(nil)

// `suicideStatePlugin` defines the required funtions from the StatePlugin for the suicide journal.
type suicideStatePlugin interface {
	// GetCodeHash returns the code hash of the given account.
	GetCodeHash(common.Address) common.Hash
	// GetBalance returns the balance of the given account.
	GetBalance(common.Address) *big.Int
	// SubBalance subtracts amount from the given account.
	SubBalance(common.Address, *big.Int)
}

type SuicidesI interface {
	// SuicidesI implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// SuicidesI implements `libtypes.Cloneable`.
	libtypes.Cloneable[SuicidesI]
	// SuicidesI marks the given address as suicided.
	Suicide(common.Address) bool
	// HasSuicided returns whether the address is suicided.
	HasSuicided(common.Address) bool
	// GetSuicides returns all suicided addresses from the tx.
	GetSuicides() []common.Address
}

// Dirty tracking of suicided accounts, we have to keep track of these manually, in order for the
// code and state to still be accessible even after the account has been deleted.
// NOTE: we are only supporting one suicided address per EVM call (and consequently per snapshot).
type suicides struct {
	// journal of suicide address per call, very rare to suicide so we alloc only 1 address
	journal ds.Stack[*common.Address]
	ssp     suicideStatePlugin
	// lastSnapshot ensures that only 1 address is being suicided per snapshot
	lastSnapshot int
}

// NewSuicides returns a new suicides journal.
func NewSuicides(ssp suicideStatePlugin) SuicidesI {
	return &suicides{
		journal:      stack.New[*common.Address](initCapacity),
		ssp:          ssp,
		lastSnapshot: -1,
	}
}

// RegistryKey implements libtypes.Registrable.
func (s *suicides) RegistryKey() string {
	return suicidesRegistryKey
}

// Suicide implements the PolarisStateDB interface by marking the given address as suicided.
// This clears the account balance, but the code and state of the address remains available
// until after Commit is called.
func (s *suicides) Suicide(addr common.Address) bool {
	// ensure only one suicide per snapshot call
	if s.journal.Size() > s.lastSnapshot {
		// pushed one suicide for this contract call, can do no more
		return false
	}

	// only smart contracts can commit suicide
	ch := s.ssp.GetCodeHash(addr)
	if (ch == common.Hash{}) || ch == emptyCodeHash {
		return false
	}

	// Reduce it's balance to 0.
	s.ssp.SubBalance(addr, s.ssp.GetBalance(addr))

	// add to journal.
	s.journal.Push(&addr)
	return true
}

// HasSuicided implements the PolarisStateDB interface by returning if the contract was suicided
// in current transaction.
func (s *suicides) HasSuicided(addr common.Address) bool {
	for i := s.journal.Size() - 1; i >= 0; i-- {
		if *s.journal.PeekAt(i) == addr {
			return true
		}
	}
	return false
}

// GetSuicides implements state.SuicidesJournal.
func (s *suicides) GetSuicides() []common.Address {
	var suicidalAddrs []common.Address
	for i := 0; i < s.journal.Size(); i++ {
		suicidalAddrs = append(suicidalAddrs, *s.journal.PeekAt(i))
	}
	return suicidalAddrs
}

// Snapshot implements libtypes.Controllable.
func (s *suicides) Snapshot() int {
	s.lastSnapshot = s.journal.Size()
	return s.lastSnapshot
}

// RevertToSnapshot implements libtypes.Controllable.
func (s *suicides) RevertToSnapshot(id int) {
	s.journal.PopToSize(id)
}

// Finalize implements libtypes.Controllable.
func (s *suicides) Finalize() {
	*s = *utils.MustGetAs[*suicides](NewSuicides(s.ssp))
}

// Clone implements libtypes.Cloneable.
func (s *suicides) Clone() SuicidesI {
	size := s.journal.Size()
	clone := &suicides{
		journal:      stack.New[*common.Address](size),
		ssp:          s.ssp,
		lastSnapshot: s.lastSnapshot,
	}

	// copy every address from the journal
	for i := 0; i < size; i++ {
		cpy := new(common.Address)
		*cpy = *s.journal.PeekAt(i)
		clone.journal.Push(cpy)
	}

	return clone
}
