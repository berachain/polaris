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
	"github.com/ethereum/go-ethereum/node"
)

// Node is a wrapper around the go-ethereum node.Node object, that allows us to conform to the
// NetworkingStack interface, we have to do some hacky stuff to initialize the graphql service,
// TODO: deprecate this and use a more elegant solution.
type Node struct {
	*node.Node
}

// NewGetNetworkingStack creates a new NetworkingStack instance for use on an underlying blockchain.
func NewGethNetworkingStack(config *node.Config) (NetworkingStack, error) {
	node, err := node.New(config)
	if err != nil {
		return nil, err
	}

	return &Node{
		Node: node,
	}, nil
}

// Start starts the networking stack.
func (n *Node) Start() error {
	// We then start the underlying node.
	return n.Node.Start()
}
