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
		ctrl = utils.MustGetAs[*controller[string, libtypes.Controllable[string]]](NewController[string, libtypes.Controllable[string]]())
		object1 = typesmock.NewControllableMock1()
		object2 = typesmock.NewControllableMock2()
	})

	When("adding a new object", func() {
		BeforeEach(func() {
			err := ctrl.Register(object1)
			Expect(err).To(BeNil())
		})
		It("should add the object", func() {
			obj, err := ctrl.Get("object1")
			Expect(err).To(BeNil())
			Expect(obj).To(Equal(object1))
		})
		When("adding a new object with the same name", func() {
			It("should return an error", func() {
				err := ctrl.Register(object1)
				Expect(err).To(MatchError(libtypes.ErrObjectAlreadyExists))
			})
		})

		When("calling Get on an uncontrolled object", func() {
			It("should return nil", func() {
				obj, err := ctrl.Get("object2")
				Expect(err.Error()).To(Equal("item object2 not found"))
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
						Expect(ctrl.Register(object2)).To(BeNil())
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
