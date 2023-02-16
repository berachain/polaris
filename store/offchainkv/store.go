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

package types

import (
	"path/filepath"

	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"

	dbm "github.com/cosmos/cosmos-db"

	pruningtypes "cosmossdk.io/store/pruning/types"

	cachekv "cosmossdk.io/store/cachekv"
)

// Compile-time checks to ensure types are correct.
var _ storetypes.CommitKVStore = (*Store)(nil)

// TODO: Upgrade this implementation to use a tree based structure to store the data off-chain.

// `Store` represents a store used for storing persistent data off-chain while also
// utilizing the multistore. We must register this store with the.
type Store struct {
	*cachekv.Store
}

// `New` creates a new store and connects it to an file system based database located at:
// <flags.FlagHome/data/<name>.
func New(name string, appOpts types.AppOptions) *Store {
	dbDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", name)
	db, err := dbm.NewDB(name, server.GetAppDBBackend(appOpts), dbDir)
	if err != nil {
		panic(err)
	}
	return &Store{
		cachekv.NewStore(&dbadapter.Store{DB: db}),
	}
}

// `Commit` commits the current store state and returns a CommitID with the new
// version and hash.
func (st *Store) Commit() storetypes.CommitID {
	// When we commit, we call the CacheKV's write, which in this implementation,
	// writes all changes in memory to the underlying database, since the parent
	// is simply a `dbadapter.Store`
	st.Write()
	return storetypes.CommitID{}
}

// `LastCommitID` implements Committer.
func (st *Store) LastCommitID() storetypes.CommitID {
	return storetypes.CommitID{}
}

// `SetPruning` panics as pruning options should be provided at initialization
// since IAVl accepts pruning options directly.
func (st *Store) SetPruning(_ pruningtypes.PruningOptions) {
	panic("cannot set pruning options on the offchainkv")
}

// `GetPruning` panics as pruning options should be provided at initialization
// since IAVl accepts pruning options directly.
func (st *Store) GetPruning() pruningtypes.PruningOptions {
	panic("cannot get pruning options on the offchainkv")
}
