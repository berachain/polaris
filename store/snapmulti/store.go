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
	"github.com/berachain/stargazer/x/evm/plugins/state/types"
	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

const storeRegistryKey = `snapmultistore`

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var _ types.ControllableMultiStore = (*Store)(nil)

// `Store` is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It journal revisions by cache-wrapping the cachekv journal on a call to `Snapshot`.
type Store struct {
	storetypes.MultiStore
	journal ds.Stack[map[storetypes.StoreKey]storetypes.CacheKVStore]
}

// `NewStoreFrom` creates and returns a new `Store` from a given MultiStore.
func NewStoreFrom(ms storetypes.MultiStore) *Store {
	journal := stack.New[map[storetypes.StoreKey]storetypes.CacheKVStore](32)
	journal.Push(make(map[storetypes.StoreKey]storetypes.CacheKVStore))
	return &Store{
		MultiStore: ms,
		journal:    journal,
	}
}

func (s *Store) RegistryKey() string {
	return storeRegistryKey
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	curr := s.journal.Peek()
	if cacheKVStore, exists := curr[key]; exists {
		return cacheKVStore
	}

	// get kvstore from cachemultistore and set cachekv to memory
	curr[key] = sdkcachekv.NewStore(s.MultiStore.GetKVStore(key))
	return curr[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *Store) Snapshot() int {
	curr := s.journal.Peek()
	next := make(map[storetypes.StoreKey]storetypes.CacheKVStore)
	for key := range curr {
		next[key] = utils.MustGetAs[storetypes.CacheKVStore](curr[key].CacheWrap())
	}
	defer func() {
		s.journal.Push(next)
	}()

	return s.journal.Size()
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *Store) RevertToSnapshot(id int) {
	s.journal.PopToSize(id)
	if id == 0 {
		s.journal.Push(make(map[storetypes.StoreKey]storetypes.CacheKVStore))
	}
}

// `Write` commits each of the individual cachekv journal to its corresponding parent kv journal.
//
// `Write` implements Cosmos SDK `storetypes.CacheMultiStore`.
func (s *Store) Write() {
	for i := s.journal.Size() - 1; i >= 0; i-- {
		revision := s.journal.Pop()

		// Safe from non-determinism, since order in which
		// we write to the parent kv journal does not matter.
		//
		//#nosec:G705
		for key, store := range revision {
			store.Write()
			delete(revision, key)
		}
	}
}
