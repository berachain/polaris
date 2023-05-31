package abci

import (
	"errors"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
)

type Mempool = sdkmempool.Mempool

type ProposalHandler struct {
	miner  *Miner
	logger log.Logger
}

func NewProposalHandler(txVerifier baseapp.ProposalTxVerifier, mempool Mempool, logger log.Logger) *ProposalHandler {
	return &ProposalHandler{
		miner:  NewMiner(mempool, txVerifier, logger),
		logger: logger,
	}
}

func (h *ProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		var (
			selectedTxs  [][]byte
			totalTxBytes int64
		)

		iterator := h.miner.mempool.Select(ctx, req.Txs)

		for iterator != nil {
			memTx := iterator.Tx()

			// NOTE: Since transaction verification was already executed in CheckTx,
			// which calls mempool.Insert, in theory everything in the pool should be
			// valid. But some mempool implementations may insert invalid txs, so we
			// check again.
			bz, err := h.miner.txVerifier.PrepareProposalVerifyTx(memTx)
			if err != nil {
				err := h.miner.mempool.Remove(memTx)
				if err != nil && !errors.Is(err, sdkmempool.ErrTxNotFound) {
					panic(err)
				}
			} else {
				txSize := int64(len(bz))
				if totalTxBytes += txSize; totalTxBytes <= req.MaxTxBytes {
					selectedTxs = append(selectedTxs, bz)
				} else {
					// We've reached capacity per req.MaxTxBytes so we cannot select any
					// more transactions.
					break
				}
			}

			iterator = iterator.Next()
		}

		return &abci.ResponsePrepareProposal{Txs: selectedTxs}, nil
	}
}

func (h *ProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		for _, txBytes := range req.Txs {
			_, err := h.miner.txVerifier.ProcessProposalVerifyTx(txBytes)
			if err != nil {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}
