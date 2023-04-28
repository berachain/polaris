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
	"context"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

// EthTxPool is a mempool for Ethereum transactions. It wraps a PriorityNonceMempool and caches
// transactions that are added to the mempool by ethereum transaction hash.
type EthTxPool struct {
	// The underlying mempool implementation.
	sdkmempool.Mempool

	// The Polaris mempool implementation.
	core.PolarisTxPool

	mu sync.Mutex
}

// NewEthTxPoolFrom is called when the mempool is created.
func NewEthTxPoolFrom(m sdkmempool.Mempool) *EthTxPool {
	return &EthTxPool{Mempool: m}
}

// SetPolarisTxPool sets the polaris txpool for the mempool.
func (etp *EthTxPool) SetPolarisTxPool(ptxp core.PolarisTxPool) {
	etp.PolarisTxPool = ptxp
}

// Insert is called when a transaction is added to the mempool.
func (etp *EthTxPool) Insert(ctx context.Context, tx sdk.Tx) error {
	etp.mu.Lock()
	defer etp.mu.Unlock()

	// Add to the base mempool
	if err := etp.Mempool.Insert(ctx, tx); err != nil {
		return err
	}

	// Add to the Polaris TxPool if it is an eth tx
	etr, ok := utils.GetAs[*types.EthTransactionRequest](tx.GetMsgs()[0])
	if !ok {
		return nil
	}
	if err := etp.PolarisTxPool.AddLocal(etr.AsTransaction()); err != nil {
		// if the eth tx cannot be added to the polaris txpool, remove it from the base mempool
		return errorslib.Wrapf(err, "removing from base mempool %v", etp.Mempool.Remove(tx))
	}

	return nil
}

// Remove is called when a transaction is removed from the mempool.
func (etp *EthTxPool) Remove(tx sdk.Tx) error {
	etp.mu.Lock()
	defer etp.mu.Unlock()

	// Call the base mempool's Remove method
	if err := etp.Mempool.Remove(tx); err != nil {
		return err
	}

	// Verify that this tx is an EthTx
	if etr, ok := utils.GetAs[*types.EthTransactionRequest](tx); ok {
		// Remove from the Polaris TxPool if the base mempool removed it
		etp.PolarisTxPool.RemoveTx(etr.AsTransaction().Hash(), true)
	}

	return nil
}
