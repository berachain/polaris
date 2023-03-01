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
	"fmt"
	"time"

	pruningtypes "cosmossdk.io/store/pruning/types"
	cdb "github.com/cosmos/cosmos-db"
	baseapp "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ethhd "pkg.berachain.dev/stargazer/crypto/hd"
	ethkeyring "pkg.berachain.dev/stargazer/crypto/keyring"
	app "pkg.berachain.dev/stargazer/testutil/app"
	config "pkg.berachain.dev/stargazer/testutil/app/config"
)

type (
	Network = network.Network
	Config  = network.Config
)

const (
	thousand    = 1000
	fivehundred = 500
	onehundred  = 100
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
	encoding := config.MakeEncodingConfig(app.ModuleBasics)
	cfg := network.Config{
		Codec:             encoding.Codec,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.ValidatorI) servertypes.Application {
			return app.NewSimApp(
				val.GetCtx().Logger, cdb.NewMemDB(), nil, true, sims.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
				baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			)
		},
		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Codec),
		TimeoutCommit:   2 * time.Second, //nolint:gomnd // 2 seconds is the default.
		ChainID:         "stargazer-2061",
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
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
