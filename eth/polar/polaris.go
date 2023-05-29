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

package polar

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/log"
	polarapi "pkg.berachain.dev/polaris/eth/polar/api"
	"pkg.berachain.dev/polaris/eth/rpc"
)

// TODO: break out the node into a seperate package and then fully use the
// abstracted away networking stack, by extension we will need to improve the registration
// architecture.

var defaultEthConfig = ethconfig.Config{
	SyncMode:           0,
	FilterLogCacheSize: 0,
}

type NetworkingStack interface {
	// RegisterAPIs registers JSON-RPC handlers for the networking stack.
	RegisterAPIs([]rpc.API)

	// Start starts the networking stack.
	Start() error
}

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	// Handlers
	// NetworkingStack represents the networking stack responsible for exposes the JSON-RPC APIs.
	// Although possible, it does not handle p2p networking like its sibling in geth would.
	stack NetworkingStack

	// txPool     *txpool.TxPool
	blockchain core.Blockchain
	backend    polarapi.Backend
}

// New creates a new `PolarisEVM` instance for use on an underlying blockchain.
func New(
	configPath string,
	dataDir string,
	host core.PolarisHostChain,
	logHandler log.Handler,
) *Polaris {
	// Load the config file.
	cfg, err := LoadConfigFromFilePath(configPath)
	if err != nil {
		cfg = DefaultConfig()
	}

	// set the data dir
	cfg.NodeConfig.DataDir = dataDir

	// Create the Polaris Provider.
	return NewWithConfig(cfg, host, logHandler)
}

// New creates a new `PolarisEVM` instance for use on an underlying blockchain.
func NewWithConfig(
	cfg *Config,
	host core.PolarisHostChain,
	logHandler log.Handler,
) *Polaris {
	pl := &Polaris{}
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		log.Root().SetHandler(logHandler)
	}

	// Build the chain from the host.
	pl.blockchain = core.NewChain(host)

	// Build and set the RPC Backend.
	pl.backend = polarapi.NewBackend(pl.blockchain, &cfg.RPCConfig, &cfg.NodeConfig)

	var err error
	pl.stack, err = node.New(&cfg.NodeConfig)
	if err != nil {
		panic(err)
	}

	return pl
}

// APIs return the collection of RPC services the polar package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (pl *Polaris) APIs() []rpc.API {
	// Grab a bunch of the apis from go-ethereum (thx bae)
	apis := polarapi.GethAPIs(pl.backend, pl.blockchain)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "net",
			Service:   polarapi.NewNetAPI(pl.backend),
		},
		{
			Namespace: "web3",
			Service:   polarapi.NewWeb3API(pl.backend),
		},
	}...)
}

// StartServices starts the standard go-ethereum node-services (i.e json-rpc).
func (pl *Polaris) StartServices() error {
	// Register the JSON-RPCs with the node
	pl.stack.RegisterAPIs(pl.APIs())

	// Register the filter API separately in order to get access to the filterSystem
	// TODO: this should be made cleaner.
	filterSystem := utils.RegisterFilterAPI(pl.stack.(*node.Node), pl.backend, &defaultEthConfig)
	// this should be a flag rather than make every node default to using it
	utils.RegisterGraphQLService(pl.stack.(*node.Node), pl.backend, filterSystem, pl.stack.(*node.Node).Config())

	// Start the services (json-rpc, graphql, etc)
	return pl.stack.Start()
}
