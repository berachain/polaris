// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package errors_test

import (
	"errors"
	"testing"

	liberrors "github.com/berachain/polaris/lib/errors"

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
