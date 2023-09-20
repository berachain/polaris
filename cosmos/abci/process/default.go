package proposal

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	// ProposalTxVerifier defines the interface that is implemented by BaseApp,
	// that any custom ABCI PrepareProposal and ProcessProposal handler can use
	// to verify a transaction.
	ProposalTxVerifier interface {
		PrepareProposalVerifyTx(tx sdk.Tx) ([]byte, error)
		ProcessProposalVerifyTx(txBz []byte) (sdk.Tx, error)
	}

	// GasTx defines the contract that a transaction with a gas limit must implement.
	GasTx interface {
		GetGas() uint64
	}
)

type Handler struct {
	txVerifier ProposalTxVerifier
}

func NewHandler(txVerifier ProposalTxVerifier) Handler {
	return Handler{
		txVerifier: txVerifier,
	}
}

func (h *Handler) ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	var totalTxGas uint64

	var maxBlockGas int64
	if b := ctx.ConsensusParams().Block; b != nil {
		maxBlockGas = b.MaxGas
	}

	for _, txBytes := range req.Txs {
		tx, err := h.txVerifier.ProcessProposalVerifyTx(txBytes)
		if err != nil {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
		}

		if maxBlockGas > 0 {
			gasTx, ok := tx.(GasTx)
			if ok {
				totalTxGas += gasTx.GetGas()
			}

			if totalTxGas > uint64(maxBlockGas) {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
		}
	}

	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}
