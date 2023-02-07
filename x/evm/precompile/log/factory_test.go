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
	"strconv"

	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	var f *Factory
	var valAddr sdk.ValAddress
	var delAddr sdk.AccAddress
	var amt sdk.Coin
	var creationHeight int64

	BeforeEach(func() {
		f = NewFactory()
		valAddr = sdk.ValAddress([]byte("alice"))
		delAddr = sdk.AccAddress([]byte("bob"))
		creationHeight = int64(10)
		amt = sdk.NewCoin("denom", sdk.NewInt(10))

		Expect(func() {
			f.RegisterEvent(common.BytesToAddress([]byte{0x01}), mockDefaultAbiEvent(), nil)
		}).ToNot(Panic())
		Expect(func() {
			f.RegisterEvent(common.BytesToAddress([]byte{0x02}), mockCustomAbiEvent(), cvd)
		}).ToNot(Panic())

		When("building Eth logs", func() {
			It("should not build for an unregistered event", func() {
				event := sdk.NewEvent("unbonding_delegation")
				log, err := f.Build(&event)
				Expect(err.Error()).To(Equal("no Ethereum event was registered for this Cosmos event"))
				Expect(log).To(BeNil())
			})

			It("should error on invalid attributes", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator", valAddr.String()),
				)
				log, err := f.Build(&event)
				Expect(err.Error()).To(Equal("not enough event attributes provided"))
				Expect(log).To(BeNil())
			})

			It("should correctly build a log for valid event", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator", valAddr.String()),
					sdk.NewAttribute("amount", amt.String()),
					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
					sdk.NewAttribute("delegator", delAddr.String()),
				)
				log, err := f.Build(&event)
				Expect(err).To(BeNil())
				Expect(log).ToNot(BeNil())
				Expect(log.Address).To(Equal(common.BytesToAddress([]byte{0x01})))
				Expect(log.Topics).To(HaveLen(3))
				Expect(log.Topics[0]).To(Equal(
					crypto.Keccak256Hash(
						[]byte("CancelUnbondingDelegation(address,address,uint256,int64)"),
					),
				))
				Expect(log.Topics[1]).To(Equal(common.BytesToHash(valAddr.Bytes())))
				Expect(log.Topics[2]).To(Equal(common.BytesToHash(delAddr.Bytes())))
				packedData, err := mockDefaultAbiEvent().Inputs.NonIndexed().Pack(
					amt.Amount.BigInt(), creationHeight,
				)
				Expect(err).To(BeNil())
				Expect(log.Data).To(Equal(packedData))
			})
		})
	})
})

// MOCKS BELOW.

func mockCustomAbiEvent() abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	return abi.NewEvent(
		"CustomUnbondingDelegation",
		"CustomUnbondingDelegation",
		false,
		abi.Arguments{
			{
				Name:    "customValidator",
				Type:    addrType,
				Indexed: true,
			},
			{
				Name:    "customAmount",
				Type:    uint256Type,
				Indexed: false,
			},
		},
	)
}

var cvd = ValueDecoders{
	"customValidator": func(val string) (interface{}, error) {
		return common.BytesToAddress([]byte(val)), nil
	},
	"customAmount": func(val string) (interface{}, error) {
		return common.BytesToAddress([]byte(val)), nil
	},
}
