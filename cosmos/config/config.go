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
	"time"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/common"
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

//nolint:funlen,gocognit,gocyclo,cyclop // TODO break up later.
func ReadConfigFromAppOpts(opts servertypes.AppOptions) (*Config, error) {
	var err error
	var val int64
	conf := &Config{}

	// Define little error handler.
	var handleError = func(err error, flag string) error {
		if err != nil {
			return fmt.Errorf("error while reading configuration: %w flag: %s", err, flag)
		}
		return nil
	}

	// Wrapping casting functions to return both value and error
	getCommonAddressList := func(key string) []common.Address {
		addresses := make([]common.Address, 0)
		addressStrs := cast.ToStringSlice(opts.Get(key))
		for _, addressStr := range addressStrs {
			address := common.HexToAddress(addressStr)
			addresses = append(addresses, address)
		}
		return addresses
	}
	getString := func(key string) (string, error) { return cast.ToStringE(opts.Get(key)) }
	getInt := func(key string) (int, error) { return cast.ToIntE(opts.Get(key)) }
	getInt64 := func(key string) (int64, error) { return cast.ToInt64E(opts.Get(key)) }
	getUint64 := func(key string) (uint64, error) { return cast.ToUint64E(opts.Get(key)) }
	getUint64Ptr := func(key string) (*uint64, error) {
		if num, _ := cast.ToStringE((opts.Get(key))); num == "<nil>" || num == "null" {
			return nil, nil //nolint:nilnil // intentional.
		}

		num, _err := cast.ToUint64E((opts.Get(key)))
		if _err != nil {
			return nil, _err
		}
		return &num, nil
	}
	getBigInt := func(key string) (*big.Int, error) {
		str, _err := cast.ToStringE(opts.Get(key))
		if str, _ := cast.ToStringE((opts.Get(key))); str == "<nil>" || str == "null" {
			return nil, nil //nolint:nilnil // intentional.
		}
		if _err != nil {
			return nil, _err
		}
		num, ok := new(big.Int).SetString(str, 10) //nolint:gomnd // base 10.
		if !ok {
			return nil, fmt.Errorf("invalid big.Int string: %s", str)
		}
		return num, nil
	}
	getFloat64 := func(key string) (float64, error) { return cast.ToFloat64E(opts.Get(key)) }
	getBool := func(key string) (bool, error) { return cast.ToBoolE(opts.Get(key)) }
	getStringSlice := func(key string) ([]string, error) { return cast.ToStringSliceE(opts.Get(key)) }
	getTimeDuration := func(key string) (time.Duration, error) { return cast.ToDurationE(opts.Get(key)) }

	// Polar settings
	if conf.Polar.RPCGasCap, err = getUint64(flagRPCGasCap); err != nil {
		return nil, handleError(err, flagRPCGasCap)
	}
	if conf.Polar.RPCEVMTimeout, err = getTimeDuration(flagRPCEvmTimeout); err != nil {
		return nil, handleError(err, flagRPCEvmTimeout)
	}
	if conf.Polar.RPCTxFeeCap, err = getFloat64(flagRPCTxFeeCap); err != nil {
		return nil, handleError(err, flagRPCTxFeeCap)
	}

	// Polar Chain settings
	if conf.Polar.Chain.ChainID, err = getBigInt(flagChainID); err != nil {
		return nil, handleError(err, flagChainID)
	}
	if conf.Polar.Chain.HomesteadBlock, err = getBigInt(flagHomesteadBlock); err != nil {
		return nil, handleError(err, flagHomesteadBlock)
	}
	if conf.Polar.Chain.DAOForkBlock, err = getBigInt(flagDAOForkBlock); err != nil {
		return nil, handleError(err, flagDAOForkBlock)
	}
	if conf.Polar.Chain.DAOForkSupport, err = getBool(flagDAOForkSupport); err != nil {
		return nil, handleError(err, flagDAOForkSupport)
	}
	if conf.Polar.Chain.EIP150Block, err = getBigInt(flagEIP150Block); err != nil {
		return nil, handleError(err, flagEIP150Block)
	}
	if conf.Polar.Chain.EIP155Block, err = getBigInt(flagEIP155Block); err != nil {
		return nil, handleError(err, flagEIP155Block)
	}
	if conf.Polar.Chain.EIP158Block, err = getBigInt(flagEIP158Block); err != nil {
		return nil, handleError(err, flagEIP158Block)
	}
	if conf.Polar.Chain.ByzantiumBlock, err = getBigInt(flagByzantiumBlock); err != nil {
		return nil, handleError(err, flagByzantiumBlock)
	}
	if conf.Polar.Chain.ConstantinopleBlock, err = getBigInt(flagConstantinopleBlock); err != nil {
		return nil, handleError(err, flagConstantinopleBlock)
	}
	if conf.Polar.Chain.PetersburgBlock, err = getBigInt(flagPetersburgBlock); err != nil {
		return nil, handleError(err, flagPetersburgBlock)
	}
	if conf.Polar.Chain.IstanbulBlock, err = getBigInt(flagIstanbulBlock); err != nil {
		return nil, handleError(err, flagIstanbulBlock)
	}
	if conf.Polar.Chain.MuirGlacierBlock, err = getBigInt(flagMuirGlacierBlock); err != nil {
		return nil, handleError(err, flagMuirGlacierBlock)
	}
	if conf.Polar.Chain.BerlinBlock, err = getBigInt(flagBerlinBlock); err != nil {
		return nil, handleError(err, flagBerlinBlock)
	}
	if conf.Polar.Chain.LondonBlock, err = getBigInt(flagLondonBlock); err != nil {
		return nil, handleError(err, flagLondonBlock)
	}
	if conf.Polar.Chain.ArrowGlacierBlock, err = getBigInt(flagArrowGlacierBlock); err != nil {
		return nil, handleError(err, flagArrowGlacierBlock)
	}
	if conf.Polar.Chain.GrayGlacierBlock, err = getBigInt(flagGrayGlacierBlock); err != nil {
		return nil, handleError(err, flagGrayGlacierBlock)
	}
	if conf.Polar.Chain.MergeNetsplitBlock, err = getBigInt(flagMergeNetsplitBlock); err != nil {
		return nil, handleError(err, flagMergeNetsplitBlock)
	}
	if conf.Polar.Chain.ShanghaiTime, err = getUint64Ptr(flagShanghaiTime); err != nil {
		return nil, handleError(err, flagShanghaiTime)
	}

	fmt.Println("SHANGHAI TIME SHOULD BE NIL", conf.Polar.Chain.ShanghaiTime)

	if conf.Polar.Chain.CancunTime, err = getUint64Ptr(flagCancunTime); err != nil {
		return nil, handleError(err, flagCancunTime)
	}
	if conf.Polar.Chain.PragueTime, err = getUint64Ptr(flagPragueTime); err != nil {
		return nil, handleError(err, flagPragueTime)
	}
	if conf.Polar.Chain.VerkleTime, err = getUint64Ptr(flagVerkleTime); err != nil {
		return nil, handleError(err, flagVerkleTime)
	}

	if conf.Polar.Chain.TerminalTotalDifficulty, err = getBigInt(flagTerminalTotalDifficulty); err != nil {
		return nil, handleError(err, flagTerminalTotalDifficulty)
	}
	if conf.Polar.Chain.TerminalTotalDifficultyPassed, err = getBool(flagTerminalTotalDifficultyPassed); err != nil {
		return nil, handleError(err, flagTerminalTotalDifficultyPassed)
	}
	if conf.Polar.Chain.IsDevMode, err = getBool(flagIsDevMode); err != nil {
		return nil, handleError(err, flagIsDevMode)
	}

	// Polar.GPO settings
	if conf.Polar.GPO.Blocks, err = getInt(flagBlocks); err != nil {
		return nil, handleError(err, flagBlocks)
	}
	if conf.Polar.GPO.Percentile, err = getInt(flagPercentile); err != nil {
		return nil, handleError(err, flagPercentile)
	}
	if conf.Polar.GPO.MaxHeaderHistory, err = getUint64(flagMaxHeaderHistory); err != nil {
		return nil, handleError(err, flagMaxHeaderHistory)
	}
	if conf.Polar.GPO.MaxBlockHistory, err = getUint64(flagMaxBlockHistory); err != nil {
		return nil, handleError(err, flagMaxBlockHistory)
	}
	if val, err = getInt64(flagDefault); err != nil {
		return nil, handleError(err, flagDefault)
	}
	conf.Polar.GPO.Default = big.NewInt(val)

	if val, err = getInt64(flagMaxPrice); err != nil {
		return nil, handleError(err, flagMaxPrice)
	}
	conf.Polar.GPO.MaxPrice = big.NewInt(val)

	if val, err = getInt64(flagIgnorePrice); err != nil {
		return nil, handleError(err, flagIgnorePrice)
	}
	conf.Polar.GPO.IgnorePrice = big.NewInt(val)

	// LegacyPool
	conf.Polar.LegacyTxPool.Locals = getCommonAddressList(flagDefault)

	if conf.Polar.LegacyTxPool.NoLocals, err = getBool(flagNoLocals); err != nil {
		return nil, handleError(err, flagNoLocals)
	}

	if conf.Polar.LegacyTxPool.Journal, err = getString(flagJournal); err != nil {
		return nil, handleError(err, flagJournal)
	}

	if conf.Polar.LegacyTxPool.Rejournal, err = getTimeDuration(flagReJournal); err != nil {
		return nil, handleError(err, flagReJournal)
	}

	if conf.Polar.LegacyTxPool.PriceLimit, err = getUint64(flagPriceLimit); err != nil {
		return nil, handleError(err, flagPriceLimit)
	}

	if conf.Polar.LegacyTxPool.PriceBump, err = getUint64(flagPriceBump); err != nil {
		return nil, handleError(err, flagPriceBump)
	}

	if conf.Polar.LegacyTxPool.AccountSlots, err = getUint64(flagAccountSlots); err != nil {
		return nil, handleError(err, flagAccountSlots)
	}

	if conf.Polar.LegacyTxPool.GlobalSlots, err = getUint64(flagGlobalSlots); err != nil {
		return nil, handleError(err, flagGlobalSlots)
	}

	if conf.Polar.LegacyTxPool.AccountQueue, err = getUint64(flagAccountQueue); err != nil {
		return nil, handleError(err, flagAccountQueue)
	}

	if conf.Polar.LegacyTxPool.GlobalQueue, err = getUint64(flagGlobalQueue); err != nil {
		return nil, handleError(err, flagGlobalQueue)
	}

	if conf.Polar.LegacyTxPool.Lifetime, err = getTimeDuration(flagLifetime); err != nil {
		return nil, handleError(err, flagLifetime)
	}

	// Node settings
	if conf.Node.Name, err = getString(flagName); err != nil {
		return nil, handleError(err, flagName)
	}
	if conf.Node.UserIdent, err = getString(flagUserIdent); err != nil {
		return nil, handleError(err, flagUserIdent)
	}
	if conf.Node.Version, err = getString(flagVersion); err != nil {
		return nil, handleError(err, flagVersion)
	}
	if conf.Node.DataDir, err = getString(flagDataDir); err != nil {
		return nil, handleError(err, flagDataDir)
	}
	if conf.Node.DataDir == "" {
		conf.Node.DataDir, err = getString(flags.FlagHome)
		if err != nil {
			return nil, handleError(err, flags.FlagHome)
		}
	}
	if conf.Node.KeyStoreDir, err = getString(flagKeyStoreDir); err != nil {
		return nil, handleError(err, flagKeyStoreDir)
	}
	if conf.Node.ExternalSigner, err = getString(flagExternalSigner); err != nil {
		return nil, handleError(err, flagExternalSigner)
	}
	if conf.Node.UseLightweightKDF, err = getBool(flagUseLightweightKdf); err != nil {
		return nil, handleError(err, flagUseLightweightKdf)
	}
	if conf.Node.InsecureUnlockAllowed, err = getBool(flagInsecureUnlockAllowed); err != nil {
		return nil, handleError(err, flagInsecureUnlockAllowed)
	}
	if conf.Node.USB, err = getBool(flagUsb); err != nil {
		return nil, handleError(err, flagUsb)
	}
	if conf.Node.SmartCardDaemonPath, err = getString(flagSmartCardDaemonPath); err != nil {
		return nil, handleError(err, flagSmartCardDaemonPath)
	}
	if conf.Node.IPCPath, err = getString(flagIpcPath); err != nil {
		return nil, handleError(err, flagIpcPath)
	}
	if conf.Node.HTTPHost, err = getString(flagHTTPHost); err != nil {
		return nil, handleError(err, flagHTTPHost)
	}
	if conf.Node.HTTPPort, err = getInt(flagHTTPPort); err != nil {
		return nil, handleError(err, flagHTTPPort)
	}
	if conf.Node.HTTPCors, err = getStringSlice(flagHTTPCors); err != nil {
		return nil, handleError(err, flagHTTPCors)
	}
	if conf.Node.HTTPVirtualHosts, err = getStringSlice(flagHTTPVirtualHosts); err != nil {
		return nil, handleError(err, flagHTTPVirtualHosts)
	}
	if conf.Node.HTTPModules, err = getStringSlice(flagHTTPModules); err != nil {
		return nil, handleError(err, flagHTTPModules)
	}
	if conf.Node.HTTPPathPrefix, err = getString(flagHTTPPathPrefix); err != nil {
		return nil, handleError(err, flagHTTPPathPrefix)
	}
	if conf.Node.AuthAddr, err = getString(flagAuthAddr); err != nil {
		return nil, handleError(err, flagAuthAddr)
	}
	if conf.Node.AuthPort, err = getInt(flagAuthPort); err != nil {
		return nil, handleError(err, flagAuthPort)
	}
	if conf.Node.AuthVirtualHosts, err = getStringSlice(flagAuthVirtualHosts); err != nil {
		return nil, handleError(err, flagAuthVirtualHosts)
	}
	if conf.Node.WSHost, err = getString(flagWsHost); err != nil {
		return nil, handleError(err, flagWsHost)
	}
	if conf.Node.WSPort, err = getInt(flagWsPort); err != nil {
		return nil, handleError(err, flagWsPort)
	}
	if conf.Node.WSPathPrefix, err = getString(flagWsPathPrefix); err != nil {
		return nil, handleError(err, flagWsPathPrefix)
	}
	if conf.Node.WSOrigins, err = getStringSlice(flagWsOrigins); err != nil {
		return nil, handleError(err, flagWsOrigins)
	}
	if conf.Node.WSModules, err = getStringSlice(flagWsModules); err != nil {
		return nil, handleError(err, flagWsModules)
	}
	if conf.Node.WSExposeAll, err = getBool(flagWsExposeAll); err != nil {
		return nil, handleError(err, flagWsExposeAll)
	}
	if conf.Node.GraphQLCors, err = getStringSlice(flagGraphqlCors); err != nil {
		return nil, handleError(err, flagGraphqlCors)
	}
	if conf.Node.GraphQLVirtualHosts, err = getStringSlice(flagGraphqlVirtualHosts); err != nil {
		return nil, handleError(err, flagGraphqlVirtualHosts)
	}
	if conf.Node.AllowUnprotectedTxs, err = getBool(flagAllowUnprotectedTxs); err != nil {
		return nil, handleError(err, flagAllowUnprotectedTxs)
	}
	if conf.Node.BatchRequestLimit, err = getInt(flagBatchRequestLimit); err != nil {
		return nil, handleError(err, flagBatchRequestLimit)
	}
	if conf.Node.BatchResponseMaxSize, err = getInt(flagBatchResponseMaxSize); err != nil {
		return nil, handleError(err, flagBatchResponseMaxSize)
	}
	if conf.Node.JWTSecret, err = getString(flagJwtSecret); err != nil {
		return nil, handleError(err, flagJwtSecret)
	}
	if conf.Node.DBEngine, err = getString(flagDBEngine); err != nil {
		return nil, handleError(err, flagDBEngine)
	}

	// Node.HTTPTimeouts settings
	if conf.Node.HTTPTimeouts.ReadTimeout, err = getTimeDuration(flagReadTimeout); err != nil {
		return nil, handleError(err, flagReadTimeout)
	}
	if conf.Node.HTTPTimeouts.ReadHeaderTimeout, err = getTimeDuration(flagReadHeaderTimeout); err != nil {
		return nil, handleError(err, flagReadHeaderTimeout)
	}
	if conf.Node.HTTPTimeouts.WriteTimeout, err = getTimeDuration(flagWriteTimeout); err != nil {
		return nil, handleError(err, flagWriteTimeout)
	}
	if conf.Node.HTTPTimeouts.IdleTimeout, err = getTimeDuration(flagIdleTimeout); err != nil {
		return nil, handleError(err, flagIdleTimeout)
	}

	return conf, nil
}
