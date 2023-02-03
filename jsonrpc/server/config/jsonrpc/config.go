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
	"github.com/berachain/stargazer/jsonrpc/server/config"
)

// `Config` defines configuration for the JSON-RPC server.Config struct.
type Config struct {
	// `API` defines a list of JSON-RPC namespaces to be enabled.
	API []string `mapstructure:"api"`

	// `Address` defines the HTTP server to listen on.
	Address string `mapstructure:"address"`

	// `WsAddress` defines the WebSocket server to listen on.
	WSAddress string `mapstructure:"ws-address"`

	// MetricsAddress defines the metrics server to listen on.
	MetricsAddress string `mapstructure:"metrics-address"`

	// `BaseRoute` defines the base path for the JSON-RPC server.
	BaseRoute string `mapstructure:"base-path"`
}

// `DefaultConfig` returns the default TLS configuration.
func DefaultConfig() *Config {
	return &Config{
		API:            config.DefaultAPINamespaces,
		Address:        config.DefaultJSONRPCAddress,
		WSAddress:      config.DefaultJSONRPCWSAddress,
		MetricsAddress: config.DefaultJSONRPCMetricsAddress,
		BaseRoute:      config.DefaultJSONRPCBaseRoute,
	}
}
