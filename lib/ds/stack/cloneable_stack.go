// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package stack

import (
	"github.com/berachain/polaris/lib/ds"
	libtypes "github.com/berachain/polaris/lib/types"
)

// cloneableStack is a struct that holds a slice of CacheEntry instances.
type cloneableStack[T libtypes.Cloneable[T]] struct {
	// The `cloneableStack` is a `ds.Stack`.
	ds.Stack[T]
}

// NewCloneable creates and returns a new cloneableStack instance.
func NewCloneable[T libtypes.Cloneable[T]](
	capacity int,
) cloneableStack[T] { //nolint:revive // it's ok.
	return cloneableStack[T]{
		New[T](capacity),
	}
}

// Clone returns a cloned journal by deep copyign each CacheEntry.
//
// Clone implements `CloneableStack[T]`.
func (cs cloneableStack[T]) Clone() ds.CloneableStack[T] {
	newcloneableStack := NewCloneable[T](cs.Capacity())
	for i := 0; i < cs.Size(); i++ {
		newcloneableStack.Push(cs.PeekAt(i).Clone())
	}
	return newcloneableStack
}
