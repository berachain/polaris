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
	"github.com/cosmos/cosmos-sdk/client"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/handler"
	mempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/txpool"
)

// Compile-time type assertion.
var _ Plugin = (*plugin)(nil)

// Plugin defines the required functions of the transaction pool plugin.
type Plugin interface {
	plugins.Base
	core.TxPoolPlugin
	SetClientContext(client.Context)
}

// plugin represents the transaction pool plugin.
type plugin struct {
	clientCtx client.Context
	handler   *handler.Handler
}

// NewPlugin returns a new transaction pool plugin.
func NewPlugin(cp mempool.ConfigurationPlugin, ethTxMempool *mempool.WrappedGethTxPool) Plugin {
	p := &plugin{}
	p.handler = handler.NewHandler(ethTxMempool, p)
	ethTxMempool.Setup(cp, p)
	return p
}

// SetClientContext implements the Plugin interface.
func (p *plugin) SetClientContext(ctx client.Context) {
	p.clientCtx = ctx
	p.handler.SetClientContext(ctx)
}

func (p *plugin) GetHandler() txpool.Handler {
	return p.handler
}

func (p *plugin) IsPlugin() {}
