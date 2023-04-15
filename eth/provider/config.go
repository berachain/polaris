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
	"fmt"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/eth/rpc"
)

// DefaultConfig returns the default configuration for the provider.
func DefaultConfig() *Config {
	c := Config{}
	c.ChainConfig = *params.DefaultChainConfig
	nodeCfg := node.DefaultConfig
	nodeCfg.P2P.MaxPeers = 0
	nodeCfg.HTTPModules = append(nodeCfg.HTTPModules, "eth")
	nodeCfg.WSModules = append(nodeCfg.WSModules, "eth")
	nodeCfg.HTTPHost = "localhost"
	nodeCfg.WSHost = "localhost"
	nodeCfg.WSOrigins = []string{"*"}
	c.NodeConfig = nodeCfg
	c.RPCConfig = *rpc.DefaultConfig()
	return &c
}

// Config represents the configurable parameters for Polaris.
type Config struct {
	ChainConfig params.ChainConfig
	NodeConfig  node.Config
	RPCConfig   rpc.Config
}

// ReadConfigFile reads in a Polaris config file from the fileystem.
func ReadConfigFile(filename string) (*Config, error) {
	var config Config

	// Read the TOML file
	bytes, err := os.ReadFile(filename) //#nosec: G304 // required.
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filename, err)
	}

	// Unmarshal the TOML data into a struct
	if err = toml.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("error parsing TOML data: %w", err)
	}

	return &config, nil
}
