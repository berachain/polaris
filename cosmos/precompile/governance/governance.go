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

package governance

import (
	"context"
	"encoding/json"
	"strconv"

	"cosmossdk.io/core/address"

	cbindings "github.com/berachain/polaris/contracts/bindings/cosmos/lib"
	generated "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "github.com/berachain/polaris/cosmos/lib"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/precompile/log"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/vm"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EventTypeProposalSubmitted = `proposal_submitted`
	EventTypeProposalVoted     = `proposal_voted`
	AttributeProposalSender    = `proposal_sender`
	AttributeProposalVote      = `proposal_vote`
)

// Contract is the precompile contract for the governance module.
type Contract struct {
	ethprecompile.BaseContract

	addressCodec address.Codec
	msgServer    v1.MsgServer
	querier      v1.QueryServer
	ir           codectypes.InterfaceRegistry
}

// NewPrecompileContract creates a new precompile contract for the governance module.
func NewPrecompileContract(
	ak cosmlib.CodecProvider, m v1.MsgServer, q v1.QueryServer,
	ir codectypes.InterfaceRegistry) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.GovernanceModuleMetaData.ABI,
			// Precompile Address: 0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2
			common.BytesToAddress(authtypes.NewModuleAddress(govtypes.ModuleName)),
		),
		addressCodec: ak.AddressCodec(),
		msgServer:    m,
		querier:      q,
		ir:           ir,
	}
}

// CustomValueDecoders implements the `ethprecompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		AttributeProposalSender: log.ConvertCommonHexAddress,
		AttributeProposalVote:   ConvertStringToVote,
		sdk.AttributeKeySender:  c.ConvertAccAddressFromString,
	}
}

// SubmitProposal is the method for the `submitProposal` method of the
// governance precompile contract.
func (c *Contract) SubmitProposal(
	ctx context.Context,
	proposal any,
) (uint64, error) {
	// Convert the submit proposal msg into v1.MsgSubmitProposal.
	msgSubmitProposal, err := cosmlib.ConvertMsgSubmitProposalToSdk(proposal, c.ir, c.addressCodec)
	if err != nil {
		return 0, err
	}

	// Send it to the governance module.
	res, err := c.msgServer.SubmitProposal(ctx, msgSubmitProposal)
	if err != nil {
		return 0, err
	}

	// Emit an event at the end of this successful proposal submission.
	polarCtx := vm.UnwrapPolarContext(ctx)
	sdk.UnwrapSDKContext(polarCtx.Context()).EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeProposalSubmitted,
			sdk.NewAttribute(
				govtypes.AttributeKeyProposalID, strconv.FormatUint(res.ProposalId, 10)),
			sdk.NewAttribute(AttributeProposalSender, polarCtx.MsgSender().Hex()),
		),
	)

	// Return the proposal ID.
	return res.ProposalId, nil
}

// CancelProposal is the method for the `cancelProposal` method of the
// governance precompile contract.
func (c *Contract) CancelProposal(
	ctx context.Context,
	id uint64,
) (uint64, uint64, error) {
	caller, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return 0, 0, err
	}

	res, err := c.msgServer.CancelProposal(ctx, &v1.MsgCancelProposal{
		ProposalId: id,
		Proposer:   caller,
	})
	if err != nil {
		return 0, 0, err
	}

	return uint64(res.CanceledTime.Unix()), res.CanceledHeight, nil
}

// Vote is the method for the `vote` method of the governance precompile contract.
func (c *Contract) Vote(
	ctx context.Context,
	proposalID uint64,
	options int32,
	metadata string,
) (bool, error) {
	caller, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	// Submit the vote.
	if _, err = c.msgServer.Vote(ctx, &v1.MsgVote{
		ProposalId: proposalID,
		Voter:      caller,
		Option:     v1.VoteOption(options),
		Metadata:   metadata,
	}); err != nil {
		return false, err
	}

	// Emit the proposal voted event.
	voteBz, err := json.Marshal(generated.IGovernanceModuleVote{
		ProposalId: proposalID,
		Voter:      vm.UnwrapPolarContext(ctx).MsgSender(),
		Options: []generated.IGovernanceModuleWeightedVoteOption{
			{
				VoteOption: options,
				Weight:     "",
			},
		},
		Metadata: metadata,
	})
	if err != nil {
		return false, err
	}
	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeProposalVoted,
			sdk.NewAttribute(AttributeProposalVote, string(voteBz)),
		),
	)

	return true, nil
}

// VoteWeighted is the method for the `voteWeighted` method of the governance precompile contract.
func (c *Contract) VoteWeighted(
	ctx context.Context,
	proposalID uint64,
	options []generated.IGovernanceModuleWeightedVoteOption,
	metadata string,
) (bool, error) {
	// Convert the options to v1.WeightedVoteOption.
	msgOptions := make([]*v1.WeightedVoteOption, len(options))
	for i, option := range options {
		msgOptions[i] = &v1.WeightedVoteOption{
			Option: v1.VoteOption(option.VoteOption),
			Weight: option.Weight,
		}
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	// Submit the vote.
	if _, err = c.msgServer.VoteWeighted(
		ctx, &v1.MsgVoteWeighted{
			ProposalId: proposalID,
			Voter:      caller,
			Options:    msgOptions,
			Metadata:   metadata,
		},
	); err != nil {
		return false, err
	}

	// Emit the proposal voted event.
	voteBz, err := json.Marshal(generated.IGovernanceModuleVote{
		ProposalId: proposalID,
		Voter:      vm.UnwrapPolarContext(ctx).MsgSender(),
		Options:    options,
		Metadata:   metadata,
	})
	if err != nil {
		return false, err
	}
	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeProposalVoted,
			sdk.NewAttribute(AttributeProposalVote, string(voteBz)),
		),
	)

	return true, nil
}

// GetProposal is the method for the `getProposal` method of the governance precompile contract.
func (c *Contract) GetProposal(
	ctx context.Context,
	proposalID uint64,
) (generated.IGovernanceModuleProposal, error) {
	res, err := c.querier.Proposal(ctx, &v1.QueryProposalRequest{
		ProposalId: proposalID,
	})
	if err != nil {
		return generated.IGovernanceModuleProposal{}, err
	}

	return cosmlib.SdkProposalToGovProposal(res.Proposal, c.addressCodec)
}

// GetProposals is the method for the `getProposal`
//
//	method of the governance precompile contract.
func (c *Contract) GetProposals(
	ctx context.Context,
	proposalStatus int32,
	pagination any,
) ([]generated.IGovernanceModuleProposal, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.Proposals(ctx, &v1.QueryProposalsRequest{
		ProposalStatus: v1.ProposalStatus(proposalStatus),
		Pagination:     cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	govProposals := make([]generated.IGovernanceModuleProposal, len(res.Proposals))
	for i, sdkProposal := range res.Proposals {
		var govProposal generated.IGovernanceModuleProposal
		if govProposal, err = cosmlib.SdkProposalToGovProposal(sdkProposal, c.addressCodec); err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}
		govProposals[i] = govProposal
	}

	return govProposals, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetProposalDeposits is the method for the `getProposalDeposits`
// method of the governance precompile contract.
func (c *Contract) GetProposalDeposits(
	ctx context.Context,
	proposalID uint64,
) ([]generated.CosmosCoin, error) {
	res, err := c.querier.Proposal(ctx, &v1.QueryProposalRequest{
		ProposalId: proposalID,
	})
	if err != nil {
		return nil, err
	}
	deposits := make([]generated.CosmosCoin, 0)
	for _, deposit := range res.Proposal.TotalDeposit {
		deposits = append(deposits, generated.CosmosCoin{
			Denom:  deposit.Denom,
			Amount: deposit.Amount.BigInt(),
		})
	}

	return deposits, nil
}

// GetProposalDepositsByDepositor is the method for the `getProposalDepositsByDepositor` method
// of the governance precompile contract.
func (c *Contract) GetProposalDepositsByDepositor(
	ctx context.Context,
	proposalID uint64,
	depositor common.Address,
) ([]generated.CosmosCoin, error) {
	depositorBech32, err := cosmlib.StringFromEthAddress(c.addressCodec, depositor)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.Deposit(ctx, &v1.QueryDepositRequest{
		ProposalId: proposalID,
		Depositor:  depositorBech32,
	})
	if err != nil {
		return nil, err
	}

	deposits := make([]generated.CosmosCoin, 0)
	for _, deposit := range res.Deposit.Amount {
		deposits = append(deposits, generated.CosmosCoin{
			Denom:  deposit.Denom,
			Amount: deposit.Amount.BigInt(),
		})
	}

	return deposits, nil
}

// GetProposalVotes is the method for the `getProposalVotes` method of the
// governance precompile contract.
func (c *Contract) GetProposalTallyResult(
	ctx context.Context,
	proposalID uint64,
) (generated.IGovernanceModuleTallyResult, error) {
	res, err := c.querier.TallyResult(ctx, &v1.QueryTallyResultRequest{
		ProposalId: proposalID,
	})
	if err != nil {
		return generated.IGovernanceModuleTallyResult{}, err
	}

	return generated.IGovernanceModuleTallyResult{
		YesCount:        res.Tally.YesCount,
		AbstainCount:    res.Tally.AbstainCount,
		NoCount:         res.Tally.NoCount,
		NoWithVetoCount: res.Tally.NoWithVetoCount,
	}, nil
}

// GetProposalVotes is the method for the `getProposalVotes` method of the
// governance precompile contract.
func (c *Contract) GetProposalVotes(
	ctx context.Context,
	proposalID uint64,
	pagination any,
) ([]generated.IGovernanceModuleVote, cbindings.CosmosPageResponse, error) {
	res, err := c.querier.Votes(ctx, &v1.QueryVotesRequest{
		ProposalId: proposalID,
		Pagination: cosmlib.ExtractPageRequestFromInput(pagination),
	})
	if err != nil {
		return nil, cbindings.CosmosPageResponse{}, err
	}

	votes := make([]generated.IGovernanceModuleVote, 0)
	for _, vote := range res.Votes {
		voteOptions := make([]generated.IGovernanceModuleWeightedVoteOption, 0)
		for _, option := range vote.Options {
			voteOptions = append(
				voteOptions,
				generated.IGovernanceModuleWeightedVoteOption{
					VoteOption: int32(option.Option),
					Weight:     option.Weight,
				},
			)
		}
		var voter common.Address
		voter, err = cosmlib.EthAddressFromString(c.addressCodec, vote.Voter)
		if err != nil {
			return nil, cbindings.CosmosPageResponse{}, err
		}
		votes = append(votes, generated.IGovernanceModuleVote{
			ProposalId: proposalID,
			Voter:      voter,
			Options:    voteOptions,
			Metadata:   vote.Metadata,
		})
	}

	return votes, cosmlib.SdkPageResponseToEvmPageResponse(res.Pagination), nil
}

// GetProposalVotesByVoter is the method for the `getProposalVotesByVoter`
// method of the governance precompile contract.
func (c *Contract) GetProposalVotesByVoter(
	ctx context.Context,
	proposalID uint64,
	voter common.Address,
) (generated.IGovernanceModuleVote, error) {
	voterBech32, err := cosmlib.StringFromEthAddress(c.addressCodec, voter)
	if err != nil {
		return generated.IGovernanceModuleVote{}, err
	}

	res, err := c.querier.Vote(ctx, &v1.QueryVoteRequest{
		ProposalId: proposalID,
		Voter:      voterBech32,
	})
	if err != nil {
		return generated.IGovernanceModuleVote{}, err
	}

	voteOptions := make([]generated.IGovernanceModuleWeightedVoteOption, 0)
	for _, option := range res.Vote.Options {
		voteOptions = append(
			voteOptions,
			generated.IGovernanceModuleWeightedVoteOption{
				VoteOption: int32(option.Option),
				Weight:     option.Weight,
			},
		)
	}
	return generated.IGovernanceModuleVote{
		ProposalId: proposalID,
		Voter:      voter,
		Options:    voteOptions,
		Metadata:   res.Vote.Metadata,
	}, nil
}

// GetProposalVoteByVoter is the method for the `getProposalVoteByVoter`
// method of the governance precompile contract.
func (c *Contract) GetParams(
	ctx context.Context,
) (generated.IGovernanceModuleParams, error) {
	res, err := c.querier.Params(ctx, &v1.QueryParamsRequest{})
	if err != nil {
		return generated.IGovernanceModuleParams{}, err
	}

	minDeposit := make([]generated.CosmosCoin, 0)
	for _, coin := range res.Params.MinDeposit {
		minDeposit = append(minDeposit, generated.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}
	expeditedMinDeposit := make([]generated.CosmosCoin, 0)
	for _, coin := range res.Params.ExpeditedMinDeposit {
		expeditedMinDeposit = append(expeditedMinDeposit, generated.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return generated.IGovernanceModuleParams{
		MinDeposit:                 minDeposit,
		MaxDepositPeriod:           uint64(res.Params.MaxDepositPeriod.Abs()),
		VotingPeriod:               uint64(res.Params.VotingPeriod.Abs()),
		Quorum:                     res.Params.Quorum,
		Threshold:                  res.Params.Threshold,
		VetoThreshold:              res.Params.VetoThreshold,
		MinInitialDepositRatio:     res.Params.MinInitialDepositRatio,
		ProposalCancelRatio:        res.Params.ProposalCancelRatio,
		ProposalCancelDest:         res.Params.ProposalCancelDest,
		ExpeditedVotingPeriod:      uint64(res.Params.ExpeditedVotingPeriod.Abs()),
		ExpeditedThreshold:         res.Params.ExpeditedThreshold,
		ExpeditedMinDeposit:        expeditedMinDeposit,
		BurnVoteQuorum:             res.Params.BurnVoteQuorum,
		BurnProposalDepositPrevote: res.Params.BurnProposalDepositPrevote,
		BurnVoteVeto:               res.Params.BurnVoteVeto,
	}, nil
}

// GetDepositParams is the method for the `getDepositParams` method of the
// governance precompile contract.
func (c *Contract) GetDepositParams(
	ctx context.Context,
) (generated.IGovernanceModuleDepositParams, error) {
	res, err := c.querier.Params(ctx, &v1.QueryParamsRequest{})
	if err != nil {
		return generated.IGovernanceModuleDepositParams{}, err
	}

	minDeposit := make([]generated.CosmosCoin, 0)
	for _, coin := range res.Params.MinDeposit {
		minDeposit = append(minDeposit, generated.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return generated.IGovernanceModuleDepositParams{
		MinDeposit: minDeposit,
	}, nil
}

// GetVotingParams is the method for the `getVotingParams` method of the
// governance precompile contract.
func (c *Contract) GetVotingParams(
	ctx context.Context,
) (generated.IGovernanceModuleVotingParams, error) {
	res, err := c.querier.Params(ctx, &v1.QueryParamsRequest{})
	if err != nil {
		return generated.IGovernanceModuleVotingParams{}, err
	}

	return generated.IGovernanceModuleVotingParams{
		VotingPeriod: uint64(res.Params.VotingPeriod.Abs()),
	}, nil
}

// GetTallyParams is the method for the `getTallyParams` method of the
// governance precompile contract.
func (c *Contract) GetTallyParams(
	ctx context.Context,
) (generated.IGovernanceModuleTallyParams, error) {
	res, err := c.querier.Params(ctx, &v1.QueryParamsRequest{})
	if err != nil {
		return generated.IGovernanceModuleTallyParams{}, err
	}

	return generated.IGovernanceModuleTallyParams{
		Quorum:        res.Params.Quorum,
		Threshold:     res.Params.Threshold,
		VetoThreshold: res.Params.VetoThreshold,
	}, nil
}

// GetConstitution is the method for the `getConstitution` method of the
// governance precompile contract.
func (c *Contract) GetConstitution(
	ctx context.Context,
) (string, error) {
	res, err := c.querier.Constitution(ctx, &v1.QueryConstitutionRequest{})
	if err != nil {
		return "", err
	}

	return res.Constitution, nil
}

// ConvertStringToVote converts a string (json marshalled) Vote into a geth-binding Vote type.
func ConvertStringToVote(attributeValue string) (any, error) {
	var vote generated.IGovernanceModuleVote
	if err := json.Unmarshal([]byte(attributeValue), &vote); err != nil {
		return nil, err
	}
	return vote, nil
}

// ConvertAccAddressFromString converts a Cosmos string representing a account address to a
// common.Address.
func (c *Contract) ConvertAccAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.addressCodec, attributeValue)
}
