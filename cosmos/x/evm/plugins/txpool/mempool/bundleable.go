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
	"github.com/skip-mev/pob/mempool"
	mevante "github.com/skip-mev/pob/x/builder/ante"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
	eth "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool/eth"
	"pkg.berachain.dev/polaris/eth/common"
)

// BuilderMempool is an interface that combines the Mempool and mevante.Mempool interfaces.
// This allows us to have a single mempool that can be used for both the auction and the
// ethereum mempool.
type BundleableMempool interface {
	sdkmempool.Mempool
	mevante.Mempool
}

// Compile-time interface assertion.
var _ mevante.Mempool = (*EthBundlePool)(nil)

// The EthBundlePool is a mempool for Ethereum transactions. It wraps a POB Auction mempool and caches
// transactions that are added to the mempool by ethereum transaction hash.
type EthBundlePool struct {
	BundleableMempool
}

// NewEthTxPoolFrom is called when the mempool is created.
func NewEthBundlePool(
	m sdkmempool.Mempool, builderAddress common.Address, txDecoder sdk.TxDecoder,
	txEncoder sdk.TxEncoder, serializer Serializer, evmDenom string,
) *eth.Mempool {
	// Create the tx config used to route transactions to the correct mempool
	txConfig := NewMempoolConfig(builderAddress, txDecoder, serializer, evmDenom)

	// We use the auction mempool as the base mempool for the ethereum wrapper.
	return eth.NewMempoolFrom(mempool.NewAuctionMempool(txDecoder, txEncoder, 0, txConfig))
}
