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
	"net/http"
	"time"

	"github.com/berachain/stargazer/jsonrpc/api"
	"github.com/berachain/stargazer/jsonrpc/cosmos"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	// `cosmosClient` provides the gRPC connection to the Cosmos node.
	cosmosClient *cosmos.Client
	// `rpcserver` is the externally facing JSON-RPC Server.
	rpcserver *ethrpc.Server
	// `server` is the HTTP server that serves the JSON-RPC server.
	server *http.Server
	// `logger` is the logger for the service.
	logger libtypes.Logger[zap.Field]
	// `notify` is the channel that is used to notify the service has stopped.
	notify chan error
	// `shutdownTimeout` is the delay between the service being stopped and the HTTP server being shutdown.
	shutdownTimeout time.Duration
	// `config` is the configuration for the service.
	config Config
}

// `New` returns a new `Service` object.
func New(ctx context.Context, logger libtypes.Logger[zap.Field], config Config, clientCtx client.Context) *Service {

	r := gin.Default()

	// Configure the JSON-RPC API.
	s := &Service{
		cosmosClient: cosmos.New(ctx, clientCtx, logger),
		rpcserver:    ethrpc.NewServer(),
		config:       config,
		logger:       logger,
		notify:       make(chan error, 1),
	}

	// Set the JSON-RPC server to use thea base route.
	r.Any(s.config.rpc.BaseRoute, gin.WrapH(ethrpc.NewServer()))
	r.ServeHTTP()
	// Configure the HTTP server.
	s.server = &http.Server{
		Addr:              config.rpc.Address,
		ReadHeaderTimeout: 5 * time.Second,  //nolint:gomnd // TODO: make this configurable
		ReadTimeout:       10 * time.Second, //nolint:gomnd // TODO: make this configurable
		WriteTimeout:      10 * time.Second, //nolint:gomnd // TODO: make this configurable
		Handler:           r,
	}

	// Register the JSON-RPC API namespaces.
	for _, namespace := range config.rpc.API {
		if err := s.RegisterAPI(api.Build(namespace, s.cosmosClient, logger)); err != nil {
			panic(err)
		}
	}

	return s
}

// `Start` stops the service.
func (s *Service) Start() {
	go func() {
		s.logger.Info("Starting JSON-RPC server at", zap.String("address", s.config.rpc.Address))
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// `Notify` returns a channel that is used to notify the service has stopped.
func (s *Service) Notify() <-chan error {
	return s.notify
}

// `Shutdown` stops the service.
func (s *Service) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.shutdownTimeout,
	)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// `RegisterAPI` registers a service with the JSON-RPC server.
func (s *Service) RegisterAPI(service api.Service) error {
	return s.rpcserver.RegisterName(service.Namespace(), service)
}
