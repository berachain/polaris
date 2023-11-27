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

package stack_test

import (
	"testing"

	"github.com/berachain/polaris/lib/ds"
	"github.com/berachain/polaris/lib/ds/stack"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/ds/stack")
}

var _ = Describe("Stack", func() {
	var s ds.Stack[int]

	BeforeEach(func() {
		s = stack.New[int](1)
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
