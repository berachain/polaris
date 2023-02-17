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

package testutil

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/libs/log"
	cometproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/berachain/stargazer/eth/common"
	mock "github.com/berachain/stargazer/testutil/mock"
)

var (
	AccKey     = storetypes.NewKVStoreKey("acc")
	BankKey    = storetypes.NewKVStoreKey("bank")
	EvmKey     = storetypes.NewKVStoreKey("evm")
	StakingKey = storetypes.NewKVStoreKey("staking")
	Alice      = common.BytesToAddress([]byte("alice"))
	Bob        = common.BytesToAddress([]byte("bob"))
)

// `NewContext` creates a SDK context and mounts a mock multistore.
func NewContext() sdk.Context {
	return sdk.NewContext(mock.NewMultiStore(), cometproto.Header{}, false, log.TestingLogger())
}

func NewContextWithMultiStore(ms storetypes.MultiStore) sdk.Context {
	return sdk.NewContext(ms, cometproto.Header{}, false, log.TestingLogger())
}

// `SetupMinimalKeepers` creates and returns keepers for the base SDK modules.
func SetupMinimalKeepers() (
	sdk.Context,
	authkeeper.AccountKeeper,
	bankkeeper.BaseKeeper,
	stakingkeeper.Keeper,
) {
	ctx := NewContext()

	encodingConfig := testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
	)

	ak := authkeeper.NewAccountKeeper(
		encodingConfig.Codec,
		AccKey,
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.NotBondedPoolName: {authtypes.Minter, authtypes.Burner},
			stakingtypes.BondedPoolName:    {authtypes.Minter, authtypes.Burner},
			"evm":                          {authtypes.Minter, authtypes.Burner},
			"staking":                      {authtypes.Minter, authtypes.Burner},
		},
		"bera",
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	ak.SetModuleAccount(ctx,
		authtypes.NewEmptyModuleAccount("evm", authtypes.Minter, authtypes.Burner))

	bk := bankkeeper.NewBaseKeeper(
		encodingConfig.Codec,
		BankKey,
		ak,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	sk := stakingkeeper.NewKeeper(
		encodingConfig.Codec,
		StakingKey,
		ak,
		bk,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	return ctx, ak, bk, *sk
}
