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

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/registry"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `manager` retrieves and runs precompile containers with an ephemeral context.
type manager struct {
	// `Registry` allows the `Controller` to search for a precompile container at an address.
	libtypes.Registry[common.Address, vm.PrecompileContainer]
	// `ctx` is the ephemeral native context, updated on every state transition.
	ctx context.Context
	// `runner` will run the precompile in a custom precompile environment for a given context.
	runner Runner
	// `ldb` is a reference to the StateDB used to add Eth logs from the precompile's execution.
	ldb LogsDB
}

// `NewManager` creates and returns a `Controller` with a native precompile runner and logs DB.
func NewManager(runner Runner, ldb LogsDB) vm.PrecompileManager {
	return &manager{
		Registry: registry.NewMap[common.Address, vm.PrecompileContainer](),
		runner:   runner,
		ldb:      ldb,
	}
}

// `Reset` sets the precompile's native environment context.
//
// `Reset` implements `vm.PrecompileController`.
func (m *manager) Reset(ctx context.Context) {
	m.ctx = ctx
}

// `Run` runs the precompile container using its runner and its ephemeral context.
//
// `Run` implements `vm.PrecompileController`.
func (m *manager) Run(
	pc vm.PrecompileContainer, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	return m.runner.Run(m.ctx, m.ldb, pc, input, caller, value, suppliedGas, readonly)
}
