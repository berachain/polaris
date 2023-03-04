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

package mock

import (
	"context"
	"errors"

	"pkg.berachain.dev/stargazer/eth/core"
)

type GasPluginMock struct {
	txGasUsed     uint64
	blockGasUsed  uint64
	txGasLimit    uint64
	blockGasLimit uint64
}

func NewGasPluginMock() *GasPluginMock {
	return &GasPluginMock{}
}

func (w *GasPluginMock) Prepare(context.Context) {
	w.blockGasUsed = 0
}

func (w *GasPluginMock) Reset(context.Context) {
	w.txGasUsed = 0
}

func (w *GasPluginMock) ConsumeGas(amount uint64) error {
	if w.txGasUsed+amount > w.txGasLimit {
		return errors.New("gas limit exceeded")
	}
	if w.blockGasUsed+amount > w.blockGasLimit {
		return core.ErrBlockOutOfGas
	}

	w.txGasUsed += amount
	return nil
}

func (w *GasPluginMock) CumulativeGasUsed() uint64 {
	return w.txGasUsed + w.blockGasUsed
}

func (w *GasPluginMock) TxGasRemaining() uint64 {
	return w.txGasLimit - w.txGasUsed
}

func (w *GasPluginMock) TxGasUsed() uint64 {
	return w.txGasUsed
}

func (w *GasPluginMock) TxRefundGas(amount uint64) {
	if w.txGasUsed < amount {
		w.txGasUsed = 0
	} else {
		w.txGasUsed -= amount
	}
}

func (w *GasPluginMock) SetTxGasLimit(amount uint64) error {
	w.txGasLimit = amount
	if w.txGasLimit < w.txGasUsed {
		return errors.New("gas limit is below currently consumed")
	}
	return nil
}

func (w *GasPluginMock) SetBlockGasLimit(amount uint64) {
	w.blockGasLimit = amount
}

func (w *GasPluginMock) BlockGasLimit() uint64 {
	return w.blockGasLimit
}

func (w *GasPluginMock) ConsumeGasToBlockLimit() error {
	delta := w.blockGasLimit - w.blockGasUsed
	if w.txGasUsed+delta > w.txGasLimit {
		return errors.New("tx gas limit exceeded")
	}

	w.txGasUsed += delta
	w.blockGasUsed += delta
	return nil
}
