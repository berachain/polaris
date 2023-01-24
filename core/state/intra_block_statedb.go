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

package state

import (
	"context"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
)

// `IntraBlockStateDB` is a wrapper around `StateDB` that is used to persist data between
// transactions within a block.
type IntraBlockStateDB struct {
	// StateDB is the underlying state database.
	vm.StargazerStateDB

	// blockBloom is scratch space for building the
	// bloom filter for an entire block
	blockBloom *coretypes.Bloom
}

// `Reset` clears the journal and other state objects. It also clears the
// refund counter and the access list.
func (ibs *IntraBlockStateDB) Reset(ctx context.Context) {
	ibs.StargazerStateDB.Reset(ctx)
	ibs.blockBloom = new(coretypes.Bloom)
}

// `BuildBloomFilterForTxn` builds the bloom filter for the current transaction.
// It also adds the bloom filter to the block bloom filter.
func (ibs *IntraBlockStateDB) BuildBloomFilterForTxn() coretypes.Bloom {
	// Calculate bloom for current transaction.
	txBloomBz := coretypes.LogsBloom(ibs.StargazerStateDB.Logs())

	// Add bloom to block bloom filter.
	ibs.blockBloom.Add(txBloomBz)

	// Convert bytes to bloom and return this tx's bloom.
	return coretypes.BytesToBloom(txBloomBz)
}

// `GetBlockBloom` returns the currently built bloom filter.
func (ibs *IntraBlockStateDB) GetBlockBloom() *coretypes.Bloom {
	return ibs.blockBloom
}
