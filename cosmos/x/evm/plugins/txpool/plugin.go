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

package txpool

import (
	"math/big"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	mempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Compile-time type assertion.
var _ Plugin = (*plugin)(nil)

// Plugin defines the required functions of the transaction pool plugin.
type Plugin interface {
	plugins.Base
	Start()
	SetLogger(logger log.Logger)
	Setup(*txpool.TxPool, client.Context)
	Prepare(*big.Int, coretypes.Signer)
}

// plugin represents the transaction pool plugin.
type plugin struct {
	*mempool.WrappedGethTxPool
	*handler
	serializer *serializer
}

// NewPlugin returns a new transaction pool plugin.
func NewPlugin(wrappedGethTxPool *mempool.WrappedGethTxPool) Plugin {
	return &plugin{
		WrappedGethTxPool: wrappedGethTxPool,
	}
}

// Setup implements the Plugin interface.
func (p *plugin) Setup(txpool *txpool.TxPool, ctx client.Context) {
	p.serializer = newSerializer(ctx)
	p.WrappedGethTxPool.Setup(txpool, p.serializer)
	p.handler = newHandler(ctx, txpool, p.serializer)
}

func (p *plugin) IsPlugin() {}
