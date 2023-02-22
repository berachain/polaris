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

package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/subtle"

	cmcrypto "github.com/cometbft/cometbft/crypto"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"pkg.berachain.dev/stargazer/eth/crypto"
)

const (
	// `PrivKeyNumBytes` defines the length of the PrivKey byte array.
	PrivKeyNumBytes = 32
	// `PubKeyNumBytes` defines the length of the PubKey byte array.
	PubKeyNumBytes = PrivKeyNumBytes + 1
	// `KeyType` is the string constant for the secp256k1 algorithm.
	KeyType = "eth_secp256k1"
)

// =====================================================================================================
// Public Key
// ====================================================================================================

// `Pubkey` is a wrapper around the Ethereum secp256k1 public key type. This wrapper conforms to
// `crypotypes.Pubkey` to allow for the use of the Ethereum secp256k1 public key type within the Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.PubKey = &EthSecp256K1PubKey{}

// `Address` returns the address of the ECDSA public key.
// The function will return an empty address if the public key is invalid.
func (pubKey EthSecp256K1PubKey) Address() cmcrypto.Address {
	pubk, err := crypto.DecompressPubkey(pubKey.Key)
	if err != nil {
		return nil
	}

	return cmcrypto.Address(crypto.PubkeyToAddress(*pubk).Bytes())
}

// `Bytes` returns the raw bytes of the ECDSA public key.
func (pubKey EthSecp256K1PubKey) Bytes() []byte {
	bz := make([]byte, len(pubKey.Key))
	copy(bz, pubKey.Key)

	return bz
}

// `Type` returns eth_secp256k1.
func (pubKey EthSecp256K1PubKey) Type() string {
	return KeyType
}

// `Equals` returns true if the pubkey type is the same and their bytes are deeply equal.
func (pubKey EthSecp256K1PubKey) Equals(other cryptotypes.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}

// `VerifySignature` verifies that the ECDSA public key created a given signature over
// the provided message. The signature should be in [R || S] format.
func (pubKey EthSecp256K1PubKey) VerifySignature(msg, sig []byte) bool {
	if len(sig) == crypto.SignatureLength {
		// remove recovery ID (V) if contained in the signature
		sig = sig[:len(sig)-1]
	}

	// The signature needs to be in [R || S] format when provided to VerifySignature.
	return crypto.VerifySignature(pubKey.Key, crypto.Keccak256Hash(msg).Bytes(), sig)
}

// =====================================================================================================
// Private Key
// ====================================================================================================

// `PrivKey` is a wrapper around the Ethereum secp256k1 private key type. This wrapper conforms to
// `crypotypes.Pubkey` to allow for the use of the Ethereum secp256k1 private key type within the Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.PrivKey = &EthSecp256K1PrivKey{}

// `GenerateKey` generates a new random private key. It returns an error upon
// failure.
func GenerateKey() (*EthSecp256K1PrivKey, error) {
	priv, err := crypto.GenerateEthKey()
	if err != nil {
		return nil, err
	}

	return &EthSecp256K1PrivKey{
		Key: crypto.FromECDSA(priv),
	}, nil
}

// `Bytes` returns the byte representation of the ECDSA Private Key.
func (privKey EthSecp256K1PrivKey) Bytes() []byte {
	bz := make([]byte, len(privKey.Key))
	copy(bz, privKey.Key)
	return bz
}

// `PubKey` returns the ECDSA private key's public key. If the privkey is not valid
// it returns a nil value.
func (privKey EthSecp256K1PrivKey) PubKey() cryptotypes.PubKey {
	ecdsaPrivKey, err := privKey.ToECDSA()
	if err != nil {
		return nil
	}

	return &EthSecp256K1PubKey{
		Key: crypto.CompressPubkey(&ecdsaPrivKey.PublicKey),
	}
}

// `Equals` returns true if two ECDSA private keys are equal and false otherwise.
func (privKey EthSecp256K1PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	return privKey.Type() == other.Type() && subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

// `Type` returns eth_secp256k1.
func (privKey EthSecp256K1PrivKey) Type() string {
	return KeyType
}

// `Sign` creates a recoverable ECDSA signature on the secp256k1 curve over the
// provided hash of the message. The produced signature is 65 bytes
// where the last byte contains the recovery ID.
func (privKey EthSecp256K1PrivKey) Sign(digestBz []byte) ([]byte, error) {
	key, err := privKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	return crypto.EthSign(digestBz, key)
}

// `ToECDSA` returns the ECDSA private key as a reference to ecdsa.PrivateKey type.
func (privKey EthSecp256K1PrivKey) ToECDSA() (*ecdsa.PrivateKey, error) {
	return crypto.ToECDSA(privKey.Bytes())
}
