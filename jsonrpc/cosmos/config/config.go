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

const (
	// `DefaultCMRPCEndpoint` is the default address of the Comet RPC server.
	DefaultCMRPCEndpoint = "http://0.0.0.0:26657"
	// `DefaultGRPCAddress` is the default address the gRPC server binds to.
	DefaultGRPCAddress = "http://0.0.0.0:9900"
	// `DefaultRPCTimeout` is the default timeout for the RPC server.
	DefaultRPCTimeout = "10s"
	// `DefaultChainID` is the default chain ID.
	DefaultChainID = "berachain_420-1"
)

type (
	// RPC defines RPC configuration of both the gRPC and CometBFT nodes.
	RPC struct {
		CMRPCEndpoint string `mapstructure:"cmrpc-endpoint" validate:"required"`
		GRPCEndpoint  string `mapstructure:"grpc-endpoint" validate:"required"`
		RPCTimeout    string `mapstructure:"rpc-timeout" validate:"required"`
		ChainID       string `mapstructure:"chain-id" validate:"required"`
	}
)

// DefaultRPC returns the default RPC configuration.
func DefaultRPC() *RPC {
	return &RPC{
		CMRPCEndpoint: DefaultCMRPCEndpoint,
		GRPCEndpoint:  DefaultGRPCAddress,
		RPCTimeout:    DefaultRPCTimeout,
		ChainID:       DefaultChainID,
	}
}
