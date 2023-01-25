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

package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/lib/ds"
)

var _ = Describe("Stack", func() {
	var stack ds.Stack[int]

	BeforeEach(func() {
		stack = ds.NewStack[int](1)
	})

	When("pushing an element", func() {
		BeforeEach(func() {
			stack.Push(1)
		})

		It("should not be empty", func() {
			Expect(stack.Size()).To(Equal(1))
		})

		It("should return the correct element", func() {
			Expect(stack.Peek()).To(Equal(1))
		})

		It("should return the correct element", func() {
			Expect(stack.PeekAt(0)).To(Equal(1))
		})
		It("should return the correct element", func() {
			Expect(stack.Pop()).To(Equal(1))
		})

		When("popping an element", func() {
			BeforeEach(func() {
				stack.Pop()
			})

			It("should be empty", func() {
				Expect(stack.Size()).To(BeZero())
			})
		})

		When("pushing more elements", func() {
			BeforeEach(func() {
				stack.Push(2)
				stack.Push(3)
			})

			It("should return the correct element", func() {
				Expect(stack.Peek()).To(Equal(3))
				Expect(stack.PeekAt(2)).To(Equal(3))
				Expect(stack.PeekAt(1)).To(Equal(2))
			})

			It("should have the correct size", func() {
				Expect(stack.Size()).To(Equal(3))
			})

			When("calling poptosize with a size smaller than the current size", func() {
				BeforeEach(func() {
					stack.PopToSize(1)
				})

				It("should have the correct size", func() {
					Expect(stack.Size()).To(Equal(1))
				})

				It("should return the correct element", func() {
					Expect(stack.Peek()).To(Equal(1))
					Expect(stack.PeekAt(0)).To(Equal(1))
				})
			})
		})
	})
})
