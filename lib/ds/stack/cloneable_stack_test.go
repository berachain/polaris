// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	typesmock "github.com/berachain/stargazer/lib/types/mock"
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
