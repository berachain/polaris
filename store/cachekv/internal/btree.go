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

package internal

import (
	"bytes"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tidwall/btree"
)

const (
	// The approximate number of items and children per B-tree node. Tuned with benchmarks.
	// copied from memdb.
	bTreeDegree = 32
)

var ErrKeyEmpty = errors.New("key cannot be empty")

// BTree implements the sorted cache for cachekv store,
// we don't use MemDB here because cachekv is used extensively in sdk core path,
// we need it to be as fast as possible, while `MemDB` is mainly used as a mocking db in unit tests.
//
// We choose tidwall/btree over google/btree here because it provides API to implement step iterator directly.
type BTree struct {
	tree *btree.BTreeG[item]
}

// NewBTree creates a wrapper around `btree.BTreeG`.
func NewBTree() BTree {
	return BTree{
		tree: btree.NewBTreeGOptions(byKeys, btree.Options{
			Degree:  bTreeDegree,
			NoLocks: false,
		}),
	}
}

func (bt BTree) Set(key, value []byte) {
	bt.tree.Set(newItem(key, value))
}

func (bt BTree) Get(key []byte) []byte {
	i, found := bt.tree.Get(newItem(key, nil))
	if !found {
		return nil
	}
	return i.value
}

func (bt BTree) Delete(key []byte) {
	bt.tree.Delete(newItem(key, nil))
}

//nolint:nolintlint,ireturn
func (bt BTree) Iterator(start, end []byte) (types.Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, ErrKeyEmpty
	}
	return newMemIterator(start, end, bt, true), nil
}

//nolint:nolintlint,ireturn
func (bt BTree) ReverseIterator(start, end []byte) (types.Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, ErrKeyEmpty
	}
	return newMemIterator(start, end, bt, false), nil
}

// Copy the tree. This is a copy-on-write operation and is very fast because
// it only performs a shadowed copy.
func (bt BTree) Copy() BTree {
	return BTree{
		tree: bt.tree.Copy(),
	}
}

// item is a btree item with byte slices as keys and values.
type item struct {
	key   []byte
	value []byte
}

// byKeys compares the items by key.
func byKeys(a, b item) bool {
	return bytes.Compare(a.key, b.key) == -1
}

// newItem creates a new pair item.
func newItem(key, value []byte) item {
	return item{key: key, value: value}
}
