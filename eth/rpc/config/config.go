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

type (
	// Server defines the configuration for the JSON-RPC server.
	Server struct {
		// Enabled defines whether the JSON-RPC server should be enabled.
		Enabled bool `mapstructure:"enabled"`

		// API defines a list of JSON-RPC namespaces to be enabled.
		EnabledAPIs []string `mapstructure:"api"`

		// HTTPBaseRoute defines the base path for the JSON-RPC server.
		BaseRoute string `mapstructure:"base-path"`

		// RPCGasCap is the global gas cap for eth-call variants.
		RPCGasCap uint64 `mapstructure:"rpc-gas-cap"`

		// RPCEVMTimeout is the global timeout for eth-call.
		RPCEVMTimeout time.Duration `mapstructure:"rpc-evm-timeout"`

		// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
		// send-transaction variants. The unit is ether.
		RPCTxFeeCap float64 `mapstructure:"rpc-tx-fee-cap"`

		// OriginAllowlist defines a list of allowed origins for the JSON-RPC server.
		OriginAllowlist []string `mapstructure:"origin-allowlist"`
	}
)
