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

import "github.com/berachain/stargazer/store/journal"

type (
	// DeleteCacheValue is a struct that contains information needed to delete a value
	// from a cache.
	DeleteCacheValue struct {
		Store *Store  // pointer to the cache Store
		Key   string  // Key of the value to be deleted
		Prev  *cValue // deep copy of object in cache map
	}

	// SetCacheValue is a struct that contains information needed to set a value in a cache.
	SetCacheValue struct {
		Store *Store  // pointer to the cache Store
		Key   string  // Key of the value to be set
		Prev  *cValue // deep copy of object in cache map
	}
)

// Revert restores the previous cache entry for the Key, if it exists.
//
// implements journal.CacheEntry
func (dcv *DeleteCacheValue) Revert() {
	// Check if there is a previous entry for the Key being deleted
	if dcv.Prev != nil {
		// If the previous entry is not dirty
		// (i.e., it has not been modified since it was added to the cache)
		if !dcv.Prev.dirty {
			// Remove the Key being deleted from the cache and restore the previous entry
			delete(dcv.Store.UnsortedCache, dcv.Key)
			dcv.Store.Cache[dcv.Key] = dcv.Prev
		} else {
			// If the previous entry is dirty and has a non-nil value
			if dcv.Prev.value != nil {
				// Remove the Key being deleted from the "deleted" set and restore the
				// previous entry to the cache
				dcv.Store.Cache[dcv.Key] = dcv.Prev
			}
		}
	} else {
		// If there is no previous entry for the Key being deleted, remove the Key from the cache
		delete(dcv.Store.Cache, dcv.Key)
		delete(dcv.Store.UnsortedCache, dcv.Key)
	}
}

// Clone creates a deep copy of the DeleteCacheValue object.
// The deep copy contains the same Store and Key fields as the original,
// and a deep copy of the Prev field, if it is not nil.
//
// implements journal.CacheEntry
func (dcv *DeleteCacheValue) Clone() journal.CacheEntry {
	// Create a deep copy of the Prev field, if it is not nil
	var prevDeepCopy *cValue
	if dcv.Prev != nil {
		prevDeepCopy = dcv.Prev.deepCopy()
	}

	// Return a new DeleteCacheValue object with the same Store and Key fields as the original,
	// and the Prev field set to the deep copy of the original Prev field
	// (or nil if the original was nil)
	return &DeleteCacheValue{
		Store: dcv.Store,
		Key:   dcv.Key,
		Prev:  prevDeepCopy,
	}
}

// Revert reverts a set operation on a cache entry by setting the previous value of the entry as
// the current value in the cache map.
//
// implements journal.CacheEntry
func (scv *SetCacheValue) Revert() {
	// If there was a previous value, set it as the current value in the cache map
	if scv.Prev != nil {
		scv.Store.Cache[scv.Key] = scv.Prev
		// If the previous value was not dirty, remove the Key from the unsorted cache set
		if !scv.Prev.dirty {
			delete(scv.Store.UnsortedCache, scv.Key)
		}
	} else {
		// If there was no previous value, remove the Key from the
		// cache map and the unsorted cache set
		delete(scv.Store.Cache, scv.Key)
		delete(scv.Store.UnsortedCache, scv.Key)
	}
}

// Clone creates a deep copy of the SetCacheValue object.
// The deep copy contains the same Store and Key fields as the original,
// and a deep copy of the Prev field, if it is not nil.
//
// implements journal.CacheEntry
func (scv *SetCacheValue) Clone() journal.CacheEntry {
	// Create a deep copy of the Prev field, if it is not nil
	var prevDeepCopy *cValue
	if scv.Prev != nil {
		prevDeepCopy = scv.Prev.deepCopy()
	}

	// Return a new SetCacheValue object with the same Store and Key fields as the original,
	// and the Prev field set to the deep copy of the original Prev field
	// (or nil if the original was nil)
	return &SetCacheValue{
		Store: scv.Store,
		Key:   scv.Key,
		Prev:  prevDeepCopy,
	}
}
