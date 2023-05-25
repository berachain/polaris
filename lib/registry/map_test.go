package registry_test

import (
	"pkg.berachain.dev/polaris/lib/registry"
	"pkg.berachain.dev/polaris/lib/registry/mock"
	libtypes "pkg.berachain.dev/polaris/lib/types"

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
