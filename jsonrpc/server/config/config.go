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

package config

import "time"

type (
	// `Server` defines the configuration for the JSON-RPC server.
	Server struct {
		// `API` defines a list of JSON-RPC namespaces to be enabled.
		EnableAPIs []string `mapstructure:"api"`

		// `Address` defines the HTTP server to listen on.
		Address string `mapstructure:"address"`

		// `WsAddress` defines the WebSocket server to listen on.
		WSAddress string `mapstructure:"ws-address"`

		// `MetricsAddress` defines the metrics server to listen on.
		MetricsAddress string `mapstructure:"metrics-address"`

		// `HTTPReadHeaderTimeout` is the read timeout of http json-rpc server.
		HTTPReadHeaderTimeout time.Duration `mapstructure:"http-read-header-timeout"`

		// `HTTPReadTimeout` is the read timeout of http json-rpc server.
		HTTPReadTimeout time.Duration `mapstructure:"http-read-timeout"`

		// `HTTPWriteTimeout` is the write timeout of http json-rpc server.
		HTTPWriteTimeout time.Duration `mapstructure:"http-write-timeout"`

		// HTTPIdleTimeout is the idle timeout of http json-rpc server.
		HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`

		// `HTTPBaseRoute` defines the base path for the JSON-RPC server.
		BaseRoute string `mapstructure:"base-path"`

		// `TLSConfig` defines the TLS configuration for the JSON-RPC server.
		TLSConfig *TLSConfig `mapstructure:"tls-config"`
	}

	// `TLSConfig` defines a certificate and matching private key for the server.
	TLSConfig struct {
		// `CertPath` the file path for the certificate .pem file
		CertPath string `mapstructure:"cert-path"`

		// KeyPath the file path for the key .pem file
		KeyPath string `toml:"key-path"`
	}
)
