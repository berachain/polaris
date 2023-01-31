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
	"fmt"
	"net/http"
	"time"

	"github.com/berachain/stargazer/jsonrpc/api"
	"github.com/cosmos/cosmos-sdk/client"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type Service struct {
	*ethrpc.Server
	config    Config
	clientCtx client.Context
}

func New(config Config, clientCtx client.Context) *Service {
	s := ethrpc.NewServer()
	return &Service{
		Server:    s,
		config:    config,
		clientCtx: clientCtx,
	}
}

func (s *Service) Start(errCh chan error) {
	// TODO: move into `./jsonrpc` and add configuration file.
	httpSrv := &http.Server{
		Addr:    s.config.rpc.Address,
		Handler: s,
		// TODO is this correct?
		ReadHeaderTimeout: time.Second, // s.config.rpc.HTTPTimeout,
		// TLSConfig:         s.config.tls.TLSConfig(),
		WriteTimeout: time.Second, // s.config.rpc.HTTPTimeout,
	}

	// // TODO: move these to a proper spot
	// if err := s.RegisterService(node.NewAPI()); err != nil {
	// 	errCh <- err
	// 	return
	// }

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
		errCh <- err
	}
}

// `ClientContext` returns the client context.
func (s *Service) ClientContext() client.Context {
	return s.clientCtx
}

// `RegisterService` registers a service with the server.
func (s *Service) RegisterService(service api.Service) error {
	return s.Server.RegisterName(service.Namespace(), service)
}
