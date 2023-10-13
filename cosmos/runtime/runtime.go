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

package runtime

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/cosmos/config"
	libtx "pkg.berachain.dev/polaris/cosmos/lib/tx"
	"pkg.berachain.dev/polaris/cosmos/miner"
	"pkg.berachain.dev/polaris/cosmos/txpool"
	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/polar"
)

// EVMKeeper is an interface that defines the methods needed for the EVM setup.
type EVMKeeper interface {
	// Setup initializes the EVM keeper.
	Setup(evmkeeper.Blockchain) error
}

// Polaris is a struct that wraps the Polaris struct from the polar package.
// It also includes wrapped versions of the Geth Miner and TxPool.
type Polaris struct {
	*polar.Polaris

	// WrappedMiner is a wrapped version of the Miner component.
	WrappedMiner *miner.Miner
	// WrappedTxPool is a wrapped version of the Mempool component.
	WrappedTxPool *txpool.Mempool
}

// ProvidePolarisRuntime creates a new Polaris runtime from the provided
// dependencies.
func New(cfg *config.Config, logger log.Logger, host core.PolarisHostChain) *Polaris {
	node, err := polar.NewGethNetworkingStack(&cfg.Node)
	if err != nil {
		panic(err)
	}

	polaris := polar.NewWithNetworkingStack(
		&cfg.Polar, host, node, LoggerFuncHandler(logger),
	)

	return &Polaris{
		Polaris: polaris,
	}
}

// Setup is a function that sets up the Polaris struct.
// It takes a BaseApp and an EVMKeeper as arguments.
// It returns an error if the setup fails.
func (p *Polaris) Setup(bApp *baseapp.BaseApp, ek EVMKeeper) error {
	p.WrappedTxPool = txpool.New(p.TxPool())
	bApp.SetMempool(p.WrappedTxPool)

	p.WrappedMiner = miner.New(p.Miner())
	bApp.SetPrepareProposal(p.WrappedMiner.PrepareProposal)

	if err := ek.Setup(p.Blockchain()); err != nil {
		return err
	}

	// Set the ante handler to nil, since it is not needed.
	bApp.SetAnteHandler(nil)

	return nil
}

// Init is a function that initializes the Polaris struct.
// It takes a client context and a logger as arguments.
// It returns an error if the initialization fails.
func (p *Polaris) Init(clientCtx client.Context, logger log.Logger) error {
	// Initialize services.
	p.WrappedMiner.Init(libtx.NewSerializer[*engine.ExecutionPayloadEnvelope](
		clientCtx.TxConfig, evmtypes.WrapPayload))

	p.WrappedTxPool.Init(logger, clientCtx, libtx.NewSerializer[*coretypes.Transaction](
		clientCtx.TxConfig, evmtypes.WrapTx))

	// Register services with Polaris.
	p.RegisterServices(clientCtx, []node.Lifecycle{
		p.WrappedTxPool,
	})
	return nil
}

// RegisterServices is a function that allows for the application to register lifecycles with the evm
// networking stack. It takes a client context and a slice of node.Lifecycle as arguments.
func (p *Polaris) RegisterServices(_ client.Context, lcs []node.Lifecycle) {
	// Register the services with polaris.
	for _, lc := range lcs {
		p.RegisterService(lc)
	}

	// Start the services.
	if err := p.StartServices(); err != nil {
		panic(err)
	}
}

// LoadLastState is a function that loads the last state of the Polaris struct.
// It takes a CommitMultiStore and an appHeight as arguments.
// It returns an error if the loading fails.
func (p *Polaris) LoadLastState(cms storetypes.CommitMultiStore, appHeight uint64) error {
	cmsCtx := sdk.Context{}.
		WithMultiStore(cms).
		WithGasMeter(storetypes.NewInfiniteGasMeter()).
		WithBlockGasMeter(storetypes.NewInfiniteGasMeter()).WithEventManager(sdk.NewEventManager())
	return p.Blockchain().LoadLastState(cmsCtx, appHeight)
}
