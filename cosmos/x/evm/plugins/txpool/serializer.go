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
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	evmante "pkg.berachain.dev/polaris/cosmos/x/evm/ante"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// SerializeToSdkTx converts an ethereum transaction to a Cosmos native transaction.
func SerializeToSdkTx(
	clientCtx client.Context, signedTx *coretypes.Transaction,
) (sdk.Tx, error) {
	// TODO: do we really need to use extensions for anything? Since we
	// are using the standard ante handler stuff I don't think we actually need to.
	tx := clientCtx.TxConfig.NewTxBuilder()

	// We can also retrieve the gaslimit for the transaction from the ethereum transaction.
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

	pk := ethsecp256k1.PubKey{Key: pkBz}

	// Create the WrappedEthereumTransaction message.
	wrappedEthTx := types.NewFromTransaction(signedTx)

	// fuck cosmos: https://github.com/cosmos/cosmos-sdk/pull/16340/files
	wrappedEthTx.HackyFixCauseCosmos = sdk.MustBech32ifyAddressBytes(
		sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	sig, err := wrappedEthTx.GetSignature()
	if err != nil {
		return nil, err
	}

	// Lastly, we set the signature. We can pull the sequence from the nonce of the ethereum tx.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			Data: &signingtypes.SingleSignatureData{
				// TODO: this is ghetto af.
				SignMode: signingtypes.SignMode(int32(evmante.SignMode_SIGN_MODE_ETHEREUM)),
				// We retrieve the hash of the signed transaction from the ethereum transaction
				// objects, as this was the bytes that were signed. We pass these into the
				// SingleSignatureData as the SignModeHandler needs to know what data was signed
				// over so that it can verify the signature in the ante handler.
				Signature: sig,
			},
			PubKey: &pk,
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
func SerializeToBytes(
	clientCtx client.Context, signedTx *coretypes.Transaction,
) ([]byte, error) {
	// First, we convert the Ethereum transaction to a Cosmos transaction.
	cosmosTx, err := SerializeToSdkTx(clientCtx, signedTx)
	if err != nil {
		return nil, err
	}

	// Then we use the clientCtx.TxConfig.TxEncoder() to encode the Cosmos transaction into bytes.
	txBytes, err := clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		return nil, err
	}

	// Finally, we return the txBytes.
	return txBytes, nil
}
