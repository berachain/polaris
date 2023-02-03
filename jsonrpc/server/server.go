// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package server

import (
	"context"
	"time"

	"github.com/berachain/stargazer/jsonrpc/api"
	"github.com/berachain/stargazer/jsonrpc/cosmos"
	"github.com/berachain/stargazer/jsonrpc/logger"
	"github.com/berachain/stargazer/jsonrpc/server/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	// `cosmosClient` provides the gRPC connection to the Cosmos node.
	cosmosClient *cosmos.Client
	// `rpcserver` is the externally facing JSON-RPC Server.
	rpcserver *ethrpc.Server
	// `engine` is the gin engine responsible for handling the JSON-RPC requests.
	engine *gin.Engine
	// `logger` is the logger for the service.
	logger logger.Zap
	// `notify` is the channel that is used to notify the service has stopped.
	notify chan error
	// `shutdownTimeout` is the delay between the service being stopped and the HTTP server being shutdown.
	shutdownTimeout time.Duration
	// `config` is the configuration for the service.
	config config.Server
}

// `New` returns a new `Service` object.
func New(ctx context.Context, logger logger.Zap, client *cosmos.Client, cfg config.Server) *Service {
	// Configure the JSON-RPC API.
	s := &Service{
		cosmosClient: client,
		rpcserver:    ethrpc.NewServer(),
		config:       cfg,
		logger:       logger,
		notify:       make(chan error, 1),
		engine:       gin.Default(),
	}

	// Set the JSON-RPC server to use the BaseRoute.
	s.engine.Any(s.config.BaseRoute, gin.WrapH(s.rpcserver))

	// Register the JSON-RPC API namespaces.
	for _, namespace := range cfg.EnableAPIs {
		if err := s.RegisterAPI(api.Build(namespace, s.cosmosClient, logger)); err != nil {
			panic(err)
		}
	}

	return s
}

// `Start` stops the service.
func (s *Service) Start() {
	go func() {
		s.logger.Info("Starting JSON-RPC server at", zap.String("address", s.config.Address))
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

// `RegisterAPI` registers a service with the JSON-RPC server.
func (s *Service) RegisterAPI(service api.Service) error {
	return s.rpcserver.RegisterName(service.Namespace(), service)
}
