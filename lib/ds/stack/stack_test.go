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
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
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
	})
})
