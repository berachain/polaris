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

package ethsecp256k1

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/subtle"

	cmcrypto "github.com/cometbft/cometbft/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const (
	// PrivKeySize defines the length of the PrivKey byte array.
	PrivKeySize = 32
	// PubKeySize defines the length of the PubKey byte array.
	PubKeySize = PrivKeySize + 1
	// KeyType is the string constant for the secp256k1 algorithm.
	KeyType = "eth_secp256k1"
)

// ===============================================================================================
// Private Key
// ===============================================================================================

// PrivKey is a wrapper around the Ethereum secp256k1 private key type. This wrapper conforms to
// crypotypes.Pubkey to allow for the use of the Ethereum secp256k1 private key type within the
// Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.PrivKey = &PrivKey{}

// Bytes returns the byte representation of the ECDSA Private Key.
func (privKey PrivKey) Bytes() []byte {
	return privKey.Key
}

// PubKey returns the ECDSA private key's public key. If the privkey is not valid
// it returns a nil value.
func (privKey PrivKey) PubKey() cryptotypes.PubKey {
	ecdsaPrivKey, err := privKey.ToECDSA()
	if err != nil {
		return nil
	}

	return &PubKey{
		Key: ethcrypto.CompressPubkey(&ecdsaPrivKey.PublicKey),
	}
}

// Equals returns true if two ECDSA private keys are equal and false otherwise.
func (privKey PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	return privKey.Type() == other.Type() &&
		subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

// Type returns eth_secp256k1.
func (privKey PrivKey) Type() string {
	return KeyType
}

// GenPrivKey generates a new random private key. It returns an error upon
// failure.
func GenPrivKey() (*PrivKey, error) {
	priv, err := ethcrypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &PrivKey{
		Key: ethcrypto.FromECDSA(priv),
	}, nil
}

// ToECDSA returns the ECDSA private key as a reference to ecdsa.PrivateKey type.
func (privKey PrivKey) ToECDSA() (*ecdsa.PrivateKey, error) {
	return ethcrypto.ToECDSA(privKey.Bytes())
}

// ===============================================================================================
// Public Key
// ===============================================================================================

// Pubkey is a wrapper around the Ethereum secp256k1 public key type. This wrapper conforms to
// crypotypes.Pubkey to allow for the use of the Ethereum secp256k1 public key type within the
// Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.PubKey = &PubKey{}

// Address returns the address of the ECDSA public key.
// The function will return an empty address if the public key is invalid.
func (pubKey PubKey) Address() cmcrypto.Address {
	key, err := ethcrypto.DecompressPubkey(pubKey.Key)
	if err != nil {
		return nil
	}

	return cmcrypto.Address(ethcrypto.PubkeyToAddress(*key).Bytes())
}

// Bytes returns the pubkey byte format.
func (pubKey *PubKey) Bytes() []byte {
	return pubKey.Key
}

// Type returns eth_secp256k1.
func (pubKey *PubKey) Type() string {
	return KeyType
}

// Equals returns true if the pubkey type is the same and their bytes are deeply equal.
func (pubKey PubKey) Equals(other cryptotypes.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}
