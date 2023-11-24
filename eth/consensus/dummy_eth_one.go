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
package consensus

import (
	"math/big"

	"github.com/berachain/polaris/eth/core/state"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
)

type Engine consensus.Engine

// DummyEthOne is a dummy implementation of the consensus.Engine interface.
var _ Engine = (*DummyEthOne)(nil)

// DummyEthOne is a mock implementation of the Engine interface.
type DummyEthOne struct{}

// Author is a mock implementation.
func (m *DummyEthOne) Author(header *ethtypes.Header) (common.Address, error) {
	return common.Address{}, nil
}

// VerifyHeader is a mock implementation.
func (m *DummyEthOne) VerifyHeader(
	chain consensus.ChainHeaderReader,
	header *ethtypes.Header,
) error {
	// Set the correct difficulty
	header.Difficulty = new(big.Int).SetUint64(1)
	return nil
}

// VerifyHeaders is a mock implementation.
func (m *DummyEthOne) VerifyHeaders(
	chain consensus.ChainHeaderReader, headers []*ethtypes.Header) (chan<- struct{}, <-chan error) {
	for _, h := range headers {
		if err := m.VerifyHeader(chain, h); err != nil {
			return nil, nil
		}
	}
	return nil, nil
}

// VerifyUncles is a mock implementation.
func (m *DummyEthOne) VerifyUncles(chain consensus.ChainReader, block *ethtypes.Block) error {
	return nil
}

// Prepare is a mock implementation.
func (m *DummyEthOne) Prepare(chain consensus.ChainHeaderReader, header *ethtypes.Header) error {
	header.Difficulty = new(big.Int).SetUint64(0)
	return nil
}

// Finalize is a mock implementation.
func (m *DummyEthOne) Finalize(chain consensus.ChainHeaderReader,
	header *ethtypes.Header, state state.StateDB, txs []*ethtypes.Transaction,
	uncles []*ethtypes.Header, withdrawals []*ethtypes.Withdrawal) {
}

// FinalizeAndAssemble is a mock implementation.
func (m *DummyEthOne) FinalizeAndAssemble(chain consensus.ChainHeaderReader,
	header *ethtypes.Header, state state.StateDB, txs []*ethtypes.Transaction,
	uncles []*ethtypes.Header, receipts []*ethtypes.Receipt,
	withdrawals []*ethtypes.Withdrawal) (*ethtypes.Block, error) {
	return ethtypes.NewBlock(header, txs, uncles, receipts, trie.NewStackTrie(nil)), nil
}

// Seal is a mock implementation.
func (m *DummyEthOne) Seal(chain consensus.ChainHeaderReader,
	block *ethtypes.Block, results chan<- *ethtypes.Block, stop <-chan struct{}) error {
	sealedBlock := block // .seal()
	results <- sealedBlock
	return nil
}

// SealHash is a mock implementation.
func (m *DummyEthOne) SealHash(header *ethtypes.Header) common.Hash {
	return header.Hash()
}

// CalcDifficulty is a mock implementation.
func (m *DummyEthOne) CalcDifficulty(chain consensus.ChainHeaderReader,
	time uint64, parent *ethtypes.Header) *big.Int {
	return big.NewInt(0)
}

// APIs is a mock implementation.
func (m *DummyEthOne) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return nil
}

// Close is a mock implementation.
func (m *DummyEthOne) Close() error {
	return nil
}
