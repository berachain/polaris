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

package snapshot

import (
	"testing"

	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/lib/utils"

	typesmock "github.com/berachain/stargazer/lib/types/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnapshot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/snapshot")
}

var _ = Describe("Controller", func() {
	var ctrl *controller[string, libtypes.Controllable[string]]
	var object1 *typesmock.ControllableMock[string]
	var object2 *typesmock.ControllableMock[string]

	BeforeEach(func() {
		ctrl = utils.MustGetAs[*controller[string, libtypes.Controllable[string]]](
			NewController[string, libtypes.Controllable[string]](),
		)
		object1 = typesmock.NewControllableMock1()
		object2 = typesmock.NewControllableMock2()
	})

	When("adding a new object", func() {
		BeforeEach(func() {
			Expect(ctrl.Register(object1)).To(BeNil())
		})
		It("should add the object", func() {
			obj := ctrl.Get("object1")
			Expect(obj).To(Equal(object1))
		})

		When("calling Get on an uncontrolled object", func() {
			It("should return nil", func() {
				obj := ctrl.Get("object2")
				Expect(obj).To(BeNil())
			})
		})

		When("calling snapshot on the controller", func() {
			BeforeEach(func() {
				object1.SnapshotFunc = func() int { return 5 }
				ctrl.Snapshot()
			})
			It("should call snapshot on the controlled object", func() {
				Expect(object1.SnapshotCalls()).To(HaveLen(1))
				snaps := ctrl.journal.PeekAt(0)
				Expect(snaps).To(HaveLen(1))
				Expect(snaps["object1"]).To(Equal(5))
				snaps = ctrl.journal.Peek()
				Expect(snaps).To(HaveLen(1))
				Expect(snaps["object1"]).To(Equal(5))
			})

			When("calling snapshot on the controller again", func() {
				BeforeEach(func() {
					object1.SnapshotFunc = func() int { return 12 }
					ctrl.Snapshot()
				})
				It("should call snapshot on the controlled object again", func() {
					Expect(object1.SnapshotCalls()).To(HaveLen(2))
					snaps := ctrl.journal.PeekAt(0)
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(5))
					snaps = ctrl.journal.PeekAt(1)
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(12))
					snaps = ctrl.journal.Peek()
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(12))
				})
				When("we start controlling a new object", func() {
					BeforeEach(func() {
						Expect(ctrl.Register(object2)).Error().NotTo(HaveOccurred())
					})
					It("should have the correct number of snapshot calls still", func() {
						Expect(object1.SnapshotCalls()).To(HaveLen(2))
						Expect(object2.SnapshotCalls()).To(HaveLen(0))
					})
					When("we snapshot again", func() {
						BeforeEach(func() {
							object2.SnapshotFunc = func() int { return 7 }
							ctrl.Snapshot()
						})
						It("should have the correct number of snapshot calls", func() {
							Expect(object1.SnapshotCalls()).To(HaveLen(3))
							Expect(object2.SnapshotCalls()).To(HaveLen(1))
						})
						It("should have the correct historical revisions", func() {
							snaps := ctrl.journal.PeekAt(0)
							Expect(snaps).To(HaveLen(1))
							Expect(snaps["object1"]).To(Equal(5))
							snaps = ctrl.journal.PeekAt(1)
							Expect(snaps).To(HaveLen(1))
							Expect(snaps["object1"]).To(Equal(12))
							snaps = ctrl.journal.PeekAt(2)
							Expect(snaps).To(HaveLen(2))
							Expect(snaps["object1"]).To(Equal(12))
							Expect(snaps["object2"]).To(Equal(7))
							snaps = ctrl.journal.Peek()
							Expect(snaps).To(HaveLen(2))
							Expect(snaps["object1"]).To(Equal(12))
							Expect(snaps["object2"]).To(Equal(7))
						})
						It("should correctly finalize", func() {
							ctrl.Finalize()
							Expect(len(object1.FinalizeCalls())).To(Equal(1))
							Expect(len(object2.FinalizeCalls())).To(Equal(1))
						})
						When("we call revert on the controller", func() {
							It("should have the correct historical revisions", func() {
								ctrl.RevertToSnapshot(2)
								Expect(object1.RevertToSnapshotCalls()).To(HaveLen(1))
								Expect(object2.RevertToSnapshotCalls()).To(HaveLen(1))
								snaps := ctrl.journal.PeekAt(0)
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(5))
								snaps = ctrl.journal.PeekAt(1)
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(12))
								Expect(func() {
									ctrl.journal.PeekAt(2)
								}).To(Panic())
								snaps = ctrl.journal.Peek()
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(12))
							})
						})
					})
				})
			})
		})
	})
})
