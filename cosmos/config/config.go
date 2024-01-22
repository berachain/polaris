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
	"fmt"
	"math/big"

	"github.com/berachain/polaris/cosmos/config/flags"
	"github.com/berachain/polaris/eth"
	"github.com/berachain/polaris/eth/accounts"

	sdkflags "github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	version "github.com/cosmos/cosmos-sdk/version"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Config = eth.Config

// SetupCosmosConfig sets up the Cosmos SDK configuration to be compatible with the
// semantics of etheruem.
func SetupCosmosConfig() {
	// set the address prefixes
	config := sdk.GetConfig()

	// We use CoinType == 60 to match Ethereum.
	// This is not strictly necessary, though highly recommended.
	config.SetCoinType(accounts.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)
	config.Seal()
}

// MustReadConfigFromAppOpts reads the configuration options from the given
// application options. Panics if the configuration cannot be read.
func MustReadConfigFromAppOpts(opts servertypes.AppOptions) *Config {
	cfg, err := ReadConfigFromAppOpts(opts)
	if err != nil {
		panic(err)
	}
	return cfg
}

// ReadConfigFromAppOpts reads the configuration options from the given
// application options.
func ReadConfigFromAppOpts(opts servertypes.AppOptions) (*Config, error) {
	return readConfigFromAppOptsParser(AppOptionsParser{AppOptions: opts})
}

func readConfigFromAppOptsParser(parser AppOptionsParser) (*Config, error) {
	var (
		err  error
		conf = &Config{}
	)

	// 🅱️onad mode.
	if conf.OptimisticExecution, err = parser.GetBool(flags.OptimisticExecution); err != nil {
		return nil, err
	}

	if err = readPolarisCoreSettings(parser, conf); err != nil {
		return nil, err
	}

	if err = readPolarisMinterSettings(parser, conf); err != nil {
		return nil, err
	}

	if err = readPolarisChainSettings(parser, conf); err != nil {
		return nil, err
	}

	if err = readPolarisGPOSettings(parser, conf); err != nil {
		return nil, err
	}

	if err = readPolarisLegacyPool(parser, conf); err != nil {
		return nil, err
	}

	if err = readNodeSettings(parser, conf); err != nil {
		return nil, err
	}

	if err = readNodeHTTPTimeoutsSettings(parser, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func readPolarisCoreSettings(parser AppOptionsParser, conf *Config) error {
	var err error

	if conf.Polar.RPCGasCap, err =
		parser.GetUint64(flags.RPCGasCap); err != nil {
		return err
	}
	if conf.Polar.RPCEVMTimeout, err =
		parser.GetTimeDuration(flags.RPCEvmTimeout); err != nil {
		return err
	}
	if conf.Polar.RPCTxFeeCap, err =
		parser.GetFloat64(flags.RPCTxFeeCap); err != nil {
		return err
	}

	return nil
}

func readPolarisMinterSettings(parser AppOptionsParser, conf *Config) error {
	var err error
	if conf.Polar.Miner.Etherbase, err =
		parser.GetCommonAddress(flags.MinerEtherbase); err != nil {
		return err
	}
	if conf.Polar.Miner.ExtraData, err =
		parser.GetHexutilBytes(flags.MinerExtraData); err != nil {
		return err
	}

	if len(conf.Polar.Miner.ExtraData) == 0 {
		commit := version.NewInfo().GitCommit
		if len(commit) != 40 { //nolint:gomnd // its okay.
			return fmt.Errorf("invalid git commit length: %d", len(commit))
		}
		conf.Polar.Miner.ExtraData = hexutil.Bytes(
			commit[32:40],
		)
	}

	if conf.Polar.Miner.GasFloor, err =
		parser.GetUint64(flags.MinerGasFloor); err != nil {
		return err
	}
	if conf.Polar.Miner.GasCeil, err =
		parser.GetUint64(flags.MinerGasCeil); err != nil {
		return err
	}
	if conf.Polar.Miner.GasPrice, err =
		parser.GetBigInt(flags.MinerGasPrice); err != nil {
		return err
	}
	if conf.Polar.Miner.Recommit, err =
		parser.GetTimeDuration(flags.MinerRecommit); err != nil {
		return err
	}

	if conf.Polar.Miner.NewPayloadTimeout, err =
		parser.GetTimeDuration(flags.MinerNewPayloadTimeout); err != nil {
		return err
	}
	return nil
}

//nolint:gocognit
func readPolarisChainSettings(parser AppOptionsParser, conf *Config) error {
	var err error
	if conf.Polar.Chain.ChainID, err =
		parser.GetBigInt(flags.ChainID); err != nil {
		return err
	}
	if conf.Polar.Chain.HomesteadBlock, err =
		parser.GetBigInt(flags.HomesteadBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.DAOForkBlock, err =
		parser.GetBigInt(flags.DAOForkBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.DAOForkSupport, err =
		parser.GetBool(flags.DAOForkSupport); err != nil {
		return err
	}
	if conf.Polar.Chain.EIP150Block, err =
		parser.GetBigInt(flags.EIP150Block); err != nil {
		return err
	}
	if conf.Polar.Chain.EIP155Block, err =
		parser.GetBigInt(flags.EIP155Block); err != nil {
		return err
	}
	if conf.Polar.Chain.EIP158Block, err =
		parser.GetBigInt(flags.EIP158Block); err != nil {
		return err
	}
	if conf.Polar.Chain.ByzantiumBlock, err =
		parser.GetBigInt(flags.ByzantiumBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.ConstantinopleBlock, err =
		parser.GetBigInt(flags.ConstantinopleBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.PetersburgBlock, err =
		parser.GetBigInt(flags.PetersburgBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.IstanbulBlock, err =
		parser.GetBigInt(flags.IstanbulBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.MuirGlacierBlock, err =
		parser.GetBigInt(flags.MuirGlacierBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.BerlinBlock, err =
		parser.GetBigInt(flags.BerlinBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.LondonBlock, err =
		parser.GetBigInt(flags.LondonBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.ArrowGlacierBlock, err =
		parser.GetBigInt(flags.ArrowGlacierBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.GrayGlacierBlock, err =
		parser.GetBigInt(flags.GrayGlacierBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.MergeNetsplitBlock, err =
		parser.GetBigInt(flags.MergeNetsplitBlock); err != nil {
		return err
	}
	if conf.Polar.Chain.ShanghaiTime, err =
		parser.GetUint64Ptr(flags.ShanghaiTime); err != nil {
		return err
	}

	if conf.Polar.Chain.CancunTime, err =
		parser.GetUint64Ptr(flags.CancunTime); err != nil {
		return err
	}
	if conf.Polar.Chain.PragueTime, err =
		parser.GetUint64Ptr(flags.PragueTime); err != nil {
		return err
	}
	if conf.Polar.Chain.VerkleTime, err =
		parser.GetUint64Ptr(flags.VerkleTime); err != nil {
		return err
	}

	if conf.Polar.Chain.TerminalTotalDifficulty, err =
		parser.GetBigInt(
			flags.TerminalTotalDifficulty); err != nil {
		return err
	}

	if conf.Polar.Chain.TerminalTotalDifficultyPassed, err =
		parser.GetBool(
			flags.TerminalTotalDifficultyPassed); err != nil {
		return err
	}

	return nil
}

func readPolarisGPOSettings(parser AppOptionsParser, conf *Config) error {
	var (
		err error
		val int64
	)

	if conf.Polar.GPO.Blocks, err =
		parser.GetInt(flags.Blocks); err != nil {
		return err
	}
	if conf.Polar.GPO.Percentile, err =
		parser.GetInt(flags.Percentile); err != nil {
		return err
	}
	if conf.Polar.GPO.MaxHeaderHistory, err =
		parser.GetUint64(flags.MaxHeaderHistory); err != nil {
		return err
	}
	if conf.Polar.GPO.MaxBlockHistory, err =
		parser.GetUint64(flags.MaxBlockHistory); err != nil {
		return err
	}
	if val, err =
		parser.GetInt64(flags.Default); err != nil {
		return err
	}
	conf.Polar.GPO.Default = big.NewInt(val)

	if val, err =
		parser.GetInt64(flags.MaxPrice); err != nil {
		return err
	}
	conf.Polar.GPO.MaxPrice = big.NewInt(val)

	if val, err =
		parser.GetInt64(flags.IgnorePrice); err != nil {
		return err
	}
	conf.Polar.GPO.IgnorePrice = big.NewInt(val)

	return nil
}

func readPolarisLegacyPool(parser AppOptionsParser, conf *Config) error {
	var err error

	if conf.Polar.LegacyTxPool.Locals, err =
		parser.GetCommonAddressList(flags.Locals); err != nil {
		return err
	}
	if conf.Polar.LegacyTxPool.NoLocals, err =
		parser.GetBool(flags.NoLocals); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.Journal, err =
		parser.GetString(flags.Journal); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.Journal == "" {
		conf.Polar.LegacyTxPool.Journal, err =
			parser.GetString(sdkflags.FlagHome)
		if err != nil {
			return err
		}
		conf.Polar.LegacyTxPool.Journal += "/data/transactions.rlp"
	}

	if conf.Polar.LegacyTxPool.Rejournal, err =
		parser.GetTimeDuration(flags.ReJournal); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.PriceLimit, err =
		parser.GetUint64(flags.PriceLimit); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.PriceBump, err =
		parser.GetUint64(flags.PriceBump); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.AccountSlots, err =
		parser.GetUint64(flags.AccountSlots); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.GlobalSlots, err =
		parser.GetUint64(flags.GlobalSlots); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.AccountQueue, err =
		parser.GetUint64(flags.AccountQueue); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.GlobalQueue, err =
		parser.GetUint64(flags.GlobalQueue); err != nil {
		return err
	}

	if conf.Polar.LegacyTxPool.Lifetime, err =
		parser.GetTimeDuration(flags.Lifetime); err != nil {
		return err
	}

	return err
}

//nolint:gocognit,funlen,cyclop,gocyclo // its okay.
func readNodeSettings(parser AppOptionsParser, conf *Config) error {
	var err error

	if conf.Node.Name, err =
		parser.GetString(flags.Name); err != nil {
		return err
	}
	if conf.Node.UserIdent, err =
		parser.GetString(flags.UserIdent); err != nil {
		return err
	}
	if conf.Node.Version, err =
		parser.GetString(flags.Version); err != nil {
		return err
	}
	if conf.Node.DataDir, err =
		parser.GetString(flags.DataDir); err != nil {
		return err
	}
	if conf.Node.DataDir == "" {
		conf.Node.DataDir, err =
			parser.GetString(sdkflags.FlagHome)
		if err != nil {
			return err
		}
	}
	if conf.Node.KeyStoreDir, err =
		parser.GetString(flags.KeyStoreDir); err != nil {
		return err
	}
	if conf.Node.ExternalSigner, err =
		parser.GetString(flags.ExternalSigner); err != nil {
		return err
	}
	if conf.Node.UseLightweightKDF, err =
		parser.GetBool(flags.UseLightweightKdf); err != nil {
		return err
	}
	if conf.Node.InsecureUnlockAllowed, err =
		parser.GetBool(flags.InsecureUnlockAllowed); err != nil {
		return err
	}
	if conf.Node.USB, err =
		parser.GetBool(flags.Usb); err != nil {
		return err
	}
	if conf.Node.SmartCardDaemonPath, err =
		parser.GetString(flags.SmartCardDaemonPath); err != nil {
		return err
	}
	if conf.Node.IPCPath, err =
		parser.GetString(flags.IpcPath); err != nil {
		return err
	}
	if conf.Node.HTTPHost, err =
		parser.GetString(flags.HTTPHost); err != nil {
		return err
	}
	if conf.Node.HTTPPort, err =
		parser.GetInt(flags.HTTPPort); err != nil {
		return err
	}
	if conf.Node.HTTPCors, err =
		parser.GetStringSlice(flags.HTTPCors); err != nil {
		return err
	}
	if conf.Node.HTTPVirtualHosts, err =
		parser.GetStringSlice(flags.HTTPVirtualHosts); err != nil {
		return err
	}
	if conf.Node.HTTPModules, err =
		parser.GetStringSlice(flags.HTTPModules); err != nil {
		return err
	}
	if conf.Node.HTTPPathPrefix, err =
		parser.GetString(flags.HTTPPathPrefix); err != nil {
		return err
	}
	if conf.Node.AuthAddr, err =
		parser.GetString(flags.AuthAddr); err != nil {
		return err
	}
	if conf.Node.AuthPort, err =
		parser.GetInt(flags.AuthPort); err != nil {
		return err
	}
	if conf.Node.AuthVirtualHosts, err =
		parser.GetStringSlice(flags.AuthVirtualHosts); err != nil {
		return err
	}
	if conf.Node.WSHost, err =
		parser.GetString(flags.WsHost); err != nil {
		return err
	}
	if conf.Node.WSPort, err =
		parser.GetInt(flags.WsPort); err != nil {
		return err
	}
	if conf.Node.WSPathPrefix, err =
		parser.GetString(flags.WsPathPrefix); err != nil {
		return err
	}
	if conf.Node.WSOrigins, err =
		parser.GetStringSlice(flags.WsOrigins); err != nil {
		return err
	}
	if conf.Node.WSModules, err =
		parser.GetStringSlice(flags.WsModules); err != nil {
		return err
	}
	if conf.Node.WSExposeAll, err =
		parser.GetBool(flags.WsExposeAll); err != nil {
		return err
	}
	if conf.Node.GraphQLCors, err =
		parser.GetStringSlice(flags.GraphqlCors); err != nil {
		return err
	}
	if conf.Node.GraphQLVirtualHosts, err =
		parser.GetStringSlice(flags.GraphqlVirtualHosts); err != nil {
		return err
	}
	if conf.Node.AllowUnprotectedTxs, err =
		parser.GetBool(flags.AllowUnprotectedTxs); err != nil {
		return err
	}
	if conf.Node.BatchRequestLimit, err =
		parser.GetInt(flags.BatchRequestLimit); err != nil {
		return err
	}
	if conf.Node.BatchResponseMaxSize, err =
		parser.GetInt(flags.BatchResponseMaxSize); err != nil {
		return err
	}
	if conf.Node.JWTSecret, err =
		parser.GetString(flags.JwtSecret); err != nil {
		return err
	}
	if conf.Node.DBEngine, err =
		parser.GetString(flags.DBEngine); err != nil {
		return err
	}

	return nil
}

func readNodeHTTPTimeoutsSettings(parser AppOptionsParser, conf *Config) error {
	var err error

	if conf.Node.HTTPTimeouts.ReadTimeout, err =
		parser.GetTimeDuration(flags.ReadTimeout); err != nil {
		return err
	}
	if conf.Node.HTTPTimeouts.ReadHeaderTimeout, err =
		parser.GetTimeDuration(
			flags.ReadHeaderTimeout); err != nil {
		return err
	}
	if conf.Node.HTTPTimeouts.WriteTimeout, err =
		parser.GetTimeDuration(flags.WriteTimeout); err != nil {
		return err
	}
	if conf.Node.HTTPTimeouts.IdleTimeout, err =
		parser.GetTimeDuration(flags.IdleTimeout); err != nil {
		return err
	}

	return nil
}
