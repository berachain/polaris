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

package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/lib/ds"
	"pkg.berachain.dev/stargazer/lib/ds/stack"
	typesmock "pkg.berachain.dev/stargazer/lib/types/mock"
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
