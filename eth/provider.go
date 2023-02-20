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

package eth

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/berachain/stargazer/eth/rpc"
	"github.com/berachain/stargazer/eth/rpc/config"

	"github.com/berachain/stargazer/eth/api"
	"github.com/berachain/stargazer/eth/core"
	ethlog "github.com/berachain/stargazer/eth/log"
)

type StargazerProvider struct {
	api.Chain
	rpcService *rpc.Service
}

// `NewStargazerProvider` creates a new `StargazerEVM` instance for use on an underlying blockchain.
func NewStargazerProvider(
	host core.StargazerHostChain,
	logHandler ethlog.Handler,
) *StargazerProvider {
	// When creating a Stargazer EVM, we allow the implementing chain
	// to specify their own log handler. If logHandler is nil then we
	// we use the default geth log handler.
	if logHandler != nil {
		// Root is a global in geth that is used by the evm to emit logs.
		ethlog.Root().SetHandler(ethlog.FuncHandler(logHandler))
	}

	return &StargazerProvider{
		Chain: core.NewChain(host),
	}
}

// `StartRPC` starts the RPC service for the Stargazer EVM.
func (p *StargazerProvider) StartRPC() error {
	var err error
	// We need to start the RPC service so that the Stargazer EVM can
	// make RPC calls to the underlying blockchain.
	// TODO: gate behind configuration
	rpcConfig := config.DefaultServer()
	p.rpcService, err = rpc.NewService(*rpcConfig, rpc.NewBackend(p.Chain, rpcConfig))
	if err != nil {
		return err
	}

	// TODO: handle graceful shutdown.
	go func() {
		p.rpcService.Start()

		// Waiting signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		// Wait for interrupt signal or an error to gracefully shutdown the server.
		select {
		case sig := <-interrupt:
			_ = sig // todo fix
		case err = <-p.rpcService.Notify():
			panic(err) // todo fix
		}

		// Ensure that if the switch statement outputs an error, we return it to the CLI.
		if err != nil {
			panic(err)
		}

		// Shutdown the server.
		if sErr := p.rpcService.Shutdown(); sErr != nil {
			// s.logger.Error(sErr.Error()/
			panic(err)
		}
	}()

	return nil
}
