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

package erc20

import (
	"math/big"

	"github.com/holiman/uint256"

	sdk "github.com/cosmos/cosmos-sdk/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// salt is used to generate unique contract addresses for each ERC20 token.
var salt uint64

// deployPolarisERC20Contract deploys a new ERC20 token contract by calling back into the EVM.
func (c *Contract) deployPolarisERC20Contract(
	ctx sdk.Context,
	evm ethprecompile.EVM,
	deployer common.Address,
	name string,
	endowment *big.Int,
) (common.Address, error) {
	plugin := c.GetPlugin()
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

	// deploy new ERC20 token contract
	code := common.FromHex(generated.PolarisERC20MetaData.Bin)
	args, err := c.polarisERC20ABI.Pack("", name, name)
	if err != nil {
		return common.Address{}, err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	_, contractAddr, gasRemaining, err := evm.Create2(
		vm.AccountRef(deployer), append(code, args...), suppliedGas, endowment, uint256.NewInt(salt),
	)
	salt++

	// consume gas used by EVM during ERC20 deployment
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, "Polaris ERC20 deployment")
	return contractAddr, err
}

// callPolarisERC20Mint mints ERC20 tokens by calling back into the EVM.
func (c *Contract) callPolarisERC20Mint(
	ctx sdk.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	token common.Address,
	to common.Address,
	amount *big.Int,
) error {
	plugin := c.GetPlugin()
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

	// call ERC20 contract to mint
	input, err := c.polarisERC20ABI.Pack("mint", to, amount)
	if err != nil {
		return err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	_, gasRemaining, err := evm.Call(
		vm.AccountRef(caller), token, input, suppliedGas, big.NewInt(0),
	)

	// consume gas used by EVM during ERC20 mint
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, "Polaris ERC20 mint")
	return err
}

// callERC20TransferFrom transfers ERC20 tokens by calling back into the EVM.
func (c *Contract) callERC20TransferFrom(
	ctx sdk.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	token common.Address,
	from common.Address,
	to common.Address,
	amount *big.Int,
) error {
	plugin := c.GetPlugin()
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

	// call ERC20 contract to transferFrom
	input, err := c.polarisERC20ABI.Pack("transferFrom", from, to, amount)
	if err != nil {
		return err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	_, gasRemaining, err := evm.Call(
		vm.AccountRef(caller), token, input, suppliedGas, big.NewInt(0),
	)

	// consume gas used by EVM during ERC20 transferFrom
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, "ERC20 transferFrom")
	return err
}

func (c *Contract) callERC20Approve(
	ctx sdk.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	token common.Address,
	spender common.Address,
	amount *big.Int,
) error {
	plugin := c.GetPlugin()
	plugin.EnableReentrancy(ctx)
	defer plugin.DisableReentrancy(ctx)

	// call ERC20 contract to approve
	input, err := c.polarisERC20ABI.Pack("approve", spender, amount)
	if err != nil {
		return err
	}
	suppliedGas := ctx.GasMeter().GasRemaining()
	_, gasRemaining, err := evm.Call(
		vm.AccountRef(caller), token, input, suppliedGas, big.NewInt(0),
	)

	// consume gas used by EVM during ERC20 approve
	ctx.GasMeter().ConsumeGas(suppliedGas-gasRemaining, "ERC20 approve")
	return err
}

// ConvertCommonHexAddress is a value decoder.
var _ ethprecompile.ValueDecoder = ConvertCommonHexAddress

// ConvertCommonHexAddress converts a common hex address attribute to a common.Address and returns
// it as type any.
func ConvertCommonHexAddress(attributeValue string) (any, error) {
	return common.HexToAddress(attributeValue), nil
}
