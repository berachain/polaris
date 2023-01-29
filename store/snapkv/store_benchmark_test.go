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

package snapkv_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/x/evm/plugins/state/store/cachemulti"
)

func DoBenchmarkGet(b *testing.B, custom bool, keyStr string) {
	key := storetypes.NewKVStoreKey(keyStr)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	_ = cms.LoadLatestVersion()
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	if custom {
		ctx = ctx.WithMultiStore(cachemulti.NewStoreFrom(cms))
	} else {
		ctx = ctx.WithMultiStore(cms.CacheMultiStore())
	}
	store := ctx.KVStore(key)

	for i := 0; i < b.N; i++ {
		store.Set([]byte("key"), []byte("value"))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Get([]byte("key"))
	}
}

func BenchmarkGetStandardSdkCache(b *testing.B) {
	DoBenchmarkGet(b, false, "test")
}

func BenchmarkGetCustomCache(b *testing.B) {
	DoBenchmarkGet(b, true, "test")
}

func BenchmarkGetCustomEvmCache(b *testing.B) {
	DoBenchmarkGet(b, true, "evm")
}
