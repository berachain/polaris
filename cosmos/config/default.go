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

import (
	"github.com/berachain/polaris/eth/node"
	"github.com/berachain/polaris/eth/polar"

	cmtcfg "github.com/cometbft/cometbft/config"

	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// RecommendedCometBFTConfig returns the recommended CometBFT config
// for the application.
func RecommendedCometBFTConfig() *cmtcfg.Config {
	cfg := cmtcfg.DefaultConfig()
	cfg.Mempool.Size = 50000
	cfg.Mempool.CacheSize = 50000
	cfg.Mempool.Recheck = true
	cfg.Mempool.Type = "flood"

	cfg.P2P.MaxNumInboundPeers = 10
	cfg.P2P.MaxNumOutboundPeers = 15

	cfg.TxIndex.Indexer = "null"

	cfg.Instrumentation.Prometheus = true
	return cfg
}

// RecommendedServerConfig returns the recommended server config.
func RecommendedServerConfig() *serverconfig.Config {
	cfg := serverconfig.DefaultConfig()
	cfg.MinGasPrices = "0abera"
	cfg.API.Enable = true
	cfg.Telemetry.Enabled = true
	cfg.Telemetry.PrometheusRetentionTime = 180
	cfg.Telemetry.EnableHostnameLabel = true
	cfg.Telemetry.GlobalLabels = [][]string{}
	cfg.IAVLCacheSize = 20000
	return cfg
}

// DefaultPolarisConfig returns the default polaris config.
func DefaultPolarisConfig() *Config {
	nodeCfg := node.DefaultConfig()
	nodeCfg.DataDir = ""
	nodeCfg.KeyStoreDir = ""
	return &Config{
		Polar: *polar.DefaultConfig(),
		Node:  *nodeCfg,
	}
}
