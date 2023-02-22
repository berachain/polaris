// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package trees_test

import (
	"encoding/binary"
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/lib/ds"
	"pkg.berachain.dev/stargazer/lib/ds/trees"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/ds/trees")
}

var _ = Describe("GetSetDelete", func() {
	var db ds.BTree

	BeforeEach(func() {
		db = trees.NewBTree()
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
	var db ds.BTree

	BeforeEach(func() {
		db = trees.NewBTree()

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
		Expect(err).To(Equal(trees.ErrKeyEmpty))
		_, err = db.ReverseIterator(nil, []byte{})
		Expect(err).To(Equal(trees.ErrKeyEmpty))
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

func verifyIterator(itr dbm.Iterator, expected []int64, _ string) {
	i := 0
	for itr.Valid() {
		key := itr.Key()
		Expect(expected[i]).To(Equal(bytes2Int64(key)))
		itr.Next()
		i++
	}
	Expect(i).To(Equal(len(expected)))
	Expect(itr.Close()).To(BeNil())
}

func int642Bytes(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func bytes2Int64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
