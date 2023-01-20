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

package types_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

var _ = Describe("Container Types", func() {
	Context("Basic - ValidateBasic Tests", func() {
		It("should error on missing Abi function signature", func() {
			pfgMissingSig := &types.Method{
				Execute:     mockPrecompileFn,
				RequiredGas: 10,
			}
			err := pfgMissingSig.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should error on missing (or 0) RequireGas", func() {
			pfgMissingGas := &types.Method{
				AbiSig:  "contractFunc(address)",
				Execute: mockPrecompileFn,
			}
			err := pfgMissingGas.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should error on missing precompile function", func() {
			pfgMissingFunc := &types.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
			}
			err := pfgMissingFunc.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should error on given abi method", func() {
			pfgMissingFunc := &types.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
				Execute:     mockPrecompileFn,
				AbiMethod:   &abi.Method{},
			}
			err := pfgMissingFunc.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Abi Signature verification - ValidateBasic tests", func() {
		var pfg = &types.Method{
			Execute:     mockPrecompileFn,
			RequiredGas: 10,
		}

		It("should not error on valid abi signatures", func() {
			pfg.AbiSig = "contractFunc(address)"
			err := pfg.ValidateBasic()
			Expect(err).To(BeNil())
			pfg.AbiSig = "getOutputPartial()"
			err = pfg.ValidateBasic()
			Expect(err).To(BeNil())
			pfg.AbiSig = "cancelUnbondingDelegation(address,uint256,int64)"
			err = pfg.ValidateBasic()
			Expect(err).To(BeNil())
			pfg.AbiSig = "$$_$3fads343(address,int64,int)"
			err = pfg.ValidateBasic()
			Expect(err).To(BeNil())
		})

		It("should error on invalid abi signatures", func() {
			pfg.AbiSig = ""
			err := pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "()"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "(int64)"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "(address,uint256,int64)"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "4fsd$_$2f(address)"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "func(324fds)"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "func"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "func())"
			err = pfg.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})
	})
})

func mockPrecompileFn(
	ctx sdk.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	return nil, nil
}
