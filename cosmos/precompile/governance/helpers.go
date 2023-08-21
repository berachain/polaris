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
	"fmt"
	"strconv"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

const (
	EventTypeProposalSubmitted = "proposal_submitted"
	AttributeProposalSender    = "proposal_sender"
)

// submitProposalHelper is a helper function for the `SubmitProposal` method of the
// governance precompile contract.
func (c *Contract) submitProposalHelper(
	ctx context.Context,
	proposalBz []byte,
	message []*codectypes.Any,
) (uint64, error) {
	// Decode the proposal.
	var proposal v1.MsgSubmitProposal
	if err := proposal.Unmarshal(proposalBz); err != nil {
		return 0, fmt.Errorf("failed to unmarshal proposal: %w", err)
	}
	proposal.Messages = message

	// Submit the message.
	res, err := c.msgServer.SubmitProposal(ctx, &proposal)
	if err != nil {
		return 0, err
	}

	// emit an event at the end of this successful proposal submission
	polarCtx := vm.UnwrapPolarContext(ctx)
	sdk.UnwrapSDKContext(polarCtx.Context()).EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeProposalSubmitted,
			sdk.NewAttribute(govtypes.AttributeKeyProposalID, strconv.FormatUint(res.ProposalId, 10)),
			sdk.NewAttribute(AttributeProposalSender, polarCtx.MsgSender().Hex()),
		),
	)

	return res.ProposalId, nil
}

// cancelProposalHelper is a helper function for the `CancelProposal` method of the
// governance precompile contract.
func (c *Contract) cancelProposalHelper(
	ctx context.Context,
	proposer sdk.AccAddress,
	proposalID uint64,
) (uint64, uint64, error) {
	res, err := c.msgServer.CancelProposal(ctx, &v1.MsgCancelProposal{
		ProposalId: proposalID,
		Proposer:   proposer.String(),
	})
	if err != nil {
		return 0, 0, err
	}

	return uint64(res.CanceledTime.Unix()), res.CanceledHeight, nil
}

// voteHelper is a helper function for the `Vote` method of the governance precompile contract.
func (c *Contract) voteHelper(
	ctx context.Context,
	voter sdk.AccAddress,
	proposalID uint64,
	option int32,
	metadata string,
) (bool, error) {
	_, err := c.msgServer.Vote(ctx, &v1.MsgVote{
		ProposalId: proposalID,
		Voter:      voter.String(),
		Option:     v1.VoteOption(option),
		Metadata:   metadata,
	})
	return err == nil, err
}

// voteWeighted is a helper function for the `VoteWeighted` method of the
// governance precompile contract.
func (c *Contract) voteWeightedHelper(
	ctx context.Context,
	voter sdk.AccAddress,
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

	_, err := c.msgServer.VoteWeighted(
		ctx, &v1.MsgVoteWeighted{
			ProposalId: proposalID,
			Voter:      voter.String(),
			Options:    msgOptions,
			Metadata:   metadata,
		},
	)
	return err == nil, err
}

// getProposalHelper is a helper function for the `GetProposal` method of the
// governance precompile contract.
func (c *Contract) getProposalHelper(
	ctx context.Context,
	proposalID uint64,
) (generated.IGovernanceModuleProposal, error) {
	res, err := c.querier.Proposal(ctx, &v1.QueryProposalRequest{
		ProposalId: proposalID,
	})
	if err != nil {
		return generated.IGovernanceModuleProposal{}, err
	}
	return transformProposalToABIProposal(*res.Proposal), nil
}

// getProposalsHelper is a helper function for the `GetProposal` method of the
// governance precompile contract.
func (c *Contract) getProposalsHelper(
	ctx context.Context,
	proposalStatus int32,
) ([]generated.IGovernanceModuleProposal, error) {
	res, err := c.querier.Proposals(ctx, &v1.QueryProposalsRequest{
		ProposalStatus: v1.ProposalStatus(proposalStatus),
	})
	if err != nil {
		return nil, err
	}

	proposals := make([]generated.IGovernanceModuleProposal, 0)
	for _, proposal := range res.Proposals {
		proposals = append(proposals, transformProposalToABIProposal(*proposal))
	}

	return proposals, nil
}

// getProposalDepositsHelper is a helper function for the `GetProposalDeposits` method of the
// governance precompile contract.
func (c *Contract) getProposalDepositsHelper(
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

// getProposalDepositsByDepositorHelper is a helper function for the `GetProposalDepositsByDepositor` method of the
// governance precompile contract.
func (c *Contract) getProposalDepositsByDepositorHelper(
	ctx context.Context,
	depositor common.Address,
	proposalID uint64,
) ([]generated.CosmosCoin, error) {
	res, err := c.querier.Deposit(ctx, &v1.QueryDepositRequest{
		ProposalId: proposalID,
		Depositor:  cosmlib.Bech32FromEthAddress(depositor),
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

// getProposalTallyResultHelper is a helper function for the `GetProposalTallyResult` method of the
// governance precompile contract.
func (c *Contract) getProposalTallyResultHelper(
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

// getProposalVotesHelper is a helper function for the `GetProposalVotes` method of the
// governance precompile contract.
func (c *Contract) getProposalVotesHelper(
	ctx context.Context,
	proposalID uint64,
	// pagination
) ([]generated.IGovernanceModuleVote, error) {
	res, err := c.querier.Votes(ctx, &v1.QueryVotesRequest{
		ProposalId: proposalID,
		// Pagination: pagination,
	})
	if err != nil {
		return nil, err
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
		votes = append(votes, generated.IGovernanceModuleVote{
			ProposalId: proposalID,
			Voter:      cosmlib.EthAddressFromBech32(vote.Voter),
			Options:    voteOptions,
			Metadata:   vote.Metadata,
		})
	}

	return votes, nil
}

// getProposalVotesByVoterHelper is a helper function for the `GetProposalVotesByVoter` method of the
// governance precompile contract.
func (c *Contract) getProposalVotesByVoterHelper(
	ctx context.Context,
	proposalID uint64,
	voter common.Address,
) (generated.IGovernanceModuleVote, error) {
	res, err := c.querier.Vote(ctx, &v1.QueryVoteRequest{
		ProposalId: proposalID,
		Voter:      cosmlib.Bech32FromEthAddress(voter),
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

// getParamsHelper is a helper function for the `GetParams` method of the
// governance precompile contract.
func (c *Contract) getParamsHelper(
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

// getDepositParamsHelper is a helper function for the `GetDepositParams` method of the
// governance precompile contract.
func (c *Contract) getDepositParamsHelper(
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

// getVotingParamsHelper is a helper function for the `GetVotingParams` method of the
// governance precompile contract.
func (c *Contract) getVotingParamsHelper(
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

// getTallyParamsHelper is a helper function for the `GetTallyParams` method of the
// governance precompile contract.
func (c *Contract) getTallyParamsHelper(
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

// getConstitutionHelper is a helper function for the `GetConstitution` method of the
// governance precompile contract.
func (c *Contract) getConstitutionHelper(
	ctx context.Context,
) (string, error) {
	res, err := c.querier.Constitution(ctx, &v1.QueryConstitutionRequest{})
	if err != nil {
		return "", err
	}

	return res.Constitution, nil
}

// transformProposalToABIProposal is a helper function to transform a `v1.Proposal`
// to an `IGovernanceModule.Proposal`.
func transformProposalToABIProposal(proposal v1.Proposal) generated.IGovernanceModuleProposal {
	message := make([]byte, 0)
	for _, msg := range proposal.Messages {
		message = append(message, msg.Value...)
	}

	totalDeposit := make([]generated.CosmosCoin, 0)
	for _, coin := range proposal.TotalDeposit {
		totalDeposit = append(totalDeposit, generated.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return generated.IGovernanceModuleProposal{
		Id:      proposal.Id,
		Message: message,
		Status:  int32(proposal.Status), // Status is an alias for int32.
		FinalTallyResult: generated.IGovernanceModuleTallyResult{
			YesCount:        proposal.FinalTallyResult.YesCount,
			AbstainCount:    proposal.FinalTallyResult.AbstainCount,
			NoCount:         proposal.FinalTallyResult.NoCount,
			NoWithVetoCount: proposal.FinalTallyResult.NoWithVetoCount,
		},
		SubmitTime:     uint64(proposal.SubmitTime.Unix()),
		DepositEndTime: uint64(proposal.DepositEndTime.Unix()),
		TotalDeposit:   totalDeposit,
		Metadata:       proposal.Metadata,
		Title:          proposal.Title,
		Summary:        proposal.Summary,
		Proposer:       proposal.Proposer,
	}
}
