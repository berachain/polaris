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

package hd

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	bip39 "github.com/cosmos/go-bip39"

	"pkg.berachain.dev/stargazer/crypto/keys/ethsecp256k1"
	"pkg.berachain.dev/stargazer/eth/crypto"
)

const (
	// `EthSecp256k1Type` defines the ECDSA secp256k1 used on Ethereum.
	EthSecp256k1Type = hd.PubKeyType(ethsecp256k1.KeyType)
)

var (
	// Compile-time type assertion.
	_ keyring.SignatureAlgo = EthSecp256k1
	// EthSecp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	EthSecp256k1 = ethSecp256k1Algo{}
)

// `ethSecp256k1Algo` implements the `keyring.SignatureAlgo` interface for the eth_secp256k1 algorithm.
type ethSecp256k1Algo struct{}

// `Name` returns eth_secp256k1.
func (s ethSecp256k1Algo) Name() hd.PubKeyType {
	return EthSecp256k1Type
}

// `Derive` derives and returns the eth_secp256k1 private key for the given mnemonic and HD path.
func (s ethSecp256k1Algo) Derive() hd.DeriveFn {
	return func(mnemonic, bip39Passphrase, hdPath string) ([]byte, error) {
		seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
		if err != nil {
			return nil, err
		}

		masterPriv, ch := hd.ComputeMastersFromSeed(seed)
		if len(hdPath) == 0 {
			return ECDSAify(masterPriv[:])
		}

		derivedKey, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hdPath)
		if err != nil {
			return nil, err
		}

		return ECDSAify(derivedKey[:])
	}
}

// `Generate` generates a eth_secp256k1 private key from the given bytes.
func (s ethSecp256k1Algo) Generate() hd.GenerateFn {
	return func(bz []byte) cryptotypes.PrivKey {
		bzArr := make([]byte, ethsecp256k1.PrivKeySize)
		copy(bzArr, bz)
		return &ethsecp256k1.PrivKey{
			Key: bzArr,
		}
	}
}

// `ECDSAify` converts a private key to an ECDSA private key.
func ECDSAify(key []byte) ([]byte, error) {
	// Convert the private key to an ECDSA private key.
	x, err := ethsecp256k1.PrivKey{Key: key}.ToECDSA()
	if err != nil {
		return nil, err
	}

	// Return the private key as a byte slice.
	return crypto.FromECDSA(x), nil
}
