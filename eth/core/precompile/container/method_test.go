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

package container_test

import (
	"context"
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/eth/core/precompile/container"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

var _ = Describe("Method", func() {
	Context("Basic - ValidateBasic Tests", func() {
		It("should error on missing Abi function signature", func() {
			methodMissingSig := &container.Method{
				Execute:     mockExecutable,
				RequiredGas: 10,
			}
			err := methodMissingSig.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should error on missing precompile executable", func() {
			methodMissingFunc := &container.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
			}
			err := methodMissingFunc.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should error on given abi method", func() {
			methodMissingFunc := &container.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
				Execute:     mockExecutable,
				AbiMethod:   &abi.Method{},
			}
			err := methodMissingFunc.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Abi Signature verification - ValidateBasic tests", func() {
		var method = &container.Method{
			Execute:     mockExecutable,
			RequiredGas: 10,
		}

		It("should not error on valid abi signatures", func() {
			method.AbiSig = "contractFunc(address)"
			err := method.ValidateBasic()
			Expect(err).To(BeNil())
			method.AbiSig = "getOutputPartial()"
			err = method.ValidateBasic()
			Expect(err).To(BeNil())
			method.AbiSig = "cancelUnbondingDelegation(address,uint256,int64)"
			err = method.ValidateBasic()
			Expect(err).To(BeNil())
			method.AbiSig = "$$_$3fads343(address,int64,int)"
			err = method.ValidateBasic()
			Expect(err).To(BeNil())
		})

		It("should error on invalid abi signatures", func() {
			method.AbiSig = ""
			err := method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "()"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "(int64)"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "(address,uint256,int64)"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "4fsd$_$2f(address)"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "func(324fds)"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "func"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
			method.AbiSig = "func())"
			err = method.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})
	})
})

// MOCKS BELOW.

func mockExecutable(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, []*coretypes.Log, error) {
	return nil, nil, nil
}
