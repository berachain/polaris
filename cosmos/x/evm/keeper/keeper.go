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

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/cosmos/config"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/engine"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/polar"
)

var _ Host = (*host)(nil)

type Keeper struct {
	// provider is the struct that houses the Polaris EVM.
	polaris *polar.Polaris

	// host represents the host chain
	host Host

	// TODO: remove this, because it's hacky af.
	storeKey storetypes.StoreKey
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	ak state.AccountKeeper,
	sk block.StakingKeeper,
	storeKey storetypes.StoreKey,
	pcs func() *ethprecompile.Injector,
	qc func() func(height int64, prove bool) (sdk.Context, error),
	logger log.Logger,
	polarisCfg *config.Config,
) *Keeper {
	host := NewHost(
		*polarisCfg,
		storeKey,
		ak,
		sk,
		pcs,
		qc,
		logger,
	)

	node, err := polar.NewGethNetworkingStack(&polarisCfg.Node)
	if err != nil {
		panic(err)
	}

	polaris := polar.NewWithNetworkingStack(&polarisCfg.Polar, host, node, ethlog.FuncHandler(
		func(r *ethlog.Record) error {
			polarisGethLogger := logger.With("module", "polaris-geth")
			switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
			case ethlog.LvlTrace, ethlog.LvlDebug:
				polarisGethLogger.Debug(r.Msg, r.Ctx...)
			case ethlog.LvlInfo, ethlog.LvlWarn:
				polarisGethLogger.Info(r.Msg, r.Ctx...)
			case ethlog.LvlError, ethlog.LvlCrit:
				polarisGethLogger.Error(r.Msg, r.Ctx...)
			}
			return nil
		}),
	)

	return &Keeper{
		polaris:  polaris,
		host:     host,
		storeKey: storeKey,
	}
}

// SetupPrecompiles initializes precompiles and the polaris node.
func (k *Keeper) SetupPrecompiles() error {
	return k.host.SetupPrecompiles()
}

// Init calls Init on the underlying polaris struct.
func (k *Keeper) Init() error {
	return k.polaris.Init()
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}

// GetPolaris returns the Polaris instance.
func (k *Keeper) Polaris() *polar.Polaris {
	return k.polaris
}

// Register Services allows for the application to register lifecycles with the evm
// networking stack.
func (k *Keeper) RegisterServices(clientContext client.Context, lcs []node.Lifecycle) {
	// TODO: probably get rid of engine plugin or something and handle rpc methods better.
	k.host.GetEnginePlugin().(engine.Plugin).Start(clientContext)

	// Register the services with polaris.
	for _, lc := range lcs {
		k.polaris.RegisterService(lc)
	}

	// Start the services.
	if err := k.polaris.StartServices(); err != nil {
		panic(err)
	}
}
