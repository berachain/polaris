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

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	runtime "pkg.berachain.dev/polaris/cosmos/runtime"
	config "pkg.berachain.dev/polaris/cosmos/runtime/config"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/params"
)

const (
	numTestAccounts = 10
)

var (
	TestKeys      [numTestAccounts]*ethsecp256k1.PrivKey
	TestAddresses [numTestAccounts]common.Address
	Signer        = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
)

func init() {
	for i := 0; i < numTestAccounts; i++ {
		key, _ := ethsecp256k1.GenPrivKey()
		TestKeys[i] = key
		TestAddresses[i] = common.Address(key.PubKey().Address().Bytes())
	}
}

// BuildGenesisState returns a genesis state for the runtime module.
func BuildGenesisState() map[string]json.RawMessage {
	encoding := config.MakeEncodingConfig(runtime.ModuleBasics)
	genState := runtime.ModuleBasics.DefaultGenesis(encoding.Codec)

	// Auth + Bank module
	var bankState banktypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[banktypes.ModuleName], &bankState)
	var authState authtypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[authtypes.ModuleName], &authState)
	for i := 0; i < numTestAccounts; i++ {
		newAccount, err := authtypes.NewBaseAccountWithPubKey(TestKeys[i].PubKey())
		if err != nil {
			panic(err)
		}
		accounts, _ := authtypes.PackAccounts([]authtypes.GenesisAccount{newAccount})
		authState.Accounts = append(authState.Accounts, accounts[0])
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: newAccount.Address,
			Coins:   sdk.NewCoins(sdk.NewCoin("abera", sdk.NewInt(megamoney))),
		})
	}
	genState[authtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&authState)
	genState[banktypes.ModuleName] = encoding.Codec.MustMarshalJSON(&bankState)

	// Staking module
	var stakingState stakingtypes.GenesisState
	encoding.Codec.MustUnmarshalJSON(genState[stakingtypes.ModuleName], &stakingState)
	stakingState.Params.BondDenom = "abera"
	genState[stakingtypes.ModuleName] = encoding.Codec.MustMarshalJSON(&stakingState)

	return genState
}
