// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package distribution

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtestutil "github.com/cosmos/cosmos-sdk/x/distribution/testutil"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	libgenerated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDistributionPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/distribution")
}

// Test Reporter to use governance module tests with Ginkgo.
type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func setup() (sdk.Context, *distrkeeper.Keeper, *stakingkeeper.Keeper, *bankkeeper.BaseKeeper) {
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()

	encCfg := cosmostestutil.MakeTestEncodingConfig(
		distribution.AppModuleBasic{},
	)

	dk := distrkeeper.NewKeeper(
		encCfg.Codec,
		runtime.NewKVStoreService(storetypes.NewKVStoreKey(distributiontypes.StoreKey)),
		ak,
		bk,
		sk,
		"gov",
		authtypes.NewModuleAddress("gov").String(),
	)

	params := distributiontypes.DefaultParams()
	params.WithdrawAddrEnabled = true
	err := dk.Params.Set(ctx, params)
	Expect(err).ToNot(HaveOccurred())

	err = dk.FeePool.Set(ctx, distributiontypes.InitialFeePool())
	Expect(err).ToNot(HaveOccurred())
	return ctx, &dk, &sk, &bk
}

var _ = Describe("Distribution Precompile Test", func() {
	var (
		contract *Contract
		valAddr  sdk.ValAddress
		f        *log.Factory
		amt      sdk.Coin

		ctx sdk.Context
		dk  *distrkeeper.Keeper
		sk  *stakingkeeper.Keeper
		bk  *bankkeeper.BaseKeeper
	)

	BeforeEach(func() {
		valAddr = sdk.ValAddress([]byte("val"))
		amt = sdk.NewCoin("abera", sdkmath.NewInt(100))

		// Set up the contracts and keepers.
		ctx, dk, sk, bk = setup()
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(
			distrkeeper.NewMsgServerImpl(*dk),
			distrkeeper.NewQuerier(*dk),
		))

		// Register the events.
		f = log.NewFactory([]ethprecompile.Registrable{contract})
	})

	It("should register the withdraw event", func() {
		event := sdk.NewEvent(
			distributiontypes.EventTypeWithdrawRewards,
			sdk.NewAttribute(distributiontypes.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amt.String()),
		)

		log, err := f.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	When("PrecompileMethods", func() {
		It("should return the correct methods", func() {
			Expect(ethprecompile.GeneratePrecompileMethods(contract.ABIMethods(), reflect.ValueOf(contract))).To(HaveLen(5))
		})
	})

	When("SetWithdrawAddress", func() {
		It("should fail if not common address", func() {
			res, err := contract.SetWithdrawAddress(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidHexAddress))
			Expect(res).To(BeNil())
		})

		It("should succeed", func() {
			res, err := contract.SetWithdrawAddress(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
				false,
				testutil.Bob,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("SetWithdrawAddressBech32", func() {
		It("should fail if not string", func() {
			res, err := contract.SetWithdrawAddress(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
				false,
				1,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})

		It("should fail if not bech32 string", func() {
			res, err := contract.SetWithdrawAddress0(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})

		It("should succeed", func() {
			res, err := contract.SetWithdrawAddress0(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
				false,
				cosmlib.AddressToAccAddress(testutil.Bob).String(),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("Withdraw Delegator Rewards", func() {
		var addr sdk.AccAddress
		var tokens sdk.DecCoins

		BeforeEach(func() {
			// Set the previous proposer.
			dk.SetPreviousProposerConsAddr(
				ctx,
				sdk.ConsAddress(testutil.Alice.Bytes()),
			)

			PKS := simtestutil.CreateTestPubKeys(5)
			valConsPk0 := PKS[0]
			valConsAddr0 := sdk.ConsAddress(valConsPk0.Address())
			valAddr = sdk.ValAddress(valConsAddr0)
			addr = sdk.AccAddress(valAddr)
			val, err := distrtestutil.CreateValidator(valConsPk0, sdkmath.NewInt(100))
			Expect(err).ToNot(HaveOccurred())

			// Set the validator.
			sk.SetValidator(ctx, val)

			// Create the delegation.
			sk.SetDelegation(ctx, stakingtypes.Delegation{
				DelegatorAddress: addr.String(),
				ValidatorAddress: valAddr.String(),
				Shares:           val.DelegatorShares,
			})

			// Run the hooks.
			err = dk.Hooks().AfterValidatorCreated(ctx, valAddr)
			Expect(err).ToNot(HaveOccurred())
			err = dk.Hooks().BeforeDelegationCreated(ctx, addr, valAddr)
			Expect(err).ToNot(HaveOccurred())
			err = dk.Hooks().AfterDelegationModified(ctx, addr, valAddr)
			Expect(err).ToNot(HaveOccurred())

			// Next Block.
			ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

			// Allocate some rewards.
			initial := sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)
			tokens = sdk.DecCoins{sdk.NewDecCoin(sdk.DefaultBondDenom, initial)}

			// Allocate the rewards.
			err = dk.AllocateTokensToValidator(ctx, val, tokens)
			Expect(err).ToNot(HaveOccurred())
			// Historical Count should be 2.
			Expect(dk.GetValidatorHistoricalReferenceCount(ctx)).To(Equal(uint64(2)))

			// Fund the distribution module account.
			coins, _ := tokens.TruncateDecimal()
			err = bk.MintCoins(ctx, distributiontypes.ModuleName, coins)
			Expect(err).ToNot(HaveOccurred())
		})

		When("Withdraw Delegator Rewards common address", func() {
			It("should fail if not common address", func() {
				res, err := contract.WithdrawDelegatorReward(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					"0x0000000000",
					cosmlib.ValAddressToEthAddress(valAddr),
				)
				Expect(err).To(MatchError(precompile.ErrInvalidHexAddress))
				Expect(res).To(BeNil())
			})

			It("should fail if validator address not common.address", func() {
				res, err := contract.WithdrawDelegatorReward(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					cosmlib.AccAddressToEthAddress(addr),
					"0x0000000000",
				)
				Expect(err).To(MatchError(precompile.ErrInvalidHexAddress))
				Expect(res).To(BeNil())
			})

			It("Success", func() {
				res, err := contract.WithdrawDelegatorReward(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					cosmlib.AccAddressToEthAddress(addr),
					cosmlib.ValAddressToEthAddress(valAddr),
				)
				Expect(err).ToNot(HaveOccurred())
				resTyped := utils.MustGetAs[[]libgenerated.CosmosCoin](res[0])
				Expect(resTyped[0].Denom).To(Equal(sdk.DefaultBondDenom))
				rewards, _ := tokens.TruncateDecimal()
				Expect(resTyped[0].Amount).To(Equal(rewards[0].Amount.BigInt()))
			})
		})

		When("Withdraw Delegator Rewards bech32 address", func() {
			It("should fail if delegator address not string", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					1,
					valAddr.String(),
				)
				Expect(err).To(MatchError(precompile.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if validator address not string", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					addr.String(),
					1,
				)
				Expect(err).To(MatchError(precompile.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if delegator address not bech32", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					"invalid",
					valAddr.String(),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should fail if validator address not bech32", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					addr.String(),
					"invalid",
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should fail if delegator address not found", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					testutil.Bob.String(),
					valAddr.String(),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("Success", func() {
				res, err := contract.WithdrawDelegatorReward0(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
					false,
					addr.String(),
					valAddr.String(),
				)
				Expect(err).ToNot(HaveOccurred())
				resTyped := utils.MustGetAs[[]libgenerated.CosmosCoin](res[0])
				Expect(resTyped[0].Denom).To(Equal(sdk.DefaultBondDenom))
				rewards, _ := tokens.TruncateDecimal()
				Expect(resTyped[0].Amount).To(Equal(rewards[0].Amount.BigInt()))
			})

		})
		When("Reading Params", func() {
			It("Should get if withdraw forwarding is enabled", func() {
				res, err := contract.GetWithdrawEnabled(ctx, nil, testutil.Alice, big.NewInt(0), true)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal([]any{true}))
			})
		})
		When("Base Precompile Features", func() {
			It("Should have custom value decoders", func() {
				Expect(contract.CustomValueDecoders()).ToNot(BeNil())
			})
			It("Should have correct amount of precompile methods", func() {
				Expect(ethprecompile.GeneratePrecompileMethods(contract.ABIMethods(), reflect.ValueOf(contract))).To(HaveLen(5))
			})
		})
	})
})
