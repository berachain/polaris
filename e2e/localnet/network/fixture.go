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

package localnet

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// defaultTimeout is the default timeout for the test fixture.
const (
	fiveHundredError        = 500
	defaultNumberOfAccounts = 3
	defaultWaitForHeight    = 5

	onehundred = 100
	examoney   = 1000000000000000000

	defaultAccountsFile = "default_accounts.json"
)

var defaultAccountNames = []string{"alice", "bob", "charlie"}

type TestingT interface {
	Fatal(args ...interface{})
	Cleanup(func())
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	TempDir() string
}

// TestFixture is a testing fixture that can be used to test the
// Ethereum JSON-RPC API.
type TestFixture struct {
	t       TestingT
	c       ContainerizedNode
	keysMap map[string]*ethsecp256k1.PrivKey
}

// NewTestFixture creates a new TestFixture.
func NewTestFixture(t TestingT) *TestFixture {
	types.SetupCosmosConfig()

	// Always setup numberOfAccounts accounts.
	keysMap := make(map[string]*ethsecp256k1.PrivKey)
	setupTestAccounts(keysMap)

	containerizedNode, err := NewContainerizedNode(
		"localnet",
		"latest",
		"goodcontainer",
		"8545/tcp",
		"8546/tcp",
		[]string{
			"GO_VERSION=1.20.4",
			"GENESIS_PATH=config",
			"BASE_IMAGE=polard/base:v0.0.0",
			"DEFAULT_ACCOUNTS=" + defaultAccountsFile,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Build and return the Test Fixture.
	return &TestFixture{
		t:       t,
		c:       containerizedNode,
		keysMap: keysMap,
	}
}

// GenerateTransactOpts generates a new transaction options object for a key by it's name.
func (tf *TestFixture) GenerateTransactOpts(name string) *bind.TransactOpts {
	// Get the nonce from the RPC.
	nonce, err := tf.c.EthClient().PendingNonceAt(context.Background(), tf.Address(name))
	if err != nil {
		tf.t.Fatal(err)
	}

	// Get the ChainID from the RPC.
	chainID, err := tf.c.EthClient().ChainID(context.Background())
	if err != nil {
		tf.t.Fatal(err)
	}

	// Build transaction opts object.
	auth, err := bind.NewKeyedTransactorWithChainID(tf.PrivKey(name), chainID)
	if err != nil {
		tf.t.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	return auth
}

func (tf *TestFixture) PrivKey(name string) *ecdsa.PrivateKey {
	newECDSATestKey, _ := tf.keysMap[name].ToECDSA()
	return newECDSATestKey
}

func (tf *TestFixture) Address(name string) common.Address {
	return crypto.PubkeyToAddress(tf.PrivKey(name).PublicKey)
}

func (tf *TestFixture) CreateKeyWithName(name string) {
	newKey, _ := ethsecp256k1.GenPrivKey()
	tf.keysMap[name] = newKey
}

type AccountInfo struct {
	Name          string    `json:"name"`
	Bech32Address string    `json:"bech32Address"`
	EthAddress    string    `json:"ethAddress"`
	Coins         sdk.Coins `json:"coins"`
}

func setupTestAccounts(keysMap map[string]*ethsecp256k1.PrivKey) {
	var accounts []AccountInfo
	for _, name := range defaultAccountNames {
		newKey, _ := ethsecp256k1.GenPrivKey()
		keysMap[name] = newKey
		privateKey, _ := newKey.ToECDSA()

		accounts = append(
			accounts,
			AccountInfo{
				Name:          name,
				Bech32Address: cosmlib.Bech32FromEthAddress((crypto.PubkeyToAddress(privateKey.PublicKey))),
				EthAddress:    crypto.PubkeyToAddress(privateKey.PublicKey).Hex()[2:],
				Coins:         getCoinsForAccount(name),
			},
		)
	}

	jsonBytes, _ := json.MarshalIndent(accounts, "", "   ")
	if err := os.WriteFile(defaultAccountsFile, jsonBytes, 0644); err != nil {
		panic(err)
	}
}

func getCoinsForAccount(name string) sdk.Coins {
	switch name {
	case "alice":
		return sdk.NewCoins(
			sdk.NewCoin("abera", sdkmath.NewInt(examoney)),
			sdk.NewCoin("bATOM", sdkmath.NewInt(examoney)),
			sdk.NewCoin("bAKT", sdkmath.NewInt(12345)), //nolint:gomnd // its okay.
			sdk.NewCoin("stake", sdkmath.NewInt(examoney)),
			sdk.NewCoin("bOSMO", sdkmath.NewInt(12345*2)), //nolint:gomnd // its okay.
			sdk.NewCoin("atoken", sdkmath.NewInt(examoney)),
			sdk.NewCoin("eth", sdkmath.NewInt(examoney)),
			// do not change the supply of this coin
			sdk.NewCoin("asupply", sdkmath.NewInt(examoney)),
		)
	case "bob":
		return sdk.NewCoins(
			sdk.NewCoin("abera", sdkmath.NewInt(onehundred)),
			sdk.NewCoin("atoken", sdkmath.NewInt(onehundred)),
			sdk.NewCoin("stake", sdkmath.NewInt(examoney)),
			sdk.NewCoin("eth", sdkmath.NewInt(examoney)),
		)
	case "charlie":
		return sdk.NewCoins(
			sdk.NewCoin("abera", sdkmath.NewInt(examoney)),
			sdk.NewCoin("eth", sdkmath.NewInt(examoney)),
		)
	default:
		return sdk.NewCoins(sdk.NewCoin("abera", sdkmath.NewInt(examoney)))
	}
}
