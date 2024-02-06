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
	"math/big"
	"time"

	"github.com/berachain/polaris/eth/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
	ethparams "github.com/ethereum/go-ethereum/params"
)

const (
	// gpoDefault is the default gpo starting point.
	gpoDefault = 1000000000

	// developmentCoinbase is the address used for development.
	// DO NOT USE IN PRODUCTION.
	// 0xf8637fa70e8e329ecb8463b788d96914f8cfe191d15ae36f161227629e3f5693.
	developmentCoinbase = "0xAf15f95bed0D3913a29092Fd7837451Ce4de64D3"
)

// DefaultConfig returns the default JSON-RPC config.
func DefaultConfig() *Config {
	gpoConfig := ethconfig.FullNodeGPO
	gpoConfig.Default = big.NewInt(gpoDefault)
	gpoConfig.MaxPrice = big.NewInt(ethparams.GWei * 10000) //nolint:gomnd // default.
	minerCfg := miner.DefaultConfig
	minerCfg.Etherbase = common.HexToAddress(developmentCoinbase)
	minerCfg.GasPrice = big.NewInt(1)
	legacyPool := legacypool.DefaultConfig
	legacyPool.NoLocals = true
	legacyPool.PriceLimit = 8 // to handle the low base fee.
	legacyPool.Journal = ""

	return &Config{
		Chain:         *params.DefaultChainConfig,
		Miner:         minerCfg,
		GPO:           gpoConfig,
		LegacyTxPool:  legacyPool,
		RPCGasCap:     ethconfig.Defaults.RPCGasCap,
		RPCTxFeeCap:   ethconfig.Defaults.RPCTxFeeCap,
		RPCEVMTimeout: ethconfig.Defaults.RPCEVMTimeout,
	}
}

// SafetyMessage is a safety check for the JSON-RPC config.
func (c *Config) SafetyMessage() {
	if c.Miner.Etherbase == common.HexToAddress(developmentCoinbase) {
		log.Error(
			"development etherbase in use - please verify this is intentional", "address",
			c.Miner.Etherbase,
		)
	}
}

// Config represents the configurable parameters for Polaris.
type Config struct {
	// The chain configuration to use.
	Chain ethparams.ChainConfig

	// Mining options
	Miner miner.Config

	// Gas Price Oracle config.
	GPO gasprice.Config

	// Transaction pool options
	LegacyTxPool legacypool.Config

	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap uint64

	// RPCEVMTimeout is the global timeout for eth-call.
	RPCEVMTimeout time.Duration

	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transaction variants. The unit is ether.
	RPCTxFeeCap float64
}
