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

package distribution

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/precompile/contracts/solidity/generated"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	coreprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// `Contract` is the precompile contract for the distribution module.
type Contract struct {
	precompile.BaseContract

	msgServer distributiontypes.MsgServer
	querier   distributiontypes.QueryServer
}

// `NewPrecompileContract` returns a new instance of the bank precompile contract.
func NewPrecompileContract(dk **distrkeeper.Keeper) coreprecompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.DistributionModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	rk := cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(distributiontypes.ModuleName))
	return &Contract{
		BaseContract: precompile.NewBaseContract(contractAbi, rk),
		msgServer:    distrkeeper.NewMsgServerImpl(**dk),
		querier:      distrkeeper.NewQuerier(**dk),
	}
}

// `PrecompileMethods` implements the `coreprecompile.StatefulImpl` interface.
func (c *Contract) PrecompileMethods() coreprecompile.Methods {
	return coreprecompile.Methods{
		{
			AbiSig:  "setWithdrawAddress(address)",
			Execute: c.SetWithdrawAddress,
		},
		{
			AbiSig:  "setWithdrawAddress(string)",
			Execute: c.SetWithdrawAddressBech32,
		},
		{
			AbiSig:  "withdrawDelegatorReward(address,address)",
			Execute: c.WithdrawDelegatorReward,
		},
		{
			AbiSig:  "withdrawDelegatorReward(string,string)",
			Execute: c.SetWithdrawAddressBech32,
		},
	}
}

// `SetWithdrawAddress` is the precompile contract method for the `setWithdrawAddress(address)` method.
func (c *Contract) SetWithdrawAddress(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	withdrawAddr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.setWithdrawAddressHelper(ctx, sdk.AccAddress(caller.Bytes()), sdk.AccAddress(withdrawAddr.Bytes()))
}

// `SetWithdrawAddressBech32` is the precompile contract method for the `setWithdrawAddress(string)` method.
func (c *Contract) SetWithdrawAddressBech32(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	withdrawAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	addr, err := sdk.AccAddressFromBech32(withdrawAddr)
	if err != nil {
		return nil, err
	}

	return c.setWithdrawAddressHelper(ctx, sdk.AccAddress(caller.Bytes()), addr)
}

// `WithdrawDelegatorReward` is the precompile contract method for the `withdrawDelegatorReward(address,address)`
// method.
func (c *Contract) WithdrawDelegatorReward(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	delegator, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	validator, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.withdrawDelegatorRewardsHelper(ctx, sdk.AccAddress(delegator.Bytes()), sdk.ValAddress(validator.Bytes()))
}

func (c *Contract) WithdrawDelegatorRewardBech32(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	delegator, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	validator, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	delegatorAddr, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		return nil, err
	}
	validatorAddr, err := sdk.ValAddressFromBech32(validator)
	if err != nil {
		return nil, err
	}

	return c.withdrawDelegatorRewardsHelper(ctx, delegatorAddr, validatorAddr)
}
