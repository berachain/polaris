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
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"pkg.berachain.dev/stargazer/lib/errors"
)

// `SetAnteHandler` sets the required ante handler for a Stargazer Cosmos SDK Chain.
func SetAnteHandler(
	ak ante.AccountKeeper,
	bk authtypes.BankKeeper,
	fgk ante.FeegrantKeeper,
	txCfg client.TxConfig,
) func(bApp *baseapp.BaseApp) {
	return func(bApp *baseapp.BaseApp) {
		opt := ante.HandlerOptions{
			AccountKeeper:   ak,
			BankKeeper:      bk,
			SignModeHandler: txCfg.SignModeHandler(),
			FeegrantKeeper:  fgk,
			SigGasConsumer:  SigVerificationGasConsumer,
		}
		ch, _ := ante.NewAnteHandler(
			opt,
		)
		bApp.SetAnteHandler(
			ch,
		)
	}
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options ante.HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		// ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper), // in geth intrinsic
		// ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper,
		// options.FeegrantKeeper, options.TxFeeChecker), // in geth state transition
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator m
		// ust be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),

		// SIG VERIFY BUG: https://github.com/berachain/stargazer/issues/354, possible solution is to
		// check signatures in the application side mempool and kick them out. The only downside to this,
		// is that is that we are letting transactions with bad sigs into the mempool and we need to potentially
		// worry about spam.
		// ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		// ante.NewIncrementSequenceDecorator(options.AccountKeeper), // in state tranisiton
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
