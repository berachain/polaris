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

	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// querier defines the required functions from the governance module's querier.
type querier interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(context.Context, *govtypesv1.QueryProposalRequest) (*govtypesv1.QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(context.Context, *govtypesv1.QueryProposalsRequest) (*govtypesv1.QueryProposalsResponse, error)
}

// msgServer defines the required functions from the governance module's msg server.
type msgServer interface {
	// CancelProposal defines a method to cancel governance proposal. Since: cosmos-sdk 0.48
	CancelProposal(context.Context, *govtypesv1.MsgCancelProposal) (*govtypesv1.MsgCancelProposalResponse, error)
	// SubmitProposal defines a method to create new proposal given the messages.
	SubmitProposal(context.Context, *govtypesv1.MsgSubmitProposal) (*govtypesv1.MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(context.Context, *govtypesv1.MsgVote) (*govtypesv1.MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	VoteWeighted(context.Context, *govtypesv1.MsgVoteWeighted) (*govtypesv1.MsgVoteWeightedResponse, error)
}
