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

package client

import (
	"context"

	"github.com/berachain/stargazer/jsonrpc/cosmos/client/block"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"
)

// `NodeClient` is a wrapper around the Cosmos SDK `client.Context` that implements querying and
// transaction capabilities for the Cosmos SDK.
type NodeClient struct {
	ctx       context.Context
	clientCtx client.Context
	logger    libtypes.Logger[zap.Field]

	blockClient block.Client
}

// `New` creates a new `NodeClient`.
func New(
	ctx context.Context,
	clientCtx client.Context,
	logger libtypes.Logger[zap.Field],
) *NodeClient {
	return &NodeClient{
		ctx:         ctx,
		clientCtx:   clientCtx,
		logger:      logger,
		blockClient: utils.MustGetAs[block.Client](clientCtx.Client),
	}
}
