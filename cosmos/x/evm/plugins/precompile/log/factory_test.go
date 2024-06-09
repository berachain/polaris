// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package log

import (
	"errors"
	"strconv"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"

	cosmlib "github.com/berachain/polaris/cosmos/lib"
	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/eth/accounts/abi"
	"github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/precompile/mock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var sk stakingkeeper.Keeper

var _ = Describe("Factory", func() {
	var (
		f              *Factory
		valAddr        sdk.ValAddress
		delAddr        sdk.AccAddress
		amt            sdk.Coin
		creationHeight int64
		pc             *mock.StatefulImplMock
	)

	BeforeEach(func() {
		_, _, _, sk = testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))
		valAddr = sdk.ValAddress([]byte("alice"))
		delAddr = sdk.AccAddress([]byte("bob"))
		creationHeight = int64(10)
		amt = sdk.NewCoin("denom", sdkmath.NewInt(10))
		pc = mock.NewStatefulImpl()

		Expect(func() {
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x02})
			}
			pc.ABIEventsFunc = mockCustomAbiEvent
			pc.CustomValueDecodersFunc = func() precompile.ValueDecoders {
				return cvd
			}
			f = NewFactory([]precompile.Registrable{pc})
		}).ToNot(Panic())
		Expect(func() {
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x01})
			}
			pc.ABIEventsFunc = func() map[string]abi.Event {
				return map[string]abi.Event{
					"CancelUnbondingDelegation": mockDefaultAbiEvent(),
				}
			}
			f = NewFactory([]precompile.Registrable{pc})
		}).ToNot(Panic())
	})

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
				sdk.NewAttribute("amount", amt.String()),
				sdk.NewAttribute("creation_height", strconv.FormatInt(creationHeight, 10)),
				sdk.NewAttribute("option", delAddr.String()),
			)
			log, err := f.Build(&event)
			Expect(err).ToNot(HaveOccurred())
			Expect(log).ToNot(BeNil())
			Expect(log.Address).To(Equal(common.BytesToAddress([]byte{0x01})))
			Expect(log.Topics).To(HaveLen(2))
			Expect(log.Topics[0]).To(Equal(
				crypto.Keccak256Hash(
					[]byte("CancelUnbondingDelegation(string,(uint256,string)[],int64)"),
				),
			))
			Expect(log.Topics[1]).To(Equal(crypto.Keccak256Hash([]byte(delAddr.String()))))
			packedData, err := mockDefaultAbiEvent().Inputs.NonIndexed().Pack(
				cosmlib.SdkCoinsToEvmCoins(sdk.NewCoins(amt)), creationHeight,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(log.Data).To(Equal(packedData))
		})

		It("should correctly build a log for valid event with custom decoder", func() {
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x02})
			}
			pc.ABIEventsFunc = mockCustomAbiEvent
			pc.CustomValueDecodersFunc = func() precompile.ValueDecoders {
				return cvd
			}
			f = NewFactory([]precompile.Registrable{pc})

			event := sdk.NewEvent(
				"custom_unbonding_delegation",
				sdk.NewAttribute("custom_validator", valAddr.String()),
				sdk.NewAttribute("custom_amount", amt.String()),
			)
			log, err := f.Build(&event)
			Expect(err).ToNot(HaveOccurred())
			Expect(log).ToNot(BeNil())
			Expect(log.Address).To(Equal(common.BytesToAddress([]byte{0x02})))
			Expect(log.Topics).To(HaveLen(2))
			Expect(log.Topics[0]).To(Equal(
				crypto.Keccak256Hash(
					[]byte("CustomUnbondingDelegation(address,(uint256,string)[])"),
				),
			))
			Expect(log.Topics[1]).To(Equal(common.BytesToHash(valAddr.Bytes())))
			packedData, err := mockCustomAbiEvent()["CustomUnbondingDelegation"].
				Inputs.NonIndexed().Pack(cosmlib.SdkCoinsToEvmCoins(sdk.NewCoins(amt)))
			Expect(err).ToNot(HaveOccurred())
			Expect(log.Data).To(Equal(packedData))
		})
	})

	When("building invalid Cosmos events", func() {
		It("should not find the custom value decoder", func() {
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x02})
			}
			pc.ABIEventsFunc = mockBadAbiEvent
			pc.CustomValueDecodersFunc = func() precompile.ValueDecoders {
				return cvd
			}
			f = NewFactory([]precompile.Registrable{pc})

			event := sdk.NewEvent(
				"custom_unbonding_delegation",
				sdk.NewAttribute("custom_validator", valAddr.String()),
				sdk.NewAttribute("custom_amount", amt.String()),
				sdk.NewAttribute("invalid_arg", amt.String()),
			)
			log, err := f.Build(&event)
			Expect(log).To(BeNil())
			Expect(err.Error()).To(
				Equal("no value decoder function is found for event attribute key: invalid_arg"))
		})

		It("should error on decoders returning errors", func() {
			badCvd := make(precompile.ValueDecoders)
			badCvd["custom_amount"] = func(val string) (any, error) {
				coin, err := sdk.ParseCoinNormalized(val)
				if err != nil {
					return nil, err
				}
				return coin.Amount.BigInt(), nil
			}
			badCvd["custom_validator"] = func(val string) (any, error) {
				return nil, errors.New("invalid validator address")
			}
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x03})
			}
			pc.ABIEventsFunc = mockCustomAbiEvent
			pc.CustomValueDecodersFunc = func() precompile.ValueDecoders {
				return badCvd
			}
			f = NewFactory([]precompile.Registrable{pc})
			event := sdk.NewEvent(
				"custom_unbonding_delegation",
				sdk.NewAttribute("custom_validator", valAddr.String()),
				sdk.NewAttribute("custom_amount", amt.String()),
			)
			log, err := f.Build(&event)
			Expect(log).To(BeNil())
			Expect(err.Error()).To(Equal("invalid validator address"))

			badCvd = make(precompile.ValueDecoders)
			badCvd["custom_validator"] = func(val string) (any, error) {
				return cosmlib.EthAddressFromString(sk.ValidatorAddressCodec(), val)
			}
			badCvd["custom_amount"] = func(val string) (any, error) {
				return nil, errors.New("invalid amount")
			}
			f = NewFactory([]precompile.Registrable{pc})
			log, err = f.Build(&event)
			Expect(log).To(BeNil())
			Expect(err.Error()).To(Equal("invalid amount"))
		})

		It("should not find attribute key", func() {
			pc.RegistryKeyFunc = func() common.Address {
				return common.BytesToAddress([]byte{0x03})
			}
			pc.ABIEventsFunc = mockCustomAbiEvent
			pc.CustomValueDecodersFunc = func() precompile.ValueDecoders {
				return cvd
			}
			f = NewFactory([]precompile.Registrable{pc})
			event := sdk.NewEvent(
				"custom_unbonding_delegation",
				sdk.NewAttribute("custom_validator", valAddr.String()),
				sdk.NewAttribute("custom_amount_bad", amt.String()),
			)
			log, err := f.Build(&event)
			Expect(log).To(BeNil())
			Expect(err.Error()).To(Equal(
				"this Ethereum event argument has no matching Cosmos attribute key: customAmount"))

			event = sdk.NewEvent(
				"custom_unbonding_delegation",
				sdk.NewAttribute("custom_validator_bad", valAddr.String()),
				sdk.NewAttribute("custom_amount", amt.String()),
			)
			log, err = f.Build(&event)
			Expect(log).To(BeNil())
			Expect(err.Error()).To(Equal(
				"this Ethereum event argument has no matching Cosmos attribute key: customValidator"))
		})
	})
})

// MOCKS BELOW.

func mockCustomAbiEvent() map[string]abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	coinType, _ := abi.NewType("tuple[]", "structIStakingModule.Coin[]", []abi.ArgumentMarshaling{
		{
			Name:         "amount",
			Type:         "uint256",
			InternalType: "uint256",
			Components:   nil,
			Indexed:      false,
		},
		{
			Name:         "denom",
			Type:         "string",
			InternalType: "string",
			Components:   nil,
			Indexed:      false,
		},
	})
	return map[string]abi.Event{
		"CustomUnbondingDelegation": abi.NewEvent(
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
					Type:    coinType,
					Indexed: false,
				},
			},
		),
	}
}

var cvd = precompile.ValueDecoders{
	"custom_validator": func(val string) (any, error) {
		return cosmlib.EthAddressFromString(sk.ValidatorAddressCodec(), val)
	},
	"custom_amount": func(val string) (any, error) {
		coin, err := sdk.ParseCoinNormalized(val)
		if err != nil {
			return nil, err
		}
		evmCoins := cosmlib.SdkCoinsToEvmCoins(sdk.Coins{coin})
		return evmCoins, nil
	},
}

func mockBadAbiEvent() map[string]abi.Event {
	addrType, _ := abi.NewType("address", "address", nil)
	coinType, _ := abi.NewType("tuple[]", "structIStakingModule.Coin[]", []abi.ArgumentMarshaling{
		{
			Name:         "amount",
			Type:         "uint256",
			InternalType: "uint256",
			Components:   nil,
			Indexed:      false,
		},
		{
			Name:         "denom",
			Type:         "string",
			InternalType: "string",
			Components:   nil,
			Indexed:      false,
		},
	})
	return map[string]abi.Event{
		"CustomUnbondingDelegation": abi.NewEvent(
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
					Type:    coinType,
					Indexed: false,
				},
				{
					Name:    "invalidArg",
					Type:    coinType,
					Indexed: false,
				},
			},
		),
	}
}
