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

package tx

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"pkg.berachain.dev/stargazer/x/evm/types"
)

// `SignMode_SIGN_MODE_ETHEREUM` defines the sign mode for Ethereum transactions.
//
//nolint:revive,stylecheck // underscores used for sign modes.
const SignMode_SIGN_MODE_ETHEREUM signingtypes.SignMode = 42069

// `CustomSignModeHandlers` returns the custom sign mode handlers for the EVM module.
func CustomSignModeHandlers() []signing.SignModeHandler {
	return []signing.SignModeHandler{
		SignModeEthTxHandler{},
	}
}

var _ signing.SignModeHandlerWithContext = (*SignModeEthTxHandler)(nil)

// `SignModeEthTx` defines the sign mode for Ethereum transactions.
type SignModeEthTxHandler struct{}

// `Modes` returns the sign modes for the EVM module.
func (s SignModeEthTxHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{s.DefaultMode()}
}

// `DefaultMode` returns the default sign mode for the EVM module.
func (s SignModeEthTxHandler) DefaultMode() signingtypes.SignMode {
	return SignMode_SIGN_MODE_ETHEREUM
}

// `GetSignBytes` returns the sign bytes for the given sign mode and transaction.
func (s SignModeEthTxHandler) GetSignBytes(
	mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx,
) ([]byte, error) {
	ethTx, ok := tx.GetMsgs()[0].(*types.EthTransactionRequest)
	if !ok {
		return nil, errors.New("expected EthTransactionRequest")
	}
	return ethTx.GetSignBytes()
}

// `GetSignBytes` returns the sign bytes for the given sign mode and transaction.
func (s SignModeEthTxHandler) GetSignBytesWithContext(_ context.Context,
	mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	return s.GetSignBytes(mode, data, tx)
}
