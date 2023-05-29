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
	"github.com/ethereum/go-ethereum/eth/ethconfig"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/log"
	polarapi "pkg.berachain.dev/polaris/eth/polar/api"
	"pkg.berachain.dev/polaris/eth/rpc"
)

// TODO: break out the node into a separate package and then fully use the
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
	// NetworkingStack represents the networking stack responsible for exposes the JSON-RPC APIs.
	// Although possible, it does not handle p2p networking like its sibling in geth would.
	stack NetworkingStack

	// txPool     *txpool.TxPool
	// blockchain represents the canonical chain.
	blockchain core.Blockchain

	// backend is utilize by the api handlers as a middleware between the JSON-RPC APIs and the blockchain.
	backend polarapi.Backend
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

	// TODO: decouple the networking stack from node.Node hardtype to allow for
	// alternative networking stacks, using node.Node is kinda ghetto ngl.
	var err error
	pl.stack, err = NewGethNetworkingStack(&cfg.NodeConfig, pl.backend)
	if err != nil {
		panic(err)
	}
	return pl
}

// SetNetworkingStack sets the networking stack for the polaris node.
func (pl *Polaris) SetNetworkingStack(stack NetworkingStack) {
	pl.stack = stack
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

// StartServices notifies the NetworkStack to spin up (i.e json-rpc).
func (pl *Polaris) StartServices() error {
	// Register the JSON-RPCs with the networking stack.
	pl.stack.RegisterAPIs(pl.APIs())

	// Start the services (json-rpc, graphql, etc)
	return pl.stack.Start()
}
