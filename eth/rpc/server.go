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

	"pkg.berachain.dev/polaris/eth/rpc/config"
)

// Service is the interface for the JSON-RPC service.
type Service interface {
	SetBackend(PolarisBackend)
	RegisterAPIs(extraAPIs func(PolarisBackend) []API) error
	GetHTTP() *Server
	GetWS() *Server
	GetConfig() *config.Server
}

// Service is a wrapper around go-ethereum JSON-RPC server(s). That also
// supplies a backend to handle the requests.
type service struct {
	// backend is the backend for the service.
	backend PolarisBackend
	// config is the configuration for the service.
	config *config.Server
	// http is the externally facing JSON-RPC Server.
	http *Server
	// ws is the externally facing JSON-RPC Server.
	ws *Server
}

// New returns a new `Service` object.
func NewService(cfg *config.Server) Service {
	return &service{
		backend: nil,
		config:  cfg,
		http:    NewServer(),
		ws:      NewServer(),
	}
}

// RegisterAPIs registers the JSON-RPC APIs with the API service.
func (s *service) RegisterAPIs(extraAPIs func(PolarisBackend) []API) error {
	apis := append(GetAPIs(s.backend), extraAPIs(s.backend)...)
	for _, srv := range []*Server{s.http, s.ws} {
		if err := node.RegisterApis(apis, s.config.EnabledAPIs, srv); err != nil {
			return err
		}
	}
	return nil
}

// SetBackend sets the backend for the service.
func (s *service) SetBackend(backend PolarisBackend) {
	s.backend = backend
}

// GetConfig returns the configuration for the service.
func (s *service) GetConfig() *config.Server {
	return s.config
}

// GetHTTP returns the HTTP server.
func (s *service) GetHTTP() *Server {
	return s.http
}

// GetWS returns the WS server.
func (s *service) GetWS() *Server {
	return s.ws
}
