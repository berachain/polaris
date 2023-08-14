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
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/lib/errors"
)

// ValidateBasicDecorator will call tx.ValidateBasic and return any non-nil error.
// If ValidateBasic passes, decorator calls next AnteHandler in chain. Note,
// ValidateBasicDecorator decorator will not get executed on ReCheckTx since it
// is not dependent on application state.
type IsEvmTxDecorator struct{}

func NewIsEvmTxDecorator() IsEvmTxDecorator {
	return IsEvmTxDecorator{}
}

func (ietd IsEvmTxDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	// Only 1 transaction per message.
	if len(tx.GetMsgs()) != 1 {
		return ctx, errors.Wrap(sdkerrors.ErrUnknownRequest, "tx must contain exactly one message")
	}

	if _, ok := tx.GetMsgs()[0].(*types.WrappedEthereumTransaction); !ok {
		// For handling genesis state. (hacky af but yolo).

		// This allows for the genutil to call deliverTx on the genesis state.
		// Another option: / if ctx.BlockHeight() != 0
		// TODO: verify if this is safe condition.
		if ctx.IsCheckTx() || ctx.IsReCheckTx() {
			return ctx, errors.Wrap(sdkerrors.ErrUnknownRequest, "tx must be an Ethereum transaction")
		}
	}

	return next(ctx, tx, simulate)
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options ante.HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),                   // outermost AnteDecorator. SetUpContext must be called first
		IsEvmTxDecorator{},                                // only accept evm transactions on polaris chains.
		ante.NewValidateBasicDecorator(),                  // run validate basic on the evm transaction.
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // we can probably remove?
	}
	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
