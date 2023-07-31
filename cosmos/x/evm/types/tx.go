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
	"sync"

	"google.golang.org/protobuf/proto"

	"cosmossdk.io/x/tx/signing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"

	v1alpha1evm "pkg.berachain.dev/polaris/cosmos/api/polaris/evm/v1alpha1"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// WrappedEthereumTransaction defines a Cosmos SDK message for Ethereum transactions.
var _ sdk.Msg = (*WrappedEthereumTransaction)(nil)

// NewFromTransaction sets the transaction data from an `coretypes.Transaction`.
func NewFromTransaction(tx *coretypes.Transaction) *WrappedEthereumTransaction {
	bz, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return &WrappedEthereumTransaction{
		Data: bz,
	}
}

// AsTransaction extracts the transaction as an `coretypes.Transaction`.
func (etr *WrappedEthereumTransaction) AsTransaction() *coretypes.Transaction {
	tx := new(coretypes.Transaction)
	if err := tx.UnmarshalBinary(etr.Data); err != nil {
		return nil
	}
	return tx
}

// GetSignBytes returns the bytes to sign over for the transaction.
func (etr *WrappedEthereumTransaction) GetSignBytes() ([]byte, error) {
	tx := etr.AsTransaction()
	return coretypes.LatestSignerForChainID(tx.ChainId()).
		Hash(tx).Bytes(), nil
}

// GetSender extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *WrappedEthereumTransaction) GetSender() (common.Address, error) {
	tx := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(tx.ChainId())
	return signer.Sender(tx)
}

// GetSender extracts the sender address from the signature values using the latest signer for the given chainID.
func (etr *WrappedEthereumTransaction) GetSignature() ([]byte, error) {
	tx := etr.AsTransaction()
	signer := coretypes.LatestSignerForChainID(tx.ChainId())
	return signer.Signature(tx)
}

// GetGasPrice returns the gas price of the transaction.
func (etr *WrappedEthereumTransaction) ValidateBasic() error {
	// Ensure the transaction is signed properly
	tx := etr.AsTransaction()
	if tx == nil {
		return errors.New("transaction data is invalid")
	}

	// Ensure the transaction does not have a negative value.
	if tx.Value().Sign() < 0 {
		return txpool.ErrNegativeValue
	}

	// Sanity check for extremely large numbers.
	if tx.GasFeeCap().BitLen() > 256 { //nolint:gomnd // 256 bits.
		return core.ErrFeeCapVeryHigh
	}

	// Sanity check for extremely large numbers.
	if tx.GasTipCap().BitLen() > 256 { //nolint:gomnd // 256 bits.
		return core.ErrTipVeryHigh
	}

	// Ensure gasFeeCap is greater than or equal to gasTipCap.
	if tx.GasFeeCapIntCmp(tx.GasTipCap()) < 0 {
		return core.ErrTipAboveFeeCap
	}

	return nil
}

// GetAsEthTx is a helper function to get an EthTx from a sdk.Tx.
func GetAsEthTx(tx sdk.Tx) *coretypes.Transaction {
	if len(tx.GetMsgs()) == 0 {
		return nil
	}
	etr, ok := utils.GetAs[*WrappedEthereumTransaction](tx.GetMsgs()[0])
	if !ok {
		return nil
	}
	return etr.AsTransaction()
}

// ProvideEthereumTransactionGetSigners defines a custom function for
// utilizing custom signer handling for `WrappedEthereumTransaction`s.
func ProvideEthereumTransactionGetSigners() signing.CustomGetSigner {
	// Utilize a sync pool to reduce memory usage.
	txSyncPool := sync.Pool{
		New: func() any { return new(coretypes.Transaction) },
	}

	// The actual function.
	return signing.CustomGetSigner{
		MsgType: proto.MessageName(&v1alpha1evm.WrappedEthereumTransaction{}),
		Fn: func(msg proto.Message) ([][]byte, error) {
			// Pull the raw ethereum bytes from pulsar.
			ethTxData := msg.(*v1alpha1evm.WrappedEthereumTransaction).Data

			// Get a new empty Transaction.
			ethTx, ok := txSyncPool.Get().(*coretypes.Transaction)
			if !ok {
				return nil, errors.New("failed to get sync pool when getting signers")
			}

			// Fill it with the data.
			if err := ethTx.UnmarshalBinary(ethTxData); err != nil {
				return nil, err
			}

			// Extract the signer from the signature.
			signer, err := coretypes.LatestSignerForChainID(ethTx.ChainId()).Sender(ethTx)
			if err != nil {
				return nil, err
			}

			// Return the signer in the required format.
			return [][]byte{signer.Bytes()}, nil
		},
	}
}
