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

package client

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"pkg.berachain.dev/stargazer/crypto/keys/ethsecp256k1"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

type Wrapper struct {
	signing.Tx
}

func NewWrapper(builder signing.Tx) *Wrapper {
	return &Wrapper{builder}
}

func (w *Wrapper) GetPubKeys() ([]cryptotypes.PubKey, error) {
	fmt.Println("GETTING CALLED")
	msgs := w.GetMsgs()
	t, ok := msgs[0].(*types.EthTransactionRequest)
	if !ok {
		return nil, errors.New("not an eth transaction")
	}

	bz, err := t.GetPubKey()
	if err != nil {
		return nil, err
	}
	return []cryptotypes.PubKey{&ethsecp256k1.PubKey{Key: bz}}, nil
}

// func (w *Wrapper) GetTx() authsigning.Tx {
// 	return w
// }

// func (w *Wrapper) FeeGranter() sdk.AccAddress {
// 	return sdk.AccAddress{}
// }

// func (w *Wrapper) FeePayer() sdk.AccAddress {
// 	return sdk.AccAddress{}
// }

// func (w *Wrapper) GetFee() sdk.Coins {
// 	return sdk.Coins{sdk.NewCoin("stargazer", sdk.NewIntFromUint64(0))}
// }

// func (w *Wrapper) GetGas() uint64 {
// 	return w.signedTx.Gas()
// }

// func (w *Wrapper) GetMemo() string {
// 	return ""
// }

// func (w *Wrapper) GetMsgs() []sdk.Msg {
// 	return []sdk.Msg{types.NewFromTransaction(w.signedTx)}
// }

// func (w *Wrapper) GetSignaturesV2() ([]signingtypes.SignatureV2, error) {
// 	return w.GetTx().GetSignaturesV2()
// }

// func (w *Wrapper) GetSigners() []sdk.AccAddress {
// 	types.Sig
// 	return []sdk.AccAddress{sdk.AccAddress()}
// }

// `NewEthTxBuilder` returns a new instance of EthTxBuilder.
func NewEthTxBuilder(signedTx *coretypes.Transaction, evmDenom string, clientCtx client.Context) (sdk.Tx, error) {
	txBuilder, ok := clientCtx.TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		return nil, errors.New("unsupported builder")
	}

	// txBuilder := NewWrapper(txb, signedTx)

	option, err := codectypes.NewAnyWithValue(&types.ExtensionOptionsEthTransaction{})
	if err != nil {
		return nil, err
	}

	txBuilder.SetExtensionOptions(option)

	txBuilder.SetGasLimit(signedTx.Gas())

	// Second, we attach the required fees to the Cosmos Tx. This is simply done,
	// by calling Cost() on the types.Transaction and setting the fee amount to that.
	fees := make(sdk.Coins, 0)
	feeAmt := sdkmath.NewIntFromBigInt(signedTx.Cost())
	if feeAmt.Sign() > 0 {
		fees = append(fees, sdk.NewCoin(evmDenom, feeAmt))
	}
	txBuilder.SetFeeAmount(fees)
	txBuilder.SetGasLimit(signedTx.Gas())

	// Thirdly, we set the nonce equal to the nonce of the transaction and also derive the PubKey
	// from the V,R,S values of the transaction. This allows us for a little trick to allow
	// ethereum transactions to work in the standard cosmos app-side mempool with no modifications.
	// Some gigabrain shit tbh.
	pk, err := PubkeyFromTx(signedTx, coretypes.LatestSignerForChainID(signedTx.ChainId()))
	if err != nil {
		return nil, err
	}
	if err = txBuilder.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			PubKey:   pk,
		},
	); err != nil {
		return nil, err
	}

	// Lastly, we inject the signed ethereum transaction as a message into the Cosmos Tx.
	if err = txBuilder.SetMsgs(types.NewFromTransaction(signedTx)); err != nil {
		return nil, err
	}

	tx := txBuilder.GetTx()
	return tx, nil

	// // Finally, we set the extension options to the builder. (ExtensionOptionsEthTransaction)
	// txBuilder.SetExtensionOptions(option)

	// // First, we attach the required fees to the Cosmos Tx. This is simply done,
	// // by calling Cost() on the types.Transaction and setting the fee amount to that.
	// fees := make(sdk.Coins, 0)
	// feeAmt := sdkmath.NewIntFromBigInt(signedTx.Cost())
	// if feeAmt.Sign() > 0 {
	// 	fees = append(fees, sdk.NewCoin(evmDenom, feeAmt))
	// }

	// txBuilder.SetFeeAmount(fees)
	// if err != nil {
	// 	return nil, errorslib.Wrap(err, "failed to set fee amount")
	// }

	// // TODO: Use SetTip() once we create the abstraction to not collect fees in "/eth"
	// // we can introduce setting the priority fee / base fee separately here.
	// // etb.SetTip(signedTx.EffectiveGasTip())
	// // etb.SetFeesAmount(signedTx.Cost()-signedTx.EffectiveGasTip())
	// // This will allow using native cosmos tipping.

	// // Secondly we set the gas limit, again extracted from ethereum transaction.
	// txBuilder.SetGasLimit(signedTx.Gas())

	// // We recover the public key from the transaction and set it in the
	// pk, err := PubkeyFromTx(signedTx, coretypes.LatestSignerForChainID(signedTx.ChainId()))
	// if err != nil {
	// 	return nil, err
	// }

	// // Thirdly, we set the nonce equal to the nonce of the transaction and also derive the PubKey
	// // from the V,R,S values of the transaction. This allows us for a little trick to allow
	// // ethereum transactions to work in the standard cosmos app-side mempool with no modifications.
	// // Some gigabrain shit tbh.
	// if err = txBuilder.SetSignatures(
	// 	signingtypes.SignatureV2{
	// 		Sequence: signedTx.Nonce(),
	// 		PubKey:   pk,
	// 	},
	// ); err != nil {
	// 	return nil, err
	// }

	// // We build a new EthereumTransaction and set give it to the builder.
	// if err = txBuilder.SetMsgs(x); err != nil {
	// 	return nil, err
	// }

	// return txBuilder.GetTx(), nil
}
