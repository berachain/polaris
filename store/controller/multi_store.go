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

package controller

import (
	"github.com/berachain/stargazer/lib/snapshot"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/store/snapkv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var (
	_ storetypes.CacheMultiStore = (*MultiStore)(nil)
	_ libtypes.Snapshottable     = (*MultiStore)(nil)
)

// SNAP KVs
// Snapshottable KV Stores approach
type MultiStore struct {
	*snapshot.Controller[libtypes.Snapshottable]
	storetypes.MultiStore

	storeKeys []string
}

func NewMultiStoreFrom(ms storetypes.MultiStore) *MultiStore {
	return &MultiStore{
		Controller: snapshot.NewController[libtypes.Snapshottable](),
		MultiStore: ms,
		storeKeys:  make([]string, 10),
	}
}

func (ms *MultiStore) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	name := key.Name()
	sst := ms.Controller.Get(name)
	if sst == nil {
		store := snapkv.NewStore(ms.MultiStore.GetKVStore(key))
		ms.Controller.Register(name, store)
		ms.storeKeys = append(ms.storeKeys, name)
		return store
	}
	return sst.(storetypes.KVStore)
}

func (ms *MultiStore) Write() {
	for _, key := range ms.storeKeys {
		ms.Controller.Get(key).(storetypes.CacheKVStore).Write()
	}
}
