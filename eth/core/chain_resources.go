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
	"errors"
	"fmt"

	"github.com/berachain/polaris/eth/core/state"

	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/params"
)

// ChainResources is the interface that defines functions for code paths within the chain to
// acquire resources to use in execution such as StateDBss and EVMss.
type ChainResources interface {
	// state snapshots
	StateAtBlockNumber(uint64) (state.StateDB, error)
	StateAt(root common.Hash) (state.StateDB, error)
	GetOverridenState() (state.StateDB, error)

	// state for tracing
	StateAtBlock(
		_ context.Context, block *ethtypes.Block, _ uint64,
		_ state.StateDB, _ bool, _ bool,
	) (state.StateDB, tracers.StateReleaseFunc, error)
	StateAtTransaction(
		ctx context.Context, block *ethtypes.Block,
		txIndex int, reexec uint64,
	) (*gethcore.Message, vm.BlockContext, state.StateDB, tracers.StateReleaseFunc, error)

	// vm/chain config
	GetVMConfig() *vm.Config
	Config() *params.ChainConfig
}

// StateAt returns a statedb configured to read what the state of the blockchain is/was at a given.
func (bc *blockchain) StateAt(common.Hash) (state.StateDB, error) {
	return nil, errors.New("StateAt is not implemented in polaris due state root")
}

// Used by geth miner to build the block (can rename to GetMinerState).
func (bc *blockchain) GetOverridenState() (state.StateDB, error) {
	return state.NewStateDB(bc.spf.NewPluginWithMode(state.Miner), bc.pp), nil
}

// StateAtBlockNumber returns a statedb configured to read what the state of the blockchain is/was
// at a given block number.
func (bc *blockchain) StateAtBlockNumber(number uint64) (state.StateDB, error) {
	sp, err := bc.spf.NewPluginAtBlockNumber(int64(number))
	if err != nil {
		return nil, err
	}

	return state.NewStateDB(sp, bc.pp), nil
}

// HasBlockAndState checks if the blockchain has a block and its state at
// a given hash and number.
func (bc *blockchain) HasBlockAndState(hash common.Hash, number uint64) bool {
	// Check for State.
	if sdb, err := bc.StateAt(hash); sdb == nil || err == nil {
		sdb, err = bc.StateAtBlockNumber(number)
		if sdb == nil || err != nil {
			return false
		}
	}

	// Check for Block.
	if block := bc.GetBlockByNumber(number); block == nil {
		block = bc.GetBlockByHash(hash)
		if block == nil {
			return false
		}
	}
	return true
}

// StateAtBlock retrieves the state database associated with a certain block.
// If no state is locally available for the given block, a number of blocks
// are attempted to be reexecuted to generate the desired state. The optional
// base layer statedb can be provided which is regarded as the statedb of the
// parent block.
//
// An additional release function will be returned if the requested state is
// available. Release is expected to be invoked when the returned state is no
// longer needed. Its purpose is to prevent resource leaking. Though it can be
// noop in some cases.
//
// Parameters:
//   - block:      The block for which we want the state(state = block.Root)
//   - reexec:     The maximum number of blocks to reprocess trying to obtain the desired state
//   - base:       If the caller is tracing multiple blocks, the caller can provide the parent
//     state continuously from the callsite.
//   - readOnly:   If true, then the live 'blockchain' state database is used. No mutation should
//     be made from caller, e.g. perform Commit or other 'save-to-disk' changes.
//     Otherwise, the trash generated by caller may be persisted permanently.
//   - preferDisk: This arg can be used by the caller to signal that even though the 'base' is
//     provided, it would be preferable to start from a fresh state, if we have it
//     on disk.
func (bc *blockchain) StateAtBlock(
	_ context.Context, block *ethtypes.Block, _ uint64,
	_ state.StateDB, _ bool, _ bool,
) (state.StateDB, tracers.StateReleaseFunc, error) {
	// Check if the requested state is available in the live chain.
	statedb, err := bc.StateAtBlockNumber(block.Number().Uint64())
	if err != nil {
		// If there is an error, it means the state is not available.
		// TODO: Historic state is not supported in path-based scheme.
		// Fully archive node in pbss will be implemented by relying
		// on state history, but needs more work on top.
		return nil, nil, errors.Join(
			err, errors.New("historical state not available in path scheme yet"),
		)
	}

	// If there is no error, return the state, a no-op function, and no error.
	return statedb, func() {}, nil
}

// StateAtTransaction returns the execution environment of a certain transaction.
func (bc *blockchain) StateAtTransaction(
	ctx context.Context, block *ethtypes.Block,
	txIndex int, reexec uint64,
) (*gethcore.Message, vm.BlockContext, state.StateDB, tracers.StateReleaseFunc, error) {
	// Short circuit if it's genesis block.
	if block.NumberU64() == 0 {
		return nil, vm.BlockContext{}, nil, nil, errors.New("no transaction in genesis")
	}
	// Create the parent state database
	parent := bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return nil, vm.BlockContext{}, nil, nil, fmt.Errorf("parent %#x not found", block.ParentHash())
	}
	// Lookup the statedb of parent block from the live database,
	// otherwise regenerate it on the flight.
	statedb, release, err := bc.StateAtBlock(ctx, parent, reexec, nil, true, false)
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, err
	}
	if txIndex == 0 && len(block.Transactions()) == 0 {
		return nil, vm.BlockContext{}, statedb, release, nil
	}
	// Recompute transactions up to the target index.
	signer := ethtypes.MakeSigner(bc.Config(), block.Number(), block.Time())
	for idx, tx := range block.Transactions() {
		// Assemble the transaction call message and return if the requested offset
		msg, _ := gethcore.TransactionToMessage(tx, signer, block.BaseFee())
		txContext := gethcore.NewEVMTxContext(msg)
		context := gethcore.NewEVMBlockContext(block.Header(), bc, nil)
		if idx == txIndex {
			return msg, context, statedb, release, nil
		}
		// Not yet the searched for transaction, execute on top of the current state
		vmenv := vm.NewEVM(context, txContext, statedb, bc.Config(), vm.Config{})
		statedb.SetTxContext(tx.Hash(), idx)
		if _, err = gethcore.ApplyMessage(vmenv,
			msg, new(gethcore.GasPool).AddGas(tx.Gas())); err != nil {
			return nil, vm.BlockContext{}, nil, nil,
				fmt.Errorf("transaction %s failed: %w", tx.Hash().Hex(), err)
		}
		// Ensure any modifications are committed to the state
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		statedb.Finalise(vmenv.ChainConfig().IsEIP158(block.Number()))
	}
	return nil, vm.BlockContext{}, nil, nil,
		fmt.Errorf("transaction index %d out of range for block %#x", txIndex, block.Hash())
}

// GetVMConfig returns the vm.Config for the current chain.
func (bc *blockchain) GetVMConfig() *vm.Config {
	return bc.vmConfig
}

// Config returns the Ethereum chain config from the host chain.
func (bc *blockchain) Config() *params.ChainConfig {
	return bc.config
}
