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
	"fmt"

	"github.com/ethereum/go-ethereum/beacon/engine"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// WrapTx sets the transaction data from an `coretypes.Transaction`.
func WrapTx(tx *ethtypes.Transaction) (*WrappedEthereumTransaction, error) {
	bz, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to wrap transaction: %w", err)
	}

	return &WrappedEthereumTransaction{
		Data: bz,
	}, nil
}

// Unwrap extracts the transaction as an `coretypes.Transaction`.
func (etr *WrappedEthereumTransaction) Unwrap() *ethtypes.Transaction {
	tx := new(ethtypes.Transaction)
	if err := tx.UnmarshalBinary(etr.Data); err != nil {
		return nil
	}
	return tx
}

// WrapPayload sets the payload data from an `engine.ExecutionPayloadEnvelope`.
func WrapPayload(envelope *engine.ExecutionPayloadEnvelope) (*WrappedPayloadEnvelope, error) {
	bz, err := envelope.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to wrap payload: %w", err)
	}

	return &WrappedPayloadEnvelope{
		Data: bz,
	}, nil
}

// AsPayload extracts the payload as an `engine.ExecutionPayloadEnvelope`.
func (wpe *WrappedPayloadEnvelope) UnwrapPayload() *engine.ExecutionPayloadEnvelope {
	payload := new(engine.ExecutionPayloadEnvelope)
	if err := payload.UnmarshalJSON(wpe.Data); err != nil {
		return nil
	}
	return payload
}
