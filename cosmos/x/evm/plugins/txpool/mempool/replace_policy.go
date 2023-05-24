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

package mempool

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// NewEthTxReplacement serves as a tx replacement policy. It is called when a new transaction is
// added to the mempool and a transaction with the same nonce already exists. It returns true if
// the new transaction should replace the existing transaction.
func NewEthTxReplacement[C comparable](priceBump uint64) func(op, np C, oTx, nTx sdk.Tx) bool {
	a := big.NewInt(100 + int64(priceBump)) //nolint:gomnd // a = 100 + priceBump
	b := big.NewInt(100)                    //nolint:gomnd // b = 100

	// Source: https://github.com/ethereum/go-ethereum/blob/9231770811cda0473a7fa4e2bccc95bf62aae634/core/txpool/list.go#L284
	//
	//nolint:lll // url.
	return func(op, np C, oTx, nTx sdk.Tx) bool {
		// Convert the transactions to Ethereum transactions.
		oldEthTx := evmtypes.GetAsEthTx(oTx)
		newEthTx := evmtypes.GetAsEthTx(nTx)
		if oldEthTx == nil || newEthTx == nil ||
			oldEthTx.GasFeeCapCmp(newEthTx) >= 0 || oldEthTx.GasTipCapCmp(newEthTx) >= 0 {
			return false
		}

		// thresholdFeeCap = oldFC * (100 + priceBump) / 100 = (oldFC * a) / b
		aFeeCap := new(big.Int).Mul(a, oldEthTx.GasFeeCap())
		thresholdFeeCap := aFeeCap.Div(aFeeCap, b)

		// thresholdTip = oldTip * (100 + priceBump) / 100 = (oldTip * a) / b
		aTip := new(big.Int).Mul(a, oldEthTx.GasTipCap())
		thresholdTip := aTip.Div(aTip, b)

		// We have to ensure that both the new fee cap and tip are higher than the old ones as well
		// as checking the percentage threshold to ensure that this is accurate for low (Wei-level)
		// gas price replacements.
		if newEthTx.GasFeeCapIntCmp(thresholdFeeCap) < 0 || newEthTx.GasTipCapIntCmp(thresholdTip) < 0 {
			return false
		}

		// If we get here, the new transaction has a higher fee cap and tip than the old one, and
		// the percentage threshold has been met, so we can replace it.
		return true
	}
}
