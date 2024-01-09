package abci

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// ValidatorCommands is a struct that provides the abci functions required
// for validators to submit cosmos-sdk transactions to the chain.
type ValidatorCommands struct {
	txDecoder      baseapp.ProposalTxVerifier
	valTxSelector  baseapp.TxSelector
	allowedValMsgs map[string]sdk.Msg
}

// processValidatorMsgs processes the validator messages.
func (m *ValidatorCommands) processValidatorMsgs(
	ctx sdk.Context, maxTxBytes int64, ethGasUsed uint64, txs [][]byte,
) ([][]byte, error) { //nolint:unparam // should be handled better.
	var maxBlockGas uint64
	if b := ctx.ConsensusParams().Block; b != nil {
		maxBlockGas = uint64(b.MaxGas)
	}

	blockGasRemaining := maxBlockGas - ethGasUsed

	for _, txBz := range txs {
		tx, err := m.txDecoder.TxDecode(txBz)
		if err != nil {
			continue
		}

		includeTx := true
		for _, msg := range tx.GetMsgs() {
			if _, ok := m.allowedValMsgs[proto.MessageName(msg)]; !ok {
				includeTx = false
				break
			}
		}

		if includeTx {
			stop := m.valTxSelector.SelectTxForProposal(
				ctx, uint64(maxTxBytes), blockGasRemaining, tx, txBz,
			)
			if stop {
				break
			}
		}
	}
	return m.valTxSelector.SelectedTxs(ctx), nil
}
