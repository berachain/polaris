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

var (
	// `DefaultAPINamespaces` is the default namespaces the JSON-RPC server exposes.
	DefaultAPINamespaces = []string{"eth", "net", "web3"}
)

const (
	// `DefaultGRPCAddress` is the default address the gRPC server binds to.
	DefaultGRPCAddress = "0.0.0.0:9900"

	// `DefaultJSONRPCAddress` is the default address the JSON-RPC server binds to.
	DefaultJSONRPCAddress = "127.0.0.1:8545"

	// `DefaultJSONRPCWSAddress` is the default address the JSON-RPC WebSocket server binds to.
	DefaultJSONRPCWSAddress = "127.0.0.1:8546"

	// `DefaultJSOPNRPCMetricsAddress` is the default address the JSON-RPC Metrics server binds to.
	DefaultJSONRPCMetricsAddress = "127.0.0.1:6065"

	// `DefaultBaseRoute` is the default base path for the JSON-RPC server.
	DefaultJSONRPCBaseRoute = "/"
)

type Config struct {
	rpc *JSONRPCConfig
	tls *TLSConfig
}

func DefaultConfig() *Config {
	return &Config{
		rpc: DefaultJSONRPCConfig(),
		tls: DefaultTLSConfig(),
	}
}

// `JSONRPCConfig` defines configuration for the JSON-RPC server.Config struct.
type JSONRPCConfig struct {
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
func DefaultJSONRPCConfig() *JSONRPCConfig {
	return &JSONRPCConfig{
		API:            DefaultAPINamespaces,
		Address:        DefaultJSONRPCAddress,
		WSAddress:      DefaultJSONRPCWSAddress,
		MetricsAddress: DefaultJSONRPCMetricsAddress,
		BaseRoute:      DefaultJSONRPCBaseRoute,
	}
}

// `TLSConfig` defines a certificate and matching private key for the server.
type TLSConfig struct {
	// `CertPath` the file path for the certificate .pem file
	CertPath string `mapstructure:"cert-path"`

	// KeyPath the file path for the key .pem file
	KeyPath string `toml:"key-path"`
}

// DefaultConfig returns the default TLS configuration.
func DefaultTLSConfig() *TLSConfig {
	return &TLSConfig{
		CertPath: "",
		KeyPath:  "",
	}
}
