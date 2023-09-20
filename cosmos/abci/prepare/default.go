package prepare

import (
	"errors"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
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
	mempool    sdkmempool.Mempool
	txVerifier ProposalTxVerifier
}

func NewHandler(mempool sdkmempool.Mempool, txVerifier ProposalTxVerifier) Handler {
	return Handler{
		mempool:    mempool,	
		txVerifier: txVerifier,
	}
}

func (h *Handler) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	var maxBlockGas int64
	if b := ctx.ConsensusParams().Block; b != nil {
		maxBlockGas = b.MaxGas
	}

	var (
		selectedTxs  [][]byte
		totalTxBytes int64
		totalTxGas   uint64
	)

	iterator := h.mempool.Select(ctx, req.Txs)

	for iterator != nil {
		memTx := iterator.Tx()

		// NOTE: Since transaction verification was already executed in CheckTx,
		// which calls mempool.Insert, in theory everything in the pool should be
		// valid. But some mempool implementations may insert invalid txs, so we
		// check again.
		bz, err := h.txVerifier.PrepareProposalVerifyTx(memTx)
		if err != nil {
			err := h.mempool.Remove(memTx)
			if err != nil && !errors.Is(err, sdkmempool.ErrTxNotFound) {
				return nil, err
			}
		} else {
			var txGasLimit uint64
			txSize := int64(len(bz))

			gasTx, ok := memTx.(GasTx)
			if ok {
				txGasLimit = gasTx.GetGas()
			}

			// only add the transaction to the proposal if we have enough capacity
			if (txSize + totalTxBytes) < req.MaxTxBytes {
				// If there is a max block gas limit, add the tx only if the limit has
				// not been met.
				if maxBlockGas > 0 {
					if (txGasLimit + totalTxGas) <= uint64(maxBlockGas) {
						totalTxGas += txGasLimit
						totalTxBytes += txSize
						selectedTxs = append(selectedTxs, bz)
					}
				} else {
					totalTxBytes += txSize
					selectedTxs = append(selectedTxs, bz)
				}
			}

			// Check if we've reached capacity. If so, we cannot select any more
			// transactions.
			if totalTxBytes >= req.MaxTxBytes || (maxBlockGas > 0 && (totalTxGas >= uint64(maxBlockGas))) {
				break
			}
		}

		iterator = iterator.Next()
	}

	return &abci.ResponsePrepareProposal{Txs: selectedTxs}, nil
}
