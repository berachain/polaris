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
package snapkv

import (
	"bytes"
	"io"
	"sort"
	"sync"

	"github.com/cosmos/cosmos-sdk/store/listenkv"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/tendermint/tendermint/libs/math"
	dbm "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	"github.com/berachain/stargazer/lib/ds/trees"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/store/snapkv/internal"

	libtypes "github.com/berachain/stargazer/lib/types"
)

type SnapshotKVStore interface {
	storetypes.CacheKVStore
	libtypes.Snapshottable
	GetParent() storetypes.KVStore
}

// Compile-time type assertion.
var _ SnapshotKVStore = (*Store)(nil)

// Store wraps an in-memory cache around an underlying storetypes.KVStore.
// If a cached value is nil but deleted is defined for the corresponding key,
// it means the parent doesn't have the key. (No need to delete upon Write()).
type Store struct {
	mtx           sync.RWMutex
	cache         map[string]*cacheValue
	unsortedCache map[string]struct{}
	sortedCache   ds.BTree // always ascending sorted
	parent        storetypes.KVStore
	journal       ds.CloneableStack[CacheEntry]
}

// NewStore creates a new Store object.
func NewStore(parent storetypes.KVStore, journal ds.CloneableStack[CacheEntry]) *Store {
	return &Store{
		cache:         make(map[string]*cacheValue),
		unsortedCache: make(map[string]struct{}),
		sortedCache:   trees.NewBTree(),
		parent:        parent,
		journal:       journal,
	}
}

// GetStoreType implements storetypes.KVStore.
func (store *Store) GetStoreType() storetypes.StoreType {
	return store.parent.GetStoreType()
}

// Get implements storetypes.KVStore.
// This function retrieves a value associated with the specified key in the store.
func (store *Store) Get(key []byte) []byte {
	var bz []byte
	store.mtx.RLock()
	defer store.mtx.RUnlock()
	// Assert that the key is valid.
	storetypes.AssertValidKey(key)

	// Check if the key is in the store's cache.
	if cacheValue, found := store.cache[utils.UnsafeBytesToStr(key)]; !found {
		bz = store.parent.Get(key)
		store.setCacheValue(key, bz, false)
	} else {
		bz = cacheValue.value
	}

	return bz
}

func (store *Store) GetParent() storetypes.KVStore {
	return store.parent
}

// Set implements storetypes.KVStore.
func (store *Store) Set(key []byte, value []byte) {
	storetypes.AssertValidKey(key)
	storetypes.AssertValidValue(value)
	store.mtx.Lock()
	defer store.mtx.Unlock()
	store.setCacheValue(key, value, true)
}

// Has implements storetypes.KVStore.
func (store *Store) Has(key []byte) bool {
	value := store.Get(key)
	return value != nil
}

// Delete implements storetypes.KVStore.
func (store *Store) Delete(key []byte) {
	storetypes.AssertValidKey(key)

	store.mtx.Lock()
	defer store.mtx.Unlock()
	store.setCacheValue(key, nil, true)
}

// Implements Cachetypes.KVStore.
func (store *Store) Write() {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	if len(store.cache) == 0 && len(store.unsortedCache) == 0 {
		store.sortedCache = trees.NewBTree()
		return
	}

	// We need a copy of all of the keys.
	// Not the best, but probably not a bottleneck depending.
	keys := make([]string, 0, len(store.cache))

	for key, dbValue := range store.cache {
		if dbValue.dirty {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	// TODO: Consider allowing usage of Batch, which would allow the write to
	// at least happen atomically.
	for _, key := range keys {
		// We use []byte(key) instead of utils.UnsafeStrToBytes because we cannot
		// be sure if the underlying store might do a save with the byteslice or
		// not. Once we get confirmation that .Delete is guaranteed not to
		// save the byteslice, then we can assume only a read-only copy is sufficient.
		cacheValue := store.cache[key]
		if cacheValue.value != nil {
			// It already exists in the parent, hence update it.
			store.parent.Set([]byte(key), cacheValue.value)
		} else {
			store.parent.Delete([]byte(key))
		}
	}

	// Clear the journal entries
	// Todo: size this properly / consider using a pool.
	store.journal = stack.NewCloneable[CacheEntry](store.journal.Size())

	// Clear the cache using the map clearing idiom
	// and not allocating fresh objects.
	// Please see https://bencher.orijtech.com/perfclinic/mapclearing/
	for key := range store.cache {
		delete(store.cache, key)
	}
	for key := range store.unsortedCache {
		delete(store.unsortedCache, key)
	}

	store.sortedCache = trees.NewBTree()
}

// CacheWrap implements CacheWrapper.
func (store *Store) CacheWrap() storetypes.CacheWrap {
	return NewStore(store, store.journal.Clone())
}

// CacheWrapWithTrace implements the CacheWrapper interface.
func (store *Store) CacheWrapWithTrace(
	w io.Writer,
	tc storetypes.TraceContext,
) storetypes.CacheWrap {
	return NewStore(tracekv.NewStore(store, w, tc), store.journal.Clone())
}

// CacheWrapWithListeners implements the CacheWrapper interface.
func (store *Store) CacheWrapWithListeners(
	storeKey storetypes.StoreKey,
	listeners []storetypes.WriteListener,
) storetypes.CacheWrap {
	return NewStore(listenkv.NewStore(store, storeKey, listeners), store.journal.Clone())
}

// ====================================================================
// types.Snapshottable
// ====================================================================

// `Snapshot` implements `libtypes.Snapshottable`.
func (store *Store) Snapshot() int {
	return store.journal.Size()
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (store *Store) RevertToSnapshot(revision int) {
	journal := store.journal
	// Revert and discard all journal entries with an index greater than or equal to revision.
	for i := journal.Size() - 1; i >= revision; i-- {
		store.RevertEntry(journal.PeekAt(i))
	}
	journal.PopToSize(revision)
}

// `RevertEntry` reverts a set operation on a cache entry by setting the previous value of the entry as
// the current value in the cache map.
func (store *Store) RevertEntry(ce CacheEntry) {
	key := ce.Key()
	prev := ce.Prev()
	// If there was a previous value, set it as the current value in the cache map
	if prev == nil {
		// If there was no previous value, remove the Key from the
		// cache map and the unsorted cache set
		delete(store.cache, key)
		delete(store.unsortedCache, key)
		return
	}

	// If there was a previous value, set it as the current value in the cache map.
	store.cache[key] = prev

	// If the previous value was not dirty, remove the Key from the unsorted cache set
	if !prev.dirty {
		delete(store.unsortedCache, key)
	}
}

// ================================================
// Iteration

// Iterator implements storetypes.KVStore.
func (store *Store) Iterator(start, end []byte) storetypes.Iterator {
	return store.iterator(start, end, true)
}

// ReverseIterator implements storetypes.KVStore.
func (store *Store) ReverseIterator(start, end []byte) storetypes.Iterator {
	return store.iterator(start, end, false)
}

func (store *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	store.dirtyItems(start, end)
	isoSortedCache := store.sortedCache.Copy()

	var (
		err           error
		parent, cache types.Iterator
	)

	if ascending {
		parent = store.parent.Iterator(start, end)
		cache, err = isoSortedCache.Iterator(start, end)
	} else {
		parent = store.parent.ReverseIterator(start, end)
		cache, err = isoSortedCache.ReverseIterator(start, end)
	}
	if err != nil {
		panic(err)
	}

	return internal.NewCacheMergeIterator(parent, cache, ascending)
}

// Constructs a slice of dirty items, to use w/ memIterator.
func (store *Store) dirtyItems(start, end []byte) {
	startStr, endStr := utils.UnsafeBytesToStr(start), utils.UnsafeBytesToStr(end)
	if end != nil && startStr > endStr {
		// Nothing to do here.
		return
	}

	n := len(store.unsortedCache)
	unsorted := make([]*kv.Pair, 0)
	// If the unsortedCache is too big, its costs too much to determine
	// whats in the subset we are concerned about.
	// If you are interleaving iterator calls with writes, this can easily become an
	// O(N^2) overhead.
	// Even without that, too many range checks eventually becomes more expensive
	// than just not having the cache.
	if n < minSortSize {
		for key := range store.unsortedCache {
			// dbm.IsKeyInDomain is nil safe and returns true iff key is greater than start
			if dbm.IsKeyInDomain(utils.UnsafeStrToBytes(key), start, end) {
				cacheValue := store.cache[key]
				unsorted = append(unsorted, &kv.Pair{Key: []byte(key), Value: cacheValue.value})
			}
		}
		store.clearUnsortedCacheSubset(unsorted, stateUnsorted)
		return
	}

	// Otherwise it is large so perform a modified binary search to find
	// the target ranges for the keys that we should be looking for.
	strL := make([]string, 0, n)
	for key := range store.unsortedCache {
		strL = append(strL, key)
	}
	sort.Strings(strL)

	// Now find the values within the domain
	//  [start, end)
	startIndex := findStartIndex(strL, startStr)
	if startIndex < 0 {
		startIndex = 0
	}

	var endIndex int
	if end == nil {
		endIndex = len(strL) - 1
	} else {
		endIndex = findEndIndex(strL, endStr)
	}
	if endIndex < 0 {
		endIndex = len(strL) - 1
	}

	// Since we spent cycles to sort the values, we should process and remove a reasonable amount
	// ensure start to end is at least minSortSize in size
	// if below minSortSize, expand it to cover additional values
	// this amortizes the cost of processing elements across multiple calls
	if endIndex-startIndex < minSortSize {
		endIndex = math.MinInt(startIndex+minSortSize, len(strL)-1)
		if endIndex-startIndex < minSortSize {
			startIndex = math.MaxInt(endIndex-minSortSize, 0)
		}
	}

	kvL := make([]*kv.Pair, 0, 1+endIndex-startIndex)
	for i := startIndex; i <= endIndex; i++ {
		key := strL[i]
		cacheValue := store.cache[key]
		kvL = append(kvL, &kv.Pair{Key: []byte(key), Value: cacheValue.value})
	}

	// kvL was already sorted so pass it in as is.
	store.clearUnsortedCacheSubset(kvL, stateAlreadySorted)
}

func (store *Store) clearUnsortedCacheSubset(unsorted []*kv.Pair, sortState sortState) {
	n := len(store.unsortedCache)
	if len(unsorted) == n { // This pattern allows the Go compiler to emit the map clearing idiom for the entire map.
		for key := range store.unsortedCache {
			delete(store.unsortedCache, key)
		}
	} else { // Otherwise, normally delete the unsorted keys from the map.
		for _, kv := range unsorted {
			delete(store.unsortedCache, utils.UnsafeBytesToStr(kv.Key))
		}
	}

	if sortState == stateUnsorted {
		sort.Slice(unsorted, func(i, j int) bool {
			return bytes.Compare(unsorted[i].Key, unsorted[j].Key) < 0
		})
	}

	for _, item := range unsorted {
		// sortedCache is able to store `nil` value to represent deleted items.
		store.sortedCache.Set(item.Key, item.Value)
	}
}

// ================================================
// etc

// Only entrypoint to mutate store.Cache.
func (store *Store) setCacheValue(key, value []byte, dirty bool) {
	keyStr := utils.UnsafeBytesToStr(key)

	// Append a new journal entry if the value is dirty, in order to remember the previous state.
	// Also add the key to the unsorted cache.
	if dirty {
		store.journal.Push(newCacheEntry(keyStr, store.cache[keyStr]))
		store.unsortedCache[keyStr] = struct{}{}
	}

	// Cache the value for the key.
	store.cache[keyStr] = newCacheValue(value, dirty)
}
