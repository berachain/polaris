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
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the distribution module.
type Contract struct {
	precompile.BaseContract

	msgServer distributiontypes.MsgServer
	querier   distributiontypes.QueryServer
}

// `NewPrecompileContract` returns a new instance of the distribution module precompile contract.
func NewPrecompileContract(
	m distributiontypes.MsgServer, q distributiontypes.QueryServer,
) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: precompile.NewBaseContract(
			generated.DistributionModuleMetaData.ABI,
			// 0x93354845030274cD4bf1686Abd60AB28EC52e1a7
			// cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(distributiontypes.ModuleName)),
			common.BytesToAddress([]byte{0x69}), // TODO: choose a better address.
		),
		msgServer: m,
		querier:   q,
	}
}

// `CustomValueDecoders` overrides the `coreprecompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		distributiontypes.AttributeKeyWithdrawAddress: log.ConvertBech32AccAddressToEth,
		sdk.AttributeKeyAmount:                        log.ConvertInt64,
		distributiontypes.AttributeKeyValidator:       log.ConvertBech32ValAddressToEth,
	}
}

// `PrecompileMethods` implements the `coreprecompile.StatefulImpl` interface.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
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
		{
			AbiSig:  "getWithdrawEnabled()",
			Execute: c.GetWithdrawAddrEnabled,
		},
	}
}

// `SetWithdrawAddress` is the precompile contract method for the `setWithdrawAddress(address)` method.
func (c *Contract) SetWithdrawAddress(
	ctx context.Context,
	_ ethprecompile.EVM,
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
	_ ethprecompile.EVM,
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
	_ ethprecompile.EVM,
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

// `WithdrawDelegatorRewardBech32` is the precompile contract method for the `withdrawDelegatorReward(string,string)`.
func (c *Contract) WithdrawDelegatorRewardBech32(
	ctx context.Context,
	_ ethprecompile.EVM,
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

// `GetWithdrawAddrEnabled` is the precompile contract method for the `getWithdrawEnabled()` method.
func (c *Contract) GetWithdrawAddrEnabled(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	return c.getWithdrawAddrEnabled(ctx)
}
