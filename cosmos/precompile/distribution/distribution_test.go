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

	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/lib/utils"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/core/vm"

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

func setup() (sdk.Context, *distrkeeper.Keeper) {
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()

	encCfg := cosmostestutil.MakeTestEncodingConfig(
		distribution.AppModuleBasic{},
	)

	ak.SetModuleAccount(ctx, authtypes.NewEmptyModuleAccount(distributiontypes.ModuleName, authtypes.Minter, authtypes.Burner))

	dk := distrkeeper.NewKeeper(
		encCfg.Codec,
		testutil.EvmKey,
		ak,
		bk,
		sk,
		"gov",
		authtypes.NewModuleAddress("gov").String(),
	)

	params := distributiontypes.DefaultParams()
	params.WithdrawAddrEnabled = true
	dk.SetParams(ctx, params)

	return ctx, &dk
}

var _ = Describe("Distribution Precompile Test", func() {
	var (
		contract *Contract
		valAddr  sdk.ValAddress
		f        *log.Factory
		amt      sdk.Coin

		ctx sdk.Context
		dk  *distrkeeper.Keeper
	)

	BeforeEach(func() {
		valAddr = sdk.ValAddress([]byte("val"))
		amt = sdk.NewCoin("denom", sdk.NewInt(100))

		// Set up the contracts and keepers.
		ctx, dk = setup()
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(&dk))

		// Register the events.
		f = log.NewFactory([]vm.RegistrablePrecompile{contract})
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

	When("SetWithdrawAddress", func() {
		It("should fail if not common address", func() {
			res, err := contract.SetWithdrawAddress(
				ctx,
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
			res, err := contract.SetWithdrawAddressBech32(
				ctx,
				testutil.Alice,
				big.NewInt(0),
				false,
				1,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})

		It("should fail if not bech32 string", func() {
			res, err := contract.SetWithdrawAddressBech32(
				ctx,
				testutil.Alice,
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})

		It("should succeed", func() {
			res, err := contract.SetWithdrawAddressBech32(
				ctx,
				testutil.Alice,
				big.NewInt(0),
				false,
				cosmlib.AddressToAccAddress(testutil.Bob).String(),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})
})
