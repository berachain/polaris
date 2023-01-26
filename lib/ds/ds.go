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

import dbm "github.com/tendermint/tm-db"

// `Stack` is an interface that defines the methods that an items Stack must implement.
// items Stacks support holding cache entries and reverting to a certain index.
type Stack[Item any] interface {
	// `Peek` returns the Item at the top of the stack
	Peek() Item

	// `PeekAt` returns the Item at the given index.
	PeekAt(index int) Item

	// `Push` adds a new Item to the top of the stack. The Size method returns the current
	// number of entries in the items.
	Push(i Item)

	// `Pop` returns the Item at the top of the stack and removes it from the stack.
	Pop() Item

	// `PopToSize` discards all items entries after and including the given size.
	PopToSize(newSize int)

	// `Size` returns the current number of entries in the items.
	Size() int
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
