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
	"reflect"

	"cosmossdk.io/core/address"
	sdkmath "cosmossdk.io/math"

	libgenerated "github.com/berachain/polaris/contracts/bindings/cosmos/lib"
	"github.com/berachain/polaris/contracts/bindings/cosmos/precompile/governance"
	"github.com/berachain/polaris/contracts/bindings/cosmos/precompile/staking"
	"github.com/berachain/polaris/cosmos/precompile"
	"github.com/berachain/polaris/lib/utils"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"
)

/**
 * This file contains conversions between native Cosmos SDK types and go-ethereum ABI types.
 */

// SdkCoinsToEvmCoins converts sdk.Coins into []libgenerated.CosmosCoin.
func SdkCoinsToEvmCoins(sdkCoins sdk.Coins) []libgenerated.CosmosCoin {
	evmCoins := make([]libgenerated.CosmosCoin, len(sdkCoins))
	for i, coin := range sdkCoins {
		evmCoins[i] = SdkCoinToEvmCoin(coin)
	}
	return evmCoins
}

// SdkCoinToEvmCoin converts sdk.Coin into libgenerated.CosmosCoin.
func SdkCoinToEvmCoin(coin sdk.Coin) libgenerated.CosmosCoin {
	evmCoin := libgenerated.CosmosCoin{
		Amount: coin.Amount.BigInt(),
		Denom:  coin.Denom,
	}
	return evmCoin
}

func SdkPageResponseToEvmPageResponse(
	pageResponse *query.PageResponse,
) libgenerated.CosmosPageResponse {
	if pageResponse == nil {
		return libgenerated.CosmosPageResponse{}
	}
	return libgenerated.CosmosPageResponse{
		NextKey: string(pageResponse.GetNextKey()),
		Total:   pageResponse.GetTotal(),
	}
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

	sdkCoins := sdk.Coins{}
	for _, evmCoin := range amounts {
		sdkCoin := sdk.Coin{
			Denom: evmCoin.Denom, Amount: sdkmath.NewIntFromBigInt(evmCoin.Amount),
		}
		if !sdkCoin.IsZero() {
			// remove any 0 amounts
			sdkCoins = append(sdkCoins, sdkCoin)
		}
	}
	if len(sdkCoins) == 0 {
		return nil, precompile.ErrInvalidCoin
	}

	// sort the coins by denom, as Cosmos expects
	return sdkCoins.Sort(), nil
}

func ExtractPageRequestFromInput(pageRequest any) *query.PageRequest {
	// note: we have to use unnamed struct here, otherwise the compiler cannot cast
	// the any type input into the contract's generated type.
	pageReq, ok := utils.GetAs[struct {
		Key        string `json:"key"`
		Offset     uint64 `json:"offset"`
		Limit      uint64 `json:"limit"`
		CountTotal bool   `json:"count_total"`
		Reverse    bool   `json:"reverse"`
	}](pageRequest)
	if !ok {
		return nil
	}

	return &query.PageRequest{
		Key:        []byte(pageReq.Key),
		Offset:     pageReq.Offset,
		Limit:      pageReq.Limit,
		CountTotal: pageReq.CountTotal,
		Reverse:    pageReq.Reverse,
	}
}

// ExtractCoinFromInputToCoin converts a coin from input (of type any) into sdk.Coins.
func ExtractCoinFromInputToCoin(coin any) (sdk.Coin, error) {
	// note: we have to use unnamed struct here, otherwise the compiler cannot cast
	// the any type input into IBankModuleCoin.
	amounts, ok := utils.GetAs[struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}](coin)
	if !ok {
		return sdk.Coin{}, precompile.ErrInvalidCoin
	}

	sdkCoin := sdk.Coin{
		Denom:  amounts.Denom,
		Amount: sdkmath.NewIntFromBigInt(amounts.Amount),
	}
	if err := sdkCoin.Validate(); err != nil {
		return sdk.Coin{}, err
	}
	return sdkCoin, nil
}

// SdkUDEToStakingUDE converts a Cosmos SDK Unbonding Delegation Entry list to a geth compatible
// list of Unbonding Delegation Entries.
func SdkUDEToStakingUDE(
	ude []stakingtypes.UnbondingDelegationEntry,
) []staking.IStakingModuleUnbondingDelegationEntry {
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
func SdkREToStakingRE(
	re []stakingtypes.RedelegationEntry,
) []staking.IStakingModuleRedelegationEntry {
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
func SdkValidatorsToStakingValidators(
	valAddrCodec address.Codec,
	vals []stakingtypes.Validator,
) ([]staking.IStakingModuleValidator, error) {
	valsOut := make([]staking.IStakingModuleValidator, len(vals))
	for i, val := range vals {
		operEthAddr, err := EthAddressFromString(valAddrCodec, val.OperatorAddress)
		if err != nil {
			return nil, err
		}
		pubKey, err := val.ConsPubKey()
		if err != nil {
			return nil, err
		}
		valsOut[i] = staking.IStakingModuleValidator{
			OperatorAddr:    operEthAddr,
			ConsAddr:        pubKey.Address(),
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

// SdkProposalToGovProposal is a helper function to transform a `v1.Proposal` to an
// `IGovernanceModule.Proposal`.
func SdkProposalToGovProposal(
	proposal *v1.Proposal, addressCodec address.Codec,
) (governance.IGovernanceModuleProposal, error) {
	messages := make([]governance.CosmosCodecAny, len(proposal.Messages))
	for i, msg := range proposal.Messages {
		messages[i] = governance.CosmosCodecAny{
			Value:   msg.Value,
			TypeURL: msg.TypeUrl,
		}
	}

	totalDeposit := make([]governance.CosmosCoin, len(proposal.TotalDeposit))
	for i, coin := range proposal.TotalDeposit {
		totalDeposit[i] = governance.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		}
	}

	proposerAddr, err := EthAddressFromString(addressCodec, proposal.Proposer)
	if err != nil {
		return governance.IGovernanceModuleProposal{}, err
	}

	output := governance.IGovernanceModuleProposal{
		Id:       proposal.Id,
		Messages: messages,
		Status:   int32(proposal.Status), // Status is an alias for int32.
		FinalTallyResult: governance.IGovernanceModuleTallyResult{
			YesCount:        proposal.FinalTallyResult.YesCount,
			AbstainCount:    proposal.FinalTallyResult.AbstainCount,
			NoCount:         proposal.FinalTallyResult.NoCount,
			NoWithVetoCount: proposal.FinalTallyResult.NoWithVetoCount,
		},

		TotalDeposit: totalDeposit,
		Metadata:     proposal.Metadata,
		Title:        proposal.Title,
		Summary:      proposal.Summary,
		Proposer:     proposerAddr,
	}

	if proposal.SubmitTime != nil {
		output.SubmitTime = uint64(proposal.SubmitTime.Unix())
	}

	if proposal.DepositEndTime != nil {
		output.DepositEndTime = uint64(proposal.DepositEndTime.Unix())
	}

	if proposal.VotingStartTime != nil {
		output.VotingStartTime = uint64(proposal.VotingStartTime.Unix())
	}

	if proposal.VotingEndTime != nil {
		output.VotingEndTime = uint64(proposal.VotingEndTime.Unix())
	}

	return output, nil
}

// ConvertMsgSubmitProposalToSdk is a helper function to convert a `MsgSubmitProposal` to the gov
// `v1.MsgSubmitProposal`.
func ConvertMsgSubmitProposalToSdk(
	prop any, ir codectypes.InterfaceRegistry, addressCodec address.Codec,
) (*v1.MsgSubmitProposal, error) {
	// Convert the prop object to the desired unnamed struct using reflection.
	propValue := reflect.ValueOf(prop)

	// Build the proposal messages.
	proposalMsgs, err := decodeProposalMessages(propValue.FieldByName("Messages").Interface())
	if err != nil {
		return nil, err
	}
	messages := make([]*codectypes.Any, len(proposalMsgs))
	for i, genCodecAny := range proposalMsgs {
		messages[i] = &codectypes.Any{
			Value:   genCodecAny.Value,
			TypeUrl: genCodecAny.TypeURL,
		}
		var msg sdk.Msg
		if err = ir.UnpackAny(messages[i], &msg); err != nil {
			return nil, err
		}
	}

	// Build the initial deposit.
	initDeposit, err := decodeInitialDeposit(propValue.FieldByName("InitialDeposit").Interface())
	if err != nil {
		return nil, err
	}
	initialDeposit := make(sdk.Coins, len(initDeposit))
	for i, coin := range initDeposit {
		initialDeposit[i] = sdk.Coin{
			Denom:  coin.Denom,
			Amount: sdkmath.NewIntFromBigInt(coin.Amount),
		}
	}

	// Return the v1.MsgSubmitProposal with all string fields attached.
	proposer, err := StringFromEthAddress(
		addressCodec, propValue.FieldByName("Proposer").Interface().(common.Address),
	)
	if err != nil {
		return nil, err
	}
	return &v1.MsgSubmitProposal{
		Messages:       messages,
		InitialDeposit: initialDeposit,
		Proposer:       proposer,
		Metadata:       propValue.FieldByName("Metadata").Interface().(string),
		Title:          propValue.FieldByName("Title").Interface().(string),
		Summary:        propValue.FieldByName("Summary").Interface().(string),
		Expedited:      propValue.FieldByName("Expedited").Interface().(bool),
	}, nil
}

// decodeProposalMessages is a helper function to convert the unnamed type of Messages into a
// usable Messages slice.
func decodeProposalMessages(propMsgs any) ([]struct {
	TypeURL string `json:"typeURL"`
	Value   []byte `json:"value"`
}, error) {
	var proposalMsgs []struct {
		TypeURL string `json:"typeURL"`
		Value   []byte `json:"value"`
	}

	switch reflect.TypeOf(propMsgs) {
	case reflect.TypeOf([]governance.CosmosCodecAny{}):
		for _, propMsg := range utils.MustGetAs[[]governance.CosmosCodecAny](propMsgs) {
			proposalMsgs = append(proposalMsgs, struct {
				TypeURL string `json:"typeURL"`
				Value   []byte `json:"value"`
			}{
				TypeURL: propMsg.TypeURL,
				Value:   propMsg.Value,
			})
		}
	case reflect.TypeOf([]struct {
		TypeURL string `json:"typeURL"`
		Value   []byte `json:"value"`
	}{}):
		proposalMsgs = utils.MustGetAs[[]struct {
			TypeURL string `json:"typeURL"`
			Value   []byte `json:"value"`
		}](propMsgs)
	default:
		return nil, precompile.ErrInvalidSubmitProposal
	}

	return proposalMsgs, nil
}

// decodeInitialDeposit is a helper function to convert the unnamed type of InitialDeposit into a
// usable coins slice.
func decodeInitialDeposit(initDep any) ([]struct {
	Amount *big.Int `json:"amount"`
	Denom  string   `json:"denom"`
}, error) {
	var initDeposit []struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}

	switch reflect.TypeOf(initDep) {
	case reflect.TypeOf([]governance.CosmosCoin{}):
		for _, coin := range utils.MustGetAs[[]governance.CosmosCoin](initDep) {
			initDeposit = append(initDeposit, struct {
				Amount *big.Int `json:"amount"`
				Denom  string   `json:"denom"`
			}{
				Amount: coin.Amount,
				Denom:  coin.Denom,
			})
		}
	case reflect.TypeOf([]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}{}):
		initDeposit = utils.MustGetAs[[]struct {
			Amount *big.Int `json:"amount"`
			Denom  string   `json:"denom"`
		}](initDep)
	default:
		return nil, precompile.ErrInvalidSubmitProposal
	}

	return initDeposit, nil
}
