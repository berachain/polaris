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
	"math/big"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// import (
// 	"math"
// 	"math/big"

// 	"github.com/berachain/stargazer/lib/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/trie"
// )

type (
	// 	EvmHeader struct {
	// 		Number     *big.Int
	// 		Hash       common.Hash
	// 		ParentHash common.Hash
	// 		Root       common.Hash
	// 		TxHash     common.Hash
	// 		Time       uint64
	// 		Coinbase   common.Address

	// 		GasLimit uint64
	// 		GasUsed  uint64

	// 		BaseFee *big.Int
	// 	}

	StargazerEVMBlock struct {
		*types.StargazerHeader

		Transactions types.Transactions
		Receipts     []*types.Receipt
	}
)

// `NewStargazerEVMBlock` creates a new EvmBlock from the given header and transactions.
func NewStargazerEVMBlock(h *types.StargazerHeader, txs types.Transactions) *StargazerEVMBlock {
	b := &StargazerEVMBlock{
		StargazerHeader: h,
		Transactions:    txs,
	}

	if len(txs) == 0 {
		b.StargazerHeader.TxHash = types.EmptyRootHash
	} else {
		b.StargazerHeader.TxHash = types.DeriveSha(txs, trie.NewStackTrie(nil))
	}

	return b
}

// // ToEvmHeader converts inter.Block to EvmHeader.
// func ToEvmHeader(block *inter.Block, index idx.Block, prevHash hash.Event, rules opera.Rules) *EvmHeader {
// 	baseFee := rules.Economy.MinGasPrice
// 	if !rules.Upgrades.London {
// 		baseFee = nil
// 	}
// 	return &EvmHeader{
// 		Hash:       common.Hash(block.Atropos),
// 		ParentHash: common.Hash(prevHash),
// 		Root:       common.Hash(block.Root),
// 		Number:     big.NewInt(int64(index)),
// 		Time:       block.Time,
// 		GasLimit:   math.MaxUint64,
// 		GasUsed:    block.GasUsed,
// 		BaseFee:    baseFee,
// 	}
// }

// `Header` returns the header of the block.
func (b *StargazerEVMBlock) Header() *types.StargazerHeader {
	if b == nil {
		return nil
	}
	// copy values
	h := *b.StargazerHeader

	// copy refs
	h.Number = new(big.Int).Set(b.Number)
	if b.BaseFee != nil {
		h.BaseFee = new(big.Int).Set(b.BaseFee)
	}

	return &h
}

// `NumberU64` returns the block number as a uint64.
func (b *StargazerEVMBlock) NumberU64() uint64 {
	return b.Number.Uint64()
}

// `EthBlock` returns the underlying ethereum block.
func (b *StargazerEVMBlock) EthBlock() *types.Block {
	if b == nil {
		return nil
	}
	return types.NewBlock(b.StargazerHeader.Header, b.Transactions, nil, nil, trie.NewStackTrie(nil))
}
