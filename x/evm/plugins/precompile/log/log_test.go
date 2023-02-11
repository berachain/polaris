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

package log

import (
	"testing"

	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "x/evm/plugins/precompile/log")
}

var _ = Describe("precompileLog", func() {
	It("should properly create a new precompile log", func() {
		var pl *precompileLog
		Expect(func() {
			pl = newPrecompileLog(common.BytesToAddress([]byte{1}), mockDefaultAbiEvent())
		}).ToNot(Panic())
		Expect(pl.RegistryKey()).To(Equal("cancel_unbonding_delegation"))
		Expect(pl.id).To(Equal(crypto.Keccak256Hash(
			[]byte("CancelUnbondingDelegation(address,address,uint256,int64)"),
		)))
		Expect(pl.precompileAddr).To(Equal(common.BytesToAddress([]byte{1})))
		Expect(len(pl.indexedInputs)).To(Equal(2))
		Expect(len(pl.nonIndexedInputs)).To(Equal(2))
	})
})

// MOCKS BELOW.

func mockDefaultAbiEvent() abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	int64Type, _ := abi.NewType("int64", "int64", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			{
				Name:    "validator",
				Type:    addrType,
				Indexed: true,
			},
			{
				Name:    "delegator",
				Type:    addrType,
				Indexed: true,
			},
			{
				Name:    "amount",
				Type:    uint256Type,
				Indexed: false,
			},
			{
				Name:    "creationHeight",
				Type:    int64Type,
				Indexed: false,
			},
		},
	)
}
