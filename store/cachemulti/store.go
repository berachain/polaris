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
	libtypes "github.com/berachain/stargazer/lib/types"
	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var _ storetypes.CacheMultiStore = (*Store)(nil)
var _ libtypes.Snapshottable = (*Store)(nil)

// `Store` is a wrapper around the Cosmos SDK `MultiStore` which injects a custom EVM CacheKVStore.
type Store struct {
	storetypes.MultiStore
	stores     map[storetypes.StoreKey]storetypes.CacheKVStore
	revisionID int
}

// `NewStoreFrom` creates and returns a new `Store` from a given MultiStore.
func NewStoreFrom(ms storetypes.MultiStore) *Store {
	return &Store{
		MultiStore: ms,
		stores:     make(map[storetypes.StoreKey]storetypes.CacheKVStore),
	}
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	if cacheKVStore, exists := s.stores[key]; exists {
		return cacheKVStore
	}
	// get kvstore from cachemultistore and set cachekv to memory
	kvstore := s.MultiStore.GetKVStore(key)
	s.stores[key] = sdkcachekv.NewStore(kvstore)
	return s.stores[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *Store) Snapshot() int {
	for key := range s.stores {
		s.stores[key] = s.stores[key].CacheWrap().(storetypes.CacheKVStore)
	}
	snapshot := s.revisionID
	s.revisionID++
	return snapshot
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *Store) RevertToSnapshot(revision int) {
	for key := range s.stores {
		revertTo := s.stores[key]
		for i := s.revisionID; i >= revision; i-- {
			revertTo = revertTo.
		}
	}

	for i := s.revisionID; i >= revision; i-- {
		for key := range s.stores {
			s.stores[key] = s.stores[key].CacheWrap().(storetypes.CacheKVStore)
		}
	}
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
