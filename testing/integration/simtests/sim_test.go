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

package sim_test

import (
	"testing"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/testutil/network"
)

var (
	// dummyContract  = network.DummyContract
	testKey        = network.TestKey
	addressFromKey = network.AddressFromKey
	signer         = network.Signer

	txData = network.TxData
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testutil/sim:integration")
}

var _ = Describe("BlockAPIs", func() {
	var net *network.Network
	var client *ethclient.Client

	BeforeEach(func() {
		cfg := network.DefaultConfig()

		var authState authtypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authState)
		newAccount := authtypes.NewBaseAccount(addressFromKey.Bytes(), testKey.PubKey(), 99, 0)
		accounts, err := authtypes.PackAccounts([]authtypes.GenesisAccount{newAccount})
		Expect(err).To(BeNil())
		authState.Accounts = append(authState.Accounts, accounts[0])
		cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authState)

		var bankState banktypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankState)
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: sdk.MustBech32ifyAddressBytes("cosmos", addressFromKey.Bytes()),
			Coins:   sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000000000000))),
		})
		cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankState)

		net = network.New(GinkgoT(), cfg)
		_, err = net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).To(BeNil())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())

	})
})
