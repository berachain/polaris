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

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"pkg.berachain.dev/polaris/cosmos/precompile/contracts/solidity/generated"
)

// `submitProposalHelper` is a helper function for the `SubmitProposal` method of the governance precompile contract.
func (c *Contract) submitProposalHelper(
	ctx context.Context,
	messages []*codectypes.Any,
	initialDeposit []generated.Coin,
	proposer sdk.AccAddress,
	metadata, title, summary string,
	expedited bool,
) ([]any, error) {
	coins := []sdk.Coin{}

	// Convert the initial deposit to sdk.Coin.
	for _, coin := range initialDeposit {
		coins = append(coins, sdk.NewCoin(coin.Denom, sdk.NewIntFromBigInt(coin.Amount)))
	}

	res, err := c.msgServer.SubmitProposal(ctx, &v1.MsgSubmitProposal{
		Messages:       messages,
		InitialDeposit: coins,
		Proposer:       proposer.String(),
		Metadata:       metadata,
		Title:          title,
		Summary:        summary,
		Expedited:      expedited,
	})
	if err != nil {
		return nil, err
	}

	return []any{res.ProposalId}, nil
}

// `cancelProposalHelper` is a helper function for the `CancelProposal` method of the governance precompile contract.
func (c *Contract) cancelProposalHelper(
	ctx context.Context,
	proposer sdk.AccAddress,
	proposalID uint64,
) ([]any, error) {
	res, err := c.msgServer.CancelProposal(ctx, &v1.MsgCancelProposal{
		ProposalId: proposalID,
		Proposer:   proposer.String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{uint64(res.CanceledTime.Unix()), res.CanceledHeight}, nil
}
