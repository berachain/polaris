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
	"context"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/graphql"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
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

// NetworkingStack defines methods that allow a Polaris chain to build and expose JSON-RPC apis.
type NetworkingStack interface {
	// IsExtRPCEnabled returns true if the networking stack is configured to expose JSON-RPC APIs.
	ExtRPCEnabled() bool

	// RegisterHandler manually registers a new handler into the networking stack.
	RegisterHandler(string, string, http.Handler)

	// RegisterAPIs registers JSON-RPC handlers for the networking stack.
	RegisterAPIs([]rpc.API)

	// Start starts the networking stack.
	Start() error
}

// PolarisBase defines methods that any implementation of Polaris must support
type PolarisBase interface {
	// APIs return the collection of RPC services the polar package offers.
	APIs() []rpc.API

	// StartServices notifies the NetworkStack to spin up (i.e json-rpc).
	StartServices() error

	// Prepare prepares the Polaris chain for processing a new block at the given height.
	Prepare(ctx context.Context, number uint64)

	// ProcessTransaction processes the given transaction and returns the receipt.
	ProcessTransaction(ctx context.Context, tx *types.Transaction) (*core.ExecutionResult, error)

	// Finalize finalizes the current block.
	Finalize(ctx context.Context) error
}

// Polaris is the only object that an implementing chain should use.
type Polaris struct {
	cfg *Config
	// NetworkingStack represents the networking stack responsible for exposes the JSON-RPC APIs.
	// Although possible, it does not handle p2p networking like its sibling in geth would.
	stack NetworkingStack

	// txPool     *txpool.TxPool
	// blockchain represents the canonical chain.
	blockchain core.Blockchain

	// backend is utilize by the api handlers as a middleware between the JSON-RPC APIs and the blockchain.
	backend Backend

	// filterSystem is the filter system that is used by the filter API.
	// TODO: relocate
	filterSystem *filters.FilterSystem
}

func NewWithNetworkingStack(
	cfg *Config,
	host core.PolarisHostChain,
	stack NetworkingStack,
	logHandler log.Handler,
) *Polaris {
	pl := &Polaris{
		cfg:        cfg,
		blockchain: core.NewChain(host),
		stack:      stack,
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

	// Build and set the RPC Backend.
	pl.backend = NewBackend(pl, stack.ExtRPCEnabled(), cfg)
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

// StartServices notifies the NetworkStack to spin up (i.e json-rpc).
func (pl *Polaris) StartServices() error {
	// Register the JSON-RPCs with the networking stack.
	pl.stack.RegisterAPIs(pl.APIs())

	// Register the filter API separately in order to get access to the filterSystem
	pl.filterSystem = utils.RegisterFilterAPI(pl.stack, pl.backend, &defaultEthConfig)

	// Register the GraphQL API (todo update cors stuff)
	// TODO: gate this behind a flag
	if err := graphql.New(pl.stack, pl.backend, pl.filterSystem, []string{"*"}, []string{"*"}); err != nil {
		return err
	}

	go func() {
		// TODO: unhack this.
		time.Sleep(2 * time.Second) //nolint:gomnd // we will fix this eventually.
		if pl.stack.Start() != nil {
			os.Exit(1)
		}
	}()
	return nil
}
