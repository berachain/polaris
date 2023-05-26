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

package test

import (
	"fmt"
	"math/big"

	"github.com/golang/mock/gomock"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"

	//nolint:stylecheck,revive // Ginkgo is the testing framework.
	. "github.com/onsi/ginkgo/v2"
)

// Test Reporter to use governance module tests with Ginkgo.
type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

// Helper functions for setting up the tests.
func Setup(ctrl *gomock.Controller, caller sdk.AccAddress) (sdk.Context, bankkeeper.Keeper, *governancekeeper.Keeper) {
	// Setup the keepers and context.
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Register the governance module account.
	ak.SetModuleAccount(
		ctx,
		authtypes.NewEmptyModuleAccount(governancetypes.ModuleName, authtypes.Minter),
	)

	// Create the codec.
	encCfg := cosmostestutil.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
		bank.AppModuleBasic{},
	)

	// Create the base app msgRouter.
	msr := baseapp.NewMsgServiceRouter()

	// Create the governance keeper.
	gk := governancekeeper.NewKeeper(
		encCfg.Codec,
		runtime.NewKVStoreService(storetypes.NewKVStoreKey(governancetypes.StoreKey)),
		ak,
		bk,
		sk,
		dk,
		msr,
		governancetypes.DefaultConfig(),
		authtypes.NewModuleAddress(governancetypes.ModuleName).String(),
	)

	// Register the msg Service Handlers.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v1.RegisterMsgServer(msr, governancekeeper.NewMsgServerImpl(gk))
	banktypes.RegisterMsgServer(msr, bankkeeper.NewMsgServerImpl(bk))

	// Set the Params and first proposal ID.
	params := v1.DefaultParams()
	err := gk.Params.Set(ctx, params)
	if err != nil {
		panic(err)
	}
	// gk.SetProposalID(ctx, 1)

	// Fund the caller with some coins.
	err = lib.MintCoinsToAddress(
		//nolint:gomnd // magic number is fine here.
		ctx, bk, governancetypes.ModuleName, lib.AccAddressToEthAddress(caller), "abera", big.NewInt(100000000),
	)
	if err != nil {
		panic(err)
	}

	return ctx, bk, gk
}
