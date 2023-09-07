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

	"cosmossdk.io/core/address"

	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/distribution"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile/staking"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the precompile contract for the distribution module.
type Contract struct {
	ethprecompile.BaseContract

	addressCodec address.Codec
	vs           staking.ValidatorStore
	msgServer    distributiontypes.MsgServer
	querier      distributiontypes.QueryServer
}

// NewPrecompileContract returns a new instance of the distribution module precompile contract.
func NewPrecompileContract(
	ak cosmlib.CodecProvider,
	vs staking.ValidatorStore,
	m distributiontypes.MsgServer,
	q distributiontypes.QueryServer,
) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.DistributionModuleMetaData.ABI,
			common.BytesToAddress([]byte{0x69}),
		),
		addressCodec: ak.AddressCodec(),
		vs:           vs,
		msgServer:    m,
		querier:      q,
	}
}

func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		distributiontypes.AttributeKeyValidator:       c.ConvertValAddressFromString,
		distributiontypes.AttributeKeyWithdrawAddress: c.ConvertAccAddressFromString,
	}
}

// SetWithdrawAddress is the precompile contract method for the `setWithdrawAddress(address)` method.
func (c *Contract) SetWithdrawAddress(
	ctx context.Context,
	withdrawAddress common.Address,
) (bool, error) {
	delAddr, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}
	withdrawAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, withdrawAddress)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.SetWithdrawAddress(ctx, &distributiontypes.MsgSetWithdrawAddress{
		DelegatorAddress: delAddr,
		WithdrawAddress:  withdrawAddr,
	})
	return err == nil, err
}

// GetWithdrawAddress is the precompile contract method for the `getWithdrawAddress(address)` method.
func (c *Contract) GetWithdrawAddress(
	ctx context.Context,
	delegator common.Address,
) (common.Address, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, delegator)
	if err != nil {
		return common.Address{}, err
	}

	resp, err := c.querier.DelegatorWithdrawAddress(
		ctx,
		&distributiontypes.QueryDelegatorWithdrawAddressRequest{
			DelegatorAddress: delAddr,
		},
	)
	if err != nil {
		return common.Address{}, err
	}

	withdrawAddr, err := cosmlib.EthAddressFromString(c.addressCodec, resp.WithdrawAddress)
	if err != nil {
		return common.Address{}, err
	}
	return withdrawAddr, nil
}

// GetWithdrawEnabled is the precompile contract method for the `getWithdrawEnabled()` method.
func (c *Contract) GetWithdrawEnabled(
	ctx context.Context,
) (bool, error) {
	res, err := c.querier.Params(ctx, &distributiontypes.QueryParamsRequest{})
	return res.Params.WithdrawAddrEnabled, err
}

// WithdrawDelegatorReward is the precompile contract method for the
// `withdrawDelegatorReward(address,address)` method.
func (c *Contract) WithdrawDelegatorReward(
	ctx context.Context,
	delegator common.Address,
	validator common.Address,
) ([]lib.CosmosCoin, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, delegator)
	if err != nil {
		return nil, err
	}
	valAddr, err := cosmlib.StringFromEthAddress(c.vs.ValidatorAddressCodec(), validator)
	if err != nil {
		return nil, err
	}

	res, err := c.msgServer.WithdrawDelegatorReward(ctx, &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
	})
	if err != nil {
		return nil, err
	}

	amount := make([]lib.CosmosCoin, 0)
	for _, coin := range res.Amount {
		amount = append(amount, lib.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return amount, nil
}

// GetDelegatorReward implements `getAllDelegatorRewards(address)`.
func (c *Contract) GetAllDelegatorRewards(
	ctx context.Context,
	delegator common.Address,
) ([]generated.IDistributionModuleValidatorReward, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, delegator)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.DelegationTotalRewards( // performs writes to the context kv stores
		ctx,
		&distributiontypes.QueryDelegationTotalRewardsRequest{
			DelegatorAddress: delAddr,
		},
	)
	if err != nil {
		return nil, err
	}

	rewards := make([]generated.IDistributionModuleValidatorReward, 0, len(res.Rewards))
	for _, reward := range res.Rewards {
		var amount []generated.CosmosCoin
		for _, coin := range reward.Reward {
			amount = append(amount, generated.CosmosCoin{
				Denom:  coin.Denom,
				Amount: coin.Amount.TruncateInt().BigInt(),
			})
		}
		var valAddr common.Address
		valAddr, err = cosmlib.EthAddressFromString(
			c.vs.ValidatorAddressCodec(), reward.ValidatorAddress,
		)
		if err != nil {
			return nil, err
		}
		rewards = append(rewards, generated.IDistributionModuleValidatorReward{
			Validator: valAddr,
			Rewards:   amount,
		})
	}
	return rewards, nil
}

// GetDelegatorReward implements `getTotalDelegatorReward(address)`.
func (c *Contract) GetTotalDelegatorReward(
	ctx context.Context,
	delegator common.Address,
) ([]lib.CosmosCoin, error) {
	delAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, delegator)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.DelegationTotalRewards( // performs writes to the context kv stores
		ctx,
		&distributiontypes.QueryDelegationTotalRewardsRequest{
			DelegatorAddress: delAddr,
		},
	)
	if err != nil {
		return nil, err
	}

	amount := make([]lib.CosmosCoin, 0, len(res.Total))
	for _, coin := range res.Total {
		amount = append(amount, lib.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.TruncateInt().BigInt(),
		})
	}
	return amount, nil
}

// ConvertValAddressFromBech32 converts a Cosmos string representing a validator address to a
// common.Address.
func (c *Contract) ConvertValAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.vs.ValidatorAddressCodec(), attributeValue)
}

// ConvertAccAddressFromString converts a Cosmos string representing a account address to a
// common.Address.
func (c *Contract) ConvertAccAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.addressCodec, attributeValue)
}
