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
)

var _ = Describe("Appendable Stack", func() {
	var s ds.Stack[int]

	BeforeEach(func() {
		s = stack.NewA[int]()
	})

	When("pushing an element", func() {
		BeforeEach(func() {
			s.Push(1)
		})
		It("should not be empty", func() {
			Expect(s.Size()).To(Equal(1))
		})
		It("should return the correct element", func() {
			Expect(s.Peek()).To(Equal(1))
		})
		It("should return the correct element", func() {
			Expect(s.PeekAt(0)).To(Equal(1))
		})
		It("should return the correct element", func() {
			Expect(s.Pop()).To(Equal(1))
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
				s.Push(2)
				s.Push(3)
			})

			It("should return the correct element", func() {
				Expect(s.Peek()).To(Equal(3))
				Expect(s.PeekAt(2)).To(Equal(3))
				Expect(s.PeekAt(1)).To(Equal(2))
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
					Expect(s.Peek()).To(Equal(1))
					Expect(s.PeekAt(0)).To(Equal(1))
				})
			})
		})
		When("PopToSize is called with a size larger than the current size", func() {
			It("should panic", func() {
				Expect(func() {
					s.PopToSize(2)
				}).To(Panic())
			})
		})
		When("pop to size zero is called", func() {
			BeforeEach(func() {
				s.PopToSize(0)
			})

			It("should be empty", func() {
				Expect(s.Size()).To(BeZero())
			})
		})
		When("calling pop on an empty stack", func() {
			It("should return an empty stack", func() {
				s.Pop()
				Expect(s.Size()).To(BeZero())
			})

			It("should return a nil element", func() {
				Expect(s.Pop()).To(Equal(1))
				Expect(s.Pop()).To(BeZero())
			})
		})
		When("calling peek on an empty stack", func() {
			It("should return an empty stack", func() {
				s.Pop()
				Expect(s.Peek()).To(BeZero())
				Expect(s.Size()).To(BeZero())
			})

			It("should return a nil element", func() {
				s.PopToSize(0)
				Expect(s.Peek()).To(BeZero())
			})
		})
		When("calling peekat with an index too large", func() {
			It("should panic", func() {
				Expect(func() {
					s.PeekAt(10)
				}).To(Panic())
			})
		})
	})
})
