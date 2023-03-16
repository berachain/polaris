package distribution

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	generated "pkg.berachain.dev/polaris/cosmos/precompile/contracts/solidity/generated"
)

// `setWithdrawAddressHelper` is a helper function for the `SetWithdrawAddress` method.
func (c *Contract) setWithdrawAddressHelper(ctx context.Context, delegator, withdrawer sdk.AccAddress) ([]any, error) {
	_, err := c.msgServer.SetWithdrawAddress(ctx, &distributiontypes.MsgSetWithdrawAddress{
		DelegatorAddress: delegator.String(),
		WithdrawAddress:  withdrawer.String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}

// `withdrawDelegatorRewards` is a helper function for the `WithdrawDelegatorRewards` method.
func (c *Contract) withdrawDelegatorRewardsHelper(ctx context.Context, delegator sdk.AccAddress, validator sdk.ValAddress) ([]any, error) {
	res, err := c.msgServer.WithdrawDelegatorReward(ctx, &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: delegator.String(),
		ValidatorAddress: validator.String(),
	})
	if err != nil {
		return nil, err
	}

	amount := make([]generated.IBankModuleCoin, 0)
	for _, coin := range res.Amount {
		amount = append(amount, generated.IBankModuleCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Uint64(),
		})
	}

	return []any{amount}, nil
}
