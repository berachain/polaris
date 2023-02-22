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
	"math/big"

	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
)

// NewTx returns a reference to a new Ethereum transaction message.
func NewEthereumTxRequest(
	chainID *big.Int, nonce uint64, to *common.Address, amount *big.Int,
	gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, input []byte, accesses *coretypes.AccessList,
) *EthTransactionRequest {
	return newEthereumTxRequest(chainID, nonce, to, amount, gasLimit, gasPrice, gasFeeCap, gasTipCap, input, accesses)
}

func newEthereumTxRequest(
	chainID *big.Int, nonce uint64, to *common.Address, amount *big.Int,
	gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, input []byte, accesses *coretypes.AccessList,
) *EthTransactionRequest {
	var (
		txData coretypes.TxData
	)

	switch {
	case accesses == nil:
		txData = &coretypes.LegacyTx{
			Nonce:    nonce,
			To:       to,
			Value:    amount,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Data:     input,
		}
		//nolint:govet // not a tautoology.
	case accesses != nil && gasFeeCap != nil && gasTipCap != nil:
		txData = &coretypes.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			To:        to,
			Value:     amount,
			Gas:       gasLimit,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Data:      input,
			// Accesses:  NewAccessList(accesses),
		}
	case accesses != nil:
		// coming soon
	default:
	}

	tx := coretypes.NewTx(txData)
	bz, err := tx.MarshalJSON()
	if err != nil {
		panic(err)
	}

	msg := EthTransactionRequest{Data: string(bz)}
	// msg.Hash = tx.Hash()
	return &msg
}
