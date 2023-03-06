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

package ante

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"pkg.berachain.dev/polaris/crypto/keys/ethsecp256k1"
)

const (
	// `secp256k1GasCostEIP155` is the cost of a secp256k1 signature verification
	// with the `EIP155` replay protection.
	secp256k1GasCostEIP155 = 21000
)

// `SigVerificationGasConsumer` is a custom gas consumer for Cosmos-SDK chains that
// use Ethereum secp256k1 signatures.
func SigVerificationGasConsumer(
	meter storetypes.GasMeter, sig signing.SignatureV2, params authtypes.Params,
) error {
	// Then check to see if the pubkey is a secp256k1 pubkey
	switch pubkey := sig.PubKey.(type) {
	case *ethsecp256k1.PubKey:
		meter.ConsumeGas(secp256k1GasCostEIP155, "ante verify: secp256k1")
		return nil
	default:
		// If we are using any other key type, we will use the default gas consumer.
		if err := ante.DefaultSigVerificationGasConsumer(meter, sig, params); err == nil {
			return fmt.Errorf("unsupported pubkey type: %T", pubkey)
		}
	}
	return nil
}
