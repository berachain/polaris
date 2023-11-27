// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package registry_test

import (
	"github.com/berachain/polaris/lib/registry"
	"github.com/berachain/polaris/lib/registry/mock"
	libtypes "github.com/berachain/polaris/lib/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
			Expect(r.Register(item)).To(Succeed())
		})

		It("should be a no-op if the item already exists", func() {
			// Register the same item again.
			mr := mock.NewMockRegistrable("foo", "bar2")
			Expect(r.Register(mr)).To(Succeed())
			Expect(r.Iterate()).To(HaveLen(1))
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
