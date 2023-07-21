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
	"pkg.berachain.dev/polaris/eth/core/vm"
)

const (
	EventTypeProposalSubmitted = "proposal_submitted"
)

// submitProposalHelper is a helper function for the `SubmitProposal` method of the
// governance precompile contract.
func (c *Contract) submitProposalHelper(
	ctx context.Context,
	proposalBz []byte,
	message []*codectypes.Any,
) (uint64, error) {
	polarCtx := vm.UnwrapPolarContext(ctx)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

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
	proposalID := strconv.FormatUint(res.ProposalId, 10)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeProposalSubmitted,
			sdk.NewAttribute(govtypes.AttributeKeyProposalID, proposalID),
			sdk.NewAttribute("proposal_sender", polarCtx.MsgSender().Hex()),
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
