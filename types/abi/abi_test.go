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

	"github.com/berachain/stargazer/types/abi"
)

func TestABI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ABI Test Suite")
}

var _ = Describe("ABI Test Suite", func() {
	Describe("Test ToMixedCase", func() {
		It("should correctly convert under_score strings to mixedCase", func() {
			Expect(abi.ToMixedCase("creation_height")).To(Equal("creationHeight"))
			Expect(abi.ToMixedCase("creation_height_arg")).To(Equal("creationHeightArg"))
		})
	})

	Describe("Test GetIndexed", func() {
		It("should correctly filter out indexed arguments", func() {
			allArgs := abi.Arguments{
				abi.Argument{Indexed: false},
				abi.Argument{
					Name:    "1",
					Indexed: true,
				},
				abi.Argument{
					Name:    "2",
					Indexed: true,
				},
				abi.Argument{Indexed: false},
				abi.Argument{
					Name:    "3",
					Indexed: true,
				},
				abi.Argument{Indexed: false},
				abi.Argument{Indexed: false},
				abi.Argument{
					Name:    "4",
					Indexed: true,
				},
			}
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
				abi.Argument{
					Name:    "4",
					Indexed: true,
				},
			}
			Expect(abi.GetIndexed(allArgs)).To(Equal(indexedArgs))
		})
	})
})
