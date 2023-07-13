// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package vm

import (
	"context"
	"math/big"
	"time"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/lib/utils"
)

// ContextKey defines a type alias for a stdlib Context key.
type ContextKey string

// PolarContextKey is the key in the context.Context which holds the polar.Context.
const PolarContextKey ContextKey = "polar-context"

// Compile-time assertion that polar.Context implements context.Context.
var _ context.Context = (*Context)(nil)

// Context is the context for a Polaris EVM execution.
type Context struct {
	baseCtx   context.Context
	evm       PrecompileEVM
	msgSender common.Address
	msgValue  *big.Int
}

// NewPolarContext creates a new polar.Context given an EVM call request.
func NewPolarContext(baseCtx context.Context, evm PrecompileEVM, msgSender common.Address, msgValue *big.Int) *Context {
	return &Context{
		baseCtx:   baseCtx,
		evm:       evm,
		msgSender: msgSender,
		msgValue:  msgValue,
	}
}

// =============================================================================
// polar.Context implementation
// =============================================================================

func (c *Context) Context() context.Context {
	return c.baseCtx
}

func (c *Context) Evm() PrecompileEVM {
	return c.evm
}

func (c *Context) MsgSender() common.Address {
	return c.msgSender
}

func (c *Context) MsgValue() *big.Int {
	return c.msgValue
}

func (c *Context) Block() *BlockContext {
	return c.evm.GetContext()
}

// =============================================================================
// context.Context implementation
// =============================================================================

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.baseCtx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.baseCtx.Done()
}

func (c Context) Err() error {
	return c.baseCtx.Err()
}

func (c Context) Value(key any) any {
	if key == PolarContextKey {
		return c
	}

	return c.baseCtx.Value(key)
}

// UnwrapPolarContext retrieves a Context from a context.Context instance attached with a
// PolarContext. It panics if a Context was not properly attached.
func UnwrapPolarContext(ctx context.Context) *Context {
	if polarCtx, ok := utils.GetAs[*Context](ctx); ok {
		return polarCtx
	}
	return utils.MustGetAs[*Context](ctx.Value(PolarContextKey))
}
