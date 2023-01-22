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
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/crypto"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "core/vm/precompile/log")
}

var _ = Describe("Precompile Log", func() {
	var precompileLog *PrecompileLog
	var precompileAddr = common.BytesToAddress([]byte("my precompile address"))
	var valAddr = sdk.ValAddress([]byte("alice"))
	var delAddr = sdk.AccAddress([]byte("bob"))
	var amt sdk.Coin
	var creationHeight int64

	Context("No value decoder issues", func() {
		BeforeEach(func() {
			var err error
			precompileLog, err = NewPrecompileLog(precompileAddr, mockDefaultAbiEvent(), nil)
			Expect(err).To(BeNil())
			amt = sdk.NewCoin("denom", sdk.NewInt(1))
			creationHeight = int64(1234)
		})

		Describe("Valid Cosmos Event", func() {
			It("should handle all allowed cosmos types", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator", valAddr.String()),
					sdk.NewAttribute("amount", amt.String()),
					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
					sdk.NewAttribute("delegator", delAddr.String()),
				)
				eventID := crypto.Keccak256Hash(
					[]byte("CancelUnbondingDelegation(address,address,uint256,int64)"),
				)

				err := precompileLog.ValidateAttributes(&event)
				Expect(err).To(BeNil())

				addr := precompileLog.precompileAddr
				Expect(addr).To(Equal(addr))

				topics, err := precompileLog.MakeTopics(&event)
				Expect(err).To(BeNil())
				Expect(len(topics)).To(Equal(3))
				Expect(topics[0]).To(Equal(eventID))
				Expect(topics[1]).To(Equal(
					common.BytesToHash(valAddr.Bytes()),
				))
				Expect(topics[2]).To(Equal(
					common.BytesToHash(delAddr.Bytes()),
				))

				data, err := precompileLog.MakeData(&event)
				Expect(err).To(BeNil())
				packedData, err := mockDefaultAbiEvent().Inputs.NonIndexed().PackValues(
					[]any{
						amt.Amount.BigInt(),
						creationHeight,
					},
				)
				Expect(err).To(BeNil())
				Expect(data).To(Equal(packedData))
			})
		})

		Describe("Invalid Cosmos Events", func() {
			It("should fail on incorrect number of attributes given", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator", valAddr.String()),
					sdk.NewAttribute("amount", amt.String()),
					sdk.NewAttribute("delegator", delAddr.String()),
				)
				err := precompileLog.ValidateAttributes(&event)
				Expect(err.Error()).To(Equal("not enough event attributes provided"))
			})

			It("should fail on invalid (indexed) attribute key given", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator!", valAddr.String()),
					sdk.NewAttribute("amount", amt.String()),
					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
					sdk.NewAttribute("delegator", delAddr.String()),
				)
				_, err := precompileLog.MakeTopics(&event)
				Expect(err.Error()).To(Equal("this Ethereum event argument has no matching Cosmos attribute key: validator"))
			})

			It("should fail on invalid (non-indexed) attribute key given", func() {
				event := sdk.NewEvent(
					"cancel_unbonding_delegation",
					sdk.NewAttribute("validator", valAddr.String()),
					sdk.NewAttribute("amount!", amt.String()),
					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
					sdk.NewAttribute("delegator", delAddr.String()),
				)
				_, err := precompileLog.MakeData(&event)
				Expect(err.Error()).To(Equal("this Ethereum event argument has no matching Cosmos attribute key: amount"))
			})

			Context("bad attribute values", func() {
				It("should error on bad validator address", func() {
					event := sdk.NewEvent(
						"cancel_unbonding_delegation",
						sdk.NewAttribute("validator", "bad validator string"),
						sdk.NewAttribute("amount", amt.String()),
						sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
						sdk.NewAttribute("delegator", delAddr.String()),
					)
					_, err := precompileLog.MakeTopics(&event)
					Expect(err).ToNot(BeNil())
				})

				It("should error on bad amount value", func() {
					event := sdk.NewEvent(
						"cancel_unbonding_delegation",
						sdk.NewAttribute("validator", valAddr.String()),
						sdk.NewAttribute("amount", "bad amount value"),
						sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
						sdk.NewAttribute("delegator", delAddr.String()),
					)
					_, err := precompileLog.MakeData(&event)
					Expect(err).ToNot(BeNil())
				})

				It("should error on bad account address", func() {
					event := sdk.NewEvent(
						"cancel_unbonding_delegation",
						sdk.NewAttribute("validator", valAddr.String()),
						sdk.NewAttribute("amount", amt.String()),
						sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
						sdk.NewAttribute("delegator", "bad acc string"),
					)
					_, err := precompileLog.MakeTopics(&event)
					Expect(err).ToNot(BeNil())
				})

				It("should error on bad creation height", func() {
					event := sdk.NewEvent(
						"cancel_unbonding_delegation",
						sdk.NewAttribute("validator", valAddr.String()),
						sdk.NewAttribute("amount", amt.String()),
						sdk.NewAttribute("creation_height", "bad creation height"),
						sdk.NewAttribute("delegator", delAddr.String()),
					)
					_, err := precompileLog.MakeData(&event)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})

	Context("value decoder issues", func() {
		BeforeEach(func() {
			amt = sdk.NewCoin("denom", sdk.NewInt(1))
			creationHeight = int64(1234)
		})

		It("should error on no value decoder func", func() {
			var err error
			precompileLog, err = NewPrecompileLog(precompileAddr, mockBadAbiEvent(), nil)
			Expect(err).To(BeNil())

			event := sdk.NewEvent(
				"cancel_unbonding_delegation",
				sdk.NewAttribute("validator_bad_arg", "bad validator value"),
			)
			_, err = precompileLog.MakeTopics(&event)
			Expect(err.Error()).To(Equal("no value decoder function is found for event attribute key: validator_bad_arg"))
		})

		It("should find the custom value decoders", func() {
			var err error
			precompileLog, err = NewPrecompileLog(
				precompileAddr,
				mockBadAbiEvent(),
				ValueDecoders{
					"validator_bad_arg": func(s string) (any, error) {
						return common.ValAddressToEthAddress(valAddr), nil
					},
				},
			)
			Expect(err).To(BeNil())

			event := sdk.NewEvent(
				"cancel_unbonding_delegation",
				sdk.NewAttribute("validator_bad_arg", "any validator value"),
			)
			topics, err := precompileLog.MakeTopics(&event)
			Expect(err).To(BeNil())
			Expect(topics[1]).To(Equal(common.BytesToHash(valAddr.Bytes())))
		})
	})
})

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

func mockBadAbiEvent() abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			{
				Name:    "validatorBadArg",
				Type:    addrType,
				Indexed: true,
			},
		},
	)
}
