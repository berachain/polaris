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

package fake

import (
	"bytes"
	"io"
	"sort"
	"sync"

	"github.com/berachain/stargazer/testutil/mock/interfaces"
	"github.com/berachain/stargazer/testutil/mock/interfaces/mock"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MultiStore is a simple multistore used for testing.
type MultiStore struct {
	kvstore map[string]interfaces.KVStore
	*mock.MultiStoreMock
}

// CachedMultiStore is a simple chached multistore for testing.
type CachedMultiStore struct {
	kvstore map[string]interfaces.KVStore
	*mock.CacheMultiStoreMock
}

// NewMultiStore returns a new Multistore instance used for testing.
func NewMultiStore() sdk.MultiStore {
	ms := MultiStore{
		kvstore:        map[string]interfaces.KVStore{},
		MultiStoreMock: &mock.MultiStoreMock{},
	}
	ms.GetKVStoreFunc = func(storeKey types.StoreKey) types.KVStore {
		if store, ok := ms.kvstore[storeKey.String()]; ok {
			return store
		}
		store := newTestKVStore()
		ms.kvstore[storeKey.String()] = store
		return store
	}

	ms.CacheMultiStoreFunc = func() types.CacheMultiStore {
		return NewCachedMultiStore(ms)
	}

	return ms
}

// NewCachedMultiStore returns a new CacheMultiStore instance for testing.
func NewCachedMultiStore(ms MultiStore) types.CacheMultiStore {
	kvstore := map[string]interfaces.KVStore{}

	for key, store := range ms.kvstore {
		var ok bool
		kvstore[key], ok = store.CacheWrap().(interfaces.KVStore)
		if !ok {
			panic("not a kv store")
		}
	}

	cached := CachedMultiStore{
		kvstore:             kvstore,
		CacheMultiStoreMock: &mock.CacheMultiStoreMock{},
	}
	cached.GetKVStoreFunc = func(storeKey types.StoreKey) types.KVStore {
		if store, ok := cached.kvstore[storeKey.String()]; ok {
			return store
		}
		store := newTestKVStore()
		store.write = func() {
			ms.kvstore[storeKey.String()] = store
		}
		cached.kvstore[storeKey.String()] = store
		return store
	}

	cached.WriteFunc = func() {
		for _, store := range cached.kvstore {
			store.(*TestKVStore).Write()
		}
	}
	return cached
}

// TestKVStore is a kv store for testing.
type TestKVStore struct {
	mutex *sync.RWMutex
	store map[string][]byte
	write func()
}

func (t TestKVStore) Write() {
	t.write()
}

// newTestKVStore returns a new kv store instance for testing.
func newTestKVStore() *TestKVStore {
	return &TestKVStore{
		mutex: &sync.RWMutex{},
		store: map[string][]byte{},
		write: func() {},
	}
}

// GetStoreType is not implemented.
func (t TestKVStore) GetStoreType() types.StoreType {
	panic("implement me")
}

// CacheWrap is not implemented.
func (t *TestKVStore) CacheWrap() types.CacheWrap {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	cache := &TestKVStore{
		mutex: &sync.RWMutex{},
		store: map[string][]byte{},
	}
	cache.write = func() { t.store = cache.store }

	for key, val := range t.store {
		cache.store[key] = val
	}

	return cache
}

// CacheWrapWithTrace is not implemented.
func (t TestKVStore) CacheWrapWithTrace(_ io.Writer, _ sdk.TraceContext) types.CacheWrap {
	panic("implement me")
}

// CacheWrapWithListeners is not implemented.
func (t TestKVStore) CacheWrapWithListeners(storeKey types.StoreKey, listeners []types.WriteListener) types.CacheWrap {
	panic("implement me")
}

// Get returns the value of the given key, nil if it does not exist.
func (t TestKVStore) Get(key []byte) []byte {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	val, ok := t.store[string(key)]

	if !ok {
		return nil
	}
	return val
}

// Has checks if an entry for the given key exists.
func (t TestKVStore) Has(key []byte) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	_, ok := t.store[string(key)]
	return ok
}

// Set stores the given key value pair.
func (t TestKVStore) Set(key, value []byte) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.store[string(key)] = value
}

// Delete deletes a key if it exists.
func (t TestKVStore) Delete(key []byte) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	delete(t.store, string(key))
}

// Iterator returns an interator over the given key domain.
func (t TestKVStore) Iterator(start, end []byte) sdk.Iterator {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return newMockIterator(start, end, t.store)
}

// ReverseIterator returns an iterator that iterates over all keys in the given domain in reverse order.
func (t TestKVStore) ReverseIterator(start, end []byte) sdk.Iterator {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	iter := newMockIterator(start, end, t.store)

	// reverse the order of the iterator, which is returned already
	// sorted in ascending order
	for i, j := 0, len(iter.keys)-1; i < j; i, j = i+1, j-1 {
		iter.keys[i], iter.keys[j] = iter.keys[j], iter.keys[i]
		iter.values[i], iter.values[j] = iter.values[j], iter.values[i]
	}

	iter.start = end
	iter.end = start

	return iter
}

// fake iterator.
type mockIterator struct {
	keys       [][]byte
	values     [][]byte
	index      int
	start, end []byte
}

func newMockIterator(start, end []byte, content map[string][]byte) *mockIterator {
	keys := make([][]byte, 0)

	// select the keys according to the specified domain
	for k := range content {
		b := []byte(k)

		if (start == nil && end == nil) || (bytes.Compare(b, start) >= 0 && bytes.Compare(b, end) < 0) {
			// make sure data is a copy so that there is no concurrent writing
			temp := make([]byte, len(k))
			copy(temp, k)
			keys = append(keys, temp)
		}
	}

	// Sort the keys in ascending order
	sort.Slice(keys, func(i, j int) bool {
		return bytes.Compare(keys[i], keys[j]) < 0
	})

	// With the keys chosen and sorted, we can now populate the slice of values
	values := make([][]byte, len(keys))

	for i := 0; i < len(keys); i++ {
		// make sure data is a copy so that there is no concurrent writing
		value := content[string(keys[i])]
		temp := make([]byte, len(value))
		copy(temp, value)

		values[i] = temp
	}

	return &mockIterator{
		keys:   keys,
		values: values,
		index:  0,
		start:  start,
		end:    end,
	}
}

// Domain returns the key domain of the iterator.
// The start & end (exclusive) limits to iterate over.
// If end < start, then the Iterator goes in reverse order.
//
// A domain of ([]byte{12, 13}, []byte{12, 14}) will iterate
// over anything with the prefix []byte{12, 13}.
//
// The smallest key is the empty byte array []byte{} - see BeginningKey().
// The largest key is the nil byte array []byte(nil) - see EndingKey().
// CONTRACT: start, end readonly []byte.
func (mi mockIterator) Domain() ([]byte, []byte) {
	return mi.start, mi.end
}

// Valid returns whether the current position is valid.
// Once invalid, an Iterator is forever invalid.
func (mi mockIterator) Valid() bool {
	return mi.index < len(mi.keys)
}

// Next moves the iterator to the next sequential key in the database, as
// defined by order of iteration.
// If Valid returns false, this method will panic.
func (mi *mockIterator) Next() {
	mi.index++
}

// Key returns the key of the cursor.
// If Valid returns false, this method will panic.
// CONTRACT: key readonly []byte.
func (mi mockIterator) Key() []byte {
	if !mi.Valid() {
		panic("Iterator position out of bounds")
	}

	return mi.keys[mi.index]
}

// Value returns the value of the cursor.
// If Valid returns false, this method will panic.
// CONTRACT: value readonly []byte.
func (mi mockIterator) Value() []byte {
	if !mi.Valid() {
		panic("Iterator position out of bounds")
	}

	return mi.values[mi.index]
}

func (mi mockIterator) Error() error {
	return nil
}

// Close releases the Iterator.
func (mi mockIterator) Close() error {
	return nil
}
