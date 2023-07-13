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
	"errors"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	ginkgo "github.com/onsi/ginkgo/v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/crypto"
)

const relativeKeysPath = "../config/ethkeys/"

// TestFixture is a testing fixture that runs a single Polaris validator node in a Docker container.
type TestFixture struct {
	ContainerizedNode
	t       ginkgo.FullGinkgoTInterface
	keysMap map[string]*ecdsa.PrivateKey
}

// NewTestFixture creates a new TestFixture.
func NewTestFixture(t ginkgo.FullGinkgoTInterface) *TestFixture {
	// load all the test accounts
	keysMap := make(map[string]*ecdsa.PrivateKey)
	if err := setupTestAccounts(keysMap); err != nil {
		t.Fatal(err)
	}

	// start the docker container
	containerizedNode, err := NewContainerizedNode(
		"localnet",
		"latest",
		"goodcontainer",
		"8545/tcp",
		"8546/tcp",
		[]string{
			"GO_VERSION=1.20.4",
			"BASE_IMAGE=polard/base:v0.0.0",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// build and return the TestFixture
	return &TestFixture{
		ContainerizedNode: containerizedNode,
		t:                 t,
		keysMap:           keysMap,
	}
}

func (tf *TestFixture) Teardown() error {
	if err := tf.Stop(); err != nil {
		return err
	}
	return tf.Remove()
}

// GenerateTransactOpts generates a new transaction options object for a key by it's name.
func (tf *TestFixture) GenerateTransactOpts(name string) *bind.TransactOpts {
	// Get the nonce from the RPC.
	nonce, err := tf.EthClient().PendingNonceAt(context.Background(), tf.Address(name))
	if err != nil {
		tf.t.Fatal(err)
	}

	// Get the ChainID from the RPC.
	chainID, err := tf.EthClient().ChainID(context.Background())
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
	return tf.keysMap[name]
}

func (tf *TestFixture) Address(name string) common.Address {
	privKey, found := tf.keysMap[name]
	if !found {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(privKey.PublicKey)
}

// setupTestAccounts loads all the test account private keys from the keys directory.
func setupTestAccounts(keysMap map[string]*ecdsa.PrivateKey) error {
	// get filePath of fixture.go
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		return errors.New("unable to get key files")
	}

	keysPath := filepath.Dir(filePath) + "/" + relativeKeysPath
	keyFiles, err := os.ReadDir(keysPath)
	if err != nil {
		return err
	}

	for _, keyFile := range keyFiles {
		var privKey *ecdsa.PrivateKey
		privKey, err = crypto.LoadECDSA(keysPath + keyFile.Name())
		if err != nil {
			return err
		}

		keysMap[strings.Split(keyFile.Name(), ".")[0]] = privKey
	}

	return nil
}
