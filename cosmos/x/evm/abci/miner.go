package abci

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
)

// Mempool defines a mempool interface that can be used to query transactions
type Miner struct {
	txVerifier baseapp.ProposalTxVerifier
	mempool    Mempool
	logger     log.Logger
}

// NewMiner returns a new instance of a miner.
func NewMiner(mempool Mempool, txVerifier baseapp.ProposalTxVerifier, logger log.Logger) *Miner {
	return &Miner{
		mempool: mempool,
		logger:  logger,
	}
}

// func loop() {
// 	for {

// 	}
// }
