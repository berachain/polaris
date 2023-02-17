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

package storage_test

import (
	"testing"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/x/evm/plugins/state/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "x/evm/plugins/state/storage")
}

var _ = Describe("StorageTest", func() {
	When("storage is empty", func() {
		It("should not return an error", func() {
			slots := storage.Storage{}
			Expect(slots.ValidateBasic()).To(BeNil())
		})
	})
	When("storage is not empty", func() {
		var slots storage.Storage

		BeforeEach(func() {
			slots = storage.Storage{
				storage.NewSlot(common.BytesToHash([]byte{1, 2, 3}), common.BytesToHash([]byte{1, 2, 3})),
			}
		})

		It("should not return an error", func() {
			Expect(slots.ValidateBasic()).To(BeNil())
		})

		When("a storage key is empty", func() {
			BeforeEach(func() {
				slots[0].Key = ""
			})

			It("should return an error", func() {
				Expect(slots.ValidateBasic()).NotTo(BeNil())
			})
		})

		It("should be Cloneable", func() {
			clone := slots.Clone()
			Expect(clone).To(Equal(slots))
			Expect(clone).NotTo(BeIdenticalTo(slots))
		})

		When("a storage key is duplicated", func() {
			BeforeEach(func() {
				slots = append(slots, storage.NewSlot(
					common.BytesToHash([]byte{1, 2, 3}),
					common.BytesToHash([]byte{1, 2, 3}),
				))
			})

			It("should return an error", func() {
				Expect(slots.ValidateBasic()).NotTo(BeNil())
			})
		})

		It("should be printable", func() {
			Expect(slots.String()).To(ContainSubstring("key:" +
				"\"0x0000000000000000000000000000000000000000000000000000000000010203\" value:" +
				"\"0x0000000000000000000000000000000000000000000000000000000000010203\"",
			))
		})
	})
})
