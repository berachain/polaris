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
	"errors"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

// `Serializer` defines that interface that allows serializes an Ethereum transactions
// to Cosmos native transaction types / formats.
type Serializer interface {
	// `Serialize` serializes the given transaction into a byte slice.
	Serialize(tx *coretypes.Transaction) ([]byte, error)
	// 'SerializeToSdkTx converts an ethereum transaction to a Cosmos transaction.
	SerializeToSdkTx(tx *coretypes.Transaction) (sdk.Tx, error)
}

// `serializer` represents the transaction pool plugin.
type serializer struct {
	clientCtx client.Context
}

// `NewSerializer` returns a new `Serializer`.
func NewSerializer(clientCtx client.Context) Serializer {
	return &serializer{
		clientCtx: clientCtx,
	}
}

// `Serialize` converts an Ethereum transaction to txBytes which allows for it to
// broadcast it to CometBFT.
func (s *serializer) Serialize(signedTx *coretypes.Transaction) ([]byte, error) {
	cosmosTx, err := s.SerializeToSdkTx(signedTx)
	if err != nil {
		return nil, err
	}

	txBytes, err := s.clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		// b.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return nil, err
	}
	return txBytes, nil
}

// `BuildCosmosTxFromEthTx` converts an ethereum transaction to a Cosmos
// transaction.
func (s *serializer) SerializeToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error) {
	// TODO: do we really need to use extensions for anything? Since we
	// are using the standard ante handler stuff I don't think we actually need to.
	tx, ok := s.clientCtx.TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		return nil, errors.New("unsupported builder")
	}

	// First, we attach the required fees to the Cosmos Tx. This is simply done,
	// by calling Cost() on the types.Transaction and setting the fee amount to that.
	option, err := codectypes.NewAnyWithValue(&types.ExtensionOptionsEthTransaction{})
	if err != nil {
		return nil, err
	}
	tx.SetExtensionOptions(option)

	// Second, we attach the required fees to the Cosmos Tx. This is simply done,
	// by calling Cost() on the types.Transaction and setting the fee amount to that
	fees := make(sdk.Coins, 0)
	feeAmt := sdkmath.NewIntFromBigInt(signedTx.Cost())
	if feeAmt.Sign() > 0 {
		// TODO: properly get evm denomination.
		fees = append(fees, sdk.NewCoin("abera", feeAmt))
	}
	tx.SetFeeAmount(fees)

	// We can also retrieve the gaslimit for the transaction from the ethereum transaction.
	tx.SetGasLimit(signedTx.Gas())

	// Thirdly, we set the nonce equal to the nonce of the transaction and also derive the PubKey
	// from the V,R,S values of the transaction. This allows us for a little trick to allow
	// ethereum transactions to work in the standard cosmos app-side mempool with no modifications.
	// Some gigabrain shit tbh.
	pk, err := PubkeyFromTx(signedTx, coretypes.LatestSignerForChainID(signedTx.ChainId()))
	if err != nil {
		return nil, err
	}

	// Lastly, we set the signature. We can pull the sequence from the nonce of the ethereum tx.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			Data: &signingtypes.SingleSignatureData{
				// TODO: this will fail, need to define custom signmode.
				SignMode: signingtypes.SignMode_SIGN_MODE_DIRECT,
				// We retrieve the hash of the signed transaction from the ethereum transaction
				// objects, as this was the bytes that were signed. We pass these into the
				// SingleSignatureData as the SignModeHandler needs to know what data was signed
				// over so that it can verify the signature in the ante handler.
				Signature: coretypes.LatestSignerForChainID(signedTx.ChainId()).
					Hash(signedTx).Bytes(),
			},
			PubKey: pk,
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = tx.SetMsgs(types.NewFromTransaction(signedTx)); err != nil {
		return nil, err
	}

	// Finally, we return the Cosmos Tx
	return tx.GetTx(), nil
}
