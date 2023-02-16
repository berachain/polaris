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

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// `initialTransactionsCapacity` is the initial capacity of the transactions, receipts slice.
// TODO: figre out optimal value.
const initialTransactionsCapacity = 256

// `StargazerBlock` represents a ethereum-like block that can be encoded to raw bytes.
//
//go:generate rlpgen -type StargazerBlock -out block.rlpgen.go -decoder
type StargazerBlock struct {
	*StargazerHeader
	txs      Transactions
	receipts Receipts
	// `logIndex` is the index of the current log in the current block
	logIndex uint
}

// `NewStargazerBlock` creates a new StargazerBlock from the given header.
func NewStargazerBlock(header *StargazerHeader) *StargazerBlock {
	return &StargazerBlock{
		StargazerHeader: header,
		txs:             make(Transactions, 0, initialTransactionsCapacity),
		receipts:        make(Receipts, 0, initialTransactionsCapacity),
	}
}

// `TxIndex` returns the current transaction index in the block.
func (sb *StargazerBlock) TxIndex() uint {
	return uint(len(sb.txs))
}

func (sb *StargazerBlock) LogIndex() uint {
	return sb.logIndex
}

// `AppendTx` appends a transaction and receipt to the block.
func (sb *StargazerBlock) AppendTx(tx *Transaction, receipt *Receipt) {
	sb.txs = append(sb.txs, tx)
	sb.receipts = append(sb.receipts, receipt)
	sb.logIndex += uint(len(receipt.Logs))
}

// `UnmarshalBinary` decodes a block from the Ethereum RLP format.
func (sb *StargazerBlock) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, sb)
}

// `MarshalBinary` encodes the block into the Ethereum RLP format.
func (sb *StargazerBlock) MarshalBinary() ([]byte, error) {
	bz, err := rlp.EncodeToBytes(sb)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// `GetReceiptsForStorage` converts a list of `Receipt`s to a `StargazerReceipts`.
func (sb *StargazerBlock) GetReceiptsForStorage() []*ReceiptForStorage {
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	return *(*[]*ReceiptForStorage)(unsafe.Pointer(&sb.receipts))
}

// `Finalize` sets the gas used, transaction hash, receipt hash, and optionally bloom of the block
// header.
func (sb *StargazerBlock) Finalize(gasUsed uint64) {
	hasher := trie.NewStackTrie(nil)
	sb.Header.GasUsed = gasUsed
	if len(sb.txs) == 0 {
		sb.Header.TxHash = EmptyRootHash
		sb.Header.ReceiptHash = EmptyRootHash
	} else {
		sb.Header.TxHash = DeriveSha(sb.txs, hasher)
		sb.Header.ReceiptHash = DeriveSha(sb.receipts, hasher)
		sb.Header.Bloom = CreateBloom(sb.receipts)
	}
}
