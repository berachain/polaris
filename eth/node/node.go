// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
