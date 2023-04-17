// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package contracts

//go:generate abigen --pkg precompile --abi ./out/staking.sol/IStakingModule.abi.json --bin ./out/staking.sol/IStakingModule.bin --out ./bindings/cosmos/precompile/i_staking_module.abigen.go --type StakingModule
//go:generate abigen --pkg precompile --abi ./out/bank.sol/IBankModule.abi.json --bin ./out/bank.sol/IbankModule.bin --out ./bindings/cosmos/precompile/i_bank_module.abigen.go --type BankModule
//go:generate abigen --pkg precompile --abi ./out/auth.sol/IAuthModule.abi.json --bin ./out/auth.sol/IAuthModule.bin --out ./bindings/cosmos/precompile/i_auth_module.abigen.go --type AuthModule
//go:generate abigen --pkg precompile --abi ./out/distribution.sol/IDistributionModule.abi.json --bin ./out/distribution.sol/IDistributionModule.bin --out ./bindings/cosmos/precompile/i_distribution_module.abigen.go --type DistributionModule
//go:generate abigen --pkg precompile --abi ./out/governance.sol/IGovernanceModule.abi.json --bin ./out/governance.sol/IGovernanceModule.bin --out ./bindings/cosmos/precompile/i_governance_module.abigen.go --type GovernanceModule
//go:generate abigen --pkg precompile --abi ./out/ERC20.sol/IERC20Module.abi.json --bin ./out/ERC20.sol/IERC20Module.bin --out ./bindings/cosmos/precompile/i_erc20_module.abigen.go --type ERC20Module

//go:generate abigen --pkg polaris --abi ./out/PolarisERC20.sol/PolarisERC20.abi.json --bin ./out/PolarisERC20.sol/PolarisERC20.bin --out ./bindings/polaris/polaris_erc20.abigen.go --type PolarisERC20

//go:generate abigen --pkg testing --abi ./out/SolmateERC20.sol/SolmateERC20.abi.json --bin ./out/SolmateERC20.sol/SolmateERC20.bin --out ./bindings/testing/solmate_erc20.abigen.go --type SolmateERC20
//go:generate abigen --pkg testing --abi ./out/MockPrecompileInterface.sol/MockPrecompileInterface.abi.json --out ./bindings/testing/mock_precompile_interface.abigen.go --type MockPrecompile
//go:generate abigen --pkg testing --abi ./out/NonRevertableTx.sol/NonRevertableTx.abi.json --bin ./out/NonRevertableTx.sol/NonRevertableTx.bin --out ./bindings/testing/non_revertable_tx.abigen.go --type NonRevertableTx
//go:generate abigen --pkg testing --abi ./out/ConsumeGas.sol/ConsumeGas.abi.json --bin ./out/ConsumeGas.sol/ConsumeGas.bin --out ./bindings/testing/consume_gas.abigen.go --type ConsumeGas

//go:generate abigen --pkg testing --abi ./out/LiquidStaking.sol/LiquidStaking.abi.json --bin ./out/LiquidStaking.sol/LiquidStaking.bin --out ./bindings/testing/liquid_staking.abigen.go --type LiquidStaking
//go:generate abigen --pkg testing --abi ./out/Swapper.sol/Swapper.abi.json --bin ./out/Swapper.sol/Swapper.bin --out ./bindings/testing/swapper.abigen.go --type Swapper
