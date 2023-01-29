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

package snapshot_test

import (
	"testing"

	"github.com/berachain/stargazer/lib/snapshot"
	typesmock "github.com/berachain/stargazer/lib/types/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnapshot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/snapshot")
}

var _ = Describe("Controller", func() {
	var ctrl *snapshot.Controller
	var object1 *typesmock.SnapshottableMock
	var object2 *typesmock.SnapshottableMock
	// var object3 *typesmock.SnapshottableMock
	BeforeEach(func() {
		ctrl = snapshot.NewController()
		object1 = typesmock.NewSnapshottableMock()
		object2 = typesmock.NewSnapshottableMock()
		// object3 = typesmock.NewSnapshottableMock()
		// ctrl.Control("object1", object1)
		// ctrl.Control("object2", object2)
		// ctrl.Control("object3", object3)
	})

	When("adding a new object", func() {
		BeforeEach(func() {
			err := ctrl.Control("object1", object1)
			Expect(err).To(BeNil())
		})
		It("should add the object", func() {
			obj := ctrl.Get("object1")
			Expect(obj).To(Equal(object1))
		})
		When("adding a new object with the same name", func() {
			It("should return an error", func() {
				err := ctrl.Control("object1", object1)
				Expect(err).To(MatchError(snapshot.ErrObjectAlreadyExists))
			})
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
				snaps := ctrl.Revision(1)
				Expect(snaps).To(HaveLen(1))
				Expect(snaps["object1"]).To(Equal(5))
				snaps = ctrl.LatestRevision()
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
					snaps := ctrl.Revision(1)
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(5))
					snaps = ctrl.Revision(2)
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(12))
					snaps = ctrl.LatestRevision()
					Expect(snaps).To(HaveLen(1))
					Expect(snaps["object1"]).To(Equal(12))
				})
				When("we start controlling a new object", func() {
					BeforeEach(func() {
						Expect(ctrl.Control("object2", object2)).To(BeNil())
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
							snaps := ctrl.Revision(1)
							Expect(snaps).To(HaveLen(1))
							Expect(snaps["object1"]).To(Equal(5))
							snaps = ctrl.Revision(2)
							Expect(snaps).To(HaveLen(1))
							Expect(snaps["object1"]).To(Equal(12))
							snaps = ctrl.Revision(3)
							Expect(snaps).To(HaveLen(2))
							Expect(snaps["object1"]).To(Equal(12))
							Expect(snaps["object2"]).To(Equal(7))
							snaps = ctrl.LatestRevision()
							Expect(snaps).To(HaveLen(2))
							Expect(snaps["object1"]).To(Equal(12))
							Expect(snaps["object2"]).To(Equal(7))
						})
						When("we call revert on the controller", func() {
							It("should have the correct historical revisions", func() {
								ctrl.RevertToSnapshot(2)
								Expect(object1.RevertToSnapshotCalls()).To(HaveLen(1))
								Expect(object2.RevertToSnapshotCalls()).To(HaveLen(1))
								snaps := ctrl.Revision(1)
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(5))
								snaps = ctrl.Revision(2)
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(12))
								Expect(func() {
									ctrl.Revision(3)
								}).To(Panic())
								snaps = ctrl.LatestRevision()
								Expect(snaps).To(HaveLen(1))
								Expect(snaps["object1"]).To(Equal(12))
							})
						})
					})
				})
				It("should not panic on calling finalize", func() {
					Expect(func() {
						ctrl.Finalize()
					}).ToNot(Panic())
				})
			})
		})
	})
})
