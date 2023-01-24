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
	"math/big"

	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

type StateProcessor struct {
	// engine *Engine

	// Contextual Variables (updated once per block)
	// signer types.Signer

	statedb vm.StargazerStateDB

	// the blockHash of the current block being processed
	blockHash   common.Hash
	blockNumber *big.Int

	// st *StateTransitioner
}

func (sp *StateProcessor) Prepare(ctx context.Context) {

}

func (sp *StateProcessor) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	// var err error
	sp.statedb.Prepare(tx.Hash(), 0)

	// Build Receipt
	result := ExecutionResult{}
	receipt := &types.Receipt{Type: tx.Type() /*, PostState: root, CumulativeGasUsed: *usedGas*/}
	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas

	// // If the transaction created a contract, store the creation address in the receipt.
	// if msg.To() == nil {
	// 	receipt.ContractAddress = crypto.CreateAddress(evm.TxContext.Origin, tx.Nonce())
	// }

	// // Set the receipt logs and create the bloom filter.
	// receipt.Logs = sp.statedb.GetLogs(tx.Hash(), sp.blockHash)
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = sp.blockHash
	receipt.BlockNumber = sp.blockNumber
	// receipt.TransactionIndex = uint(sp.statedb.TxIndex())
	return receipt, nil
}

func (sp *StateProcessor) Finalize(ctx context.Context) {

}
