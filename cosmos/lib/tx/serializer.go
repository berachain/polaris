// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package libtx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// TxSerializer provides an interface to serialize ethereum transactions
// to sdk.Tx's and bytes that can be used by CometBFT.
type TxSerializer[I any] interface {
	ToSdkTx(input I, gasLimit uint64) (sdk.Tx, error)
	ToSdkTxBytes(input I, gasLimit uint64) ([]byte, error)
}

type serializer[I any, O sdk.Msg] struct {
	txConfig client.TxConfig
	wrapFn   func(I) (O, error)
}

// NewSerializer returns a new instance of TxSerializer.
func NewSerializer[I any, O sdk.Msg](
	txConfig client.TxConfig, wrapFn func(I) (O, error),
) TxSerializer[I] {
	return &serializer[I, O]{
		txConfig: txConfig,
		wrapFn:   wrapFn,
	}
}

func (s *serializer[I, O]) ToSdkTx(input I, gasLimit uint64) (sdk.Tx, error) {
	var err error
	// TODO: do we really need to use extensions for anything? Since we
	// are using the standard ante handler stuff I don't think we actually need to.
	tx := s.txConfig.NewTxBuilder()

	// Set the tx gas limit to the block gas limit in the payload
	tx.SetGasLimit(gasLimit)

	wrapped, err := s.wrapFn(input)
	if err != nil {
		return nil, err
	}
	// TODO: figure out if we can ignore setting sigs.
	if err = tx.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: 0,
			Data: &signingtypes.SingleSignatureData{
				Signature: []byte{0x01},
			},
			PubKey: &secp256k1.PubKey{Key: []byte{0x01}},
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = tx.SetMsgs(wrapped); err != nil {
		return nil, err
	}

	// Finally, we return the Cosmos Tx.
	return tx.GetTx(), nil
}

// SerializeToBytes converts an Ethereum transaction to Cosmos formatted
// txBytes which allows for it to broadcast it to CometBFT.
func (s *serializer[I, O]) ToSdkTxBytes(
	input I, gasLimit uint64,
) ([]byte, error) {
	// First, we convert the Ethereum transaction to a Cosmos transaction.
	cosmosTx, err := s.ToSdkTx(input, gasLimit)
	if err != nil {
		return nil, err
	}

	// Then we use the clientCtx.TxConfig.TxEncoder() to encode the Cosmos transaction into bytes.
	return s.txConfig.TxEncoder()(cosmosTx)
}
