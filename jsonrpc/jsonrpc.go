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

package jsonrpc

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"

	"github.com/berachain/stargazer/jsonrpc/server"
)

// `Service` is a JSON-RPC endpoint service.
type Service struct {
	logger *zap.Logger
	server *server.Service
}

// `New` is a constructor for `Service`.
func New(config server.Config, clientCtx client.Context) *Service {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint: errcheck // ignore error
	return &Service{
		logger: logger,
		server: server.New(config, logger, clientCtx),
	}
}

// `Start` starts the service.
func (s *Service) Start() error {
	// errCh := make(chan error)
	// 1. Build CosmosClient to connect to node
	// TODO: implement

	// 2. Setup JSONRPC Server to provide endpoint
	s.server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal or an error to gracefully shutdown the server.
	var err error
	select {
	case sig := <-interrupt:
		s.logger.Info("app - Run - signal: " + sig.String())
	case err = <-s.server.Notify():
		s.logger.Error(err.Error())
	}

	// Shutdown
	if sErr := s.server.Shutdown(); sErr != nil {
		s.logger.Error(sErr.Error())
		return sErr
	}

	// Ensure that if the switch statement outputs an error, we return it to the CLI.
	return err
}
