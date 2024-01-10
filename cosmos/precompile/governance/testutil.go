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

package governance

import (
	"fmt"
	"math/big"

	"github.com/golang/mock/gomock"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	testutils "github.com/berachain/polaris/cosmos/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"

	//nolint:stylecheck,revive // Ginkgo is the testing framework.
	. "github.com/onsi/ginkgo/v2"
)

// Test Reporter to use governance module tests with Ginkgo.
type ginkgoTestReporter struct{}

func (g ginkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g ginkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

// Helper functions for setting up the tests.
// TODO: deprecate this garbage.
func setupGovTest(ctrl *gomock.Controller, caller sdk.AccAddress) (
	sdk.Context, authkeeper.AccountKeeperI, bankkeeper.Keeper, *governancekeeper.Keeper,
) {
	// Setup the keepers and context.
	govKey := storetypes.NewKVStoreKey(governancetypes.StoreKey)
	ctx, ak, bk, sk := testutils.SetupMinimalKeepers(
		log.NewTestLogger(GinkgoT()), []storetypes.StoreKey{govKey}...,
	)
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Create the codec.
	encCfg := testutils.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
		bank.AppModuleBasic{},
	)

	// Create the base app msgRouter.
	msr := baseapp.NewMsgServiceRouter()

	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = "abera"
	err := sk.SetParams(ctx, stakingParams)
	if err != nil {
		panic(err)
	}

	// Create the governance keeper.
	authority, err := ak.AddressCodec().BytesToString(
		authtypes.NewModuleAddress(governancetypes.ModuleName))
	if err != nil {
		panic(err)
	}
	gk := governancekeeper.NewKeeper(
		encCfg.Codec,
		runtime.NewKVStoreService(govKey),
		ak,
		bk,
		sk,
		dk,
		msr,
		governancetypes.DefaultConfig(),
		authority,
	)

	// Register the msg Service Handlers.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v1.RegisterMsgServer(msr, governancekeeper.NewMsgServerImpl(gk))
	banktypes.RegisterMsgServer(msr, bankkeeper.NewMsgServerImpl(bk))

	// Set the Params and first proposal ID.
	params := v1.DefaultParams()
	err = gk.Params.Set(ctx, params)
	if err != nil {
		panic(err)
	}
	// gk.SetProposalID(ctx, 1)

	// Fund the caller with some coins.
	err = testutils.MintCoinsToAddress(

		ctx, bk, governancetypes.ModuleName,
		common.BytesToAddress(caller), "abera", big.NewInt(100000000), //nolint:gomnd // its okay.
	)
	if err != nil {
		panic(err)
	}

	return ctx, ak, bk, gk
}

func SdkCoinsToEvmCoins(sdkCoins sdk.Coins) []struct {
	Amount *big.Int `json:"amount"`
	Denom  string   `json:"denom"`
} {
	evmCoins := make([]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}, len(sdkCoins))
	for i, coin := range sdkCoins {
		evmCoins[i] = struct {
			Amount *big.Int `json:"amount"`
			Denom  string   `json:"denom"`
		}{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		}
	}
	return evmCoins
}
