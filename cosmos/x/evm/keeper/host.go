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
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/config"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/block"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/historical"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/precompile"
	pclog "github.com/berachain/polaris/cosmos/x/evm/plugins/precompile/log"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/eth/core"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

// Compile-time interface assertion.
var _ core.PolarisHostChain = (*Host)(nil)

type Host struct {
	// The various plugins that are are used to implement core.PolarisHostChain.
	bp  block.Plugin
	hp  historical.Plugin
	pp  precompile.Plugin
	sp  state.Plugin
	spf *state.SPFactory

	pcs func() *ethprecompile.Injector
}

// Newhost creates new instances of the plugin host.
func NewHost(
	cfg config.Config,
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	precompiles func() *ethprecompile.Injector,
	qc func() func(height int64, prove bool) (sdk.Context, error),
) *Host {
	// We setup the host with some Cosmos standard sauce.
	h := &Host{
		bp: block.NewPlugin(
			storeKey, qc,
		),
		pcs: precompiles,
		pp:  precompile.NewPlugin(),
		sp:  state.NewPlugin(ak, storeKey, qc, nil),
	}

	// historical plugin requires block plugin.
	h.hp = historical.NewPlugin(&cfg.Polar.Chain, h.bp, nil, storeKey)
	h.spf = state.NewSPFactory(ak, storeKey, qc)
	return h
}

// SetupPrecompiles initializes the precompile contracts.
func (h *Host) SetupPrecompiles() error {
	// Set the query context function for the block and state plugins
	pcs := h.pcs().GetPrecompiles()

	if err := h.pp.RegisterPrecompiles(pcs); err != nil {
		return err
	}

	h.sp.SetPrecompileLogFactory(pclog.NewFactory(pcs))
	h.spf.SetPrecompileLogFactory(pclog.NewFactory(pcs))
	return nil
}

// GetBlockPlugin returns the header plugin.
func (h *Host) GetBlockPlugin() core.BlockPlugin {
	return h.bp
}

// GetHistoricalPlugin returns the historical plugin.
func (h *Host) GetHistoricalPlugin() core.HistoricalPlugin {
	return h.hp
}

// GetPrecompilePlugin returns the precompile plugin.
func (h *Host) GetPrecompilePlugin() core.PrecompilePlugin {
	return h.pp
}

func (h *Host) GetStatePluginFactory() core.StatePluginFactory {
	return h.spf
}

// GetAllPlugins returns all the plugins.
func (h *Host) GetAllPlugins() []any {
	return []any{h.bp, h.hp, h.pp, h.sp}
}

// Version returns the version of the host chain.
func (h *Host) Version() string {
	versionInfo := version.NewInfo()
	return versionInfo.AppName + "/" + version.Version + ":" + "cosmos/" +
		versionInfo.CosmosSdkVersion
}
