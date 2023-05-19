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
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// RefundI is a `Store` that tracks the refund counter.
type RefundI interface {
	// RefundI implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// RefundI implements `libtypes.Cloneable`.
	libtypes.Cloneable[RefundI]
	// GetRefund returns the current value of the refund counter.
	GetRefund() uint64
	// AddRefund sets the refund counter to the given `gas`.
	AddRefund(gas uint64)
	// SubRefund subtracts the given `gas` from the refund counter.
	SubRefund(gas uint64)
}

// refund is a `Store` that tracks the refund counter.
type refund struct {
	ds.Stack[uint64] // journal of historical refunds.
}

// NewRefund creates and returns a `refund` journal.
func NewRefund() RefundI {
	return &refund{
		Stack: stack.New[uint64](initCapacity),
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

// Snapshot returns the current size of the refund counter, which is used to
// revert the refund counter to a previous value.
//
// Snapshot implements `libtypes.Snapshottable`.
func (r *refund) Snapshot() int {
	return r.Size()
}

// RevertToSnapshot reverts the refund counter to the value at the given `snap`.
//
// RevertToSnapshot implements `libtypes.Snapshottable`.
func (r *refund) RevertToSnapshot(id int) {
	r.PopToSize(id)
}

// Finalize implements `libtypes.Controllable`.
func (r *refund) Finalize() {
	r.Stack = stack.New[uint64](initCapacity)
}

// Clone implements `libtypes.Cloneable`.
func (r *refund) Clone() RefundI {
	size := r.Size()
	copy := &refund{
		Stack: stack.New[uint64](size),
	}
	for i := 0; i < size; i++ {
		copy.Push(r.PeekAt(i))
	}
	return copy
}
