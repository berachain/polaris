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
	coretypes "github.com/berachain/stargazer/core/types"
)

// `BloomBuilder` is a wrapper around `StateDB` that is used to persist data between
// transactions within a block.
type BloomBuilder struct {
	// blockBloom is scratch space for building the
	// bloom filter for an entire block
	bloom *coretypes.Bloom
}

// `AddLogsToBloom` builds the bloom filter for the provided set of logs.
// It also adds the bloom filter to the block bloom filter.
func (bb *BloomBuilder) AddLogsToBloom(logs []*coretypes.Log) *coretypes.Bloom {
	// Calculate bloom for current transaction.
	txBloomBz := coretypes.LogsBloom(logs)

	// Add bloom to block bloom filter.
	bb.bloom.Add(txBloomBz)

	// Convert bytes to bloom and return this tx's bloom.
	bloom := coretypes.BytesToBloom(txBloomBz)
	return &bloom
}

// `Reset` resets the bloom builder.
func (bb *BloomBuilder) Reset() {
	bb.bloom = &coretypes.Bloom{}
}

// `GetBloom` returns the currently built bloom filter.
func (bb *BloomBuilder) GetBloom() *coretypes.Bloom {
	return bb.bloom
}
