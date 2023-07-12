package precompile

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
)

type PolarContext struct {
	// Registrable is the base precompile implementation.
	Registrable
	ctx    context.Context
	evm    EVM
	caller common.Address
	value  *big.Int
}
