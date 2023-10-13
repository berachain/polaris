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
	"context"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/config"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/params"
)

type Blockchain interface {
	PreparePlugins(context.Context)
	Config() *params.ChainConfig
	core.ChainWriter
	core.ChainReader
}

type Keeper struct {
	// host represents the host chain
	*Host

	// provider is the struct that houses the Polaris EVM.
	chain Blockchain

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
	return &Keeper{
		Host:     host,
		storeKey: storeKey,
	}
}

func (k *Keeper) Setup(chain Blockchain) error {
	k.chain = chain
	return k.SetupPrecompiles()
}

func (k *Keeper) StartEnginePlguin(ctx client.Context) {
	k.ep.Start(ctx)
}

// SetBlock sets the underlying ethereum blockchain on the keeper.
func (k *Keeper) SetBlockchain(chain Blockchain) {
	k.chain = chain
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}
