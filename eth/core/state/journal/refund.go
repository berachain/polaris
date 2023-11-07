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
)

// Refund is a `Store` that tracks the refund counter.
type Refund interface {
	// Refund implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// Refund implements `libtypes.Cloneable`.
	libtypes.Cloneable[Refund]
	// GetRefund returns the current value of the refund counter.
	GetRefund() uint64
	// AddRefund sets the refund counter to the given `gas`.
	AddRefund(gas uint64)
	// SubRefund subtracts the given `gas` from the refund counter.
	SubRefund(gas uint64)
}

// refund is a `Store` that tracks the refund counter.
type refund struct {
	baseJournal[uint64] // journal of historical refunds.

}

// NewRefund creates and returns a `refund` journal.
func NewRefund() Refund {
	return &refund{
		baseJournal: newBaseJournal[uint64](initCapacity),
	}
}

// RegistryKey implements `libtypes.Registrable`.
func (r *refund) RegistryKey() string {
	return refundRegistryKey
}

// GetRefund returns the current value of the refund counter.
func (r *refund) GetRefund() uint64 {
	// When the refund counter is empty, the stack will return 0 by design.
	return r.Peek()
}

// AddRefund sets the refund counter to the given `gas`.
func (r *refund) AddRefund(gas uint64) {
	r.Push(r.Peek() + gas)
}

// SubRefund subtracts the given `gas` from the refund counter.
func (r *refund) SubRefund(gas uint64) {
	r.Push(r.Peek() - gas)
}

// Finalize implements `libtypes.Controllable`.
func (r *refund) Finalize() {
	r.baseJournal = newBaseJournal[uint64](initCapacity)
}

// Clone implements `libtypes.Cloneable`.
func (r *refund) Clone() Refund {
	clone := &refund{
		baseJournal: newBaseJournal[uint64](initCapacity),
	}
	for i := 0; i < r.Size(); i++ {
		clone.Push(r.PeekAt(i))
	}
	return clone
}
