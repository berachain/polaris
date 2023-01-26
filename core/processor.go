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

	"github.com/berachain/stargazer/core/block"
	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
	"github.com/berachain/stargazer/params"
)

type StateProcessor struct {
	// The Host provides the underlying application the EVM is running in
	// as well an underlying consensus engine.
	host Host

	// Contextual Variables (updated once per block)
	// signer types.Signer
	config  *params.EthChainConfig
	vmf     vm.EVMFactory
	evm     vm.StargazerEVM
	statedb vm.StargazerStateDB

	// the blockHash of the current block being processed
	blockHash    common.Hash
	blockContext vm.BlockContext

	// `receipts` are stored in the state processor to be returned to the caller.
	receipts types.Receipts
}

func NewStateProcessor(
	config *params.EthChainConfig,
	host Host,
) *StateProcessor {
	return &StateProcessor{
		config: config,
		host:   host,
		vmf:    *vm.NewEVMFactory(nil),
	}
}

func (sp *StateProcessor) Prepare(ctx context.Context, block *block.Data) {
	// Build block context.
	sp.blockContext = NewEVMBlockContext(block, sp.host.GetBlockHashFunc(ctx))

	// Save the block hash to prevent having to recalculate it later.
	sp.blockHash = sp.blockContext.GetHash(sp.blockContext.BlockNumber.Uint64())

	// Build a new EVM to use for this block.
	sp.evm = sp.vmf.Build(sp.statedb, sp.blockContext, &params.EthChainConfig{}, false)

	// Store direct pointers to structs in the evm in order to save a little computation.
	sp.statedb = sp.evm.StateDB()
}

func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	msg, err := tx.AsMessage(types.MakeSigner(sp.config, sp.blockContext.BlockNumber), sp.blockContext.BaseFee)
	if err != nil {
		return nil, fmt.Errorf("could not apply tx %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Create a new context to be used in the EVM environment.
	txContext := NewEVMTxContext(msg)
	sp.evm.Reset(txContext, sp.statedb)

	// var err error
	sp.statedb.Prepare(tx.Hash(), 0)

	// Apply the state transition.
	result, err := (&StateTransitioner{}).ApplyMessageAndCommit(sp.evm, msg)
	if err != nil {
		return nil, fmt.Errorf("could apply message %d [%v]: %w", 0, tx.Hash().Hex(), err)
	}

	// Build Receipt
	receipt := &types.Receipt{Type: tx.Type() /*, PostState: root, CumulativeGasUsed: *usedGas*/}
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
	receipt.Logs = sp.statedb.GetLogs(tx.Hash(), sp.blockHash)
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = sp.blockHash
	receipt.BlockNumber = sp.blockContext.BlockNumber
	// receipt.TransactionIndex = uint(sp.statedb.TxIndex())

	sp.receipts = append(sp.receipts, receipt)
	return receipt, nil
}

func (sp *StateProcessor) Finalize(ctx context.Context) (types.Receipts, types.Bloom, error) {
	return sp.receipts, types.CreateBloom(sp.receipts), nil
}
