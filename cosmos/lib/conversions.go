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

package lib

import (
	"math/big"
	"time"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	"pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/auth"
	"pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

/**
 * This file contains conversions between native Cosmos SDK types and go-ethereum ABI types.
 */

// SdkCoinsToEvmCoins converts sdk.Coins into []generated.CosmosCoin.
func SdkCoinsToEvmCoins(sdkCoins sdk.Coins) []generated.CosmosCoin {
	evmCoins := make([]generated.CosmosCoin, len(sdkCoins))
	for i, coin := range sdkCoins {
		evmCoins[i] = generated.CosmosCoin{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		}
	}
	return evmCoins
}

// ExtractCoinsFromInput converts coins from input (of type any) into sdk.Coins.
func ExtractCoinsFromInput(coins any) (sdk.Coins, error) {
	// note: we have to use unnamed struct here, otherwise the compiler cannot cast
	// the any type input into IBankModuleCoin.
	amounts, ok := utils.GetAs[[]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}](coins)
	if !ok {
		return nil, precompile.ErrInvalidCoin
	}

	sdkCoins := sdk.NewCoins()
	for _, evmCoin := range amounts {
		sdkCoins = append(sdkCoins, sdk.NewCoin(evmCoin.Denom, sdkmath.NewIntFromBigInt(evmCoin.Amount)))
	}
	// sort the coins by denom, as Cosmos expects.
	sdkCoins = sdkCoins.Sort()

	return sdkCoins, nil
}

// SdkCoinsToUnnamedCoins converts sdk.Coins into an unnamed struct.
func SdkCoinsToUnnamedCoins(coins sdk.Coins) any {
	unnamedCoins := []struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}{}
	for _, coin := range coins {
		unnamedCoins = append(unnamedCoins, struct {
			Amount *big.Int `json:"amount"`
			Denom  string   `json:"denom"`
		}{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		})
	}
	return unnamedCoins
}

// GetGrantAsSendAuth maps a list of grants to a list of send authorizations.
func GetGrantAsSendAuth(
	grants []*authz.Grant, blocktime time.Time,
) ([]*banktypes.SendAuthorization, error) {
	var sendAuths []*banktypes.SendAuthorization
	for _, grant := range grants {
		// Check that the expiration is still valid.
		if grant.Expiration == nil || grant.Expiration.After(blocktime) {
			sendAuth, ok := utils.GetAs[*banktypes.SendAuthorization](grant.Authorization.GetCachedValue())
			if !ok {
				return nil, precompile.ErrInvalidGrantType
			}
			sendAuths = append(sendAuths, sendAuth)
		}
	}
	return sendAuths, nil
}

// SdkUDEToStakingUDE converts a Cosmos SDK Unbonding Delegation Entry list to a geth compatible
// list of Unbonding Delegation Entries.
func SdkUDEToStakingUDE(ude []stakingtypes.UnbondingDelegationEntry) []staking.IStakingModuleUnbondingDelegationEntry {
	entries := make([]staking.IStakingModuleUnbondingDelegationEntry, len(ude))
	for i, entry := range ude {
		entries[i] = staking.IStakingModuleUnbondingDelegationEntry{
			CreationHeight: entry.CreationHeight,
			CompletionTime: entry.CompletionTime.String(),
			InitialBalance: entry.InitialBalance.BigInt(),
			Balance:        entry.Balance.BigInt(),
		}
	}
	return entries
}

// SdkREToStakingRE converts a Cosmos SDK Redelegation Entry list to a geth compatible list of
// Redelegation Entries.
func SdkREToStakingRE(re []stakingtypes.RedelegationEntry) []staking.IStakingModuleRedelegationEntry {
	entries := make([]staking.IStakingModuleRedelegationEntry, len(re))
	for i, entry := range re {
		entries[i] = staking.IStakingModuleRedelegationEntry{
			CreationHeight: entry.CreationHeight,
			CompletionTime: entry.CompletionTime.String(),
			InitialBalance: entry.InitialBalance.BigInt(),
			SharesDst:      entry.SharesDst.BigInt(),
		}
	}
	return entries
}

// SdkValidatorsToStakingValidators converts a Cosmos SDK Validator list to a geth compatible list
// of Validators.
func SdkValidatorsToStakingValidators(vals []stakingtypes.Validator) (
	[]staking.IStakingModuleValidator, error,
) {
	valsOut := make([]staking.IStakingModuleValidator, len(vals))
	for i, val := range vals {
		pubKey, err := val.ConsPubKey()
		if err != nil {
			return nil, err
		}
		valsOut[i] = staking.IStakingModuleValidator{
			OperatorAddress: val.OperatorAddress,
			ConsensusPubkey: pubKey.Bytes(),
			Jailed:          val.Jailed,
			Status:          val.Status.String(),
			Tokens:          val.Tokens.BigInt(),
			DelegatorShares: val.DelegatorShares.BigInt(),
			Description:     staking.IStakingModuleDescription(val.Description),
			UnbondingHeight: val.UnbondingHeight,
			UnbondingTime:   val.UnbondingTime.String(),
			Commission: staking.IStakingModuleCommission{
				CommissionRates: staking.IStakingModuleCommissionRates{
					Rate:          val.Commission.CommissionRates.Rate.BigInt(),
					MaxRate:       val.Commission.CommissionRates.MaxRate.BigInt(),
					MaxChangeRate: val.Commission.CommissionRates.MaxChangeRate.BigInt(),
				},
			},
			MinSelfDelegation:       val.MinSelfDelegation.BigInt(),
			UnbondingOnHoldRefCount: val.UnbondingOnHoldRefCount,
			UnbondingIds:            val.UnbondingIds,
		}
	}
	return valsOut, nil
}

// SdkAccountToAuthAccount converts a Cosmos SDK Base Account to a geth compatible Base Account.
func SdkAccountToAuthAccount(acc *authtypes.BaseAccount) auth.IAuthModuleBaseAccount {
	return auth.IAuthModuleBaseAccount{
		Addr:          EthAddressFromBech32(acc.Address),
		PubKey:        acc.GetPubKey().Bytes(),
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}
}
