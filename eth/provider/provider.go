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

package provider

import (
	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/api"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/rpc"
)

// PolarisProvider is the only object that an implementing chain should use.
type PolarisProvider struct {
	api.Chain
	backend rpc.PolarisBackend
	Node    *node.Node
}

// NewPolarisProvider creates a new `PolarisEVM` instance for use on an underlying blockchain.
func NewPolarisProvider(
	cfg *node.Config,
	host core.PolarisHostChain,
	logHandler log.Handler,
) *PolarisProvider {
	sp := &PolarisProvider{}
	// When creating a Polaris EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		log.Root().SetHandler(logHandler)
	}

	// Build the chain from the host.
	sp.Chain = core.NewChain(host)

	// Build and set the RPC Backend.
	sp.backend = rpc.NewPolarisBackend(sp.Chain, rpc.DefaultConfig())

	var err error
	sp.Node, err = node.New(cfg)
	if err != nil {
		panic(err)
	}

	return sp
}

// StartServices starts the standard go-ethereum node-services (i.e json-rpc).
func (sp *PolarisProvider) StartServices() error {
	sp.Node.RegisterAPIs(rpc.GetAPIs(sp.backend))
	return sp.Node.Start()
}
