// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
