package mempool

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/ethereum/go-ethereum/core/txpool"
)

var _ mempool.Mempool = (*WrappedGethTxPool)(nil)

type WrappedGethTxPool struct {
	txpool.TxPool
}

func (wgtp *WrappedGethTxPool) CountTx() int {
	return 0
}

func (wgtp *WrappedGethTxPool) Insert(_ context.Context, tx sdk.Tx) error {
	return nil
}

func (wgtp *WrappedGethTxPool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

func (wgtp *WrappedGethTxPool) Remove(sdk.Tx) error {
	return nil
}
