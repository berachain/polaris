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

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/x/evm/plugins"
	"github.com/berachain/stargazer/x/evm/plugins/block"
	"github.com/berachain/stargazer/x/evm/plugins/configuration"
	"github.com/berachain/stargazer/x/evm/plugins/gas"
	"github.com/berachain/stargazer/x/evm/plugins/precompile"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	"github.com/berachain/stargazer/x/evm/types"

	"github.com/cometbft/cometbft/libs/log"

	precompilelog "github.com/berachain/stargazer/x/evm/plugins/precompile/log"
)

// Compile-time interface assertion.
var _ core.StargazerHostChain = (*Keeper)(nil)

type Keeper struct {
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey

	stargazer *eth.StargazerProvider

	// sk is used to retrieve infofrmation about the current / past
	// blocks and associated validator information.
	// sk StakingKeeper

	authority string

	// plugins
	bp block.Plugin
	cp configuration.Plugin
	gp gas.Plugin
	pp precompile.Plugin
	sp state.Plugin
}

// NewKeeper creates new instances of the stargazer Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	bk state.BankKeeper,
	authority string,
) *Keeper {
	k := &Keeper{
		authority: authority,
		storeKey:  storetypes.NewKVStoreKey(types.StoreKey),
	}

	k.bp = block.NewPlugin(k)

	k.cp = configuration.NewPlugin(storeKey)

	k.gp = gas.NewPlugin()

	k.pp = precompile.NewPlugin()
	// TODO: register precompiles

	plf := precompilelog.NewFactory()
	// TODO: register precompile events/logs

	k.sp = state.NewPlugin(ak, bk, k.storeKey, types.ModuleName, plf)

	k.stargazer = eth.NewStargazerProvider(k, nil)
	// TODO: provide cosmos ctx logger.

	return k
}

// `Logger` returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

func (k *Keeper) GetBlockPlugin() core.BlockPlugin {
	return k.bp
}

func (k *Keeper) GetConfigurationPlugin() core.ConfigurationPlugin {
	return k.cp
}

func (k *Keeper) GetGasPlugin() core.GasPlugin {
	return k.gp
}

func (k *Keeper) GetPrecompilePlugin() core.PrecompilePlugin {
	return k.pp
}

func (k *Keeper) GetStatePlugin() core.StatePlugin {
	return k.sp
}

func (k *Keeper) GetStargazer() *eth.StargazerProvider {
	return k.stargazer
}

func (k *Keeper) GetAllPlugins() []plugins.BaseCosmosStargazer {
	return []plugins.BaseCosmosStargazer{
		k.bp,
		k.cp,
		k.gp,
		k.pp,
		k.sp,
	}
}
