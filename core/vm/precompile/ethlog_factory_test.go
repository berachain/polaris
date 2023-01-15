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
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/crypto"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

var _ = Describe("Events Registry", func() {
	var factory *precompile.EthereumLogFactory
	var stakingModuleAddr common.Address
	var valAddr sdk.ValAddress
	var delAddr sdk.AccAddress
	var amt sdk.Coin
	var creationHeight int64

	BeforeEach(func() {
		stakingModuleAddr = common.BytesToAddress(authtypes.NewModuleAddress("staking").Bytes())
		factory = precompile.NewEthereumLogFactory()
		factory.RegisterEvent(stakingModuleAddr, getMockAbiEvent(), nil)
		valAddr = sdk.ValAddress([]byte("alice"))
		delAddr = sdk.AccAddress([]byte("bob"))
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
			log, err := factory.BuildLog(&event)
			eventID := crypto.Keccak256Hash(
				[]byte("CancelUnbondingDelegation(address,address,uint256,int64)"),
			)

			Expect(err).To(BeNil())
			Expect(log.Address).To(Equal(stakingModuleAddr))
			Expect(len(log.Topics)).To(Equal(3))
			Expect(log.Topics[0]).To(Equal(eventID))
			Expect(log.Topics[1]).To(Equal(
				common.BytesToHash(valAddr.Bytes()),
			))
			Expect(log.Topics[2]).To(Equal(
				common.BytesToHash(delAddr.Bytes()),
			))
			ethEvent := getMockAbiEvent()
			Expect(ethEvent).ToNot(BeNil())
			packedData, err := ethEvent.Inputs.NonIndexed().PackValues(
				[]any{
					amt.Amount.BigInt(),
					creationHeight,
				},
			)
			Expect(err).To(BeNil())
			Expect(log.Data).To(Equal(packedData))
		})
	})

	Describe("Invalid Cosmos Events", func() {
		It("should fail on non-registered event name", func() {
			event := sdk.NewEvent("cancel-unbonding-delegation")
			_, err := factory.BuildLog(&event)
			Expect(err.Error()).To(Equal("the Ethereum event corresponding to Cosmos event cancel-unbonding-delegation was not registered")) //nolint:lll
		})

		It("should fail on incorrect number of attributes given", func() {
			event := sdk.NewEvent(
				"cancel_unbonding_delegation",
				sdk.NewAttribute("validator", valAddr.String()),
				sdk.NewAttribute("amount", amt.String()),
				sdk.NewAttribute("delegator", delAddr.String()),
			)
			_, err := factory.BuildLog(&event)
			Expect(err.Error()).To(Equal("not enough event attributes provided for event cancel_unbonding_delegation")) //nolint:lll
		})

		It("should fail on invalid (indexed) attribute key given", func() {
			event := sdk.NewEvent(
				"cancel_unbonding_delegation",
				sdk.NewAttribute("validator!", valAddr.String()),
				sdk.NewAttribute("amount", amt.String()),
				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
				sdk.NewAttribute("delegator", delAddr.String()),
			)
			_, err := factory.BuildLog(&event)
			Expect(err.Error()).To(Equal("no attribute key found for event cancel_unbonding_delegation argument validator")) //nolint:lll
		})

		It("should fail on invalid (non-indexed) attribute key given", func() {
			event := sdk.NewEvent(
				"cancel_unbonding_delegation",
				sdk.NewAttribute("validator", valAddr.String()),
				sdk.NewAttribute("amount!", amt.String()),
				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
				sdk.NewAttribute("delegator", delAddr.String()),
			)
			_, err := factory.BuildLog(&event)
			Expect(err.Error()).To(Equal("no attribute key found for event cancel_unbonding_delegation argument amount")) //nolint:lll
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
				_, err := factory.BuildLog(&event)
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
				_, err := factory.BuildLog(&event)
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
				_, err := factory.BuildLog(&event)
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
				_, err := factory.BuildLog(&event)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})

func getMockAbiEvent() abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	int64Type, _ := abi.NewType("int64", "int64", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			abi.Argument{
				Name:    "validator",
				Type:    addrType,
				Indexed: true,
			},
			abi.Argument{
				Name:    "delegator",
				Type:    addrType,
				Indexed: true,
			},
			abi.Argument{
				Name:    "amount",
				Type:    uint256Type,
				Indexed: false,
			},
			abi.Argument{
				Name:    "creationHeight",
				Type:    int64Type,
				Indexed: false,
			},
		},
	)
}
