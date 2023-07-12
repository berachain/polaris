package precompile

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
)

type PolarContext interface {
	// Ctx returns the context of the current Precompile call.
	Ctx() context.Context
	// Evm returns the EVM instance of the current Precompile call.
	Evm() EVM
	// Caller returns the caller of the current Precompile call.
	Caller() common.Address
	// Value returns the value of the current Precompile call.
	Value() *big.Int
}

type polarCtx struct {
	ctx    context.Context
	evm    EVM
	caller common.Address
	value  *big.Int
}

func NewPolarContext(ctx context.Context, evm EVM, caller common.Address, value *big.Int) PolarContext {
	return &polarCtx{
		ctx:    ctx,
		evm:    evm,
		caller: caller,
		value:  value,
	}
}

func (p *polarCtx) Ctx() context.Context {
	return p.ctx
}

func (p *polarCtx) Evm() EVM {
	return p.evm
}

func (p *polarCtx) Caller() common.Address {
	return p.caller
}

func (p *polarCtx) Value() *big.Int {
	return p.value
}
