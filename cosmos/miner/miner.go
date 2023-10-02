package miner

import (
	"context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/params"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

type Engine interface {
	consensus.Engine
}

var _ baseapp.TxSelector = (*Miner)(nil)

type Miner struct {
	mux *event.TypeMux
	*miner.Miner
	serializer evmtypes.TxSerializer
}

func NewMiner(eth miner.Backend, config *miner.Config, chainConfig *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, isLocalBlock func(header *types.Header) bool) *Miner {
	return &Miner{
		mux:   mux,
		Miner: miner.New(eth, config, chainConfig, mux, engine, isLocalBlock),
	}
}

func (m *Miner) SetSerializer(serializer evmtypes.TxSerializer) {
	m.serializer = serializer
}

func (m *Miner) SelectedTxs() [][]byte {
	payload, err := m.BuildPayload(&miner.BuildPayloadArgs{
		// TODO: properly fill in the rest of the payload.
		Timestamp: m.PendingBlock().Time() + 2,
	})
	if err != nil {
		panic(err)
	}

	// This blocks.
	executionPayload := payload.ResolveFull()

	ethTxBzs := executionPayload.ExecutionPayload.Transactions
	txs := make([][]byte, len(executionPayload.ExecutionPayload.Transactions))

	// encode to sdk.txs and then
	for i, ethTxBz := range ethTxBzs {
		var tx types.Transaction
		if err := tx.UnmarshalBinary(ethTxBz); err != nil {
			return nil
		}
		bz, err := m.serializer.SerializeToBytes(&tx)
		if err != nil {
			panic(err)
		}
		txs[i] = bz
	}
	return txs
}

func (m *Miner) Clear() {
	// no-op
}

func (m *Miner) SelectTxForProposal(_ context.Context, maxTxBytes, maxBlockGas uint64, memTx sdk.Tx, txBz []byte) bool {
	return true
}
