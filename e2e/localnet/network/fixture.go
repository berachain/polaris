// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package localnet

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	ginkgo "github.com/onsi/ginkgo/v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	relativeKeysPath = "ethkeys/"
	genFilePath      = "genesis.json"
)

// FixtureConfig is a type defining the configuration of a TestFixture.
type FixtureConfig struct {
	path string

	baseImage     string
	localnetImage string

	containerName string
	httpAddress   string
	wsAdddress    string
	goVersion     string
}

// NewFixtureConfig creates a new FixtureConfig and infers the config directory
// absolute path from given relative path.
// requires: the configRelativePath to be relative to the file calling NewFixtureConfig.
func NewFixtureConfig(
	configRelativePath,
	baseImage,
	localnetImage,
	containerName,
	httpAddress,
	wsAdddress,
	goVersion string,
) *FixtureConfig {
	// Get file path of the caller of NewFixtureConfig.
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to get caller")
	}
	configPath, err := filepath.Abs(filepath.Join(filepath.Dir(caller), configRelativePath))
	if err != nil {
		panic(err)
	}
	return &FixtureConfig{
		path:          configPath,
		baseImage:     baseImage,
		localnetImage: localnetImage,
		containerName: containerName,
		httpAddress:   httpAddress,
		wsAdddress:    wsAdddress,
		goVersion:     goVersion,
	}
}

// TestFixture is a testing fixture that runs a single Polaris validator
// node in a Docker container.
type TestFixture struct {
	ContainerizedNode
	t       ginkgo.FullGinkgoTInterface
	keysMap map[string]*ecdsa.PrivateKey
	valAddr common.Address
}

// NewTestFixture creates a new TestFixture.
func NewTestFixture(t ginkgo.FullGinkgoTInterface, config *FixtureConfig) *TestFixture {
	tf := &TestFixture{
		t:       t,
		keysMap: make(map[string]*ecdsa.PrivateKey),
	}

	err := tf.setupTestAccounts(config)
	if err != nil {
		t.Fatal(err)
	}

	localnetImage := strings.Split(config.localnetImage, ":")
	tf.ContainerizedNode, err = NewContainerizedNode(
		localnetImage[0],
		localnetImage[1],
		config.containerName,
		config.httpAddress,
		config.wsAdddress,
		[]string{
			"GO_VERSION=" + config.goVersion,
			"BASE_IMAGE=" + config.baseImage,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	return tf
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

func (tf *TestFixture) ValAddr() common.Address {
	return tf.valAddr
}

// setupTestAccounts loads the test account private keys and validator public key.
func (tf *TestFixture) setupTestAccounts(config *FixtureConfig) error {
	// read the test account private keys from the keys directory
	keysPath := filepath.Join(config.path, relativeKeysPath)
	keyFiles, err := os.ReadDir(filepath.Clean(keysPath))
	if err != nil {
		return err
	}
	for _, keyFile := range keyFiles {
		keyFileName := keyFile.Name()

		var privKey *ecdsa.PrivateKey
		privKey, err = crypto.LoadECDSA(filepath.Join(keysPath, keyFile.Name()))
		if err != nil {
			return err
		}

		tf.keysMap[strings.Split(keyFileName, ".")[0]] = privKey
	}

	return nil
}
