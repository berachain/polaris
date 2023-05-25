package errors_test

import (
	"errors"
	"testing"

	liberrors "pkg.berachain.dev/polaris/lib/errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestErrorsLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/errors")
}

var _ = Describe("Wrap", func() {
	_error := errors.New("myFunkyError123")

	When("we call Wrap", func() {
		It("should match", func() {
			err := liberrors.Wrap(_error, "myErrorMessage")
			Expect(err.Error()).To(Equal("myFunkyError123: myErrorMessage"))
			Expect(errors.Unwrap(err)).To(Equal(_error))
		})
	})

	When("we call Wrapf", func() {
		var err error
		BeforeEach(func() {
			err = liberrors.Wrapf(_error, "myErrorMessage %s", "456")
		})

		It("should match", func() {
			Expect(err.Error()).To(Equal("myFunkyError123: myErrorMessage 456"))
			Expect(errors.Unwrap(err)).To(Equal(_error))
		})

		When("we wrap again", func() {
			It("should match", func() {
				err2 := liberrors.Wrapf(err, "myErrorMessage2 %s", "789")
				Expect(err2.Error()).To(Equal("myFunkyError123: myErrorMessage 456: myErrorMessage2 789"))
				Expect(errors.Unwrap(err2)).To(Equal(err))
				Expect(errors.Unwrap(errors.Unwrap(err2))).To(Equal(_error))
			})
		})
	})
})
