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
	"github.com/ethereum/go-ethereum/node"
	"github.com/gin-gonic/gin"

	"pkg.berachain.dev/stargazer/eth/rpc/config"
)

type Service struct {
	backend StargazerBackend
	// `http` is the externally facing JSON-RPC Server.
	http *Server
	// `ws` is the externally facing JSON-RPC Server.
	ws *Server
	// `engine` is the gin engine responsible for handling the JSON-RPC requests.
	// engine *gin.Engine
	// `notify` is the channel that is used to notify the service has stopped.
	notify chan error
	// `config` is the configuration for the service.
	config config.Server
}

type Logger interface {
	gin.HandlerFunc
}

// `New` returns a new `Service` object.
func NewService(cfg config.Server, backend StargazerBackend) (*Service, error) {
	// Create the service object.
	s := &Service{
		http:    NewServer(),
		ws:      NewServer(),
		config:  cfg,
		notify:  make(chan error, 1),
		backend: backend,
	}

	return s, nil
}

func (s *Service) SetupAPIs() error {
	apis := GetAPIs(s.backend)
	for _, srv := range []*Server{s.http, s.ws} {
		if err := node.RegisterApis(apis, s.config.SetupAPIs, srv); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetHTTP() *Server {
	return s.http
}

func (s *Service) GetWS() *Server {
	return s.ws
}
