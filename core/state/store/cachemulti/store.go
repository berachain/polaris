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
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/journal"
	statetypes "github.com/berachain/stargazer/core/state/types"
)

var _ storetypes.CacheMultiStore = (*Store)(nil)

type Store struct {
	storetypes.MultiStore
	stores     map[storetypes.StoreKey]storetypes.CacheKVStore
	JournalMgr *journal.Manager
}

func NewStoreFrom(ms storetypes.MultiStore) *Store {
	return &Store{
		MultiStore: ms,
		stores:     make(map[storetypes.StoreKey]storetypes.CacheKVStore),

		JournalMgr: journal.NewManager(),
	}
}

// GetKVStore shadows cosmos sdk storetypes.MultiStore function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	if cacheKVStore, exists := s.stores[key]; exists {
		return cacheKVStore
	}
	// get kvstore from cachemultistore and set cachekv to memory
	kvstore := s.MultiStore.GetKVStore(key)
	s.stores[key] = s.newCacheKVStore(key, kvstore)
	return s.stores[key]
}

// implements cosmos sdk storetypes.CacheMultiStore
// Write commits each of the individual cachekv stores to its corresponding parent kv stores.
func (s *Store) Write() {
	// Safe from non-determinism, since order in which
	// we write to the parent kv stores does not matter.
	//
	//#nosec:G705
	for _, cacheKVStore := range s.stores {
		cacheKVStore.Write()
	}
}

func (s *Store) newCacheKVStore(
	key storetypes.StoreKey,
	kvstore storetypes.KVStore,
) storetypes.CacheKVStore {
	if key.Name() == statetypes.EvmStoreKey {
		return cachekv.NewEvmStore(kvstore, s.JournalMgr)
	}
	return cachekv.NewStore(kvstore, s.JournalMgr)
}
