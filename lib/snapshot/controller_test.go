package snapshot

import (
	"testing"

	libtypes "pkg.berachain.dev/polaris/lib/types"
	typesmock "pkg.berachain.dev/polaris/lib/types/mock"
	"pkg.berachain.dev/polaris/lib/utils"

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
			Expect(ctrl.Register(object1)).To(Succeed())
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
						Expect(object2.SnapshotCalls()).To(BeEmpty())
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
							Expect(object1.FinalizeCalls()).To(HaveLen(1))
							Expect(object2.FinalizeCalls()).To(HaveLen(1))
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
