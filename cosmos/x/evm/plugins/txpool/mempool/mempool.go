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
	"sync"

	"github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// EthTxPool is a mempool for Ethereum transactions. It wraps a PriorityNonceMempool and caches
// transactions that are added to the mempool by ethereum transaction hash.
type EthTxPool struct {
	// The underlying mempool implementation.
	*PriorityNonceMempool[int64]

	// ethTxCache caches transactions that are added to the mempool so that they can be retrieved
	// later
	ethTxCache  map[common.Hash]*coretypes.Transaction
	nonceToHash map[uint64]common.Hash

	// NonceRetriever is used to retrieve the nonce for a given address (this is typically a
	// reference to the StateDB).
	nr NonceRetriever

	// We have a mutex to protect the ethTxCache and nonces maps since they are accessed
	// concurrently by multiple goroutines.
	mu sync.RWMutex
}

// NewPolarisEthereumTxPool creates a new Ethereum transaction pool.
func NewPolarisEthereumTxPool() *EthTxPool {
	config := mempool.DefaultPriorityNonceMempoolConfig()
	config.TxReplacement = EthereumTxReplacePolicy[int64]{
		PriceBump: 10, //nolint:gomnd // 10% to match geth.
	}.Func
	return &EthTxPool{
		PriorityNonceMempool: NewPriorityMempool[int64](config),
		nonceToHash:          make(map[uint64]common.Hash),
		ethTxCache:           make(map[common.Hash]*coretypes.Transaction),
	}
}
