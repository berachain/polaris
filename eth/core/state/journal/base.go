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
	"github.com/berachain/polaris/lib/ds"
	"github.com/berachain/polaris/lib/ds/stack"
)

// baseJournal is a struct that holds a stack of items.
type baseJournal[T any] struct {
	ds.Stack[T]
}

// newBaseJournal returns a new `baseJournal` with the given initial capacity.
func newBaseJournal[T any](initialCapacity int) baseJournal[T] {
	return baseJournal[T]{
		Stack: stack.New[T](initialCapacity),
	}
}

// Snapshot takes a snapshot of the `Logs` store.
//
// Snapshot implements `libtypes.Snapshottable`.
func (j *baseJournal[T]) Snapshot() int {
	return j.Size()
}

// RevertToSnapshot reverts the `Logs` store to a given snapshot id.
//
// RevertToSnapshot implements `libtypes.Snapshottable`.
func (j *baseJournal[T]) RevertToSnapshot(id int) {
	j.PopToSize(id)
}
