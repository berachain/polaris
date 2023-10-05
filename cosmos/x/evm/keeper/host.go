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
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/config"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/configuration"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/engine"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/historical"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile"
	pclog "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// Compile-time interface assertion.
var _ core.PolarisHostChain = (*host)(nil)

// Host is the interface that must be implemented by the host.
// It includes core.PolarisHostChain and functions that are called in other packages.
type Host interface {
	core.PolarisHostChain
	GetAllPlugins() []any
	SetupPrecompiles()
}

type host struct {
	// The various plugins that are are used to implement core.PolarisHostChain.
	bp     block.Plugin
	cp     configuration.Plugin
	ep     engine.Plugin
	hp     historical.Plugin
	pp     precompile.Plugin
	sp     state.Plugin
	logger log.Logger

	ak       state.AccountKeeper
	storeKey storetypes.StoreKey
	pcs      func() *ethprecompile.Injector
	qc       func() func(height int64, prove bool) (sdk.Context, error)
}

// Newhost creates new instances of the plugin host.
func NewHost(
	cfg config.Config,
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	sk block.StakingKeeper,
	precompiles func() *ethprecompile.Injector,
	qc func() func(height int64, prove bool) (sdk.Context, error),
	logger log.Logger,
) Host {
	// We setup the host with some Cosmos standard sauce.
	h := &host{}

	// Build the Plugins
	h.bp = block.NewPlugin(storeKey, sk)
	h.cp = configuration.NewPlugin(&cfg.Polar.Chain)
	h.ep = engine.NewPlugin()
	h.pcs = precompiles
	h.storeKey = storeKey
	h.ak = ak
	h.qc = qc
	h.logger = logger

	// Setup the state, precompile, historical, and txpool plugins
	// TODO: re-enable historical plugin using ABCI listener.
	h.hp = historical.NewPlugin(h.cp, h.bp, nil, h.storeKey)
	h.pp = precompile.NewPlugin()
	h.sp = state.NewPlugin(h.ak, h.storeKey, nil)
	h.bp.SetQueryContextFn(h.qc)
	h.sp.SetQueryContextFn(h.qc)

	return h
}

// SetupPrecompiles intializes the precompile contracts.
func (h *host) SetupPrecompiles() {
	// Set the query context function for the block and state plugins
	pcs := h.pcs().GetPrecompiles()
	h.pp.RegisterPrecompiles(pcs)
	h.sp.SetPrecompileLogFactory(pclog.NewFactory(pcs))
}

// GetBlockPlugin returns the header plugin.
func (h *host) GetBlockPlugin() core.BlockPlugin {
	return h.bp
}

// GetConfigurationPlugin returns the configuration plugin.
func (h *host) GetConfigurationPlugin() core.ConfigurationPlugin {
	return h.cp
}

// GetEnginePlugin returns the engine plugin.
func (h *host) GetEnginePlugin() core.EnginePlugin {
	return h.ep
}
func (h *host) GetHistoricalPlugin() core.HistoricalPlugin {
	return h.hp
}

// GetPrecompilePlugin returns the precompile plugin.
func (h *host) GetPrecompilePlugin() core.PrecompilePlugin {
	return h.pp
}

// GetStatePlugin returns the state plugin.
func (h *host) GetStatePlugin() core.StatePlugin {
	return h.sp
}

// GetAllPlugins returns all the plugins.
func (h *host) GetAllPlugins() []any {
	return []any{h.bp, h.cp, h.hp, h.pp, h.sp}
}
