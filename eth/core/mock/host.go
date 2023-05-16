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

package mock

import "pkg.berachain.dev/polaris/eth/core"

//go:generate moq -out ./host.mock.go -pkg mock ../ PolarisHostChain

func NewMockHostAndPlugins() (
	*PolarisHostChainMock, *BlockPluginMock, *ConfigurationPluginMock, *GasPluginMock,
	*HistoricalPluginMock, *PrecompilePluginMock, *StatePluginMock, *TxPoolPluginMock,
) {
	bp := NewBlockPluginMock()
	cp := NewConfigurationPluginMock()
	gp := NewGasPluginMock()
	hp := NewHistoricalPluginMock()
	pp := NewPrecompilePluginMock()
	sp := NewStatePluginMock()
	tp := &TxPoolPluginMock{}
	mockedPolarisHostChain := &PolarisHostChainMock{
		GetBlockPluginFunc: func() core.BlockPlugin {
			return bp
		},
		GetConfigurationPluginFunc: func() core.ConfigurationPlugin {
			return cp
		},
		GetGasPluginFunc: func() core.GasPlugin {
			return gp
		},
		GetHistoricalPluginFunc: func() core.HistoricalPlugin {
			return hp
		},
		GetPrecompilePluginFunc: func() core.PrecompilePlugin {
			return pp
		},
		GetStatePluginFunc: func() core.StatePlugin {
			return sp
		},
		GetTxPoolPluginFunc: func() core.TxPoolPlugin {
			return tp
		},
	}
	return mockedPolarisHostChain, bp, cp, gp, hp, pp, sp, tp
}
