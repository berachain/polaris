// SPDX-License-Identifier: Apache-2.0
//

package stack

import (
	"pkg.berachain.dev/polaris/lib/ds"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// `cloneableStack` is a struct that holds a slice of CacheEntry instances.
type cloneableStack[T libtypes.Cloneable[T]] struct {
	// The `cloneableStack` is a `ds.Stack`.
	ds.Stack[T]
}

// `NewCloneable` creates and returns a new cloneableStack instance.
func NewCloneable[T libtypes.Cloneable[T]](capacity int) cloneableStack[T] { //nolint:revive // it's ok.
	return cloneableStack[T]{
		New[T](capacity),
	}
}

// `Clone` returns a cloned journal by deep copyign each CacheEntry.
//
// `Clone` implements `CloneableStack[T]`.
func (cs cloneableStack[T]) Clone() ds.CloneableStack[T] {
	newcloneableStack := NewCloneable[T](cs.Capacity())
	for i := 0; i < cs.Size(); i++ {
		newcloneableStack.Push(cs.PeekAt(i).Clone())
	}
	return newcloneableStack
}
