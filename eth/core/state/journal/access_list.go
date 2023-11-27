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
	libtypes "github.com/berachain/polaris/lib/types"
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/common"
	ethstate "github.com/ethereum/go-ethereum/core/state"
)

type Accesslist interface {
	// Accesslist implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// Accesslist implements `libtypes.Cloneable`.
	libtypes.Cloneable[Accesslist]
	// `AddAddressToAccessList` adds the given address to the access list.
	AddAddressToAccessList(common.Address)
	// `AddSlotToAccessList` adds the given slot to the access list for the given address.
	AddSlotToAccessList(common.Address, common.Hash)
	// `SlotInAccessList` returns whether the given address and slot are in the access list.
	SlotInAccessList(common.Address, common.Hash) (addressPresent bool, slotPresent bool)
	// `AddressInAccessList` returns whether the given address is in the access list.
	AddressInAccessList(common.Address) bool
}

// accessList is a `baseJournal` that tracks the access list.
type accessList struct {
	baseJournal[*ethstate.AccessList] // journal of access lists.
}

// NewAccesslist returns a new `accessList` journal.
func NewAccesslist() Accesslist {
	journal := newBaseJournal[*ethstate.AccessList](initCapacity)
	journal.Push(ethstate.NewAccessList())
	return &accessList{
		baseJournal: journal,
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (al *accessList) RegistryKey() string {
	return accessListRegistryKey
}

// AddAddressToAccessList implements `state.AccessListJournal`.
func (al *accessList) AddAddressToAccessList(addr common.Address) {
	al.Peek().AddAddress(addr)
}

// AddSlotToAccessList implements `state.AccessListJournal`.
func (al *accessList) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	al.Peek().AddSlot(addr, slot)
}

// AddressInAccessList implements `state.AccessListJournal`.
func (al *accessList) AddressInAccessList(addr common.Address) bool {
	return al.Peek().ContainsAddress(addr)
}

// SlotInAccessList implements `state.AccessListJournal`.
func (al *accessList) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	return al.Peek().Contains(addr, slot)
}

// Snapshot implements `libtypes.Snapshottable`.
func (al *accessList) Snapshot() int {
	al.Push(al.Peek().Copy())
	// Accesslist is size minus one, since we want to revert to the place in the stack
	// where snapshot was called, which since we need to push a copy on the stack, is
	// the size minus one.
	return al.baseJournal.Size() - 1
}

// Finalize implements `libtypes.Controllable`.
func (al *accessList) Finalize() {
	*al = *utils.MustGetAs[*accessList](NewAccesslist())
}

// Clone implements `libtypes.Cloneable`.
func (al *accessList) Clone() Accesslist {
	cpy := &accessList{
		baseJournal: newBaseJournal[*ethstate.AccessList](al.Capacity()),
	}

	for i := 0; i < al.Size(); i++ { // skip the root, already pushed above
		cpy.Push(al.PeekAt(i).Copy())
	}

	return cpy
}
