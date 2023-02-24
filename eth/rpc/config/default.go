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

package config

import "time"

var (
	// `DefaultAPINamespaces` is the default namespaces the JSON-RPC server exposes.
	DefaultAPINamespaces = []string{"eth", "net"}
)

const (
	// `DefaultBaseRoute` is the default base path for the JSON-RPC server.
	DefaultBaseRoute = "/"

	// `DefaultRPCGasCap` is the default gas cap for the JSON-RPC server.
	DefaultRPCGasCap = 50000000

	// `DefaultRPCTxFeeCap` is the default transaction fee cap for the JSON-RPC server.
	DefaultRPCTxFeeCap = 1.0 // 1 ether

	// `DefaultRPCEVMTimeout` is the default timeout for the JSON-RPC server.
	DefaultRPCEVMTimeout = 5 * time.Second
)

// `DefaultServer` returns the default JSON-RPC server config.
func DefaultServer() *Server {
	return &Server{
		EnabledAPIs:   DefaultAPINamespaces,
		BaseRoute:     DefaultBaseRoute,
		RPCGasCap:     DefaultRPCGasCap,
		RPCTxFeeCap:   DefaultRPCTxFeeCap,
		RPCEVMTimeout: DefaultRPCEVMTimeout,
	}
}
