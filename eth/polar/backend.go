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
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/rpc"
)

var defaultEthConfig = ethconfig.Config{
	SyncMode:           0,
	FilterLogCacheSize: 0,
}

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	config *ethconfig.Config

	// Handlers
	node       *node.Node
	txPool     *txpool.TxPool
	blockchain core.Blockchain
	backend    rpc.PolarisBackend

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	eventMux *event.TypeMux
	// engine         consensus.Engine
	accountManager *accounts.Manager

	APIBackend *EthAPIBackend

	// miner     *miner.Miner
	gasPrice  *big.Int
	etherbase common.Address

	networkID     uint64
	netRPCService *ethapi.NetAPI

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
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
	sp := &Polaris{}
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		log.Root().SetHandler(logHandler)
	}

	// Build the chain from the host.
	sp.blockchain = core.NewChain(host)

	// Build and set the RPC Backend.
	sp.backend = rpc.NewPolarisBackend(sp.blockchain, &cfg.RPCConfig, &cfg.NodeConfig)

	var err error
	sp.node, err = node.New(&cfg.NodeConfig)
	if err != nil {
		panic(err)
	}

	return sp
}

// APIs return the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Polaris) APIs() []rpc.API {
	apis := ethapi.GetAPIs(s.APIBackend, core.NewChainContext(s.blockchain))
	apis = append(apis, rpc.GetAPIs(s.backend)...)
	// Append any APIs exposed explicitly by the consensus engine
	// apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		// {
		// 	Namespace: "eth",
		// 	Service:   eth.NewEthereumAPI(s),
		// }, {
		// 	Namespace: "debug",
		// 	Service:   eth.NewDebugAPI(s),
		// }, {
		{
			Namespace: "net",
			Service:   s.netRPCService,
		},
	}...)
}

// StartServices starts the standard go-ethereum node-services (i.e json-rpc).
func (sp *Polaris) StartServices() error {
	// Register the filter API separately in order to get access to the filterSystem
	// TODO: this should be made cleaner.
	filterSystem := utils.RegisterFilterAPI(sp.node, sp.backend, &defaultEthConfig)
	// this should be a flag rather than make every node default to using it
	utils.RegisterGraphQLService(sp.node, sp.backend, filterSystem, sp.node.Config())
	return sp.node.Start()
}
