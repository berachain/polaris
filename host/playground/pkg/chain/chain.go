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

package chain

import (
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/playground/pkg/plugins"
)

// The playground chain implements the polaris host chain interface.
var _ core.PolarisHostChain = (*Playground)(nil)

// `Playground` is the playground chain.
type Playground struct {
	blockProducer *blockProducer
	mempool       MempoolReader
}

// `NewPlayground` creates a new playground chain.
func NewPlayground(mempool MempoolReader) *Playground {
	playground := &Playground{
		mempool: mempool,
	}
	playground.blockProducer = &blockProducer{
		polaris: core.NewChain(playground),
	}
	return playground
}

func (p *Playground) ProduceBlock() (*coretypes.Block, error) {
	return p.blockProducer.ProduceBlock()
}

// `GetBlockPlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetBlockPlugin() core.BlockPlugin {
	return plugins.NewBlockPlugin()
}

// `GetConfigurationPlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetConfigurationPlugin() core.ConfigurationPlugin {
	return plugins.NewConfigurationPlugin()
}

// `GetGasPlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetGasPlugin() core.GasPlugin {
	return plugins.NewGasPlugin()
}

// `GetStatePlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetStatePlugin() core.StatePlugin {
	return plugins.NewStatePlugin()
}

// `GetTxPoolPlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetTxPoolPlugin() core.TxPoolPlugin {
	return plugins.NewTxPoolPlugin()
}

// The Playground Host Chain does not support stateful precompiles.
//
// `GetPrecompilePlugin` implements `core.PolarisHostChain`.
func (p *Playground) GetPrecompilePlugin() core.PrecompilePlugin {
	return nil
}
