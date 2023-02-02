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
	"github.com/berachain/stargazer/jsonrpc/api/node"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	rpcserver       *ethrpc.Server
	server          *http.Server
	logger          libtypes.Logger[zap.Field]
	notify          chan error
	shutdownTimeout time.Duration
	config          Config
	clientCtx       client.Context
}

func New(config Config, logger libtypes.Logger[zap.Field], clientCtx client.Context) *Service {
	s := &Service{
		rpcserver: ethrpc.NewServer(),
		config:    config,
		logger:    logger,
		clientCtx: clientCtx,
		notify:    make(chan error, 1),
	}

	s.server = &http.Server{
		Addr:    config.rpc.Address,
		Handler: s.rpcserver,
		// TODO is this correct?
		ReadHeaderTimeout: time.Second, // s.config.rpc.HTTPTimeout,
		// TLSConfig:         s.config.tls.TLSConfig(),
		WriteTimeout: time.Second, // s.config.rpc.HTTPTimeout,
	}

	// TODO: move these to a proper spot
	if err := s.RegisterAPI(node.NewAPI(logger)); err != nil {
		panic(err)
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

// Notify -.
func (s *Service) Notify() <-chan error {
	return s.notify
}

func (s *Service) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.shutdownTimeout,
	)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// `RegisterService` registers a service with the server.
func (s *Service) RegisterAPI(service api.Service) error {
	return s.rpcserver.RegisterName(service.Namespace(), service)
}
