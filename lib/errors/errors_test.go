// SPDX-License-Identifier: Apache-2.0
//

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
			Expect(err.Error()).To(Equal("myErrorMessage: myFunkyError123"))
			Expect(errors.Unwrap(err)).To(Equal(_error))
		})
	})

	When("we call Wrapf", func() {
		var err error
		BeforeEach(func() {
			err = liberrors.Wrapf(_error, "myErrorMessage %s", "456")
		})

		It("should match", func() {
			Expect(err.Error()).To(Equal("myErrorMessage 456: myFunkyError123"))
			Expect(errors.Unwrap(err)).To(Equal(_error))
		})

		When("we wrap again", func() {
			It("should match", func() {
				err2 := liberrors.Wrapf(err, "myErrorMessage2 %s", "789")
				Expect(err2.Error()).To(Equal("myErrorMessage2 789: myErrorMessage 456: myFunkyError123"))
				Expect(errors.Unwrap(err2)).To(Equal(err))
				Expect(errors.Unwrap(errors.Unwrap(err2))).To(Equal(_error))
			})
		})
	})
})
