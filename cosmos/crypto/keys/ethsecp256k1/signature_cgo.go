// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
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

package ethsecp256k1

import (
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Sign signs the provided message using the ECDSA private key. It returns an error if the
// Sign creates a recoverable ECDSA signature on the `secp256k1` curve over the
// provided hash of the message. The produced signature is 65 bytes
// where the last byte contains the recovery ID.
func (privKey PrivKey) Sign(digestBz []byte) ([]byte, error) {
	// We hash the provided input since EthSign expects a 32byte hash.
	if len(digestBz) != ethcrypto.DigestLength {
		digestBz = ethcrypto.Keccak256(digestBz)
	}

	key, err := privKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	return ethcrypto.Sign(digestBz, key)
}

// VerifySignature verifies that the ECDSA public key created a given signature over
// the provided message. The signature should be in [R || S] format.
func (pubKey PubKey) VerifySignature(msg, sig []byte) bool {
	// This is a little hacky, but in order to work around the fact that the Cosmos-SDK typically
	// does not hash messages, we have to accept an unhashed message and hash it.
	// NOTE: this function will not work correctly if a msg of length 32 is provided, that is actually
	// the hash of the message that was signed.
	if len(msg) != ethcrypto.DigestLength {
		msg = ethcrypto.Keccak256(msg)
	}

	// The signature length must be correct.
	if len(sig) == ethcrypto.SignatureLength {
		// remove recovery ID (V) if contained in the signature
		sig = sig[:len(sig)-1]
	}

	// The signature needs to be in [R || S] format when provided to `VerifySignature`.
	return ethcrypto.VerifySignature(pubKey.Key, msg, sig)
}
