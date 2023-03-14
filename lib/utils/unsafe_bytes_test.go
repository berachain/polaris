// SPDX-License-Identifier: Apache-2.0
//

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
