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
	"context"

	"github.com/berachain/stargazer/eth/core/types"
)

// `Chain` defines the methods that the Stargazer Ethereum API exposes. This is the only interface
// that an implementing chain should use.
// TODO: rename.
type Chain interface {
	// `Prepare` prepares the chain for a new block. This method is called before the first tx in
	// the block.
	Prepare(ctx context.Context, header *types.StargazerHeader)
	// `ProcessTransaction` processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
	// `Finalize` finalizes the block and returns the block. This method is called after the last
	// tx in the block.
	Finalize(ctx context.Context) (*types.StargazerBlock, error)
}
