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
	"google.golang.org/grpc"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/client"

	"pkg.berachain.dev/polaris/cosmos/runtime/polaris/mempool"
	"pkg.berachain.dev/polaris/cosmos/runtime/polaris/miner"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/polar"
)

// Ensure that the Runtime implements the HasServices interface.
var _ appmodule.HasServices = &Runtime{}

// Runtime represents the runtime used by the Polaris EVM.
type Runtime struct {
	logger    log.Logger
	clientCtx client.Context

	polaris   *polar.Polaris
	miner     *miner.Miner
	mempool   *mempool.WrappedGethTxPool
	hostChain core.PolarisHostChain
}

// CreateRuntime creates a new Polaris runtime.
func CreateRuntime(
	logger log.Logger,
	miner *miner.Miner,
	mempool *mempool.WrappedGethTxPool,
	hostChain core.PolarisHostChain,
) *Runtime {
	return &Runtime{
		logger:    logger,
		miner:     miner,
		mempool:   mempool,
		hostChain: hostChain,
	}
}

// Load is called on application initialization and provides an opportunity to
// perform initialization logic. It returns an error if initialization fails.
// We shadow the Load function from cosmos-sdk/runtime/app.go in order to prime the blockchain.
func (r *Runtime) Load() error {
	// Load the runtime config from disk.
	nodeCfg, err := r.LoadConfigFromDisk()
	if err != nil {
		return err
	}

	node, err := polar.NewGethNetworkingStack(nodeCfg)
	if err != nil {
		panic(err)
	}

	r.polaris = polar.New(
		polar.DefaultConfig(),
		r.hostChain,
		node,
		ethlog.FuncHandler(
			func(record *ethlog.Record) error {
				polarisGethLogger := r.logger.With("module", "polaris-geth")
				switch record.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
				case ethlog.LvlTrace, ethlog.LvlDebug:
					polarisGethLogger.Debug(record.Msg, record.Ctx...)
				case ethlog.LvlInfo, ethlog.LvlWarn:
					polarisGethLogger.Info(record.Msg, record.Ctx...)
				case ethlog.LvlError, ethlog.LvlCrit:
					polarisGethLogger.Error(record.Msg, record.Ctx...)
				}
				return nil
			}),
		nil,
	)
	return nil
}

// TODO: deprecate this eventually.
func (r *Runtime) Polaris() *polar.Polaris {
	return r.polaris
}
func (r *Runtime) SetClientCtx(clientCtx client.Context) {
	r.clientCtx = clientCtx
}

// RegisterServices registers the services that are used by the Polaris EVM.
func (r *Runtime) RegisterServices(grpc.ServiceRegistrar) error {
	// We don't actually need the services to be registered
	// TODO: TxPool initialization

	// Intializating the txpool here seems to be safe, but I think there may be a race condition here.
	// What we should really be doing is reading the ChainConfig directly from the database after `Load()` is called.
	// Opposed to getting it through the blockchain object, which won't be prepare'd until after BeginBlock(), I think
	// it is just happenstance that RegisterAPIRoutes is happen after the blockchain is prepared.
	//
	// Note: Once we are properly loading the state of the blockchain in `Load()` this issue should be formally resolved.
	txPool := txpool.NewTxPool(txpool.DefaultConfig, r.polaris.Blockchain().Config(), r.polaris.Blockchain())
	r.polaris.SetTxPool(txPool)

	// Now that we have the client context and the txpool, we can setup the mempool and miner.
	r.mempool.Setup(txPool, r.hostChain.GetConfigurationPlugin(), mempool.NewTxSerializer(r.clientCtx))

	// We set the handler.
	r.polaris.SetHandler(mempool.NewHandler(r.logger, r.clientCtx, txPool))

	// Note: this is a bit of an awkward place to put this, but if you look in the Cosmos-SDK server:
	// https://github.com/cosmos/cosmos-sdk/blob/3db9528efb5fec1cccdb4e6f084c24ed195951b1/server/start.go#L504
	// You'll see that the API server is started right after `RegisterAPIRoutes` is called. So starting the
	// Polaris services here is quite oddly a semi logical place to do it (in lieu of having a custom
	// server even though it feels a little strange.
	// TODO: Really we should create a way for runtime modules to register services with the server package.
	// We suggest this to @tac0turtle.

	r.miner.Start() // move somewhere better
	return r.polaris.StartServices()
}
