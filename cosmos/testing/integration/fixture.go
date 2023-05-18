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

package integration

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// defaultTimeout is the default timeout for the test fixture.
const (
	fiveHundredError        = 500
	defaultTimeout          = 10 * time.Second
	defaultNumberOfAccounts = 3
)

var defaultAccountNames = []string{"alice", "bob", "charlie"}

// TestFixture is a testing fixture that can be used to test the
// Ethereum JSON-RPC API.
type TestFixture struct {
	t                network.TestingT
	Network          *network.Network
	EthClient        *ethclient.Client
	EthWsClient      *ethclient.Client
	EthGraphQLClient *http.Client
	HTTPAddr         string
	WsAddr           string
	keysMap          map[string]*ethsecp256k1.PrivKey
}

// NewTestFixture creates a new TestFixture.
func NewTestFixture(t network.TestingT) *TestFixture {
	// For now we just use a context.Background() but we may want to
	// add some timeout functionality in the future.
	ctx := context.Background()

	// Always setup numberOfAccounts accounts.
	keysMap := make(map[string]*ethsecp256k1.PrivKey)
	setupTestAccounts(keysMap)

	// Build Testing Network.
	net := network.New(t, network.DefaultConfig(keysMap))
	_, err := net.WaitForHeightWithTimeout(1, defaultTimeout)
	if err != nil {
		t.Fatal(err)
	}

	apiAddr := strings.Split(net.Validators[0].APIAddress, ":")[1]

	// Dial the Ethereum HTTP Endpoint
	httpAddr := "http:" + apiAddr + ":8545"
	client, _ := ethclient.DialContext(ctx, httpAddr)

	// Dial the Ethereum WS Endpoint
	wsaddr := "ws:" + apiAddr + ":8546"
	wsClient, _ := ethclient.DialContext(ctx, wsaddr)

	graphQLClient := &http.Client{}

	// Build and return the Test Fixture.
	return &TestFixture{
		t:                t,
		Network:          net,
		EthClient:        client,
		EthWsClient:      wsClient,
		EthGraphQLClient: graphQLClient,
		HTTPAddr:         httpAddr,
		WsAddr:           wsaddr,
		keysMap:          keysMap,
	}
}

// GenerateTransactOpts generates a new transaction options object for a key by it's name.
func (tf *TestFixture) GenerateTransactOpts(name string) *bind.TransactOpts {
	// Get the nonce from the RPC.
	nonce, err := tf.EthClient.PendingNonceAt(context.Background(), tf.Address(name))
	if err != nil {
		tf.t.Fatal(err)
	}

	// Get the ChainID from the RPC.
	chainID, err := tf.EthClient.ChainID(context.Background())
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

func (tf *TestFixture) SendGraphQLRequest(query string) (string, int, error) {
	url := tf.HTTPAddr + "/graphql"
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		return "", fiveHundredError, err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fiveHundredError, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := tf.EthGraphQLClient.Do(req)
	if err != nil {
		return "", fiveHundredError, err
	}
	defer resp.Body.Close() //#nosec:G307 // only used in testing.

	responseBody, err := io.ReadAll(resp.Body)
	responseStatusCode := resp.StatusCode
	if err != nil {
		return "", fiveHundredError, err
	}

	return string(responseBody), responseStatusCode, nil
}

func setupTestAccounts(keysMap map[string]*ethsecp256k1.PrivKey) {
	for i := 0; i < defaultNumberOfAccounts; i++ {
		newKey, _ := ethsecp256k1.GenPrivKey()
		keysMap[defaultAccountNames[i]] = newKey
	}
}
