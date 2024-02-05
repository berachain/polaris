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

package node

import (
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/node"
)

type (
	// Lifecycle represents a lifecycle object.
	Lifecycle = node.Lifecycle

	// GethExecutionNode is the entrypoint for the evm execution environment.
	GethExecutionNode struct {
		*node.Node
	}
)

// New creates a new execution layer node with the provided backend.
func New(config *Config) (*GethExecutionNode, error) {
	gethNode, err := node.New(&config.Config)
	if err != nil {
		return nil, err
	}

	// In Polaris we don't use P2P at the geth level.
	gethNode.SetP2PDisabled(true)

	return &GethExecutionNode{
		Node: gethNode,
	}, nil
}

// ExtRPCEnabled returns whether or not the external RPC service is enabled.
func (n *GethExecutionNode) ExtRPCEnabled() bool {
	return n.Node.Config().ExtRPCEnabled()
}

// EventMux retrieves the event multiplexer used by all the network services in
// the current protocol stack.
func (n *GethExecutionNode) EventMux() *event.TypeMux { //nolint:staticcheck // still in geth.
	return n.Node.EventMux()
}

// DefaultGethNodeConfig returns the default configuration for the provider.
func DefaultGethNodeConfig() *node.Config {
	nodeCfg := node.DefaultConfig
	nodeCfg.P2P.NoDiscovery = true
	nodeCfg.P2P.MaxPeers = 0
	nodeCfg.Name = clientIdentifier
	nodeCfg.HTTPModules = append(nodeCfg.HTTPModules, "eth", "txpool")
	nodeCfg.WSModules = append(nodeCfg.WSModules, "eth")
	nodeCfg.HTTPHost = "0.0.0.0"
	nodeCfg.WSHost = "0.0.0.0"
	nodeCfg.WSOrigins = []string{"*"}
	nodeCfg.HTTPCors = []string{"*"}
	nodeCfg.HTTPVirtualHosts = []string{"*"}
	return &nodeCfg
}
