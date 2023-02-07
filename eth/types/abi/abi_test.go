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

package abi_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/eth/types/abi"
)

func TestABI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/types/abi")
}

var _ = Describe("ABI Test Suite", func() {
	Describe("Test ToMixedCase", func() {
		It("should correctly convert under_score strings to mixedCase", func() {
			Expect(abi.ToMixedCase("creation_height")).To(Equal("creationHeight"))
			Expect(abi.ToMixedCase("creation_height_arg")).To(Equal("creationHeightArg"))
		})
	})

	Describe("Test GetIndexed", func() {
		var allArgs abi.Arguments
		BeforeEach(func() {
			allArgs = abi.Arguments{
				abi.Argument{},
				abi.Argument{
					Name:    "1",
					Indexed: true,
				},
				abi.Argument{
					Name:    "2",
					Indexed: true,
				},
				abi.Argument{},
				abi.Argument{},
				abi.Argument{},
				abi.Argument{
					Name:    "3",
					Indexed: true,
				},
			}
		})

		It("should correctly filter out indexed arguments", func() {
			indexedArgs := abi.Arguments{
				abi.Argument{
					Name:    "1",
					Indexed: true,
				},
				abi.Argument{
					Name:    "2",
					Indexed: true,
				},
				abi.Argument{
					Name:    "3",
					Indexed: true,
				},
			}
			args := abi.GetIndexed(allArgs)
			Expect(args).To(Equal(indexedArgs))
		})

		It("should panic if more than 3 indexed args are given", func() {
			Expect(func() { abi.GetIndexed(append(allArgs, abi.Argument{Indexed: true})) }).To(Panic())
		})
	})

	Describe("Test ToUnderScore", func() {
		It("should correctly convert mixedCase strings to under_score", func() {
			Expect(abi.ToUnderScore("Creation4Height")).To(Equal("creation4_height"))
			Expect(abi.ToUnderScore("creationHeight")).To(Equal("creation_height"))
			Expect(abi.ToUnderScore("creationHeightArg")).To(Equal("creation_height_arg"))
			Expect(abi.ToUnderScore("creation")).To(Equal("creation"))
			Expect(abi.ToUnderScore("creation_height")).To(Equal("creation_height"))
		})
	})
})
