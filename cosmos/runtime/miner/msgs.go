// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
