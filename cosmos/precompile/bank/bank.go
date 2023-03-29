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

package bank

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// Contract is the precompile contract for the bank module.
type Contract struct {
	precompile.BaseContract

	msgServer banktypes.MsgServer
}

// NewPrecompileContract returns a new instance of the bank precompile contract.
func NewPrecompileContract(bk bankkeeper.Keeper) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			generated.BankModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
		),
		msgServer: bankkeeper.NewMsgServerImpl(bk),
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		// {
		// 	AbiSig:  "getDelegation(address,address)",
		// 	Execute: c.GetDelegationAddrInput,
		// },
	}
}

// GetBalance implements `GetBalance(address)` method.
func (c *Contract) GetBalance(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	val, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.getDelegationHelper(
		ctx, cosmlib.AddressToAccAddress(del), cosmlib.AddressToValAddress(val),
	)
}