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
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/crypto"
)

// `StateProcessor` is responsible for processing blocks, transactions, and updating the state.
type StateProcessor struct {
	// `mtx` is used to make sure we don't try to prepare a new block before finalizing the
	// previous one.
	mtx sync.Mutex
	// `bp` provides block functions from the underlying chain the EVM is running on
	bp BlockPlugin
	// `gp` provides gas functions from the underlying chain the EVM is running on
	gp GasPlugin
	// `cp` provides configuration functions from the underlying chain the EVM is running on
	cp ConfigurationPlugin
	// `pp` is responsible for keeping track of the stateful precompile containers that are
	// available to the EVM and executing them.
	pp PrecompilePlugin
	// `signer` is the signer used to verify transaction signatures.
	signer types.Signer
	// `evm ` is the EVM that is used to process transactions.
	evm vm.StargazerEVM
	// `vmConfig` is the configuration for the EVM.
	vmConfig vm.Config
	// `statedb` is the state database that is used to mange state during transactions.
	statedb vm.StargazerStateDB
	// `block` represents the current block being processed.
	block *types.StargazerBlock
	// `commit` indicates whether the state processor should commit the state after processing a tx
	commit bool
}

// `NewStateProcessor` creates a new state processor with the given host, statedb, vmConfig, and
// commit flag.
func NewStateProcessor(
	host StargazerHostChain,
	statedb vm.StargazerStateDB,
	vmConfig vm.Config,
	commit bool,
) *StateProcessor {
	sp := &StateProcessor{
		mtx:      sync.Mutex{},
		bp:       host.GetBlockPlugin(),
		gp:       host.GetGasPlugin(),
		cp:       host.GetConfigurationPlugin(),
		pp:       host.GetPrecompilePlugin(),
		vmConfig: vmConfig,
		statedb:  statedb,
		commit:   commit,
	}
	return sp
}

// ==============================================================================
// Block, Tx Lifecycle
// ==============================================================================

// `Prepare` prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(ctx context.Context, header *types.StargazerHeader) {
	// We lock the state processor as a safety measure to ensure that Prepare is not called again
	// before finalize.
	sp.mtx.Lock()

	// Prepare the plugins for the new block.
	sp.bp.Prepare(ctx)
	sp.cp.Prepare(ctx)
	sp.gp.Prepare(ctx)

	// Build a block object so we can track that status of the block as we process it.
	sp.block = types.NewStargazerBlock(header)
	chainConfig := sp.cp.ChainConfig()

	// We must re-create the signer since we are processing a new block and the block number has increased.
	sp.signer = types.MakeSigner(chainConfig, sp.block.Number)

	// Setup the EVM for this block.
	sp.vmConfig.ExtraEips = sp.cp.ExtraEips()
	sp.evm = vm.NewStargazerEVM(
		sp.newEVMBlockContext(ctx),
		vm.TxContext{},
		sp.statedb,
		chainConfig,
		sp.vmConfig,
		sp.pp,
	)
}

// `ProcessTransaction` applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	msg, err := tx.AsMessage(sp.signer, sp.block.BaseFee)
	if err != nil {
		return nil, fmt.Errorf("could not apply tx %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Create a new context to be used in the EVM environment. We also must reset the StateDB and
	// precompile manager, which resets the state and precompile plugins, for the tx.
	txContext := NewEVMTxContext(msg)
	sp.evm.SetTxContext(txContext)
	sp.statedb.Reset(ctx)
	sp.pp.Reset(ctx)

	// Apply the state transition.
	result, err := ApplyMessage(sp.evm, sp.gp, msg, sp.commit)
	if err != nil {
		return nil, fmt.Errorf("could apply message %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	receipt := &types.Receipt{
		Type:             tx.Type(),
		PostState:        nil, // TODO: Should we do something with PostState?
		TxHash:           tx.Hash(),
		GasUsed:          result.UsedGas,
		BlockHash:        sp.block.Hash(),
		BlockNumber:      sp.block.Number,
		TransactionIndex: sp.block.TxIndex(),
	}

	// Gas from this transaction was added to the gasPlugin in `ApplyMessageAndCommit`
	// And thus CumulativeGasUsed should include gas from all prior transactions in the
	// block, plus the gas consumed during this one.
	receipt.CumulativeGasUsed = sp.gp.CumulativeGasUsed()

	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(txContext.Origin, tx.Nonce())
	}

	// Set the receipt logs and create the bloom filter.
	receipt.Logs = sp.statedb.BuildLogsAndClear(
		receipt.TxHash, receipt.BlockHash, sp.block.TxIndex(), sp.block.LogIndex(),
	)
	receipt.Bloom = types.BytesToBloom(types.LogsBloom(receipt.Logs))

	// Update the block information.
	sp.block.AppendTx(tx, receipt)
	return receipt, nil
}

// `Finalize` finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(ctx context.Context) (*types.StargazerBlock, error) {
	// We unlock the state processor to ensure that the state is consistent.
	defer sp.mtx.Unlock()

	sp.block.Finalize(sp.gp.CumulativeGasUsed())
	return sp.block, nil
}

// ===========================================================================
// Utilities
// ===========================================================================

// `newEVMBlockContext` creates a new block context for use in the EVM.
func (sp *StateProcessor) newEVMBlockContext(ctx context.Context) vm.BlockContext {
	var baseFee *big.Int
	// Copy the baseFee to avoid side effects.
	if sp.block.StargazerHeader.BaseFee != nil {
		baseFee = new(big.Int).Set(sp.block.StargazerHeader.BaseFee)
	}

	return vm.BlockContext{
		CanTransfer: state.CanTransfer,
		Transfer:    state.Transfer,
		GetHash:     sp.getHashFn(ctx),
		Coinbase:    sp.block.StargazerHeader.Coinbase,
		BlockNumber: new(big.Int).Set(sp.block.StargazerHeader.Number),
		Time:        new(big.Int).SetUint64(sp.block.StargazerHeader.Time),
		Difficulty:  new(big.Int), // not used by stargazer.
		BaseFee:     baseFee,
		GasLimit:    sp.block.StargazerHeader.GasLimit,
		Random:      &common.Hash{}, // TODO: find a source of randomness
	}
}

// `getHashFn` returns a `GetHashFunc` which retrieves header hashes by number.
func (sp *StateProcessor) getHashFn(ctx context.Context) vm.GetHashFunc {
	// Cache will initially contain [refHash.parent],
	// Then fill up with [refHash.p, refHash.pp, refHash.ppp, ...]
	var cache []common.Hash

	return func(n uint64) common.Hash {
		// If there's no hash cache yet, make one
		if len(cache) == 0 {
			cache = append(cache, sp.block.StargazerHeader.ParentHash)
		}
		if idx := sp.block.StargazerHeader.Number.Uint64() - n - 1; idx < uint64(len(cache)) {
			return cache[idx]
		}
		// No luck in the cache, but we can start iterating from the last element we already know
		var lastKnownHash common.Hash
		lastKnownNumber := sp.block.StargazerHeader.Number.Uint64() - uint64(len(cache))
		for {
			header := sp.bp.GetStargazerHeaderAtHeight(ctx, lastKnownNumber)
			if header == nil {
				break
			}
			cache = append(cache, header.ParentHash)
			lastKnownHash = header.ParentHash
			lastKnownNumber = header.Number.Uint64() - 1
			if n == lastKnownNumber {
				return lastKnownHash
			}
		}
		return common.Hash{}
	}
}
