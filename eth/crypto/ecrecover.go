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
	"errors"
	"math/big"

	"pkg.berachain.dev/stargazer/eth/common"
)

// `RecoverPubkey` recovers the public key from the given signature.
func RecoverPubkey(sighash common.Hash, rb, sb, vb *big.Int, homestead bool) ([]byte, error) {
	if vb.BitLen() > 8 { //nolint:gomnd // 8 is the length of a byte.
		return nil, errors.New("invalid signature length")
	}
	v := byte(vb.Uint64() - 27) //nolint:gomnd // 27 is the offset for the recovery id.
	if !ValidateSignatureValues(v, rb, sb, homestead) {
		return nil, errors.New("invalid signature values")
	}
	// encode the signature in uncompressed format
	r, s := rb.Bytes(), sb.Bytes()
	sig := make([]byte, SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = v
	// recover the public key from the signature
	pub, err := Ecrecover(sighash[:], sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, errors.New("invalid public key")
	}
	return pub, nil
}
