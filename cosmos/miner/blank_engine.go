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
//
//nolint:revive // boilerplate for now.
package miner

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	consensus "github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
)

// MockEngine is a mock implementation of the Engine interface.
type MockEngine struct{}

// Author is a mock implementation.
func (m *MockEngine) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

// VerifyHeader is a mock implementation.
func (m *MockEngine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	// Set the correct difficulty
	header.Difficulty = new(big.Int).SetUint64(1)

	return nil
}

// VerifyHeaders is a mock implementation.
func (m *MockEngine) VerifyHeaders(
	chain consensus.ChainHeaderReader, headers []*types.Header) (chan<- struct{}, <-chan error) {
	for _, h := range headers {
		if err := m.VerifyHeader(chain, h); err != nil {
			return nil, nil
		}
	}
	return nil, nil
}

// VerifyUncles is a mock implementation.
func (m *MockEngine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return nil
}

// Prepare is a mock implementation.
func (m *MockEngine) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	header.Difficulty = new(big.Int).SetUint64(0)
	return nil
}

// Finalize is a mock implementation.
func (m *MockEngine) Finalize(chain consensus.ChainHeaderReader,
	header *types.Header, state state.StateDBI, txs []*types.Transaction,
	uncles []*types.Header, withdrawals []*types.Withdrawal) {
}

// FinalizeAndAssemble is a mock implementation.
func (m *MockEngine) FinalizeAndAssemble(chain consensus.ChainHeaderReader,
	header *types.Header, state state.StateDBI, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt,
	withdrawals []*types.Withdrawal) (*types.Block, error) {
	return types.NewBlock(header, txs, uncles, receipts, trie.NewStackTrie(nil)), nil
}

// Seal is a mock implementation.
func (m *MockEngine) Seal(chain consensus.ChainHeaderReader,
	block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	sealedBlock := block // .seal()
	results <- sealedBlock
	return nil
}

// SealHash is a mock implementation.
func (m *MockEngine) SealHash(header *types.Header) common.Hash {
	return header.Hash()
}

// CalcDifficulty is a mock implementation.
func (m *MockEngine) CalcDifficulty(chain consensus.ChainHeaderReader,
	time uint64, parent *types.Header) *big.Int {
	return big.NewInt(0)
}

// APIs is a mock implementation.
func (m *MockEngine) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return nil
}

// Close is a mock implementation.
func (m *MockEngine) Close() error {
	return nil
}
