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

//go:generate abigen --pkg staking --abi ./out/Staking.sol/IStakingModule.abi.json --bin ./out/Staking.sol/IStakingModule.bin --out ./bindings/cosmos/precompile/staking/i_staking_module.abigen.go --type StakingModule
//go:generate abigen --pkg bank --abi ./out/Bank.sol/IBankModule.abi.json --bin ./out/Bank.sol/IBankModule.bin --out ./bindings/cosmos/precompile/bank/i_bank_module.abigen.go --type BankModule
//go:generate abigen --pkg distribution --abi ./out/Distribution.sol/IDistributionModule.abi.json --bin ./out/Distribution.sol/IDistributionModule.bin --out ./bindings/cosmos/precompile/distribution/i_distribution_module.abigen.go --type DistributionModule --exc "IBankModuleCoin"
//go:generate abigen --pkg governance --abi ./out/Governance.sol/IGovernanceModule.abi.json --bin ./out/Governance.sol/IGovernanceModule.bin --out ./bindings/cosmos/precompile/governance/i_governance_module.abigen.go --type GovernanceModule
//go:generate abigen --pkg lib --abi ./out/CosmosTypes.sol/CosmosTypes.abi.json --bin ./out/CosmosTypes.sol/CosmosTypes.bin --out ./bindings/cosmos/lib/cosmos_types.abigen.go --type CosmosTypes
//go:generate abigen --pkg testing --abi ./out/SolmateERC20.sol/SolmateERC20.abi.json --bin ./out/SolmateERC20.sol/SolmateERC20.bin --out ./bindings/testing/solmate_erc20.abigen.go --type SolmateERC20
//go:generate abigen --pkg testing --abi ./out/MockPrecompileInterface.sol/MockPrecompileInterface.abi.json --out ./bindings/testing/mock_precompile_interface.abigen.go --type MockPrecompile
//go:generate abigen --pkg testing --abi ./out/PrecompileConstructor.sol/PrecompileConstructor.abi.json --bin ./out/PrecompileConstructor.sol/PrecompileConstructor.bin --out ./bindings/testing/precompile_constructor.abigen.go --type PrecompileConstructor
//go:generate abigen --pkg testing --abi ./out/ConsumeGas.sol/ConsumeGas.abi.json --bin ./out/ConsumeGas.sol/ConsumeGas.bin --out ./bindings/testing/consume_gas.abigen.go --type ConsumeGas
//go:generate abigen --pkg testing --abi ./out/LiquidStaking.sol/LiquidStaking.abi.json --bin ./out/LiquidStaking.sol/LiquidStaking.bin --out ./bindings/testing/liquid_staking.abigen.go --type LiquidStaking
//go:generate abigen --pkg testing_governance --abi ./out/GovernanceWrapper.sol/GovernanceWrapper.abi.json --bin ./out/GovernanceWrapper.sol/GovernanceWrapper.bin --out ./bindings/testing/governance/governance_wrapper.abigen.go --type GovernanceWrapper
//go:generate abigen --pkg testing --abi ./out/DistributionWrapper.sol/DistributionWrapper.abi.json --bin ./out/DistributionWrapper.sol/DistributionWrapper.bin --out ./bindings/testing/distribution_testing_helper.abigen.go --type DistributionWrapper
//go:generate abigen --pkg testing --abi ./out/MockMethods.sol/MockMethods.abi.json --out ./bindings/testing/mock_methods.abigen.go --type MockMethods
