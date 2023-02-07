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

package types

import (
	"github.com/berachain/stargazer/lib/common"
)

// `StargazerHeader` represents a wrapped Ethereum header that allows for specifying a custom
// blockhash to make it compatible with a non-ethereum chain.
//
//go:generate rlpgen -type StargazerHeader -out header.rlpgen.go -decoder
type StargazerHeader struct {
	// `Header` is an embedded ethereum header.
	*Header
	// `CachedHash` is the cached hash of the header
	CachedHash common.Hash
}

// `NewStargazerHeader` returns a `StargazerHeader`.
func NewStargazerHeader(header *Header, hash common.Hash) *StargazerHeader {
	return &StargazerHeader{Header: header, CachedHash: hash}
}

// `Author` returns the address of the original block producer.
func (h *StargazerHeader) Author() common.Address {
	return h.Coinbase
}

// `Hash` returns the block hash of the header, we override the geth implementation
// to use the cached hash, as the implementing chain might want to use it's real block hash
// opposed to hashing the "fake" header.
func (h *StargazerHeader) Hash() common.Hash {
	if h.CachedHash == (common.Hash{}) {
		h.CachedHash = h.Header.Hash()
	}
	return h.CachedHash
}

// `SetHash` sets the hash of the header.
func (h *StargazerHeader) SetHash(hash common.Hash) {
	h.CachedHash = hash
}
