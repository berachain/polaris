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

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"pkg.berachain.dev/stargazer/crypto"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	ethcrypto "pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

// `EthTxBuilder` is an interface that wraps the `BuildTx` method and an
// ExtensionOptionsTxBuilder.
type EthTxBuilder interface {
	authtx.ExtensionOptionsTxBuilder
	BuildTx(signedTx *coretypes.Transaction, evmDenom string) (signing.Tx, error)
}

// `ethTxBuilder` implements `EthTxBuilder` interface.
type ethTxBuilder struct {
	authtx.ExtensionOptionsTxBuilder
	option *codectypes.Any
}

// `NewEthTxBuilder` returns a new instance of EthTxBuilder.
func NewEthTxBuilder(clientCtx client.Context) (EthTxBuilder, error) {
	// All Eth transactions use the ExtensionOptionsEthTransaction extension option.
	option, err := codectypes.NewAnyWithValue(&types.ExtensionOptionsEthTransaction{})
	if err != nil {
		return nil, err
	}

	// We use the clientCtx.TxConfig to create a new TxBuilder.
	txBuilder, ok := clientCtx.TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		return nil, errors.New("clientCtx.TxConfig.NewTxBuilder returns unsupported builder")
	}

	return &ethTxBuilder{
		ExtensionOptionsTxBuilder: txBuilder,
		option:                    option,
	}, nil
}

// BuildTx builds the canonical cosmos tx from ethereum msg.
func (etb *ethTxBuilder) BuildTx(
	signedTx *coretypes.Transaction, evmDenom string,
) (signing.Tx, error) {
	// First, we attach the required fees to the Cosmos Tx. This is simply done,
	// by calling Cost() on the types.Transaction and setting the fee amount to that.
	fees := make(sdk.Coins, 0)
	feeAmt := sdkmath.NewIntFromBigInt(signedTx.Cost())
	if feeAmt.Sign() > 0 {
		fees = append(fees, sdk.NewCoin(evmDenom, feeAmt))
	}
	etb.SetFeeAmount(fees)

	// TODO: Use SetTip() once we create the abstraction to not collect fees in "/eth"
	// we can introduce setting the priority fee / base fee separately here.
	// etb.SetTip(signedTx.EffectiveGasTip())
	// etb.SetFeesAmount(signedTx.Cost()-signedTx.EffectiveGasTip())
	// This will allow using native cosmos tipping.

	// Secondly we set the gas limit, again extracted from ethereum transaction.
	etb.SetGasLimit(signedTx.Gas())

	// We recover the public key from the transaction and set it in the
	pk, err := PubkeyFromTx(signedTx, coretypes.LatestSignerForChainID(signedTx.ChainId()))
	if err != nil {
		return nil, err
	}

	// Thirdly, we set the nonce equal to the nonce of the transaction and also derive the PubKey
	// from the V,R,S values of the transaction. This allows us for a little trick to allow
	// ethereum transactions to work in the standard cosmos app-side mempool with no modifications.
	// Some gigabrain shit tbh.
	if err = etb.SetSignatures(
		signingtypes.SignatureV2{
			Sequence: signedTx.Nonce(),
			PubKey:   &pk,
		},
	); err != nil {
		return nil, err
	}

	// We build a new EthereumTransaction and set give it to the builder.
	if err = etb.SetMsgs(types.NewFromTransaction(signedTx)); err != nil {
		return nil, err
	}

	// Finally, we set the extension options to the builder. (ExtensionOptionsEthTransaction)
	etb.SetExtensionOptions(etb.option)
	return etb.GetTx(), nil
}

// `PubkeyFromTx` returns the public key of the signer of the transaction.
func PubkeyFromTx(signedTx *coretypes.Transaction, signer coretypes.Signer) (crypto.EthSecp256K1PubKey, error) {
	hash := signer.Hash(signedTx)
	v, r, s := signedTx.RawSignatureValues()
	pk, err := ethcrypto.RecoverPubkey(hash, r, s, v, true)
	if err != nil {
		return crypto.EthSecp256K1PubKey{}, err
	}
	return crypto.EthSecp256K1PubKey{Key: pk}, nil
}
