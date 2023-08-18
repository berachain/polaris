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

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the governance precompile contract.
type Contract struct {
	ethprecompile.BaseContract

	msgServer   v1.MsgServer
	queryServer v1.QueryServer
}

// NewPrecompileContract creates a new governance precompile contract.
func NewPrecompileContract(m v1.MsgServer, q v1.QueryServer) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.GovernanceModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(govtypes.ModuleName)),
		),
		msgServer:   m,
		queryServer: q,
	}
}

const (
	EventTypeProposalSubmitted = "proposal_submitted"
	AttributeProposalSender    = "proposal_sender"
)

// CustomValueDecoders implements the `ethprecompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		AttributeProposalSender: log.ConvertCommonHexAddress,
	}
}

// SubmitProposal is the method for the `submitProposal` function.
func (c *Contract) SubmitProposal(ctx context.Context, proposalMsg []byte) (uint64, error) {
	// Decode the proposal bytes into  v1.Proposal.
	var p v1.MsgSubmitProposal
	if err := p.Unmarshal(proposalMsg); err != nil {
		return 0, err
	}

	// Create the proposal.
	res, err := c.msgServer.SubmitProposal(ctx, &p)

	// Return the proposal ID.
	return res.ProposalId, err
}

// CancelProposal is the method for the `cancelProposal` function.
func (c *Contract) CancelProposal(ctx context.Context, id uint64) (uint64, uint64, error) {
	// Get the proposer from the context.
	proposer := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())

	// Cancel the vote.
	res, err := c.msgServer.CancelProposal(ctx, v1.NewMsgCancelProposal(id, proposer.String()))
	if err != nil {
		return 0, 0, err
	}

	return uint64(res.CanceledTime.Unix()), res.CanceledHeight, nil
}

// Vote is the method for the `vote` function.
func (c *Contract) Vote(ctx context.Context, id uint64, option int32, metadata string) (bool, error) {
	// Get the voter from the context.
	voter := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())

	// Push through the vote.
	_, err := c.msgServer.Vote(ctx, v1.NewMsgVote(voter, id, v1.VoteOption(option), metadata))

	return err == nil, err
}

// VoteWeighted is the method for the `voteWeighted` function.
func (c *Contract) VoteWeighted(
	ctx context.Context,
	id uint64,
	options []generated.IGovernanceModuleWeightedVoteOption,
	metadata string,
) (bool, error) {
	// Get the voter from the context.
	voter := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())

	// Decode the evm type into a cosmos vote options type.
	weightedOptions := []*v1.WeightedVoteOption{}
	for _, option := range options {
		weightedOptions = append(weightedOptions, &v1.WeightedVoteOption{
			Option: v1.VoteOption(option.VoteOption),
			Weight: option.Weight,
		})
	}

	// Push through the vote.
	_, err := c.msgServer.VoteWeighted(ctx, v1.NewMsgVoteWeighted(voter, id, weightedOptions, metadata))

	return err == nil, err
}

// GetProposal is the method for the `getProposal` function.
func (c *Contract) GetProposal(ctx context.Context, id uint64) (generated.IGovernanceModuleProposal, error) {
	// Get the proposal from the querier.
	res, err := c.queryServer.Proposal(ctx, &v1.QueryProposalRequest{ProposalId: id})
	if err != nil {
		return generated.IGovernanceModuleProposal{}, err
	}

	// Transform the cosmos type to an evm type and return.
	return cosmlib.SdkToEvmProposal(res.Proposal), nil
}

// GetProposals is the method for the `getProposals` function.
func (c *Contract) GetProposals(
	ctx context.Context, proposalStatus int32) ([]generated.IGovernanceModuleProposal, error) {
	// Get the proposals from the query server.
	res, err := c.queryServer.Proposals(
		ctx,
		&v1.QueryProposalsRequest{ProposalStatus: v1.ProposalStatus(proposalStatus)},
	)
	if err != nil {
		return nil, err
	}

	// Transform the cosmos type to an evm type and return.
	proposals := make([]generated.IGovernanceModuleProposal, 0, len(res.Proposals))
	for _, proposal := range res.Proposals {
		proposals = append(proposals, cosmlib.SdkToEvmProposal(proposal))
	}

	return proposals, nil
}
