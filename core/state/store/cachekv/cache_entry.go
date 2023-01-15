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

// Compile-time check to ensure `cacheEntry` implements `journal.CacheEntry`.
var _ journal.CacheEntry = (*cacheEntry)(nil)

// `cacheEntry` is a struct that contains information needed to set a value in a cache.
type cacheEntry struct {
	Store *Store      // Pointer to the cache store.
	Key   string      // Key of the value to be set.
	Prev  *cacheValue // Deep copy of object in cache map.
}

// `newCacheEntry` creates a new `cacheEntry` object for the given `store`, `key`, and `prev`
// cache value.
func newCacheEntry(store *Store, key string, prev *cacheValue) *cacheEntry {
	// create a deep copy of the Prev field, if it is not nil.
	if prev != nil {
		prev = prev.Clone()
	}

	return &cacheEntry{
		Store: store,
		Key:   key,
		Prev:  prev,
	}
}

// `Revert` reverts a set operation on a cache entry by setting the previous value of the entry as
// the current value in the cache map.
//
// `Revert` implements journal.cacheEntry.
func (ce *cacheEntry) Revert() {
	// If there was a previous value, set it as the current value in the cache map
	if ce.Prev == nil {
		// If there was no previous value, remove the Key from the
		// cache map and the unsorted cache set
		delete(ce.Store.Cache, ce.Key)
		delete(ce.Store.UnsortedCache, ce.Key)
		return
	}

	// If there was a previous value, set it sas the current value in the cache map.
	ce.Store.Cache[ce.Key] = ce.Prev

	// If the previous value was not dirty, remove the Key from the unsorted cache set
	if !ce.Prev.dirty {
		delete(ce.Store.UnsortedCache, ce.Key)
	}
}

// `Clone` creates a deep copy of the cacheEntry object.
// The deep copy contains the same Store and Key fields as the original,
// and a deep copy of the Prev field, if it is not nil.s
//
// `Clone` implements `journal.cacheEntry`.
//
//nolint:nolintlint,ireturn // by design.
func (ce *cacheEntry) Clone() journal.CacheEntry {
	// Return a new cacheEntry object with the same Store and Key fields as the original,
	// and the Prev field set to the deep copy of the original Prev field (or nil if the original
	// was nil).
	return newCacheEntry(ce.Store, ce.Key, ce.Prev)
}
