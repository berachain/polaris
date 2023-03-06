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
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// We must implement the `sdk.Msg` interface to be able to use the `sdk.Msg` type
// in the `sdk.Msg` field of the `sdk.Tx` interface.
var _ ante.GasTx = (*EthTransactionRequest)(nil)
var _ sdk.Tx = (*EthTransactionRequest)(nil)
var _ sdk.Msg = (*EthTransactionRequest)(nil)

// var _ authsigning.Tx = (*EthTransactionRequest)(nil)

// `NewFromTransaction` sets the transaction data from an `coretypes.Transaction`.
func NewFromTransaction(tx *coretypes.Transaction) *EthTransactionRequest {
	etr := new(EthTransactionRequest)
	bz, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	etr.Data = bz
	return etr
}

// `GetMsgs` returns the message(s) contained in the transaction.
func (etr *EthTransactionRequest) GetMsgs() []sdk.Msg {
	return []sdk.Msg{etr}
}

// `GetSigners` returns the address(es) that must sign over the transaction.
func (etr *EthTransactionRequest) GetSigners() []sdk.AccAddress {
	sender, err := etr.GetSender()
	if err != nil {
		panic(err)
	}

	signer := sdk.AccAddress(sender.Bytes())
	signers := []sdk.AccAddress{signer}
	return signers
}

// `AsTransaction` extracts the transaction as an `coretypes.Transaction`.
func (etr *EthTransactionRequest) AsTransaction() *coretypes.Transaction {
	t := new(coretypes.Transaction)
	err := t.UnmarshalBinary(etr.Data)
	if err != nil {
		return nil
	}
	return t
}

func (etr *EthTransactionRequest) GetSignBytes() ([]byte, error) {
	t := etr.AsTransaction()
	return coretypes.LatestSignerForChainID(t.ChainId()).
		Hash(t).Bytes(), nil
}

// `GetSender` extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *EthTransactionRequest) GetSender() (common.Address, error) {
	t := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(t.ChainId())
	return signer.Sender(t)
}

// `GetSender` extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *EthTransactionRequest) GetPubKey() ([]byte, error) {
	t := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(t.ChainId())
	return signer.PubKey(t)
}

// `GetSender` extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *EthTransactionRequest) GetSignature() ([]byte, error) {
	t := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(t.ChainId())
	return signer.Signature(t)
}

// `GetGas` returns the gas limit of the transaction.
func (etr *EthTransactionRequest) GetGas() uint64 {
	tx := etr.AsTransaction()
	if tx == nil {
		return 0
	}
	return tx.Gas()
}

// `GetGasPrice` returns the gas price of the transaction.
func (etr *EthTransactionRequest) ValidateBasic() error {
	if len(etr.Data) == 0 {
		return errors.New("transaction data cannot be empty")
	}

	if etr.AsTransaction() == nil {
		return errors.New("transaction data is invalid")
	}

	if etr.GetGas() == 0 {
		return errors.New("gas limit cannot be zero")
	}

	return nil
}
