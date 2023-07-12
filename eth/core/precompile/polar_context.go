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
