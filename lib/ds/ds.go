// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ds

import (
	libtypes "github.com/berachain/stargazer/lib/types"
	dbm "github.com/tendermint/tm-db"
)

// `Stack` is an interface represent a FILO data structure.
type Stack[Item any] interface {
	// `Peek` returns the Item at the top of the stack
	Peek() Item

	// `PeekAt` returns the Item at the given index.
	PeekAt(index int) Item

	// `Push` adds a new Item to the top of the stack and returns the new size.
	Push(i Item) int

	// `Pop` returns the Item at the top of the stack and removes it from the stack.
	Pop() Item

	// `PopToSize` discards all items entries after and including the given size and returns the
	// last element that was popped.
	PopToSize(newSize int) Item

	// `Size` returns the current number of entries in the items.
	Size() int

	// `Capacity` returns the size of the allocated buffer for the stack.
	Capacity() int

	// `Slice` returns the stack in the form of a slice.
	Slice() []Item
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
