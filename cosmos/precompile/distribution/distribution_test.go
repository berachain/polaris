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
	"testing"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	pclog "github.com/berachain/polaris/cosmos/x/evm/plugins/precompile/log"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/vm"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/runtime"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtestutil "github.com/cosmos/cosmos-sdk/x/distribution/testutil"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"

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

func setup() (
	sdk.Context,
	authkeeper.AccountKeeperI,
	*distrkeeper.Keeper,
	*stakingkeeper.Keeper,
	*bankkeeper.BaseKeeper,
) {
	distrKey := storetypes.NewKVStoreKey(distributiontypes.StoreKey)
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()),
		[]storetypes.StoreKey{distrKey}...)

	encCfg := cosmostestutil.MakeTestEncodingConfig(
		distribution.AppModuleBasic{},
	)

	dk := distrkeeper.NewKeeper(
		encCfg.Codec,
		runtime.NewKVStoreService(distrKey),
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
	return ctx, ak, &dk, &sk, &bk
}

var _ = Describe("Distribution Precompile Test", func() {
	var (
		contract *Contract
		valAddr  sdk.ValAddress
		f        *pclog.Factory
		amt      sdk.Coin

		ctx sdk.Context
		ak  authkeeper.AccountKeeperI
		dk  *distrkeeper.Keeper
		sk  *stakingkeeper.Keeper
		bk  *bankkeeper.BaseKeeper
		sf  *ethprecompile.StatefulFactory
	)

	BeforeEach(func() {
		valAddr = sdk.ValAddress([]byte("val"))
		amt = sdk.NewCoin("abera", sdkmath.NewInt(100))

		// Set up the contracts and keepers.
		ctx, ak, dk, sk, bk = setup()
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(
			ak, sk,
			distrkeeper.NewMsgServerImpl(*dk),
			distrkeeper.NewQuerier(*dk),
		))

		// Register the events.
		f = pclog.NewFactory([]ethprecompile.Registrable{contract})

		// Set up the stateful factory.
		sf = ethprecompile.NewStatefulFactory()
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
			_, err := sf.Build(contract, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("SetWithdrawAddress", func() {
		It("should succeed", func() {
			pCtx := vm.NewPolarContext(
				ctx,
				nil,
				testutil.Alice,
				big.NewInt(0),
			)

			res, err := contract.SetWithdrawAddress(
				pCtx,
				testutil.Bob,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("Withdraw Delegator Rewards", func() {
		var addr sdk.AccAddress
		var tokens sdk.DecCoins
		var val stakingtypes.Validator

		BeforeEach(func() {
			// Set the previous proposer.
			Expect(dk.SetPreviousProposerConsAddr(
				ctx,
				sdk.ConsAddress(testutil.Alice.Bytes()),
			)).To(Succeed())

			PKS := simtestutil.CreateTestPubKeys(5)
			valConsPk0 := PKS[0]
			valConsAddr0 := sdk.ConsAddress(valConsPk0.Address())
			valAddr = sdk.ValAddress(valConsAddr0)
			addr = sdk.AccAddress(valAddr)
			var err error
			val, err = distrtestutil.CreateValidator(valConsPk0, sdkmath.NewInt(100))
			Expect(err).ToNot(HaveOccurred())

			// Set the validator.
			Expect(sk.SetValidator(ctx, val)).To(Succeed())

			// Create the delegation.
			Expect(sk.SetDelegation(ctx, stakingtypes.Delegation{
				DelegatorAddress: addr.String(),
				ValidatorAddress: valAddr.String(),
				Shares:           val.DelegatorShares,
			})).To(Succeed())

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
			It("Success", func() {
				pCtx := vm.NewPolarContext(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
				)

				res1, err := contract.GetTotalDelegatorReward(pCtx, common.BytesToAddress(addr))
				Expect(err).ToNot(HaveOccurred())
				Expect(res1[0].Denom).To(Equal(sdk.DefaultBondDenom))
				rewards, _ := tokens.TruncateDecimal()
				expectedReward := rewards[0].Amount.BigInt()
				Expect(res1[0].Amount.Cmp(expectedReward)).To(Equal(0))

				res3, err := contract.GetAllDelegatorRewards(pCtx, common.BytesToAddress(addr))
				Expect(err).ToNot(HaveOccurred())
				Expect(res3[0].Validator).To(Equal(common.BytesToAddress(valAddr)))
				Expect(res3[0].Rewards[0].Amount.Cmp(expectedReward)).To(Equal(0))

				res4, err := contract.GetDelegatorValidatorReward(
					pCtx, common.BytesToAddress(addr), common.BytesToAddress(valAddr),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res4[0].Denom).To(Equal(sdk.DefaultBondDenom))
				Expect(res4[0].Amount.Cmp(expectedReward)).To(Equal(0))

				res2, err := contract.WithdrawDelegatorReward(
					pCtx, common.BytesToAddress(addr), common.BytesToAddress(valAddr),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res2[0].Denom).To(Equal(sdk.DefaultBondDenom))
				Expect(res2[0].Amount).To(Equal(expectedReward))
			})
		})

		When("Reading Params", func() {
			It("Should get if withdraw forwarding is enabled", func() {
				pCtx := vm.NewPolarContext(
					ctx,
					nil,
					testutil.Alice,
					big.NewInt(0),
				)
				res, err := contract.GetWithdrawEnabled(pCtx)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(BeTrue())
			})
		})

		When("Base Precompile Features", func() {
			It("Should not have custom value decoders", func() {
				Expect(contract.CustomValueDecoders()).To(HaveLen(2))
			})

		})
	})
})
