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
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
)

type accessList struct {
	*AccessList                       // current access list, always the head of journal stack.
	journal     ds.Stack[*AccessList] // journal of access lists.
}

// NewAccesslist returns a new `accessList` journal.
//
//nolint:revive // only used as a `state.AccessListJournal`.
func NewAccesslist() *accessList {
	journal := stack.New[*AccessList](initCapacity)
	journal.Push(NewAccessList())
	return &accessList{
		AccessList: journal.Peek(),
		journal:    journal,
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (al *accessList) RegistryKey() string {
	return accessListRegistryKey
}

// AddAddressToAccessList implements `state.AccessListJournal`.
func (al *accessList) AddAddressToAccessList(addr common.Address) {
	al.AddAddress(addr)
}

// AddSlotToAccessList implements `state.AccessListJournal`.
func (al *accessList) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	al.AddSlot(addr, slot)
}

// AddressInAccessList implements `state.AccessListJournal`.
func (al *accessList) AddressInAccessList(addr common.Address) bool {
	return al.ContainsAddress(addr)
}

// SlotInAccessList implements `state.AccessListJournal`.
func (al *accessList) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	return al.Contains(addr, slot)
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (al *accessList) Snapshot() int {
	al.AccessList = al.AccessList.Copy()
	al.journal.Push(al.AccessList)
	return al.journal.Size() - 1
}

// RevertToSnapshot implements `libtypes.Snapshottable`.
func (al *accessList) RevertToSnapshot(id int) {
	al.journal.PopToSize(id)
	al.AccessList = al.journal.Peek()
}

// Finalize implements `libtypes.Controllable`.
func (al *accessList) Finalize() {
	al.journal = stack.New[*AccessList](initCapacity)
	al.journal.Push(NewAccessList())
	al.AccessList = al.journal.Peek()
}
