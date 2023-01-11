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

package internal_test

import (
	"testing"

	"github.com/berachain/stargazer/store/cachekv/internal"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cachekv/internal")
}

var _ = Describe("GetSetDelete", func() {
	var db internal.BTree

	BeforeEach(func() {
		db = internal.NewBTree()
	})

	It("should return nil for a nonexistent key", func() {
		value := db.Get([]byte("a"))
		Expect(value).To(BeNil())
	})

	It("should set and get a value", func() {
		db.Set([]byte("a"), []byte{0x01})
		db.Set([]byte("b"), []byte{0x02})

		value := db.Get([]byte("a"))
		Expect(value).To(Equal([]byte{0x01}))

		value = db.Get([]byte("b"))
		Expect(value).To(Equal([]byte{0x02}))
	})

	It("should delete a value", func() {
		db.Set([]byte("a"), []byte{0x01})
		db.Set([]byte("b"), []byte{0x02})

		db.Delete([]byte("a"))

		value := db.Get([]byte("a"))
		Expect(value).To(BeNil())

		db.Delete([]byte("b"))

		value = db.Get([]byte("b"))
		Expect(value).To(BeNil())
	})
})
var _ = Describe("DBIterator", func() {
	var db internal.BTree

	BeforeEach(func() {
		db = internal.NewBTree()

		for i := 0; i < 10; i++ {
			if i != 6 { // but skip 6.
				db.Set(int642Bytes(int64(i)), []byte{})
			}
		}
	})

	It("should be deep copyable", func() {
		db2 := db.Copy()

		itr, err := db.Iterator(nil, nil)
		Expect(err).ToNot(HaveOccurred())
		defer itr.Close()

		for itr.Valid() {
			key := itr.Key()
			value := itr.Value()
			itr.Next()
			value2 := db2.Get(key)
			Expect(value).To(Equal(value2))
		}
	})

	It("should error with blank iterator keys", func() {
		_, err := db.ReverseIterator([]byte{}, nil)
		Expect(err).To(Equal(internal.ErrKeyEmpty))
		_, err = db.ReverseIterator(nil, []byte{})
		Expect(err).To(Equal(internal.ErrKeyEmpty))
	})

	It("should iterate forward", func() {
		itr, err := db.Iterator(nil, nil)
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64{0, 1, 2, 3, 4, 5, 7, 8, 9}, "forward iterator")
	})

	It("should iterate reverse", func() {
		ritr, err := db.ReverseIterator(nil, nil)
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(ritr, []int64{9, 8, 7, 5, 4, 3, 2, 1, 0}, "reverse iterator")
	})

	It("should iterate to 0", func() {
		itr, err := db.Iterator(nil, int642Bytes(0))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64(nil), "forward iterator to 0")
	})

	It("should iterate from 10 (ex)", func() {
		ritr, err := db.ReverseIterator(int642Bytes(10), nil)
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(ritr, []int64(nil), "reverse iterator from 10 (ex)")
	})

	It("should iterate from 0", func() {
		itr, err := db.Iterator(int642Bytes(0), nil)
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64{0, 1, 2, 3, 4, 5, 7, 8, 9}, "forward iterator from 0")
	})

	It("should iterate from 1", func() {
		itr, err := db.Iterator(int642Bytes(1), nil)
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64{1, 2, 3, 4, 5, 7, 8, 9}, "forward iterator from 1")
	})

	It("should iterate reverse from 10 (ex)", func() {
		ritr, err := db.ReverseIterator(nil, int642Bytes(10))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(ritr, []int64{9, 8, 7, 5, 4, 3, 2, 1, 0}, "reverse iterator from 10 (ex)")
	})

	It("should iterate reverse from 9 (ex)", func() {
		ritr, err := db.ReverseIterator(nil, int642Bytes(9))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(ritr, []int64{8, 7, 5, 4, 3, 2, 1, 0}, "reverse iterator from 9 (ex)")
	})

	It("should iterate reverse from 8 (ex)", func() {
		ritr, err := db.ReverseIterator(nil, int642Bytes(8))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(ritr, []int64{7, 5, 4, 3, 2, 1, 0}, "reverse iterator from 8 (ex)")
	})

	It("should iterate forward from 5 to 6", func() {
		itr, err := db.Iterator(int642Bytes(5), int642Bytes(6))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64{5}, "forward iterator from 5 to 6")
	})

	It("should iterate forward from 5 to 7", func() {
		itr, err := db.Iterator(int642Bytes(5), int642Bytes(7))
		Expect(err).ToNot(HaveOccurred())
		verifyIterator(itr, []int64{5}, "forward iterator from 5 to 7")
	})

	It("should reverse iterator from 5 (ex) to 4", func() {
		ritr, err := db.ReverseIterator(int642Bytes(4), int642Bytes(5))
		Expect(err).Should(BeNil())
		verifyIterator(ritr, []int64{4}, "reverse iterator from 5 (ex) to 4")
	})

	It("should reverse iterator from 6 (ex) to 4", func() {
		ritr, err := db.ReverseIterator(int642Bytes(4), int642Bytes(6))
		Expect(err).Should(BeNil())
		verifyIterator(ritr, []int64{5, 4}, "reverse iterator from 6 (ex) to 4")
	})

	It("should return a reverse iterator from 9 (ex)", func() {
		ritr, err := db.ReverseIterator(nil, int642Bytes(9))
		Expect(err).NotTo(HaveOccurred())
		verifyIterator(ritr, []int64{8, 7, 5, 4, 3, 2, 1, 0}, "reverse iterator from 9 (ex)")
	})

})

var _ = Describe("Copy", func() {

})

func verifyIterator(itr types.Iterator, expected []int64, _ string) {
	i := 0
	t := GinkgoT()
	for itr.Valid() {
		key := itr.Key()
		require.Equal(t, expected[i], bytes2Int64(key), "iterator: %d mismatches", i)
		itr.Next()
		i++
	}
	require.Equal(t, i, len(expected), "expected to have fully iterated over all the elements in iter")
	require.NoError(t, itr.Close())
}

func int642Bytes(i int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(i))
}

func bytes2Int64(buf []byte) int64 {
	return int64(sdk.BigEndianToUint64(buf))
}
