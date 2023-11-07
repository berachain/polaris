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

package ds

import (
	libtypes "github.com/berachain/polaris/lib/types"
)

type StackFactory[T any] interface {
	// NewStack returns a new stack with the given initial capacity.
	NewStack() Stack[T]
}

// Stack is an interface represent a FILO data structure.
type Stack[Item any] interface {
	// Peek returns the Item at the top of the stack
	Peek() Item

	// PeekAt returns the Item at the given index.
	PeekAt(index int) Item

	// Push adds a new Item to the top of the stack and returns the size of the stack after the
	// push.
	Push(i Item) int

	// Pop returns the Item at the top of the stack and removes it from the stack.
	Pop() Item

	// PopToSize discards all items entries after and including the given size. It returns the
	// item at index `newSize`.
	PopToSize(newSize int) Item

	// Size returns the current number of entries in the items.
	Size() int

	// Capacity returns the size of the allocated buffer for the stack.
	Capacity() int
}

// CloneableStack is an interface that extends `Stack` to allow for deep copying.
// As such, the items in the stack must implement `Cloneable`.
type CloneableStack[T libtypes.Cloneable[T]] interface {
	// CloneableStack is a `Stack`.
	Stack[T]

	// CloneableStack implements `Cloneable`.
	libtypes.Cloneable[CloneableStack[T]]
}
