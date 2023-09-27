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
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core/txpool"

	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Compile-time type assertion.
var _ Plugin = (*plugin)(nil)

type Serializer interface {
	SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error)
	SerializeToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error)
}

// Plugin defines the required functions of the transaction pool plugin.
type Plugin interface {
	core.TxPoolPlugin
	Start(log.Logger, *txpool.TxPool, client.Context)
	// Prepare(*big.Int, coretypes.Signer)
	SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error)
	GetTxPool() *txpool.TxPool
}

// plugin represents the transaction pool plugin.
type plugin struct {
	*txpool.TxPool
	*handler
	serializer *serializer
}

// NewPlugin returns a new transaction pool plugin.
func NewPlugin() Plugin {
	return &plugin{}
}

// GetHandler implements the Plugin interface.
func (p *plugin) GetHandler() core.Handler {
	return p.handler
}

func (p *plugin) GetTxPool() *txpool.TxPool {
	return p.TxPool
}

func (p *plugin) SerializeToBytes(signedTx *coretypes.Transaction) ([]byte, error) {
	return p.serializer.SerializeToBytes(signedTx)
}

// Setup implements the Plugin interface.
func (p *plugin) Start(logger log.Logger, txpool *txpool.TxPool, ctx client.Context) {
	p.serializer = newSerializer(ctx)
	p.TxPool = txpool
	p.handler = newHandler(ctx, txpool, p.serializer, logger)

	// TODO: register all these starting things somewhere better.
	p.handler.Start()
}
