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
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/gasprice"
)

const (
	// clientIdentifier is the identifier string for the client.
	clientIdentifier = "polaris-geth"

	// gpoDefault is the default gpo starting point.
	gpoDefault = 1000000000
)

// DefaultConfig returns the default JSON-RPC config.
func DefaultConfig() *Config {
	gpoConfig := ethconfig.FullNodeGPO
	gpoConfig.Default = big.NewInt(gpoDefault)
	return &Config{
		GPO:           &gpoConfig,
		RPCGasCap:     ethconfig.Defaults.RPCGasCap,
		RPCTxFeeCap:   ethconfig.Defaults.RPCTxFeeCap,
		RPCEVMTimeout: ethconfig.Defaults.RPCEVMTimeout,
	}
}

// Config represents the configurable parameters for Polaris.
type Config struct {
	// Gas Price Oracle config.
	GPO *gasprice.Config

	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap uint64 `toml:""`

	// RPCEVMTimeout is the global timeout for eth-call.
	RPCEVMTimeout time.Duration `toml:""`

	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transaction variants. The unit is ether.
	RPCTxFeeCap float64 `toml:""`
}

// LoadConfigFromFilePath reads in a Polaris config file from the fileystem.
func LoadConfigFromFilePath(filename string) (*Config, error) {
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
