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

package jsonrpc

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/berachain/stargazer/jsonrpc/server"
)

// `Service` is a JSON-RPC endpoint service.
type Service struct {
	logger *zap.Logger
	server server.Service
}

// `New` is a constructor for `Service`.
func New(config Config) *Service {
	ctx := context.Background()

	// Create a new logger instance.
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint: errcheck // ignore error

	return &Service{
		logger: logger,
		server: *server.New(ctx, logger, config.Server),
	}
}

// `Start` starts the service.
func (s *Service) Start() error {
	// 2. Setup JSONRPC Server to provide endpoint
	s.server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal or an error to gracefully shutdown the server.
	var err error
	select {
	case sig := <-interrupt:
		s.logger.Info(sig.String())
	case err = <-s.server.Notify():
		s.logger.Error(err.Error())
	}

	// Ensure that if the switch statement outputs an error, we return it to the CLI.
	if err != nil {
		return err
	}

	// Shutdown the server.
	if sErr := s.server.Shutdown(); sErr != nil {
		s.logger.Error(sErr.Error())
		return sErr
	}

	return nil
}
