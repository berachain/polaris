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

package snapkv

import (
	"bytes"
	"fmt"
	"testing"

	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	dbm "github.com/tendermint/tm-db"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnapKV(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "store/snapkv")
}

func newParent() types.CacheKVStore {
	return sdkcachekv.NewStore(dbadapter.Store{DB: dbm.NewMemDB()})
}

func newCacheKVStoreFromParent(parent types.CacheKVStore) types.CacheKVStore {
	return NewStore(parent)
}

func TestSdkConsistency(t *testing.T) {
	sdkParent := dbadapter.Store{DB: dbm.NewMemDB()}
	sdkCacheKV := sdkcachekv.NewStore(sdkParent)
	parent := newParent()
	cacheKV := newCacheKVStoreFromParent(parent)

	// do an op, test the iterator
	for i := 0; i < 2000; i++ {
		doRandomOp(t, cacheKV, sdkCacheKV, 1000)
		assertIterateDomainCompare(t, cacheKV, sdkCacheKV)
	}
}

func TestGetStoreType(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)
	require.Equal(t, parent.GetStoreType(), st.GetStoreType())
}

func TestHas(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)
	st.Set(keyFmt(1), keyFmt(1))
	require.True(t, st.Has(keyFmt(1)))
	require.False(t, parent.Has(keyFmt(1)))
	st.Write()
	require.True(t, parent.Has(keyFmt(1)))
}

func TestCacheKVReverseIterator(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	// Use the parent to check values on the merge iterator
	setRange(t, st, parent, 0, 40)
	st.Write()

	itr1 := st.ReverseIterator(nil, nil)
	itr2 := parent.ReverseIterator(nil, nil)
	checkIterators(t, itr1, itr2)
}

func TestCacheWrap(t *testing.T) {
	st := NewStore(newParent())

	// test before initializing cache wraps
	st.Set(keyFmt(1), valFmt(1))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))

	stWrap, ok := st.CacheWrap().(types.KVStore)
	require.True(t, ok)
	stTrace, ok := st.CacheWrapWithTrace(
		bytes.NewBuffer(nil),
		types.TraceContext(map[string]interface{}{"blockHeight": 64}),
	).(types.KVStore)
	require.True(t, ok)
	stListeners, ok := st.CacheWrapWithListeners(
		types.NewKVStoreKey("acc"),
		[]types.WriteListener{},
	).(types.KVStore)
	require.True(t, ok)

	// test after initializing cache wraps
	st.Set(keyFmt(2), valFmt(2))
	require.Equal(t, valFmt(2), st.Get(keyFmt(2)))

	// both keys 1 and 2 should be set in all cache wraps
	require.Equal(t, valFmt(1), stWrap.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), stTrace.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), stListeners.Get(keyFmt(1)))
	require.Equal(t, valFmt(2), stWrap.Get(keyFmt(2)))
	require.Equal(t, valFmt(2), stTrace.Get(keyFmt(2)))
	require.Equal(t, valFmt(2), stListeners.Get(keyFmt(2)))
}

// Tests below taken from Cosmos SDK and adapted to our custom CacheKVStore to ensure
// the same logical behavior

func bz(s string) []byte { return []byte(s) }

func keyFmt(i int) []byte { return bz(fmt.Sprintf("key%0.8d", i)) }

func valFmt(i int) []byte { return bz(fmt.Sprintf("value%0.8d", i)) }

func TestCacheKVStore(t *testing.T) {
	parent := newParent()
	st := newCacheKVStoreFromParent(parent)

	require.Empty(t, st.Get(keyFmt(1)), "Expected `key1` to be empty")

	// put something in parent and in cache
	parent.Set(keyFmt(1), valFmt(1))
	st.Set(keyFmt(1), valFmt(1))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))

	// update it in cache, shoudn't change parent
	st.Set(keyFmt(1), valFmt(2))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), parent.Get(keyFmt(1)))

	// write it. should change parent
	st.Write()
	require.Equal(t, valFmt(2), parent.Get(keyFmt(1)))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))

	// more writes and checks
	st.Write()
	st.Write()
	require.Equal(t, valFmt(2), parent.Get(keyFmt(1)))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))

	// make a new one, check it
	st = newCacheKVStoreFromParent(parent)
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))

	// make a new one and delete - should not be removed from parent
	st = newCacheKVStoreFromParent(parent)
	st.Delete(keyFmt(1))
	require.Empty(t, st.Get(keyFmt(1)))
	require.Equal(t, parent.Get(keyFmt(1)), valFmt(2))

	// Write. should now be removed from both
	st.Write()
	require.Empty(t, st.Get(keyFmt(1)), "Expected `key1` to be empty")
	require.Empty(t, parent.Get(keyFmt(1)), "Expected `key1` to be empty")
}

func TestCacheKVStoreNoNilSet(t *testing.T) {
	parent := newParent()
	st := newCacheKVStoreFromParent(parent)
	require.Panics(t, func() { st.Set([]byte("key"), nil) }, "setting a nil value should panic")
	require.Panics(t, func() { st.Set(nil, []byte("value")) }, "setting a nil key should panic")
	require.Panics(
		t,
		func() { st.Set([]byte(""), []byte("value")) },
		"setting an empty key should panic",
	)
}

func TestCacheKVStoreNested(t *testing.T) {
	parent := newParent()
	st := newCacheKVStoreFromParent(parent)

	// set. check its there on st and not on parent.
	st.Set(keyFmt(1), valFmt(1))
	require.Empty(t, parent.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))

	// make a new from st and check
	st2 := sdkcachekv.NewStore(st)
	require.Equal(t, valFmt(1), st2.Get(keyFmt(1)))

	// update the value on st2, check it only effects st2
	st2.Set(keyFmt(1), valFmt(3))
	require.Equal(t, []byte(nil), parent.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))
	require.Equal(t, valFmt(3), st2.Get(keyFmt(1)))

	// st2 writes to its parent, st. doesnt effect parent
	st2.Write()
	require.Equal(t, []byte(nil), parent.Get(keyFmt(1)))
	require.Equal(t, valFmt(3), st.Get(keyFmt(1)))

	// updates parent
	st.Write()
	require.Equal(t, valFmt(3), parent.Get(keyFmt(1)))
}

func TestCacheKVIteratorBounds(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	// set some items
	nItems := 5
	for i := 0; i < nItems; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}

	// iterate over all of them
	itr := st.Iterator(nil, nil)
	i := 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, nItems, i)

	// iterate over none
	itr = st.Iterator(bz("money"), nil)
	i = 0
	for ; itr.Valid(); itr.Next() {
		i++
	}
	require.Equal(t, 0, i)

	// iterate over lower
	itr = st.Iterator(keyFmt(0), keyFmt(3))
	i = 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, 3, i)

	// iterate over upper
	itr = st.Iterator(keyFmt(2), keyFmt(4))
	i = 2
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, 4, i)
}

func TestCacheKVMergeIteratorBasics(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	// set and delete an item in the cache, iterator should be empty
	k, v := keyFmt(0), valFmt(0)
	st.Set(k, v)
	st.Delete(k)
	assertIterateDomain(t, st, 0)

	// now set it and assert its there
	st.Set(k, v)
	assertIterateDomain(t, st, 1)

	// write it and assert its there
	st.Write()
	assertIterateDomain(t, st, 1)

	// remove it in cache and assert its not
	st.Delete(k)
	assertIterateDomain(t, st, 0)

	// write the delete and assert its not there
	st.Write()
	assertIterateDomain(t, st, 0)

	// add two keys and assert theyre there
	k1, v1 := keyFmt(1), valFmt(1)
	st.Set(k, v)
	st.Set(k1, v1)
	assertIterateDomain(t, st, 2)

	// write it and assert theyre there
	st.Write()
	assertIterateDomain(t, st, 2)

	// remove one in cache and assert its not
	st.Delete(k1)
	assertIterateDomain(t, st, 1)

	// write the delete and assert its not there
	st.Write()
	assertIterateDomain(t, st, 1)

	// delete the other key in cache and asserts its empty
	st.Delete(k)
	assertIterateDomain(t, st, 0)
}

func TestCacheKVMergeIteratorDeleteLast(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	// set some items and write them
	nItems := 5
	for i := 0; i < nItems; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}
	st.Write()

	// set some more items and leave dirty
	for i := nItems; i < nItems*2; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}

	// iterate over all of them
	assertIterateDomain(t, st, nItems*2)

	// delete them all
	for i := 0; i < nItems*2; i++ {
		last := nItems*2 - 1 - i
		st.Delete(keyFmt(last))
		assertIterateDomain(t, st, last)
	}
}

func TestCacheKVMergeIteratorDeletes(t *testing.T) {
	parent := newParent()
	st := newCacheKVStoreFromParent(parent)

	// set some items and write them
	nItems := 10
	for i := 0; i < nItems; i++ {
		doOp(t, st, parent, opSet, i)
	}
	st.Write()

	// delete every other item, starting from 0
	for i := 0; i < nItems; i += 2 {
		doOp(t, st, parent, opDel, i)
		assertIterateDomainCompare(t, st, parent)
	}

	// reset
	parent = newParent()
	st = newCacheKVStoreFromParent(parent)

	// set some items and write them
	for i := 0; i < nItems; i++ {
		doOp(t, st, parent, opSet, i)
	}
	st.Write()

	// delete every other item, starting from 1
	for i := 1; i < nItems; i += 2 {
		doOp(t, st, parent, opDel, i)
		assertIterateDomainCompare(t, st, parent)
	}
}

func TestCacheKVMergeIteratorChunks(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	// Use the parent to check values on the merge iterator
	// sets to the parent
	setRange(t, st, parent, 0, 20)
	setRange(t, st, parent, 40, 60)
	st.Write()

	// sets to the cache
	setRange(t, st, parent, 20, 40)
	setRange(t, st, parent, 60, 80)
	assertIterateDomainCheck(t, st, parent, []keyRange{{0, 80}})

	// remove some parents and some cache
	deleteRange(t, st, parent, 15, 25)
	assertIterateDomainCheck(t, st, parent, []keyRange{{0, 15}, {25, 80}})

	// remove some parents and some cache
	deleteRange(t, st, parent, 35, 45)
	assertIterateDomainCheck(t, st, parent, []keyRange{{0, 15}, {25, 35}, {45, 80}})

	// write, add more to the cache, and delete some cache
	st.Write()
	setRange(t, st, parent, 38, 42)
	deleteRange(t, st, parent, 40, 43)
	assertIterateDomainCheck(t, st, parent, []keyRange{{0, 15}, {25, 35}, {38, 40}, {45, 80}})
}

func TestCacheKVMergeIteratorRandom(t *testing.T) {
	parent := newParent()
	st := NewStore(parent)

	start, end := 25, 975
	max := 1000
	setRange(t, st, parent, start, end)

	// do an op, test the iterator
	for i := 0; i < 2000; i++ {
		doRandomOp(t, st, parent, max)
		assertIterateDomainCompare(t, st, parent)
	}
}

// ============================================================-
// do some random ops

const (
	opSet      = 0
	opSetRange = 1
	opDel      = 2
	opDelRange = 3
	opWrite    = 4

	totalOps = 5 // number of possible operations
)

func randInt(n int) int {
	return tmrand.NewRand().Int() % n
}

// useful for replaying a error case if we find one.
func doOp(t *testing.T, st types.CacheKVStore, truth types.KVStore, op int, args ...int) {
	switch op {
	case opSet:
		k := args[0]
		st.Set(keyFmt(k), valFmt(k))
		truth.Set(keyFmt(k), valFmt(k))
	case opSetRange:
		start := args[0]
		end := args[1]
		setRange(t, st, truth, start, end)
	case opDel:
		k := args[0]
		st.Delete(keyFmt(k))
		truth.Delete(keyFmt(k))
	case opDelRange:
		start := args[0]
		end := args[1]
		deleteRange(t, st, truth, start, end)
	case opWrite:
		st.Write()
	}
}

func doRandomOp(t *testing.T, st types.CacheKVStore, truth types.KVStore, maxKey int) {
	r := randInt(totalOps)
	switch r {
	case opSet:
		k := randInt(maxKey)
		st.Set(keyFmt(k), valFmt(k))
		truth.Set(keyFmt(k), valFmt(k))
	case opSetRange:
		start := randInt(maxKey - 2)
		end := randInt(maxKey-start) + start
		setRange(t, st, truth, start, end)
	case opDel:
		k := randInt(maxKey)
		st.Delete(keyFmt(k))
		truth.Delete(keyFmt(k))
	case opDelRange:
		start := randInt(maxKey - 2)
		end := randInt(maxKey-start) + start
		deleteRange(t, st, truth, start, end)
	case opWrite:
		st.Write()
	}
}

// ============================================================-

// iterate over whole domain.
func assertIterateDomain(t *testing.T, st types.KVStore, expectedN int) {
	itr := st.Iterator(nil, nil)
	i := 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, expectedN, i)
}

func assertIterateDomainCheck(t *testing.T, st types.KVStore, mem types.KVStore, r []keyRange) {
	// iterate over each and check they match the other
	itr := st.Iterator(nil, nil)
	itr2 := mem.Iterator(nil, nil) // ground truth

	krc := newKeyRangeCounter(r)
	i := 0

	for ; krc.valid(); krc.next() {
		require.True(t, itr.Valid())
		require.True(t, itr2.Valid())

		// check the key/val matches the ground truth
		k, v := itr.Key(), itr.Value()
		k2, v2 := itr2.Key(), itr2.Value()
		require.Equal(t, k, k2)
		require.Equal(t, v, v2)

		// check they match the counter
		require.Equal(t, k, keyFmt(krc.key()))

		itr.Next()
		itr2.Next()
		i++
	}

	require.False(t, itr.Valid())
	require.False(t, itr2.Valid())
}

func assertIterateDomainCompare(t *testing.T, st types.KVStore, mem types.KVStore) {
	// iterate over each and check they match the other
	itr := st.Iterator(nil, nil)
	itr2 := mem.Iterator(nil, nil) // ground truth
	checkIterators(t, itr, itr2)
	checkIterators(t, itr2, itr)
}

func checkIterators(t *testing.T, itr, itr2 types.Iterator) {
	for ; itr.Valid(); itr.Next() {
		require.True(t, itr2.Valid())
		k, v := itr.Key(), itr.Value()
		k2, v2 := itr2.Key(), itr2.Value()
		require.Equal(t, k, k2)
		require.Equal(t, v, v2)
		itr2.Next()
	}
	require.False(t, itr.Valid())
	require.False(t, itr2.Valid())
}

// ====================================--

func setRange(_ *testing.T, st types.KVStore, mem types.KVStore, start, end int) {
	for i := start; i < end; i++ {
		st.Set(keyFmt(i), valFmt(i))
		mem.Set(keyFmt(i), valFmt(i))
	}
}

func deleteRange(_ *testing.T, st types.KVStore, mem types.KVStore, start, end int) {
	for i := start; i < end; i++ {
		st.Delete(keyFmt(i))
		mem.Delete(keyFmt(i))
	}
}

// ====================================--

type keyRange struct {
	start int
	end   int
}

func (kr keyRange) len() int {
	return kr.end - kr.start
}

func newKeyRangeCounter(kr []keyRange) *keyRangeCounter {
	return &keyRangeCounter{keyRanges: kr}
}

// we can iterate over this and make sure our real iterators have all the right keys.
type keyRangeCounter struct {
	rangeIdx  int
	idx       int
	keyRanges []keyRange
}

func (krc *keyRangeCounter) valid() bool {
	maxRangeIdx := len(krc.keyRanges) - 1
	maxRange := krc.keyRanges[maxRangeIdx]

	// if we're not in the max range, we're valid
	if krc.rangeIdx <= maxRangeIdx &&
		krc.idx < maxRange.len() {
		return true
	}

	return false
}

func (krc *keyRangeCounter) next() {
	thisKeyRange := krc.keyRanges[krc.rangeIdx]
	if krc.idx == thisKeyRange.len()-1 {
		krc.rangeIdx++
		krc.idx = 0
	} else {
		krc.idx++
	}
}

func (krc *keyRangeCounter) key() int {
	thisKeyRange := krc.keyRanges[krc.rangeIdx]
	return thisKeyRange.start + krc.idx
}
