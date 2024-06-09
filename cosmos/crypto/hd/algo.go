// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package hd

import (
	"github.com/berachain/polaris/cosmos/crypto/keys/ethsecp256k1"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

const (
	// EthSecp256k1Type defines the ECDSA secp256k1 used on Ethereum.
	EthSecp256k1Type = hd.PubKeyType(ethsecp256k1.KeyType)
)

var (
	// Compile-time type assertion.
	_ keyring.SignatureAlgo = EthSecp256k1
	// EthSecp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	EthSecp256k1 = ethSecp256k1Algo{}
)

// ethSecp256k1Algo implements the `keyring.SignatureAlgo` interface for the
// eth_secp256k1 algorithm.
type ethSecp256k1Algo struct{}

// Name returns eth_secp256k1.
func (s ethSecp256k1Algo) Name() hd.PubKeyType {
	return EthSecp256k1Type
}

// Derive derives and returns the eth_secp256k1 private key for the given mnemonic
// and HD path.
func (s ethSecp256k1Algo) Derive() hd.DeriveFn {
	return hd.Secp256k1.Derive()
}

// Generate generates a eth_secp256k1 private key from the given bytes.
func (s ethSecp256k1Algo) Generate() hd.GenerateFn {
	return func(bz []byte) cryptotypes.PrivKey {
		bzArr := make([]byte, ethsecp256k1.PrivKeySize)
		copy(bzArr, bz)
		return &ethsecp256k1.PrivKey{
			Key: bzArr,
		}
	}
}
