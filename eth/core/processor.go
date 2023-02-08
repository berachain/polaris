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
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"

	vmmock "github.com/berachain/stargazer/eth/core/vm/mock"
)

type StateProcessor struct {
	// The Host provides the underlying application the EVM is running in
	// as well an underlying consensus engine
	host StargazerHostChain

	// Contextual Variables (updated once per block)
	// signer types.Signer
	config  *params.EthChainConfig
	vmf     *vm.EVMFactory
	evm     vm.StargazerEVM
	statedb vm.StargazerStateDB

	// `blockHeader` of the current block being processed
	blockHeader *types.StargazerHeader
	// `receipts` of the current block being processed
	receipts types.Receipts
	// `transactions` of the current block being processed
	transactions types.Transactions
}

// `NewStateProcessor` creates a new state processor.
func NewStateProcessor(
	config *params.EthChainConfig,
	host StargazerHostChain,
) *StateProcessor {
	return &StateProcessor{
		config: config,
		host:   host,
		vmf:    vm.NewEVMFactory(precompile.NewManager(nil)),
	}
}

// `Prepare` prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(ctx context.Context, height uint64) {
	// Build a block to use throughout the evm.
	// NOTE: sp.blockHeader.Bloom is nil here, but it is set in `Finalize`.
	sp.blockHeader = sp.host.StargazerHeaderAtHeight(ctx, height)
	sp.receipts = types.Receipts{}
	sp.transactions = types.Transactions{}
	// todo: use a real state db
	sp.statedb = vmmock.NewEmptyStateDB()

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

	// Store direct pointers to structs in the evm in order to save a little computation.
	sp.statedb = sp.evm.StateDB()
}

// `ProcessTransaction` applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	msg, err := tx.AsMessage(types.MakeSigner(sp.config, sp.blockHeader.Number), sp.blockHeader.BaseFee)
	if err != nil {
		return nil, fmt.Errorf("could not apply tx %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Create a new context to be used in the EVM environment.
	txContext := NewEVMTxContext(msg)
	sp.evm.Reset(txContext, sp.statedb)

	// Apply the state transition.
	result, err := ApplyMessageAndCommit(sp.evm, msg)
	if err != nil {
		return nil, fmt.Errorf("could apply message %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// TODO: Should we do something with PostState?
	receipt := &types.Receipt{Type: tx.Type(), PostState: common.Hash{}.Bytes()}
	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(txContext.Origin, tx.Nonce())
	}

	// Set the receipt logs and create the bloom filter.
	receipt.BlockHash = sp.blockHeader.Hash()
	receipt.BlockNumber = sp.blockHeader.Number
	receipt.Logs = sp.statedb.BuildLogsAndClear(
		receipt.TxHash, receipt.BlockHash, uint(len(sp.receipts)), uint(0),
	)
	receipt.Bloom = types.BytesToBloom(types.LogsBloom(receipt.Logs))
	receipt.TransactionIndex = uint(len(sp.transactions))

	// Update the block information.
	sp.transactions = append(sp.transactions, tx)
	sp.receipts = append(sp.receipts, receipt)
	return receipt, nil
}

// `Finalize` finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(ctx context.Context, height uint64) (*types.StargazerBlock, error) {
	// Set the header's bloom.
	sp.blockHeader.Bloom = types.CreateBloom(sp.receipts)

	// Return a finalized block.
	return types.NewStargazerBlock(
		sp.blockHeader,
		sp.transactions,
	), nil
}
