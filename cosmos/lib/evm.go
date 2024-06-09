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

package lib

import (
	"math/big"

	"github.com/berachain/polaris/cosmos/x/evm/plugins/precompile"
	"github.com/berachain/polaris/eth/accounts/abi"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TODO: Add these functions to the EVM object itself to allow enforcing calls into
// EVM automatically (i.e. precompile cannot bypass these calls to enter the EVM via call/create
// when in read-only mode). Use gas pool to consume gas rather than Cosmos gas meter.

// DeployOnEVMFromPrecompile deploys an EVM contract from a precompile contract.
func DeployOnEVMFromPrecompile(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm vm.PrecompileEVM,
	deployer common.Address,
	contract abi.ABI,
	endowment *big.Int,
	contractCode string, // hex-encoded string
	constructorArgs ...any,
) (common.Address, []byte, error) {
	if utils.MustGetAs[precompile.MultiStore](ctx.MultiStore()).IsReadOnly() {
		return common.Address{}, nil, vm.ErrWriteProtection
	}

	plugin.EnableReentrancy(evm)
	defer plugin.DisableReentrancy(evm)

	code := common.FromHex(contractCode)
	args, err := contract.Pack("", constructorArgs...)
	if err != nil {
		return common.Address{}, nil, err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	ret, contractAddr, gasRemaining, err := evm.Create(
		vm.AccountRef(deployer), append(code, args...), suppliedGas, endowment,
	)

	// consume gas used by EVM during contract creation
	ctx.GasMeter().ConsumeGas(
		suppliedGas-gasRemaining, "EVM contract creation "+contractAddr.Hex())
	return contractAddr, ret, err
}

// CallEVMFromPrecompile calls into the EVM from a precompile contract.
func CallEVMFromPrecompile(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm vm.PrecompileEVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	value *big.Int,
	methodName string,
	args ...any,
) ([]byte, error) {
	if utils.MustGetAs[precompile.MultiStore](ctx.MultiStore()).IsReadOnly() {
		return nil, vm.ErrWriteProtection
	}

	plugin.EnableReentrancy(evm)
	defer plugin.DisableReentrancy(evm)

	input, err := contract.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	ret, gasRemaining, err := evm.Call(
		vm.AccountRef(caller), address, input, suppliedGas, value,
	)

	// consume gas used by EVM during contract call
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, methodName)
	return ret, err
}

// CallEVMFromPrecompileUnpackArgs calls into the EVM from a precompile contract and returns the
// unpacked result.
func CallEVMFromPrecompileUnpackArgs(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm vm.PrecompileEVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	value *big.Int,
	methodName string,
	args ...any,
) ([]any, error) {
	ret, err := CallEVMFromPrecompile(
		ctx, plugin, evm, caller, address, contract, value, methodName, args...)
	if err != nil {
		return nil, err
	}

	return contract.Unpack(methodName, ret)
}

// StaticCallEVMFromPrecompileUnpackArgs calls into the EVM from a precompile contract for readonly
// calls.
func StaticCallEVMFromPrecompileUnpackArgs(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm vm.PrecompileEVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	methodName string,
	args ...any,
) ([]any, error) {
	plugin.EnableReentrancy(evm)
	defer plugin.DisableReentrancy(evm)

	input, err := contract.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	ret, gasRemaining, err := evm.StaticCall(
		vm.AccountRef(caller), address, input, suppliedGas,
	)
	if err != nil {
		return nil, err
	}

	// consume gas used by EVM during contract call
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, methodName)
	return contract.Unpack(methodName, ret)
}
