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
package cachemulti

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/x/evm/plugins/state/store/cachekv"
	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var _ storetypes.CacheMultiStore = (*Store)(nil)
var _ libtypes.Snapshottable = (*Store)(nil)

// `Store` is a wrapper around the Cosmos SDK `MultiStore` which injects a custom EVM CacheKVStore.
type Store struct {
	storetypes.MultiStore
	stores ds.Stack[map[storetypes.StoreKey]storetypes.CacheKVStore]
}

// `NewStoreFrom` creates and returns a new `Store` from a given MultiStore.
func NewStoreFrom(ms storetypes.MultiStore) *Store {
	return &Store{
		MultiStore: ms,
		stores:     stack.New[map[storetypes.StoreKey]storetypes.CacheKVStore](8),
	}
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	store := s.stores.Peek()
	if cacheKVStore, exists := store[key]; exists {
		return cacheKVStore
	}
	// get kvstore from cachemultistore and set cachekv to memory
	kvstore := s.MultiStore.GetKVStore(key)
	store[key] = s.newCacheKVStore(key, kvstore)
	return store[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *Store) Snapshot() int {
	stores := s.stores.Peek()
	for store := range stores {
		stores[store] = stores[store].CacheWrap().(storetypes.CacheKVStore)
	}
	return s.stores.Size()
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *Store) RevertToSnapshot(revision int) {
	s.stores.PopToSize(revision) // nolint:errcheck
}

// `Write` commits each of the individual cachekv stores to its corresponding parent kv stores.
//
// `Write` implements Cosmos SDK `storetypes.CacheMultiStore`.
func (s *Store) Write() {
	// Safe from non-determinism, since order in which
	// we write to the parent kv stores does not matter.
	//
	//#nosec:G705
	for _, cacheKVStore := range s.stores.Peek() {
		cacheKVStore.Write()
	}

	// to allow garbage collector to vibe
	for i := s.stores.Size() - 1; i >= 0; i-- {
		store := s.stores.Pop()
		for key := range store {
			delete(store, key)
		}
	}
}

// `newCacheKVStore` returns a new CacheKVStore. If the `key` is an EVM storekey, it will return
// an EVM CacheKVStore.
func (s *Store) newCacheKVStore(
	key storetypes.StoreKey,
	kvstore storetypes.KVStore,
) storetypes.CacheKVStore {
	if key.Name() == statetypes.EvmStoreKey {
		return cachekv.NewEvmStore(kvstore)
	}
	return sdkcachekv.NewStore(kvstore)
}
