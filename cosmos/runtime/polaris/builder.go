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

package polaris

import (
	"encoding/json"
	"io"

	"cosmossdk.io/log"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"

	"pkg.berachain.dev/polaris/cosmos/runtime/polaris/mempool"
	"pkg.berachain.dev/polaris/eth/core"
)

// AppBuilder is a type that is injected into a container by the runtime module
// (as *AppBuilder) which can be used to create an app which is compatible with
// the existing app.go initialization conventions.
type AppBuilder struct {
	// Used to build the baseapp
	*runtime.AppBuilder

	//  used for export.
	polarisApp *PolarisApp
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (a *AppBuilder) DefaultGenesis() map[string]json.RawMessage {
	return a.AppBuilder.DefaultGenesis()
}

// Build builds an *App instance.
func (a *AppBuilder) Build(db dbm.DB,
	traceStore io.Writer,
	ethTxMempool *mempool.WrappedGethTxPool, logger log.Logger, host core.PolarisHostChain,
	baseAppOptions ...func(*baseapp.BaseApp)) *PolarisApp {
	a.polarisApp = &PolarisApp{
		hostChain: host,
	}

	// TODO: move this somewhere better, introduce non IAVL enforced module keys as a PR to the SDK
	// we ask @tac0turtle how 2 fix
	// offchainKey := storetypes.NewKVStoreKey("offchain-evm")

	// Build the base runtime.App (and thus baseapp.BaseApp)
	a.polarisApp.App = a.AppBuilder.Build(db, traceStore, baseAppOptions...)
	a.polarisApp.logger = logger
	a.polarisApp.mempool = ethTxMempool
	a.polarisApp.SetMempool(a.polarisApp.mempool)
	// Configure the proposal handlers to use our custom tx pool and miner.
	// a.polarisApp.wrappedTxPool = txPool
	// a.polarisApp.SetPrepareProposal(pflarisMiner.PrepareProposalHandler())
	// a.polarisApp.SetProcessProposal(polarisMiner.ProcessProposalHandler())

	// Mount our custom stores.
	// a.polarisApp.MountCustomStores(offchainKey)

	// Return the app
	return a.polarisApp
}
