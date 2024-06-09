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
