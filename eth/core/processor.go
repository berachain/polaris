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

	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/params"
	"github.com/berachain/stargazer/lib/crypto"
)

type StateProcessor struct {
	// `host` provides the underlying chain the EVM is running on
	host StargazerHostChain

	// Contextual Variables (updated once per block)
	signer  types.Signer
	config  *params.ChainConfig
	vmf     *vm.EVMFactory
	evm     vm.StargazerEVM
	statedb vm.StargazerStateDB

	// `blockHeader` of the current block being processed
	blockHeader *types.StargazerHeader
	// `receipts` of the current block being processed
	receipts types.Receipts
	// all the logs of the block, indexed by tx index
	logs [][]*types.Log
	// `logIndex` is the index of the current log in the current block
	logIndex uint
	// `transactions` of the current block being processed
	transactions types.Transactions
}

// `NewStateProcessor` creates a new state processor.
func NewStateProcessor(
	config *params.ChainConfig,
	statedb vm.StargazerStateDB,
	host StargazerHostChain,
) *StateProcessor {
	return &StateProcessor{
		config:  config,
		host:    host,
		statedb: statedb,
		vmf:     vm.NewEVMFactory(precompile.NewManager(host.GetPrecompilePlugin(), statedb)),
	}
}

// `Prepare` prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(ctx context.Context, height uint64) {
	// Build a block to use throughout the evm.
	// NOTE: sp.blockHeader.Bloom is nil here, but it is set in `Finalize`.
	// sp.blockHeader = sp.host.StargazerHeaderAtHeight(ctx, height)
	sp.receipts = types.Receipts{}
	sp.logIndex = 0
	sp.logs = make([][]*types.Log, 0)
	sp.transactions = types.Transactions{}
	sp.statedb.Reset(ctx)
	sp.signer = types.MakeSigner(sp.config, sp.blockHeader.Number)

	// Build a new EVM to use for this block.
	sp.evm = sp.vmf.Build(
		sp.statedb,
		NewEVMBlockContext(
			ctx,
			sp.blockHeader,
			sp.host,
		),
		sp.config,
		false,
	)
}

// `ProcessTransaction` applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	msg, err := tx.AsMessage(sp.signer, sp.blockHeader.BaseFee)
	if err != nil {
		return nil, fmt.Errorf("could not apply tx %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Create a new context to be used in the EVM environment.
	txContext := NewEVMTxContext(msg)
	sp.statedb.Reset(ctx)
	sp.evm.Reset(txContext, sp.statedb)

	// Apply the state transition.
	gp := sp.host.GetGasPlugin()
	result, err := ApplyMessageAndCommit(sp.evm, gp, msg)
	if err != nil {
		return nil, fmt.Errorf("could apply message %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	receipt := &types.Receipt{
		Type:             tx.Type(),
		PostState:        nil, // TODO: Should we do something with PostState?
		TxHash:           tx.Hash(),
		GasUsed:          result.UsedGas,
		BlockHash:        sp.blockHeader.Hash(),
		BlockNumber:      sp.blockHeader.Number,
		TransactionIndex: uint(len(sp.transactions)),
	}

	// Gas from this transaction was added to the gasPlugin in `ApplyMessageAndCommit`
	// And thus CumulativeGasUsed should include gas from all prior transactions in the
	// block, plus the gas consumed during this one.
	receipt.CumulativeGasUsed = gp.CumulativeGasUsed()

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
		receipt.TxHash, receipt.BlockHash, uint(len(sp.receipts)), sp.logIndex,
	)
	sp.logs = append(sp.logs, receipt.Logs)
	sp.logIndex += uint(len(receipt.Logs))
	receipt.Bloom = types.BytesToBloom(types.LogsBloom(receipt.Logs))

	// Update the block information.
	sp.transactions = append(sp.transactions, tx)
	sp.receipts = append(sp.receipts, receipt)
	return receipt, nil
}

// `Finalize` finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(ctx context.Context, height uint64) (*types.StargazerBlock, error) {
	// Update the block header with information regarding the final state of the block.
	sp.blockHeader.GasUsed = sp.host.GetGasPlugin().CumulativeGasUsed()
	sp.blockHeader.Bloom = types.CreateBloom(sp.receipts)

	// Return a finalized block.
	return types.NewStargazerBlock(
		sp.blockHeader,
		sp.transactions,
	), nil
}
