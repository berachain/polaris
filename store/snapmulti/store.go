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

package snapmulti

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	"github.com/berachain/stargazer/lib/utils"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	storeRegistryKey    = `snapmultistore`
	initJournalCapacity = 32
)

// `cacheMultiStore` represents a cached multistore, which is just a map of store keys to its
// corresponding cache kv store currently being used.
type cacheMultiStore map[storetypes.StoreKey]storetypes.CacheKVStore

// `store` is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It journals revisions by cache-wrapping the cachekv stores on a call to `Snapshot`. In this
// store's lifecycle, any operations done before the first call to snapshot will be enforced on the
// root `cacheMultiStore`.
type store struct {
	storetypes.MultiStore

	// root is the cacheMultiStore used before the first snapshot is called
	root cacheMultiStore
	// journal holds the snapshots of cachemultistores
	journal ds.Stack[cacheMultiStore]
}

// `NewStoreFrom` creates and returns a new `store` from a given Multistore `ms`.
func NewStoreFrom(ms storetypes.MultiStore) *store { //nolint:revive // its okay.
	return &store{
		MultiStore: ms,
		root:       make(cacheMultiStore),
		journal:    stack.New[cacheMultiStore](initJournalCapacity),
	}
}

// `RegistryKey` implements `libtypes.Registrable`.
func (s *store) RegistryKey() string {
	return storeRegistryKey
}

// `GetCommittedKVStore` returns the KV Store from the given Multistore. This function follows
// the Multistore's normal `GetKVStore` code path.
func (s *store) GetCommittedKVStore(key storetypes.StoreKey) storetypes.KVStore {
	return s.MultiStore.GetKVStore(key)
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	var cms cacheMultiStore
	if cms = s.journal.Peek(); cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// check if cache kv store already used
	if cacheKVStore, found := cms[key]; found {
		return cacheKVStore
	}

	// get kvstore from cacheMultiStore and set cachekv to memory
	cms[key] = cachekv.NewStore(s.GetCommittedKVStore(key))
	return cms[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *store) Snapshot() int {
	var cms cacheMultiStore
	if cms = s.journal.Peek(); cms == nil {
		// use root if the journal is empty
		cms = s.root
	}

	// build revision of cms by cachewrapping each cachekv store
	revision := make(cacheMultiStore)
	for key, cacheKVStore := range cms {
		revision[key] = utils.MustGetAs[storetypes.CacheKVStore](cacheKVStore.CacheWrap())
	}

	// push the revision to the journal and return the size BEFORE snapshot
	return s.journal.Push(revision) - 1
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *store) RevertToSnapshot(id int) {
	// `id` is the new size of the journal we want to maintain.
	s.journal.PopToSize(id)
}

// `Finalize` commits each of the individual cachekv stores to its corresponding parent cachekv stores
// in the journal. Finally it commits the root cachekv stores.
//
// `Finalize` implements `libtypes.Controllable`.
func (s *store) Finalize() {
	// write each cachekv store in the journal to its parent
	revision := s.journal.Peek()
	for ; s.journal.Size() > 0; revision = s.journal.Pop() {
		// Safe from non-determinism, since order in which
		// we write to the parent kv journal does not matter.
		//
		//#nosec:G705
		for key, cacheKVStore := range revision {
			cacheKVStore.Write()
			delete(revision, key)
		}
	}

	// write the root
	for key, cacheKVStore := range s.root {
		cacheKVStore.Write()
		delete(s.root, key)
	}
}
