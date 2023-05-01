package pob

import (
	pobabci "github.com/skip-mev/pob/abci"
	"github.com/skip-mev/pob/x/builder/ante"
)

// Mempool is the interface that the mempool must implement to be used by the builder.
type Mempool interface {
	pobabci.Mempool
	ante.Mempool
}
