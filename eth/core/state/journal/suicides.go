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
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
)

// emptyCodeHash is the Keccak256 Hash of empty code
// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
var emptyCodeHash = crypto.Keccak256Hash(nil)

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
//
//nolint:revive // only used as a state.SuicidesJournal.
func NewSuicides(ssp suicideStatePlugin) *suicides {
	return &suicides{
		journal: stack.New[*common.Address](initCapacity),
		ssp:     ssp,
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
	*s = suicides{
		journal: stack.New[*common.Address](initCapacity),
		ssp:     s.ssp,
	}
}
