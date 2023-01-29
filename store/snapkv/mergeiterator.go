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

package snapkv

import (
	"bytes"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/types"
)

// `cacheMergeIterator` merges a parent Iterator and a cache Iterator.
// The cache iterator may return nil keys to signal that an item
// had been deleted (but not deleted in the parent).
// If the cache iterator has the same key as the parent, the
// cache shadows (overrides) the parent.
//
// TODO: Optimize by memoizing.
type cacheMergeIterator struct {
	parent    types.Iterator
	cache     types.Iterator
	ascending bool

	valid bool
}

var _ types.Iterator = (*cacheMergeIterator)(nil)

// `NewCacheMergeIterator` creates a new cacheMergeIterator.
func newCacheMergeIterator(parent, cache types.Iterator, ascending bool) *cacheMergeIterator {
	iter := &cacheMergeIterator{
		parent:    parent,
		cache:     cache,
		ascending: ascending,
	}

	iter.valid = iter.skipUntilExistsOrInvalid()
	return iter
}

// Domain implements Iterator.
// Returns parent domain because cache and parent domains are the same.
func (iter *cacheMergeIterator) Domain() ([]byte, []byte) {
	return iter.parent.Domain()
}

// Valid implements Iterator.
func (iter *cacheMergeIterator) Valid() bool {
	return iter.valid
}

// Next implements Iterator.
func (iter *cacheMergeIterator) Next() {
	iter.assertValid()

	switch {
	case !iter.parent.Valid():
		// If parent is invalid, get the next cache item.
		iter.cache.Next()
	case !iter.cache.Valid():
		// If cache is invalid, get the next parent item.
		iter.parent.Next()
	default:
		// Both are valid.  Compare keys.
		keyP, keyC := iter.parent.Key(), iter.cache.Key()
		switch iter.compare(keyP, keyC) {
		case -1: // parent < cache
			iter.parent.Next()
		case 0: // parent == cache
			iter.parent.Next()
			iter.cache.Next()
		case 1: // parent > cache
			iter.cache.Next()
		}
	}
	iter.valid = iter.skipUntilExistsOrInvalid()
}

// Key implements Iterator.
func (iter *cacheMergeIterator) Key() []byte {
	iter.assertValid()

	// If parent is invalid, get the cache key.
	if !iter.parent.Valid() {
		return iter.cache.Key()
	}

	// If cache is invalid, get the parent key.
	if !iter.cache.Valid() {
		return iter.parent.Key()
	}

	// Both are valid.  Compare keys.
	keyP, keyC := iter.parent.Key(), iter.cache.Key()

	cmp := iter.compare(keyP, keyC)
	switch cmp {
	case -1: // parent < cache
		return keyP
	case 0: // parent == cache
		return keyP
	case 1: // parent > cache
		return keyC
	default:
		panic("invalid compare result")
	}
}

// Value implements Iterator.
func (iter *cacheMergeIterator) Value() []byte {
	iter.assertValid()

	// If parent is invalid, get the cache value.
	if !iter.parent.Valid() {
		return iter.cache.Value()
	}

	// If cache is invalid, get the parent value.
	if !iter.cache.Valid() {
		return iter.parent.Value()
	}

	// Both are valid.  Compare keys.
	keyP, keyC := iter.parent.Key(), iter.cache.Key()

	cmp := iter.compare(keyP, keyC)
	switch cmp {
	case -1: // parent < cache
		return iter.parent.Value()
	case 0: // parent == cache
		return iter.cache.Value()
	case 1: // parent > cache
		return iter.cache.Value()
	default:
		panic("invalid comparison result")
	}
}

// Close implements Iterator.
func (iter *cacheMergeIterator) Close() error {
	err1 := iter.cache.Close()
	if err := iter.parent.Close(); err != nil {
		return err
	}

	return err1
}

// Error returns an error if the cacheMergeIterator is invalid defined by the
// Valid method.
func (iter *cacheMergeIterator) Error() error {
	if !iter.Valid() {
		return errors.New("invalid cacheMergeIterator")
	}

	return nil
}

// If not valid, panics.
// NOTE: May have side-effect of iterating over cache.
func (iter *cacheMergeIterator) assertValid() {
	if err := iter.Error(); err != nil {
		panic(err)
	}
}

// Like bytes.Compare but opposite if not ascending.
func (iter *cacheMergeIterator) compare(a, b []byte) int {
	if iter.ascending {
		return bytes.Compare(a, b)
	}

	return bytes.Compare(a, b) * -1
}

// Skip all delete-items from the cache w/ `key < until`.  After this function,
// current cache item is a non-delete-item, or `until <= key`.
// If the current cache item is not a delete item, does nothing.
// If `until` is nil, there is no limit, and cache may end up invalid.
// CONTRACT: cache is valid.
func (iter *cacheMergeIterator) skipCacheDeletes(until []byte) {
	for iter.cache.Valid() &&
		iter.cache.Value() == nil &&
		(until == nil || iter.compare(iter.cache.Key(), until) < 0) {
		iter.cache.Next()
	}
}

// Fast forwards cache (or parent+cache in case of deleted items) until current
// item exists, or until iterator becomes invalid.
// Returns whether the iterator is valid.
func (iter *cacheMergeIterator) skipUntilExistsOrInvalid() bool {
	for {
		// If parent is invalid, fast-forward cache.
		if !iter.parent.Valid() {
			iter.skipCacheDeletes(nil)
			return iter.cache.Valid()
		}
		// Parent is valid.

		if !iter.cache.Valid() {
			return true
		}
		// Parent is valid, cache is valid.

		// Compare parent and cache.
		keyP := iter.parent.Key()
		keyC := iter.cache.Key()

		switch iter.compare(keyP, keyC) {
		case -1: // parent < cache.
			return true

		case 0: // parent == cache.
			// Skip over if cache item is a delete.
			valueC := iter.cache.Value()
			if valueC == nil {
				iter.parent.Next()
				iter.cache.Next()

				continue
			}
			// Cache is not a delete.

			return true // cache exists.
		case 1: // cache < parent
			// Skip over if cache item is a delete.
			valueC := iter.cache.Value()
			if valueC == nil {
				iter.skipCacheDeletes(keyP)
				continue
			}
			// Cache is not a delete.

			return true // cache exists.
		}
	}
}
