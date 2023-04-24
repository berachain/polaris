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

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"

	evmapi "pkg.berachain.dev/polaris/cosmos/api/polaris/evm/v1alpha1"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// SignMode_SIGN_MODE_ETHEREUM defines the sign mode for Ethereum transactions.
//
//nolint:revive,stylecheck // underscores used for sign modes.
const SignMode_SIGN_MODE_ETHEREUM signingv1beta1.SignMode = 42069

var _ txsigning.SignModeHandler = (*SignModeEthTxHandler)(nil)

// SignModeEthTx defines the sign mode for Ethereum transactions.
type SignModeEthTxHandler struct{}

// Mode implements txsigning.SignModeHandler.
func (s SignModeEthTxHandler) Mode() signingv1beta1.SignMode {
	return SignMode_SIGN_MODE_ETHEREUM
}

// TODO CONVERT ALL TXS to Pulsar (this is some hood cast shit rn)
//
// GetSignBytes implements txsigning.SignModeHandler.
func (s SignModeEthTxHandler) GetSignBytes(ctx context.Context,
	data txsigning.SignerData, txData txsigning.TxData) ([]byte, error) {
	ethTx := &evmapi.EthTransactionRequest{}
	if err := txData.Body.Messages[0].UnmarshalTo(ethTx); err != nil {
		return nil, err
	}

	return (&types.EthTransactionRequest{Data: ethTx.GetData()}).GetSignBytes()
}
