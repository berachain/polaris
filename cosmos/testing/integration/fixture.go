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
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// defaultTimeout is the default timeout for the test fixture.
const defaultTimeout = 10 * time.Second

// TestFixture is a testing fixture that can be used to test the
// Ethereum JSON-RPC API.
type TestFixture struct {
	t           network.TestingT
	Network     *network.Network
	EthClient   *ethclient.Client
	EthWsClient *ethclient.Client
}

// NewTestFixture creates a new TestFixture.
func NewTestFixture(t network.TestingT) *TestFixture {
	// For now we just use a context.Background() but we may want to
	// add some timeout functionality in the future.
	ctx := context.Background()

	// Build Testing Network.
	net := network.New(t, network.DefaultConfig())
	_, err := net.WaitForHeightWithTimeout(1, defaultTimeout)
	if err != nil {
		t.Fatal(err)
	}

	// Dial the Ethereum HTTP Endpoint
	client, _ := ethclient.DialContext(ctx, net.Validators[0].APIAddress+"/eth/rpc")

	// Dial the Ethereum WS Endpoint
	wsaddr := "ws" + strings.TrimPrefix(net.Validators[0].APIAddress+"/eth/rpc/ws/", "http")
	wsClient, _ := ethclient.DialContext(ctx, wsaddr)

	// Build and return the Test Fixture.
	return &TestFixture{
		t:           t,
		Network:     net,
		EthClient:   client,
		EthWsClient: wsClient,
	}
}

func (tf *TestFixture) GenerateTransactOpts(name string) *bind.TransactOpts {
	// Get the nonce from the RPC.
	// TODO: switch to pending once the txpool is finished. https://github.com/berachain/polaris/issues/385
	// Get the nonce from the RPC.
	blockNumber, err := tf.EthClient.BlockNumber(context.Background())
	if err != nil {
		tf.t.Fatal(err)
	}
	// nonce, err := client.PendingNonceAt(context.Background(), network.TestAddress)
	// hacky stuff to make sure the nonce is correct.
	time.Sleep(2) //nolint:gomnd,staticcheck // temporary.
	nonce, err := tf.EthClient.NonceAt(context.Background(), network.TestAddress, big.NewInt(int64(blockNumber)))
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
	return network.ECDSATestKey
}

func (tf *TestFixture) Address(name string) common.Address {
	return crypto.PubkeyToAddress(tf.PrivKey(name).PublicKey)
}
