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
	"fmt"
	"math/big"
	"strings"
	"time"

	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// defaultTimeout is the default timeout for the test fixture.
const (
	fiveHundredError        = 500
	defaultTimeout          = 30 * time.Second
	defaultNumberOfAccounts = 3
	defaultWaitForHeight    = 5
)

var defaultAccountNames = []string{"alice", "bob", "charlie"}

// TestFixture is a testing fixture that can be used to test the
// Ethereum JSON-RPC API.
type TestFixture struct {
	t           network.TestingT
	Network     *network.Network
	EthClient   *ethclient.Client
	EthWsClient *ethclient.Client
	HTTPAddr    string
	WsAddr      string
	keysMap     map[string]*ethsecp256k1.PrivKey
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
	_, err := net.WaitForHeightWithTimeout(defaultWaitForHeight, defaultTimeout)
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

	// Build and return the Test Fixture.
	return &TestFixture{
		t:           t,
		Network:     net,
		EthClient:   client,
		EthWsClient: wsClient,
		HTTPAddr:    httpAddr,
		WsAddr:      wsaddr,
		keysMap:     keysMap,
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

func setupTestAccounts(keysMap map[string]*ethsecp256k1.PrivKey) {
	for i := 0; i < defaultNumberOfAccounts; i++ {
		newKey, _ := ethsecp256k1.GenPrivKey()
		keysMap[defaultAccountNames[i]] = newKey
	}
}

func (tf *TestFixture) BankSendTx(from, to common.Address, amount int64) error {
	val := tf.Network.Validators[0]

	txBuilder := tf.Network.Config.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(&banktypes.MsgSend{
		FromAddress: cosmlib.Bech32FromEthAddress(from),
		ToAddress:   cosmlib.Bech32FromEthAddress(to),
		Amount:      sdk.Coins{sdk.NewInt64Coin(tf.Network.Config.BondDenom, amount)},
	})
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin(tf.Network.Config.BondDenom, 10)})
	txBuilder.SetGasLimit(testdata.NewTestGasLimit())
	txBuilder.SetMemo("memo")

	txFactory := clienttx.Factory{}.
		WithChainID(val.ClientCtx.ChainID).
		WithKeybase(val.ClientCtx.Keyring).
		WithTxConfig(tf.Network.Config.TxConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	err := authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
	if err != nil {
		fmt.Println("sign tx error")
		return err
	}

	txBytes, err := tf.Network.Config.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Println("tx encoding error")
		return err
	}

	res, err := val.ClientCtx.
		WithBroadcastMode(tx.BroadcastMode_BROADCAST_MODE_SYNC.String()).
		BroadcastTx(txBytes)
	if err != nil {
		fmt.Println("tx broadcast error")
		return err
	}

	if res.Code != 0 {
		return fmt.Errorf("tx not successful: %s", res.RawLog)
	}
	return nil
}
