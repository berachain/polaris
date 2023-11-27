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
