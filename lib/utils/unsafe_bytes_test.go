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

package utils_test

import (
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UnsafeStrToBytes", func() {
	When("given a valid string", func() {
		It("should return a byte array with the same content", func() {
			input := "valid string"
			expectedOutput := []byte("valid string")

			output := utils.UnsafeStrToBytes(input)
			Expect(output).To(Equal(expectedOutput))
		})
	})
})

var _ = Describe("UnsafeBytesToStr", func() {
	When("given a valid byte array", func() {
		It("should return a string with the same content", func() {
			input := []byte("valid byte array")
			expectedOutput := "valid byte array"

			output := utils.UnsafeBytesToStr(input)
			Expect(output).To(Equal(expectedOutput))
		})
	})
	When("given empty input", func() {
		It("should return empty string", func() {
			input := []byte{}
			expectedOutput := ""
			output := utils.UnsafeBytesToStr(input)
			Expect(output).To(Equal(expectedOutput))
		})
	})
})
