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
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/block"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/provider"
)

type Keeper struct {
	// ak is the reference to the AccountKeeper.
	ak state.AccountKeeper
	// bk is the reference to the BankKeeper.
	bk state.BankKeeper
	// provider is the struct that houses the Polaris EVM.
	polaris *provider.PolarisProvider
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey
	// authority is the bech32 address that is allowed to execute governance proposals.
	authority string
	// The host contains various plugins that are are used to implement `core.PolarisHostChain`.
	host Host
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	ak state.AccountKeeper,
	bk state.BankKeeper,
	sk block.StakingKeeper,
	authority string,
	ethTxMempool sdkmempool.Mempool,
	pcs func() *ethprecompile.Injector,
) *Keeper {
	// We setup the keeper with some Cosmos standard sauce.
	k := &Keeper{
		ak:        ak,
		bk:        bk,
		authority: authority,
		storeKey:  storeKey,
	}

	k.host = NewHost(
		storeKey,
		sk,
		ethTxMempool,
		pcs,
	)
	return k
}

// Setup sets up the plugins in the Host. It also build the Polaris EVM Provider.
func (k *Keeper) Setup(
	offchainStoreKey *storetypes.KVStoreKey,
	qc func(height int64, prove bool) (sdk.Context, error),
	polarisConfigPath string,
	polarisDataDir string,

) {
	// Setup plugins in the Host
	k.host.Setup(k.storeKey, offchainStoreKey, k.ak, k.bk, qc)

	// Build the Polaris EVM Provider
	k.polaris = provider.NewPolarisProvider(polarisConfigPath, polarisDataDir, k.host, nil)
}

// ConfigureGethLogger configures the Geth logger to use the Cosmos logger.
func (k *Keeper) ConfigureGethLogger(ctx sdk.Context) {
	ethlog.Root().SetHandler(ethlog.FuncHandler(func(r *ethlog.Record) error {
		logger := ctx.Logger().With("module", "polaris-geth")
		switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
		case ethlog.LvlTrace, ethlog.LvlDebug:
			logger.Debug(r.Msg, r.Ctx...)
		case ethlog.LvlInfo, ethlog.LvlWarn:
			logger.Info(r.Msg, r.Ctx...)
		case ethlog.LvlError, ethlog.LvlCrit:
			logger.Error(r.Msg, r.Ctx...)
		}
		return nil
	}))
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}

// GetHost returns the Host that contains all plugins.
func (k *Keeper) GetHost() Host {
	return k.host
}

func (k *Keeper) SetClientCtx(clientContext client.Context) {
	k.host.GetTxPoolPlugin().(txpool.Plugin).SetClientContext(clientContext)
}
