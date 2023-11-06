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

package miner

import (
	evidence "cosmossdk.io/x/evidence/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govbeta "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	// DefaultAllowedMsgs are messages that can be submitted by external users.
	DefaultAllowedMsgs = map[string]sdk.Msg{
		// crisis
		"cosmos.crisis.v1beta1.MsgVerifyInvariant":         &crisis.MsgVerifyInvariant{},
		"cosmos.crisis.v1beta1.MsgVerifyInvariantResponse": nil,

		// evidence
		"cosmos.evidence.v1beta1.Equivocation":              nil,
		"cosmos.evidence.v1beta1.MsgSubmitEvidence":         &evidence.MsgSubmitEvidence{},
		"cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse": nil,

		// gov
		"cosmos.gov.v1.MsgDeposit":                   &gov.MsgDeposit{},
		"cosmos.gov.v1.MsgDepositResponse":           nil,
		"cosmos.gov.v1.MsgVote":                      &gov.MsgVote{},
		"cosmos.gov.v1.MsgVoteResponse":              nil,
		"cosmos.gov.v1.MsgVoteWeighted":              &gov.MsgVoteWeighted{},
		"cosmos.gov.v1.MsgVoteWeightedResponse":      nil,
		"cosmos.gov.v1beta1.MsgDeposit":              &govbeta.MsgDeposit{},
		"cosmos.gov.v1beta1.MsgDepositResponse":      nil,
		"cosmos.gov.v1beta1.MsgVote":                 &govbeta.MsgVote{},
		"cosmos.gov.v1beta1.MsgVoteResponse":         nil,
		"cosmos.gov.v1beta1.MsgVoteWeighted":         &govbeta.MsgVoteWeighted{},
		"cosmos.gov.v1beta1.MsgVoteWeightedResponse": nil,
		"cosmos.gov.v1beta1.TextProposal":            nil,

		// slashing
		"cosmos.slashing.v1beta1.MsgUnjail":         &slashing.MsgUnjail{},
		"cosmos.slashing.v1beta1.MsgUnjailResponse": nil,

		// staking
		"cosmos.staking.v1beta1.MsgCreateValidator":         &staking.MsgCreateValidator{},
		"cosmos.staking.v1beta1.MsgCreateValidatorResponse": nil,
		"cosmos.staking.v1beta1.MsgEditValidator":           &staking.MsgEditValidator{},
		"cosmos.staking.v1beta1.MsgEditValidatorResponse":   nil,

		// tx
		"cosmos.tx.v1beta1.Tx": nil,

		// upgrade
		"cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal": nil,
		"cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":       nil,
	}
)
