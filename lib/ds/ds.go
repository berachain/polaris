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

package ds

import (
	dbm "github.com/cosmos/cosmos-db"

	libtypes "pkg.berachain.dev/stargazer/lib/types"
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

// `BTree` is an interface that defines the methods a binary tree must implement.
type BTree interface {
	// `Set` sets the key to value.
	Set(key, value []byte)

	// `Get` gets the value at key.
	Get(key []byte) []byte

	// `Delete` deletes key.
	Delete(key []byte)

	// `Iterator` returns an iterator between start and end.
	Iterator(start, end []byte) (dbm.Iterator, error)

	// `ReverseIterator` returns a reverse iterator between start and end.
	ReverseIterator(start, end []byte) (dbm.Iterator, error)

	// `Copy` returns a shallow copy of BTree.
	Copy() BTree
}
