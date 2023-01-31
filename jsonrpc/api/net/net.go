// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package net

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	tmclient "github.com/tendermint/tendermint/rpc/client"

	libtypes "github.com/berachain/stargazer/lib/types"
	"go.uber.org/zap/zapcore"
)

var (
	tempChainID uint64 = 420 // TODO UNTEMP
)

// `api`.
type api struct {
	logger        libtypes.Logger[zapcore.Field]
	networkClient tmclient.NetworkClient
	chainID       uint64
}

// `New` creates an instance of the public Net Web3 API.
//
//nolint:revive // by design.
func New(clientCtx client.Context, logger libtypes.Logger[zapcore.Field]) *api {
	// TODO: get chainid from the context
	// // parse the chainID from a integer string
	// chainID, err := ParseChainID(clientCtx.ChainID)
	// if err != nil {
	// 	panic(err)
	// }

	// Ensure that the client from the context is a NetworkClient.
	networkClient, ok := clientCtx.Client.(tmclient.NetworkClient)
	if !ok {
		panic("invalid client")
	}

	return &api{
		logger:        logger,
		networkClient: networkClient,
		chainID:       tempChainID, // TODO UNTEMP THIS
	}
}

// `Version` returns the current ethereum protocol version.
func (api *api) Version() string {
	// `net_version` is a fancy way of saying "what chain_id are you connected to?"
	api.logger.Debug("net_version")
	return fmt.Sprintf("%d", api.chainID)
}

// `Listening` returns whether or not a client is actively listening for network connections.
func (api *api) Listening() bool {
	api.logger.Debug("net_listening")
	netInfo, err := api.networkClient.NetInfo(context.Background())
	if err != nil {
		return false
	}
	return netInfo.Listening
}

// `PeerCount` returns the total number of peers currently connected to the client.
func (api *api) PeerCount() int {
	api.logger.Debug("net_peerCount")
	netInfo, err := api.networkClient.NetInfo(context.Background())
	if err != nil {
		return 0
	}
	return len(netInfo.Peers)
}
