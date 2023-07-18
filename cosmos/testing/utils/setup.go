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

package utils

import (
	"errors"
	"testing"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	cometproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"pkg.berachain.dev/polaris/cosmos/testing/types/mock"
	"pkg.berachain.dev/polaris/cosmos/types"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	coremock "pkg.berachain.dev/polaris/eth/core/mock"
)

var (
	AccKey     = storetypes.NewKVStoreKey("acc")
	BankKey    = storetypes.NewKVStoreKey("bank")
	EvmKey     = storetypes.NewKVStoreKey("evm")
	StakingKey = storetypes.NewKVStoreKey("staking")
	Alice      = common.BytesToAddress([]byte("alice"))
	Bob        = common.BytesToAddress([]byte("bob"))
)

// NewContext creates a SDK context and mounts a mock multistore.
func NewContext() sdk.Context {
	ms := mock.NewMultiStore()
	return sdk.NewContext(ms, cometproto.Header{}, false, log.NewTestLogger(&testing.T{}))
}

func NewContextWithMultiStore(ms storetypes.MultiStore) sdk.Context {
	return sdk.NewContext(ms, cometproto.Header{}, false, log.NewTestLogger(&testing.T{}))
}

// TestEncodingConfig defines an encoding configuration that is used for testing
// purposes. Note, MakeTestEncodingConfig takes a series of AppModuleBasic types
// which should only contain the relevant module being tested and any potential
// dependencies.
type TestEncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeTestEncodingConfig(modules ...module.AppModuleBasic) TestEncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectestutil.CodecOptions{
		AccAddressPrefix: "polar",
		ValAddressPrefix: "polarvaloper",
	}.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}

	mb := module.NewBasicManager(modules...)

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}

// SetupMinimalKeepers creates and returns keepers for the base SDK modules.
func SetupMinimalKeepers() (
	sdk.Context,
	authkeeper.AccountKeeper,
	bankkeeper.BaseKeeper,
	stakingkeeper.Keeper,
) {
	types.SetupCosmosConfig()
	ctx := NewContext().WithBlockHeight(1)

	encodingConfig := testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		authz.AppModuleBasic{},
	)

	ak := authkeeper.NewAccountKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(AccKey),
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.NotBondedPoolName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
			stakingtypes.BondedPoolName:    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
			evmtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
			erc20types.ModuleName:          {authtypes.Minter, authtypes.Burner},
			stakingtypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
			govtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
			distrtypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
		},
		// TODO: switch to eip-55 fuck bech32.
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		types.Bech32Prefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	bk := bankkeeper.NewBaseKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(BankKey),
		ak,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewTestLogger(&testing.T{}),
	)

	sk := stakingkeeper.NewKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(StakingKey),
		ak,
		bk,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	return ctx, ak, bk, *sk
}

func GetEncodingConfig() testutil.TestEncodingConfig {
	return testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		authz.AppModuleBasic{},
	)
}

func MockQueryContext(height int64, _ bool) (sdk.Context, error) {
	if height <= 0 {
		return sdk.Context{}, errors.New("cannot query context at this height")
	}
	ctx := NewContext().WithBlockHeight(height)
	header := coremock.GenerateHeaderAtHeight(height)
	headerBz, err := coretypes.MarshalHeader(header)
	if err != nil {
		return sdk.Context{}, err
	}
	ctx.KVStore(EvmKey).Set([]byte{evmtypes.HeaderKey}, headerBz)
	ctx.KVStore(EvmKey).Set(header.Hash().Bytes(), header.Number.Bytes())
	return ctx, nil
}
