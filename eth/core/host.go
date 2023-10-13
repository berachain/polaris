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

package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// PolarisHostChain defines the plugins that the chain running the Polaris EVM should implement.
type PolarisHostChain interface {
	// GetBlockPlugin returns the `BlockPlugin` of the Polaris host chain.
	GetBlockPlugin() BlockPlugin
	// GetEnginePlugin() returns the `EnginePlugin` of the Polaris host chain.
	GetEnginePlugin() EnginePlugin
	// GetHistoricalPlugin returns the OPTIONAL `HistoricalPlugin` of the Polaris host chain.
	GetHistoricalPlugin() HistoricalPlugin
	// GetPrecompilePlugin returns the OPTIONAL `PrecompilePlugin` of the Polaris host chain.
	GetPrecompilePlugin() PrecompilePlugin
	// GetStatePlugin returns the `StatePlugin` of the Polaris host chain.
	GetStatePlugin() StatePlugin
}

// =============================================================================
// Mandatory Plugins
// =============================================================================

// The following plugins should be implemented by the chain running the Polaris EVM and exposed via
// the `PolarisHostChain` interface. All plugins should be resettable with a given context.
type (
	// BlockPlugin defines the methods that the chain running the Polaris EVM should implement to
	// support getting and setting block headers.
	BlockPlugin interface {
		// BlockPlugin implements `libtypes.Preparable`. Calling `Prepare` should reset the
		// BlockPlugin to a default state.
		libtypes.Preparable
		// GetNewBlockMetadata returns a new block metadata (coinbase, timestamp) for the given
		// block number.
		GetNewBlockMetadata(uint64) (common.Address, uint64)
		// GetHeaderByNumber returns the block header at the given block number.
		GetHeaderByNumber(uint64) (*types.Header, error)
		// GetHeaderByHash returns the block header with the given block hash.
		GetHeaderByHash(common.Hash) (*types.Header, error)
		// StoreHeader stores the block header at the given block number.
		StoreHeader(*types.Header) error
		// BaseFee returns the base fee of the current block.
		BaseFee() *big.Int
	}

	// EnginePlugin defines methods that allow the chain to have insight into the underlying
	// consensus engine of the host chain.
	EnginePlugin interface {
		// SyncProgress returns the current sync progress of the host chain.
		SyncProgress(ctx context.Context) (ethereum.SyncProgress, error)
		// IsListening returns whether or not the host chain is listening for new blocks.
		Listening(ctx context.Context) (bool, error)
		// PeerCount returns the current number of peers connected to the host chain.
		PeerCount(ctx context.Context) (uint64, error)
	}

	// StatePlugin defines the methods that the chain running Polaris EVM should implement.
	StatePlugin interface {
		state.Plugin
		// StateAtBlockNumber returns the state at the given block height.
		StateAtBlockNumber(uint64) (StatePlugin, error)
	}
)

// =============================================================================
// Optional Plugins
// =============================================================================

// `The following plugins are OPTIONAL to be implemented by the chain running Polaris EVM.
type (
	// HistoricalPlugin defines the methods that the chain running Polaris EVM should implement
	// in order to support storing historical blocks, receipts, and transactions. This plugin will
	// be used by the RPC backend to support certain methods on the Ethereum JSON RPC spec.
	// Implementing this plugin is optional.
	HistoricalPlugin interface {
		// HistoricalPlugin implements `libtypes.Preparable`.
		libtypes.Preparable
		// GetBlockByNumber returns the block at the given block number.
		GetBlockByNumber(uint64) (*types.Block, error)
		// GetBlockByHash returns the block at the given block hash.
		GetBlockByHash(common.Hash) (*types.Block, error)
		// GetTransactionByHash returns the transaction lookup entry at the given transaction
		// hash.
		GetTransactionByHash(common.Hash) (*types.TxLookupEntry, error)
		// GetReceiptByHash returns the receipts at the given block hash.
		GetReceiptsByHash(common.Hash) (types.Receipts, error)
		// StoreBlock stores the given block.
		StoreBlock(*types.Block) error
		// StoreReceipts stores the receipts for the given block hash.
		StoreReceipts(common.Hash, types.Receipts) error
		// StoreTransactions stores the transactions for the given block hash.
		StoreTransactions(uint64, common.Hash, types.Transactions) error
	}

	// PrecompilePlugin defines the methods that the chain running Polaris EVM should implement
	// in order to support running their own stateful precompiled contracts. Implementing this
	// plugin is optional.
	PrecompilePlugin = precompile.Plugin
)
