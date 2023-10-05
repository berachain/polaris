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
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/registry"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// defaultPlugin is the default precompile plugin, should any chain running Polaris EVM not
// implement their own precompile plugin. Notably, this plugin can only run the default stateless
// precompiles provided by Go-Ethereum.
type defaultPlugin struct {
	libtypes.Registry[common.Address, vm.PrecompileContainer]
}

// NewDefaultPlugin returns a new instance of the default precompile plugin.
func NewDefaultPlugin() Plugin {
	return &defaultPlugin{
		Registry: registry.NewMap[common.Address, vm.PrecompileContainer](),
	}
}

// Register is a no-op for the default plugin.
func (dp *defaultPlugin) Register(vm.PrecompileContainer) error {
	// no-op
	return nil
}

func (dp *defaultPlugin) Get(_ common.Address, _ *params.Rules) (vm.PrecompiledContract, bool) {
	return nil, false
}

// GetPrecompiles implements core.PrecompilePlugin.
func (dp *defaultPlugin) GetPrecompiles(_ *params.Rules) []Registrable {
	return nil
}

// GetActive implements core.PrecompilePlugin.
func (dp *defaultPlugin) GetActive(_ params.Rules) []common.Address {
	return nil
}

// Run supports executing stateless precompiles with the background context.
//
// Run implements core.PrecompilePlugin.
func (dp *defaultPlugin) Run(
	evm vm.PrecompileEVM, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, _ bool,
) ([]byte, uint64, error) {
	gasCost := pc.RequiredGas(input)
	if gasCost > suppliedGas {
		return nil, 0, vm.ErrOutOfGas
	}

	suppliedGas -= gasCost
	output, err := pc.Run(context.Background(), evm, input, caller, value)

	return output, suppliedGas, err
}

// EnableReentrancy implements core.PrecompilePlugin.
func (dp *defaultPlugin) EnableReentrancy(vm.PrecompileEVM) {}

// DisableReentrancy implements core.PrecompilePlugin.
func (dp *defaultPlugin) DisableReentrancy(vm.PrecompileEVM) {}
