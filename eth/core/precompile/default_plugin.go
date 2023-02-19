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

// `defaultPlugin` is the default precompile plugin, should any chain running Stargazer EVM not implement
// their own precompile plugin.
type defaultPlugin struct {
	libtypes.Registry[common.Address, vm.PrecompileContainer]
}

// `NewDefaultPlugin` returns a new instance of the default precompile plugin.
func NewDefaultPlugin() Plugin {
	return &defaultPlugin{
		Registry: registry.NewMap[common.Address, vm.PrecompileContainer](),
	}
}

// `Reset` implements `core.PrecompilePlugin`.
func (dp *defaultPlugin) Reset(ctx context.Context) {
	// no-op
}

// `Run` supports executing stateless precompiles with the background context.
//
// `Run` implements `core.PrecompilePlugin`.
func (dp *defaultPlugin) Run(
	sdb vm.GethStateDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	if pc.RequiredGas(input) > suppliedGas {
		return nil, 0, vm.ErrOutOfGas
	}

	ret, err := pc.Run(context.Background(), input, caller, value, readonly)

	return ret, suppliedGas - pc.RequiredGas(input), err
}
