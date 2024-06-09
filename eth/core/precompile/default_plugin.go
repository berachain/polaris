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

package precompile

import (
	"context"
	"math/big"

	"github.com/berachain/polaris/lib/registry"
	libtypes "github.com/berachain/polaris/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// defaultPlugin is the default precompile plugin, should any chain running Polaris EVM not
// implement their own precompile plugin. Notably, this plugin can only run the default stateless
// precompiles provided by Go-Ethereum.
type defaultPlugin struct {
	libtypes.Registry[common.Address, vm.PrecompiledContract]
}

// NewDefaultPlugin returns a new instance of the default precompile plugin.
func NewDefaultPlugin() Plugin {
	return &defaultPlugin{
		Registry: registry.NewMap[common.Address, vm.PrecompiledContract](),
	}
}

// Register is a no-op for the default plugin.
func (dp *defaultPlugin) Register(vm.PrecompiledContract) error {
	// no-op
	return nil
}

func (dp *defaultPlugin) Get(_ common.Address, _ *params.Rules) (vm.PrecompiledContract, bool) {
	return nil, false
}

// GetActive implements core.PrecompilePlugin.
func (dp *defaultPlugin) GetActive(_ params.Rules) []common.Address {
	return nil
}

// Run supports executing stateless precompiles with the background context.
//
// Run implements core.PrecompilePlugin.
func (dp *defaultPlugin) Run(
	evm vm.PrecompileEVM, pc vm.PrecompiledContract, input []byte,
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
