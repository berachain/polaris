package abci

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	prepare "pkg.berachain.dev/polaris/cosmos/abci/prepare"
	process "pkg.berachain.dev/polaris/cosmos/abci/process"
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

	// DefaultProposalHandler defines the default ABCI PrepareProposal and
	// ProcessProposal handlers.
	DefaultProposalHandler struct {
		proposer  prepare.Handler
		processor process.Handler
	}
)

func NewDefaultProposalHandler(mp mempool.Mempool, txVerifier ProposalTxVerifier) DefaultProposalHandler {
	_, isNoOp := mp.(mempool.NoOpMempool)
	if mp == nil || isNoOp {
		panic("mempool must be set and cannot be a NoOpMempool")
	}
	return DefaultProposalHandler{
		proposer:  prepare.NewHandler(mp, txVerifier),
		processor: process.NewHandler(txVerifier),
	}
}

func (h DefaultProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return h.proposer.PrepareProposal
}

func (h DefaultProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return h.processor.ProcessProposal
}
