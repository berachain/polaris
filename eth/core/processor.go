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

// “.
type StateProcessor struct {
	// `bp` provides block functions from the underlying chain the EVM is running on
	bp BlockPlugin
	// `gp` provides gas functions from the underlying chain the EVM is running on
	gp GasPlugin
	// `signer` is the signer used to verify transaction signatures.
	signer types.Signer
	// `config` is the chain configuration.
	config *params.ChainConfig
	// `vmf` is the EVM factory that is used to create new EVMs.
	vmf *vm.EVMFactory
	// `evm ` is the EVM that is used to process transactions.
	evm vm.StargazerEVM
	// `statedb` is the state database that is used to mange state during transactions.
	statedb vm.StargazerStateDB
	// `block` represents the current block being processed.
	block *types.StargazerBlock
	// `logIndex` is the index of the current log in the current block
	logIndex uint
}

// `NewStateProcessor` creates a new state processor.
func NewStateProcessor(
	config *params.ChainConfig,
	statedb vm.StargazerStateDB,
	host StargazerHostChain,
) *StateProcessor {
	return &StateProcessor{
		bp:      host.GetBlockPlugin(),
		gp:      host.GetGasPlugin(),
		config:  config,
		statedb: statedb,
		vmf:     vm.NewEVMFactory(precompile.NewManager(host.GetPrecompilePlugin(), statedb)),
	}
}

// `Prepare` prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(ctx context.Context, height uint64) {
	// Build a block object so we can track that status of the block as we process it.
	sp.block = &types.StargazerBlock{
		StargazerHeader: sp.bp.GetStargazerHeaderAtHeight(ctx, height),
		Transactions:    make([]*types.Transaction, 0),
		Receipts:        make([]*types.Receipt, 0),
	}
	sp.logIndex = 0

	// We must re-create the signer since we are processing a new block and the block number has increased.
	sp.signer = types.MakeSigner(sp.config, sp.block.Number)

	// Build a new EVM to use for this block.
	sp.evm = sp.vmf.Build(
		sp.statedb,
		NewEVMBlockContext(
			ctx,
			sp.block.StargazerHeader,
			sp.bp,
		),
		sp.config,
		sp.block.BaseFee.Int64() != 0,
	)
}

// `ProcessTransaction` applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	msg, err := tx.AsMessage(sp.signer, sp.block.BaseFee)
	if err != nil {
		return nil, fmt.Errorf("could not apply tx %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Create a new context to be used in the EVM environment. We also must reset the StateDB as well as the EVM.
	txContext := NewEVMTxContext(msg)
	sp.statedb.Reset(ctx)
	sp.evm.Reset(txContext, sp.statedb)

	// Apply the state transition.
	result, err := ApplyMessageAndCommit(sp.evm, sp.gp, msg)
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
		TransactionIndex: uint(len(sp.block.Receipts)),
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
		receipt.TxHash, receipt.BlockHash, uint(len(sp.block.Receipts)), sp.logIndex,
	)
	sp.logIndex += uint(len(receipt.Logs))
	receipt.Bloom = types.BytesToBloom(types.LogsBloom(receipt.Logs))

	// Update the block information.
	sp.block.Transactions = append(sp.block.Transactions, tx)
	sp.block.Receipts = append(sp.block.Receipts, receipt)
	// sp.receipts.Append((*types.ReceiptForStorage)(receipt))
	return receipt, nil
}

// `Finalize` finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(ctx context.Context, height uint64) (*types.StargazerBlock, error) {
	sp.block.SetGasUsed(sp.gp.CumulativeGasUsed())
	sp.block.CreateBloom()
	sp.block.SetReceiptHash()
	return sp.block, nil
}
