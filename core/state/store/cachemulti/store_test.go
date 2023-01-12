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

package cachemulti_test

import (
	"reflect"
	"testing"

	sdkcachemulti "github.com/cosmos/cosmos-sdk/store/cachemulti"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/core/state/store/cachemulti"
)

var (
	byte1       = []byte{1}
	evmStoreKey = storetypes.NewKVStoreKey("evm")
	accStoreKey = storetypes.NewKVStoreKey("acc")
)

type CacheMultiSuite struct {
	suite.Suite
	ms storetypes.MultiStore
}

func TestCacheMultiSuite(t *testing.T) {
	suite.Run(t, new(CacheMultiSuite))
}

func (s *CacheMultiSuite) SetupTest() {
	stores := map[storetypes.StoreKey]storetypes.CacheWrapper{
		evmStoreKey: dbadapter.Store{DB: dbm.NewMemDB()},
		accStoreKey: dbadapter.Store{DB: dbm.NewMemDB()},
	}
	s.ms = sdkcachemulti.NewStore(
		dbm.NewMemDB(),
		stores, map[string]storetypes.StoreKey{},
		nil,
		nil,
	)
}

func (s *CacheMultiSuite) TestCorrectStoreType() {
	s.SetupTest()

	cms := cachemulti.NewStoreFrom(s.ms)
	evmStore := cms.GetKVStore(evmStoreKey)
	evmStoreType := reflect.TypeOf(evmStore).String()
	require.Equal(s.T(), "*cachekv.EvmStore", evmStoreType)

	accStore := cms.GetKVStore(accStoreKey)
	accStoreType := reflect.TypeOf(accStore).String()
	require.Equal(s.T(), "*cachekv.Store", accStoreType)
}

func (s *CacheMultiSuite) TestWriteToParent() {
	accStoreParent := s.ms.GetKVStore(accStoreKey)
	evmStoreParent := s.ms.GetKVStore(evmStoreKey)
	accStoreParent.Set(byte1, byte1)

	// simulate writes to cache store as it would through context
	cms := cachemulti.NewStoreFrom(s.ms)
	cms.GetKVStore(evmStoreKey).Set(byte1, byte1)
	cms.GetKVStore(accStoreKey).Delete(byte1)

	// should not write to parent
	require.False(s.T(), evmStoreParent.Has(byte1))
	require.Equal(s.T(), byte1, accStoreParent.Get(byte1))

	cms.Write()

	// accStore should be empty, evmStore should have key: 1 to val: 1
	require.False(s.T(), accStoreParent.Has(byte1))
	require.Equal(s.T(), byte1, evmStoreParent.Get(byte1))
}

func (s *CacheMultiSuite) TestGetCachedStore() {
	accStoreParent := s.ms.GetKVStore(accStoreKey)
	cms := cachemulti.NewStoreFrom(s.ms)
	accStoreCache := cms.GetKVStore(accStoreKey)
	accStoreCache.Set(byte1, byte1)

	// check that accStoreCache is not equal to accStoreParent
	require.True(s.T(), accStoreCache.Has(byte1))
	require.False(s.T(), accStoreParent.Has(byte1))

	// check that getting accStore from cms is not the same as parent
	accStoreCache2 := cms.GetKVStore(accStoreKey)
	require.True(s.T(), accStoreCache2.Has(byte1))
}
