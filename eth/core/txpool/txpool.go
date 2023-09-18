package txpool

import "pkg.berachain.dev/polaris/eth/core"

type PolarisTxPool interface {
	core.TxPoolPlugin
}

func NewPolarisTxPool(core.TxPoolPlugin) *polarisTxPool {
	return &polarisTxPool{}
}

type polarisTxPool struct {
	core.TxPoolPlugin
}
