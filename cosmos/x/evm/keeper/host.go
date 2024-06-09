// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
