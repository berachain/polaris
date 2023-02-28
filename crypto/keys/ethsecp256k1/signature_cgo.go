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

package ethsecp256k1

import (
	"pkg.berachain.dev/stargazer/eth/crypto"
)

// `Sign` signs the provided message using the ECDSA private key. It returns an error if the
// `Sign` creates a recoverable ECDSA signature on the `secp256k1` curve over the
// provided hash of the message. The produced signature is 65 bytes
// where the last byte contains the recovery ID.
func (privKey PrivKey) Sign(digestBz []byte) ([]byte, error) {
	if len(digestBz) != crypto.DigestLength {
		digestBz = crypto.Keccak256(digestBz)
	}
	key, err := privKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	return crypto.EthSign(digestBz, key)
}

// `VerifySignature` verifies that the ECDSA public key created a given signature over
// the provided message. The signature should be in [R || S] format.
func (pubKey PubKey) VerifySignature(msg, sig []byte) bool {
	if len(msg) != crypto.DigestLength {
		msg = crypto.Keccak256(msg)
	}
	if len(sig) == crypto.SignatureLength {
		// remove recovery ID (V) if contained in the signature
		sig = sig[:len(sig)-1]
	}

	// The signature needs to be in [R || S] format when provided to `VerifySignature`.
	x := crypto.VerifySignature(pubKey.Key, msg, sig)
	return x
}
