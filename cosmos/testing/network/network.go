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
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtestutil "github.com/cosmos/cosmos-sdk/x/distribution/testutil"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	ethhd "pkg.berachain.dev/polaris/cosmos/crypto/hd"
	ethkeyring "pkg.berachain.dev/polaris/cosmos/crypto/keyring"
	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	runtime "pkg.berachain.dev/polaris/cosmos/runtime"
	config "pkg.berachain.dev/polaris/cosmos/runtime/config"
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

	// Set up the distribution keeper.
	DistrTestSetup(&bankState, &distrState, &stakingState)

	// Governance Module.
	var governanceState v1.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[govtypes.ModuleName], &governanceState)
	// Create the proposal message.
	// subtract one hour from  time.Now .
	voteStart := time.Now().Add(-time.Hour)
	//nolint:gomnd // 2 days.
	voteEnd := voteStart.Add(time.Hour * 24 * 2)
	proposal := &v1.Proposal{
		Id:               2, //nolint:gomnd // not important.
		Proposer:         TestAddress.String(),
		Messages:         []*codectypes.Any{},
		Status:           v1.StatusVotingPeriod,
		FinalTallyResult: &v1.TallyResult{},
		SubmitTime:       &time.Time{},
		DepositEndTime:   &time.Time{},
		TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(onehundred))),
		VotingStartTime:  &voteStart,
		VotingEndTime:    &voteEnd,
		Metadata:         "metadata",
		Title:            "title",
		Summary:          "summary",
		Expedited:        false,
	}
	// Append the proposal to the governance genesis state.
	governanceState.Proposals = append(governanceState.Proposals, proposal)

	// Set the states into the genesis state.
	genState[authtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&authState)
	genState[banktypes.ModuleName] = encoding.Codec.MustMarshalJSON(&bankState)
	genState[stakingtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&stakingState)
	genState[distrtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&distrState)
	genState[govtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&governanceState)
	return genState
}

func GetDistrValidator() common.Address {
	pks := simtestutil.CreateTestPubKeys(2) //nolint: gomnd // magic numbers are fine in tests.
	valConsPk1 := pks[1]
	valConsAddr1 := sdk.ConsAddress(valConsPk1.Address())
	valAddr := sdk.ValAddress(valConsAddr1)
	return cosmlib.ValAddressToEthAddress(valAddr)
}

func DistrTestSetup(
	bk *banktypes.GenesisState,
	dk *distrtypes.GenesisState,
	sk *stakingtypes.GenesisState,
) {
	// ==============================================================================
	// Staking Keeper
	// ==============================================================================

	// Create the validator.
	//nolint: gomnd // magic numbers are fine in tests.
	pks := simtestutil.CreateTestPubKeys(2)
	valConsPk1 := pks[1]
	valConsAddr1 := sdk.ConsAddress(valConsPk1.Address())
	valAddr := sdk.ValAddress(valConsAddr1)
	val, err := distrtestutil.CreateValidator(valConsPk1, math.NewInt(onehundred))
	if err != nil {
		panic(err)
	}
	val.Status = stakingtypes.Unbonded
	val.Tokens = math.NewInt(onehundred)
	val.DelegatorShares = sdk.NewDec(onehundred)

	// Set the validator.
	sk.Validators = append(sk.Validators, val)

	// Set the delegations.
	delegator := cosmlib.AddressToAccAddress(TestAddress)
	sk.Delegations = append(sk.Delegations, stakingtypes.Delegation{
		DelegatorAddress: delegator.String(),
		ValidatorAddress: valAddr.String(),
		Shares:           val.DelegatorShares,
	})

	// ==============================================================================
	// Distribution Keeper
	// ==============================================================================

	// Set the parameters.
	dk.Params = distrtypes.DefaultParams()
	dk.Params.WithdrawAddrEnabled = true

	// Set the validator accumulated commission.
	dk.ValidatorAccumulatedCommissions = append(dk.ValidatorAccumulatedCommissions,
		distrtypes.ValidatorAccumulatedCommissionRecord{
			ValidatorAddress: valAddr.String(),
		})

	// Set the validator historical rewards.
	dk.ValidatorHistoricalRewards = append(dk.ValidatorHistoricalRewards,
		distrtypes.ValidatorHistoricalRewardsRecord{
			ValidatorAddress: valAddr.String(),
			Period:           1,
			Rewards: distrtypes.ValidatorHistoricalRewards{
				ReferenceCount: 2, //nolint:gomnd // test.
			},
		})

	// Set the validator current rewards.
	dk.ValidatorCurrentRewards = append(dk.ValidatorCurrentRewards,
		distrtypes.ValidatorCurrentRewardsRecord{
			ValidatorAddress: valAddr.String(),
			Rewards: distrtypes.ValidatorCurrentRewards{
				Rewards: sdk.NewDecCoins(sdk.NewDecCoin("abera", sdk.NewInt(onehundred))),
				Period:  2, //nolint:gomnd // test.
			},
		})

	// Set the delegator starting info.
	dk.DelegatorStartingInfos = append(dk.DelegatorStartingInfos,
		distrtypes.DelegatorStartingInfoRecord{
			DelegatorAddress: delegator.String(),
			ValidatorAddress: valAddr.String(),
			StartingInfo: distrtypes.DelegatorStartingInfo{
				PreviousPeriod: 1,
				Stake:          sdk.NewDec(onehundred),
			},
		})

	// ==============================================================================
	// Bank Keeper
	// ==============================================================================

	// Set the not bonded pool balance to 100.
	bk.Balances = append(bk.Balances, banktypes.Balance{
		Address: "polar1tygms3xhhs3yv487phx3dw4a95jn7t7l2g6um3", // Not Bonded Pool
		Coins:   sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(onehundred))),
	})
}
