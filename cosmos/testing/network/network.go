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

package network

import (
	"encoding/json"
	"fmt"
	"time"

	cdb "github.com/cosmos/cosmos-db"

	"cosmossdk.io/math"
	pruningtypes "cosmossdk.io/store/pruning/types"

	baseapp "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtestutil "github.com/cosmos/cosmos-sdk/x/distribution/testutil"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	ethhd "pkg.berachain.dev/polaris/cosmos/crypto/hd"
	ethkeyring "pkg.berachain.dev/polaris/cosmos/crypto/keyring"
	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	runtime "pkg.berachain.dev/polaris/cosmos/runtime"
	config "pkg.berachain.dev/polaris/cosmos/runtime/config"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"
)

type (
	Network = network.Network
	Config  = network.Config
)

const (
	thousand    = 1000
	fivehundred = 500
	onehundred  = 100
	megamoney   = 1000000000000000000
	two         = 2
)

var (
	DummyContract   = common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
	TestKey, _      = ethsecp256k1.GenPrivKey()
	ECDSATestKey, _ = TestKey.ToECDSA()
	AddressFromKey  = TestKey.PubKey().Address()
	Signer          = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
	TestAddress     = crypto.PubkeyToAddress(ECDSATestKey.PublicKey)
)

type TestingT interface {
	Fatal(args ...interface{})
	Cleanup(func())
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	TempDir() string
}

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t TestingT, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}

	net, err := network.New(t, t.TempDir(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig.
func DefaultConfig() network.Config {
	encoding := config.MakeEncodingConfig(runtime.ModuleBasics)
	cfg := network.Config{
		Codec:             encoding.Codec,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.ValidatorI) servertypes.Application {
			return runtime.NewPolarisApp(
				val.GetCtx().Logger, cdb.NewMemDB(), nil, true, simtestutil.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
				baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
				baseapp.SetChainID("polaris-2061"),
			)
		},
		GenesisState:    BuildGenesisState(),
		TimeoutCommit:   two * time.Second,
		ChainID:         "polaris-2061",
		NumValidators:   1,
		BondDenom:       "abera",
		MinGasPrices:    fmt.Sprintf("0.00006%s", "abera"),
		AccountTokens:   sdk.TokensFromConsensusPower(thousand, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(fivehundred, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(onehundred, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      false,
		SigningAlgo:     string(ethhd.EthSecp256k1Type),
		KeyringOptions:  []keyring.Option{ethkeyring.EthSecp256k1Option()},
	}

	return cfg
}

func BuildGenesisState() map[string]json.RawMessage {
	encoding := config.MakeEncodingConfig(runtime.ModuleBasics)
	genState := runtime.ModuleBasics.DefaultGenesis(encoding.Codec)

	// Auth module
	var authState authtypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[authtypes.ModuleName], &authState)
	newAccount, err := authtypes.NewBaseAccountWithPubKey(TestKey.PubKey())
	if err != nil {
		panic(err)
	}
	accounts, _ := authtypes.PackAccounts([]authtypes.GenesisAccount{newAccount})
	authState.Accounts = append(authState.Accounts, accounts[0])

	// Bank module
	var bankState banktypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[banktypes.ModuleName], &bankState)
	bankState.Balances = append(bankState.Balances, banktypes.Balance{
		Address: newAccount.Address,
		Coins:   sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(megamoney))),
	})

	// Staking module
	var stakingState stakingtypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[stakingtypes.ModuleName], &stakingState)
	stakingState.Params.BondDenom = "abera"

	// Distribution module
	var distrState distrtypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[distrtypes.ModuleName], &distrState)
	distrState.Params.WithdrawAddrEnabled = true

	// TODO: Fix the state invariants that are being thrown.
	// For the distribution module, we need set it up for having rewards ready to be withdrawn.
	// DistributionGenesisState(&bankState, &distrState, &stakingState)

	// Set the states into the genesis state.
	genState[authtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&authState)
	genState[banktypes.ModuleName] = encoding.Codec.MustMarshalJSON(&bankState)
	genState[stakingtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&stakingState)
	genState[distrtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&distrState)
	return genState
}

// Pass in the genesis state of the bank module and the distribution module.
func DistributionGenesisState(
	bankGenState *banktypes.GenesisState,
	distrGenState *distrtypes.GenesisState,
	stakingGenState *stakingtypes.GenesisState,
) {
	delegator := cosmlib.AddressToAccAddress(TestAddress)

	// Set the allow set withdraw address to true.
	distrGenState.Params.WithdrawAddrEnabled = true

	// Set the previous proposer.
	distrGenState.PreviousProposer = sdk.ConsAddress(testutil.Alice.Bytes()).String()

	// Create the validator.
	//nolint: gomnd // magic numbers are fine in tests.
	pks := simtestutil.CreateTestPubKeys(5)
	valConsPk0 := pks[0]
	valConsAddr0 := sdk.ConsAddress(valConsPk0.Address())
	valAddr := sdk.ValAddress(valConsAddr0)
	val, err := distrtestutil.CreateValidator(valConsPk0, math.NewInt(onehundred))
	if err != nil {
		panic(err)
	}
	operator, err := sdk.ValAddressFromBech32(val.OperatorAddress)
	if err != nil {
		panic(err)
	}

	// Set the outstanding rewards.
	reward := sdk.NewDecCoinsFromCoins(sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(onehundred)))...)
	distrGenState.OutstandingRewards = append(distrGenState.OutstandingRewards,
		distrtypes.ValidatorOutstandingRewardsRecord{
			ValidatorAddress:   operator.String(),
			OutstandingRewards: reward,
		})

	// Set the validator historical rewards.
	distrGenState.ValidatorHistoricalRewards = append(distrGenState.ValidatorHistoricalRewards,
		distrtypes.ValidatorHistoricalRewardsRecord{
			ValidatorAddress: operator.String(),
			Period:           1,
			Rewards: distrtypes.ValidatorHistoricalRewards{
				CumulativeRewardRatio: sdk.DecCoins{},
				ReferenceCount:        two,
			},
		})

	// Set the validator current rewards.
	distrGenState.ValidatorCurrentRewards = append(distrGenState.ValidatorCurrentRewards,
		distrtypes.ValidatorCurrentRewardsRecord{
			ValidatorAddress: operator.String(),
			Rewards: distrtypes.ValidatorCurrentRewards{
				Rewards: reward,
				Period:  two,
			},
		})

	// Set the delegator starting info.
	distrGenState.DelegatorStartingInfos = append(distrGenState.DelegatorStartingInfos,
		distrtypes.DelegatorStartingInfoRecord{
			DelegatorAddress: delegator.String(),
			ValidatorAddress: operator.String(),
			StartingInfo: distrtypes.DelegatorStartingInfo{
				PreviousPeriod: 1,
				Stake:          sdk.MustNewDecFromStr("onehundred"),
			},
		})

	// Mint the tokens in the bank distr module.
	moduleAddr, err := sdk.AccAddressFromBech32("cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl")
	if err != nil {
		panic(err)
	}
	bankGenState.Balances = append(
		bankGenState.Balances,
		banktypes.Balance{
			Address: moduleAddr.String(),
			Coins:   sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(onehundred))),
		})

	// Mint the tokens in the not bonded pool.
	// Add 100abera to the not bonded pool.
	notBondedPoolAddr, err := sdk.AccAddressFromBech32("cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r")
	if err != nil {
		panic(err)
	}
	bankGenState.Balances = append(bankGenState.Balances, banktypes.Balance{
		Address: notBondedPoolAddr.String(),
		Coins:   sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(onehundred))),
	})

	// Set Delegation in the staking module.
	stakingGenState.Delegations = append(stakingGenState.Delegations, stakingtypes.Delegation{
		DelegatorAddress: delegator.String(),
		ValidatorAddress: valAddr.String(),
		Shares:           val.DelegatorShares,
	})

	// Add the validator to the staking state.
	stakingGenState.Validators = append(stakingGenState.Validators, val)
}
