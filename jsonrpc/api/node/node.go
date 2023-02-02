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

package node

import (
	libtypes "github.com/berachain/stargazer/lib/types"
	"go.uber.org/zap/zapcore"
)

// `API` is the node API.
type api struct {
	logger libtypes.Logger[zapcore.Field]
}

func NewAPI(logger libtypes.Logger[zapcore.Field]) *api { //nolint: revive // by design.
	return &api{logger: logger}
}

// `Namespace` impements the api.Service interface.
func (api) Namespace() string {
	return "node"
}

// `Health` returns if the stargazer node is healthy.
func (api *api) Health() string {
	api.logger.Info("node_health")
	return "ok" // todo query the node status
}

// `RpcHealth` returns if the rpc server is healthy.
func (api api) RpcHealth() string { //nolint: revive // by design.
	api.logger.Info("node_rpcHealth")
	return "ok"
}
