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

package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// TxSerializer provides an interface to serialize ethereum transactions
// to sdk.Tx's and bytes that can be used by CometBFT.
type TxSerializer interface {
	SerializeToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error)
	SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error)
}

// serializer implements TxSerializer.
type serializer struct {
	txConfig client.TxConfig
}

// NewSerializer returns a new instance of TxSerializer.
func NewSerializer(txConfig client.TxConfig) TxSerializer {
	return &serializer{
		txConfig: txConfig,
	}
}

// SerializeToSdkTx converts an ethereum transaction to a Cosmos native transaction.
func (s *serializer) SerializeToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error) {
	// TODO: do we really need to use extensions for anything? Since we
	// are using the standard ante handler stuff I don't think we actually need to.
	tx := s.txConfig.NewTxBuilder()

	// We can also retrievepothe gaslimit for the transaction from the ethereum transaction.
	tx.SetGasLimit(signedTx.Gas())

	// Thirdly, we set the nonce equal to the nonce of the transaction and also derive the PubKey
	// from the V,R,S values of the transaction. This allows us for a little trick to allow
	// ethereum transactions to work in the standard cosmos app-side mempool with no modifications.
	// Some gigabrain shit tbh.
	pkBz, err := coretypes.PubkeyFromTx(
		signedTx, coretypes.LatestSignerForChainID(signedTx.ChainId()),
	)
	if err != nil {
		return nil, err
	}

	// Create the WrappedEthereumTransaction message.
	wrappedEthTx := NewFromTransaction(signedTx)
	sig, err := wrappedEthTx.GetSignature()
	if err != nil {
		return nil, err
	}

	// Lastly, we set the signature. We can pull the sequence from the nonce of the ethereum tx.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			Data: &signingtypes.SingleSignatureData{
				// We retrieve the hash of the signed transaction from the ethereum transaction
				// objects, as this was the bytes that were signed. We pass these into the
				// SingleSignatureData as the SignModeHandler needs to know what data was signed
				// over so that it can verify the signature in the ante handler.
				Signature: sig,
			},
			PubKey: &ethsecp256k1.PubKey{Key: pkBz},
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = tx.SetMsgs(wrappedEthTx); err != nil {
		return nil, err
	}

	// Finally, we return the Cosmos Tx.
	return tx.GetTx(), nil
}

// SerializeToBytes converts an Ethereum transaction to Cosmos formatted txBytes which allows for
// it to broadcast it to CometBFT.
func (s *serializer) SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error) {
	// First, we convert the Ethereum transaction to a Cosmos transaction.
	cosmosTx, err := s.SerializeToSdkTx(signedTx)
	if err != nil {
		return nil, err
	}

	// Then we use the clientCtx.TxConfig.TxEncoder() to encode the Cosmos transaction into bytes.
	txBytes, err := s.txConfig.TxEncoder()(cosmosTx)
	if err != nil {
		return nil, err
	}

	// Finally, we return the txBytes.
	return txBytes, nil
}
