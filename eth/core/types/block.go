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
	"unsafe"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// `StargazerBlock` represents a ethereum-like block that can be encoded to raw bytes.
//
//go:generate rlpgen -type StargazerBlock -out block.rlpgen.go -decoder
type StargazerBlock struct {
	*StargazerHeader
	Transactions Transactions
	Receipts     StargazerReceipts
}

// `NewStargazerBlock` creates a new StargazerBlock from the given header and transactions.
func NewStargazerBlock(h *StargazerHeader, txs Transactions, rs StargazerReceipts) *StargazerBlock {
	b := &StargazerBlock{
		StargazerHeader: h,
		Transactions:    txs,
		Receipts:        rs,
	}

	return b
}

func (b *StargazerBlock) SetGasUsed(gas uint64) {
	b.GasUsed = gas
}

func (b *StargazerBlock) SetReceiptHash() {
	if b.Receipts.Len() > 0 {
		b.StargazerHeader.ReceiptHash = DeriveSha(
			*(*(Receipts))((unsafe.Pointer(&b.Receipts.Receipts))), trie.NewStackTrie(nil), //#nosec:G103
		)
	} else {
		b.StargazerHeader.ReceiptHash = EmptyRootHash
	}
}

// `CreateBloom` creates the bloom filter for the block.
func (b *StargazerBlock) CreateBloom() {
	//#nosec:G103
	b.StargazerHeader.Bloom = types.CreateBloom(*(*(Receipts))((unsafe.Pointer(&b.Receipts.Receipts))))
}

// `UnmarshalBinary` decodes a block from the Ethereum RLP format.
func (b *StargazerBlock) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, b)
}

// `MarshalBinary` encodes the block into the Ethereum RLP format.
func (b *StargazerBlock) MarshalBinary() ([]byte, error) {
	bz, err := rlp.EncodeToBytes(b)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// `EthBlock` returns the block as an Ethereum Block.
func (b *StargazerBlock) EthBlock() *Block {
	if b == nil {
		return nil
	}
	return NewBlock(b.StargazerHeader.Header, b.Transactions, nil, nil, trie.NewStackTrie(nil))
}
