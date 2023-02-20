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
	"crypto/rand"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/x/evm/plugins/state/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("x/evm/plugins/state/storage", func() {
	var slot *storage.Slot
	key := common.Hash{}.Bytes()
	value := common.Hash{}.Bytes()

	BeforeEach(func() {
		rand.Read(key)
		rand.Read(value)
		slot = storage.NewSlot(common.BytesToHash(key), common.BytesToHash(value))
	})

	It("should return the correct key", func() {
		Expect(slot.Key).To(Equal(common.BytesToHash(key).Hex()))
	})

	It("should return the correct value", func() {
		Expect(slot.Value).To(Equal(common.BytesToHash(value).Hex()))
	})

	It("should have valid slot", func() {
		Expect(slot.ValidateBasic()).To(BeNil())
	})

	When("slot key is empty", func() {
		BeforeEach(func() {
			slot.Key = ""
		})

		It("should return an error", func() {
			Expect(slot.ValidateBasic()).NotTo(BeNil())
		})
	})

	When("slot key has leading or trailing spaces", func() {
		When("slot key is not empty", func() {
			BeforeEach(func() {
				slot.Key = " bingbong "
			})

			It("should not return an error", func() {
				Expect(slot.ValidateBasic()).To(BeNil())
			})
		})

		When("slot key is empty", func() {
			BeforeEach(func() {
				slot.Key = "       "
			})

			It("should return an error", func() {
				Expect(slot.ValidateBasic()).NotTo(BeNil())
			})
		})
	})

	It("is cloneable", func() {
		clone := slot.Clone()
		Expect(clone).To(Equal(slot))
		Expect(&clone).NotTo(BeIdenticalTo(&slot))
	})
})
