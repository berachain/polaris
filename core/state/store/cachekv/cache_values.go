// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
package cachekv

import "github.com/berachain/stargazer/core/state/store/journal"

type (
	// `DeleteCacheValue` is a struct that contains information needed to delete a value
	// from a cache.
	DeleteCacheValue struct {
		Store *Store      // Pointer to the cache store.
		Key   string      // Key of the value to be deleted.
		Prev  *cacheValue // Deep copy of object in cache map.
	}

	// `SetCacheValue` is a struct that contains information needed to set a value in a cache.
	SetCacheValue struct {
		Store *Store      // Pointer to the cache store.
		Key   string      // Key of the value to be set.
		Prev  *cacheValue // Deep copy of object in cache map.
	}
)

// `NewDeleteCacheValue` creates a new `DeleteCacheValue` object for the given `store`, `key`, and
// `prev` cache value.
func NewDeleteCacheValue(store *Store, key string, prev *cacheValue) *DeleteCacheValue {
	return &DeleteCacheValue{
		Store: store,
		Key:   key,
		Prev:  prev,
	}
}

// `Revert` restores the previous cache entry for the Key, if it exists.
//
// `Revert` implements journal.CacheEntry.
func (dcv *DeleteCacheValue) Revert() {
	// If the previous entry is nil, remove the Key from the cache
	if dcv.Prev == nil {
		delete(dcv.Store.Cache, dcv.Key)
		delete(dcv.Store.UnsortedCache, dcv.Key)
		return
	}

	// If there is no previous entry for the Key being deleted, remove the Key from the cache
	// If the previous entry is not dirty
	// (i.e., it has not been modified since it was added to the cache)
	if !dcv.Prev.dirty {
		// Remove the Key being deleted from the cache and restore the previous entry
		delete(dcv.Store.UnsortedCache, dcv.Key)
		dcv.Store.Cache[dcv.Key] = dcv.Prev
		return
	}

	// If the previous entry is dirty and has a non-nil value
	if dcv.Prev.value != nil {
		// Remove the Key being deleted from the "deleted" set and restore the
		// previous entry to the cache
		dcv.Store.Cache[dcv.Key] = dcv.Prev
	}
}

// `Clone` creates a deep copy of the DeleteCacheValue object.
// The deep copy contains the same Store and Key fields as the original,
// and a deep copy of the Prev field, if it is not nil.
//
// `Clone` implements journal.CacheEntry.
//
//nolint:nolintlint,ireturn // by design.
func (dcv *DeleteCacheValue) Clone() journal.CacheEntry {
	// Create a deep copy of the Prev field, if it is not nil
	var prevDeepCopy *cacheValue
	if dcv.Prev != nil {
		prevDeepCopy = dcv.Prev.Clone()
	}

	// Return a new DeleteCacheValue object with the same Store and Key fields as the original,
	// and the Prev field set to the deep copy of the original Prev field
	// (or nil if the original was nil)
	return NewDeleteCacheValue(dcv.Store, dcv.Key, prevDeepCopy)
}

// `NewSetCacheValue` creates a new `SetCacheValue` object for the given `store`, `key`, and `prev`
// cache value.
func NewSetCacheValue(store *Store, key string, prev *cacheValue) *SetCacheValue {
	return &SetCacheValue{
		Store: store,
		Key:   key,
		Prev:  prev,
	}
}

// `Revert` reverts a set operation on a cache entry by setting the previous value of the entry as
// the current value in the cache map.
//
// `Revert` implements journal.CacheEntry.
func (scv *SetCacheValue) Revert() {
	// If there was a previous value, set it as the current value in the cache map
	if scv.Prev == nil {
		// If there was no previous value, remove the Key from the
		// cache map and the unsorted cache set
		delete(scv.Store.Cache, scv.Key)
		delete(scv.Store.UnsortedCache, scv.Key)
		return
	}

	// If there was a previous value, set it as the current value in the cache map.
	scv.Store.Cache[scv.Key] = scv.Prev

	// If the previous value was not dirty, remove the Key from the unsorted cache set
	if !scv.Prev.dirty {
		delete(scv.Store.UnsortedCache, scv.Key)
	}
}

// `Clone` creates a deep copy of the SetCacheValue object.
// The deep copy contains the same Store and Key fields as the original,
// and a deep copy of the Prev field, if it is not nil.
//
// `Clone` implements `journal.CacheEntry`.
//
//nolint:nolintlint,ireturn // by design.
func (scv *SetCacheValue) Clone() journal.CacheEntry {
	// Create a deep copy of the Prev field, if it is not nil
	var prevDeepCopy *cacheValue
	if scv.Prev != nil {
		prevDeepCopy = scv.Prev.Clone()
	}

	// Return a new SetCacheValue object with the same Store and Key fields as the original,
	// and the Prev field set to the deep copy of the original Prev field
	// (or nil if the original was nil)
	return NewSetCacheValue(scv.Store, scv.Key, prevDeepCopy)
}
