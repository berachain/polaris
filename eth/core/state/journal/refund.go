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
