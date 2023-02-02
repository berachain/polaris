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
	"fmt"
	"net/http"
	"time"

	"github.com/berachain/stargazer/jsonrpc/api"
	"github.com/berachain/stargazer/jsonrpc/api/node"
	"github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	rpcserver       *ethrpc.Server
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	config          Config
	clientCtx       client.Context
}

func New(config Config, clientCtx client.Context) *Service {
	ethrpc := ethrpc.NewServer()
	api := node.NewAPI(zap.NewNop())

	err := ethrpc.RegisterName(api.Namespace(), api)
	if err != nil {
		panic(err)
	}

	fmt.Println("jsonrpc server listening on", config.rpc.Address)

	httpSrv := &http.Server{
		Addr:    config.rpc.Address,
		Handler: ethrpc,
		// TODO is this correct?
		ReadHeaderTimeout: time.Second, // s.config.rpc.HTTPTimeout,
		// TLSConfig:         s.config.tls.TLSConfig(),
		WriteTimeout: time.Second, // s.config.rpc.HTTPTimeout,
	}

	return &Service{
		rpcserver:       ethrpc,
		server:          httpSrv,
		notify:          make(chan error, 1),
		config:          config,
		shutdownTimeout: time.Second,
		clientCtx:       clientCtx,
	}
}

// `Start` stops the service.
func (s *Service) Start() {
	// TODO: move into `./jsonrpc` and add configuration file.
	httpSrv := &http.Server{
		Addr:    s.config.rpc.Address,
		Handler: s.rpcserver,
		// TODO is this correct?
		ReadHeaderTimeout: time.Second, // s.config.rpc.HTTPTimeout,
		// TLSConfig:         s.config.tls.TLSConfig(),
		WriteTimeout: time.Second, // s.config.rpc.HTTPTimeout,
	}

	// TODO: move these to a proper spot
	if err := s.RegisterService(node.NewAPI(zap.NewNop())); err != nil {
		s.notify <- err
		return
	}

	httpSrvDone := make(chan struct{}, 1)
	if err := httpSrv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			close(httpSrvDone)
			return
		}
		// TODO: proper logger
		//nolint: forbidigo // temp.
		fmt.Println("failed to start JSON-RPC server", "error", err.Error())
		// 	fmt.Println("failed to start JSON-RPC server", "error", err.Error())
		s.notify <- err
	}
}

// Notify -.
func (s *Service) Notify() <-chan error {
	return s.notify
}

func (s *Service) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

// `ClientContext` returns the client context.
func (s *Service) ClientContext() client.Context {
	return s.clientCtx
}

// `RegisterService` registers a service with the server.
func (s *Service) RegisterService(service api.Service) error {
	return s.rpcserver.RegisterName(service.Namespace(), service)
}
