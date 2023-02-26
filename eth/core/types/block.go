// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import (
	"unsafe"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// var _ ethapi.Block = &StargazerBlock{}

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
func (sb *StargazerBlock) TxIndex() int {
	return len(sb.txs)
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

// `GetReceipts` returns the receipts of the block.
func (sb *StargazerBlock) GetReceipts() Receipts {
	// TODO: fill bloom if empty (for old blocks) that were
	// marshaled without bloom.
	return sb.receipts
}

// `GetTransactions` returns the transactions of the block.
func (sb *StargazerBlock) GetTransactions() Transactions {
	return sb.txs
}

// `Finalize` sets the gas used, transaction hash, receipt hash, and optionally bloom of the block
// header.
func (sb *StargazerBlock) Finalize(gasUsed uint64) {
	hasher := trie.NewStackTrie(nil)
	sb.StargazerHeader.GasUsed = gasUsed
	if len(sb.txs) == 0 {
		sb.StargazerHeader.TxHash = EmptyRootHash
		sb.StargazerHeader.ReceiptHash = EmptyRootHash
	} else {
		sb.StargazerHeader.TxHash = DeriveSha(sb.txs, hasher)
		sb.StargazerHeader.ReceiptHash = DeriveSha(sb.receipts, hasher)
		sb.StargazerHeader.Bloom = CreateBloom(sb.receipts)
	}
}

// `EthBlock` represents a ethereum-like block that can be encoded to raw bytes.
func (sb *StargazerBlock) EthBlock() *Block {
	if sb == nil {
		return nil
	}
	eb := NewBlock(sb.Header, sb.txs, nil, sb.receipts, trie.NewStackTrie(nil))
	return eb
}
