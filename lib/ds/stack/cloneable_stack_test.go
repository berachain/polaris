// SPDX-License-Identifier: Apache-2.0
//

package stack_test

import (
	"pkg.berachain.dev/polaris/lib/ds"
	"pkg.berachain.dev/polaris/lib/ds/stack"
	typesmock "pkg.berachain.dev/polaris/lib/types/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cloneable Stack", func() {
	var s ds.CloneableStack[*typesmock.WrappedCloneableMock]
	item1 := typesmock.NewWrappedCloneableMock[typesmock.WrappedCloneableMock](1)
	item2 := typesmock.NewWrappedCloneableMock[typesmock.WrappedCloneableMock](2)
	item3 := typesmock.NewWrappedCloneableMock[typesmock.WrappedCloneableMock](3)

	BeforeEach(func() {
		s = stack.NewCloneable[*typesmock.WrappedCloneableMock](1000)
	})

	When("pushing an element", func() {
		BeforeEach(func() {
			s.Push(item1)
		})

		It("should not be empty", func() {
			Expect(s.Size()).To(Equal(1))
		})

		It("should return the correct element", func() {
			Expect(s.Peek()).To(Equal(item1))
		})

		It("should return the correct element", func() {
			Expect(s.PeekAt(0)).To(Equal(item1))
		})
		It("should return the correct element", func() {
			Expect(s.Pop()).To(Equal(item1))
		})

		When("popping an element", func() {
			BeforeEach(func() {
				s.Pop()
			})

			It("should be empty", func() {
				Expect(s.Size()).To(BeZero())
			})
		})

		When("pushing more elements", func() {
			BeforeEach(func() {
				s.Push(item2)
				s.Push(item3)
			})

			It("should return the correct element", func() {
				Expect(s.Peek()).To(Equal(item3))
				Expect(s.PeekAt(2)).To(Equal(item3))
				Expect(s.PeekAt(1)).To(Equal(item2))
			})

			It("should have the correct size", func() {
				Expect(s.Size()).To(Equal(3))
			})

			When("calling poptosize with a size smaller than the current size", func() {
				BeforeEach(func() {
					s.PopToSize(1)
				})

				It("should have the correct size", func() {
					Expect(s.Size()).To(Equal(1))
				})

				It("should return the correct element", func() {
					Expect(s.Peek()).To(Equal(item1))
					Expect(s.PeekAt(0)).To(Equal(item1))
				})
			})

			When("we call clone", func() {
				var s2 ds.CloneableStack[*typesmock.WrappedCloneableMock]
				BeforeEach(func() {
					s2 = s.Clone()
				})

				It("should have the same size", func() {
					Expect(s.Size()).To(Equal(s2.Size()))
				})

				It("should have the same elements", func() {
					Expect(s.Peek().Val()).To(Equal(s2.Peek().Val()))
					Expect(s.PeekAt(0).Val()).To(Equal(s2.PeekAt(0).Val()))
					Expect(s.PeekAt(1).Val()).To(Equal(s2.PeekAt(1).Val()))
					Expect(s.PeekAt(2).Val()).To(Equal(s2.PeekAt(2).Val()))
				})

				It("items should have different memory addresses", func() {
					Expect(s.Peek()).NotTo(BeIdenticalTo(s2.Peek()))
					Expect(s.PeekAt(0)).NotTo(BeIdenticalTo(s2.PeekAt(0)))
					Expect(s.PeekAt(1)).NotTo(BeIdenticalTo(s2.PeekAt(1)))
					Expect(s.PeekAt(2)).NotTo(BeIdenticalTo(s2.PeekAt(2)))
				})
			})
		})
	})
})
