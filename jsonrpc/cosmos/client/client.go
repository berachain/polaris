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

	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"
)

// `CosmosClient` is a wrapper around the Cosmos SDK `client.Context` that implements querying and
// transaction capabilities for the Cosmos SDK.
type CosmosClient struct {
	// `ctx` is the context instance.
	ctx context.Context
	// `clientCtx` is the Cosmos SDK `client.Context` instance.
	clientCtx client.Context

	// `cbCtx` is the `CometBlockClient` context.
	cbc CometBlockClient

	// `logger` is the logger instance.
	logger libtypes.Logger[zap.Field]
}

// `New` creates a new `CosmosClient`.
func New(
	ctx context.Context,
	clientCtx client.Context,
	logger libtypes.Logger[zap.Field],
) *CosmosClient {
	return &CosmosClient{
		ctx:       ctx,
		clientCtx: clientCtx,
		logger:    logger,
		cbc:       utils.MustGetAs[CometBlockClient](clientCtx.Client),
	}
}
