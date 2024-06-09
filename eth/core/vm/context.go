// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vm

import (
	"context"
	"math/big"
	"time"

	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ContextKey defines a type alias for a stdlib Context key.
type ContextKey string

// PolarContextKey is the key in the context.Context which holds the PolarContext.
const PolarContextKey ContextKey = "polar-context"

// Compile-time assertion that PolarContext implements context.Context.
var _ context.Context = (*PolarContext)(nil)

// Context is the context for a Polaris EVM execution.
type PolarContext struct {
	baseCtx   context.Context
	evm       vm.PrecompileEVM
	msgSender common.Address
	msgValue  *big.Int
}

// NewPolarContext creates a new PolarContext given an EVM call request.
func NewPolarContext(
	baseCtx context.Context,
	evm vm.PrecompileEVM,
	msgSender common.Address,
	msgValue *big.Int,
) *PolarContext {
	return &PolarContext{
		baseCtx:   baseCtx,
		evm:       evm,
		msgSender: msgSender,
		msgValue:  msgValue,
	}
}

// =============================================================================
// vm.PolarContext implementation
// =============================================================================

func (c *PolarContext) Context() context.Context {
	return c.baseCtx
}

func (c *PolarContext) Evm() vm.PrecompileEVM {
	return c.evm
}

func (c *PolarContext) MsgSender() common.Address {
	return c.msgSender
}

func (c *PolarContext) MsgValue() *big.Int {
	return c.msgValue
}

func (c *PolarContext) Block() *vm.BlockContext {
	return c.evm.GetContext()
}

// =============================================================================
// context.Context implementation
// =============================================================================

func (c *PolarContext) Deadline() (time.Time, bool) {
	return c.baseCtx.Deadline()
}

func (c *PolarContext) Done() <-chan struct{} {
	return c.baseCtx.Done()
}

func (c *PolarContext) Err() error {
	return c.baseCtx.Err()
}

func (c *PolarContext) Value(key any) any {
	if key == PolarContextKey {
		return c
	}

	return c.baseCtx.Value(key)
}

// WithValue attaches a value to the context.
func (c *PolarContext) WithValue(key, value any) *PolarContext {
	c.baseCtx = context.WithValue(c.baseCtx, key, value)
	return c
}

// UnwrapPolarContext retrieves a Context from a context.Context instance attached with a
// PolarContext. It panics if a Context was not properly attached.
func UnwrapPolarContext(ctx context.Context) *PolarContext {
	if polarCtx, ok := utils.GetAs[*PolarContext](ctx); ok {
		return polarCtx
	}
	return utils.MustGetAs[*PolarContext](ctx.Value(PolarContextKey))
}
