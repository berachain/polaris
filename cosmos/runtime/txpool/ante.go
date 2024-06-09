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

package txpool

import (
	"errors"
	"time"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// AnteHandle implements sdk.AnteHandler.
// It is used to determine whether transactions should be ejected
// from the comet mempool.
func (m *Mempool) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	// The transaction put into this function by CheckTx
	// is a transaction from CometBFT mempool.
	telemetry.IncrCounter(float32(1), MetricKeyCometPoolTxs)
	msgs := tx.GetMsgs()

	// TODO: Record the time it takes to build a payload.

	// We only want to eject transactions from comet on recheck.
	if ctx.ExecMode() == sdk.ExecModeCheck || ctx.ExecMode() == sdk.ExecModeReCheck {
		if wet, ok := utils.GetAs[*types.WrappedEthereumTransaction](msgs[0]); ok {
			ethTx := wet.Unwrap()
			if shouldEject := m.shouldEjectFromCometMempool(
				ctx.BlockTime().Unix(), ethTx,
			); shouldEject {
				telemetry.IncrCounter(float32(1), MetricKeyAnteEjectedTxs)
				return ctx, errors.New("eject from comet mempool")
			}
		}
	}
	return next(ctx, tx, simulate)
}

// shouldEject returns true if the transaction should be ejected from the CometBFT mempool.
func (m *Mempool) shouldEjectFromCometMempool(
	currentTime int64, tx *ethtypes.Transaction,
) bool {
	defer telemetry.MeasureSince(time.Now(), MetricKeyTimeShouldEject)
	if tx == nil {
		return false
	}

	// First check things that are stateless.
	if m.validateStateless(tx, currentTime) {
		return true
	}

	// Then check for things that are stateful.
	return m.validateStateful(tx)
}

// validateStateless returns whether the tx of the given hash is stateless.
func (m *Mempool) validateStateless(tx *ethtypes.Transaction, currentTime int64) bool {
	txHash := tx.Hash()
	// 1. If the transaction has been in the mempool for longer than the configured timeout.
	// 2. If the transaction's gas params are less than or equal to the configured limit.
	expired := currentTime-m.crc.TimeFirstSeen(txHash) > m.lifetime
	priceLeLimit := tx.GasPrice().Cmp(m.priceLimit) <= 0

	if expired {
		telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectExpiredTx)
	}
	if priceLeLimit {
		telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectPriceLimit)
	}

	return expired || priceLeLimit
}

// includedCanonicalChain returns whether the tx of the given hash is included in the canonical
// Eth chain.
func (m *Mempool) validateStateful(tx *ethtypes.Transaction) bool {
	// // 1. If the transaction has been included in a block.
	// signer := ethtypes.LatestSignerForChainID(m.chainConfig.ChainID)
	// if _, err := ethtypes.Sender(signer, tx); err != nil {
	// 	return true
	// }

	// tx.Nonce() <
	included := m.chain.GetTransactionLookup(tx.Hash()) != nil
	telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectInclusion)
	return included
}
