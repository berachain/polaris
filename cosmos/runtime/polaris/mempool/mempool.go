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
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/core/txpool"
)

// Compile-time interface assertion.
var _ sdkmempool.Mempool = (*WrappedGethTxPool)(nil)

// WrappedGethTxPool is a mempool for Ethereum transactions. It wraps a Geth TxPool.
// NOTE: currently does not support adding `sdk.Tx`s that do NOT have a `WrappedEthereumTransaction`
// as the tx Msg.
type WrappedGethTxPool struct {
	// The underlying Geth mempool implementation.
	*txpool.TxPool

	// serializer converts eth txs to sdk txs when being iterated over.
	serializer SdkTxSerializer

	// cp is used to retrieve the current chain config.
	cp ConfigurationPlugin

	// // block data for the pending block.
	// blockNumber *big.Int
	// blockTime   uint64
	// baseFee     *big.Int
}

// NewWrappedGethTxPool creates a new Ethereum transaction pool.
func NewWrappedGethTxPool() *WrappedGethTxPool {
	return &WrappedGethTxPool{}
}

// Setup sets the chain config and sdk tx serializer on the wrapped Geth TxPool.
func (gtp *WrappedGethTxPool) Setup(txPool *txpool.TxPool, cp ConfigurationPlugin, serializer SdkTxSerializer) {
	gtp.TxPool = txPool
	gtp.cp = cp
	gtp.serializer = serializer
}

// // Prepare updates the mempool for the current block. Sets the block number, block time, and base
// // fee.
// func (gtp *WrappedGethTxPool) Prepare(header *types.Header) {
// 	gtp.blockNumber = header.Number
// 	gtp.blockTime = header.Time
// 	gtp.baseFee = header.BaseFee
// }
