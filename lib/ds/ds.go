// SPDX-License-Identifier: Apache-2.0
//

package ds

import (
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// `Stack` is an interface represent a FILO data structure.
type Stack[Item any] interface {
	// `Peek` returns the Item at the top of the stack
	Peek() Item

	// `PeekAt` returns the Item at the given index.
	PeekAt(index int) Item

	// `Push` adds a new Item to the top of the stack. The Size method returns the current
	// number of entries in the items.
	Push(i Item) int

	// `Pop` returns the Item at the top of the stack and removes it from the stack.
	Pop() Item

	// `PopToSize` discards all items entries after and including the given size. It returns the
	// item at index `newSize`.
	PopToSize(newSize int) Item

	// `Size` returns the current number of entries in the items.
	Size() int

	// `Capacity` returns the size of the allocated buffer for the stack.
	Capacity() int
}

// `CloneableStack` is an interface that extends `Stack` to allow for deep copying.
// As such, the items in the stack must implement `Cloneable`.
type CloneableStack[T libtypes.Cloneable[T]] interface {
	// `CloneableStack` is a `Stack`.
	Stack[T]

	// `CloneableStack` implements `Cloneable`.
	libtypes.Cloneable[CloneableStack[T]]
}
