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

package block

import (
	"context"

	"go.uber.org/zap"

	libtypes "github.com/berachain/stargazer/lib/types"

	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// `Client` is a query client that can be used to gather block information from
// a Cosmos SDK based blockchain.
type Client struct {
	// `ctx` is the context instance.
	ctx context.Context
	// `logger` is the logger instance.
	logger libtypes.Logger[zap.Field]
	// `cbCtx` is the `CometBlockClient` context.
	cbc CometBlockClient
}

// `NewClient` creates a new `Client` instance.
func NewClient(cbc CometBlockClient) Client {
	return Client{cbc: cbc}
}

// `LatestBlockNumber` returns the the latest block number as reported at the application layer.
func (c *Client) LatestBlockNumber() (uint64, error) {
	res, err := c.cbc.ABCIInfo(c.ctx)
	if err != nil {
		return 0, err
	}
	return uint64(res.Response.LastBlockHeight), nil
}

// CometBlockByNumber returns a CometBFT-formatted block at a given chain height.
func (c *Client) CometBlockByNumber(height int64) (*tmrpctypes.ResultBlock, error) {
	if height <= 0 {
		// fetch the latest block number from the app state, more accurate than the tendermint block store state.
		n, err := c.LatestBlockNumber()
		if err != nil {
			return nil, err
		}
		height = int64(n)
	}

	resBlock, err := c.cbc.Block(c.ctx, &height)
	if err != nil {
		c.logger.Debug("CometBlockClient client failed to get block",
			zap.Int64("height", height), zap.String("error", err.Error()))
		return nil, err
	} else if resBlock.Block == nil {
		c.logger.Debug("CometBlockByNumber: block not found", zap.Int64("height", height))
		return nil, nil //nolint:nilnil // not finding the block isn't nessarily an error.
	}

	return resBlock, nil
}
