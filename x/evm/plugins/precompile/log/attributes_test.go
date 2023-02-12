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
	"math/big"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/berachain/stargazer/eth/common"
)

var _ = Describe("Attributes", func() {
	var gethValue any
	var err error

	Describe("Test Default Attribute Value Decoder Functions", func() {
		It("should correctly convert sdk coin strings to big.Int", func() {
			denom10 := sdk.NewCoin("denom", sdk.NewInt(10))
			gethValue, err = ConvertSdkCoin(denom10.String())
			Expect(err).To(BeNil())
			bigVal, ok := gethValue.(*big.Int)
			Expect(ok).To(BeTrue())
			Expect(bigVal).To(Equal(big.NewInt(10)))
		})

		It("should correctly convert creation height to int64", func() {
			creationHeightStr := strconv.FormatInt(55, 10)
			gethValue, err = ConvertInt64(creationHeightStr)
			Expect(err).To(BeNil())
			int64Val, ok := gethValue.(int64)
			Expect(ok).To(BeTrue())
			Expect(int64Val).To(Equal(int64(55)))
		})

		It("should correctly convert ValAddress to common.Address", func() {
			valAddr := sdk.ValAddress([]byte("alice"))
			gethValue, err = ConvertValAddressFromBech32(valAddr.String())
			Expect(err).To(BeNil())
			valAddrVal, ok := gethValue.(common.Address)
			Expect(ok).To(BeTrue())
			Expect(valAddrVal).To(Equal(common.ValAddressToEthAddress(valAddr)))
		})

		It("should correctly convert AccAddress to common.Address", func() {
			accAddr := sdk.AccAddress([]byte("alice"))
			gethValue, err = ConvertAccAddressFromBech32(accAddr.String())
			Expect(err).To(BeNil())
			accAddrVal, ok := gethValue.(common.Address)
			Expect(ok).To(BeTrue())
			Expect(accAddrVal).To(Equal(common.AccAddressToEthAddress(accAddr)))
		})
	})

	Describe("Test Search Attributes for Argument", func() {
		var attributes = []abci.EventAttribute{
			{Key: "k0"},
			{Key: "k1"},
			{Key: "k2"},
			{Key: "k3"},
			{Key: "k4"},
		}

		It("should return the correct index if it contains the argument name", func() {
			Expect(searchAttributesForArg(&attributes, "k0")).To(Equal(0))
			Expect(searchAttributesForArg(&attributes, "k3")).To(Equal(3))
			Expect(searchAttributesForArg(&attributes, "k4")).To(Equal(4))
		})

		It("should return -1 if it does not contain the argument name", func() {
			Expect(searchAttributesForArg(&attributes, "")).To(Equal(-1))
			Expect(searchAttributesForArg(&attributes, "k6")).To(Equal(-1))
		})
	})
})
