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

package journal_test

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/store/journal"
	"github.com/berachain/stargazer/store/journal/mock"
)

func TestJournalManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "store/journal")
}

var _ = Describe("Journal", func() {
	var jm journal.ManagerI[*journal.Manager]
	var entries []*mock.CacheEntry

	BeforeEach(func() {
		entries = make([]*mock.CacheEntry, 10)
		jm = journal.NewManager()
		for i := 0; i < 10; i++ {
			entries[i] = mock.NewCacheEntry()
		}
	})

	When("the journal is Pushed to", func() {
		BeforeEach(func() {
			jm.Push(entries[0])
		})

		It("should have a size of 1", func() {
			Expect(jm.Size()).To(Equal(1))
		})

		When("the journal is reverted to size 0", func() {
			BeforeEach(func() {
				jm.PopToSize(0)
			})

			It("should have a size of 0", func() {
				Expect(jm.Size()).To(Equal(0))
			})
		})

		When("the journal is Pushed to 9 more times", func() {
			BeforeEach(func() {
				for i := 1; i <= 9; i++ {
					jm.Push(entries[i])
				}
			})

			It(fmt.Sprintf("should have a size of %d", 10), func() {
				Expect(jm.Size()).To(Equal(10))
			})

			size := rand.Int() % 10
			When(fmt.Sprintf("the journal is reverted to size, %d", size), func() {
				BeforeEach(func() {
					jm.PopToSize(size)
				})

				It(fmt.Sprintf("should have a size of %d", size), func() {
					Expect(jm.Size()).To(Equal(size))
				})
			})

			When("the journal is reverted to size 5", func() {
				BeforeEach(func() {
					jm.PopToSize(5)
				})

				It("should have a size of 5", func() {
					Expect(jm.Size()).To(Equal(5))
				})

				It("should have called revert on last 5 entries", func() {
					for i := len(entries) - 1; i >= 5; i-- {
						Expect(entries[i].RevertCallCount()).To(Equal(1))
					}
				})

				It("should not have called revert on the first 5 entries", func() {
					for i := 4; i >= 0; i-- {
						Expect(entries[i].RevertCallCount()).To(Equal(0))
					}
				})

				When("the journal is cloned", func() {
					var jm2 journal.ManagerI[*journal.Manager]
					BeforeEach(func() {
						jm2 = jm.Clone()
					})

					It("should have a size of 5", func() {
						Expect(jm2.Size()).To(Equal(5))
					})

					It("should be a deep copy", func() {
						for i := 0; i < 5; i++ {
							Expect(jm2.PeekAt(i)).To(Equal(jm.PeekAt(i)))
							Expect(jm2.PeekAt(0)).ToNot(BeIdenticalTo(jm.PeekAt(0)))
						}
					})

					When("the original journal is reverted to size 0", func() {
						BeforeEach(func() {
							jm.PopToSize(0)
						})

						It("the clone should stillhave a size of 5", func() {
							Expect(jm2.Size()).To(Equal(5))
						})
					})
				})
			})
		})
	})
})
