package precompile

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
)

type PolarContext interface {
	Ctx() context.Context
	Evm() EVM
	Caller() common.Address
	Value() *big.Int
}

type PolarContextImpl struct {
	ctx    context.Context
	evm    EVM
	caller common.Address
	value  *big.Int
}

func (pCtx *PolarContextImpl) Ctx() context.Context {
	return pCtx.ctx
}

func (pCtx *PolarContextImpl) Evm() EVM {
	return pCtx.evm
}

func (pCtx *PolarContextImpl) Caller() common.Address {
	return pCtx.caller
}

func (pCtx *PolarContextImpl) Value() *big.Int {
	return pCtx.value
}
