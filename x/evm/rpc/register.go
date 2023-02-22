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

package rpc

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	ethrpc "pkg.berachain.dev/stargazer/eth/rpc"
	ethrpcconfig "pkg.berachain.dev/stargazer/eth/rpc/config"
)

const (
	httpPath = "/eth/rpc"
	wsPath   = "/eth/rpc"
)

type Provider interface {
	GetHTTP() *ethrpc.Server
	GetWS() *ethrpc.Server
	Ready() bool
	SetupAPIs() error
}

type provider struct {
	ethrpc.Service
	backend ethrpc.Backend
}

// `NewProvider` returns a new `Provider` object. The provider object is used to
// register the JSON-RPC servers with the API server.
func NewProvider(cfg ethrpcconfig.Server, backend ethrpc.Backend) Provider {
	service, err := ethrpc.NewService(cfg, backend)
	if err != nil {
		panic(err)
	}
	return &provider{
		Service: *service,
		backend: backend,
	}
}

func (p *provider) Ready() bool {
	// TODO: there is might a race condition? maybe not? we can maybe just remove this.
	return true
}

// `RegisterJSONRPCServer` provides a common function which registers the ethereum rpc servers
// with routes on the native Cosmos API Server.
func RegisterJSONRPCServer(ctx client.Context, rtr *mux.Router, provider Provider) error {
	if !provider.Ready() {
		return errors.New("JSONRPC provider not ready")
	}
	httpSrv := provider.GetHTTP()
	wsSrv := provider.GetWS()
	rtr.PathPrefix(httpPath).Handler(httpSrv)
	rtr.PathPrefix(httpPath + "/").Handler(httpSrv)
	rtr.PathPrefix(wsPath).Handler(wsSrv)
	rtr.PathPrefix(wsPath + "/").Handler(wsSrv)
	return nil
}
