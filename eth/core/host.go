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

	"github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/state"
	"github.com/berachain/polaris/eth/core/types"
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// PolarisHostChain defines the plugins that the chain running the Polaris EVM should implement.
type PolarisHostChain interface {
	// GetBlockPlugin returns the `BlockPlugin` of the Polaris host chain.
	GetBlockPlugin() BlockPlugin
	// GetHistoricalPlugin returns the OPTIONAL `HistoricalPlugin` of the Polaris host chain.
	GetHistoricalPlugin() HistoricalPlugin
	// GetPrecompilePlugin returns the OPTIONAL `PrecompilePlugin` of the Polaris host chain.
	GetPrecompilePlugin() PrecompilePlugin
	// GetStatePlugin returns the `StatePlugin` of the Polaris host chain.
	GetStatePluginFactory() StatePluginFactory
	// Version()
	Version() string
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
		// BlockPlugin to a default.
		libtypes.Preparable
		// GetHeaderByNumber returns the block header at the given block number.
		GetHeaderByNumber(uint64) (*ethtypes.Header, error)
		// GetHeaderByHash returns the block header with the given block hash.
		GetHeaderByHash(common.Hash) (*ethtypes.Header, error)
		// StoreHeader stores the block header at the given block number.
		StoreHeader(*ethtypes.Header) error
	}

	// StatePlugin defines the methods that the chain running Polaris EVM should implement.
	StatePlugin interface {
		state.Plugin
		// StateAtBlockNumber returns the state at the given block height.
		StateAtBlockNumber(uint64) (StatePlugin, error)
		SetStateOverride(ctx context.Context)
		GetOverridenState() StatePlugin
	}

	StatePluginFactory interface {
		NewPluginAtBlockNumber(int64) (StatePlugin, error)
		NewPluginWithMode(state.Mode) StatePlugin
		NewPluginFromContext(context.Context) StatePlugin
		SetLatestQueryContext(ctx context.Context)
		SetLatestMiningContext(ctx context.Context)
		SetInsertChainContext(ctx context.Context)
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
		GetBlockByNumber(uint64) (*ethtypes.Block, error)
		// GetBlockByHash returns the block at the given block hash.
		GetBlockByHash(common.Hash) (*ethtypes.Block, error)
		// GetTransactionByHash returns the transaction lookup entry at the given transaction
		// hash.
		GetTransactionByHash(common.Hash) (*types.TxLookupEntry, error)
		// GetReceiptByHash returns the receipts at the given block hash.
		GetReceiptsByHash(common.Hash) (ethtypes.Receipts, error)
		// StoreBlock stores the given block.
		StoreBlock(*ethtypes.Block) error
		// StoreReceipts stores the receipts for the given block hash.
		StoreReceipts(common.Hash, ethtypes.Receipts) error
		// StoreTransactions stores the transactions for the given block hash.
		StoreTransactions(uint64, common.Hash, ethtypes.Transactions) error
	}

	// PrecompilePlugin defines the methods that the chain running Polaris EVM should implement
	// in order to support running their own stateful precompiled contracts. Implementing this
	// plugin is optional.
	PrecompilePlugin = precompile.Plugin
)
