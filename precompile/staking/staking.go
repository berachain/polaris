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

package staking

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"pkg.berachain.dev/stargazer/eth/accounts/abi"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/precompile/contracts/solidity/generated"
	evmutils "pkg.berachain.dev/stargazer/x/evm/utils"
)

// `Contract` is the precompile contract for the staking module.
type Contract struct {
	contractAbi *abi.ABI

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer
}

// `NewContract` is the constructor of the staking contract.
func NewPrecompileContract(sk *stakingkeeper.Keeper) precompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		contractAbi: &contractAbi,
		msgServer:   stakingkeeper.NewMsgServerImpl(sk),
		querier:     stakingkeeper.Querier{Keeper: sk},
	}
}

// `RegistryKey` implements StatefulImpl.
func (c *Contract) RegistryKey() common.Address {
	// Contract Address: 0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF
	return evmutils.AccAddressToEthAddress(authtypes.NewModuleAddress(stakingtypes.ModuleName))
}

// `ABIMethods` implements StatefulImpl.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `PrecompileMethods` implements StatefulImpl.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		{
			AbiSig:  "getDelegation(address,address)",
			Execute: c.GetDelegationAddrInput,
		},
		{
			AbiSig:  "getDelegation(string,string)",
			Execute: c.GetDelegationStringInput,
		},
		{
			AbiSig:  "getUnbondingDelegation(address,address)",
			Execute: c.GetUnbondingDelegationAddrInput,
		},
		{
			AbiSig:  "getUnbondingDelegation(string,string)",
			Execute: c.GetUnbondingDelegationStringInput,
		},
		{
			AbiSig:  "getRedelegations(address,address,address)",
			Execute: c.GetRedelegationsAddrInput,
		},
		{
			AbiSig:  "getRedelegations(string,string,string)",
			Execute: c.GetRedelegationsStringInput,
		},
		{
			AbiSig:  "delegate(address,uint256)",
			Execute: c.DelegateAddrInput,
		},
		{
			AbiSig:  "delegate(string,uint256)",
			Execute: c.DelegateStringInput,
		},
		{
			AbiSig:  "undelegate(address,uint256)",
			Execute: c.UndelegateAddrInput,
		},
		{
			AbiSig:  "undelegate(string,uint256)",
			Execute: c.UndelegateStringInput,
		},
		{
			AbiSig:  "beginRedelegate(address,address,uint256)",
			Execute: c.BeginRedelegateAddrInput,
		},
		{
			AbiSig:  "beginRedelegate(string,string,uint256)",
			Execute: c.BeginRedelegateStringInput,
		},
		{
			AbiSig:  "cancelUnbondingDelegation(address,uint256,int64)",
			Execute: c.CancelUnbondingDelegationAddrInput,
		},
		{
			AbiSig:  "cancelUnbondingDelegation(string,uint256,int64)",
			Execute: c.CancelUnbondingDelegationStringInput,
		},
		{
			AbiSig:  "getActiveValidators()",
			Execute: c.GetActiveValidators,
		},
	}
}

// `ABIEvents` implements StatefulImpl.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulImpl.
func (c *Contract) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}

// `GetDelegationAddrInput` implements `getDelegation(address)` method.
func (c *Contract) GetDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidDelegatorAddr
	}
	val, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.getDelegationHelper(
		ctx, evmutils.AddressToAccAddress(del), evmutils.AddressToValAddress(val),
	)
}

// `GetDelegationStringInput` implements `getDelegation(string)` method.
func (c *Contract) GetDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	bech32ValAddr, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}
	val, err := sdk.ValAddressFromBech32(bech32ValAddr)
	if err != nil {
		return nil, err
	}

	return c.getDelegationHelper(ctx, del, val)
}

// `GetUnbondingDelegationAddrInput` implements the `getUnbondingDelegation(address)` method.
func (c *Contract) GetUnbondingDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidDelegatorAddr
	}
	val, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.getUnbondingDelegationHelper(
		ctx, evmutils.AddressToAccAddress(del), evmutils.AddressToValAddress(val),
	)
}

// `GetUnbondingDelegationStringInput` implements the `getUnbondingDelegation(string)` method.
func (c *Contract) GetUnbondingDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	bech32ValAddr, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}
	val, err := sdk.ValAddressFromBech32(bech32ValAddr)
	if err != nil {
		return nil, err
	}

	return c.getUnbondingDelegationHelper(ctx, del, val)
}

// `GetRedelegationsAddrInput` implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegationsAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidDelegatorAddr
	}
	srcVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	dstVal, ok := utils.GetAs[common.Address](args[2])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.getRedelegationsHelper(
		ctx,
		evmutils.AddressToAccAddress(del),
		evmutils.AddressToValAddress(srcVal),
		evmutils.AddressToValAddress(dstVal),
	)
}

// `GetRedelegationsStringInput` implements the `getRedelegations(string,string)` method.
func (c *Contract) GetRedelegationsStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32DelAddr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	srcVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, ErrInvalidString
	}
	del, err := sdk.AccAddressFromBech32(bech32DelAddr)
	if err != nil {
		return nil, err
	}
	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return c.getRedelegationsHelper(ctx, del, src, dst)
}

// `DelegateAddrInput` implements the `delegate(address,uint256)` method.
func (c *Contract) DelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.delegateHelper(ctx, caller, amount, evmutils.AddressToValAddress(val))
}

// `DelegateStringInput` implements the `delegate(string,uint256)` method.
func (c *Contract) DelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.delegateHelper(ctx, caller, amount, val)
}

// `UndelegateAddrInput` implements the `undelegate(address,uint256)` method.
func (c *Contract) UndelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.undelegateHelper(ctx, caller, amount, evmutils.AddressToValAddress(val))
}

// `UndelegateStringInput` implements the `undelegate(string,uint256)` method.
func (c *Contract) UndelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.undelegateHelper(ctx, caller, amount, val)
}

// `BeginRedelegateAddrInput` implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	dstVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.beginRedelegateHelper(
		ctx,
		caller,
		amount,
		evmutils.AddressToValAddress(srcVal),
		evmutils.AddressToValAddress(dstVal),
	)
}

// `BeginRedelegateStringInput` implements the `beginRedelegate(string,string,uint256)` method.
func (c *Contract) BeginRedelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return nil, c.beginRedelegateHelper(ctx, caller, amount, src, dst)
}

// `CancelRedelegateAddrInput` implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, ErrInvalidInt64
	}

	return nil, c.cancelUnbondingDelegationHelper(ctx, caller, amount, evmutils.AddressToValAddress(val), creationHeight)
}

// `CancelRedelegateStringInput` implements the `cancelRedelegate(string,string,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, ErrInvalidInt64
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.cancelUnbondingDelegationHelper(ctx, caller, amount, val, creationHeight)
}

// `GetActiveValidators` implements the `getActiveValidators()` method.
func (c *Contract) GetActiveValidators(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	return c.activeValidatorsHelper(ctx)
}
