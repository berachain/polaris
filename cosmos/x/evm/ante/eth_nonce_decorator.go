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
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

var (
	_ sdk.AnteDecorator = NonceDecorator{}
)

type (
	AccountKeeper interface {
		GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	}

	// NonceDecorator decouples the nonce verification logic from the signature verification logic seen
	// in the SigVerificationDecorator. This decorator will enforce that transactions are ordered
	// respecting the upcoming nonce for an account.
	NonceDecorator struct {
		ak AccountKeeper
	}
)

func NewNonceDecorator(ak AccountKeeper) NonceDecorator {
	return NonceDecorator{
		ak: ak,
	}
}

func (nonceDecorator NonceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, fmt.Errorf("invalid transaction type")
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, fmt.Errorf("failed to get signatures: %w", err)
	}

	signerAddrs := sigTx.GetSigners()
	if len(sigs) != len(signerAddrs) {
		return ctx, fmt.Errorf("invalid number of signer; expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc := nonceDecorator.ak.GetAccount(ctx, signerAddrs[i])
		if acc == nil {
			return ctx, fmt.Errorf("account %s does not exist", signerAddrs[i])
		}

		// NOTE: Is this problematic for accounts that haven't been created?
		pubKey := acc.GetPubKey()
		if pubKey == nil {
			return ctx, fmt.Errorf("account %s has no pubkey", signerAddrs[i])
		}

		if sig.Sequence != acc.GetSequence() {
			return ctx, fmt.Errorf("invalid sequence number: got %d, expected %d", sig.Sequence, acc.GetSequence())
		}
	}

	return next(ctx, tx, simulate)
}
