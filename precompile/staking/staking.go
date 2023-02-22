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
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/types/abi"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/precompile/contracts/solidity/generated"
)

var _ precompile.StatefulPrecompileImpl = (*Contract)(nil)

// `Contract` is the precompile contract for the staking module.
type Contract struct {
	vm.PrecompileContainer

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer

	contractAbi abi.ABI
}

// `NewContract` is the constructor of the staking contract.
func NewContract(sk *stakingkeeper.Keeper) *Contract {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		msgServer:   stakingkeeper.NewMsgServerImpl(sk),
		querier:     stakingkeeper.Querier{Keeper: sk},
		contractAbi: contractAbi,
	}
}

// `ABIMethods` implements StatefulPrecompileImpl.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `PrecompileMethods` implements StatefulPrecompileImpl.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		&precompile.Method{
			AbiSig:  "getDelegation(address)",
			Execute: c.GetDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getDelegation(string)",
			Execute: c.GetDelegationStringInput,
		},
		&precompile.Method{
			AbiSig:  "getUnbondingDelegation(address)",
			Execute: c.GetUnbondingDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getUnbondingDelegation(string)",
			Execute: c.GetUnbondingDelegationStringInput,
		},
		&precompile.Method{
			AbiSig:  "getRedelegations(address,address)",
			Execute: c.GetRedelegationsAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getRedelegations(string,string)",
			Execute: c.GetRedelegationsStringInput,
		},
		&precompile.Method{
			AbiSig:  "delegate(address,uint256)",
			Execute: c.DelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "delegate(string,uint256)",
			Execute: c.DelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "undelegate(address,uint256)",
			Execute: c.UndelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "undelegate(string,uint256)",
			Execute: c.UndelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "beginRedelegate(address,address,uint256)",
			Execute: c.BeginRedelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "beginRedelegate(string,string,uint256)",
			Execute: c.BeginRedelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "cancelUnbondingDelegation(address,uint256,int64)",
			Execute: c.CancelUnbondingDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "cancelUnbondingDelegation(string,uint256,int64)",
			Execute: c.CancelUnbondingDelegationStringInput,
		},
	}
}

// `ABIEvents` implements StatefulPrecompileImpl.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulPrecompileImpl.
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
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.delegationHelper(ctx, caller, sdk.ValAddress(val.Bytes()))
}

// `GetDelegationStringInput` implements `getDelegation(string)` method.
func (c *Contract) GetDelegationStringInput(
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

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.delegationHelper(ctx, caller, val)
}

// `GetUnbondingDelegationAddrInput` implements the `getUnbondingDelegation(address)` method.
func (c *Contract) GetUnbondingDelegationAddrInput(
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

	return c.getUnbondingDelegationHelper(ctx, caller, sdk.ValAddress(val.Bytes()))
}

// `GetUnbondingDelegationStringInput` implements the `getUnbondingDelegation(string)` method.
func (c *Contract) GetUnbondingDelegationStringInput(
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

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.getUnbondingDelegationHelper(ctx, caller, val)
}

// `GetRedelegationsAddrInput` implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegationsAddrInput(
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

	return c.getRedelegationsHelper(ctx, caller, sdk.ValAddress(srcVal.Bytes()), sdk.ValAddress(dstVal.Bytes()))
}

// `GetRedelegationsStringInput` implements the `getRedelegations(string,string)` method.
func (c *Contract) GetRedelegationsStringInput(
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

	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return c.getRedelegationsHelper(ctx, caller, src, dst)
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

	return nil, c.delegateHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()))
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

	return nil, c.undelegateHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()))
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
		sdk.ValAddress(srcVal.Bytes()),
		sdk.ValAddress(dstVal.Bytes()),
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

	return nil, c.cancelUnbondingDelegationHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()), creationHeight)
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
