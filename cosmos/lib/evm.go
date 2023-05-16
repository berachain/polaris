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

package lib

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// DeployOnEVMFromPrecompile deploys an EVM contract from a precompile contract.
func DeployOnEVMFromPrecompile(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm ethprecompile.EVM,
	deployer common.Address,
	contract abi.ABI,
	endowment *big.Int,
	contractCode string, // hex-encoded string
	constructorArgs ...any,
) (common.Address, []byte, error) {
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

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
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, "EVM contract creation "+contractAddr.Hex())
	return contractAddr, ret, err
}

// CallEVMFromPrecompile calls into the EVM from a precompile contract.
func CallEVMFromPrecompile(
	ctx sdk.Context,
	plugin ethprecompile.Plugin,
	evm ethprecompile.EVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	value *big.Int,
	methodName string,
	args ...any,
) ([]byte, error) {
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

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
	evm ethprecompile.EVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	value *big.Int,
	methodName string,
	args ...any,
) ([]any, error) {
	ret, err := CallEVMFromPrecompile(ctx, plugin, evm, caller, address, contract, value, methodName, args...)
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
	evm ethprecompile.EVM,
	caller common.Address,
	address common.Address,
	contract abi.ABI,
	methodName string,
	args ...any,
) ([]any, error) {
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

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
