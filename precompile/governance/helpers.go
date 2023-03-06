package governance

import (
	"context"
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
)

// `submitProposalHelper` is a helper function for the `SubmitProposal` method of the governance precompile contract.
func (c *Contract) submitProposalHelper(
	ctx context.Context,
	messages []*codectypes.Any,
	initialDeposit []generated.IGovernanceModuleCoin,
	proposer sdk.AccAddress,
	metadata, title, summary string,
	expedited bool,
) ([]any, error) {
	var coins []sdk.Coin

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

	return []any{big.NewInt(int64(res.ProposalId))}, nil
}

// `cancelProposalHelper` is a helper function for the `CancelProposal` method of the governance precompile contract.
func (c *Contract) cancelProposalHelper(
	ctx context.Context,
	proposer sdk.AccAddress,
	proposalID *big.Int,
) ([]any, error) {
	res, err := c.msgServer.CancelProposal(ctx, &v1.MsgCancelProposal{
		ProposalId: proposalID.Uint64(),
		Proposer:   proposer.String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{big.NewInt(int64(res.ProposalId))}, nil
}

// `execLegacyContentHelper` is a helper function for the `ExecLegacyContent` method of the governance precompile contract.
func (c *Contract) execLegacyContentHelper(ctx context.Context, content *codectypes.Any, authority string) ([]any, error) {
	_, err := c.msgServer.ExecLegacyContent(ctx, &v1.MsgExecLegacyContent{
		Content:   content,
		Authority: authority,
	})
	if err != nil {
		return nil, err
	}
	return []any{}, nil
}

// `voteHelper` is a helper function for the `Vote` method of the governance precompile contract.
func (c *Contract) voteHelper(
	ctx context.Context,
	voter sdk.AccAddress,
	proposalId *big.Int,
	option int32,
	metadata string,
) ([]any, error) {
	_, err := c.msgServer.Vote(ctx, &v1.MsgVote{
		ProposalId: proposalId.Uint64(),
		Voter:      voter.String(),
		Option:     v1.VoteOption(option),
		Metadata:   metadata,
	})
	if err != nil {
		return nil, err
	}
	return []any{}, nil
}

// `voteWeighted` is a helper function for the `VoteWeighted` method of the governance precompile contract.
func (c *Contract) voteWeightedHelper(
	ctx context.Context,
	voter sdk.AccAddress,
	proposalId *big.Int,
	options []generated.IGovernanceModuleWeightedVoteOption,
	metadata string,
) ([]any, error) {
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
			ProposalId: proposalId.Uint64(),
			Voter:      voter.String(),
			Options:    msgOptions,
			Metadata:   metadata,
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}
