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

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the staking module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer
}

// NewContract is the constructor of the staking contract.
func NewPrecompileContract(sk *stakingkeeper.Keeper) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.StakingModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(stakingtypes.ModuleName)),
		),
		msgServer: stakingkeeper.NewMsgServerImpl(sk),
		querier:   stakingkeeper.Querier{Keeper: sk},
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "getDelegation(address,address)",
			Execute: c.GetDelegationAddrInput,
		},
		{
			AbiSig:  "getUnbondingDelegation(address,address)",
			Execute: c.GetUnbondingDelegationAddrInput,
		},
		{
			AbiSig:  "getRedelegations(address,address,address)",
			Execute: c.GetRedelegationsAddrInput,
		},
		{
			AbiSig:  "delegate(address,uint256)",
			Execute: c.DelegateAddrInput,
		},
		{
			AbiSig:  "undelegate(address,uint256)",
			Execute: c.UndelegateAddrInput,
		},
		{
			AbiSig:  "beginRedelegate(address,address,uint256)",
			Execute: c.BeginRedelegateAddrInput,
		},
		{
			AbiSig:  "cancelUnbondingDelegation(address,uint256,int64)",
			Execute: c.CancelUnbondingDelegationAddrInput,
		},
		{
			AbiSig:  "getActiveValidators()",
			Execute: c.GetActiveValidators,
		},
		{
			AbiSig:  "getValidators()",
			Execute: c.GetValidators,
		},
		{
			AbiSig:  "getValidator(address)",
			Execute: c.GetValidatorAddrInput,
		},
		{
			AbiSig:  "getDelegatorValidators(address)",
			Execute: c.GetDelegatorValidatorsAddrInput,
		},
	}
}

// GetDelegationAddrInput implements `getDelegation(address)` method.
func (c *Contract) GetDelegationAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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

// GetUnbondingDelegationAddrInput implements the `getUnbondingDelegation(address)` method.
func (c *Contract) GetUnbondingDelegationAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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

	return c.getUnbondingDelegationHelper(
		ctx, cosmlib.AddressToAccAddress(del), cosmlib.AddressToValAddress(val),
	)
}

// GetRedelegationsAddrInput implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegationsAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	srcVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	dstVal, ok := utils.GetAs[common.Address](args[2])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.getRedelegationsHelper(
		ctx,
		cosmlib.AddressToAccAddress(del),
		cosmlib.AddressToValAddress(srcVal),
		cosmlib.AddressToValAddress(dstVal),
	)
}

// DelegateAddrInput implements the `delegate(address,uint256)` method.
func (c *Contract) DelegateAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.delegateHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val))
}

// UndelegateAddrInput implements the `undelegate(address,uint256)` method.
func (c *Contract) UndelegateAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.undelegateHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val))
}

// BeginRedelegateAddrInput implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegateAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	dstVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.beginRedelegateHelper(
		ctx,
		caller,
		amount,
		cosmlib.AddressToValAddress(srcVal),
		cosmlib.AddressToValAddress(dstVal),
	)
}

// CancelRedelegateAddrInput implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegationAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, precompile.ErrInvalidInt64
	}

	return c.cancelUnbondingDelegationHelper(ctx, caller, amount, cosmlib.AddressToValAddress(val), creationHeight)
}

// GetActiveValidators implements the `getActiveValidators()` method.
func (c *Contract) GetActiveValidators(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return c.activeValidatorsHelper(ctx)
}

// GetValidators implements the `getValidators()` method.
func (c *Contract) GetValidators(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return c.validatorsHelper(ctx)
}

// GetValidators implements the `getValidator(address)` method.
func (c *Contract) GetValidatorAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.validatorHelper(ctx, sdk.ValAddress(val[:]).String())
}

// GetDelegatorValidatorsAddrInput implements the `getDelegatorValidators(address)` method.
func (c *Contract) GetDelegatorValidatorsAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	del, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	return c.delegatorValidatorsHelper(ctx, cosmlib.Bech32FromEthAddress(del))
}
