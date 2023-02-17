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

package server

import (
	"context"
	"time"

	ethlog "github.com/berachain/stargazer/eth/log"
	"github.com/berachain/stargazer/jsonrpc/server/config"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	// `rpcserver` is the externally facing JSON-RPC Server.
	rpcserver *ethrpc.Server
	// `engine` is the gin engine responsible for handling the JSON-RPC requests.
	engine *gin.Engine
	// `logger` is the logger for the service.
	logger *zap.Logger
	// `notify` is the channel that is used to notify the service has stopped.
	notify chan error
	// `shutdownTimeout` is the delay between the service being stopped and the HTTP server being shutdown.
	shutdownTimeout time.Duration
	// `config` is the configuration for the service.
	config config.Server
}

// `New` returns a new `Service` object.
func New(ctx context.Context, logger *zap.Logger, cfg config.Server) *Service {
	// Create the service object.
	s := &Service{
		rpcserver: ethrpc.NewServer(),
		config:    cfg,
		logger:    logger,
		notify:    make(chan error, 1),
		engine:    gin.New(),
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	s.engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Like with gin, we must use middleware to have ethlog output to our logger.
	ethlog.Root().SetHandler(ethlog.FuncHandler(func(r *ethlog.Record) error {
		sugared := logger.Sugar()
		switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
		case ethlog.LvlTrace, ethlog.LvlDebug:
			sugared.Debug(r.Msg, r.Ctx)
		case ethlog.LvlInfo, ethlog.LvlWarn:
			sugared.Info(r.Msg, r.Ctx)
		case ethlog.LvlError, ethlog.LvlCrit:
			sugared.Error(r.Msg, r.Ctx)
		}
		return nil
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	s.engine.Use(gin.Recovery())

	// Set the JSON-RPC server to use the BaseRoute.
	s.engine.Any(s.config.BaseRoute, gin.WrapH(s.rpcserver))

	return s
}

// `Start` stops the service.
func (s *Service) Start() {
	go func() {
		s.logger.Info("Starting JSON-RPC server at:", zap.String("address", s.config.Address))
		s.notify <- s.engine.Run(s.config.Address)
		close(s.notify)
	}()
}

// `Notify` returns a channel that is used to notify the service has stopped.
func (s *Service) Notify() <-chan error {
	return s.notify
}

// `Shutdown` stops the service.
func (s *Service) Shutdown() error {
	_, cancel := context.WithTimeout(
		context.Background(),
		s.shutdownTimeout,
	)
	defer cancel()
	// Stop the RPC Server
	s.rpcserver.Stop()
	// TODO: stop the gin server
	return nil
}
