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

package stack

import (
	"pkg.berachain.dev/stargazer/lib/ds"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
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
