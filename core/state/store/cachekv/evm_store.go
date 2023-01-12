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

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/berachain/stargazer/core/state/store/journal"
	"github.com/berachain/stargazer/lib/utils"
)

var _ storetypes.CacheKVStore = (*EvmStore)(nil)

// Avoid the mutex lock for EVM stores (codes/storage)
// Writes to the EVM are thread-safe because the EVM interpreter is guaranteed to be single
// threaded. All entry points to the EVM check that only a single execution context is running.
type EvmStore struct {
	*Store
}

// NewEvmStore creates a new Store object.
func NewEvmStore(parent storetypes.KVStore, journalMgr *journal.Manager) *EvmStore {
	return &EvmStore{
		NewStore(parent, journalMgr),
	}
}

// Get shadows Store.Get
// This function retrieves a value associated with the specified key in the store.
func (store *EvmStore) Get(key []byte) []byte {
	var bz []byte
	// Check if the key is in the store's cache.
	cacheValue, ok := store.Cache[utils.UnsafeBytesToStr(key)]
	if ok {
		// If the key is in the cache, return the value.
		return cacheValue.value
	}

	// If the key is not found in the cache, query the parent store.
	bz = store.Parent.Get(key)

	// Add the key-value pair to the cache.
	store.setCacheValue(key, bz, false)

	return bz
}

// Set shadows Store.Set.
func (store *EvmStore) Set(key []byte, value []byte) {
	store.setCacheValue(key, value, true)
}

// Delete shadows Store.Delete.
func (store *EvmStore) Delete(key []byte) {
	store.setCacheValue(key, nil, true)
}
