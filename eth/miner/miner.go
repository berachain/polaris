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

package miner

import (
	"context"

	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// Backend wraps all methods required for mining. Only full node is capable
// to offer all the functions here.
type Backend interface {
	// Blockchain returns the blockchain instance.
	Blockchain() core.Blockchain
}

// Miner defines the interface for a Polaris miner.
type Miner interface {
	// Prepare prepares the miner for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, uint64) *types.Header

	// ProcessTransaction processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*core.ExecutionResult, error)

	// Finalize is called after the last tx in the block.
	Finalize(context.Context) error
}

// miner implements the Miner interface.
type miner struct {
	backend Backend
}

// NewMiner creates a new Miner instance.
func NewMiner(backend Backend) Miner {
	return &miner{
		backend: backend,
	}
}

// Prepare prepares the blockchain for processing a new block at the given height.
func (m *miner) Prepare(ctx context.Context, number uint64) *types.Header {
	return m.backend.Blockchain().Prepare(ctx, number)
}

// ProcessTransaction processes the given transaction and returns the receipt after applying
// the state transition. This method is called for each tx in the block.
func (m *miner) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*core.ExecutionResult, error) {
	return m.backend.Blockchain().ProcessTransaction(ctx, tx)
}

// Finalize is called after the last tx in the block.
func (m *miner) Finalize(ctx context.Context) error {
	block, receipts, logs, err := m.backend.Blockchain().GetProcessor().Finalize(ctx)
	if err != nil {
		return err
	}

	return m.backend.Blockchain().InsertBlock(block, receipts, logs)
}
