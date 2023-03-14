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

//go:generate abigen --pkg generated --abi ./out/staking.sol/IStakingModule.abi.json --bin ./out/staking.sol/IStakingModule.bin --out ./bindings/cosmos/precompile/i_staking_module.abigen.go --type StakingModule
//go:generate abigen --pkg generated --abi ./out/bank.sol/IBankModule.abi.json --bin ./out/bank.sol/IbankModule.bin --out ./bindings/cosmos/precompile/i_bank_module.abigen.go --type BankModule
//go:generate abigen --pkg generated --abi ./out/address.sol/IAddress.abi.json --bin ./out/address.sol/IAddress.bin --out ./bindings/cosmos/precompile/i_address.abigen.go --type Address
//go:generate abigen --pkg generated --abi ./out/distribution.sol/IDistributionModule.abi.json --bin ./out/distribution.sol/IDistributionModule.bin --out ./bindings/cosmos/precompile/i_distribution_module.abigen.go --type DistributionModule
//go:generate abigen --pkg generated --abi ./out/governance.sol/IGovernanceModule.abi.json --bin ./out/governance.sol/IGovernanceModule.bin --out ./bindings/cosmos/precompile/i_governance_module.abigen.go --type GovernanceModule
