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

package utils_test

import (
	"github.com/berachain/stargazer/lib/utils"
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
