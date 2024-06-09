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

package testutil

import (
	"testing"

	cdb "github.com/cosmos/cosmos-db"

	"cosmossdk.io/log"
	"cosmossdk.io/store/metrics"
	"cosmossdk.io/store/rootmulti"
	storetypes "cosmossdk.io/store/types"

	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"

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
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"
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
func NewContext(logger log.Logger, storekeys ...storetypes.StoreKey) sdk.Context {
	cdb := cdb.NewMemDB()
	rms := rootmulti.NewStore(cdb, logger, metrics.NewNoOpMetrics())

	// Register defaults
	rms.MountStoreWithDB(AccKey, storetypes.StoreTypeIAVL, cdb)
	rms.MountStoreWithDB(BankKey, storetypes.StoreTypeIAVL, cdb)
	rms.MountStoreWithDB(EvmKey, storetypes.StoreTypeIAVL, cdb)
	rms.MountStoreWithDB(StakingKey, storetypes.StoreTypeIAVL, cdb)

	// Allow extending the
	for _, storeKey := range storekeys {
		rms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, cdb)
	}
	_ = rms.LoadLatestVersion()
	return NewContextWithMultiStore(rms, logger)
}

func NewContextWithMultiStore(ms storetypes.MultiStore, logger log.Logger) sdk.Context {
	return sdk.NewContext(ms, cometproto.Header{}, false, logger)
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
		AccAddressPrefix: "cosmos",
		ValAddressPrefix: "cosmosvaloper",
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
func SetupMinimalKeepers(logger log.Logger, keys ...storetypes.StoreKey) (
	sdk.Context,
	authkeeper.AccountKeeper,
	bankkeeper.BaseKeeper,
	stakingkeeper.Keeper,
) {
	ctx := NewContext(logger, keys...).WithBlockHeight(1)

	encodingConfig := testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
	)

	addrCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	authority, err := addrCodec.BytesToString(authtypes.NewModuleAddress(govtypes.ModuleName))
	if err != nil {
		panic(err)
	}
	ak := authkeeper.NewAccountKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(AccKey),
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.NotBondedPoolName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
			stakingtypes.BondedPoolName:    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
			evmtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
			stakingtypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
			govtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
			distrtypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
		},
		addrCodec,
		"cosmos",
		authority,
	)

	bk := bankkeeper.NewBaseKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(BankKey),
		ak,
		nil,
		authority,
		log.NewTestLogger(&testing.T{}),
	)

	sk := stakingkeeper.NewKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(StakingKey),
		ak,
		bk,
		authority,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	return ctx, ak, bk, *sk
}

func GetEncodingConfig() testutil.TestEncodingConfig {
	return testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
	)
}
