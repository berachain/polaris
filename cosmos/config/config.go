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
	"math/big"

	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/polar"
)

// DefaultConfig returns the default configuration for a polaris chain.
func DefaultConfig() *Config {
	nodeCfg := *polar.DefaultGethNodeConfig()
	nodeCfg.DataDir = ""
	nodeCfg.KeyStoreDir = ""
	return &Config{
		Polar: *polar.DefaultConfig(),
		Node:  nodeCfg,
	}
}

type Config struct {
	Polar polar.Config
	Node  node.Config
}

func MustReadConfigFromAppOpts(opts servertypes.AppOptions) *Config {
	cfg, err := ReadConfigFromAppOpts(opts)
	if err != nil {
		panic(err)
	}
	return cfg
}

func ReadConfigFromAppOpts(opts servertypes.AppOptions) (*Config, error) {
	return readConfigFromAppOptsParser(AppOptionsParser{AppOptions: opts})
}

//nolint:funlen,gocognit,gocyclo,cyclop // TODO break up later.
func readConfigFromAppOptsParser(parser AppOptionsParser) (*Config, error) {
	var err error
	var val int64
	conf := &Config{}

	// Polar settings
	if conf.Polar.RPCGasCap, err =
		parser.GetUint64(flagRPCGasCap); err != nil {
		return nil, err
	}
	if conf.Polar.RPCEVMTimeout, err =
		parser.GetTimeDuration(flagRPCEvmTimeout); err != nil {
		return nil, err
	}
	if conf.Polar.RPCTxFeeCap, err =
		parser.GetFloat64(flagRPCTxFeeCap); err != nil {
		return nil, err
	}

	// Polar Miner settings
	if conf.Polar.Miner.Etherbase, err =
		parser.GetCommonAddress(flagMinerEtherbase); err != nil {
		return nil, err
	}
	if conf.Polar.Miner.ExtraData, err =
		parser.GetHexutilBytes(flagMinerExtraData); err != nil {
		return nil, err
	}
	if conf.Polar.Miner.GasFloor, err =
		parser.GetUint64(flagMinerGasFloor); err != nil {
		return nil, err
	}
	if conf.Polar.Miner.GasCeil, err =
		parser.GetUint64(flagMinerGasCeil); err != nil {
		return nil, err
	}
	if conf.Polar.Miner.GasPrice, err =
		parser.GetBigInt(flagMinerGasPrice); err != nil {
		return nil, err
	}
	if conf.Polar.Miner.Recommit, err =
		parser.GetTimeDuration(flagMinerRecommit); err != nil {
		return nil, err
	}

	if conf.Polar.Miner.NewPayloadTimeout, err =
		parser.GetTimeDuration(flagMinerNewPayloadTimeout); err != nil {
		return nil, err
	}

	// Polar Chain settings
	if conf.Polar.Chain.ChainID, err =
		parser.GetBigInt(flagChainID); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.HomesteadBlock, err =
		parser.GetBigInt(flagHomesteadBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.DAOForkBlock, err =
		parser.GetBigInt(flagDAOForkBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.DAOForkSupport, err =
		parser.GetBool(flagDAOForkSupport); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.EIP150Block, err =
		parser.GetBigInt(flagEIP150Block); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.EIP155Block, err =
		parser.GetBigInt(flagEIP155Block); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.EIP158Block, err =
		parser.GetBigInt(flagEIP158Block); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.ByzantiumBlock, err =
		parser.GetBigInt(flagByzantiumBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.ConstantinopleBlock, err =
		parser.GetBigInt(flagConstantinopleBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.PetersburgBlock, err =
		parser.GetBigInt(flagPetersburgBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.IstanbulBlock, err =
		parser.GetBigInt(flagIstanbulBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.MuirGlacierBlock, err =
		parser.GetBigInt(flagMuirGlacierBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.BerlinBlock, err =
		parser.GetBigInt(flagBerlinBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.LondonBlock, err =
		parser.GetBigInt(flagLondonBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.ArrowGlacierBlock, err =
		parser.GetBigInt(flagArrowGlacierBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.GrayGlacierBlock, err =
		parser.GetBigInt(flagGrayGlacierBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.MergeNetsplitBlock, err =
		parser.GetBigInt(flagMergeNetsplitBlock); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.ShanghaiTime, err =
		parser.GetUint64Ptr(flagShanghaiTime); err != nil {
		return nil, err
	}

	if conf.Polar.Chain.CancunTime, err =
		parser.GetUint64Ptr(flagCancunTime); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.PragueTime, err =
		parser.GetUint64Ptr(flagPragueTime); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.VerkleTime, err =
		parser.GetUint64Ptr(flagVerkleTime); err != nil {
		return nil, err
	}

	if conf.Polar.Chain.TerminalTotalDifficulty, err =
		parser.GetBigInt(
			flagTerminalTotalDifficulty); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.TerminalTotalDifficultyPassed, err =
		parser.GetBool(
			flagTerminalTotalDifficultyPassed); err != nil {
		return nil, err
	}
	if conf.Polar.Chain.IsDevMode, err =
		parser.GetBool(flagIsDevMode); err != nil {
		return nil, err
	}

	// Polar.GPO settings
	if conf.Polar.GPO.Blocks, err =
		parser.GetInt(flagBlocks); err != nil {
		return nil, err
	}
	if conf.Polar.GPO.Percentile, err =
		parser.GetInt(flagPercentile); err != nil {
		return nil, err
	}
	if conf.Polar.GPO.MaxHeaderHistory, err =
		parser.GetUint64(flagMaxHeaderHistory); err != nil {
		return nil, err
	}
	if conf.Polar.GPO.MaxBlockHistory, err =
		parser.GetUint64(flagMaxBlockHistory); err != nil {
		return nil, err
	}
	if val, err =
		parser.GetInt64(flagDefault); err != nil {
		return nil, err
	}
	conf.Polar.GPO.Default = big.NewInt(val)

	if val, err =
		parser.GetInt64(flagMaxPrice); err != nil {
		return nil, err
	}
	conf.Polar.GPO.MaxPrice = big.NewInt(val)

	if val, err =
		parser.GetInt64(flagIgnorePrice); err != nil {
		return nil, err
	}
	conf.Polar.GPO.IgnorePrice = big.NewInt(val)

	// LegacyPool
	if conf.Polar.LegacyTxPool.Locals, err =
		parser.GetCommonAddressList(flagLocals); err != nil {
		return nil, err
	}
	if conf.Polar.LegacyTxPool.NoLocals, err =
		parser.GetBool(flagNoLocals); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.Journal, err =
		parser.GetString(flagJournal); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.Rejournal, err =
		parser.GetTimeDuration(flagReJournal); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.PriceLimit, err =
		parser.GetUint64(flagPriceLimit); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.PriceBump, err =
		parser.GetUint64(flagPriceBump); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.AccountSlots, err =
		parser.GetUint64(flagAccountSlots); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.GlobalSlots, err =
		parser.GetUint64(flagGlobalSlots); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.AccountQueue, err =
		parser.GetUint64(flagAccountQueue); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.GlobalQueue, err =
		parser.GetUint64(flagGlobalQueue); err != nil {
		return nil, err
	}

	if conf.Polar.LegacyTxPool.Lifetime, err =
		parser.GetTimeDuration(flagLifetime); err != nil {
		return nil, err
	}

	// Node settings
	if conf.Node.Name, err =
		parser.GetString(flagName); err != nil {
		return nil, err
	}
	if conf.Node.UserIdent, err =
		parser.GetString(flagUserIdent); err != nil {
		return nil, err
	}
	if conf.Node.Version, err =
		parser.GetString(flagVersion); err != nil {
		return nil, err
	}
	if conf.Node.DataDir, err =
		parser.GetString(flagDataDir); err != nil {
		return nil, err
	}
	if conf.Node.DataDir == "" {
		conf.Node.DataDir, err =
			parser.GetString(flags.FlagHome)
		if err != nil {
			return nil, err
		}
	}
	if conf.Node.KeyStoreDir, err =
		parser.GetString(flagKeyStoreDir); err != nil {
		return nil, err
	}
	if conf.Node.ExternalSigner, err =
		parser.GetString(flagExternalSigner); err != nil {
		return nil, err
	}
	if conf.Node.UseLightweightKDF, err =
		parser.GetBool(flagUseLightweightKdf); err != nil {
		return nil, err
	}
	if conf.Node.InsecureUnlockAllowed, err =
		parser.GetBool(flagInsecureUnlockAllowed); err != nil {
		return nil, err
	}
	if conf.Node.USB, err =
		parser.GetBool(flagUsb); err != nil {
		return nil, err
	}
	if conf.Node.SmartCardDaemonPath, err =
		parser.GetString(flagSmartCardDaemonPath); err != nil {
		return nil, err
	}
	if conf.Node.IPCPath, err =
		parser.GetString(flagIpcPath); err != nil {
		return nil, err
	}
	if conf.Node.HTTPHost, err =
		parser.GetString(flagHTTPHost); err != nil {
		return nil, err
	}
	if conf.Node.HTTPPort, err =
		parser.GetInt(flagHTTPPort); err != nil {
		return nil, err
	}
	if conf.Node.HTTPCors, err =
		parser.GetStringSlice(flagHTTPCors); err != nil {
		return nil, err
	}
	if conf.Node.HTTPVirtualHosts, err =
		parser.GetStringSlice(flagHTTPVirtualHosts); err != nil {
		return nil, err
	}
	if conf.Node.HTTPModules, err =
		parser.GetStringSlice(flagHTTPModules); err != nil {
		return nil, err
	}
	if conf.Node.HTTPPathPrefix, err =
		parser.GetString(flagHTTPPathPrefix); err != nil {
		return nil, err
	}
	if conf.Node.AuthAddr, err =
		parser.GetString(flagAuthAddr); err != nil {
		return nil, err
	}
	if conf.Node.AuthPort, err =
		parser.GetInt(flagAuthPort); err != nil {
		return nil, err
	}
	if conf.Node.AuthVirtualHosts, err =
		parser.GetStringSlice(flagAuthVirtualHosts); err != nil {
		return nil, err
	}
	if conf.Node.WSHost, err =
		parser.GetString(flagWsHost); err != nil {
		return nil, err
	}
	if conf.Node.WSPort, err =
		parser.GetInt(flagWsPort); err != nil {
		return nil, err
	}
	if conf.Node.WSPathPrefix, err =
		parser.GetString(flagWsPathPrefix); err != nil {
		return nil, err
	}
	if conf.Node.WSOrigins, err =
		parser.GetStringSlice(flagWsOrigins); err != nil {
		return nil, err
	}
	if conf.Node.WSModules, err =
		parser.GetStringSlice(flagWsModules); err != nil {
		return nil, err
	}
	if conf.Node.WSExposeAll, err =
		parser.GetBool(flagWsExposeAll); err != nil {
		return nil, err
	}
	if conf.Node.GraphQLCors, err =
		parser.GetStringSlice(flagGraphqlCors); err != nil {
		return nil, err
	}
	if conf.Node.GraphQLVirtualHosts, err =
		parser.GetStringSlice(flagGraphqlVirtualHosts); err != nil {
		return nil, err
	}
	if conf.Node.AllowUnprotectedTxs, err =
		parser.GetBool(flagAllowUnprotectedTxs); err != nil {
		return nil, err
	}
	if conf.Node.BatchRequestLimit, err =
		parser.GetInt(flagBatchRequestLimit); err != nil {
		return nil, err
	}
	if conf.Node.BatchResponseMaxSize, err =
		parser.GetInt(flagBatchResponseMaxSize); err != nil {
		return nil, err
	}
	if conf.Node.JWTSecret, err =
		parser.GetString(flagJwtSecret); err != nil {
		return nil, err
	}
	if conf.Node.DBEngine, err =
		parser.GetString(flagDBEngine); err != nil {
		return nil, err
	}

	// Node.HTTPTimeouts settings
	if conf.Node.HTTPTimeouts.ReadTimeout, err =
		parser.GetTimeDuration(flagReadTimeout); err != nil {
		return nil, err
	}
	if conf.Node.HTTPTimeouts.ReadHeaderTimeout, err =
		parser.GetTimeDuration(
			flagReadHeaderTimeout); err != nil {
		return nil, err
	}
	if conf.Node.HTTPTimeouts.WriteTimeout, err =
		parser.GetTimeDuration(flagWriteTimeout); err != nil {
		return nil, err
	}
	if conf.Node.HTTPTimeouts.IdleTimeout, err =
		parser.GetTimeDuration(flagIdleTimeout); err != nil {
		return nil, err
	}

	return conf, nil
}
