// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

package events_test

// import (
// 	"strconv"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"

// 	"github.com/berachain/berachain-node/testing/utils"
// 	"github.com/berachain/stargazer/common"
// 	stakingprecompile "github.com/berachain/stargazer/core/vm/precompile/contracts/staking" //nolint:lll
// 	"github.com/berachain/stargazer/core/vm/precompile/events"
// 	"github.com/berachain/stargazer/crypto"
// 	"github.com/berachain/stargazer/testutil"
// 	"github.com/berachain/stargazer/types/abi"
// )

// func TestEvents(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Events Registry Suite")
// }

// var _ = Describe("Events Registry", func() {
// 	var registry *events.Registry
// 	var stakingModuleAddr common.Address
// 	var valAddr sdk.ValAddress
// 	var delAddr sdk.AccAddress
// 	var amt sdk.Coin
// 	var creationHeight int64

// 	BeforeEach(func() {
// 		_, _, _, sk := testutil.SetupMinimalKeepers()
// 		stakingModuleAddr = common.BytesToAddress(authtypes.NewModuleAddress("staking").Bytes())
// 		registry = events.NewRegistry()
// 		registry.RegisterModule(&stakingModuleAddr, stakingprecompile.NewContract(&sk))
// 		valAddr = getValidatorAddress()
// 		delAddr = utils.CreateRandomAccounts(1)[0]
// 		amt = sdk.NewCoin("denom", sdk.NewInt(1))
// 		creationHeight = int64(1234)
// 	})

// 	Describe("Valid Cosmos Event", func() {
// 		It("should handle all allowed cosmos types", func() {
// 			event := sdk.NewEvent(
// 				"cancel_unbonding_delegation",
// 				sdk.NewAttribute("validator", valAddr.String()),
// 				sdk.NewAttribute("amount", amt.String()),
// 				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 				sdk.NewAttribute("delegator", delAddr.String()),
// 			)
// 			log, err := registry.BuildEthLog(&event)
// 			eventID := crypto.Keccak256Hash(
// 				[]byte("CancelUnbondingDelegation(address,address,uint256,int64)"),
// 			)

// 			Expect(err).To(BeNil())
// 			Expect(log.Address).To(Equal(stakingModuleAddr))
// 			Expect(len(log.Topics)).To(Equal(3))
// 			Expect(log.Topics[0]).To(Equal(eventID))
// 			Expect(log.Topics[1]).To(Equal(
// 				common.BytesToHash(valAddr.Bytes()),
// 			))
// 			Expect(log.Topics[2]).To(Equal(
// 				common.BytesToHash(delAddr.Bytes()),
// 			))
// 			var stakingEventsABI abi.ABI
// 			err = stakingEventsABI.UnmarshalJSON([]byte(stakingprecompile.EventsMetaData.ABI))
// 			Expect(err).To(BeNil())
// 			ethEvent, err := stakingEventsABI.EventByID(eventID)
// 			Expect(err).To(BeNil())
// 			packedData, err := ethEvent.Inputs.NonIndexed().PackValues(
// 				[]any{
// 					amt.Amount.BigInt(),
// 					creationHeight,
// 				},
// 			)
// 			Expect(err).To(BeNil())
// 			Expect(log.Data).To(Equal(packedData))
// 		})
// 	})

// 	Describe("Invalid Cosmos Events", func() {
// 		It("should fail on non-registered event name", func() {
// 			event := sdk.NewEvent("cancel-unbonding-delegation")
// 			_, err := registry.BuildEthLog(&event)
// 			Expect(err.Error()).To(Equal("the Eth event corresponding to
// Cosmos event cancel-unbonding-delegation has not been registered")) //nolint:lll
// 		})

// 		It("should fail on incorrect number of attributes given", func() {
// 			event := sdk.NewEvent(
// 				"cancel_unbonding_delegation",
// 				sdk.NewAttribute("validator", valAddr.String()),
// 				sdk.NewAttribute("amount", amt.String()),
// 				sdk.NewAttribute("delegator", delAddr.String()),
// 			)
// 			_, err := registry.BuildEthLog(&event)
// 			Expect(err.Error()).To(Equal("not enough event
//  attributes provided for event cancel_unbonding_delegation")) //nolint:lll
// 		})

// 		It("should fail on invalid (indexed) attribute key given", func() {
// 			event := sdk.NewEvent(
// 				"cancel_unbonding_delegation",
// 				sdk.NewAttribute("validator!", valAddr.String()),
// 				sdk.NewAttribute("amount", amt.String()),
// 				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 				sdk.NewAttribute("delegator", delAddr.String()),
// 			)
// 			_, err := registry.BuildEthLog(&event)
// 			Expect(err.Error()).To(Equal("no attribute key
// found for event cancel_unbonding_delegation argument validator")) //nolint:lll
// 		})

// 		It("should fail on invalid (non-indexed) attribute key given", func() {
// 			event := sdk.NewEvent(
// 				"cancel_unbonding_delegation",
// 				sdk.NewAttribute("validator", valAddr.String()),
// 				sdk.NewAttribute("amount!", amt.String()),
// 				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 				sdk.NewAttribute("delegator", delAddr.String()),
// 			)
// 			_, err := registry.BuildEthLog(&event)
// 			Expect(err.Error()).To(Equal("no attribute key fou
// nd for event cancel_unbonding_delegation argument amount")) //nolint:lll
// 		})

// 		Context("bad attribute values", func() {
// 			It("should error on bad validator address", func() {
// 				event := sdk.NewEvent(
// 					"cancel_unbonding_delegation",
// 					sdk.NewAttribute("validator", "bad validator string"),
// 					sdk.NewAttribute("amount", amt.String()),
// 					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 					sdk.NewAttribute("delegator", delAddr.String()),
// 				)
// 				_, err := registry.BuildEthLog(&event)
// 				Expect(err).ToNot(BeNil())
// 			})

// 			It("should error on bad amount value", func() {
// 				event := sdk.NewEvent(
// 					"cancel_unbonding_delegation",
// 					sdk.NewAttribute("validator", valAddr.String()),
// 					sdk.NewAttribute("amount", "bad amount value"),
// 					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 					sdk.NewAttribute("delegator", delAddr.String()),
// 				)
// 				_, err := registry.BuildEthLog(&event)
// 				Expect(err).ToNot(BeNil())
// 			})

// 			It("should error on bad account address", func() {
// 				event := sdk.NewEvent(
// 					"cancel_unbonding_delegation",
// 					sdk.NewAttribute("validator", valAddr.String()),
// 					sdk.NewAttribute("amount", amt.String()),
// 					sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
// 					sdk.NewAttribute("delegator", "bad acc string"),
// 				)
// 				_, err := registry.BuildEthLog(&event)
// 				Expect(err).ToNot(BeNil())
// 			})

// 			It("should error on bad creation height", func() {
// 				event := sdk.NewEvent(
// 					"cancel_unbonding_delegation",
// 					sdk.NewAttribute("validator", valAddr.String()),
// 					sdk.NewAttribute("amount", amt.String()),
// 					sdk.NewAttribute("creation_height", "bad creation height"),
// 					sdk.NewAttribute("delegator", delAddr.String()),
// 				)
// 				_, err := registry.BuildEthLog(&event)
// 				Expect(err).ToNot(BeNil())
// 			})
// 		})
// 	})
// })

// func getValidatorAddress() sdk.ValAddress {
// 	return sdk.ValAddress([]byte("hello I am a validator"))
// }
