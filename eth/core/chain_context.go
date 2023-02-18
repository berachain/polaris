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

package core

import (
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/ethereum/go-ethereum/consensus"
)

// Compile-time interface assertion.
var _ ChainContext = (*chainContext)(nil)

// `chainContext` is a wrapper around `StateProcessor` that implements the `ChainContext` interface.
type chainContext struct {
	*StateProcessor
}

// `GetHeader` returns the header for the given hash and height. This is used by the `GetHashFn`.
func (cc *chainContext) GetHeader(_ common.Hash, height uint64) *types.Header {
	if header := cc.StateProcessor.bp.GetStargazerHeaderAtHeight(int64(height)); header != nil {
		return header.Header
	}
	return nil
}

// `Engine` returns the consensus engine. For our use case, this never gets called.
func (cc *chainContext) Engine() consensus.Engine {
	return nil
}
