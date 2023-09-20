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

	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/node"

	"pkg.berachain.dev/polaris/eth/polar"
)

// DefaultConfig returns the default configuration for a polaris chain.
func DefaultConfig() *Config {
	return &Config{
		Polar: *polar.DefaultConfig(),
		Node:  node.DefaultConfig,
	}
}

type Config struct {
	Polar polar.Config
	Node  node.Config
}

//nolint:funlen,gocognit,gocyclo,cyclop // TODO break up later.
func ReadConfigFromAppOpts(opts servertypes.AppOptions) (*Config, error) {
	var err error
	var val int64
	conf := &Config{}

	// Define little error handler.
	var handleError = func(err error) error {
		if err != nil {
			return fmt.Errorf("error while reading configuration: %w", err)
		}
		return nil
	}

	// Wrapping casting functions to return both value and error
	getString := func(key string) (string, error) { return cast.ToStringE(opts.Get(key)) }
	getInt := func(key string) (int, error) { return cast.ToIntE(opts.Get(key)) }
	getInt64 := func(key string) (int64, error) { return cast.ToInt64E(opts.Get(key)) }
	getUint64 := func(key string) (uint64, error) { return cast.ToUint64E(opts.Get(key)) }
	getFloat64 := func(key string) (float64, error) { return cast.ToFloat64E(opts.Get(key)) }
	getBool := func(key string) (bool, error) { return cast.ToBoolE(opts.Get(key)) }
	getStringSlice := func(key string) ([]string, error) { return cast.ToStringSliceE(opts.Get(key)) }
	getTimeDuration := func(key string) (time.Duration, error) { return cast.ToDurationE(opts.Get(key)) }

	// Polar settings
	if conf.Polar.RPCGasCap, err = getUint64(flagRPCGasCap); err != nil {
		return nil, handleError(err)
	}
	if conf.Polar.RPCEVMTimeout, err = getTimeDuration(flagRPCEvmTimeout); err != nil {
		return nil, handleError(err)
	}
	if conf.Polar.RPCTxFeeCap, err = getFloat64(flagRPCTxFeeCap); err != nil {
		return nil, handleError(err)
	}

	// Polar.GPO settings
	if conf.Polar.GPO.Blocks, err = getInt(flagBlocks); err != nil {
		return nil, handleError(err)
	}
	if conf.Polar.GPO.Percentile, err = getInt(flagPercentile); err != nil {
		return nil, handleError(err)
	}
	if conf.Polar.GPO.MaxHeaderHistory, err = getUint64(flagMaxHeaderHistory); err != nil {
		return nil, handleError(err)
	}
	if conf.Polar.GPO.MaxBlockHistory, err = getUint64(flagMaxBlockHistory); err != nil {
		return nil, handleError(err)
	}
	if val, err = getInt64(flagDefault); err != nil {
		return nil, handleError(err)
	}
	conf.Polar.GPO.Default = big.NewInt(val)

	if val, err = getInt64(flagDefault); err != nil {
		return nil, handleError(err)
	}
	conf.Polar.GPO.MaxPrice = big.NewInt(val)

	if val, err = getInt64(flagDefault); err != nil {
		return nil, handleError(err)
	}
	conf.Polar.GPO.IgnorePrice = big.NewInt(val)

	// Node settings
	if conf.Node.Name, err = getString(flagName); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.UserIdent, err = getString(flagUserIdent); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.Version, err = getString(flagVersion); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.DataDir, err = getString(flagDataDir); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.KeyStoreDir, err = getString(flagKeyStoreDir); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.ExternalSigner, err = getString(flagExternalSigner); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.UseLightweightKDF, err = getBool(flagUseLightweightKdf); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.InsecureUnlockAllowed, err = getBool(flagInsecureUnlockAllowed); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.USB, err = getBool(flagUsb); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.SmartCardDaemonPath, err = getString(flagSmartCardDaemonPath); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.IPCPath, err = getString(flagIpcPath); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPHost, err = getString(flagHTTPHost); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPPort, err = getInt(flagHTTPPort); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPCors, err = getStringSlice(flagHTTPCors); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPVirtualHosts, err = getStringSlice(flagHTTPVirtualHosts); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPModules, err = getStringSlice(flagHTTPModules); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPPathPrefix, err = getString(flagHTTPPathPrefix); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.AuthAddr, err = getString(flagAuthAddr); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.AuthPort, err = getInt(flagAuthPort); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.AuthVirtualHosts, err = getStringSlice(flagAuthVirtualHosts); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSHost, err = getString(flagWsHost); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSPort, err = getInt(flagWsPort); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSPathPrefix, err = getString(flagWsPathPrefix); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSOrigins, err = getStringSlice(flagWsOrigins); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSModules, err = getStringSlice(flagWsModules); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.WSExposeAll, err = getBool(flagWsExposeAll); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.GraphQLCors, err = getStringSlice(flagGraphqlCors); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.GraphQLVirtualHosts, err = getStringSlice(flagGraphqlVirtualHosts); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.AllowUnprotectedTxs, err = getBool(flagAllowUnprotectedTxs); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.BatchRequestLimit, err = getInt(flagBatchRequestLimit); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.BatchResponseMaxSize, err = getInt(flagBatchResponseMaxSize); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.JWTSecret, err = getString(flagJwtSecret); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.DBEngine, err = getString(flagDBEngine); err != nil {
		return nil, handleError(err)
	}

	// Node.HTTPTimeouts settings
	if conf.Node.HTTPTimeouts.ReadTimeout, err = getTimeDuration(flagReadTimeout); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPTimeouts.ReadHeaderTimeout, err = getTimeDuration(flagReadHeaderTimeout); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPTimeouts.WriteTimeout, err = getTimeDuration(flagWriteTimeout); err != nil {
		return nil, handleError(err)
	}
	if conf.Node.HTTPTimeouts.IdleTimeout, err = getTimeDuration(flagIdleTimeout); err != nil {
		return nil, handleError(err)
	}

	return conf, nil
}
