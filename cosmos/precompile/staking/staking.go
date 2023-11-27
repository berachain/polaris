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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/core/address"
	sdkmath "cosmossdk.io/math"

	cbindings "github.com/berachain/polaris/contracts/bindings/cosmos/lib"
	generated "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "github.com/berachain/polaris/cosmos/lib"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	pvm "github.com/berachain/polaris/eth/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"
)

type ValidatorStore interface {
	ValidatorAddressCodec() address.Codec
	ValidatorByConsAddr(ctx context.Context, addr sdk.ConsAddress) (stakingtypes.ValidatorI, error)
	ConsensusAddressCodec() address.Codec
	GetValidator(ctx context.Context, addr sdk.ValAddress) (stakingtypes.Validator, error)
	IterateBondedValidatorsByPower(
		ctx context.Context, fn func(index int64, validator stakingtypes.ValidatorI) bool,
	) error
}

// Contract is the precompile contract for the staking module.
type Contract struct {
	ethprecompile.BaseContract

	accAddrCodec address.Codec
	vs           ValidatorStore
	msgServer    stakingtypes.MsgServer
	querier      stakingtypes.QueryServer
}

// NewContract is the constructor of the staking contract.
func NewPrecompileContract(ak cosmlib.CodecProvider, sk *stakingkeeper.Keeper) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.StakingModuleMetaData.ABI,
			common.BytesToAddress(authtypes.NewModuleAddress(stakingtypes.ModuleName)),
		),
		accAddrCodec: ak.AddressCodec(),
		vs:           sk,
		msgServer:    stakingkeeper.NewMsgServerImpl(sk),
		querier:      stakingkeeper.Querier{Keeper: sk},
	}
}

func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		stakingtypes.AttributeKeyDelegator:    c.ConvertAccAddressFromString,
		stakingtypes.AttributeKeyValidator:    c.ConvertValAddressFromString,
		stakingtypes.AttributeKeySrcValidator: c.ConvertValAddressFromString,
		stakingtypes.AttributeKeyDstValidator: c.ConvertValAddressFromString,
	}
}

func (c *Contract) GetValAddressFromConsAddress(
	ctx context.Context,
	consAddress []byte,
) (common.Address, error) {
	val, err := c.vs.ValidatorByConsAddr(ctx, consAddress)
	if err != nil {
		return common.Address{}, err
	}
	return cosmlib.EthAddressFromString(c.vs.ValidatorAddressCodec(), val.GetOperator())
}

// GetBondedValidators implements the `getBondedValidators(PageRequest)` method.
func (c *Contract) GetBondedValidators(
	ctx context.Context,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Status:     stakingtypes.BondStatusBonded,
		Pagination: cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(
		c.vs.ValidatorAddressCodec(), res.GetValidators(),
	)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	pageResponse := cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination)
	return vals, pageResponse, nil
}

// GetBondedValidatorsByPoweer implements the `getBondedValidatorsByPower()` method.
func (c *Contract) GetBondedValidatorsByPower(
	ctx context.Context,
) ([]common.Address, error) {
	var (
		vals []common.Address
		err  error
	)

	iteratorErr := c.vs.IterateBondedValidatorsByPower(
		ctx,
		func(_ int64, validator stakingtypes.ValidatorI) bool {
			var valOperAddr common.Address
			valOperAddr, err = cosmlib.EthAddressFromString(
				c.vs.ValidatorAddressCodec(), validator.GetOperator(),
			)
			if err != nil {
				return true
			}
			vals = append(vals, valOperAddr)
			return false
		},
	)
	if iteratorErr != nil {
		return nil, iteratorErr
	}

	return vals, err
}

// GetValidators implements the `getValidators(PageRequest)` method.
func (c *Contract) GetValidators(
	ctx context.Context,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		Pagination: cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(
		c.vs.ValidatorAddressCodec(), res.GetValidators(),
	)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	pageResponse := cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination)
	return vals, pageResponse, nil
}

// GetValidators implements the `getValidator(address)` method.
func (c *Contract) GetValidator(
	ctx context.Context,
	validatorAddress common.Address,
) (generated.IStakingModuleValidator, error) {
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return generated.IStakingModuleValidator{}, err
	}
	res, err := c.querier.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: valAddr,
	})
	if err != nil {
		return generated.IStakingModuleValidator{}, err
	}

	val, err := cosmlib.SdkValidatorsToStakingValidators(
		c.vs.ValidatorAddressCodec(), []stakingtypes.Validator{res.GetValidator()},
	)
	if err != nil {
		return generated.IStakingModuleValidator{}, err
	}

	// guaranteed not to panic because val is guaranteed to have length 1.
	return val[0], nil
}

// GetDelegatorValidators implements the `getDelegatorValidators(address)` method.
func (c *Contract) GetDelegatorValidators(
	ctx context.Context,
	delegatorAddr common.Address,
	pagination any,
) ([]generated.IStakingModuleValidator, cbindings.CosmosPageResponse, error) {
	delegator, err := cosmlib.StringFromEthAddress(c.accAddrCodec, delegatorAddr)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	res, err := c.querier.DelegatorValidators(ctx, &stakingtypes.QueryDelegatorValidatorsRequest{
		DelegatorAddr: delegator,
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	vals, err := cosmlib.SdkValidatorsToStakingValidators(
		c.vs.ValidatorAddressCodec(), res.GetValidators(),
	)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	return vals, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetValidatorDelegations implements the `getValidatorDelegations(address,PageRequest)` method.
func (c *Contract) GetValidatorDelegations(
	ctx context.Context,
	validatorAddress common.Address,
	pagination any,
) ([]generated.IStakingModuleDelegation, cbindings.CosmosPageResponse, error) {
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	res, err := c.querier.ValidatorDelegations(ctx, &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: valAddr,
		Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleDelegation{}, cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	delegations := make([]generated.IStakingModuleDelegation, 0)
	for _, d := range res.GetDelegationResponses() {
		var delegator common.Address
		delegator, err = cosmlib.EthAddressFromString(c.accAddrCodec, d.Delegation.DelegatorAddress)
		if err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}
		delegations = append(delegations, generated.IStakingModuleDelegation{
			Delegator: delegator,
			Balance:   d.Balance.Amount.BigInt(),
			Shares:    d.Delegation.Shares.BigInt(),
		})
	}

	return delegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetDelegation implements `getDelegation(address)` method.
func (c *Contract) GetDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) (*big.Int, error) {
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return nil, err
	}
	delAddr, err := cosmlib.StringFromEthAddress(c.accAddrCodec, delegatorAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.Delegation(ctx, &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	})
	if status.Code(err) == codes.NotFound {
		// handle the case where the delegation does not exist
		return big.NewInt(0), nil
	} else if err != nil {
		return nil, err
	}

	delegation := res.GetDelegationResponse()
	if delegation == nil {
		return big.NewInt(0), nil
	}

	return delegation.Balance.Amount.BigInt(), nil
}

// GetUnbondingDelegation implements the `getUnbondingDelegation(address,address)` method.
func (c *Contract) GetUnbondingDelegation(
	ctx context.Context,
	delegatorAddress common.Address,
	validatorAddress common.Address,
) ([]generated.IStakingModuleUnbondingDelegationEntry, error) {
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return nil, err
	}
	delAddr, err := cosmlib.StringFromEthAddress(c.accAddrCodec, delegatorAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.UnbondingDelegation(ctx, &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleUnbondingDelegationEntry{}, nil
	} else if err != nil {
		return nil, err
	}

	return cosmlib.SdkUDEToStakingUDE(res.GetUnbond().Entries), nil
}

// GetDelegatorUnbondingDelegations implements the `getDelegatorUnbondingDelegations(address)`
// method.
func (c *Contract) GetDelegatorUnbondingDelegations(
	ctx context.Context,
	delegatorAddress common.Address,
	pagination any,
) ([]generated.IStakingModuleUnbondingDelegation, cbindings.CosmosPageResponse, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.accAddrCodec, delegatorAddress)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	res, err := c.querier.DelegatorUnbondingDelegations(ctx,
		&stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
			DelegatorAddr: delAddr,
			Pagination:    cosmlib.ExtractPageRequestFromInput(pagination),
		})
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleUnbondingDelegation{},
			cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	unbondingDelegations := make([]generated.IStakingModuleUnbondingDelegation, 0)
	for _, u := range res.GetUnbondingResponses() {
		var (
			valAddr   common.Address
			delegator common.Address
		)

		valAddr, err = cosmlib.EthAddressFromString(
			c.vs.ValidatorAddressCodec(), u.ValidatorAddress)
		if err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}
		delegator, err = cosmlib.EthAddressFromString(c.accAddrCodec, u.DelegatorAddress)
		if err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}

		unbondingDelegations = append(unbondingDelegations,
			generated.IStakingModuleUnbondingDelegation{
				DelegatorAddress: delegator,
				ValidatorAddress: valAddr,
				Entries:          cosmlib.SdkUDEToStakingUDE(u.Entries),
			},
		)
	}

	return unbondingDelegations, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetRedelegations implements the `getRedelegations(address,address)` method.
func (c *Contract) GetRedelegations(
	ctx context.Context,
	delegatorAddress common.Address,
	srcValidator common.Address,
	dstValidator common.Address,
	pagination any,
) ([]generated.IStakingModuleRedelegationEntry, cbindings.CosmosPageResponse, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.accAddrCodec, delegatorAddress)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	srcValAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), srcValidator)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}
	destValAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), dstValidator)
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	rsp, err := c.querier.Redelegations(
		ctx,
		&stakingtypes.QueryRedelegationsRequest{
			DelegatorAddr:    delAddr,
			SrcValidatorAddr: srcValAddr,
			DstValidatorAddr: destValAddr,
			Pagination:       cosmlib.ExtractPageRequestFromInput(pagination),
		},
	)
	if status.Code(err) == codes.NotFound {
		return []generated.IStakingModuleRedelegationEntry{}, cbindings.CosmosPageResponse{}, nil
	} else if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	var redelegationEntryResponses []stakingtypes.RedelegationEntryResponse
	for _, r := range rsp.GetRedelegationResponses() {
		redel := r.GetRedelegation()
		if redel.DelegatorAddress == delAddr &&
			redel.ValidatorSrcAddress == srcValAddr &&
			redel.ValidatorDstAddress == destValAddr {
			redelegationEntryResponses = r.GetEntries()
			break
		}
	}
	redelegationEntries := make(
		[]stakingtypes.RedelegationEntry, 0, len(redelegationEntryResponses),
	)
	for _, entryRsp := range redelegationEntryResponses {
		redelegationEntries = append(redelegationEntries, entryRsp.GetRedelegationEntry())
	}

	return cosmlib.SdkREToStakingRE(redelegationEntries),
		cosmlib.SdkPageResponseToEvmPageResponse(rsp.Pagination),
		err
}

// Delegate implements the `delegate(address,uint256)` method.
func (c *Contract) Delegate(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
) (bool, error) {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.accAddrCodec, pvm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Delegate(ctx, stakingtypes.NewMsgDelegate(
		caller,
		valAddr,
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// Undelegate implements the `undelegate(address,uint256)` method.
func (c *Contract) Undelegate(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
) (bool, error) {
	denom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.accAddrCodec, pvm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Undelegate(ctx, stakingtypes.NewMsgUndelegate(
		caller,
		valAddr,
		sdk.Coin{Denom: denom, Amount: sdkmath.NewIntFromBigInt(amount)},
	))
	return err == nil, err
}

// BeginRedelegate implements the `beginRedelegate(address,address,uint256)` method.
func (c *Contract) BeginRedelegate(
	ctx context.Context,
	srcValidator common.Address,
	dstValidator common.Address,
	amount *big.Int,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}
	srcValAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), srcValidator)
	if err != nil {
		return false, err
	}
	destValAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), dstValidator)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.accAddrCodec, pvm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.BeginRedelegate(
		ctx,
		stakingtypes.NewMsgBeginRedelegate(
			caller,
			srcValAddr,
			destValAddr,
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err == nil, err
}

// CancelRedelegate implements the `cancelRedelegate(address,address,uint256,int64)` method.
func (c *Contract) CancelUnbondingDelegation(
	ctx context.Context,
	validatorAddress common.Address,
	amount *big.Int,
	creationHeight int64,
) (bool, error) {
	bondDenom, err := c.bondDenom(ctx)
	if err != nil {
		return false, err
	}
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validatorAddress)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.accAddrCodec, pvm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.CancelUnbondingDelegation(
		ctx,
		stakingtypes.NewMsgCancelUnbondingDelegation(
			caller,
			valAddr,
			creationHeight,
			sdk.Coin{Denom: bondDenom, Amount: sdkmath.NewIntFromBigInt(amount)},
		),
	)
	return err != nil, err
}

// bondDenom returns the bond denom from the staking module.
func (c *Contract) bondDenom(ctx context.Context) (string, error) {
	res, err := c.querier.Params(ctx, &stakingtypes.QueryParamsRequest{})
	if err != nil {
		return "", err
	}

	return res.Params.BondDenom, nil
}

// ConvertValAddressFromString converts a Cosmos string representing a validator address to a
// common.Address.
func (c *Contract) ConvertValAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.vs.ValidatorAddressCodec(), attributeValue)
}

// ConvertAccAddressFromString converts a Cosmos string representing a account address to a
// common.Address.
func (c *Contract) ConvertAccAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.accAddrCodec, attributeValue)
}
