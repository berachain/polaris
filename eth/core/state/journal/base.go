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
