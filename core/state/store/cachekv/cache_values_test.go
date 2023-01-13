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
	"reflect"
	"testing"

	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/core/state/store/journal"
	"github.com/berachain/stargazer/lib/utils"
)

var (
	byte0    = []byte{0}
	byte1    = []byte{1}
	byte0Str = utils.UnsafeBytesToStr(byte0)
	byte1Str = utils.UnsafeBytesToStr(byte1)
)

type CacheValueSuite struct {
	suite.Suite
	cacheKVStore *Store
}

func TestCacheValueSuite(t *testing.T) {
	suite.Run(t, new(CacheValueSuite))
}

func (s *CacheValueSuite) SetupTest() {
	parent := sdkcachekv.NewStore(dbadapter.Store{DB: dbm.NewMemDB()})
	parent.Set(byte0, byte0)
	s.cacheKVStore = NewStore(parent, journal.NewManager())
}

func (s *CacheValueSuite) TestRevertDeleteAfterNothing() {
	// delete after nothing happened to key
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// delete key: 0
	s.cacheKVStore.Delete(byte0)
	s.Require().Equal(([]byte)(nil), s.cacheKVStore.Cache[byte0Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte0Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte0Str)
	// revert delete key: 0
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().NotContains(s.cacheKVStore.Cache, byte0Str)
	s.Require().NotContains(s.cacheKVStore.UnsortedCache, byte0Str)
}

func (s *CacheValueSuite) TestRevertDeleteAfterGet() {
	// delete after Get called on key
	_ = s.cacheKVStore.Get(byte0)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// delete key: 0
	s.cacheKVStore.Delete(byte0)
	// revert delete key: 0
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(byte0, s.cacheKVStore.Cache[byte0Str].value)
	s.Require().False(s.cacheKVStore.Cache[byte0Str].dirty)
	s.Require().NotContains(s.cacheKVStore.UnsortedCache, byte0Str)
}

func (s *CacheValueSuite) TestRevertDeleteAfterSet() {
	// delete after Set called on key
	s.cacheKVStore.Set(byte1, byte1)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// delete key: 1
	s.cacheKVStore.Delete(byte1)
	// revert delete key: 1
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(byte1, s.cacheKVStore.Cache[byte1Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte1Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte1Str)
}

func (s *CacheValueSuite) TestRevertDeleteAfterDelete() {
	// delete after Delete called on key
	s.cacheKVStore.Delete(byte0)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// delete key: 0
	s.cacheKVStore.Delete(byte0)
	// revert delete key: 0
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(([]byte)(nil), s.cacheKVStore.Cache[byte0Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte0Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte0Str)
}

func (s *CacheValueSuite) TestRevertSetAfterNothing() {
	// set after nothing happened to key
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// set key: 1
	s.cacheKVStore.Set(byte1, byte1)
	s.Require().Equal(byte1, s.cacheKVStore.Cache[byte1Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte1Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte1Str)
	// revert set key: 1
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().NotContains(s.cacheKVStore.Cache, byte1Str)
	s.Require().NotContains(s.cacheKVStore.UnsortedCache, byte1Str)
}

func (s *CacheValueSuite) TestRevertSetAfterGet() {
	// set after get called on key
	_ = s.cacheKVStore.Get(byte0)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// set key: 0 to val: 1
	s.cacheKVStore.Set(byte0, byte1)
	// revert set key: 1 to val: 1
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(byte0, s.cacheKVStore.Cache[byte0Str].value)
	s.Require().False(s.cacheKVStore.Cache[byte0Str].dirty)
	s.Require().NotContains(s.cacheKVStore.UnsortedCache, byte0Str)
}

func (s *CacheValueSuite) TestRevertSetAfterDelete() {
	// set after delete called on key
	s.cacheKVStore.Delete(byte0)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// set key: 0 to val: 0
	s.cacheKVStore.Set(byte0, byte0)
	// revert set key: 0 to val: 0
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(([]byte)(nil), s.cacheKVStore.Cache[byte0Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte0Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte0Str)
}

func (s *CacheValueSuite) TestRevertSetAfterSet() {
	// set after set called on key
	s.cacheKVStore.Set(byte1, byte1)
	snapshot := s.cacheKVStore.JournalMgr().Size()
	// set key: 1 to val: 0
	s.cacheKVStore.Set(byte1, byte0)
	// revert set key: 1 to val: 0
	s.cacheKVStore.JournalMgr().PopToSize(snapshot)
	s.Require().Equal(byte1, s.cacheKVStore.Cache[byte1Str].value)
	s.Require().True(s.cacheKVStore.Cache[byte1Str].dirty)
	s.Require().Contains(s.cacheKVStore.UnsortedCache, byte1Str)
}

func (s *CacheValueSuite) TestCloneDelete() {
	dcvNonNil := NewDeleteCacheValue(s.cacheKVStore, byte1Str, NewCacheValue(byte1, true))
	dcvNonNilClone, ok := dcvNonNil.Clone().(*DeleteCacheValue)
	s.Require().True(ok)
	s.Require().Equal(byte1Str, dcvNonNilClone.Key)
	s.Require().True(dcvNonNilClone.Prev.dirty)
	s.Require().Equal(byte1, dcvNonNilClone.Prev.value)

	dcvNil := NewDeleteCacheValue(s.cacheKVStore, "", nil)
	dcvNilClone, ok := dcvNil.Clone().(*DeleteCacheValue)
	s.Require().True(ok)
	s.Require().Equal("", dcvNilClone.Key)
	s.Require().True(reflect.ValueOf(dcvNilClone.Prev).IsNil())
	// s.Require().Equal((*cValue)(nil), dcvNilClone.Prev)
}

func (s *CacheValueSuite) TestCloneSet() {
	dcvNonNil := NewSetCacheValue(s.cacheKVStore, byte1Str, NewCacheValue(byte1, true))
	dcvNonNilClone, ok := dcvNonNil.Clone().(*SetCacheValue)
	s.Require().True(ok)
	s.Require().Equal(byte1Str, dcvNonNilClone.Key)
	s.Require().True(dcvNonNilClone.Prev.dirty)
	s.Require().Equal(byte1, dcvNonNilClone.Prev.value)

	dcvNil := NewSetCacheValue(s.cacheKVStore, "", nil)
	dcvNilClone, ok := dcvNil.Clone().(*SetCacheValue)
	s.Require().True(ok)
	s.Require().Equal("", dcvNilClone.Key)
	// s.Require().Equal((*cValue)(nil), dcvNilClone.Prev)
	s.Require().True(reflect.ValueOf(dcvNilClone.Prev).IsNil())
}
