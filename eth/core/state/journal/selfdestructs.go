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

	libtypes "github.com/berachain/polaris/lib/types"
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// emptyCodeHash is the Keccak256 Hash of empty code
// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
var emptyCodeHash = crypto.Keccak256Hash(nil)

// `selfDestructStatePlugin` defines the required functions from the StatePlugin
// for the suicide journal.
type selfDestructStatePlugin interface {
	// GetCodeHash returns the code hash of the given account.
	GetCodeHash(common.Address) common.Hash
	// GetBalance returns the balance of the given account.
	GetBalance(common.Address) *big.Int
	// SubBalance subtracts amount from the given account.
	SubBalance(common.Address, *big.Int)
}

type SelfDestructs interface {
	// SelfDestructs implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// SelfDestructs implements `libtypes.Cloneable`.
	libtypes.Cloneable[SelfDestructs]
	// SelfDestruct marks the given address as self destructed .
	SelfDestruct(common.Address)
	// Selfdestruct6780 marks the given address as self destructed post eip-6780
	Selfdestruct6780(common.Address)
	// HasSelfDestructed returns whether the address is self destructed .
	HasSelfDestructed(common.Address) bool
	// GetSelfDestructs returns all self destructed addresses from the tx.
	GetSelfDestructs() []common.Address
}

// Dirty tracking of self destructed accounts, we have to keep track of these manually,
// in order for the code and state to still be accessible even after the account has
// been deleted. NOTE: we are only supporting one self destructed address per EVM call
// (and consequently per snapshot).
type selfDestructs struct {
	// journal of suicide address per call, very rare to suicide so we alloc only 1 address
	baseJournal[*common.Address]
	ssp selfDestructStatePlugin
	// lastSnapshot ensures that only 1 address is being self destructed per snapshot
	lastSnapshot int
}

// NewSelfDestructs returns a new selfDestructs journal.
func NewSelfDestructs(ssp selfDestructStatePlugin) SelfDestructs {
	return &selfDestructs{
		baseJournal:  newBaseJournal[*common.Address](initCapacity),
		ssp:          ssp,
		lastSnapshot: -1,
	}
}

// RegistryKey implements libtypes.Registrable.
func (s *selfDestructs) RegistryKey() string {
	return suicidesRegistryKey
}

// SelfDestruct implements the PolarStateDB interface by marking the given address as self
// destructed. This clears the account balance, but the code and state of the address remains
// available until after Commit is called.
func (s *selfDestructs) SelfDestruct(addr common.Address) {
	// ensure only one suicide per snapshot call
	if s.Size() > s.lastSnapshot {
		// pushed one suicide for this contract call, can do no more
		return
	}

	// only smart contracts can commit suicide
	ch := s.ssp.GetCodeHash(addr)
	if (ch == common.Hash{}) || ch == emptyCodeHash {
		return
	}

	// Reduce it's balance to 0.
	s.ssp.SubBalance(addr, s.ssp.GetBalance(addr))

	// add to journal.
	s.Push(&addr)
}

func (s *selfDestructs) Selfdestruct6780(_ common.Address) {
	// TODO: IMPLEMENT EIP-6780
}

// HasSelfDestructed implements the PolarStateDB interface by returning if the contract was
// self destructed in current transaction.
func (s *selfDestructs) HasSelfDestructed(addr common.Address) bool {
	for i := s.Size() - 1; i >= 0; i-- {
		if *s.PeekAt(i) == addr {
			return true
		}
	}
	return false
}

// GetSelfDestructs implements state.SelfDestructsJournal.
func (s *selfDestructs) GetSelfDestructs() []common.Address {
	var suicidalAddrs []common.Address
	for i := 0; i < s.Size(); i++ {
		suicidalAddrs = append(suicidalAddrs, *s.PeekAt(i))
	}
	return suicidalAddrs
}

func (s *selfDestructs) Snapshot() int {
	s.lastSnapshot = s.Size()
	return s.baseJournal.Snapshot()
}

// Finalize implements libtypes.Controllable.
func (s *selfDestructs) Finalize() {
	*s = *utils.MustGetAs[*selfDestructs](NewSelfDestructs(s.ssp))
}

// Clone implements libtypes.Cloneable.
func (s *selfDestructs) Clone() SelfDestructs {
	clone := &selfDestructs{
		baseJournal:  newBaseJournal[*common.Address](s.Capacity()),
		ssp:          s.ssp,
		lastSnapshot: s.lastSnapshot,
	}

	// copy every address from the journal
	for i := 0; i < s.Size(); i++ {
		cpy := new(common.Address)
		*cpy = *s.PeekAt(i)
		clone.Push(cpy)
	}

	return clone
}
