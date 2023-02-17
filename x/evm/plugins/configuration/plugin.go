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

package configuration

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	paramsPrefix = []byte("params")
)

// `plugin` implements the core.ConfigurationPlugin interface.
type plugin struct {
	evmStoreKey storetypes.StoreKey
	paramsStore storetypes.KVStore
}

// `NewPlugin` returns a new plugin instance.
func NewPlugin() core.ConfigurationPlugin {
	return &plugin{}
}

// `Prepare` implements the core.ConfigurationPlugin interface.
func (p *plugin) Prepare(ctx context.Context) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	p.paramsStore = prefix.NewStore(sCtx.KVStore(p.evmStoreKey), paramsPrefix)
}

// `ChainConfig` implements the core.ConfigurationPlugin interface.
func (p *plugin) ChainConfig() *params.ChainConfig {
	return p.GetParams().EthereumChainConfig()
}

// `ExtraEips` implements the core.ConfigurationPlugin interface.
func (p *plugin) ExtraEips() []int {
	eips := make([]int, 0)
	for _, e := range p.GetParams().ExtraEIPs {
		eips = append(eips, int(e))
	}
	return eips
}
