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

package config

import (
	"time"

	"github.com/berachain/polaris/eth/node"
	"github.com/berachain/polaris/eth/polar"

	cmtcfg "github.com/cometbft/cometbft/config"

	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// RecommendedCometBFTConfig returns the recommended CometBFT config
// for the application.
func RecommendedCometBFTConfig() *cmtcfg.Config {
	cfg := cmtcfg.DefaultConfig()
	cfg.Mempool.Size = 3000
	cfg.Mempool.CacheSize = 250000
	cfg.Mempool.Recheck = true
	cfg.Mempool.Type = "flood"

	cfg.P2P.MaxNumInboundPeers = 40
	cfg.P2P.MaxNumOutboundPeers = 20

	cfg.TxIndex.Indexer = "null"

	cfg.Consensus.TimeoutPropose = 3 * time.Second //nolint:gomnd // default.
	cfg.Consensus.TimeoutPrevote = 1 * time.Second
	cfg.Consensus.TimeoutPrecommit = 1 * time.Second

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
	cfg.AppDBBackend = "pebbledb"
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
