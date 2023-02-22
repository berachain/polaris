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

package errors_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	liberrors "pkg.berachain.dev/stargazer/lib/errors"
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
