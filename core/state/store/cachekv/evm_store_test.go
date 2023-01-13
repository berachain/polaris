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

package cachekv_test

import (
	"testing"

	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state/store/cachekv"
	"github.com/berachain/stargazer/core/state/store/journal"
)

var (
	byte0           = []byte{0}
	byte1           = []byte{1}
	nonZeroCodeHash = common.BytesToHash([]byte{0x05})
	zeroCodeHash    = common.Hash{}
)

type EvmStoreSuite struct {
	suite.Suite
	parent   storetypes.KVStore
	evmStore *cachekv.EvmStore
}

func TestEvmStoreSuite(t *testing.T) {
	suite.Run(t, new(EvmStoreSuite))
	require.True(t, true)
}

func (s *EvmStoreSuite) SetupTest() {
	parent := sdkcachekv.NewStore(dbadapter.Store{DB: dbm.NewMemDB()})
	s.parent = parent
	s.evmStore = cachekv.NewEvmStore(s.parent, journal.NewManager())
}

// TODO: determine if this is allowable behavior, should be fine?
func (s *EvmStoreSuite) TestWarmSlotVia0() {
	s.SetupTest()

	// set [key: 1, val: zeroCodeHash]
	s.evmStore.Set(byte1, zeroCodeHash.Bytes())
	// write Store
	s.evmStore.Write()

	require.Equal(s.T(), zeroCodeHash.Bytes(), s.parent.Get(byte1))
}

// parent != nil, set value == zeroCodeHash -> should write key ??
func (s *EvmStoreSuite) TestWriteZeroValParentNotNil() {
	s.SetupTest()

	// set [key: 0, val: zeroCodeHash]
	s.evmStore.Set(byte0, zeroCodeHash.Bytes())
	// write Store
	s.evmStore.Write()
	// check written
	require.Equal(s.T(), zeroCodeHash.Bytes(), s.parent.Get(byte0))
}

// parent == nil, set value != zeroCodeHash -> should write key.
func (s *EvmStoreSuite) TestWriteNonZeroValParentNil() {
	s.SetupTest()

	// set [key: 1, val: nonZeroCodeHash]
	s.evmStore.Set(byte1, nonZeroCodeHash.Bytes())
	// write Store
	s.evmStore.Write()
	// check written
	require.Equal(s.T(), nonZeroCodeHash.Bytes(), s.parent.Get(byte1))
}

// parent != nil, set value != zeroCodeHash -> should write key.
func (s *EvmStoreSuite) TestWriteNonZeroValParentNotNil() {
	s.SetupTest()

	// set [key: 0, val: nonZeroCodeHash]
	s.evmStore.Set(byte0, nonZeroCodeHash.Bytes())
	// write Store
	s.evmStore.Write()
	// check written
	require.Equal(s.T(), nonZeroCodeHash.Bytes(), s.parent.Get(byte0))
}

func (s *EvmStoreSuite) TestWriteAfterDelete() {
	s.SetupTest()

	// parent store has [key: 1, value: NIL] (does not contain key: 1)
	// set [key: 1, val: zeroCodeHash] to cache evm store
	s.evmStore.Set(byte1, zeroCodeHash.Bytes())
	// the cache evm store SHOULD have value set (to zeroCodeHash) for key: 1
	require.Equal(s.T(), zeroCodeHash.Bytes(), s.evmStore.Get(byte1))
	// delete [key: 1, val: zeroCodeHash] from cache evm store
	s.evmStore.Delete(byte1)
	require.False(s.T(), s.evmStore.Has(byte1))
	s.evmStore.Write()
	// parent store SHOULD NOT be written to; stll [key: 1, value: NIL]
	require.False(s.T(), s.parent.Has(byte1))
}
