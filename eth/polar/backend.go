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
	"time"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/params"
)

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	*eth.Ethereum
	config *Config
	// NetworkingStack represents the networking stack responsible for exposes the JSON-RPC
	// APIs. Although possible, it does not handle p2p networking like its sibling in geth
	// would.
	stack *node.Node
}

func NewWithNetworkingStack(
	config *Config,
	stack *node.Node,
	logHandler log.Handler,
) *Polaris {
	defats := &ethconfig.Defaults
	defats.Genesis = core.DefaultGenesis
	defats.Genesis.Config = params.DefaultChainConfig
	ethereum, err := eth.New(stack, defats)
	if err != nil {
		panic(err)
	}
	pl := &Polaris{
		Ethereum: ethereum,
		config:   config,
		stack:    stack,
		// ethereum:     ethereum,
	}
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		log.Root().SetHandler(logHandler)
	}

	// pl.backend = NewBackend(pl, pl.config)

	return pl
}

// StartServices notifies the NetworkStack to spin up (i.e json-rpc).
func (pl *Polaris) StartServices() error {
	go func() {
		// TODO: these values are sensitive due to a race condition in the json-rpc ports opening.
		// If the JSON-RPC opens before the first block is committed, hive tests will start failing.
		// This needs to be fixed before mainnet as its ghetto af. If the block time is too long
		// and this sleep is too short, it will cause hive tests to error out.
		time.Sleep(5 * time.Second) //nolint:gomnd // as explained above.
		if err := pl.stack.Start(); err != nil {
			panic(err)
		}
	}()
	return nil
}

// RegisterService adds a service to the networking stack.
func (pl *Polaris) RegisterService(lc node.Lifecycle) {
	pl.stack.RegisterLifecycle(lc)
}

func (pl *Polaris) Close() error {
	return pl.stack.Close()
}
