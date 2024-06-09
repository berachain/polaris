// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package journal

import (
	"github.com/berachain/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Refund", func() {
	var r *refund

	BeforeEach(func() {
		r = utils.MustGetAs[*refund](NewRefund())
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

		It("should clone correctly", func() {
			r.AddRefund(1)
			Expect(r.GetRefund()).To(Equal(uint64(2)))
			Expect(r.Size()).To(Equal(2))

			r2 := utils.MustGetAs[*refund](r.Clone())
			Expect(r2.GetRefund()).To(Equal(uint64(2)))
			Expect(r2.Size()).To(Equal(2))

			r2.AddRefund(1)
			Expect(r2.GetRefund()).To(Equal(uint64(3)))
			Expect(r2.Size()).To(Equal(3))
			Expect(r.GetRefund()).To(Equal(uint64(2)))
			Expect(r.Size()).To(Equal(2))
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
						Expect(func() { r.Finalize() }).ToNot(Panic())
					})
				})
			})
		})
	})
})
