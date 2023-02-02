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

package plugin

import (
	"github.com/berachain/stargazer/eth/core/state"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Refund", func() {
	var r state.RefundPlugin

	BeforeEach(func() {
		r = NewRefund()
	})
	It("should have the correct registry key", func() {
		Expect(r.RegistryKey()).To(Equal("refund"))
	})
	When("adding a refund", func() {
		BeforeEach(func() {
			r.AddRefund(1)
		})
		It("should return the correct refund", func() {
			Expect(r.GetRefund()).To(Equal(uint64(1)))
		})
		When("subtracting a refund", func() {
			BeforeEach(func() {
				r.SubRefund(1)
			})
			It("should return the correct refund", func() {
				Expect(r.GetRefund()).To(BeZero())
			})
		})
	})
	When("pushing an element", func() {
		BeforeEach(func() {
			r.AddRefund(1)
		})
		It("should not be empty", func() {
			Expect(r.Snapshot()).To(Equal(1))
		})
		It("should return the correct element", func() {
			Expect(r.GetRefund()).To(Equal(uint64(1)))
		})
		When("subbing refund", func() {
			BeforeEach(func() {
				r.SubRefund(1)
			})

			It("should be empty", func() {
				Expect(r.GetRefund()).To(BeZero())
			})
		})
		When("pushing more elements and snapshotting", func() {
			BeforeEach(func() {
				r.AddRefund(2)
				Expect(r.Snapshot()).To(Equal(2))
				r.AddRefund(3)
				Expect(r.Snapshot()).To(Equal(3))
			})
			It("should return the correct element", func() {
				Expect(r.GetRefund()).To(Equal(uint64(6)))
			})
			When("subbing an element", func() {
				BeforeEach(func() {
					r.SubRefund(3)
				})
				It("should return the correct element", func() {
					Expect(r.GetRefund()).To(Equal(uint64(3)))
				})
				When("subbing an element", func() {
					BeforeEach(func() {
						r.SubRefund(3)
					})
					It("should return the correct element", func() {
						Expect(r.GetRefund()).To(Equal(uint64(0)))
					})
					When("taking a snapshot", func() {
						BeforeEach(func() {
							Expect(r.Snapshot()).To(Equal(5))
						})

						When("adding more elements", func() {
							BeforeEach(func() {
								r.AddRefund(1)
							})
							When("reverting to snapshot", func() {
								BeforeEach(func() {
									r.RevertToSnapshot(1)
								})
								It("should return the correct element", func() {
									Expect(r.GetRefund()).To(Equal(uint64(1)))
								})
							})
						})
					})
				})
				When("finally reverting to snapshot", func() {
					BeforeEach(func() {
						r.RevertToSnapshot(0)
					})

					It("should return the correct element", func() {
						Expect(r.GetRefund()).To(Equal(uint64(0)))
					})
				})
				When("finalize", func() {
					It("should not panic", func() {
						Expect(func() {
							r.Finalize()
						}).ToNot(Panic())
					})
				})
			})
		})
	})
})
