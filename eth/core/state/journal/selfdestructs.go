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
