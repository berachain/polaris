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
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var (
	_ storetypes.CacheMultiStore = (*Store)(nil)
	_ libtypes.Snapshottable     = (*Store)(nil)
)

// m
// `MSStore` is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It stores revisions by cache-wrapping the cachekv stores on a call to `Snapshot`.
type MSStore struct {
	storetypes.MultiStore

	stores ds.Stack[storetypes.CacheMultiStore]
}

// `NewMSStoreFrom` creates and returns a new `Store` from a given MultiStore.
func NewMSStoreFrom(ms storetypes.MultiStore) *MSStore {
	stores := stack.New[storetypes.CacheMultiStore](16)
	stores.Push(ms.CacheMultiStore())
	return &MSStore{
		MultiStore: ms,
		stores:     stores,
	}
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (ms *MSStore) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	return ms.stores.Peek().GetKVStore(key)
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (ms *MSStore) Snapshot() int {
	defer func() {
		ms.stores.Push(ms.stores.Peek().CacheMultiStore())
	}()

	return ms.stores.Size()
}

// `Revert` implements `libtypes.Snapshottable`.
func (ms *MSStore) RevertToSnapshot(id int) {
	ms.stores.PopToSize(id)
	if id == 0 {
		ms.stores.Push(ms.CacheMultiStore())
	}
}

// `Write` commits each of the individual cachekv stores to its corresponding parent kv stores.
//
// `Write` implements Cosmos SDK `storetypes.CacheMultiStore`.
func (ms *MSStore) Write() {
	// to allow garbage collector to vibe
	for i := ms.stores.Size() - 1; i >= 0; i-- {
		ms.stores.Pop().Write()
	}
}
