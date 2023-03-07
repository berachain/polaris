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

package txpool

import (
	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// `PubkeyFromTx` returns the public key of the signer of the transaction.
func PubkeyFromTx(signedTx *coretypes.Transaction, signer coretypes.Signer) (*ethsecp256k1.PubKey, error) {
	// `signer.PubKey` returns the uncompressed public key.
	uncompressed, err := signer.PubKey(signedTx)
	if err != nil {
		return &ethsecp256k1.PubKey{}, err
	}

	// We marshal it to a *ecdsa.PublicKey.
	pubKey, err := crypto.UnmarshalPubkey(uncompressed)
	if err != nil {
		return &ethsecp256k1.PubKey{}, err
	}

	// Then we can compress it to adhere to the required format.
	return &ethsecp256k1.PubKey{Key: crypto.CompressPubkey(pubKey)}, nil
}
