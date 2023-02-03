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

package api

import (
	"github.com/berachain/stargazer/jsonrpc/api/eth"
	"github.com/berachain/stargazer/jsonrpc/api/node"
	"github.com/berachain/stargazer/jsonrpc/cosmos"
	"go.uber.org/zap"
)

// `Service` is an interface that all API services must implement.
type Service interface {
	Namespace() string
}

// `Build` returns a new API service based on the given namespace.
func Build(
	namespace string,
	client *cosmos.Client,
	logger *zap.Logger,
) Service {
	switch namespace {
	case "node":
		return node.NewAPI(logger)
	case "eth":
		return eth.NewAPI(client, logger)
	default:
		return nil
	}
}
