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
