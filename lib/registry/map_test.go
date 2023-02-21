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

package registry_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/lib/registry"
	"pkg.berachain.dev/stargazer/lib/registry/mock"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

var _ = Describe("Registry", func() {
	var r libtypes.Registry[string, *mock.Registrable]

	BeforeEach(func() {
		r = registry.NewMap[string, *mock.Registrable]()
	})

	When("adding an item", func() {
		BeforeEach(func() {
			// Register an item.
			item := mock.NewMockRegistrable("foo", "bar")
			Expect(r.Register(item)).To(BeNil())
		})

		It("should be a no-op if the item already exists", func() {
			// Register the same item again.
			mr := mock.NewMockRegistrable("foo", "bar2")
			Expect(r.Register(mr)).To(BeNil())
			Expect(len(r.Iterate())).To(Equal(1))
			Expect(r.Get("foo").Data()).To(Equal("bar2"))
		})

		It("should be able to get the item", func() {
			// Get the item.
			item := r.Get("foo")
			Expect(item.RegistryKey()).To(Equal("foo"))
		})

		It("should be able to remove the item", func() {
			// Remove the item.
			r.Remove("foo")

			// Get the item.
			item := r.Get("foo")
			Expect(item).To(BeNil())
		})

		It("should be able to check if the item exists", func() {
			// Check if the item exists.
			exists := r.Has("foo")
			Expect(exists).To(BeTrue())

			// Remove the item.
			r.Remove("foo")

			// Check if the item exists.
			exists = r.Has("foo")
			Expect(exists).To(BeFalse())
		})

		It("should be able to check if an item does not exist", func() {
			// Check an item that does not exist.
			exists := r.Has("bar")
			Expect(exists).To(BeFalse())
		})

		It("should no-op when removing an item that does not exist", func() {
			// Remove an item that does not exist.
			r.Remove("bar")
		})
	})
})
