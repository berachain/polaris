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

package precompile_test

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/lib/common"
)

var _ = Describe("Precompile Types", func() {
	Context("Basic - RequireValid Tests", func() {
		It("should error on missing Abi function signature", func() {
			pfgMissingSig := &precompile.FnAndGas{
				Func:        mockPrecompileFn,
				RequiredGas: 10,
			}
			err := pfgMissingSig.RequireValid()
			Expect(err).ToNot(BeNil())
		})

		It("should error on missing (or 0) RequireGas", func() {
			pfgMissingGas := &precompile.FnAndGas{
				AbiSig: "contractFunc(address)",
				Func:   mockPrecompileFn,
			}
			err := pfgMissingGas.RequireValid()
			Expect(err).ToNot(BeNil())
		})

		It("should error on missing precompile function", func() {
			pfgMissingFunc := &precompile.FnAndGas{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
			}
			err := pfgMissingFunc.RequireValid()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Abi Signature verification - RequireValid tests", func() {
		var pfg = &precompile.FnAndGas{
			Func:        mockPrecompileFn,
			RequiredGas: 10,
		}

		It("should not error on valid abi signatures", func() {
			pfg.AbiSig = "contractFunc(address)"
			err := pfg.RequireValid()
			Expect(err).To(BeNil())
			pfg.AbiSig = "getOutputPartial()"
			err = pfg.RequireValid()
			Expect(err).To(BeNil())
			pfg.AbiSig = "cancelUnbondingDelegation(address,uint256,int64)"
			err = pfg.RequireValid()
			Expect(err).To(BeNil())
			pfg.AbiSig = "$$_$3fads343(address,int64,int)"
			err = pfg.RequireValid()
			Expect(err).To(BeNil())
		})

		It("should error on invalid abi signatures", func() {
			pfg.AbiSig = ""
			err := pfg.RequireValid()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "()"
			err = pfg.RequireValid()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "(int64)"
			err = pfg.RequireValid()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "(address,uint256,int64)"
			err = pfg.RequireValid()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "4fsd$_$2f(address)"
			err = pfg.RequireValid()
			Expect(err).ToNot(BeNil())
			pfg.AbiSig = "func(324fds)"
			err = pfg.RequireValid()
			Expect(err).ToNot(BeNil())
		})
	})
})

func mockPrecompileFn(
	ctx sdk.Context,
	evm *precompile.GethEVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (a []any, err error) {
	return
}
